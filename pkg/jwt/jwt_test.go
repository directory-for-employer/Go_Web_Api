package jwt_test

import (
	"go/web-api/pkg/jwt"
	"testing"
)

func TestJWT_Create(t *testing.T) {
	const email = "test@test.com"
	jwtService := jwt.NewJwt("/z6yxgkL9GsjgfHGjrX4X3tL95lvluPtYOdiiGWxqMfkTyxK12j79/u9Ilm5MHjG")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid!")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", email, data.Email)
	}
}
