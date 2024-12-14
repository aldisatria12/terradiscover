package repository

import (
	"context"
	"database/sql"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/aldisatria12/terradiscover/entity"
)

type UserRepository interface {
	Login(ctx context.Context, input entity.User) (entity.User, error)
	Register(ctx context.Context, input entity.User) (entity.User, error)
	IsEmailAvailable(ctx context.Context, user entity.User) error
}

type userRepository struct {
	db DBTX
}

func NewUserRepository(db DBTX) userRepository {
	return userRepository{db: db}
}

func (r userRepository) Login(ctx context.Context, input entity.User) (entity.User, error) {
	var getUser entity.User

	query := `SELECT id, password FROM users WHERE username = $1;`

	err := r.db.QueryRowContext(ctx, query, input.Username).Scan(&getUser.Id, &getUser.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, apperror.NewError(err, apperror.ErrUsernameNotFound)
		}
		return entity.User{}, apperror.NewError(err, apperror.ErrQuery)
	}

	return getUser, nil
}

func (r userRepository) Register(ctx context.Context, input entity.User) (entity.User, error) {
	result := input
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) RETURNING id;"
	err := r.db.QueryRowContext(ctx, query, input.Username, input.Password, input.Email).Scan()

	if err != nil {
		return entity.User{}, apperror.NewError(err, apperror.ErrQuery)
	}

	return result, nil
}

func (r userRepository) IsEmailAvailable(ctx context.Context, user entity.User) error {
	var getEmail entity.User

	query := `SELECT u.email FROM users u WHERE u.email = $1;`

	err := r.db.QueryRowContext(ctx, query, user.Email).Scan(&getEmail.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return apperror.NewError(err, apperror.ErrQuery)
	}

	return apperror.NewError(err, apperror.ErrRegisteredEmail)
}
