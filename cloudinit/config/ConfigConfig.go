// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Specify whether or not to base64 encode the `rendered` output.
	//
	// Defaults to `true`, and cannot be disabled if gzip is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.7/docs/resources/config#base64_encode Config#base64_encode}
	Base64Encode interface{} `field:"optional" json:"base64Encode" yaml:"base64Encode"`
	// Specify the Writer's default boundary separator. Defaults to `MIMEBOUNDARY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.7/docs/resources/config#boundary Config#boundary}
	Boundary *string `field:"optional" json:"boundary" yaml:"boundary"`
	// Specify whether or not to gzip the `rendered` output. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.7/docs/resources/config#gzip Config#gzip}
	Gzip interface{} `field:"optional" json:"gzip" yaml:"gzip"`
	// part block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/cloudinit/2.3.7/docs/resources/config#part Config#part}
	Part interface{} `field:"optional" json:"part" yaml:"part"`
}

