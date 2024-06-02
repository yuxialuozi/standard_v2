package hello3

import (
	"context"

	service "cwgo_test/biz/service/hello3"
	"cwgo_test/biz/utils"
	hello3 "cwgo_test/hertz_gen/hello3"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// HelloMethod .
// @router /hello3 [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hello3.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusBadRequest, err)
		return
	}

	resp, err := service.NewHelloMethodService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
