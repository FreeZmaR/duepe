package session

import (
	"sync"
	"sync/atomic"
)

const sessionIDDelta = 1

type managerController struct {
	nextID atomic.Int64
	pool   *sync.Pool
}

func newManagerController() *managerController {
	controller := &managerController{}
	pool := sync.Pool{New: func() any {
		return controller.nextID.Add(sessionIDDelta)
	}}

	controller.pool = &pool

	return controller
}

func (c *managerController) getID() int64 {
	return c.pool.Get().(int64)
}

func (c *managerController) putID(id int64) {
	c.pool.Put(id)
}
