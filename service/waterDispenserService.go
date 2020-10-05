package service

import (
	"phyze/repository-example/model"
	"phyze/repository-example/persistence"
	"time"
)

type WaterDispenser interface {
	Push(time.Duration) (model.CC, error)
}

func NewWaterDispenser() WaterDispenser {
	return &waterDispenserService{}
}

type waterDispenserService struct {
}

func (wds *waterDispenserService) Push(ms time.Duration) (model.CC, error) {
	pushedRange := time.Millisecond * ms
	wd := model.NewWaterDispenser()
	cc := wd.CalculateCC(pushedRange)
	pumper := persistence.NewPumping()
	if err := pumper.RemoveWaterInTank(cc); err != nil {
		return 0.0, err
	}
	return cc, nil
}
