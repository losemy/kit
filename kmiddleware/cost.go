package kmiddleware

import (
	"kit/ktime"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Cost(r *ghttp.Request) {
	path := r.URL.Path
	defer ktime.TimeCost(path)()
	r.Middleware.Next()
}
