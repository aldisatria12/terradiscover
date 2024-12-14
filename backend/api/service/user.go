package service

import (
	"context"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/aldisatria12/terradiscover/constants"
	"github.com/aldisatria12/terradiscover/dto"
	"github.com/aldisatria12/terradiscover/entity"
	"github.com/aldisatria12/terradiscover/repository"
	"github.com/aldisatria12/terradiscover/util"
)

type UserService interface {
	Login(ctx context.Context, user dto.UserLoginRequest) (dto.UserLoginResponse, error)
	Register(ctx context.Context, input dto.UserRegisterRequest) (dto.UserRegisterResponse, error)
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

func (s userService) Register(ctx context.Context, input dto.UserRegisterRequest) (dto.UserRegisterResponse, error) {
	txFunction := func(ds repository.DataStore) (any, error) {
		userRepo := ds.GetUserRepository()

		err := userRepo.IsEmailAvailable(ctx, entity.User{Email: input.Email})

		if err != nil {
			return dto.UserRegisterResponse{}, err
		}

		hashedPass, err := util.HashPassword(input.Password, constants.HashCost)

		if err != nil {
			return dto.UserRegisterResponse{}, apperror.NewError(err, apperror.ErrHashing)
		}

		newUser := dto.FromUserRegisterRequest(input)
		newUser.Password = string(hashedPass)

		userResult, err := userRepo.Register(ctx, newUser)
		if err != nil {
			return dto.UserRegisterResponse{}, err
		}

		result := dto.UserRegisterResponse{
			Email:    userResult.Email,
			Password: userResult.Password,
		}

		return result, nil
	}

	result, err := s.dataStore.StartTransaction(ctx, txFunction)

	if err != nil {
		return dto.UserRegisterResponse{}, err
	}

	parsedResult := result.(dto.UserRegisterResponse)

	return parsedResult, nil
}
