package hook

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type shutdown struct {
	hooks []func()
	once  sync.Once
}

var Shutdown = new(shutdown)

func (c *shutdown) Register(fn func()) {
	c.hooks = append(c.hooks, fn)
}

// Signal 使用方式 主线程 调用 Signal 阻塞 然后接受对应的信号完成 close
func (c *shutdown) Signal() {
	signalCh := make(chan os.Signal)

	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)

	select {
	case <-signalCh:
		// 完成对应的 任务处理
		// 其实就是这里处理退出逻辑就行了
	}
	c.Close()
}

func (c *shutdown) Close() {
	// 执行 close操作
	c.once.Do(func() {
		for _, fn := range c.hooks {
			fn()
		}
	})

}
