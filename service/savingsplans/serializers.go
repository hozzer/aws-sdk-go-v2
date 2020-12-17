// Code generated by smithy-go-codegen DO NOT EDIT.

package savingsplans

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/encoding/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/encoding/json"
	"github.com/awslabs/smithy-go/middleware"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCreateSavingsPlan struct {
}

func (*awsRestjson1_serializeOpCreateSavingsPlan) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateSavingsPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateSavingsPlanInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/CreateSavingsPlan")
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
	if err := awsRestjson1_serializeOpDocumentCreateSavingsPlanInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsCreateSavingsPlanInput(v *CreateSavingsPlanInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateSavingsPlanInput(v *CreateSavingsPlanInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("clientToken")
		ok.String(*v.ClientToken)
	}

	if v.Commitment != nil {
		ok := object.Key("commitment")
		ok.String(*v.Commitment)
	}

	if v.PurchaseTime != nil {
		ok := object.Key("purchaseTime")
		ok.Double(smithytime.FormatEpochSeconds(*v.PurchaseTime))
	}

	if v.SavingsPlanOfferingId != nil {
		ok := object.Key("savingsPlanOfferingId")
		ok.String(*v.SavingsPlanOfferingId)
	}

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	if v.UpfrontPaymentAmount != nil {
		ok := object.Key("upfrontPaymentAmount")
		ok.String(*v.UpfrontPaymentAmount)
	}

	return nil
}

type awsRestjson1_serializeOpDeleteQueuedSavingsPlan struct {
}

