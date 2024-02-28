package function

import (
	"context"
	"kis-flow-demo/kis"
	"kis-flow-demo/log"
)

type KisFunctionE struct {
	BaseFunction
}

func (e *KisFunctionE) Call(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoF("kisFunctionE, flow = %+v \n", flow)

	if err := kis.Pool().CallFunction(ctx, e.Config.FName, flow); err != nil {
		log.Logger().ErrorFX(ctx, "Function called error = %s", err)
		return err
	}

	return nil
}
