package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

// Previously generated password hash have the parameters encoded in the hash string.
// this means those parameters will not be changed retroactively.
//
// Rehashing could be a possible improvement, when the parameters are changed.
const (
	memory      = 64 * 1024
	iterations  = 3
	parallelism = 2
	saltLength  = 16
	keyLength   = 32
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a AuthService) GenerateSalt() ([]byte, error) {
	b := make([]byte, saltLength)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a AuthService) HashPassword(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, iterations, parallelism, b64Salt, b64Hash)
}
