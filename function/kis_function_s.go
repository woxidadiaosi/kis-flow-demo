package function

import (
	"context"
	"kis-flow-demo/kis"
	"kis-flow-demo/log"
)

type KisFunctionS struct {
	BaseFunction
}

func (s *KisFunctionS) Call(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoFX(ctx, "KisFunctionS, flow = %+v", flow)

	if err := kis.Pool().CallFunction(ctx, s.Config.FName, flow); err != nil {
		log.Logger().ErrorFX(ctx, "Function called error = %s", err)
		return err
	}
	return nil
}
