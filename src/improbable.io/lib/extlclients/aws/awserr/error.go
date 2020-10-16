// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package awserr

import (
	"improbable.io/lib/errors"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
)

const errMsg = "AWS error"

type Code string

const (
	// Based on common errors in http://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html
	// NOTE: If we find the error code which is pretty useful, we will update this list.

	// InvalidArgument section
	InvalidParameterErrCode            = Code(request.InvalidParameterErrCode)
	InvalidParameterCombinationErrCode = Code("InvalidParameterCombination")
	InvalidParameterValueErrCode       = Code("InvalidParameterValue")
	InvalidQueryParameterErrCode       = Code("InvalidQueryParameter")
	MalformedQueryStringErrCode        = Code("MalformedQueryString")
	MissingActionErrCode               = Code("MissingAction")
	MissingParameterErrCode            = Code("MissingParameter")
	UnknownParameterErrCode            = Code("UnknownParameter")
	MissingInputErrCode                = Code("MissingInput")

	// PermissionDenied section.
	AuthFailureErrCode                 = Code("AuthFailure")
	BlockedErrCode                     = Code("Blocked")
	IdempotentParameterMismatchErrCode = Code("IdempotentParameterMismatch")
	OptInRequiredErrCode               = Code("OptInRequired")
	PendingVerificationErrCode         = Code("PendingVerification")
	UnauthorizedOperationErrCode       = Code("UnauthorizedOperation")

	// Unauthenticated section.
	MissingAuthenticationTokenErrCode = Code("MissingAuthenticationToken")

	// DeadlineExceeded section.
	RequestExpiredErrCode = Code("RequestExpired")

	// FailedPrecondition section.
	ValidationErrorErrCode = Code("ValidationError")

	// ResourceExhausted section.
	InsufficientCapacityErrCode                 = Code("InsufficientCapacity")
	InsufficientAddressCapacityErrCode          = Code("InsufficientAddressCapacity")
	InsufficientInstanceCapacityErrCode         = Code("InsufficientInstanceCapacity")
	InsufficientHostCapacityErrCode             = Code("InsufficientHostCapacity")
	InsufficientReservedInstanceCapacityErrCode = Code("InsufficientReservedInstanceCapacity")

	// Internal section.
	InternalErrorErrCode   = Code("InternalError")
	InternalFailureErrCode = Code("InternalFailure")

	// Unavailable section.
	ServiceUnavailableErrCode = Code("ServiceUnavailable")
	UnavailableErrCode        = Code("Unavailable")

	//s3 related section
	NoSuchKey = Code("NoSuchKey")
)

// RemapError translates an `awserr` to Improbable error if possible.
func RemapError(err error) error {
	awsErr, ok := err.(awserr.Error)
	if !ok {
		// Not `awserr` type or nil.
		return err
	}

	switch Code(awsErr.Code()) {
	case InvalidParameterErrCode:
		fallthrough
	case InvalidParameterCombinationErrCode:
		fallthrough
	case InvalidParameterValueErrCode:
		fallthrough
	case InvalidQueryParameterErrCode:
		fallthrough
	case MalformedQueryStringErrCode:
		fallthrough
	case MissingActionErrCode:
		fallthrough
	case MissingParameterErrCode:
		fallthrough
	case UnknownParameterErrCode:
		fallthrough
	case MissingInputErrCode:
		return errors.New(err, errors.InvalidArgument, errMsg)

	case AuthFailureErrCode:
		fallthrough
	case BlockedErrCode:
		fallthrough
	case IdempotentParameterMismatchErrCode:
		fallthrough
	case OptInRequiredErrCode:
		fallthrough
	case PendingVerificationErrCode:
		fallthrough
	case UnauthorizedOperationErrCode:
		return errors.New(err, errors.PermissionDenied, errMsg)

	case MissingAuthenticationTokenErrCode:
		return errors.New(err, errors.Unauthenticated, errMsg)

	case RequestExpiredErrCode:
		return errors.New(err, errors.DeadlineExceeded, errMsg)

	case ValidationErrorErrCode:
		return errors.New(err, errors.FailedPrecondition, errMsg)

	case InsufficientCapacityErrCode:
		fallthrough
	case InsufficientHostCapacityErrCode:
		fallthrough
	case InsufficientReservedInstanceCapacityErrCode:
		fallthrough
	case InsufficientAddressCapacityErrCode:
		fallthrough
	case InsufficientInstanceCapacityErrCode:
		return errors.New(err, errors.ResourceExhausted, errMsg)

	case InternalErrorErrCode:
		fallthrough
	case InternalFailureErrCode:
		return errors.New(err, errors.Internal, errMsg)

	case ServiceUnavailableErrCode:
		fallthrough
	case UnavailableErrCode:
		return errors.New(err, errors.Unavailable, errMsg)
	case NoSuchKey:
		return errors.New(err, errors.NotFound, errMsg)
	}

	return errors.New(err, errors.Unknown, errMsg)
}
