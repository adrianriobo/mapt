package serverless

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/scheduler"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/awsx"
	awsxecs "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	maptContext "github.com/redhat-developer/mapt/pkg/manager/context"
	"github.com/redhat-developer/mapt/pkg/provider/aws/data"
	"github.com/redhat-developer/mapt/pkg/util"

	resourcesUtil "github.com/redhat-developer/mapt/pkg/util/resources"
)

type ServerlessManager struct{}

// Mapts requires the cluster to exist previously wit specific naming
// check hacks/aws/serverless.sh to
var (
	MaptServerlessDefaultPrefix = "mapt-serverless-manager"
)

// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#task_size
const (
	LimitCPU    = "2048"
	LimitMemory = "4096"
)

// func GetServerlessManager() *ServerlessManager {
// 	return &ServerlessManager{}
// }

var (
	ErrInvalidBackedURLForTimeout = fmt.Errorf("timeout can action can not be set due to backed url pointing to local file. Please use external storage or remote timeout option")
)

func CreateDestroyOperation(ctx *pulumi.Context,
	region, prefix, componentID string,
	target string,
	delay string) error {
	if err := checkBackedURLForServerless(); err != nil {
		return err
	}
	destroyCommand := fmt.Sprintf("aws %s destroy --project-name %s --backed-url %s --serverless",
		target,
		maptContext.ProjectName(),
		maptContext.BackedURL())
	return Create(ctx,
		region, prefix, componentID,
		destroyCommand, delay)
}

func checkBackedURLForServerless() error {
	return util.If(
		strings.HasPrefix(maptContext.BackedURL(), "file:///"),
		ErrInvalidBackedURLForTimeout,
		nil)
}

func Create(ctx *pulumi.Context,
	region, prefix, componentID string,
	command, delay string) error {
	// Get the pre configured cluster to handle serverless exectucions
	clusterArn, err := getClusterArn(ctx, region, prefix, componentID)
	if err != nil {
		return err
	}
	roleArn, err := getTaskRole(ctx, prefix, componentID)
	if err != nil {
		return err
	}
	limitCPUasInt, err := strconv.Atoi(LimitCPU)
	if err != nil {
		return err
	}
	limitMemoryasInt, err := strconv.Atoi(LimitMemory)
	if err != nil {
		return err
	}
	td, err := awsxecs.NewFargateTaskDefinition(ctx,
		resourcesUtil.GetResourceName(prefix, componentID, "fg-task"),
		&awsxecs.FargateTaskDefinitionArgs{
			Container: &awsxecs.TaskDefinitionContainerDefinitionArgs{
				Image:   pulumi.String(maptContext.OCI),
				Command: pulumi.ToStringArray(strings.Fields(command)),
				Cpu:     pulumi.Int(limitCPUasInt),
				Memory:  pulumi.Int(limitMemoryasInt),
			},
			Cpu:    pulumi.String(LimitCPU),
			Memory: pulumi.String(LimitMemory),
			ExecutionRole: &awsx.DefaultRoleWithPolicyArgs{
				RoleArn: roleArn,
			},
			TaskRole: &awsx.DefaultRoleWithPolicyArgs{
				RoleArn: roleArn,
			},
			LogGroup: &awsx.DefaultLogGroupArgs{
				Args: &awsx.LogGroupArgs{
					SkipDestroy:     pulumi.Bool(true),
					RetentionInDays: pulumi.Int(3),
				},
			},
		})
	if err != nil {
		return err
	}
	sRole, err := createSchedulerRole(ctx, prefix, componentID)
	if err != nil {
		return err
	}
	se, err := generateOneTimeScheduleExpression(region, delay)
	if err != nil {
		return err
	}
	subnetID, err := data.GetRandomPublicSubnet(region)
	if err != nil {
		return err
	}
	_, err = scheduler.NewSchedule(ctx,
		resourcesUtil.GetResourceName(prefix, componentID, "fgs"),
		&scheduler.ScheduleArgs{
			FlexibleTimeWindow: scheduler.ScheduleFlexibleTimeWindowArgs{
				Mode:                   scheduler.ScheduleFlexibleTimeWindowModeFlexible,
				MaximumWindowInMinutes: pulumi.Float64(1),
			},
			Target: scheduler.ScheduleTargetArgs{
				EcsParameters: scheduler.ScheduleEcsParametersArgs{
					TaskDefinitionArn: td.TaskDefinition.Arn(),
					LaunchType:        scheduler.ScheduleLaunchTypeFargate,
					NetworkConfiguration: scheduler.ScheduleNetworkConfigurationArgs{
						// https://github.com/aws/aws-cdk/issues/13348#issuecomment-1539336376
						AwsvpcConfiguration: scheduler.ScheduleAwsVpcConfigurationArgs{
							AssignPublicIp: scheduler.ScheduleAssignPublicIpEnabled,
							Subnets: pulumi.StringArray{
								pulumi.String(*subnetID),
							},
						},
					},
				},
				Arn:     clusterArn,
				RoleArn: sRole.Arn,
			},
			ScheduleExpression:         pulumi.String(se),
			ScheduleExpressionTimezone: pulumi.String(data.RegionTimezones[region]),
		})
	if err != nil {
		return err
	}
	return nil
}

