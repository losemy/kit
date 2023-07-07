package kconfig

import (
	"context"

	"kit/log"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetConfig 获取配置
func GetConfig[T any](name ...string) T {
	var d T
	data := g.Cfg(name...).MustData(context.Background())
	if err := gconv.Struct(data, &d); err != nil {
		log.Errorf("parse config err: %v", err)
		return d
	}
	return d
}
