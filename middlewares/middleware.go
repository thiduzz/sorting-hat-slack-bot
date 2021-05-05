package middlewares

import (
	"context"
	"encoding/json"
)

type HandlerFunc func(context.Context, json.RawMessage) (interface{}, error)

func MiddlewareFunc(next HandlerFunc) HandlerFunc {
	return HandlerFunc(func(ctx context.Context, data json.RawMessage) (interface{}, error) {
		return next(ctx, data)
	})
}