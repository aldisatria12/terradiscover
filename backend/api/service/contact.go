package service

import (
	"context"

	"github.com/aldisatria12/terradiscover/dto"
	"github.com/aldisatria12/terradiscover/repository"
)

type ContactService interface {
	GetContact(ctx context.Context, userId int) ([]dto.ContactResponse, error)
}

type contactService struct {
	dataStore         repository.DataStore
	contactRepository repository.ContactRepository
}

func NewContactService(ds repository.DataStore, cr repository.ContactRepository) contactService {
	return contactService{
		dataStore:         ds,
		contactRepository: cr,
	}
}

func (s contactService) GetContact(ctx context.Context, userId int) ([]dto.ContactResponse, error) {
	contactList, err := s.contactRepository.GetContact(ctx, userId)

	if err != nil {
		return nil, err
	}

	var result []dto.ContactResponse

	for _, contact := range contactList {
		result = append(result, dto.ToContactResponse(contact))
	}

	return result, nil
}
