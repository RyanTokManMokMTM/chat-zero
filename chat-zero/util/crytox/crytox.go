package crytox

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(password, salt string) string {
	pw, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", pw)
}
