package ktime

import (
	"time"

	"kit/log"
)

// TimeCost 计算时间消耗
// defer ktime.TimeCost("name")()
func TimeCost(methodName string) func() {
	start := time.Now()
	// 时间消耗上报
	return func() {
		tc := time.Since(start)
		timeCost := tc.Milliseconds()
		log.Infof("%v cost: %vms", methodName, timeCost)
	}
}
