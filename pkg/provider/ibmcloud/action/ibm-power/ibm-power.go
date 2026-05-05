package ibmpower

import (
	"fmt"

	"github.com/mapt-oss/pulumi-ibmcloud/sdk/go/ibmcloud"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/redhat-developer/mapt/pkg/manager"
	mc "github.com/redhat-developer/mapt/pkg/manager/context"
	ibmcloudp "github.com/redhat-developer/mapt/pkg/provider/ibmcloud"
	icdata "github.com/redhat-developer/mapt/pkg/provider/ibmcloud/data"
	"github.com/redhat-developer/mapt/pkg/provider/util/output"
	"github.com/redhat-developer/mapt/pkg/util"
	"github.com/redhat-developer/mapt/pkg/util/logging"
	resourcesUtil "github.com/redhat-developer/mapt/pkg/util/resources"
)

const (
	stackIBMPowerVS      = "icpw"
	outputHost           = "alsHost"
	outputUsername       = "alsUsername"
	outputUserPrivateKey = "alsUserPrivatekey"

	imageRHEL9   = "RHEL9-SP6"
	defaultUser  = "root"

	// Standard large build-host sizing on an s922 frame with shared processors and tier1 SSD.
	instanceMemory      = 256.0
	instanceProcs       = 8.0
	instanceProcType    = "shared"
	instanceSysType     = "s922"
	instanceStorageType = "tier1"
)

type PWArgs struct {
	Prefix      string
	NetworkID   string
	WorkspaceID string
}

type pwRequest struct {
	mCtx        *mc.Context
	prefix      *string
	networkID   string
	workspaceID string
}

// New provisions a Power VS (ppc64) instance inside an existing workspace and
// network. Both NetworkID and WorkspaceID are required.
func New(ctx *mc.ContextArgs, args *PWArgs) error {
	if args.NetworkID == "" || args.WorkspaceID == "" {
		return fmt.Errorf("--network-id and --workspace-id are required")
	}

	ibmcloudProvider := ibmcloudp.Provider()
	mCtx, err := mc.Init(ctx, ibmcloudProvider)
	if err != nil {
		return err
	}

	prefix := util.If(len(args.Prefix) > 0, args.Prefix, "main")
	r := &pwRequest{
		mCtx:        mCtx,
		prefix:      &prefix,
		networkID:   args.NetworkID,
		workspaceID: args.WorkspaceID,
	}
	cs := manager.Stack{
		StackName:           mCtx.StackNameByProject(stackIBMPowerVS),
		ProjectName:         mCtx.ProjectName(),
		BackedURL:           mCtx.BackedURL(),
		ProviderCredentials: ibmcloudp.DefaultCredentials,
		DeployFunc:          r.deploy,
	}
	sr, err := manager.UpStack(r.mCtx, cs)
	if err != nil {
		return fmt.Errorf("stack creation failed: %w", err)
	}
	return manageResults(mCtx, sr, prefix)
}

// Destroy tears down the Power VS stack identified by mCtxArgs.
func Destroy(mCtxArgs *mc.ContextArgs) (err error) {
	mCtx, err := mc.Init(mCtxArgs, ibmcloudp.Provider())
	if err != nil {
		return err
	}
	return ibmcloudp.Destroy(mCtx, stackIBMPowerVS)
}

func (r *pwRequest) deploy(ctx *pulumi.Context) error {
	pk, pki, err := piKey(ctx, r.mCtx, *r.prefix, stackIBMPowerVS, pulumi.String(r.workspaceID))
	if err != nil {
		return err
	}
	ctx.Export(fmt.Sprintf("%s-%s", *r.prefix, outputUserPrivateKey), pk.PrivateKeyPem)
	imageId, err := icdata.GetImage(r.mCtx,
		&icdata.PiImageArgs{
			CloudInstanceId: r.workspaceID,
			Name:            imageRHEL9,
		})
	if err != nil {
		return err
	}
	i, err := ibmcloud.NewPiInstance(ctx,
		resourcesUtil.GetResourceName(*r.prefix, stackIBMPowerVS, "pii"),
		&ibmcloud.PiInstanceArgs{
			PiInstanceName:    pulumi.String(r.mCtx.ProjectName()),
			PiMemory:          pulumi.Float64(instanceMemory),
			PiProcessors:      pulumi.Float64(instanceProcs),
			PiProcType:        pulumi.String(instanceProcType),
			PiSysType:         pulumi.String(instanceSysType),
			PiImageId:         pulumi.String(*imageId),
			PiHealthStatus:    pulumi.String("WARNING"),
			PiCloudInstanceId: pulumi.String(r.workspaceID),
			PiStorageType:     pulumi.String(instanceStorageType),
			PiKeyPairName:     pki.PiKeyName,
			PiNetworks: ibmcloud.PiInstancePiNetworkArray{
				&ibmcloud.PiInstancePiNetworkArgs{
					NetworkId: pulumi.String(r.networkID),
				},
			},
		})
	if err != nil {
		return err
	}
	ctx.Export(fmt.Sprintf("%s-%s", *r.prefix, outputUsername), pulumi.String(defaultUser))
	// Use ExternalIp when available (public network); fall back to IpAddress for private networks.
	ctx.Export(fmt.Sprintf("%s-%s", *r.prefix, outputHost),
		i.PiNetworks.ApplyT(func(networks []ibmcloud.PiInstancePiNetwork) (string, error) {
			if len(networks) == 0 {
				return "", fmt.Errorf("instance has no network interfaces")
			}
			if networks[0].ExternalIp != nil && *networks[0].ExternalIp != "" {
				return *networks[0].ExternalIp, nil
			}
			if networks[0].IpAddress != nil && *networks[0].IpAddress != "" {
				return *networks[0].IpAddress, nil
			}
			return "", fmt.Errorf("instance network has no IP address")
		}).(pulumi.StringOutput))
	return nil
}

// piKey creates a 4096-bit RSA TLS key pair and registers the public key as a
// Power VS SSH key in the given workspace. The private key PEM is available
// via the returned PrivateKey for export.
func piKey(ctx *pulumi.Context, mCtx *mc.Context, prefix, cId string, cloudInstanceID pulumi.StringInput) (*tls.PrivateKey, *ibmcloud.PiKey, error) {
	pk, err := tls.NewPrivateKey(
		ctx,
		resourcesUtil.GetResourceName(prefix, cId, "pk"),
		&tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA"),
			RsaBits:   pulumi.Int(4096),
		})
	if err != nil {
		return nil, nil, err
	}
	if mCtx.Debug() {
		pk.PrivateKeyPem.ApplyT(
			func(privateKey string) error {
				logging.Debugf("%s", privateKey)
				return nil
			})
	}
	pik, err := ibmcloud.NewPiKey(ctx,
		resourcesUtil.GetResourceName(prefix, cId, "pik"),
		&ibmcloud.PiKeyArgs{
			PiKeyName:         pulumi.String(mCtx.ProjectName()),
			PiCloudInstanceId: cloudInstanceID,
			PiSshKey:          pk.PublicKeyOpenssh,
		})
	return pk, pik, err
}

func manageResults(mCtx *mc.Context, stackResult auto.UpResult, prefix string) error {
	return output.Write(stackResult, mCtx.GetResultsOutputPath(), map[string]string{
		fmt.Sprintf("%s-%s", prefix, outputUsername):       "username",
		fmt.Sprintf("%s-%s", prefix, outputUserPrivateKey): "id_rsa",
		fmt.Sprintf("%s-%s", prefix, outputHost):           "host",
	})
}
