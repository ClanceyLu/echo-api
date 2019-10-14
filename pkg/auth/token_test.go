package auth

import (
	"reflect"
	"testing"
)

func TestCryptoToken(t *testing.T) {
	payload := map[string]interface{}{
		"name": "John Doe",
		"iat":  1516239022,
	}
	secret := "secret"
	token, err := CryptoToken(payload, secret)
	if err != nil {
		t.Error(err)
	}
	expectToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjIsIm5hbWUiOiJKb2huIERvZSJ9.DydTqdoRcXVw6LIJaE7qEWj5iIbdzKvVM0QVzFt_gkA"
	if token != expectToken {
		t.Errorf("期望到的 %s, 但是得到 %s", expectToken, token)
	}
}

func TestVerifyToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjIsIm5hbWUiOiJKb2huIERvZSJ9.DydTqdoRcXVw6LIJaE7qEWj5iIbdzKvVM0QVzFt_gkA"
	secret := "secret"
	payload, err := VertifyToken(token, secret)
	if err != nil {
		t.Error(err)
	}
	expectPayload := map[string]interface{}{
		"name": "John Doe",
		"iat":  1.516239022e+09,
	}
	if !reflect.DeepEqual(payload, expectPayload) {
		t.Errorf("期望得到 %+v, 但是得到 %+v", expectPayload, payload)
	}
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjIsIm5hbWUiOiJKb2huIERvZSJ9.DydTqdoRcXVw6LIJaE7qEWj5iIbdzKvVM0QVzFt_gkb"
	payload, err = VertifyToken(token, secret)
	if err == nil {
		t.Errorf("期望得到错误 token 不合法")
	}
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzQzMDgsImlhdCI6MTUxNjIzOTAyMiwibmFtZSI6IkpvaG4gRG9lIn0.P2xHRrWFiRTvSXF4Kknjc2RTuAg441jmifVfpr1Gbwg"
	payload, err = VertifyToken(token, secret)
	if err == nil {
		t.Errorf("期望得到错误 token 已过期")
	}
}
