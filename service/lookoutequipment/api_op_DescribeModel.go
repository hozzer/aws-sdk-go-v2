// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides a JSON containing the overall information about a specific machine
// learning model, including model name and ARN, dataset, training and evaluation
// information, status, and so on.
func (c *Client) DescribeModel(ctx context.Context, params *DescribeModelInput, optFns ...func(*Options)) (*DescribeModelOutput, error) {
	if params == nil {
		params = &DescribeModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModel", params, optFns, c.addOperationDescribeModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelInput struct {

	// The name of the machine learning model to be described.
	//
	// This member is required.
	ModelName *string

	noSmithyDocumentSerde
}

type DescribeModelOutput struct {

	// Indicates the end time of the inference data that has been accumulated.
	AccumulatedInferenceDataEndTime *time.Time

	// Indicates the start time of the inference data that has been accumulated.
	AccumulatedInferenceDataStartTime *time.Time

	// The name of the model version used by the inference schedular when running a
	// scheduled inference execution.
	ActiveModelVersion *int64

	// The Amazon Resource Name (ARN) of the model version used by the inference
	// scheduler when running a scheduled inference execution.
	ActiveModelVersionArn *string

	// Indicates the time and date at which the machine learning model was created.
	CreatedAt *time.Time

	// The configuration is the TargetSamplingRate , which is the sampling rate of the
	// data after post processing by Amazon Lookout for Equipment. For example, if you
	// provide data that has been collected at a 1 second level and you want the system
	// to resample the data at a 1 minute rate before training, the TargetSamplingRate
	// is 1 minute. When providing a value for the TargetSamplingRate , you must attach
	// the prefix "PT" to the rate you want. The value for a 1 second rate is therefore
	// PT1S, the value for a 15 minute rate is PT15M, and the value for a 1 hour rate
	// is PT1H
	DataPreProcessingConfiguration *types.DataPreProcessingConfiguration

	// The Amazon Resouce Name (ARN) of the dataset used to create the machine
	// learning model being described.
	DatasetArn *string

	// The name of the dataset being used by the machine learning being described.
	DatasetName *string

	// Indicates the time reference in the dataset that was used to end the subset of
	// evaluation data for the machine learning model.
	EvaluationDataEndTime *time.Time

	// Indicates the time reference in the dataset that was used to begin the subset
	// of evaluation data for the machine learning model.
	EvaluationDataStartTime *time.Time

	// If the training of the machine learning model failed, this indicates the reason
	// for that failure.
	FailedReason *string

	// The date and time when the import job was completed. This field appears if the
	// active model version was imported.
	ImportJobEndTime *time.Time

	// The date and time when the import job was started. This field appears if the
	// active model version was imported.
	ImportJobStartTime *time.Time

	// Specifies configuration information about the labels input, including its S3
	// location.
	LabelsInputConfiguration *types.LabelsInputConfiguration

	// Indicates the last time the machine learning model was updated. The type of
	// update is not specified.
	LastUpdatedTime *time.Time

	// Indicates the number of days of data used in the most recent scheduled
	// retraining run.
	LatestScheduledRetrainingAvailableDataInDays *int32

	// If the model version was generated by retraining and the training failed, this
	// indicates the reason for that failure.
	LatestScheduledRetrainingFailedReason *string

	// Indicates the most recent model version that was generated by retraining.
	LatestScheduledRetrainingModelVersion *int64

	// Indicates the start time of the most recent scheduled retraining run.
	LatestScheduledRetrainingStartTime *time.Time

	// Indicates the status of the most recent scheduled retraining run.
	LatestScheduledRetrainingStatus types.ModelVersionStatus

	// The Amazon Resource Name (ARN) of the machine learning model being described.
	ModelArn *string

	// Configuration information for the model's pointwise model diagnostics.
	ModelDiagnosticsOutputConfiguration *types.ModelDiagnosticsOutputConfiguration

	// The Model Metrics show an aggregated summary of the model's performance within
	// the evaluation time range. This is the JSON content of the metrics created when
	// evaluating the model.
	//
	// This value conforms to the media type: application/json
	ModelMetrics *string

	// The name of the machine learning model being described.
	ModelName *string

	// Provides a quality assessment for a model that uses labels. If Lookout for
	// Equipment determines that the model quality is poor based on training metrics,
	// the value is POOR_QUALITY_DETECTED . Otherwise, the value is
	// QUALITY_THRESHOLD_MET . If the model is unlabeled, the model quality can't be
	// assessed and the value of ModelQuality is CANNOT_DETERMINE_QUALITY . In this
	// situation, you can get a model quality assessment by adding labels to the input
	// dataset and retraining the model. For information about using labels with your
	// models, see Understanding labeling (https://docs.aws.amazon.com/lookout-for-equipment/latest/ug/understanding-labeling.html)
	// . For information about improving the quality of a model, see Best practices
	// with Amazon Lookout for Equipment (https://docs.aws.amazon.com/lookout-for-equipment/latest/ug/best-practices.html)
	// .
	ModelQuality types.ModelQuality

	// The date the active model version was activated.
	ModelVersionActivatedAt *time.Time

	// Indicates the date and time that the next scheduled retraining run will start
	// on. Lookout for Equipment truncates the time you provide to the nearest UTC day.
	NextScheduledRetrainingStartDate *time.Time

	// Indicates that the asset associated with this sensor has been shut off. As long
	// as this condition is met, Lookout for Equipment will not use data from this
	// asset for training, evaluation, or inference.
	OffCondition *string

	// The model version that was set as the active model version prior to the current
	// active model version.
	PreviousActiveModelVersion *int64

	// The ARN of the model version that was set as the active model version prior to
	// the current active model version.
	PreviousActiveModelVersionArn *string

	// The date and time when the previous active model version was activated.
	PreviousModelVersionActivatedAt *time.Time

	// If the model version was retrained, this field shows a summary of the
	// performance of the prior model on the new training range. You can use the
	// information in this JSON-formatted object to compare the new model version and
	// the prior model version.
	//
	// This value conforms to the media type: application/json
	PriorModelMetrics *string

	// Indicates the status of the retraining scheduler.
	RetrainingSchedulerStatus types.RetrainingSchedulerStatus

	// The Amazon Resource Name (ARN) of a role with permission to access the data
	// source for the machine learning model being described.
	RoleArn *string

	// A JSON description of the data that is in each time series dataset, including
	// names, column names, and data types.
	//
	// This value conforms to the media type: application/json
	Schema *string

	// Provides the identifier of the KMS key used to encrypt model data by Amazon
	// Lookout for Equipment.
	ServerSideKmsKeyId *string

	// The Amazon Resource Name (ARN) of the source model version. This field appears
	// if the active model version was imported.
	SourceModelVersionArn *string

	// Specifies the current status of the model being described. Status describes the
	// status of the most recent action of the model.
	Status types.ModelStatus

	// Indicates the time reference in the dataset that was used to end the subset of
	// training data for the machine learning model.
	TrainingDataEndTime *time.Time

	// Indicates the time reference in the dataset that was used to begin the subset
	// of training data for the machine learning model.
	TrainingDataStartTime *time.Time

	// Indicates the time at which the training of the machine learning model was
	// completed.
	TrainingExecutionEndTime *time.Time

	// Indicates the time at which the training of the machine learning model began.
	TrainingExecutionStartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeModel"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModel(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeModel",
	}
}
