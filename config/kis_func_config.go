package config

import (
	"errors"
	"kis-flow-demo/common"
	"kis-flow-demo/log"
	"time"
)

// FParam 在当前Flow中Function定制固定配置参数类型
type FParam map[string]string

// KisSource 表示当前Function的业务源
type KisSource struct {
	Name string   `yaml:"name"` //本层Function的数据源描述
	Must []string `yaml:"must"` //source必传字段
}

// KisFuncOption 可选配置
type KisFuncOption struct {
	CName         string        `yaml:"cname"`          //连接器Connector名称
	RetryTime     uint          `yaml:"retry_times"`    //选填,Function调度重试(不包括正常调度)最大次数
	RetryDuration time.Duration `yaml:"retry_duration"` //选填,Function调度每次重试最大时间间隔(单位:ms)
	Params        FParam        `yaml:"default_params"` //选填,在当前Flow中Function定制固定配置参数
}

type KisFuncConfig struct {
	KisType  string         `yaml:"kistype"`
	FName    string         `yaml:"fname"`
	FMode    string         `yaml:"fmode"`
	Source   *KisSource     `yaml:"source"`
	Opt      *KisFuncOption `yaml:"option"`
	connConf *KisConnConfig `yaml:"connConf"`
}

func (k *KisFuncConfig) AddConnConf(connConf *KisConnConfig) error {
	if connConf == nil {
		return errors.New("connConf is nil")
	}
	k.connConf = connConf
	return connConf.WithFunc(k)
}

func (k *KisFuncConfig) GetConnConf() (*KisConnConfig, error) {
	if k.connConf == nil {
		return nil, errors.New("connConfig is nil")
	}
	return k.connConf, nil
}

func NewKisFuncConfig(funcName string, mode common.FMode, source *KisSource, opt *KisFuncOption) *KisFuncConfig {
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
