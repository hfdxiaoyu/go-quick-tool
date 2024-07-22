package security

// 密码加密器工厂
type PasswordEncryptorFactory interface {
	CreatePasswordEncryptor() PasswordEncryptor
}

// 密码加密器
type PasswordEncryptor interface {
	Encryption(password string) string
	EncryptionBysalt(password, salt string) string
	Verify(password, encryptionPassword string) bool
	VerifyBySalt(password, salt, encryptionPassword string) bool
}
