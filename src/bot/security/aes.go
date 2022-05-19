package security

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/fsufitch/tagioalisi-bot/config"
)

// AESSupport encapsulates encrypting/decrypting data using AES
type AESSupport struct {
	Block config.AESBlock
}

// Ready returns true if a cipher block is available
func (a AESSupport) Ready() bool {
	return a.Block != nil
}

// Encrypt receives some raw byte data and returns encrypted data; the first block of the return is the IV
func (a AESSupport) Encrypt(data []byte) ([]byte, error) {
	iv := a.IV()
	stream := cipher.NewOFB(a.Block, iv)
	dest := bytes.NewBuffer(iv)
	reader := bytes.NewReader(data)

	writer := cipher.StreamWriter{S: stream, W: dest}
	if _, err := io.Copy(writer, reader); err != nil {
		return nil, err
	}

	return dest.Bytes(), nil
}

// Decrypt receives some raw encrypted data and decrypts it; the first block of the input is expected to be the IV
func (a AESSupport) Decrypt(data []byte) ([]byte, error) {
	iv := data[:a.Block.BlockSize()]
	data = data[a.Block.BlockSize():]

	stream := cipher.NewOFB(a.Block, iv)
	dest := bytes.Buffer{}
	encReader := bytes.NewReader(data)

	reader := cipher.StreamReader{S: stream, R: encReader}
	if _, err := io.Copy(&dest, reader); err != nil {
		return nil, err
	}

	return dest.Bytes(), nil
}

// IV generates a new cryptographically secure IV
func (a AESSupport) IV() []byte {
	iv := make([]byte, a.Block.BlockSize())
	rand.Read(iv)
	return iv
}
