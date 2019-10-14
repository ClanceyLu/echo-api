package auth

import (
	"crypto/md5"
	"fmt"
)

// CryptoPassword 返回加密后的密码
func CryptoPassword(password, salt string) string {
	cryptoed := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	for i := 0; i < 3; i++ {
		cryptoed += salt
		cryptoed = fmt.Sprintf("%x", md5.Sum([]byte(cryptoed)))
	}
	return cryptoed
}
