// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"likegozero/api/demo/internal/svc"
	"likegozero/api/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoHandler2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoHandler2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoHandler2Logic {
	return &DemoHandler2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoHandler2Logic) DemoHandler2(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = req.Name
	return
}
