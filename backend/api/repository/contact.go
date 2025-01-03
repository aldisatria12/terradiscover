package repository

import (
	"context"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/aldisatria12/terradiscover/entity"
)

type ContactRepository interface {
	GetContact(ctx context.Context, userId int) ([]entity.Contact, error)
	InsertContact(ctx context.Context, input entity.Contact) error
	EditContact(ctx context.Context, input entity.Contact) error
	GetContactById(ctx context.Context, userId int) (entity.Contact, error)
}

type contactRepository struct {
	db DBTX
}

func NewContactRepository(db DBTX) contactRepository {
	return contactRepository{db: db}
}

func (r contactRepository) GetContact(ctx context.Context, userId int) ([]entity.Contact, error) {
	var contactList []entity.Contact

	query := `SELECT id, name, phone, email FROM contacts WHERE user_id = $1 AND deleted_at IS NULL;`

	rows, err := r.db.QueryContext(ctx, query, userId)

	if err != nil {
		return nil, apperror.NewError(err, apperror.ErrQuery)
	}

	for rows.Next() {
		var contact entity.Contact
		err := rows.Scan(&contact.Id, &contact.Name, &contact.Phone, &contact.Email)

		if err != nil {
			return nil, apperror.NewError(err, apperror.ErrQuery)
		}

		contactList = append(contactList, contact)
	}

	return contactList, nil
}

func (r contactRepository) GetContactById(ctx context.Context, userId int) (entity.Contact, error) {
	var contactList entity.Contact

	query := `SELECT id, name, phone, email FROM contacts WHERE id = $1 AND deleted_at IS NULL;`

	err := r.db.QueryRowContext(ctx, query, userId).Scan(&contactList.Id, &contactList.Name, &contactList.Phone, &contactList.Email)

	if err != nil {
		return entity.Contact{}, apperror.NewError(err, apperror.ErrQuery)
	}

	return contactList, nil
}

func (r contactRepository) InsertContact(ctx context.Context, input entity.Contact) error {
	var userId int
	query := "INSERT INTO contacts(user_id, name, phone, email) VALUES ($1, $2, $3, $4) RETURNING id;"
	err := r.db.QueryRowContext(ctx, query, input.UserId, input.Name, input.Phone, input.Email).Scan(&userId)

	if err != nil {
		return apperror.NewError(err, apperror.ErrQuery)
	}

	return nil
}

func (r contactRepository) EditContact(ctx context.Context, input entity.Contact) error {
	query := "UPDATE contacts SET name = $1, phone = $2, email = $3, updated_at = NOW() WHERE id = $4"
	_, err := r.db.ExecContext(ctx, query, input.Name, input.Phone, input.Email, input.Id)

	if err != nil {
		return apperror.NewError(err, apperror.ErrQuery)
	}

	return nil
}
