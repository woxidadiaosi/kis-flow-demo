package kis

import (
	"context"
	"errors"
	"fmt"
	"kis-flow-demo/common"
	"kis-flow-demo/log"
	"sync"
)

var _poolOnce sync.Once

type kisPool struct {
	fnRouter funcRouter
	fnLock   sync.RWMutex

	flowRouter flowRouter
	flowLock   sync.RWMutex

	cInitRouter connInitRouter
	ciLock      sync.RWMutex

	cTree      connTree
	connectors map[string]Connector
	cLock      sync.RWMutex
}

var _pool *kisPool

func Pool() *kisPool {
	_poolOnce.Do(func() {
		_pool = &kisPool{
			fnRouter:    make(funcRouter),
			flowRouter:  make(flowRouter),
			cInitRouter: make(connInitRouter),
			cTree:       make(connTree),
			connectors:  make(map[string]Connector),
		}
	})
	return _pool
}

func (pool *kisPool) AddFlow(name string, flow Flow) {
	pool.flowLock.Lock()
	defer pool.flowLock.Unlock()
	if _, ok := pool.flowRouter[name]; !ok {
		pool.flowRouter[name] = flow
	} else {
		errString := fmt.Sprintf("Pool AddFlow Repeat FlowName=%s \n", name)
		panic(errString)
	}
}

func (pool *kisPool) GetFlow(name string) Flow {
	pool.flowLock.RLock()
	defer pool.flowLock.RUnlock()

	return pool.flowRouter[name]
}

func (pool *kisPool) Faas(fName string, f Faas) {
	pool.fnLock.Lock()
	defer pool.fnLock.Unlock()

	if _, ok := pool.fnRouter[fName]; !ok {
		pool.fnRouter[fName] = f
	} else {
		errString := fmt.Sprintf("Pool Faas Repeat FuncName = %s \n", fName)
		panic(errString)
	}
}

func (pool *kisPool) CallFunction(ctx context.Context, fName string, flow Flow) error {
	if f, ok := pool.fnRouter[fName]; ok {
		return f(ctx, flow)
	}
	log.Logger().ErrorFX(ctx, "FuncName: %s Can not found in kisPool", fName)

	return errors.New("FuncName: " + fName + " Can not find in NsPool, Not Added.")
}

func (pool *kisPool) CaasInit(cName string, c ConnInit) {
	pool.ciLock.Lock()
	defer pool.ciLock.Unlock()
	if _, ok := pool.cInitRouter[cName]; ok {
		errString := fmt.Sprintf("kisPool Reg CaasInit repeat cName = %s", cName)
		panic(errString)
	} else {
		pool.cInitRouter[cName] = c
	}
	log.Logger().InfoF("Add KisPool CaaSInit CName=%s \n", cName)
}

func (pool *kisPool) CallConnInit(c Connector) error {
	pool.ciLock.RLock()
	defer pool.ciLock.RUnlock()

	init, ok := pool.cInitRouter[c.GetName()]
	if !ok {
		panic(errors.New(fmt.Sprintf("init connector cname = %s not reg..", c.GetName())))
	}
	return init(c)
}

func (pool *kisPool) Caas(cName string, fName string, mode common.FMode, c Caas) {
	pool.cLock.Lock()
	defer pool.cLock.Unlock()

	if _, ok := pool.cTree[cName]; !ok {
		pool.cTree[cName] = make(connSL)
		pool.cTree[cName][common.L] = make(connFuncRouter)
		pool.cTree[cName][common.S] = make(connFuncRouter)
	}

	if _, ok := pool.cTree[cName][mode][fName]; !ok {
		pool.cTree[cName][mode][fName] = c
	} else {
		panic(fmt.Sprintf("Caas repeat cName=%s, fName=%s, mode=%s \n", cName, fName, mode))
	}

	log.Logger().InfoF("Add kisPool Caas cName=%s, fName=%s, mode=%s \n", cName, fName, mode)
}

func (pool *kisPool) CallConnector(ctx context.Context, c Connector, flow Flow, args interface{}) error {
	function := flow.GetThisFunction()
	mode := common.FMode(function.GetConfig().FMode)
	fName := function.GetConfig().FName
	if callback, ok := pool.cTree[c.GetName()][mode][fName]; ok {
		return callback(ctx, c, function, flow, args)
	}
	log.Logger().ErrorFX(ctx, "kisPool cTree can not found cName=%s, mode=%s, function=%s", c.GetName(), mode, fName)
	return errors.New(fmt.Sprintf("kisPool cTree can not found cName=%s, mode=%s, function=%s", c.GetName(), mode, fName))
}
