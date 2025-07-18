package data

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/redhat-developer/mapt/pkg/provider/gcp"
)

// export GOOGLE_APPLICATION_CREDENTIALS="/home/user/keys/my-gcp-sa.json"
// export GOOGLE_CLOUD_PROJECT="my-gcp-project"

func GetZones() (zones []string, err error) {
	ctx := context.Background()
	client, err := compute.NewZonesRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	it := client.List(ctx,
		&computepb.ListZonesRequest{
			Project: gcp.GetProjectID(),
		})
	for {
		zone, err := it.Next()
		if err != nil {
			break
		}
		zones = append(zones, zone.GetName())
	}
	return
}
