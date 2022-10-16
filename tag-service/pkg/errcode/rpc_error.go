package errcode

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "tag-service/proto"
)

type Status struct {
	*status.Status
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

// func FromError(err error) (s *Status, ok bool) {
// 	if err == nil {
// 		return nil, true
// 	}
// 	if se, ok := err.(interface { // TODO: 类型断言还可以使用匿名类型,
// 	熟悉这种用法
// 		GRPCStatus() *Status
// 	}); ok {
// 		return se.GRPCStatus(), true
// 	}
// 	return New(codes.Unknown, err.Error()), false
// }

func TogRPCError(err *Error) error {
	pbErr := &pb.Error{Code: int32(err.Code()), Message: err.Msg()}
	s, _ := status.New(ToRPCCode(err.Code()), err.Msg()).WithDetails(pbErr)
	return s.Err()
}

func ToRPCStatus(code int, msg string) *Status {
	pbErr := &pb.Error{Code: int32(code), Message: msg}
	s, _ := status.New(ToRPCCode(code), msg).WithDetails(pbErr)
	return &Status{s}
}

// TODO: 熟悉 gRPC 常用状态码
func ToRPCCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case Fail.Code():
		statusCode = codes.Internal
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case Unauthorized.Code():
		statusCode = codes.Unauthenticated
	case AccessDenied.Code():
		statusCode = codes.PermissionDenied
	case DeadlineExceeded.Code():
		statusCode = codes.DeadlineExceeded
	case NotFound.Code():
		statusCode = codes.NotFound
	case LimitExceed.Code():
		statusCode = codes.ResourceExhausted
	case MethodNotAllowed.Code():
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknown
	}
	return statusCode
}
