// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"likegozero/example/internal/svc"
	"likegozero/example/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoHandlerV1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoHandlerV1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoHandlerV1Logic {
	return &DemoHandlerV1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoHandlerV1Logic) DemoHandlerV1(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = req.Name
	return
}
