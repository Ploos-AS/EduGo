package authtoken

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

var ErrInvalid = errors.New("invalid token")

func New(key []byte) (string, error) {
	if len(key) < 16 {
		return "", errors.New("key too short")
	}

	payload := make([]byte, 24)
	if _, err := rand.Read(payload); err != nil {
		return "", err
	}

	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(payload)
	sig := mac.Sum(nil)

	raw := append(payload, sig...)
	return base64.RawURLEncoding.EncodeToString(raw), nil
}

func Verify(key []byte, token string) error {
	raw, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return ErrInvalid
	}
	const payloadSize = 24
	if len(raw) != payloadSize+sha256.Size {
		return ErrInvalid
	}

	payload := raw[:payloadSize]
	got := raw[payloadSize:]
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(payload)
	if !hmac.Equal(got, mac.Sum(nil)) {
		return ErrInvalid
	}
	return nil
}
