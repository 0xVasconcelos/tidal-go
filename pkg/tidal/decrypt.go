package tidal

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

var masterKey, _ = hex.DecodeString("5089534C432698B7C6A30A3F502EB4C761F8E56E8C74681345FA3FBA6838EF9E")

func (m *Media) GetKey() error {

	encryptionKey, err := m.Base64Decode(m.EncryptionKey)

	if err != nil {
		return err
	}

	block, err := aes.NewCipher(masterKey)

	if err != nil {
		return err
	}

	iv := encryptionKey[:16]
	data := encryptionKey[16:]

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(data, data)

	m.Key = data[:16]
	m.Nonce = data[16:24]

	return nil
}

func (m *Media) DecryptMedia() error {

	block, err := aes.NewCipher(m.Key)

	if err != nil {
		return err
	}

	iv := append(m.Nonce, make([]byte, aes.BlockSize-len(m.Nonce))...)

	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(m.Media, m.Media)

	return nil
}

func (m *Media) Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func (m *Media) Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
