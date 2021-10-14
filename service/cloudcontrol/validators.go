// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudcontrol

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelResourceRequest struct {
}

func (*validateOpCancelResourceRequest) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelResourceRequest) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelResourceRequestInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelResourceRequestInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateResource struct {
}

func (*validateOpCreateResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteResource struct {
}

func (*validateOpDeleteResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResource struct {
}

func (*validateOpGetResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourceRequestStatus struct {
}

func (*validateOpGetResourceRequestStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourceRequestStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourceRequestStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourceRequestStatusInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListResources struct {
}

func (*validateOpListResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateResource struct {
}

func (*validateOpUpdateResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelResourceRequestValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelResourceRequest{}, middleware.After)
}

func addOpCreateResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateResource{}, middleware.After)
}

func addOpDeleteResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteResource{}, middleware.After)
}

func addOpGetResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResource{}, middleware.After)
}

func addOpGetResourceRequestStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResourceRequestStatus{}, middleware.After)
}

func addOpListResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListResources{}, middleware.After)
}

func addOpUpdateResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateResource{}, middleware.After)
}

func validateOpCancelResourceRequestInput(v *CancelResourceRequestInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelResourceRequestInput"}
	if v.RequestToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RequestToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateResourceInput(v *CreateResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateResourceInput"}
	if v.TypeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TypeName"))
	}
	if v.DesiredState == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DesiredState"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteResourceInput(v *DeleteResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteResourceInput"}
	if v.TypeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TypeName"))
	}
	if v.Identifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Identifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourceInput(v *GetResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourceInput"}
	if v.TypeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TypeName"))
	}
	if v.Identifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Identifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourceRequestStatusInput(v *GetResourceRequestStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourceRequestStatusInput"}
	if v.RequestToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RequestToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListResourcesInput(v *ListResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListResourcesInput"}
	if v.TypeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TypeName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateResourceInput(v *UpdateResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateResourceInput"}
	if v.TypeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TypeName"))
	}
	if v.Identifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Identifier"))
	}
	if v.PatchDocument == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PatchDocument"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}