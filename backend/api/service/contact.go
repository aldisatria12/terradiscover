package service

import (
	"context"

	"github.com/aldisatria12/terradiscover/dto"
	"github.com/aldisatria12/terradiscover/repository"
)

type ContactService interface {
	GetContact(ctx context.Context, userId int) ([]dto.ContactResponse, error)
	InsertContact(ctx context.Context, input dto.NewContactRequest, id int) error
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

func (s contactService) InsertContact(ctx context.Context, input dto.NewContactRequest, id int) error {
	txFunction := func(ds repository.DataStore) (any, error) {
		contactRepo := ds.GetContactRepository()

		newContact := dto.FromNewContactRequest(input)
		newContact.UserId = id

		err := contactRepo.InsertContact(ctx, newContact)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := s.dataStore.StartTransaction(ctx, txFunction)

	if err != nil {
		return err
	}

	return nil
}

func (s contactService) EditContact(ctx context.Context, input dto.EditContactRequest) error {
	txFunction := func(ds repository.DataStore) (any, error) {
		contactRepo := ds.GetContactRepository()

		editContact := dto.FromEditContactRequest(input)

		err := contactRepo.EditContact(ctx, editContact)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := s.dataStore.StartTransaction(ctx, txFunction)

	if err != nil {
		return err
	}

	return nil
}
