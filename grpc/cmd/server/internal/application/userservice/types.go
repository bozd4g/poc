package userservice

import (
	" github.com/bozd4g/poc/grpc/cmd/server/internal/infrastructure/repository/userrepository"
	"github.com/google/uuid"
)

type IUserService interface {
	Create(entity UserCreateRequestDto) error
	GetAll() ([]UserDto, error)
	Get(id uuid.UUID) (*UserDto, error)
}

type UserService struct {
	repository userrepository.IUserRepository
}
