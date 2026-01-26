package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

func GlobalMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("global middleware---pre")
		next(w, r)
		logx.Info("global middleware---after")
	}
}
