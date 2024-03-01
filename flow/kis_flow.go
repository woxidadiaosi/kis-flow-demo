package flow

import (
	"context"
	"errors"
	"fmt"
	"kis-flow-demo/common"
	"kis-flow-demo/config"
	"kis-flow-demo/conn"
	"kis-flow-demo/function"
	"kis-flow-demo/id"
	"kis-flow-demo/kis"
	"kis-flow-demo/log"
	"sync"
)

type KisFlow struct {
	Id   string
	Name string
	Conf *config.KisFlowConfig

	Funcs          map[string]kis.Function
	FlowHead       kis.Function
	FlowTail       kis.Function
	flock          sync.Mutex
	ThisFunction   kis.Function
	ThisFunctionId string
	PrevFunctionId string

	funcParam map[string]config.FParam
	fplock    sync.Mutex

	buffer common.KisRowArr
	data   common.KisDataMap
	input  common.KisRowArr
}

func (k *KisFlow) GetConnector() (kis.Connector, error) {
	c := k.ThisFunction.GetConnector()
	if c == nil {
		return nil, errors.New("GetConnector() ThisFunction is nil")
	}
	return c, nil
}

func (k *KisFlow) GetConnectorConf() (*config.KisConnConfig, error) {
	if conn := k.ThisFunction.GetConnector(); conn != nil {
		return conn.GetConfig(), nil
	}
	return nil, errors.New("GetConnectorConf() ThisFunction is nil")
}

func (k *KisFlow) Run(ctx context.Context) error {

	if k.Conf.Status == uint(common.OFF) {
		return nil
	}

	var f kis.Function
	f = k.FlowHead
	if err := k.CommitSrcData(ctx); err != nil {
		return err
	}
	k.PrevFunctionId = common.FunctionIdFirstVirtual
	// 流式链式调用
	for f != nil {
		k.ThisFunction = f
		k.ThisFunctionId = f.GetId()

		if input, err := k.getCurData(); err != nil {
			log.Logger().ErrorFX(ctx, "flow.Run(), getCurData err = %s \n", err.Error())
			return err
		} else {
			k.input = input
		}

		if err := f.Call(ctx, k); err != nil {
			return err
		} else {
			if err := k.CommitCurData(ctx); err != nil {
				return err
			}
			k.PrevFunctionId = k.ThisFunctionId
			f = f.Next()
		}
	}
	return nil
}

func (k *KisFlow) Link(fConf *config.KisFuncConfig, fParam config.FParam) error {
	kisFunction, err := function.NewKisFunction(fConf)
	if err != nil {
		return err
	}
	// create connector
	if fConf.Opt != nil && fConf.Opt.CName != "" {
		conf, err := fConf.GetConnConf()
		if err != nil {
			panic(err)
		}
		connector := conn.NewKisConnector(conf)
		if err = connector.Init(); err != nil {
			panic(err)
		}

		kisFunction.AddConnector(connector)

	}
	// append Func
	if err := k.appFunc(kisFunction, fParam); err != nil {
		return err
	}
	return nil
}

func (k *KisFlow) appFunc(kisFunction kis.Function, fParam config.FParam) error {
	k.flock.Lock()
	defer k.flock.Unlock()
	k.Funcs[kisFunction.GetConfig().FName] = kisFunction

	kisFunction.SetFlow(k)

	if k.FlowHead == nil {
		k.FlowHead = kisFunction
		k.FlowTail = kisFunction

		kisFunction.SetP(nil)
		kisFunction.SetN(nil)

	} else {
		kisFunction.SetP(k.FlowTail)
		kisFunction.SetN(nil)

		k.FlowTail.SetN(kisFunction)
		k.FlowTail = kisFunction

	}

	//k.funcParam[fConf.FName] = fParam 不符合规范，对于fParam为引用，外部可能修改，最好单独拷贝一份
	ps := make(config.FParam) //这里参数也不一定全部不同，可能新传递的覆盖了默认的
	//先添加function 默认携带的Params参数
	if kisFunction.GetConfig().Opt != nil && kisFunction.GetConfig().Opt.Params != nil {
		for k, v := range kisFunction.GetConfig().Opt.Params {
			ps[k] = v
		}
	}
	//再添加flow携带的function定义参数(重复即覆盖)
	for k, v := range fParam {
		ps[k] = v
	}

	// 将得到的FParams存留在flow结构体中，用来function业务直接通过Hash获取
	// key 为当前Function的KisId，不用Fid的原因是为了防止一个Flow添加两个相同策略Id的Function
	k.funcParam[kisFunction.GetId()] = ps
	return nil
}

func (k *KisFlow) CommitRow(row interface{}) error {
	k.buffer = append(k.buffer, row)
	return nil
}

func (k *KisFlow) CommitSrcData(ctx context.Context) error {
	dCnt := len(k.buffer)
	batch := make(common.KisRowArr, 0, dCnt)
	for _, da := range k.buffer {
		batch = append(batch, da)
	}
	k.clearData(k.data)

	k.data[common.FunctionIdFirstVirtual] = batch

	k.buffer = k.buffer[0:0]

	log.Logger().DebugFX(ctx, "===> After CommitSrcData, flow_name = %s, flow_id = %s \n All Level Data = \n%+v\n", k.Name, k.Id, k.data)
	return nil
}

func (k *KisFlow) clearData(dataMap common.KisDataMap) {
	for k := range dataMap {
		delete(dataMap, k)
	}
}

func (k *KisFlow) CommitCurData(ctx context.Context) error {
	if len(k.buffer) == 0 {
		return nil
	}

	batch := make(common.KisRowArr, 0, len(k.buffer))

	for _, da := range k.buffer {
		batch = append(batch, da)
	}

	k.data[k.ThisFunctionId] = batch

	k.buffer = k.buffer[0:0]

	log.Logger().DebugFX(ctx, "===> After CommitCurData,flow_name = %s, flow_id = %s \n All Level Data = \n %+v \n", k.Name, k.Id, k.data)

	return nil
}

func (k *KisFlow) getCurData() (common.KisRowArr, error) {
	if k.PrevFunctionId == "" {
		return nil, errors.New("flow.PrevFunctionId is nil, maybe not set")
	}

	if _, ok := k.data[k.PrevFunctionId]; !ok {
		return nil, errors.New(fmt.Sprintf("[%s] is not in flow.data", k.PrevFunctionId))
	}

	return k.data[k.PrevFunctionId], nil
}

func (k *KisFlow) Input() common.KisRowArr {
	return k.input
}

func (k *KisFlow) GetName() string {
	return k.Name
}

func (k *KisFlow) GetThisFunction() kis.Function {
	return k.ThisFunction
}

func (k *KisFlow) GetThisFunctionConf() *config.KisFuncConfig {
	return k.ThisFunction.GetConfig()
}

func NewKisFlow(conf *config.KisFlowConfig) kis.Flow {
	return &KisFlow{
		Id:        id.KisID(common.KisIdTypeFlow),
		Name:      conf.FlowName,
		Conf:      conf,
		Funcs:     make(map[string]kis.Function),
		funcParam: make(map[string]config.FParam),
		data:      make(map[string]common.KisRowArr),
	}
}
