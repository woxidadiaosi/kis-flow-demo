package log

import (
	"context"
	"testing"
)

// go test --test.v  --test.paniconexit0 --test.run TestKisLogger
// 参数解释:
// test.v表示输出详细的测试信息，包括每个测试函数的名称和运行结果。
// test.paniconexit0表示当测试代码中出现 panic 时，程序不会崩溃，而是以 0 的状态码退出。
// test.run TestKisLogger 表示只运行包含 TestKisLogger 字符串的测试函数。
func TestKisLogger(t *testing.T) {
	ctx := context.Background()

	Logger().DebugFX(ctx, "TestKisLogger DebugFX %d", 1)
	Logger().InfoFX(ctx, "TestKisLogger InfoFX %s", "test")
	Logger().ErrorFX(ctx, "TestKisLogger ErrorFX")

	Logger().DebugF("TestKisLogger DebugF")
	Logger().InfoF("TestKisLogger InfoF")
	Logger().ErrorF("TestKisLogger ErrorF")
}
