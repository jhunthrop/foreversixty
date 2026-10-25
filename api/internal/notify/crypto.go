package notify

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"github.com/jhunthrop/foreversixty/api/internal/config"
)

const sealKeyLen = config.NotifyKeyLen

// Sealer encrypts webhook tokens at rest with AES-256-GCM. A box is the
// random nonce followed by the ciphertext.
type Sealer struct{ aead cipher.AEAD }

// NewSealer builds a Sealer from a key of exactly sealKeyLen bytes.
func NewSealer(key []byte) (*Sealer, error) {
	if len(key) != sealKeyLen {
		return nil, fmt.Errorf("notify: sealer key must be %d bytes", sealKeyLen)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("notify: cipher: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("notify: gcm: %w", err)
	}
	return &Sealer{aead: aead}, nil
}

// Seal encrypts plain under a fresh random nonce.
func (s *Sealer) Seal(plain string) ([]byte, error) {
	nonce := make([]byte, s.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("notify: nonce: %w", err)
	}
	return s.aead.Seal(nonce, nonce, []byte(plain), nil), nil
}

// Open decrypts a box; a box sealed under another key, or altered, fails.
func (s *Sealer) Open(box []byte) (string, error) {
	n := s.aead.NonceSize()
	if len(box) < n {
		return "", errors.New("notify: box too short")
	}
	plain, err := s.aead.Open(nil, box[:n], box[n:], nil)
	if err != nil {
		return "", errors.New("notify: box does not open")
	}
	return string(plain), nil
}
