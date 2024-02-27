package flow

import (
	"context"
	"kis-flow-demo/common"
	"kis-flow-demo/config"
	"kis-flow-demo/log"
	"testing"
	"time"
)

func TestKisFlow_Link(t *testing.T) {
	source := &config.KisSource{
		Name: "table",
		Must: []string{"user_id", "order_id"},
	}
	source2 := &config.KisSource{
		Name: "table2",
		Must: []string{"user_id", "order_id"},
	}
	opt := &config.KisFuncOption{
		CName:         "mysql",
		RetryTime:     3,
		RetryDuration: time.Microsecond * 3000,
		Params:        config.FParam{"key": "value"},
	}
	kisFuncConf := config.NewKisFuncConfig("kisFuncTest", common.L, source, opt)
	kisFuncConf2 := config.NewKisFuncConfig("kisFuncTest", common.S, source2, opt)

	kisFlowConf := config.NewKisFlowConfig("flow", common.ON)
	kisFlow := NewKisFlow(kisFlowConf)

	p := config.FParam{
		"key": "value2",
	}

	kisFlow.Link(kisFuncConf, p)
	kisFlow.Link(kisFuncConf2, p)

	log.Logger().InfoF("kisFlow link result = %+v", kisFlow)

	kisFlow.Run(context.Background())
}

func TestKisFlow_Run(t *testing.T) {
	source1 := &config.KisSource{
		Name: "公众号抖音商城户订单数据",
		Must: []string{"order_id", "user_id"},
	}

	source2 := &config.KisSource{
		Name: "用户订单错误率",
		Must: []string{"order_id", "user_id"},
	}

	funcConfig1 := config.NewKisFuncConfig("myFunc1", common.C, source1, nil)
	funcConfig2 := config.NewKisFuncConfig("myFunc2", common.E, source2, nil)

	flowConfig := config.NewKisFlowConfig("myFlow", common.ON)
	flow := NewKisFlow(flowConfig)

	if err := flow.Link(funcConfig1, nil); err != nil {
		t.Error(err.Error())
	}
	if err := flow.Link(funcConfig2, nil); err != nil {
		t.Error(err)
	}

	flow.CommitRow("this is test data1")
	flow.CommitRow("this is test data2")
	flow.CommitRow("this is test data3")
	ctx := context.Background()
	if err := flow.Run(ctx); err != nil {
		t.Error(err)
	}
}
