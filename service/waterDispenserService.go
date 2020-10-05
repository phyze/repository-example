package service

import (
	"phyze/repository-example/model"
	"phyze/repository-example/persistence"
	"time"
)

type WaterDispenser interface {
	AddWaterToTank(model.CC) error
	Push(time.Duration) (model.CC, error)
}

func NewWaterDispenser() WaterDispenser {
	return &waterDispenserService{
		machine: model.NewWaterDispenser(),
	}
}

type waterDispenserService struct {
	machine model.Pusher
}

func (wds *waterDispenserService) AddWaterToTank(cc model.CC) error {
	if err := wds.machine.AddWaterToTank(cc); err != nil {
		return err
	}
	pumper := persistence.NewPumping()
	return pumper.AddWaterVolume(cc)
}

func (wds *waterDispenserService) Push(ms time.Duration) (model.CC, error) {
	pushedRange := time.Millisecond * ms
	cc,err := wds.machine.Push(pushedRange)
	if err != nil {
		return 0,err
	}
	pumper := persistence.NewPumping()
	balance := model.CC(wds.machine.Balance())
	if err := pumper.UpdateWaterVolume(balance); err != nil {
		return 0, err
	}
	return cc, nil
}


