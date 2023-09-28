package interfaces

import (
	pb "Clean-Grpc/pkg/api/proto"
	"context"
)

type UserHandler interface{
	Register(ctx context.Context,req *pb.RegisterRequest)(*pb.ResgisterResponse)
	Login(ctx context.Context,req *pb.RegisterRequest)(*pb.ResgisterResponse)
}