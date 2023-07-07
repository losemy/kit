package kmiddleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	xerrors "github.com/losemy/kit/errors"
	"github.com/losemy/kit/kconstant"
	"github.com/losemy/kit/kmodel"
)

func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.ClearBuffer()
		if data, ok := err.(*xerrors.Error); ok {
			r.Response.WriteJson(kmodel.Error(data))
		} else {
			r.Response.WriteJson(kmodel.Error(kconstant.ErrSystem))
		}
	}
	r.SetError(nil)
}
