package function

import (
	"context"
	"fmt"
	"kis-flow-demo/kis"
	"kis-flow-demo/log"
)

type KisFunctionC struct {
	BaseFunction
}

func (c *KisFunctionC) Call(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoF("KisFunction, flow = %+v \n", flow)

	for i, row := range flow.Input() {
		fmt.Printf("kisFunctionC, row = %+v \n", row)
		_ = flow.CommitRow(fmt.Sprintf("Data from kisFunctionC , index = %d", i))
	}
	return nil
}
