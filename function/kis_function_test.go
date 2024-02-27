package function

import (
	"context"
	"kis-flow-demo/common"
	"kis-flow-demo/config"
	"kis-flow-demo/flow"
	"testing"
	"time"
)

func TestKisFunction(t *testing.T) {
	source := &config.KisSource{
		Name: "table",
		Must: []string{"user_id", "order_id"},
	}
	opt := &config.KisFuncOption{
		CName:         "mysql",
		RetryTime:     3,
		RetryDuration: time.Microsecond * 3000,
		Params:        config.FParam{"key": "value"},
	}
	kisFuncConf := config.NewKisFuncConfig("kisFuncTest", common.L, source, opt)

	kisFlowConf := config.NewKisFlowConfig("flow", common.ON)
	kisFlow := flow.NewKisFlow(kisFlowConf)

	function, err := NewKisFunction(kisFuncConf)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	if err = function.Call(ctx, kisFlow); err != nil {
		t.Error(err)
	}

}
