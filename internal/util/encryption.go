package util

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"golang.org/x/crypto/bcrypt"
)

func HashBCrypt(plainText string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	return string(hashedByte), err
}

func EncryptGCM(ctx context.Context, plaintext, key string) (result []byte, err error) {
	// Decode base64 key.
	bKey, err := base64.RawStdEncoding.DecodeString(key)
	if err != nil {
		LogContext(ctx).Error(err.Error())
		return
	}

	// Generate a new aes cipher using our 32 byte long key.
	c, err := aes.NewCipher(bKey)
	if err != nil {
		LogContext(ctx).Error(err.Error())
		return
	}

	// GCM or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		LogContext(ctx).Error(err.Error())
		return
	}

	// Creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		LogContext(ctx).Error(err.Error())
		return
	}

	return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
}
