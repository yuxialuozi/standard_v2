package hello3

import (
	"context"

	hello3 "cwgo_test/hertz_gen/hello3"
	"github.com/cloudwego/hertz/pkg/app"
)

type HelloMethodService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethodService(Context context.Context, RequestContext *app.RequestContext) *HelloMethodService {
	return &HelloMethodService{RequestContext: RequestContext, Context: Context}
}

func (h *HelloMethodService) Run(req *hello3.HelloReq) (resp *hello3.HelloResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
