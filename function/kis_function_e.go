package function

import (
	"context"
	"fmt"
	"kis-flow-demo/kis"
	"kis-flow-demo/log"
)

type KisFunctionE struct {
	BaseFunction
}

func (e *KisFunctionE) Call(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoF("kisFunctionE, flow = %+v \n", flow)

	for _, row := range flow.Input() {
		fmt.Printf("in kisFunctionE, row = %+v", row)
	}

	return nil
}
