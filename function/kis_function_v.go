package function

import (
	"context"
	"kis-flow-demo/kis"
	"kis-flow-demo/log"
)

type KisFunctionV struct {
	BaseFunction
}

func (v *KisFunctionV) Call(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoFX(ctx, "KisFunctionV, flow = %+v", flow)

	if err := kis.Pool().CallFunction(ctx, v.Config.FName, flow); err != nil {
		log.Logger().ErrorFX(ctx, "Function called error = %s", err)
		return err
	}
	return nil
}
