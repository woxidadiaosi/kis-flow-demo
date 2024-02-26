package config

import "kis-flow-demo/common"

type KisFuncParam struct {
	FName string `yaml:"fname"`
	param FParam `yaml:"params"`
}

type KisFlowConfig struct {
	KisType  string         `yaml:"kistype"`
	Status   uint           `yaml:"status"`
	FlowName string         `yaml:"flow_name"`
	Flows    []KisFuncParam `yaml:"flows"`
}

func NewKisFlowConfig(flowName string, enable common.KisOnOff) *KisFlowConfig {
	return &KisFlowConfig{
		KisType:  string(common.KisFlow),
		Status:   uint(enable),
		FlowName: flowName,
		Flows:    nil,
	}
}

func (f *KisFlowConfig) AppendFunctionConfig(flows KisFuncParam) {
	f.Flows = append(f.Flows, flows)
}
