package conn

import (
	"context"
	"kis-flow-demo/common"
	"kis-flow-demo/config"
	"kis-flow-demo/id"
	"kis-flow-demo/kis"
	"sync"
)

type KisConnector struct {
	CId      string
	CName    string
	Conf     *config.KisConnConfig
	onceInit sync.Once
}

func (c *KisConnector) Init() error {
	var err error
	c.onceInit.Do(func() {
		err = kis.Pool().CallConnInit(c)
	})
	return err
}

func (c *KisConnector) Call(ctx context.Context, flow kis.Flow, args interface{}) error {
	if err := kis.Pool().CallConnector(ctx, c, flow, args); err != nil {
		return err
	}
	return nil
}

func (c *KisConnector) GetId() string {
	return c.CId
}

func (c *KisConnector) GetName() string {
	return c.CName
}

func (c *KisConnector) GetConfig() *config.KisConnConfig {
	return c.Conf
}

func NewKisConnector(config *config.KisConnConfig) *KisConnector {
	return &KisConnector{
		CId:   id.KisID(common.KisIdTypeConnector),
		CName: config.CName,
		Conf:  config,
	}
}
