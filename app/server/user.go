package server

import (
	"context"
	"github.com/mf-sakura/bh_user/app/db"
	hpb "github.com/mf-sakura/bh_user/app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type UserServiceServerImpl struct {
}

func (u *UserServiceServerImpl) GetUser(ctx context.Context, req *hpb.GetUserMessage) (*hpb.GetUserResponse, error) {
	return nil, nil
}

func (u *UserServiceServerImpl) RegistUser(ctx context.Context, req *hpb.RegistUserMessage) (*hpb.RegistUserResponse, error) {
	return nil, nil
}
func (u *UserServiceServerImpl) IncrUserCounter(ctx context.Context, req *hpb.IncrUserCounterMessage) (*hpb.IncrUserCounterResponse, error) {
	if err := db.IncrementUserReservationCount(req.UserId); err != nil {
		return nil, grpc.Errorf(codes.Internal, "db.IncrementUserReservationCount failed:%v", err)
	}
	return &hpb.IncrUserCounterResponse{}, nil
}
func (u *UserServiceServerImpl) DecrUserCounter(ctx context.Context, req *hpb.DecrUserCounterMessage) (*hpb.DecrUserCounterResponse, error) {
	if err := db.DecrementUserReservationCount(req.UserId); err != nil {
		return nil, grpc.Errorf(codes.Internal, "db.DecrementUserReservationCount failed:%v", err)
	}
	return &hpb.DecrUserCounterResponse{}, nil
}
