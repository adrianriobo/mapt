package data

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	cr "github.com/redhat-developer/mapt/pkg/provider/api/compute-request"
)

type ComputeSelector struct{}

func NewComputeSelector() *ComputeSelector { return &ComputeSelector{} }

func (c *ComputeSelector) Select(
	args *cr.ComputeRequestArgs) ([]string, error) {
	return machinesTypes(args)
}

func machinesTypes(args *cr.ComputeRequestArgs) ([]string, error) {
	machineTypesClient, err := compute.NewMachineTypesRESTClient(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create machineTypes client: %w", err)
	}
	defer machineTypesClient.Close()

	// reqList := &computepb.ListMachineTypesRequest{
	// 	Project: projectID,
	// 	Zone:    zone,
	// }

	return nil, fmt.Errorf("not implementedyet")
}
