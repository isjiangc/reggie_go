package service

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/repository"
	"time"
)

type AddressbookService interface {
	UpdataAddressIsDefault(ctx context.Context, req *v1.UpdateAddressBookIsDefaultRequest, updateTime time.Time, updateUser int64) error
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

func (a *addressbookService) UpdataAddressIsDefault(ctx context.Context, req *v1.UpdateAddressBookIsDefaultRequest, updateTime time.Time, updateUser int64) error {
	rowsAffected, err := a.addressbookRepository.UpdataAddressIsDefault(ctx, req.UserId, req.Id, updateTime, updateUser)
	if err != nil || rowsAffected < 0 {
		return err
	}
	return nil
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
