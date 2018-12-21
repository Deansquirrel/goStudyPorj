package lib

import (
	"errors"
	"time"
)

type myGoTickets struct {
	total    uint32
	ticketCh chan struct{}
	active   bool
}

func NewMyGoTickets(total uint32) (IGoTickets, error) {
	gt := myGoTickets{}
	if !gt.init(total) {
		return nil, errors.New("GoTickets can not init")
	}
	return &gt, nil
}

func (gt *myGoTickets) init(total uint32) bool {
	if gt.active {
		return false
	}
	if total <= 0 {
		return false
	}
	ch := make(chan struct{}, total)
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
	gt.ticketCh = ch
	gt.total = total
	gt.active = true
	return true
}

func (gt *myGoTickets) Take() {
	if gt.Remainder() > 0 {
		<-gt.ticketCh
	} else {
		time.Sleep(time.Millisecond)
		gt.Take()
	}
}

func (gt *myGoTickets) Return() {
	if gt.Remainder() < gt.total {
		gt.ticketCh <- struct{}{}
	}
}

func (gt *myGoTickets) Active() bool {
	return gt.active
}

func (gt *myGoTickets) Total() uint32 {
	return gt.total
}

func (gt *myGoTickets) Remainder() uint32 {
	return uint32(len(gt.ticketCh))
}
