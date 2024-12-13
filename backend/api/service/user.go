package service

import (
	"context"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/aldisatria12/terradiscover/dto"
	"github.com/aldisatria12/terradiscover/repository"
	"github.com/aldisatria12/terradiscover/util"
)

type UserService interface {
	Login(ctx context.Context, user dto.UserLoginRequest) (dto.UserLoginResponse, error)
}

type userService struct {
	dataStore      repository.DataStore
	userRepository repository.UserRepository
}

func NewUserService(ds repository.DataStore, ur repository.UserRepository) userService {
	return userService{
		dataStore:      ds,
		userRepository: ur,
	}
}

func (s userService) Login(ctx context.Context, user dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	loggedUser, err := s.userRepository.Login(ctx, dto.FromUserLoginRequest(user))

	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	isPassRight, err := util.CheckPassword(user.Password, []byte(loggedUser.Password))

	if !isPassRight {
		return dto.UserLoginResponse{}, apperror.NewError(err, apperror.ErrUsernameNotFound)
	}

	if err != nil {
		return dto.UserLoginResponse{}, apperror.NewError(err, apperror.ErrHashing)
	}

	token, err := util.CreateAndSign(loggedUser.Id)

	if err != nil {
		return dto.UserLoginResponse{}, apperror.NewError(err, apperror.ErrHashing)
	}

	return dto.UserLoginResponse{Token: token}, nil
}
