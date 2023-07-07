package kmiddleware

import (
	"bytes"
	"io"
	"time"

	"github.com/losemy/kit/kconstant"
	"github.com/losemy/kit/log"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func Sign(r *ghttp.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		r.SetError(kconstant.ErrSystem)
		return
	}
	r.Request.Body.Close() // NOTE 原始的 Body 无需手动关闭,会在 response.reqBody中自动关闭的.
	br := bytes.NewReader(data)
	r.Request.Body = io.NopCloser(br)
	sign := r.Header.Get(kconstant.Sign)
	// 加上时间戳
	timestamp := r.Header.Get(kconstant.Timestamp)
	timeNow := time.Now().Unix()
	mistakeTime := g.Cfg().MustGet(r.Context(), "mistake_time").Int64()
	if timeNow-gconv.Int64(timestamp) > mistakeTime ||
		(timeNow-gconv.Int64(timestamp) < 0 && gconv.Int64(timestamp)-timeNow > mistakeTime) {
		log.Debugf("time: %d --> %d --> %s", timeNow, mistakeTime, timestamp)
		r.SetError(kconstant.ErrTimestamp)
		return
	}
	dataSign := append(data, []byte(timestamp)...)
	signC, err := gmd5.EncryptBytes(dataSign)
	if err != nil {
		r.SetError(kconstant.ErrSystem)
		return
	}
	if sign != signC {
		log.Debugf("sign need: %s", signC)
		r.SetError(kconstant.ErrSign)
		return
	}
	r.Middleware.Next()
}
