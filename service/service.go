package service

import (
	"context"

	pb "github.com/hokita/grpc_practice/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService struct
type UserService struct {
	pb.UnimplementedUserManagerServer
}

// GetUser func
func (u *UserService) GetUser(ctx context.Context, message *pb.UserRequest) (*pb.User, error) {
	users := []*pb.User{
		{
			Id:   1,
			Name: "Wasabi",
		},
		{
			Id:   2,
			Name: "Karashi",
		},
	}

	var targetUser *pb.User
	for _, u := range users {
		if u.Id == message.Id {
			targetUser = u
		}
	}

	if targetUser == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return targetUser, nil
}