// As part of the runtime for serverless invocation we need a fixed cluster spec on the region as so if
// it exists it will pick the cluster otherwise it will create and will not be deleted
func getClusterArn(ctx *pulumi.Context, region, prefix, componentID string) (*pulumi.StringOutput, error) {
	clusterName := fmt.Sprintf("%s-%s", MaptServerlessDefaultPrefix, "cluster")
	clusterArn, err := data.GetCluster(clusterName, region)
	if err != nil {
		if err == data.ErrECSClusterNotFound {
			if cluster, err := ecs.NewCluster(ctx,
				resourcesUtil.GetResourceName(prefix, componentID, "cluster"),
				&ecs.ClusterArgs{
					Tags: maptContext.ResourceTags(),
					Name: pulumi.String(clusterName),
				},
				pulumi.RetainOnDelete(true)); err != nil {
				return nil, err
			} else {
				return &cluster.Arn, nil
			}
		} else {
			return nil, fmt.Errorf("error getting cluster for serverless mode. %w", err)
		}
	}
	carn := pulumi.String(*clusterArn).ToStringOutput()
	return &carn, nil
}

// As part of the runtime for serverless invocation we need a fixed role for task execution the region as so if
// it exists it will pick the role otherwise it will create and will not be deleted
func getTaskRole(ctx *pulumi.Context, prefix, componentID string) (*pulumi.StringOutput, error) {
	roleName := fmt.Sprintf("%s-%s", MaptServerlessDefaultPrefix, "role")
	roleArn, err := data.GetRole(roleName)
	if err != nil {
		if role, err := createTaskRole(ctx, roleName, prefix, componentID); err != nil {
			return nil, err
		} else {
			return &role.Arn, nil
		}
	}
	rarn := pulumi.String(*roleArn).ToStringOutput()
	return &rarn, nil
}

// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html
// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/security-iam-roles.html
func createTaskRole(ctx *pulumi.Context, roleName, prefix, componentID string) (*iam.Role, error) {
	trustPolicyContent, err := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Effect": "Allow",
				"Principal": map[string]interface{}{
					"Service": "ecs-tasks.amazonaws.com",
				},
				"Action": "sts:AssumeRole",
			},
		},
	})
	if err != nil {
		return nil, err
	}
	// Need to creeate policies and attach
	r, err := iam.NewRole(ctx,
		resourcesUtil.GetResourceName(prefix, componentID, "role"),
		&iam.RoleArgs{
			Name:             pulumi.String(roleName),
			AssumeRolePolicy: pulumi.String(string(trustPolicyContent)),
			Tags:             maptContext.ResourceTags(),
		},
		pulumi.RetainOnDelete(true),
	)
	if err != nil {
		return nil, err
	}
	policyContent, err := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Effect": "Allow",
				"Action": []string{
					"s3:*",
					"ec2:*",
					"logs:*",
					"cloudformation:*",
					"scheduler:*",
				},
				"Resource": []string{
					"*",
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if _, err = iam.NewRolePolicy(ctx,
		resourcesUtil.GetResourceName(prefix, componentID, "ecs-role-policy"),
		&iam.RolePolicyArgs{
			Role:   r.ID(),
			Policy: pulumi.String(string(policyContent)),
		},
		pulumi.RetainOnDelete(true)); err != nil {
		return nil, err
	}
	return r, nil
}

// https://docs.aws.amazon.com/scheduler/latest/UserGuide/setting-up.html#setting-up-execution-role
func createSchedulerRole(ctx *pulumi.Context, prefix, componentID string) (*iam.Role, error) {
	trustPolicyContent, err := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Effect": "Allow",
				"Principal": map[string]interface{}{
					"Service": "scheduler.amazonaws.com",
				},
				"Action": "sts:AssumeRole",
			},
		},
	})
	if err != nil {
		return nil, err
	}
	// Need to creeate policies and attach
	r, err := iam.NewRole(ctx,
		resourcesUtil.GetResourceName(prefix, componentID, "role-sche"),
		&iam.RoleArgs{
			Name:             pulumi.String(fmt.Sprintf("mapt-sche-%s", maptContext.RunID())),
			AssumeRolePolicy: pulumi.String(string(trustPolicyContent)),
			Tags:             maptContext.ResourceTags(),
		})
	if err != nil {
		return nil, err
	}
	policyContent, err := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Effect": "Allow",
				"Action": []string{
					"s3:*",
					"ec2:*",
					"ecs:*",
					"iam:PassRole",
					"logs:*",
					"cloudformation:*",
					"scheduler:*",
				},
				"Resource": []string{
					"*",
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if _, err = iam.NewRolePolicy(ctx,
		resourcesUtil.GetResourceName(prefix, componentID, "sche-role-policy"),
		&iam.RolePolicyArgs{
			Role:   r.ID(),
			Policy: pulumi.String(string(policyContent)),
		}); err != nil {
		return nil, err
	}
	return r, nil
}

func generateOneTimeScheduleExpression(region, delay string) (string, error) {
	location, err := time.LoadLocation(data.RegionTimezones[region])
	if err != nil {
		log.Fatal("Failed to load timezone:", err)
	}
	// Get the current time in that timezone
	currentTime := time.Now().In(location)
	// Parse the timeout duration
	duration, err := time.ParseDuration(delay)
	if err != nil {
		return "", fmt.Errorf("invalid timeout format: %v", err)
	}
	// Add the duration to the current time
	futureTime := currentTime.Add(duration)
	return fmt.Sprintf("at(%s)", futureTime.Format("2006-01-02T15:04:05")), nil
}
