package config

import (
	"kis-flow-demo/common"
	"kis-flow-demo/log"
	"testing"
	"time"
)

// %+v +表示输出结构体字段+值，包括内部复杂类型的字段+值
func TestKisFuncConfig(t *testing.T) {
	source := &KisSource{
		Name: "公众号抖音商城户订单数据",
		Must: []string{"order_id", "user_id"},
	}
	opt := &KisFuncOption{
		CName:         "connectorName1",
		RetryTime:     3,
		RetryDuration: time.Microsecond * 3000,
		Params:        FParam{"param1": "value1", "param2": "value2"},
	}
	myFunc1 := NewKisFuncConfig("funcName1", common.S, source, opt)
	log.Logger().InfoF("myFunc: %+v \n", myFunc1)
}
