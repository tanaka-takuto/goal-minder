package context

import (
	"context"
)

type requestID struct{}

var requestIDKey = requestID{}

// SetRequestID リクエストIDをコンテキストに設定する
func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// GetRequestID リクエストIDをコンテキストから取得する
func GetRequestID(ctx context.Context) *string {
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return nil
	}

	return &requestID
}
