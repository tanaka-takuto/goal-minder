package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	modelContext "github.com/tanaka-takuto/goal-minder/domain/context"
)

// SetRequestID リクエストIDを設定する
func SetRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rid := uuid.New().String()
			ctx := modelContext.SetRequestID(c.Request().Context(), rid)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
