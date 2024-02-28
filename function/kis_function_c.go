package function

import (
	"context"
	"kis-flow-demo/kis"
	"kis-flow-demo/log"
)

type KisFunctionC struct {
	BaseFunction
}

func (c *KisFunctionC) Call(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoF("KisFunction, flow = %+v \n", flow)

	if err := kis.Pool().CallFunction(ctx, c.Config.FName, flow); err != nil {
		log.Logger().ErrorFX(ctx, "Function called error = %s \n", err)
		return err
	}
	return nil
}
