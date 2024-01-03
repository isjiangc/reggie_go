package service

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/repository"
)

type AddressbookService interface {
	GetAddressbook(ctx context.Context, req *v1.GetAddressBookByUserIdRequest) (*[]v1.AddressBook, error)
}

func NewAddressbookService(service *Service, addressbookRepository repository.AddressbookRepository) AddressbookService {
	return &addressbookService{
		Service:               service,
		addressbookRepository: addressbookRepository,
	}
}

type addressbookService struct {
	*Service
	addressbookRepository repository.AddressbookRepository
}

func (a *addressbookService) GetAddressbook(ctx context.Context, req *v1.GetAddressBookByUserIdRequest) (*[]v1.AddressBook, error) {
	addressBooks, err := a.addressbookRepository.FirstById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var addressBookList []v1.AddressBook
	err = copier.Copy(&addressBookList, addressBooks)
	if err != nil {
		return nil, err
	}
	return &addressBookList, nil

}
