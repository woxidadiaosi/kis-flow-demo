package function

import (
	"context"
	"fmt"
	"kis-flow-demo/kis"
)

type KisFunctionV struct {
	BaseFunction
}

func (v *KisFunctionV) Call(ctx context.Context, flow kis.Flow) error {
	fmt.Printf("kisFunctionV, flow = %+v", flow)
	//TODO 调用具体Function执行方法
	return nil
}
