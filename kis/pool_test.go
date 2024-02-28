package kis_test

import (
	"context"
	"fmt"
	"kis-flow-demo/common"
	"kis-flow-demo/config"
	"kis-flow-demo/flow"
	"kis-flow-demo/kis"
	"kis-flow-demo/log"
	"testing"
)

func handle1(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoFX(ctx, "自定义函数handle1 in...")

	for index, da := range flow.Input() {
		fmt.Printf("function= [%s], functionId = [%s], 进入的数据 =》 %v \n", flow.GetThisFunction().GetConfig().FName, flow.GetThisFunction().GetId(), da)

		//计算结果
		resultStr := fmt.Sprintf("data from functionName[%s], index = [%d]", flow.GetThisFunctionConf().FName, index)

		_ = flow.CommitRow(resultStr)
	}
	return nil
}

func handle2(ctx context.Context, flow kis.Flow) error {
	log.Logger().InfoFX(ctx, "自定义函数handle2 in ...")

	for _, da := range flow.Input() {
		fmt.Printf("funcion = [%s], functionId = [%s], 进入的数据 =》 %v \n", flow.GetThisFunctionConf().FName, flow.GetThisFunction().GetId(), da)
	}

	return nil
}

func TestPool(t *testing.T) {
	ctx := context.Background()

	kis.Pool().Faas("handle1", handle1)
	kis.Pool().Faas("handle2", handle2)

	source1 := &config.KisSource{
		Name: "source1",
		Must: []string{"order_id", "user_id"},
	}
	source2 := &config.KisSource{
		Name: "source2",
		Must: []string{"order_id", "user_id"},
	}
	fCon1 := config.NewKisFuncConfig("handle1", common.V, source1, nil)
	fCon2 := config.NewKisFuncConfig("handle2", common.C, source2, nil)

	flowConf := config.NewKisFlowConfig("flow1", common.ON)
	flow := flow.NewKisFlow(flowConf)

	flow.Link(fCon1, nil)
	flow.Link(fCon2, nil)

	flow.CommitRow("data1")
	flow.CommitRow("data2")
	flow.CommitRow("data3")
	flow.CommitRow("data4")

	if err := flow.Run(ctx); err != nil {
		t.Errorf("flow Run() error = %s", err)
	}

}
