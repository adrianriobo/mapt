package supportmatrix

import (
	"fmt"
)

var (
	OL_RHEL = SupportedHost{
		ID:                 olRHELID,
		Description:        "rhel machine supporting initialize / build openshift local",
		Type:               RHEL,
		InstaceTypes:       []string{"c5.metal", "c5d.metal", "c5n.metal"},
		ProductDescription: "Red Hat Enterprise Linux",
		Spot:               true,
		AMI: AMI{
			// https://access.redhat.com/solutions/15356
			// Pattern with composition %s is major rhel version
			RegexPattern: "RHEL-%s*-x86_64-*",
			DefaultUser:  "ec2-user",
		},
	}

	OL_Windows = SupportedHost{
		ID:                 olWindowsID,
		Description:        "windows machine supporting nested virtualization (start openshift local)",
		Type:               Windows,
		InstaceTypes:       []string{"c5.metal", "c5d.metal", "c5n.metal"},
		ProductDescription: "Windows",
		Spot:               true,
	}

	G_MAC_M1 = SupportedHost{
		ID:           gMacOSM1ID,
		Description:  "mac machine with m1 chip arm64 arch",
		Type:         MacM1,
		InstaceTypes: []string{"mac2.metal"},
		Spot:         false,
		AMI: AMI{
			RegexName:   "amzn-ec2-macos-12*",
			DefaultUser: "ec2-user",
			Owner:       "628277914472",
			Filters: map[string]string{
				"architecture": "arm64_mac"},
		},
	}

	S_BASTION = SupportedHost{
		ID:           sBastionID,
		Description:  "bastion host to access hosts on private subnets",
		InstaceTypes: []string{"t2.small"},
		Spot:         false,
		AMI: AMI{
			RegexName:   "amzn-ami-hvm-*-x86_64-ebs",
			DefaultUser: "ec2-user",
		},
	}

	// https://github.com/ptcodes/proxy-server-with-terraform/blob/master/main.tf
	S_PROXY = SupportedHost{
		ID:           sProxyID,
		Description:  "proxy host to control network http access from hosts",
		InstaceTypes: []string{"t2.small"},
		Spot:         false,
		AMI: AMI{
			RegexName:   "amzn-ami-hvm-*-x86_64-ebs",
			DefaultUser: "ec2-user",
		},
	}
)

func GetHost(id string) (*SupportedHost, error) {
	switch id {
	case olRHELID:
		return &OL_RHEL, nil
	case olWindowsID:
		return &OL_Windows, nil
	case gMacOSM1ID:
		return &G_MAC_M1, nil
	}

	return nil, fmt.Errorf("supported host id is not valid")
}