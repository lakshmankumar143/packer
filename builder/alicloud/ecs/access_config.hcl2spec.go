// Code generated by "hcl2-schema"; DO NOT EDIT.\n

package ecs

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*AlicloudAccessConfig) HCL2Schema() hcldec.ObjectSpec {
	s := map[string]hcldec.Spec{
		"AlicloudAccessKey": &hcldec.AttrSpec{Name:"AlicloudAccessKey", Type:cty.String, Required:false},
		"AlicloudSecretKey": &hcldec.AttrSpec{Name:"AlicloudSecretKey", Type:cty.String, Required:false},
		"AlicloudRegion": &hcldec.AttrSpec{Name:"AlicloudRegion", Type:cty.String, Required:false},
		"AlicloudSkipValidation": &hcldec.AttrSpec{Name:"AlicloudSkipValidation", Type:cty.Bool, Required:false},
		"SecurityToken": &hcldec.AttrSpec{Name:"SecurityToken", Type:cty.String, Required:false},
	}
	return hcldec.ObjectSpec(s)
}

