package kis

import (
	"context"
	"kis-flow-demo/common"
	"kis-flow-demo/config"
)

type Flow interface {
	Run(ctx context.Context) error
	Link(fConf *config.KisFuncConfig, fParas config.FParam) error
	CommitRow(row interface{}) error

	Input() common.KisRowArr

	GetName() string

	GetThisFunction() Function
	GetThisFunctionConf() *config.KisFuncConfig
}
