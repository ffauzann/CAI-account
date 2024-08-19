package constant

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
)

// Known gRPC errors.
var (
	// Generic errors.
	ErrNotFound         = errors.New("Not found")
	ErrNoArg            = errors.New("No argument given")
	ErrInternal         = errors.New("Internal error")
	ErrPermissionDenied = errors.New("Permission denied")
	ErrUnauthenticated  = errors.New("Unauthenticated")

	// Specific errors.
	ErrInvalidMethod              = errors.New("Invalid method")
	ErrInvalidUsernamePassword    = errors.New("Invalid username/password")
	ErrPasswordIsTooWeak          = errors.New("Password is too weak")
	ErrMalformedEmail             = errors.New("Malformed email")
	ErrInvalidUserIdType          = errors.New("Invalid user ID type")
	ErrUserNotFound               = errors.New("User not found")
	ErrUserIsNotActive            = errors.New("User is blocked/closed")
	ErrUserAlreadyExists          = errors.New("User already exists")
	ErrInvalidToken               = errors.New("Invalid/expired token")
	ErrGroupAlreadyExists         = errors.New("Group already exists")
	ErrGroupNotFound              = errors.New("Group not found")
	ErrUnspecifiedAction          = errors.New("Unspecified Action")
	ErrInvalidOTP                 = errors.New("Invalid OTP")
	ErrAccountNotFound            = errors.New("Account not found")
	ErrInsufficientAccountBalance = errors.New("Insufficient account balance")
	ErrInvalidTransferAmount      = errors.New("Invalid transfer amount: must be in negative value")
)

// All client-safe errors goes here.
var (
	MapGRPCErrCodes = map[error]codes.Code{
		// For HTTP mapping: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
		ErrInvalidMethod:              codes.InvalidArgument,
		ErrInvalidUsernamePassword:    codes.InvalidArgument,
		ErrMalformedEmail:             codes.InvalidArgument,
		ErrInvalidUserIdType:          codes.InvalidArgument,
		ErrUnspecifiedAction:          codes.FailedPrecondition,
		ErrPasswordIsTooWeak:          codes.FailedPrecondition,
		ErrNoArg:                      codes.FailedPrecondition,
		ErrInsufficientAccountBalance: codes.FailedPrecondition,
		ErrInvalidTransferAmount:      codes.FailedPrecondition,
		ErrNotFound:                   codes.NotFound,
		ErrUserNotFound:               codes.NotFound,
		ErrGroupNotFound:              codes.NotFound,
		ErrAccountNotFound:            codes.NotFound,
		ErrUserAlreadyExists:          codes.AlreadyExists,
		ErrGroupAlreadyExists:         codes.AlreadyExists,
		ErrPermissionDenied:           codes.PermissionDenied,
		ErrUserIsNotActive:            codes.PermissionDenied,
		ErrInternal:                   codes.Internal,
		ErrInvalidToken:               codes.Unauthenticated,
		ErrInvalidOTP:                 codes.Unauthenticated,
		ErrUnauthenticated:            codes.Unauthenticated,
	}
)

// NewDynamicError covers an edge case where error message is not absolute but the code is.
func NewDynamicError(code codes.Code, msg string) error {
	return fmt.Errorf("DYNAMIC_ERR:%d:%s", code, msg)
}
