package kis

import (
	"context"
	"kis-flow-demo/config"
)

type Function interface {
	Call(ctx context.Context, flow Flow) error

	SetConfig(s *config.KisFuncConfig) error

	GetConfig() *config.KisFuncConfig

	SetFlow(f Flow) error

	GetFlow() Flow

	AddConnector(c Connector) error

	GetConnector() Connector

	CreateId()

	GetId() string

	GetPrevId() string

	GetNextId() string

	Next() Function

	Prev() Function

	SetN(f Function)

	SetP(f Function)
}
