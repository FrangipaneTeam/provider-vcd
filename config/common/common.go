// Package common package contains common functions and constants that are used
/*
Copyright 2021 Upbound Inc.
*/
package common

import (
	"fmt"
	"strings"

	"github.com/FrangipaneTeam/provider-vcd/pkg/tools"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/upbound/upjet/pkg/resource"
)

var (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = fmt.Sprintf("%s/config/common", tools.ModulePath)

	// PathARNExtractor is the golang path to ARNExtractor function
	// in this package.
	PathARNExtractor = SelfPackagePath + ".ARNExtractor()"

	// PathTerraformIDExtractor is the golang path to TerraformID extractor
	// function in this package.
	PathTerraformIDExtractor = SelfPackagePath + ".TerraformID()"

	// PathTerraformIDWithoutURNExtractor is the golang path to TerraformIDWithoutURNExtractor
	PathTerraformIDWithoutURNExtractor = SelfPackagePath + ".IDWithoutURNExtractor()"

	// PathFromSpecNameExtractor is the golang path to NameFromSpec
	PathFromSpecNameExtractor = SelfPackagePath + ".NameFromSpec()"

	// VersionV1Beta1 is used for resources that meet the v1beta1 criteria
	// here: https://github.com/upbound/arch/pull/33
	VersionV1Beta1 = "v1beta1"

	// ErrFmtNoAttribute is the error format for missing attribute in Terraform state.
	ErrFmtNoAttribute = "cannot find attribute %s in Terraform state"
	// ErrFmtUnexpectedType is the error format for unexpected type of attribute in Terraform state.
	ErrFmtUnexpectedType = "unexpected type for attribute %s in Terraform state"
)

// ARNExtractor extracts ARN of the resources from "status.atProvider.arn" which
// is quite common among all AWS resources.
func ARNExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.arn")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}

// TerraformID returns the Terraform ID string of the resource without any
// manipulation.
func TerraformID() reference.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		return tr.GetID()
	}
}

// IDWithoutURNExtractor extracts ID from "status.atProvider.id"
func IDWithoutURNExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.id")
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		words := strings.Split(r, ":")
		return words[len(words)-1]
	}
}

// NameFromSpec extracts name from "metadata.name"
func NameFromSpec() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		r, err := paved.GetString("metadata.name")
		if err != nil {
			// TODO should we log this error?
			return ""
		}
		return r
	}
}
