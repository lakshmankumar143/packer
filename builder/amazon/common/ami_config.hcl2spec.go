// Code generated by "hcl2-schema"; DO NOT EDIT.\n

package common

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*AMIConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"AMIName":                 &hcldec.AttrSpec{Name: "ami_name", Type: cty.String, Required: false},
		"AMIDescription":          &hcldec.AttrSpec{Name: "ami_description", Type: cty.String, Required: false},
		"AMIVirtType":             &hcldec.AttrSpec{Name: "ami_virtualization_type", Type: cty.String, Required: false},
		"AMIUsers":                &hcldec.AttrSpec{Name: "ami_users", Type: cty.List(cty.String), Required: false},
		"AMIGroups":               &hcldec.AttrSpec{Name: "ami_groups", Type: cty.List(cty.String), Required: false},
		"AMIProductCodes":         &hcldec.AttrSpec{Name: "ami_product_codes", Type: cty.List(cty.String), Required: false},
		"AMIRegions":              &hcldec.AttrSpec{Name: "ami_regions", Type: cty.List(cty.String), Required: false},
		"AMISkipRegionValidation": &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"AMIENASupport":           &hcldec.AttrSpec{Name: "ena_support", Type: cty.Bool, Required: false},
		"AMISriovNetSupport":      &hcldec.AttrSpec{Name: "sriov_support", Type: cty.Bool, Required: false},
		"AMIForceDeregister":      &hcldec.AttrSpec{Name: "force_deregister", Type: cty.Bool, Required: false},
		"AMIForceDeleteSnapshot":  &hcldec.AttrSpec{Name: "force_delete_snapshot", Type: cty.Bool, Required: false},
		"AMIEncryptBootVolume":    &hcldec.AttrSpec{Name: "encrypt_boot", Type: cty.Bool, Required: false},
		"AMIKmsKeyId":             &hcldec.AttrSpec{Name: "kms_key_id", Type: cty.String, Required: false},
		"AMISkipBuildRegion":      &hcldec.AttrSpec{Name: "skip_save_build_region", Type: cty.Bool, Required: false},
		"SnapshotUsers":           &hcldec.AttrSpec{Name: "snapshot_users", Type: cty.List(cty.String), Required: false},
		"SnapshotGroups":          &hcldec.AttrSpec{Name: "snapshot_groups", Type: cty.List(cty.String), Required: false},
	}
	for k, v := range (*AMIConfig)(nil).AMITags.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (*AMIConfig)(nil).AMIRegionKMSKeyIDs.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (*AMIConfig)(nil).SnapshotTags.HCL2Spec() {
		s[k] = v
	}
	return s
}
