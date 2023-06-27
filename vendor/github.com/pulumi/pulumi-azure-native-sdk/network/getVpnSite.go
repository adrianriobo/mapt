// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the details of a VPN site.
// API Version: 2020-11-01.
func LookupVpnSite(ctx *pulumi.Context, args *LookupVpnSiteArgs, opts ...pulumi.InvokeOption) (*LookupVpnSiteResult, error) {
	var rv LookupVpnSiteResult
	err := ctx.Invoke("azure-native:network:getVpnSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnSiteArgs struct {
	// The resource group name of the VpnSite.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VpnSite being retrieved.
	VpnSiteName string `pulumi:"vpnSiteName"`
}

// VpnSite Resource.
type LookupVpnSiteResult struct {
	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace *AddressSpaceResponse `pulumi:"addressSpace"`
	// The set of bgp properties.
	BgpProperties *BgpSettingsResponse `pulumi:"bgpProperties"`
	// The device properties.
	DeviceProperties *DevicePropertiesResponse `pulumi:"deviceProperties"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The ip-address for the vpn-site.
	IpAddress *string `pulumi:"ipAddress"`
	// IsSecuritySite flag.
	IsSecuritySite *bool `pulumi:"isSecuritySite"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Office365 Policy.
	O365Policy *O365PolicyPropertiesResponse `pulumi:"o365Policy"`
	// The provisioning state of the VPN site resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The key for vpn-site that can be used for connections.
	SiteKey *string `pulumi:"siteKey"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan *SubResourceResponse `pulumi:"virtualWan"`
	// List of all vpn site links.
	VpnSiteLinks []VpnSiteLinkResponse `pulumi:"vpnSiteLinks"`
}

func LookupVpnSiteOutput(ctx *pulumi.Context, args LookupVpnSiteOutputArgs, opts ...pulumi.InvokeOption) LookupVpnSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnSiteResult, error) {
			args := v.(LookupVpnSiteArgs)
			r, err := LookupVpnSite(ctx, &args, opts...)
			var s LookupVpnSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnSiteResultOutput)
}

type LookupVpnSiteOutputArgs struct {
	// The resource group name of the VpnSite.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the VpnSite being retrieved.
	VpnSiteName pulumi.StringInput `pulumi:"vpnSiteName"`
}

func (LookupVpnSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnSiteArgs)(nil)).Elem()
}

// VpnSite Resource.
type LookupVpnSiteResultOutput struct{ *pulumi.OutputState }

func (LookupVpnSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnSiteResult)(nil)).Elem()
}

func (o LookupVpnSiteResultOutput) ToLookupVpnSiteResultOutput() LookupVpnSiteResultOutput {
	return o
}

func (o LookupVpnSiteResultOutput) ToLookupVpnSiteResultOutputWithContext(ctx context.Context) LookupVpnSiteResultOutput {
	return o
}

// The AddressSpace that contains an array of IP address ranges.
func (o LookupVpnSiteResultOutput) AddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *AddressSpaceResponse { return v.AddressSpace }).(AddressSpaceResponsePtrOutput)
}

// The set of bgp properties.
func (o LookupVpnSiteResultOutput) BgpProperties() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *BgpSettingsResponse { return v.BgpProperties }).(BgpSettingsResponsePtrOutput)
}

// The device properties.
func (o LookupVpnSiteResultOutput) DeviceProperties() DevicePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *DevicePropertiesResponse { return v.DeviceProperties }).(DevicePropertiesResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVpnSiteResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupVpnSiteResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The ip-address for the vpn-site.
func (o LookupVpnSiteResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// IsSecuritySite flag.
func (o LookupVpnSiteResultOutput) IsSecuritySite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *bool { return v.IsSecuritySite }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o LookupVpnSiteResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupVpnSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

// Office365 Policy.
func (o LookupVpnSiteResultOutput) O365Policy() O365PolicyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *O365PolicyPropertiesResponse { return v.O365Policy }).(O365PolicyPropertiesResponsePtrOutput)
}

// The provisioning state of the VPN site resource.
func (o LookupVpnSiteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The key for vpn-site that can be used for connections.
func (o LookupVpnSiteResultOutput) SiteKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *string { return v.SiteKey }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupVpnSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupVpnSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

// The VirtualWAN to which the vpnSite belongs.
func (o LookupVpnSiteResultOutput) VirtualWan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *SubResourceResponse { return v.VirtualWan }).(SubResourceResponsePtrOutput)
}

// List of all vpn site links.
func (o LookupVpnSiteResultOutput) VpnSiteLinks() VpnSiteLinkResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) []VpnSiteLinkResponse { return v.VpnSiteLinks }).(VpnSiteLinkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnSiteResultOutput{})
}