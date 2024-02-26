package config

import (
	"kis-flow-demo/common"
	"kis-flow-demo/log"
	"testing"
)

func TestKisFLowConfig(t *testing.T) {
	kfp1 := KisFuncParam{FName: "func1", param: FParam{"param1": "value1"}}
	kfp2 := KisFuncParam{FName: "func2", param: FParam{"param1": "value1"}}
	kisFlow := NewKisFlowConfig("test-kis-flow", common.ON)
	kisFlow.AppendFunctionConfig(kfp1)
	kisFlow.AppendFunctionConfig(kfp2)

	log.Logger().InfoF("kisFlow : %+v \n", kisFlow)
}
