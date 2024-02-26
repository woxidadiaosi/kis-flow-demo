package config

import (
	"errors"
	"fmt"
	"kis-flow-demo/common"
)

type KisConnConfig struct {
	KisType    string             `yaml:"kistype"`
	CName      string             `yaml:"cname"'`
	AddrString string             `yaml:"addrs"`
	Type       common.KisConnType `yaml:"type"`
	Key        string             `yaml:"key"`
	Param      map[string]string  `yaml:"params"`
	Load       []string           `yaml:"load"`
	Save       []string           `yaml:"save"`
}

func NewKisConnConfig(cName string, adds string, key string, t common.KisConnType, param map[string]string) *KisConnConfig {
	return &KisConnConfig{
		KisType:    string(common.C),
		CName:      cName,
		AddrString: adds,
		Type:       t,
		Key:        key,
		Param:      param,
	}
}

func (c *KisConnConfig) WithFunc(kisFuncConfig *KisFuncConfig) error {
	switch common.FMode(kisFuncConfig.FMode) {
	case common.L:
		c.Load = append(c.Load, kisFuncConfig.FName)
	case common.S:
		c.Save = append(c.Save, kisFuncConfig.FName)
	default:
		return errors.New(fmt.Sprintf("Wrong KisMode %s \n", kisFuncConfig.FMode))
	}
	return nil
}
