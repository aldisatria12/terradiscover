package repository

import (
	"context"

	"github.com/aldisatria12/terradiscover/apperror"
	"github.com/aldisatria12/terradiscover/entity"
)

type ContactRepository interface {
	GetContact(ctx context.Context, userId int) ([]entity.Contact, error)
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
