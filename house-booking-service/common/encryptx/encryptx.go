package encryptx

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func EncryptPassword(salt []byte, password []byte) string {
	pw, _ := scrypt.Key(password, salt, 32768, 8, 1, 32)
	return fmt.Sprintf("%x", pw)
}
