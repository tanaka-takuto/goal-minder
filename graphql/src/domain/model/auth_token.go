package model

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const secreatKey = "secsec"

// AuthToken 認証トークン
type AuthToken string

// Decode 認証トークンをデコードする
func (at AuthToken) Decode() (*AuthTokenClaims, error) {
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
	AccountID AccountID
}

// Encode 認証トークンクレームをエンコードする
func (atc AuthTokenClaims) Encode() AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	tokenJwt, err := token.SignedString([]byte(secreatKey))
	if err != nil {
		panic(err)
	}

	return AuthToken(tokenJwt)
}

// NewAuthTokenClaims 認証トークンクレームを作成する
func NewAuthTokenClaims(accountID AccountID) AuthTokenClaims {
	claims := AuthTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "asdf",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        "asdf",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "asdf",
			NotBefore: time.Now().Unix(),
			Subject:   "asdf",
		},
		AccountID: accountID,
	}

	return claims

}
