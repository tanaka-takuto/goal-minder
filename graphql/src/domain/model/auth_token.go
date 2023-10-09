package model

import (
	"fmt"
	"time"

	"goal-minder/config"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// AuthToken 認証トークン
type AuthToken string

// Decode 認証トークンをデコードする
func (at AuthToken) Decode() (*AuthTokenClaims, error) {
	secreatKey := config.SecretKey()

	token, err := jwt.ParseWithClaims(string(at), &AuthTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secreatKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AuthTokenClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

// AuthTokenClaims 認証トークンクレーム
type AuthTokenClaims struct {
	jwt.StandardClaims
	Payload AuthTokenPayload
}

// AuthTokenPayload 認証トークンペイロード
type AuthTokenPayload struct {
	AccountID AccountID
}

// Encode 認証トークンクレームをエンコードする
func (atc AuthTokenClaims) Encode() AuthToken {
	secreatKey := config.SecretKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	tokenJwt, err := token.SignedString([]byte(secreatKey))
	if err != nil {
		panic(err)
	}

	return AuthToken(tokenJwt)
}

// NewAuthTokenClaims 認証トークンクレームを作成する
func NewAuthTokenClaims(accountID AccountID) AuthTokenClaims {
	uuid, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	claims := AuthTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "goal-minder",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        uuid.String(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "goal-minder",
			NotBefore: time.Now().Unix(),
			Subject:   "AuthToken",
		},
		Payload: AuthTokenPayload{
			AccountID: accountID,
		},
	}

	return claims

}
