package model

import (
	"time"
)

type CC float32
type Balance float32

type Pusher interface {
	Push(duration time.Duration) (CC,error)
	AddWaterToTank(cc CC) error
	Balance() Balance
}

func NewWaterDispenser() Pusher {
	return &waterDispenser{
		macPressRate: 0.2,
		tank: &tank{
			balance: 0,
			capacity: 30000,
		},
	}
}

type waterDispenser struct {
	tank *tank
	macPressRate float32
}

func (wd *waterDispenser) Balance() Balance {
	return Balance(wd.tank.balance)
}

func (wd *waterDispenser) AddWaterToTank(cc CC) error {
	if err := wd.tank.addWater(cc); err != nil {
		return err
	}
	return nil
}

func (wd *waterDispenser) Push(d time.Duration) (CC, error) {

	cc := wd.calculateCC(d)
	if err := wd.tank.removeWaterVolume(cc); err != nil {
		return 0,err
	}
	return cc,nil
}

func (wd *waterDispenser) calculateCC(t time.Duration) CC {
	i := float32(t.Milliseconds())
	i = i * wd.macPressRate
	return CC(i)
}
