package security

import (
	"crypto/sha512"
	"encoding/hex"
)

type hash512Encryptor struct {
}

func (h *hash512Encryptor) Encryption(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func (h *hash512Encryptor) EncryptionBysalt(password, salt string) string {
	hash := sha512.New()
	hash.Write([]byte(password + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

func (h *hash512Encryptor) Verify(password, encryptionPassword string) bool {
	if h.Encryption(password) == encryptionPassword {
		return true
	}
	return false
}

func (h *hash512Encryptor) VerifyBySalt(password, salt, encryptionPassword string) bool {
	if h.EncryptionBysalt(password, salt) == encryptionPassword {
		return true
	}
	return false
}

type Hash512EncryptorFactory struct {
}

func (h *Hash512EncryptorFactory) CreatePasswordEncryptor() PasswordEncryptor {
	encryptor := new(hash256Encryptor)
	return encryptor
}