func (*awsRestjson1_serializeOpDeleteQueuedSavingsPlan) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteQueuedSavingsPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteQueuedSavingsPlanInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/DeleteQueuedSavingsPlan")
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
	if err := awsRestjson1_serializeOpDocumentDeleteQueuedSavingsPlanInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsDeleteQueuedSavingsPlanInput(v *DeleteQueuedSavingsPlanInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentDeleteQueuedSavingsPlanInput(v *DeleteQueuedSavingsPlanInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.SavingsPlanId != nil {
		ok := object.Key("savingsPlanId")
		ok.String(*v.SavingsPlanId)
	}

	return nil
}

type awsRestjson1_serializeOpDescribeSavingsPlanRates struct {
}

func (*awsRestjson1_serializeOpDescribeSavingsPlanRates) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeSavingsPlanRates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeSavingsPlanRatesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/DescribeSavingsPlanRates")
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
	if err := awsRestjson1_serializeOpDocumentDescribeSavingsPlanRatesInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsDescribeSavingsPlanRatesInput(v *DescribeSavingsPlanRatesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentDescribeSavingsPlanRatesInput(v *DescribeSavingsPlanRatesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Filters != nil {
		ok := object.Key("filters")
		if err := awsRestjson1_serializeDocumentSavingsPlanRateFilterList(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.SavingsPlanId != nil {
		ok := object.Key("savingsPlanId")
		ok.String(*v.SavingsPlanId)
	}

	return nil
}

type awsRestjson1_serializeOpDescribeSavingsPlans struct {
}

func (*awsRestjson1_serializeOpDescribeSavingsPlans) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeSavingsPlans) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeSavingsPlansInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/DescribeSavingsPlans")
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
	if err := awsRestjson1_serializeOpDocumentDescribeSavingsPlansInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsDescribeSavingsPlansInput(v *DescribeSavingsPlansInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentDescribeSavingsPlansInput(v *DescribeSavingsPlansInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Filters != nil {
		ok := object.Key("filters")
		if err := awsRestjson1_serializeDocumentSavingsPlanFilterList(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.SavingsPlanArns != nil {
		ok := object.Key("savingsPlanArns")
		if err := awsRestjson1_serializeDocumentSavingsPlanArnList(v.SavingsPlanArns, ok); err != nil {
			return err
		}
	}

	if v.SavingsPlanIds != nil {
		ok := object.Key("savingsPlanIds")
		if err := awsRestjson1_serializeDocumentSavingsPlanIdList(v.SavingsPlanIds, ok); err != nil {
			return err
		}
	}

	if v.States != nil {
		ok := object.Key("states")
		if err := awsRestjson1_serializeDocumentSavingsPlanStateList(v.States, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDescribeSavingsPlansOfferingRates struct {
}

func (*awsRestjson1_serializeOpDescribeSavingsPlansOfferingRates) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeSavingsPlansOfferingRates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeSavingsPlansOfferingRatesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/DescribeSavingsPlansOfferingRates")
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
	if err := awsRestjson1_serializeOpDocumentDescribeSavingsPlansOfferingRatesInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsDescribeSavingsPlansOfferingRatesInput(v *DescribeSavingsPlansOfferingRatesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentDescribeSavingsPlansOfferingRatesInput(v *DescribeSavingsPlansOfferingRatesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Filters != nil {
		ok := object.Key("filters")
		if err := awsRestjson1_serializeDocumentSavingsPlanOfferingRateFiltersList(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != 0 {
		ok := object.Key("maxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.Operations != nil {
		ok := object.Key("operations")
		if err := awsRestjson1_serializeDocumentSavingsPlanRateOperationList(v.Operations, ok); err != nil {
			return err
		}
	}

	if v.Products != nil {
		ok := object.Key("products")
		if err := awsRestjson1_serializeDocumentSavingsPlanProductTypeList(v.Products, ok); err != nil {
			return err
		}
	}

	if v.SavingsPlanOfferingIds != nil {
		ok := object.Key("savingsPlanOfferingIds")
		if err := awsRestjson1_serializeDocumentUUIDs(v.SavingsPlanOfferingIds, ok); err != nil {
			return err
		}
	}

	if v.SavingsPlanPaymentOptions != nil {
		ok := object.Key("savingsPlanPaymentOptions")
		if err := awsRestjson1_serializeDocumentSavingsPlanPaymentOptionList(v.SavingsPlanPaymentOptions, ok); err != nil {
			return err
		}
	}

	if v.SavingsPlanTypes != nil {
		ok := object.Key("savingsPlanTypes")
		if err := awsRestjson1_serializeDocumentSavingsPlanTypeList(v.SavingsPlanTypes, ok); err != nil {
			return err
		}
	}

	if v.ServiceCodes != nil {
		ok := object.Key("serviceCodes")
		if err := awsRestjson1_serializeDocumentSavingsPlanRateServiceCodeList(v.ServiceCodes, ok); err != nil {
			return err
		}
	}

	if v.UsageTypes != nil {
		ok := object.Key("usageTypes")
		if err := awsRestjson1_serializeDocumentSavingsPlanRateUsageTypeList(v.UsageTypes, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDescribeSavingsPlansOfferings struct {
}

func (*awsRestjson1_serializeOpDescribeSavingsPlansOfferings) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeSavingsPlansOfferings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeSavingsPlansOfferingsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/DescribeSavingsPlansOfferings")
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
	if err := awsRestjson1_serializeOpDocumentDescribeSavingsPlansOfferingsInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsDescribeSavingsPlansOfferingsInput(v *DescribeSavingsPlansOfferingsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentDescribeSavingsPlansOfferingsInput(v *DescribeSavingsPlansOfferingsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Currencies != nil {
		ok := object.Key("currencies")
		if err := awsRestjson1_serializeDocumentCurrencyList(v.Currencies, ok); err != nil {
			return err
		}
	}

	if v.Descriptions != nil {
		ok := object.Key("descriptions")
		if err := awsRestjson1_serializeDocumentSavingsPlanDescriptionsList(v.Descriptions, ok); err != nil {
			return err
		}
	}

	if v.Durations != nil {
		ok := object.Key("durations")
		if err := awsRestjson1_serializeDocumentDurationsList(v.Durations, ok); err != nil {
			return err
		}
	}

	if v.Filters != nil {
		ok := object.Key("filters")
		if err := awsRestjson1_serializeDocumentSavingsPlanOfferingFiltersList(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != 0 {
		ok := object.Key("maxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.OfferingIds != nil {
		ok := object.Key("offeringIds")
		if err := awsRestjson1_serializeDocumentUUIDs(v.OfferingIds, ok); err != nil {
			return err
		}
	}

	if v.Operations != nil {
		ok := object.Key("operations")
		if err := awsRestjson1_serializeDocumentSavingsPlanOperationList(v.Operations, ok); err != nil {
			return err
		}
	}

	if v.PaymentOptions != nil {
		ok := object.Key("paymentOptions")
		if err := awsRestjson1_serializeDocumentSavingsPlanPaymentOptionList(v.PaymentOptions, ok); err != nil {
			return err
		}
	}

	if v.PlanTypes != nil {
		ok := object.Key("planTypes")
		if err := awsRestjson1_serializeDocumentSavingsPlanTypeList(v.PlanTypes, ok); err != nil {
			return err
		}
	}

	if len(v.ProductType) > 0 {
		ok := object.Key("productType")
		ok.String(string(v.ProductType))
	}

	if v.ServiceCodes != nil {
		ok := object.Key("serviceCodes")
		if err := awsRestjson1_serializeDocumentSavingsPlanServiceCodeList(v.ServiceCodes, ok); err != nil {
			return err
		}
	}

	if v.UsageTypes != nil {
		ok := object.Key("usageTypes")
		if err := awsRestjson1_serializeDocumentSavingsPlanUsageTypeList(v.UsageTypes, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListTagsForResource struct {
}

func (*awsRestjson1_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/ListTagsForResource")
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
	if err := awsRestjson1_serializeOpDocumentListTagsForResourceInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentListTagsForResourceInput(v *ListTagsForResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	return nil
}

type awsRestjson1_serializeOpTagResource struct {
}

func (*awsRestjson1_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/TagResource")
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
	if err := awsRestjson1_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpUntagResource struct {
}

func (*awsRestjson1_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/UntagResource")
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
	if err := awsRestjson1_serializeOpDocumentUntagResourceInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUntagResourceInput(v *UntagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.TagKeys != nil {
		ok := object.Key("tagKeys")
		if err := awsRestjson1_serializeDocumentTagKeyList(v.TagKeys, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentCurrencyList(v []types.CurrencyCode, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentDurationsList(v []int64, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.Long(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentFilterValuesList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentListOfStrings(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanArnList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanDescriptionsList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanFilter(v *types.SavingsPlanFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Name) > 0 {
		ok := object.Key("name")
		ok.String(string(v.Name))
	}

	if v.Values != nil {
		ok := object.Key("values")
		if err := awsRestjson1_serializeDocumentListOfStrings(v.Values, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanFilterList(v []types.SavingsPlanFilter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentSavingsPlanFilter(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanIdList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanOfferingFilterElement(v *types.SavingsPlanOfferingFilterElement, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Name) > 0 {
		ok := object.Key("name")
		ok.String(string(v.Name))
	}

	if v.Values != nil {
		ok := object.Key("values")
		if err := awsRestjson1_serializeDocumentFilterValuesList(v.Values, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanOfferingFiltersList(v []types.SavingsPlanOfferingFilterElement, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentSavingsPlanOfferingFilterElement(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanOfferingRateFilterElement(v *types.SavingsPlanOfferingRateFilterElement, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Name) > 0 {
		ok := object.Key("name")
		ok.String(string(v.Name))
	}

	if v.Values != nil {
		ok := object.Key("values")
		if err := awsRestjson1_serializeDocumentFilterValuesList(v.Values, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanOfferingRateFiltersList(v []types.SavingsPlanOfferingRateFilterElement, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentSavingsPlanOfferingRateFilterElement(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanOperationList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanPaymentOptionList(v []types.SavingsPlanPaymentOption, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanProductTypeList(v []types.SavingsPlanProductType, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanRateFilter(v *types.SavingsPlanRateFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Name) > 0 {
		ok := object.Key("name")
		ok.String(string(v.Name))
	}

	if v.Values != nil {
		ok := object.Key("values")
		if err := awsRestjson1_serializeDocumentListOfStrings(v.Values, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanRateFilterList(v []types.SavingsPlanRateFilter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentSavingsPlanRateFilter(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanRateOperationList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanRateServiceCodeList(v []types.SavingsPlanRateServiceCode, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanRateUsageTypeList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanServiceCodeList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanStateList(v []types.SavingsPlanState, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanTypeList(v []types.SavingsPlanType, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentSavingsPlanUsageTypeList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentTagKeyList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentTagMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsRestjson1_serializeDocumentUUIDs(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}
