package function

import (
	"context"
	"errors"
	"kis-flow-demo/common"
	"kis-flow-demo/config"
	"kis-flow-demo/id"
	"kis-flow-demo/kis"
)

type BaseFunction struct {
	// Id , KisFunction的实例ID，用于KisFlow内部区分不同的实例对象
	Id     string
	Config *config.KisFuncConfig

	Flow kis.Flow

	N kis.Function
	P kis.Function
}

func (b *BaseFunction) Call(ctx context.Context, flow kis.Flow) error {
	return nil
}

func (b *BaseFunction) SetConfig(s *config.KisFuncConfig) error {
	if s == nil {
		return errors.New("kisFuncConfig is nil")
	}
	b.Config = s
	return nil
}

func (b *BaseFunction) GetConfig() *config.KisFuncConfig {
	return b.Config
}

func (b *BaseFunction) SetFlow(f kis.Flow) error {
	if f == nil {
		return errors.New("kisFlow is nil")
	}
	b.Flow = f
	return nil
}

func (b *BaseFunction) GetFlow() kis.Flow {
	return b.Flow
}

func (b *BaseFunction) CreateId() {
	b.Id = id.KisID(common.KisIdTypeFunction)
}

func (b *BaseFunction) GetId() string {
	return b.Id
}

func (b *BaseFunction) GetPrevId() string {
	if b.P == nil {
		//表示该function为首节点
		return common.FunctionIdFirstVirtual
	}
	return b.P.GetId()
}

func (b *BaseFunction) GetNextId() string {
	if b.N == nil {
		return common.FunctionIdLastVirtual
	}
	return b.N.GetId()
}

func (b *BaseFunction) Next() kis.Function {
	return b.N
}

func (b *BaseFunction) Prev() kis.Function {
	return b.P
}

func (b *BaseFunction) SetN(f kis.Function) {
	b.N = f
}

func (b *BaseFunction) SetP(f kis.Function) {
	b.P = f
}

func NewKisFunction(conf *config.KisFuncConfig) (kis.Function, error) {
	var f kis.Function
	switch common.FMode(conf.FMode) {
	case common.V:
		f = &KisFunctionV{}
	case common.S:
		f = &KisFunctionS{}
	case common.L:
		f = &KisFunctionL{}
	case common.C:
		f = &KisFunctionL{}
	case common.E:
		f = &KisFunctionE{}
	default:
		return nil, nil
	}

	f.CreateId()

	if err := f.SetConfig(conf); err != nil {
		return nil, err
	}
	return f, nil
}
