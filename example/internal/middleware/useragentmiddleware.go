// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"context"
	"fmt"
	"net/http"
)

type UserAgentMiddleware struct {
}

func NewUserAgentMiddleware() *UserAgentMiddleware {
	return &UserAgentMiddleware{}
}

func (m *UserAgentMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		fmt.Println("----------- UserAgentMiddleware add User-Agent  pre-----------")
		//将请求头中的User-Agent保存在context中
		val := r.Header.Get("User-Agent")
		reqCtx := r.Context()
		ctx := context.WithValue(reqCtx, "User-Agent", val)
		newReq := r.WithContext(ctx)
		fmt.Println("----------- UserAgentMiddleware add User-Agent  after-----------")
		next(w, newReq)
	}
}
