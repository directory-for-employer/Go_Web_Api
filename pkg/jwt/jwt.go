package jwt

import "github.com/golang-jwt/jwt/v5"

type Jwt struct {
	Secret string
}

func NewJwt(secret string) *Jwt {
	return &Jwt{
		Secret: secret,
	}
}

func (j *Jwt) Create(email string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	s, err := t.SignedString([]byte(j.Secret))

	if err != nil {
		return "", err
	}

	return s, nil
}
