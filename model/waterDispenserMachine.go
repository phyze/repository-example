package model

import "time"

type CC float32

type Pusher interface {
	CalculateCC(time.Duration) CC
}

func NewWaterDispenser() Pusher {
	return &waterDispenser{
		macPressRate: 0.2,
	}
}

type waterDispenser struct {
	macPressRate float32
}

func (wd *waterDispenser) CalculateCC(t time.Duration) CC {
	i := float32(t.Milliseconds())
	i = i * wd.macPressRate
	return CC(i)
}
