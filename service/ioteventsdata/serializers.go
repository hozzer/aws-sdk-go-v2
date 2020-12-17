// Code generated by smithy-go-codegen DO NOT EDIT.

package ioteventsdata

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/encoding/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/encoding/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsRestjson1_serializeOpBatchPutMessage struct {
}

func (*awsRestjson1_serializeOpBatchPutMessage) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpBatchPutMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BatchPutMessageInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/inputs/messages")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentBatchPutMessageInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsBatchPutMessageInput(v *BatchPutMessageInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentBatchPutMessageInput(v *BatchPutMessageInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Messages != nil {
		ok := object.Key("messages")
		if err := awsRestjson1_serializeDocumentMessages(v.Messages, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpBatchUpdateDetector struct {
}

func (*awsRestjson1_serializeOpBatchUpdateDetector) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpBatchUpdateDetector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BatchUpdateDetectorInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/detectors")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentBatchUpdateDetectorInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsBatchUpdateDetectorInput(v *BatchUpdateDetectorInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentBatchUpdateDetectorInput(v *BatchUpdateDetectorInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Detectors != nil {
		ok := object.Key("detectors")
		if err := awsRestjson1_serializeDocumentUpdateDetectorRequests(v.Detectors, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDescribeDetector struct {
}

func (*awsRestjson1_serializeOpDescribeDetector) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeDetector) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeDetectorInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/detectors/{detectorModelName}/keyValues")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDescribeDetectorInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDescribeDetectorInput(v *DescribeDetectorInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.DetectorModelName == nil || len(*v.DetectorModelName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member detectorModelName must not be empty")}
	}
	if v.DetectorModelName != nil {
		if err := encoder.SetURI("detectorModelName").String(*v.DetectorModelName); err != nil {
			return err
		}
	}

	if v.KeyValue != nil {
		encoder.SetQuery("keyValue").String(*v.KeyValue)
	}

	return nil
}

type awsRestjson1_serializeOpListDetectors struct {
}

func (*awsRestjson1_serializeOpListDetectors) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListDetectors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListDetectorsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/detectors/{detectorModelName}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListDetectorsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListDetectorsInput(v *ListDetectorsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.DetectorModelName == nil || len(*v.DetectorModelName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member detectorModelName must not be empty")}
	}
	if v.DetectorModelName != nil {
		if err := encoder.SetURI("detectorModelName").String(*v.DetectorModelName); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if v.StateName != nil {
		encoder.SetQuery("stateName").String(*v.StateName)
	}

	return nil
}

func awsRestjson1_serializeDocumentDetectorStateDefinition(v *types.DetectorStateDefinition, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.StateName != nil {
		ok := object.Key("stateName")
		ok.String(*v.StateName)
	}

	if v.Timers != nil {
		ok := object.Key("timers")
		if err := awsRestjson1_serializeDocumentTimerDefinitions(v.Timers, ok); err != nil {
			return err
		}
	}

	if v.Variables != nil {
		ok := object.Key("variables")
		if err := awsRestjson1_serializeDocumentVariableDefinitions(v.Variables, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentMessage(v *types.Message, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.InputName != nil {
		ok := object.Key("inputName")
		ok.String(*v.InputName)
	}

	if v.MessageId != nil {
		ok := object.Key("messageId")
		ok.String(*v.MessageId)
	}

	if v.Payload != nil {
		ok := object.Key("payload")
		ok.Base64EncodeBytes(v.Payload)
	}

	return nil
}

func awsRestjson1_serializeDocumentMessages(v []types.Message, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentMessage(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentTimerDefinition(v *types.TimerDefinition, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	if v.Seconds != nil {
		ok := object.Key("seconds")
		ok.Integer(*v.Seconds)
	}

	return nil
}

func awsRestjson1_serializeDocumentTimerDefinitions(v []types.TimerDefinition, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentTimerDefinition(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentUpdateDetectorRequest(v *types.UpdateDetectorRequest, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DetectorModelName != nil {
		ok := object.Key("detectorModelName")
		ok.String(*v.DetectorModelName)
	}

	if v.KeyValue != nil {
		ok := object.Key("keyValue")
		ok.String(*v.KeyValue)
	}

	if v.MessageId != nil {
		ok := object.Key("messageId")
		ok.String(*v.MessageId)
	}

	if v.State != nil {
		ok := object.Key("state")
		if err := awsRestjson1_serializeDocumentDetectorStateDefinition(v.State, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentUpdateDetectorRequests(v []types.UpdateDetectorRequest, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentUpdateDetectorRequest(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentVariableDefinition(v *types.VariableDefinition, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	if v.Value != nil {
		ok := object.Key("value")
		ok.String(*v.Value)
	}

	return nil
}

func awsRestjson1_serializeDocumentVariableDefinitions(v []types.VariableDefinition, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentVariableDefinition(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}
