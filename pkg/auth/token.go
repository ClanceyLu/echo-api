package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// CryptoToken 返回加密后的 token
func CryptoToken(payload map[string]interface{},
	secret string) (string, error) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerBase64 := base64.RawURLEncoding.EncodeToString([]byte(headerJSON))
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	payloadBase64 := base64.RawURLEncoding.EncodeToString(payloadJSON)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%s.%s", headerBase64, payloadBase64)))
	signature := base64.RawURLEncoding.EncodeToString([]byte(mac.Sum(nil)))

	return fmt.Sprintf("%s.%s.%s",
		headerBase64, payloadBase64, signature), nil
}

// VertifyToken 验证 token 是否有效并返回 payload 信息
func VertifyToken(token, secret string) (map[string]interface{}, error) {
	data := strings.Split(token, ".")
	if len(data) != 3 {
		return nil, errors.New("token 不合法")
	}

	// 判断 token 签名是否正确
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%s.%s", data[0], data[1])))
	signature := base64.RawURLEncoding.EncodeToString([]byte(mac.Sum(nil)))
	if signature != data[2] {
		return nil, errors.New("token 不合法")
	}

	payloadJSON, err := base64.RawURLEncoding.DecodeString(data[1])
	if err != nil {
		return nil, err
	}
	payload := make(map[string]interface{})
	json.Unmarshal(payloadJSON, &payload)

	// 判断 token 是否过期
	if exp, ok := payload["exp"].(float64); ok {
		expire := time.Unix(int64(exp), 0)
		now := time.Now()
		if expire.Before(now) {
			return nil, errors.New("token 已过期")
		}
	}

	return payload, nil
}
