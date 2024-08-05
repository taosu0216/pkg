package timewheel

import (
	"container/list"
	"sync"
	"time"
)

type TimeWheel struct {
	sync.Once
	interval time.Duration
	ticker   *time.Ticker

	stopc chan struct{}

	addTaskChan    chan *taskElement
	removeTaskChan chan string

	slots      []*list.List
	curSlot    int
	keyToETask map[string]*list.Element
}

type taskElement struct {
	key   string
	task  func()
	pos   int
	cycle int
}
