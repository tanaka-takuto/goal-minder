package middleware

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/tanaka-takuto/goal-minder/domain/context"
	"github.com/tanaka-takuto/goal-minder/domain/model"
)

// Authentication 認証を行う
func Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Cookieへの認証ヘルパーをコンテキストに設定
			ctx := c.Request().Context()
			ctx = context.SetAuthorizationHelper(ctx, c.Response().Writer)

			// Cookieから認証トークンを取得
			authCookie, _ := c.Request().Cookie(context.AuthorizationCookieName)
			if authCookie != nil {
				oldAuthToken := model.AuthToken(strings.ReplaceAll(authCookie.Value, "Bearer ", ""))
				oldAuthTokenClaims, _ := oldAuthToken.Decode()
				if oldAuthTokenClaims != nil {
					// コンテキストにアカウントIDを設定
					ctx = context.SetAccountID(ctx, oldAuthTokenClaims.AccountID)

					// Cookieの認証トークンを延長するため新規作成
					newAuthTokenClaims := model.NewAuthTokenClaims(oldAuthTokenClaims.AccountID)
					newAuthToken := newAuthTokenClaims.Encode()

					// 認証トークンをクリアする場合は別途処理が必要
					setter := context.GetAuthorizationHelper(ctx)
					setter.SetAuthorizationIntoCookie(newAuthToken)
				}
			}

			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
