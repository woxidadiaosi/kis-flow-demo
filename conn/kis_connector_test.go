package conn_test

import (
	"context"
	"fmt"
	"kis-flow-demo/common"
	"kis-flow-demo/config"
	"kis-flow-demo/flow"
	"kis-flow-demo/kis"
	"kis-flow-demo/log"
	"testing"
	"time"
)

func handle1(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoFX(ctx, "call function handle1 ... \n")
	for i, row := range flow.Input() {
		fmt.Printf("====>data[%+v], index[%d] \n", row, i)
		s, ok := row.(string)
		if !ok {
			continue
		}
		_ = flow.CommitRow("handle1=>" + s)
	}
	return nil
}

func handle2(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoFX(ctx, "call function handle2... \n")
	//function := flow.GetThisFunction()
	c, err := flow.GetConnector()
	if err != nil {
		log.Logger().ErrorFX(ctx, "function handle2 GetConnector error = %s \n", err)
		return err
	}
	if err = c.Init(); err != nil {
		log.Logger().ErrorFX(ctx, "function handle2 connector Init() error = %s \n", err)
		return err
	}

	for _, row := range flow.Input() {
		if err := c.Call(ctx, flow, row); err != nil {
			log.Logger().ErrorFX(ctx, "function handle2 connector Call() error = %s \n", err)
		}
		s, ok := row.(string)
		if !ok {
			continue
		}
		_ = flow.CommitRow("handle2=>" + s)
	}
	return nil
}

func cInit(conn kis.Connector) error {
	log.Logger().InfoF("connector init()... \n")
	log.Logger().InfoF("do something ... \n")
	return nil
}

func caasHandle(ctx context.Context, c kis.Connector, f kis.Function, flow kis.Flow, row interface{}) error {
	log.Logger().InfoFX(ctx, "connector caasHandle ... \n")
	fmt.Printf("conn config : %v \n", c.GetConfig())
	log.Logger().InfoFX(ctx, "do something...\n")
	return nil
}

func TestConnector(t *testing.T) {
	ctx := context.Background()
	kis.Pool().Faas("handle1", handle1)
	kis.Pool().Faas("handle2", handle2)

	kis.Pool().CaasInit("conn1", cInit)
	kis.Pool().Caas("conn1", "handle2", common.L, caasHandle)

	source1 := &config.KisSource{
		Name: "source1",
		Must: []string{"order_id", "user_id"},
	}
	source2 := &config.KisSource{
		Name: "source2",
		Must: []string{"order_id", "user_id"},
	}

	param := map[string]string{
		"param1": "value1",
	}
	connConfig := config.NewKisConnConfig("conn1", "127.0.0.1:9092", "topic1", common.Kafka, param)

	funcConfig1 := config.NewKisFuncConfig("handle1", common.V, source1, nil)
	opt := &config.KisFuncOption{CName: "conn1", RetryTime: 3, RetryDuration: 3 * time.Microsecond, Params: param}
	funcConfig2 := config.NewKisFuncConfig("handle2", common.L, source2, opt)
	funcConfig2.AddConnConf(connConfig)

	flowConf := config.NewKisFlowConfig("flow1", common.ON)
	flow := flow.NewKisFlow(flowConf)
	flow.Link(funcConfig1, nil)
	flow.Link(funcConfig2, nil)

	// 7. 提交原始数据
	_ = flow.CommitRow("This is Data1 from Test")
	_ = flow.CommitRow("This is Data2 from Test")
	_ = flow.CommitRow("This is Data3 from Test")

	// 8. 执行flow1
	if err := flow.Run(ctx); err != nil {
		panic(err)
	}
}
