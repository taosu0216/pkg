package timewheel

import "time"

func (t *TimeWheel) getPosAndCycle(execTime time.Time) (int, int) {
	delay := int(time.Until(execTime))

	cycle := delay / (len(t.slots) * int(t.interval))

	pos := (t.curSlot + delay/int(t.interval)) % len(t.slots)

	return pos, cycle
}

func (t *TimeWheel) circularIncr() {
	t.curSlot = (t.curSlot + 1) % len(t.slots)
}
