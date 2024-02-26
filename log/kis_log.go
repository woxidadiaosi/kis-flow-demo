package log

import "context"

type KisLogger interface {
	// InfoFX 有上下文的Info级别日志接口， format字符串格式
	InfoFX(ctx context.Context, str string, v ...interface{})
	// ErrorFX 有上下文的Info级别日志接口， format字符串格式
	ErrorFX(ctx context.Context, str string, v ...interface{})
	// DebugFX 有上下文的Info级别日志接口， format字符串格式
	DebugFX(ctx context.Context, str string, v ...interface{})

	// InfoF 无上下文的Info级别日志接口，format字符串格式
	InfoF(str string, v ...interface{})
	// ErrorF 无上下文的Error级别日志接口，format字符串格式
	ErrorF(str string, v ...interface{})
	// DebugF 无上下文的Debug级别日志接口，format字符串格式
	DebugF(str string, v ...interface{})
}

// KisLog 默认的KisLogger对象
var KisLog KisLogger

// SetLogger 设置KisLog对象，可以是用自定义的Logger对象
func SetLogger(newLog KisLogger) {
	KisLog = newLog
}

// Logger 获取到kisLog对象
func Logger() KisLogger {
	return KisLog
}