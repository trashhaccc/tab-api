package service

import (
	"errors"

	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user *model.User) (*model.User, error)
	Login(username, password string) (*model.User, error)
	GetUserByID(userID uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(nik string) error
	GetUserList() (*[]model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) DeleteUser(nik string) error {
	return s.userRepo.DeleteUser(nik)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *userService) GetUserList() (*[]model.User, error) {
	return s.userRepo.GetUserList()
}

func (s *userService) Register(user *model.User) (*model.User, error) {
	_, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if user.Role == "" {
		user.Role = "nasabah"
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

func (s *userService) GetUserByID(userID uint) (*model.User, error) {
	return s.userRepo.GetUserByID(userID)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}
