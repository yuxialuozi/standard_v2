// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	hello1 "cwgo_test/biz/router/hello1"
	hello2 "cwgo_test/biz/router/hello2"
	hello3 "cwgo_test/biz/router/hello3"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	hello3.Register(r)

	hello2.Register(r)

	hello1.Register(r)
}
