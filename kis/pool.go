package kis

import (
	"context"
	"errors"
	"fmt"
	"kis-flow-demo/log"
	"sync"
)

var _poolOnce sync.Once

type kisPool struct {
	fnRouter funcRouter
	fnLock   sync.RWMutex

	flowRouter flowRouter
	flowLock   sync.RWMutex
}

var _pool *kisPool

func Pool() *kisPool {
	_poolOnce.Do(func() {
		_pool = &kisPool{
			fnRouter:   make(funcRouter),
			flowRouter: make(flowRouter),
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
