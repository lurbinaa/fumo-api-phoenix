package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"os"
)

func GenerateToken() (token string, tokenHash string, err error) {
	b := make([]byte, 32)
	_, err = rand.Read(b)
	if err != nil {
		return "", "", err
	}

	token = base64.URLEncoding.EncodeToString(b)
	tokenHash, err = HashToken(token)
	if err != nil {
		return "", "", err
	}

	return token, tokenHash, nil
}

func HashToken(token string) (string, error) {
	pepper := os.Getenv("TOKEN_PEPPER")
	if pepper == "" {
		return "", errors.New("Missing environment variable TOKEN_PEPPER")
	}

	mac := hmac.New(sha256.New, []byte(pepper))
	mac.Write([]byte(token))

	return base64.URLEncoding.EncodeToString(mac.Sum(nil)), nil
}
