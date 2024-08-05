package timewheel

import (
	"container/list"
	"fmt"
	"time"
)

func NewTimeWheel(slotNum int, interval time.Duration) *TimeWheel {
	if slotNum <= 0 {
		slotNum = 10
	}
	if interval <= 0 {
		interval = time.Second
	}

	t := &TimeWheel{
		interval:       interval,
		ticker:         time.NewTicker(interval),
		stopc:          make(chan struct{}),
		addTaskChan:    make(chan *taskElement),
		removeTaskChan: make(chan string),
		slots:          make([]*list.List, 0, slotNum),
		keyToETask:     make(map[string]*list.Element),
	}

	for i := 0; i < slotNum; i++ {
		t.slots = append(t.slots, list.New())
	}

	go t.run()
	return t
}

func (t *TimeWheel) run() {
	defer func() {
		if err := recover(); err != nil {
			er := fmt.Sprintf("timeWheel run panic is: %v", err)
			panic(er)
		}
	}()

	for {
		select {
		case <-t.stopc:
			return
		case <-t.ticker.C:
			t.tick()
		case task := <-t.addTaskChan:
			t.addTask(task)
		case key := <-t.removeTaskChan:
			t.removeTask(key)
		}
	}
}

func (t *TimeWheel) Stop() {
	t.Do(func() {
		t.ticker.Stop()
		close(t.stopc)
	})
}
