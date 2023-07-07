package kmiddleware

import (
	xerrors "github.com/losemy/kit/errors"
	"github.com/losemy/kit/kconstant"
	"github.com/losemy/kit/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.ClearBuffer()
		if data, ok := err.(*xerrors.Error); ok {
			r.Response.WriteJson(&model.Response{
				Code: data.Code,
				Msg:  data.Msg,
			})
		} else {
			r.Response.WriteJson(&model.Response{
				Code: kconstant.ErrSystem.Code,
				Msg:  kconstant.ErrSystem.Msg,
			})
		}
	}
	r.SetError(nil)
}
