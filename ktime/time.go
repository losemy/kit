package ktime

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// TimeCost 计算时间消耗
// defer ktime.TimeCost("name")()
func TimeCost(methodName string) func() {
	start := time.Now()
	// 时间消耗上报
	return func() {
		tc := time.Since(start)
		timeCost := tc.Milliseconds()
		g.Log().Infof(context.Background(), "%v cost: %vms", methodName, timeCost)
	}
}
