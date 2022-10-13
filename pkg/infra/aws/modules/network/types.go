package network

import (
	"github.com/adrianriobo/qenvs/pkg/infra/aws/services/vpc/subnet"
	vpc "github.com/adrianriobo/qenvs/pkg/infra/aws/services/vpc/vpc"
)

type NetworkRequest struct {
	CIDR                string
	Name                string
	AvailabilityZones   []string
	PublicSubnetsCIDRs  []string
	PrivateSubnetsCIDRs []string
	IntraSubnetsCIDRs   []string
	SingleNatGateway    bool
}

type NetworkResources struct {
	VPCResources       *vpc.VPCResources
	PublicSNResources  []*subnet.PublicSubnetResources
	PrivateSNResources []*subnet.PrivateSubnetResources
	IntraSNResources   []*subnet.PrivateSubnetResources
}