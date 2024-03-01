package kis

import (
	"context"
	"kis-flow-demo/common"
)

// FaaS Function as a Service

type Faas func(context.Context, Flow) error

// funcRouter
// key: Function Name
// value: Function 回调自定义业务
type funcRouter map[string]Faas

// flowRouter
// key: Flow Name
// value: Flow
type flowRouter map[string]Flow

type ConnInit func(conn Connector) error

type connInitRouter map[string]ConnInit

type Caas func(context.Context, Connector, Function, Flow, interface{}) error

type connFuncRouter map[string]Caas

type connSL map[common.FMode]connFuncRouter

type connTree map[string]connSL
