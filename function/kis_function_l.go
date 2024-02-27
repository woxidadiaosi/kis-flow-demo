package function

import (
	"context"
	"fmt"
	"kis-flow-demo/kis"
)

type KisFunctionL struct {
	BaseFunction
}

func (l *KisFunctionL) Call(ctx context.Context, flow kis.Flow) error {
	fmt.Printf("kisFunctionL, flow = %+v", flow)
	//TODO 调用具体Function执行方法
	return nil
}
