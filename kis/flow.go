package kis

import (
	"context"
	"kis-flow-demo/config"
)

type Flow interface {
	Run(ctx context.Context) error
	Link(fConf *config.KisFuncConfig, fParas config.FParam) error
}
