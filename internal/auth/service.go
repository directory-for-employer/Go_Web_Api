package auth

import (
	"errors"
	"go/web-api/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Login(email, password string) (string, error) {
	existUser, err := service.UserRepository.FindByEmail(email)
	if existUser == nil {
		return "", errors.New(ErrUserNotFound)
	}
	err = bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrPasswordNotValid)
	}
	return existUser.Email, nil
}

func (service *AuthService) Register(users *user.User) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(users.Email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	userData := &user.User{
		Email:    users.Email,
		Password: string(hashedPassword),
		Name:     users.Name,
	}
	_, err = service.UserRepository.Create(userData)
	if err != nil {
		return "", err
	}
	return userData.Email, nil
}
