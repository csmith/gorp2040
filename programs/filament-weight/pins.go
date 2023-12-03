package main

import "machine"

var oledClock = machine.GPIO15

func init() {
	oledClock.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
}

var oledData = machine.GPIO14

func init() {
	oledData.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
}

var oledBus = machine.I2C1

func init() {
	_ = oledBus.Configure(machine.I2CConfig{
		SDA: oledData,
		SCL: oledClock,
	})
}

var ampData = machine.GPIO26

func init() {
	ampData.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})
}

var ampClock = machine.GPIO27

func init() {
	ampClock.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
}

var rotaryEncoderAction = machine.GPIO5

func init() {
	rotaryEncoderAction.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})
}

var rotaryEncoderA = machine.GPIO6

func init() {
	rotaryEncoderA.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})
}

var rotaryEncoderB = machine.GPIO7

func init() {
	rotaryEncoderB.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})
}
