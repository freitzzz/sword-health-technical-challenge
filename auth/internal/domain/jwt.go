package domain

import (
	"os"

	"github.com/golang-jwt/jwt"
)

const (
	jwtAlgEnvKey    = "jwt_alg"
	jwtSecretEnvKey = "jwt_secret"
)

type JWTBundle struct {
	Alg    string
	Secret string
}

func LoadJWTBundle() JWTBundle {

	alg := os.Getenv(jwtAlgEnvKey)

	secret := os.Getenv(jwtSecretEnvKey)

	return JWTBundle{
		Alg:    alg,
		Secret: secret,
	}

}

func SignUserSession(b JWTBundle, u User, exp int64) (string, error) {

	key := []byte(b.Secret)

	return jwt.NewWithClaims(
		jwt.GetSigningMethod(b.Alg),
		jwt.MapClaims{
			"uid":       u.Identifier,
			"secret":    u.Secret,
			"expiresAt": exp,
		},
	).SignedString(key)

}
