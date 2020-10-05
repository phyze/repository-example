package model

import (
	"errors"
)

type tank struct {
	balance CC
	capacity CC
}

func (t *tank) removeWaterVolume(cc CC) error {
	if cc == 0 {
		return errors.New("cc can not be zero")
	}
	if t.balance < cc {
		return errors.New("water not enough")
	}
	t.balance = t.balance - cc
	return nil
}

func (t *tank) addWater(cc CC) error {
	if t.balance == t.capacity {
		return errors.New("water full")
	}
	t.balance = t.balance + cc
	return nil
}