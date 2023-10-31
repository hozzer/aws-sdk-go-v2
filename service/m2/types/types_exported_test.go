// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/m2/types"
)

func ExampleBatchJobDefinition_outputUsage() {
	var union types.BatchJobDefinition
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.BatchJobDefinitionMemberFileBatchJobDefinition:
		_ = v.Value // Value is types.FileBatchJobDefinition

	case *types.BatchJobDefinitionMemberScriptBatchJobDefinition:
		_ = v.Value // Value is types.ScriptBatchJobDefinition

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ScriptBatchJobDefinition
var _ *types.FileBatchJobDefinition

func ExampleBatchJobIdentifier_outputUsage() {
	var union types.BatchJobIdentifier
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.BatchJobIdentifierMemberFileBatchJobIdentifier:
		_ = v.Value // Value is types.FileBatchJobIdentifier

	case *types.BatchJobIdentifierMemberS3BatchJobIdentifier:
		_ = v.Value // Value is types.S3BatchJobIdentifier

	case *types.BatchJobIdentifierMemberScriptBatchJobIdentifier:
		_ = v.Value // Value is types.ScriptBatchJobIdentifier

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3BatchJobIdentifier
var _ *types.FileBatchJobIdentifier
var _ *types.ScriptBatchJobIdentifier

func ExampleDatasetDetailOrgAttributes_outputUsage() {
	var union types.DatasetDetailOrgAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DatasetDetailOrgAttributesMemberGdg:
		_ = v.Value // Value is types.GdgDetailAttributes

	case *types.DatasetDetailOrgAttributesMemberPo:
		_ = v.Value // Value is types.PoDetailAttributes

	case *types.DatasetDetailOrgAttributesMemberPs:
		_ = v.Value // Value is types.PsDetailAttributes

	case *types.DatasetDetailOrgAttributesMemberVsam:
		_ = v.Value // Value is types.VsamDetailAttributes

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.PoDetailAttributes
var _ *types.VsamDetailAttributes
var _ *types.PsDetailAttributes
var _ *types.GdgDetailAttributes

func ExampleDataSetImportConfig_outputUsage() {
	var union types.DataSetImportConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DataSetImportConfigMemberDataSets:
		_ = v.Value // Value is []types.DataSetImportItem

	case *types.DataSetImportConfigMemberS3Location:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ []types.DataSetImportItem

func ExampleDatasetOrgAttributes_outputUsage() {
	var union types.DatasetOrgAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DatasetOrgAttributesMemberGdg:
		_ = v.Value // Value is types.GdgAttributes

	case *types.DatasetOrgAttributesMemberPo:
		_ = v.Value // Value is types.PoAttributes

	case *types.DatasetOrgAttributesMemberPs:
		_ = v.Value // Value is types.PsAttributes

	case *types.DatasetOrgAttributesMemberVsam:
		_ = v.Value // Value is types.VsamAttributes

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.PsAttributes
var _ *types.GdgAttributes
var _ *types.VsamAttributes
var _ *types.PoAttributes

func ExampleDefinition_outputUsage() {
	var union types.Definition
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DefinitionMemberContent:
		_ = v.Value // Value is string

	case *types.DefinitionMemberS3Location:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string

func ExampleExternalLocation_outputUsage() {
	var union types.ExternalLocation
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ExternalLocationMemberS3Location:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleJobIdentifier_outputUsage() {
	var union types.JobIdentifier
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.JobIdentifierMemberFileName:
		_ = v.Value // Value is string

	case *types.JobIdentifierMemberScriptName:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleStorageConfiguration_outputUsage() {
	var union types.StorageConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.StorageConfigurationMemberEfs:
		_ = v.Value // Value is types.EfsStorageConfiguration

	case *types.StorageConfigurationMemberFsx:
		_ = v.Value // Value is types.FsxStorageConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FsxStorageConfiguration
var _ *types.EfsStorageConfiguration
