package security

import "testing"

func TestCreatePasswordEncryptor512(t *testing.T) {
	hash512Factory := new(Hash512EncryptorFactory)
	encryptor := hash512Factory.CreatePasswordEncryptor()
	encryptionPassword := encryptor.Encryption("123456")
	t.Log("encryptionPassword:", encryptionPassword)
	t.Log("verify:", encryptor.Verify("123456", encryptionPassword))
}

func TestEncryptor512BySalt(t *testing.T) {
	password := "123456"
	var salt Salt
	salt = NewRandomStrSalt()
	gSalt, err := salt.GenerateSalt(10)
	if err != nil {
		t.Fatal(err)
	}
	hash512Factory := new(Hash512EncryptorFactory)
	encryptor := hash512Factory.CreatePasswordEncryptor()
	epassword := encryptor.EncryptionBysalt(password, gSalt)
	t.Log("epassword:", epassword)
	t.Log("verify:", encryptor.VerifyBySalt(password, gSalt, epassword))
}
