package context

import (
	"context"
	"net/http"
	"time"

	"github.com/tanaka-takuto/goal-minder/domain/model"
)

type authorizationHelperID struct{}

var authorizationHelperIDKey = authorizationHelperID{}

// AuthorizationHelper 認証ヘルパー
type AuthorizationHelper struct {
	writer http.ResponseWriter
}

const (
	// AuthorizationCookieName 認証トークンのCookie名
	AuthorizationCookieName = "Authorization"
)

// SetAuthorizationIntoCookie CookieにAuthorizationを設定する
func (a *AuthorizationHelper) SetAuthorizationIntoCookie(authToken model.AuthToken) {
	http.SetCookie(a.writer, &http.Cookie{
		Name:     AuthorizationCookieName,
		Value:    string(authToken),
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
	})
}

// ClearAuthorizationFromCookie CookieからAuthorizationをクリアする
func (a *AuthorizationHelper) ClearAuthorizationFromCookie() {
	http.SetCookie(a.writer, &http.Cookie{
		Name:     AuthorizationCookieName,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	})
}

// SetRequestID 認証ヘルパーをコンテキストに設定する
func SetAuthorizationHelper(ctx context.Context, writer http.ResponseWriter) context.Context {
	return context.WithValue(ctx, authorizationHelperIDKey, AuthorizationHelper{
		writer: writer,
	})
}

// GetRequestID 認証ヘルパーをコンテキストから取得する
func GetAuthorizationHelper(ctx context.Context) *AuthorizationHelper {
	authorizationSetter, ok := ctx.Value(authorizationHelperIDKey).(AuthorizationHelper)
	if !ok {
		return nil
	}

	return &authorizationSetter
}
