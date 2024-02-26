package config

import (
	"kis-flow-demo/common"
	"kis-flow-demo/log"
	"testing"
	"time"
)

func TestKisConnConfig(t *testing.T) {
	source := &kisSource{
		Name: "公众号抖音商城户订单数据",
		Must: []string{"order_id", "user_id"},
	}
	opt := &kisFuncOption{
		CName:         "connectorName1",
		RetryTime:     3,
		RetryDuration: time.Microsecond * 3000,
		Params:        FParam{"param1": "value1", "param2": "value2"},
	}
	myFunc1 := NewKisFuncConfig("funcName1", common.S, source, opt)

	kisConnConfig := NewKisConnConfig("conn-test", "topic1", "127.0.0.1:9092", common.Kafka, nil)

	kisConnConfig.WithFunc(myFunc1)
	log.Logger().InfoF("kisConfigConn-test, %+v", kisConnConfig)

}
