package main

import (
	"github.com/csmith/gorp2040/drivers/vl53l0x"
	"machine"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	err := machine.I2C1.Configure(machine.I2CConfig{
		SDA: machine.GPIO6,
		SCL: machine.GPIO7,
	})
	if err != nil {
		println("Failed to configure I2C1")
		panic(err)
	}

	dev := vl53l0x.NewVl53l0x()
	if err := dev.Reset(machine.I2C1); err != nil {
		println("Failed to reset")
		panic(err)
	}

	if err := dev.Init(machine.I2C1); err != nil {
		panic(err)
	}

	if err := dev.Config(machine.I2C1, vl53l0x.LongRange, vl53l0x.GoodAccuracy); err != nil {
		panic(err)
	}

	if err := dev.StartContinuous(machine.I2C1, 100); err != nil {
		panic(err)
	}

	for {
		read, err := dev.ReadRangeContinuousMillimeters(machine.I2C1)
		println(read-91, err)
	}
}
