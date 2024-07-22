package security

import "testing"

func TestCreatePasswordEncryptor256(t *testing.T) {
	hash256Factory := new(Hash256EncryptorFactory)
	encryptor := hash256Factory.CreatePasswordEncryptor()
	encryptionPassword := encryptor.Encryption("123456")
	t.Log("encryptionPassword:", encryptionPassword)
	t.Log("verify:", encryptor.Verify("123456", encryptionPassword))
}

func TestEncryptor256BySalt(t *testing.T) {
	password := "123456"
	var salt Salt
	salt = NewRandomStrSalt()
	gSalt, err := salt.GenerateSalt(10)
	if err != nil {
		t.Fatal(err)
	}
	hash256Factory := new(Hash256EncryptorFactory)
	encryptor := hash256Factory.CreatePasswordEncryptor()
	epassword := encryptor.EncryptionBysalt(password, gSalt)
	t.Log("epassword:", epassword)
	t.Log("verify:", encryptor.VerifyBySalt(password, gSalt, epassword))
}
