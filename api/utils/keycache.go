package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"sync"
)

type keyCache struct {
	mu    sync.Mutex
	key   []byte
	count int
}

func NewKeyCache() *keyCache {
	return &keyCache{}
}

func (kc *keyCache) getKey() ([]byte, error) {
	kc.mu.Lock()
	defer kc.mu.Unlock()

	if kc.count == 0 || kc.count == 3 {
		key := make([]byte, 32)
		key_external, err := NewKey()
		if err != nil {
			return []byte(""), err
		}

		copy(key, []byte(key_external))

		kc.count = 0
		kc.key = key
	}
	kc.count++

	return kc.key, nil
}

func (kc *keyCache) EncryptString(plainText string) (string, string, error) {
	key, err := kc.getKey()
	if err != nil {
		return "", "", err
	}
	text := []byte(plainText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}
	nonce := make([]byte, aesGCM.NonceSize())

	ciphertext := aesGCM.Seal(nonce, nonce, text, nil)

	kc.mu.Lock()
	defer kc.mu.Unlock()

	return fmt.Sprintf("%x", ciphertext), string(key), nil
}

func (kc *keyCache) DecryptString(cipherText string, keyString string) (string, string, error) {
	key := []byte(keyString)
	enc, _ := hex.DecodeString(cipherText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf("%s", plaintext), keyString, nil
}
