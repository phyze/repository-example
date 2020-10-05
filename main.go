package main

import (
	"fmt"
	"phyze/repository-example/service"
	"time"

	"github.com/opentracing/opentracing-go/log"
)

func main() {
	pushedTime := time.Duration(6000)
	wdsvc := service.NewWaterDispenser()
	cc, err := wdsvc.Push(pushedTime)
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("volume of water that pushed : %v", cc)
}
