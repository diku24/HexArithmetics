package rpc

import (
	"HexArithmatics/internal/adapters/framework/left/grpc/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const MISSING_VALUE_ERROR = "missing required"
const UNEXPECTED_ERROR = "Unexpected Error"

// GetAddtion gets the result of adding parameters a and b
func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, MISSING_VALUE_ERROR)

	}
	answer, err := grpca.api.GetAddtion(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, UNEXPECTED_ERROR)

	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

// GetSubstraction gets the result of subtracting parameters a and b
func (grpca Adapter) GetSubstraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, MISSING_VALUE_ERROR)

	}
	answer, err := grpca.api.GetSubstraction(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, UNEXPECTED_ERROR)

	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

// GetMultiplication gets the result of Multiplying parameters a and b
func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, MISSING_VALUE_ERROR)

	}
	answer, err := grpca.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, UNEXPECTED_ERROR)

	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

// GetDivision gets the result of dividing parameters a and b
func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, MISSING_VALUE_ERROR)

	}
	answer, err := grpca.api.GetDivision(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, UNEXPECTED_ERROR)

	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}
