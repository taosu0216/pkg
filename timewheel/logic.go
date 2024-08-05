package timewheel

import (
	"container/list"
	"fmt"
	"time"
)

func (t *TimeWheel) AddTask(key string, task func(), execTime time.Time) {
	pos, cycle := t.getPosAndCycle(execTime)
	t.addTaskChan <- &taskElement{pos: pos, cycle: cycle, key: key, task: task}
}

func (t *TimeWheel) addTask(task *taskElement) {
	list := t.slots[task.pos]

	if _, ok := t.keyToETask[task.key]; ok {
		t.removeTask(task.key)
	}

	eTask := list.PushBack(task)
	t.keyToETask[task.key] = eTask
}

func (t *TimeWheel) RemoveTask(key string) {
	t.removeTaskChan <- key
}

func (t *TimeWheel) removeTask(key string) {
	eTask, ok := t.keyToETask[key]
	if !ok {
		return
	}
	delete(t.keyToETask, key)
	task, _ := eTask.Value.(*taskElement)
	_ = t.slots[task.pos].Remove(eTask)
}

func (t *TimeWheel) tick() {
	taskList := t.slots[t.curSlot]
	defer t.circularIncr()
	t.execute(taskList)
}

func (t *TimeWheel) execute(taskList *list.List) {
	for e := taskList.Front(); e != nil; {
		task, _ := e.Value.(*taskElement)
		if task.cycle > 0 {
			task.cycle--
			e = e.Next()
			continue
		}
		go func() {
			defer func() {
				if err := recover(); err != nil {
					er := fmt.Sprintf("timeWheel execute panic is: %v", err)
					panic(er)
				}
			}()
			task.task()
		}()
		next := e.Next()
		taskList.Remove(e)
		delete(t.keyToETask, task.key)
		e = next
	}
}
