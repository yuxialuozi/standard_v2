package hello2

import (
	"context"

	service "cwgo_test/biz/service/hello2"
	"cwgo_test/biz/utils"
	hello2 "cwgo_test/hertz_gen/hello2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// HelloMethod .
// @router /hello2 [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hello2.HelloReq
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
