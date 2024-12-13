package repository

import (
	"context"
	"database/sql"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/aldisatria12/terradiscover/entity"
)

type UserRepository interface {
	Login(ctx context.Context, input entity.User) (entity.User, error)
}

type userRepository struct {
	db DBTX
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

func NewUserRepository(db DBTX) userRepository {
	return userRepository{db: db}
}
