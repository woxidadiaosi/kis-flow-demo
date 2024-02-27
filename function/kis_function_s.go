package function

import (
	"context"
	"fmt"
	"kis-flow-demo/kis"
)

type KisFunctionS struct {
	BaseFunction
}

func (s *KisFunctionS) Call(ctx context.Context, flow kis.Flow) error {
	fmt.Printf("kisFunctionS, flow = %+v", flow)
	//TODO 调用具体Function执行方法
	return nil
}
