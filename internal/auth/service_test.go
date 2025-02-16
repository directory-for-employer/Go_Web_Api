package auth_test

import (
	"go/web-api/internal/auth"
	"go/web-api/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{Email: "test@test.com"}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestAuthService_Register(t *testing.T) {
	const email = "test@test.com"
	authService := auth.NewAuthService(&MockUserRepository{})
	emailData, err := authService.Register(&user.User{Email: email, Password: "password", Name: "test"})
	if err != nil {
		t.Fatal(err)
	}
	if emailData != email {
		t.Fatalf("Email %s not equal %s", email, emailData)
	}
}
