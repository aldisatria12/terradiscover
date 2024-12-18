package dto

import "github.com/aldisatria12/terradiscover/entity"

type ContactResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func ToContactResponse(contact entity.Contact) ContactResponse {
	return ContactResponse{
		Id:    contact.Id,
		Name:  contact.Name,
		Phone: contact.Phone,
		Email: contact.Email,
	}
}

type NewContactRequest struct {
	Name  string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Email string `json:"email" validate:"email"`
}

func FromNewContactRequest(request NewContactRequest) entity.Contact {
	return entity.Contact{
		Name:  request.Name,
		Phone: request.Phone,
		Email: request.Email,
	}
}
