package service

import (
	"context"

	v1 "reggie_go/api/v1"
	"reggie_go/internal/model"
	"reggie_go/internal/repository"
)

type UsersService interface {
	UserLogin(ctx context.Context, req *v1.UserLoginRequest) (*v1.Users, error)
}

func NewUsersService(service *Service, usersRepository repository.UsersRepository) UsersService {
	return &usersService{
		Service:         service,
		usersRepository: usersRepository,
	}
}

type usersService struct {
	*Service
	usersRepository repository.UsersRepository
}

func (s *usersService) UserLogin(ctx context.Context, req *v1.UserLoginRequest) (*v1.Users, error) {
	count, err := s.usersRepository.FirstCountByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		users, err := s.usersRepository.GetUserByPhone(ctx, req.Phone)
		if err != nil {
			return nil, err
		} else {
			return &v1.Users{
				Phone:  users.Phone,
				Status: users.Status,
			}, err
		}
	} else {
		ret, err := s.usersRepository.SaveUser(ctx, model.Users{
			Phone:  req.Phone,
			Status: 1,
		})
		if err != nil {
			return nil, err
		}
		if ret > 0 {
			return &v1.Users{
				Phone:  req.Phone,
				Status: 1,
			}, nil
		}
	}
	return &v1.Users{
		Phone:  req.Phone,
		Status: 1,
	}, err
}
