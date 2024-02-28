package function

import (
	"context"
	"kis-flow-demo/kis"
	"kis-flow-demo/log"
)

type KisFunctionL struct {
	BaseFunction
}

func (l *KisFunctionL) Call(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoFX(ctx, "KisFunction, flow = %+v", flow)

	if err := kis.Pool().CallFunction(ctx, l.Config.FName, flow); err != nil {
		log.Logger().ErrorFX(ctx, "Function called error = %s", err)
		return err
	}
	return nil
}
