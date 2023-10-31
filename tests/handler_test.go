package handlers_test

import (
	"context"
	"testing"

	pb "github.com/atharv-bhadange/grpc-user-service/go-proto"
	"github.com/atharv-bhadange/grpc-user-service/handlers"
	"google.golang.org/grpc"
)

func TestGetUserByID(t *testing.T) {
	// create a new UserServiceServer instance
	s := &handlers.UserServiceServer{
		Users: map[int32]*pb.User{
			1: {
				Id:      1,
				Fname:   "Alice",
				City:    "Pune",
				Phone:   1234455,
				Height:  5.8,
				Married: false,
			},
			2: {
				Id:      2,
				Fname:   "Bob",
				City:    "Mumbai",
				Phone:   1234455,
				Height:  4.8,
				Married: false,
			},
		},
	}

	// create a new UserRequest instance
	req := &pb.UserRequest{
		Id: 1,
	}

	// call the GetUserByID function
	user, err := s.GetUserByID(context.Background(), req)

	// check if the returned user is correct
	if err != nil {
		t.Errorf("GetUserByID returned an error: %v", err)
		return
	}
	if user == nil {
		t.Errorf("GetUserByID returned nil user")
		return
	}
	if user.Id != 1 {
		t.Errorf("GetUserByID returned incorrect user ID: %d", user.Id)
	}
	if user.Fname != "Alice" {
		t.Errorf("GetUserByID returned incorrect user name: %s", user.Fname)
	}
	if user.City != "Pune" {
		t.Errorf("GetUserByID returned incorrect user city: %s", user.City)
	}
	if user.Phone != 1234455 {
		t.Errorf("GetUserByID returned incorrect user phone: %d", user.Phone)
	}
	if user.Height != 5.8 {
		t.Errorf("GetUserByID returned incorrect user height: %f", user.Height)
	}
	if user.Married != false {
		t.Errorf("GetUserByID returned incorrect user married status: %t", user.Married)
	}

	// test for a non-existent user
	req = &pb.UserRequest{
		Id: 3,
	}
	user, err = s.GetUserByID(context.Background(), req)
	if err == nil {
		t.Errorf("GetUserByID did not return an error for non-existent user")
	}
	if user != nil {
		t.Errorf("GetUserByID returned non-nil user for non-existent user")
	}
}

type mockUserStream struct {
	grpc.ServerStream
	users []*pb.User
}

func (m *mockUserStream) Send(user *pb.User) error {
	m.users = append(m.users, user)
	return nil
}

func (m *mockUserStream) Users() []*pb.User {
	return m.users
}

func TestGetUsersByIds(t *testing.T) {
	// create a new UserServiceServer instance
	s := &handlers.UserServiceServer{
		Users: map[int32]*pb.User{
			1: {
				Id:      1,
				Fname:   "Alice",
				City:    "Pune",
				Phone:   1234455,
				Height:  5.8,
				Married: false,
			},
			2: {
				Id:      2,
				Fname:   "Bob",
				City:    "Mumbai",
				Phone:   1234455,
				Height:  4.8,
				Married: false,
			},
			3: {
				Id:      3,
				Fname:   "Charlie",
				City:    "Delhi",
				Phone:   1234455,
				Height:  6.0,
				Married: true,
			},
		},
	}

	// check if the returned users are correct
	req := &pb.UserListRequest{
		Ids: []int32{1, 3},
	}
	stream := &mockUserStream{}
	err := s.GetUsersByIds(req, stream)
	if err != nil {
		t.Errorf("GetUsersByIds returned an error: %v", err)
		return
	}
	users := stream.Users()

	if len(users) != 2 {
		t.Errorf("GetUsersByIds returned incorrect number of users: %d", len(users))
		return
	}
	if users[0].Id != 1 {
		t.Errorf("GetUsersByIds returned incorrect user ID: %d", users[0].Id)
	}
	if users[0].Fname != "Alice" {
		t.Errorf("GetUsersByIds returned incorrect user name: %s", users[0].Fname)
	}
	if users[0].City != "Pune" {
		t.Errorf("GetUsersByIds returned incorrect user city: %s", users[0].City)
	}
	if users[0].Phone != 1234455 {
		t.Errorf("GetUsersByIds returned incorrect user phone: %d", users[0].Phone)
	}
	if users[0].Height != 5.8 {
		t.Errorf("GetUsersByIds returned incorrect user height: %f", users[0].Height)
	}
	if users[0].Married != false {
		t.Errorf("GetUsersByIds returned incorrect user married status: %t", users[0].Married)
	}
	if users[1].Id != 3 {
		t.Errorf("GetUsersByIds returned incorrect user ID: %d", users[1].Id)
	}
	if users[1].Fname != "Charlie" {
		t.Errorf("GetUsersByIds returned incorrect user name: %s", users[1].Fname)
	}
	if users[1].City != "Delhi" {
		t.Errorf("GetUsersByIds returned incorrect user city: %s", users[1].City)
	}
	if users[1].Phone != 1234455 {
		t.Errorf("GetUsersByIds returned incorrect user phone: %d", users[1].Phone)
	}
	if users[1].Height != 6.0 {
		t.Errorf("GetUsersByIds returned incorrect user height: %f", users[1].Height)
	}
	if users[1].Married != true {
		t.Errorf("GetUsersByIds returned incorrect user married status: %t", users[1].Married)
	}

	// test for an empty user list
	req = &pb.UserListRequest{
		Ids: []int32{},
	}
	stream = &mockUserStream{}
	err = s.GetUsersByIds(req, stream)
	if err == nil {
		t.Errorf("GetUsersByIds did not return an error for empty user list")
	}
}
