package kis

import (
	"context"
	"kis-flow-demo/config"
)

type Connector interface {
	Init() error
	Call(ctx context.Context, flow Flow, args interface{}) error
	GetId() string
	GetName() string
	GetConfig() *config.KisConnConfig
}
