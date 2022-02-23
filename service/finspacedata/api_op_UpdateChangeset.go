// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a FinSpace Changeset.
func (c *Client) UpdateChangeset(ctx context.Context, params *UpdateChangesetInput, optFns ...func(*Options)) (*UpdateChangesetOutput, error) {
	if params == nil {
		params = &UpdateChangesetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateChangeset", params, optFns, c.addOperationUpdateChangesetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateChangesetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to update an existing changeset.
type UpdateChangesetInput struct {

	// The unique identifier for the Changeset to update.
	//
	// This member is required.
	ChangesetId *string

	// The unique identifier for the FinSpace Dataset in which the Changeset is
	// created.
	//
	// This member is required.
	DatasetId *string

	// Options that define the structure of the source file(s) including the format
	// type (formatType), header row (withHeader), data separation character
	// (separator) and the type of compression (compression). formatType is a required
	// attribute and can have the following values:
	//
	// * PARQUET - Parquet source file
	// format.
	//
	// * CSV - CSV source file format.
	//
	// * JSON - JSON source file format.
	//
	// *
	// XML - XML source file format.
	//
	// Here is an example of how you could specify the
	// formatParams:  "formatParams": { "formatType": "CSV", "withHeader": "true",
	// "separator": ",", "compression":"None" }  Note that if you only provide
	// formatType as CSV, the rest of the attributes will automatically default to CSV
	// values as following:  { "withHeader": "true", "separator": "," }  For more
	// information about supported file formats, see Supported Data Types and File
	// Formats
	// (https://docs.aws.amazon.com/finspace/latest/userguide/supported-data-types.html)
	// in the FinSpace User Guide.
	//
	// This member is required.
	FormatParams map[string]string

	// Options that define the location of the data being ingested (s3SourcePath) and
	// the source of the changeset (sourceType). Both s3SourcePath and sourceType are
	// required attributes. Here is an example of how you could specify the
	// sourceParams:  "sourceParams": { "s3SourcePath":
	// "s3://finspace-landing-us-east-2-bk7gcfvitndqa6ebnvys4d/scratch/wr5hh8pwkpqqkxa4sxrmcw/ingestion/equity.csv",
	// "sourceType": "S3" }  The S3 path that you specify must allow the FinSpace role
	// access. To do that, you first need to configure the IAM policy on S3 bucket. For
	// more information, see Loading data from an Amazon S3 Bucket using the FinSpace
	// API
	// (https://docs.aws.amazon.com/finspace/latest/data-api/fs-using-the-finspace-api.html#access-s3-buckets)section.
	//
	// This member is required.
	SourceParams map[string]string

	// A token that ensures idempotency. This token expires in 10 minutes.
	ClientToken *string

	noSmithyDocumentSerde
}

// The response from a update changeset operation.
type UpdateChangesetOutput struct {

	// The unique identifier for the Changeset to update.
	ChangesetId *string

	// The unique identifier for the FinSpace Dataset in which the Changeset is
	// created.
	DatasetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateChangesetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateChangeset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateChangeset{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opUpdateChangesetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateChangesetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateChangeset(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpUpdateChangeset struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateChangeset) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateChangeset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateChangesetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateChangesetInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateChangesetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateChangeset{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateChangeset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "UpdateChangeset",
	}
}