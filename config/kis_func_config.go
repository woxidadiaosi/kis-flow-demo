package config

import (
	"kis-flow-demo/common"
	"kis-flow-demo/log"
	"time"
)

// FParam 在当前Flow中Function定制固定配置参数类型
type FParam map[string]string

// kisSource 表示当前Function的业务源
type kisSource struct {
	Name string   `yaml:"name"` //本层Function的数据源描述
	Must []string `yaml:"must"` //source必传字段
}

// KisFuncOption 可选配置
type kisFuncOption struct {
	CName         string        `yaml:"cname"`          //连接器Connector名称
	RetryTime     uint          `yaml:"retry_times"`    //选填,Function调度重试(不包括正常调度)最大次数
	RetryDuration time.Duration `yaml:"retry_duration"` //选填,Function调度每次重试最大时间间隔(单位:ms)
	Params        FParam        `yaml:"default_params"` //选填,在当前Flow中Function定制固定配置参数
}

type KisFuncConfig struct {
	KisType string         `yaml:"kistype"`
	FName   string         `yaml:"fname"`
	FMode   string         `yaml:"fmode"`
	Source  *kisSource     `yaml:"source"`
	Opt     *kisFuncOption `yaml:"option"`
}

func NewKisFuncConfig(funcName string, mode common.FMode, source *kisSource, opt *kisFuncOption) *KisFuncConfig {
	if source == nil {
		log.Logger().ErrorF("funcName NewConfig Error, source is nil , funcName = %s\n", funcName)
		return nil
	}
	// 如果FMode类型为S或者L 即save或load，那么opt中必须要有参数,因为要连接外部源
	if mode == common.S || mode == common.L {
		if opt == nil {
			log.Logger().ErrorF("Function S/L need option->Cid\n")
			return nil
		} else if opt.CName == "" {
			log.Logger().ErrorF("Function S/L need option->Cid\n")
			return nil
		}
	}
	return &KisFuncConfig{
		KisType: string(common.KisFunction),
		FName:   funcName,
		FMode:   string(mode),
		Source:  source,
		Opt:     opt,
	}
}
