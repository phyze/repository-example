package main

import (
	"fmt"
	"phyze/repository-example/service"
	"time"
)

func main() {
	wdsvc := service.NewWaterDispenser()
	if err := wdsvc.AddWaterToTank(30000); err != nil {
		panic(err)
	}

	pushedTime := time.Duration(6000)
	cc, err := wdsvc.Push(pushedTime)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The volume of water pumped %v cc", cc)
}
