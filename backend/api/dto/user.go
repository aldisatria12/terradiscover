package dto

import "github.com/aldisatria12/terradiscover/entity"

type UserLoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

func FromUserLoginRequest(user UserLoginRequest) entity.User {
	return entity.User{
		Email:    user.Email,
		Password: user.Password,
	}
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
