package service

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/model"
	"reggie_go/internal/repository"
	"time"
)

type AddressbookService interface {
	SaveAddressBook(ctx context.Context, req *v1.SaveAddressBookRequest) error
	GetAddressById(ctx context.Context, req *v1.GetAddressBookByIdRequest) (*v1.AddressBook, error)
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

func (a *addressbookService) SaveAddressBook(ctx context.Context, req *v1.SaveAddressBookRequest) error {
	ret, err := a.addressbookRepository.SaveAddressBook(ctx, model.AddressBook{
		UserId:     req.UserId,
		Consignee:  req.Consignee,
		Sex:        req.Sex,
		Phone:      req.Phone,
		Detail:     req.Detail,
		Label:      req.Label,
		IsDefault:  0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		CreateUser: req.CreateUser,
		UpdateUser: req.UpdateUser,
		IsDeleted:  0,
	})
	if err != nil || ret < 0 {
		return err
	}
	return nil
}

func (a *addressbookService) GetAddressById(ctx context.Context, req *v1.GetAddressBookByIdRequest) (*v1.AddressBook, error) {
	addressBook, err := a.addressbookRepository.QueryAddressById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	var v1AddressBook = v1.AddressBook{}
	err = copier.Copy(&v1AddressBook, addressBook)
	if err != nil {
		return nil, err
	}
	return &v1AddressBook, nil
}

func (a *addressbookService) UpdataAddressIsDefault(ctx context.Context, req *v1.UpdateAddressBookIsDefaultRequest, updateTime time.Time, updateUser int64) error {
	rowsAffected, err := a.addressbookRepository.UpdataAddressIsDefault(ctx, req.UserId, req.Id, updateTime, updateUser)
	if err != nil || rowsAffected < 0 {
		return err
	}
	return nil
}

func (a *addressbookService) GetAddressbook(ctx context.Context, req *v1.GetAddressBookByUserIdRequest) (*[]v1.AddressBook, error) {
	addressBooks, err := a.addressbookRepository.FirstByUserId(ctx, req.UserId)
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
