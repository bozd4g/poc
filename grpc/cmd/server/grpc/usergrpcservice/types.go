package usergrpcservice

import (
	" github.com/bozd4g/poc/grpc/cmd/server/internal/application/userservice"
	pb " github.com/bozd4g/poc/grpc/pkg/proto/userservice"
	"github.com/sirupsen/logrus"
)

type UserGrpcService struct {
	pb.UnimplementedUserServiceServer

	Logger      logrus.Logger
	UserService userservice.IUserService
}
