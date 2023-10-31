package handlers

import (
	"context"

	pb "github.com/atharv-bhadange/grpc-user-service/go-proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	Users map[int32]*pb.User
}

func (c *UserServiceServer) GetUserByID(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	user, ok := c.Users[req.GetId()]

	if !ok {
		return nil, status.Errorf(codes.NotFound, "User with id %d not found", req.GetId())
	}

	return user, nil
}

func (c *UserServiceServer) GetUsersByIds(req *pb.UserListRequest, stream pb.UserService_GetUsersByIdsServer) error {

	if len(req.Ids) < 1 {
		return status.Errorf(codes.InvalidArgument, "user IDs are required")
	}

	for _, id := range req.GetIds() {
		user, ok := c.Users[id]

		if !ok {
			return status.Errorf(codes.NotFound, "User with id %d not found", id)
		}

		stream.Send(user)
	}

	return nil
}
