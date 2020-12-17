// Code generated by smithy-go-codegen DO NOT EDIT.

package workmailmessageflow

import (
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/encoding/httpbinding"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsRestjson1_serializeOpGetRawMessageContent struct {
}

func (*awsRestjson1_serializeOpGetRawMessageContent) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetRawMessageContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetRawMessageContentInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/messages/{messageId}")
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

	if err := awsRestjson1_serializeOpHttpBindingsGetRawMessageContentInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetRawMessageContentInput(v *GetRawMessageContentInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MessageId == nil || len(*v.MessageId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member messageId must not be empty")}
	}
	if v.MessageId != nil {
		if err := encoder.SetURI("messageId").String(*v.MessageId); err != nil {
			return err
		}
	}

	return nil
}
