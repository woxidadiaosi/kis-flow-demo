package log

import (
	"context"
	"fmt"
)

type kisDefaultLog struct{}

func (l *kisDefaultLog) InfoFX(ctx context.Context, str string, v ...interface{}) {
	fmt.Println(ctx)
	fmt.Printf(str, v...)
}

func (l *kisDefaultLog) ErrorFX(ctx context.Context, str string, v ...interface{}) {
	fmt.Println(ctx)
	fmt.Printf(str, v...)
}

func (l *kisDefaultLog) DebugFX(ctx context.Context, str string, v ...interface{}) {
	fmt.Println(ctx)
	fmt.Printf(str, v...)
}

func (l *kisDefaultLog) InfoF(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}

func (l *kisDefaultLog) ErrorF(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}

func (l *kisDefaultLog) DebugF(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}

func init() {
	// 如果没有设置Logger, 则启动时使用默认的kisDefaultLog对象
	if Logger() == nil {
		SetLogger(&kisDefaultLog{})
	}
}
