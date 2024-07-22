package security

import (
	"crypto/sha256"
	"encoding/hex"
)

type hash256Encryptor struct {
}

func (h *hash256Encryptor) Encryption(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func (h *hash256Encryptor) EncryptionBysalt(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

func (h *hash256Encryptor) Verify(password, encryptionPassword string) bool {
	if h.Encryption(password) == encryptionPassword {
		return true
	}
	return false
}

func (h *hash256Encryptor) VerifyBySalt(password, salt, encryptionPassword string) bool {
	if h.EncryptionBysalt(password, salt) == encryptionPassword {
		return true
	}
	return false
}

type Hash256EncryptorFactory struct {
}

func (h *Hash256EncryptorFactory) CreatePasswordEncryptor() PasswordEncryptor {
	encryptor := new(hash256Encryptor)
	return encryptor
}
