package persistence

import (
	"fmt"
	"phyze/repository-example/model"
)

type PumpingRepo interface {
	RemoveWaterInTank(cc model.CC) error
}

func NewPumping() PumpingRepo {
	return &pumping{}
}

type pumping struct {
}

func (p *pumping) RemoveWaterInTank(cc model.CC) error {
	var volume model.CC
	volume = 50000
	fmt.Printf("balance of water in tank : %v\n", volume)
	fmt.Println("remove water in tank")
	volume = volume - cc
	fmt.Printf("balance of water in tank : %v\n", volume)
	return nil
}
