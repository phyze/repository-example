package persistence

import (
	"fmt"
	"phyze/repository-example/model"
)

type PumpingRepo interface {
	UpdateWaterVolume(cc model.CC) error
	AddWaterVolume(cc model.CC) error
}

func NewPumping() PumpingRepo {
	return &pumping{}
}

type pumping struct {
}

func (p *pumping) AddWaterVolume(cc model.CC) error {
	fmt.Println("add water to tank ",cc)
	return nil
}

func (p *pumping) UpdateWaterVolume(cc model.CC) error {
	fmt.Println("balance water in tank ",cc)
	return nil
}
