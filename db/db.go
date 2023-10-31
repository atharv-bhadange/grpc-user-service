package db

import (
	"sync"

	pb "github.com/atharv-bhadange/grpc-user-service/go-proto"
	"github.com/atharv-bhadange/grpc-user-service/handlers"
)

var once sync.Once

func GetUserServiceServer() *handlers.UserServiceServer {
	var instance *handlers.UserServiceServer
	once.Do(func() {
		instance = &handlers.UserServiceServer{
			Users: map[int32]*pb.User{
				1: {
					Id:      1,
					Fname:   "Atharv",
					City:    "Pune",
					Phone:   1234455,
					Height:  5.8,
					Married: false,
				},
				2: {
					Id:      2,
					Fname:   "Rahul",
					City:    "Mumbai",
					Phone:   1234455,
					Height:  4.8,
					Married: false,
				},
				3: {
					Id:      3,
					Fname:   "Rohit",
					City:    "Delhi",
					Phone:   1234455,
					Height:  5.2,
					Married: true,
				},
				4: {
					Id:      4,
					Fname:   "Raj",
					City:    "Banglore",
					Phone:   1234455,
					Height:  6.2,
					Married: false,
				},
			},
		}
	})
	return instance
}
