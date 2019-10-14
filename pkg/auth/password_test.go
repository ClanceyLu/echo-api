package auth

import "testing"

func TestCryptoPassword(t *testing.T) {
	password := "123456abc"
	salt := "salt"
	cryptoed := CryptoPassword(password, salt)
	expectResult := "2f75698f2ff26956ba0b01cc42e9639b"
	if cryptoed != expectResult {
		t.Errorf("加密错误，预计得到 %s, 但是得到 %s", expectResult, cryptoed)
	}
}
