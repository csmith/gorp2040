package main

import (
	"machine"
	"time"
)

type event int

const (
	eventLoadCellReady event = iota
	eventRotaryEncoderMovedLeft
	eventRotaryEncoderMovedRight

	eventActionButtonPressed
)

func main() {
	time.Sleep(time.Second)

	loadWeightSettings()

	configureDisplay()

	data := make(chan event, 10)

	monitorAmp(data)

	_ = rotaryEncoderA.SetInterrupt(machine.PinFalling, func(machine.Pin) {
		if rotaryEncoderB.Get() {
			data <- eventRotaryEncoderMovedRight
		} else {
			data <- eventRotaryEncoderMovedLeft
		}
	})

	_ = rotaryEncoderAction.SetInterrupt(machine.PinFalling, func(machine.Pin) {
		data <- eventActionButtonPressed
	})

	for {
		event := <-data

		switch event {
		case eventLoadCellReady:
			readWeight()
			renderScreen()
			monitorAmp(data)
		case eventRotaryEncoderMovedLeft:
			handleLeftButton()
		case eventRotaryEncoderMovedRight:
			handleRightButton()
		case eventActionButtonPressed:
			handleActionButton()
		}
	}
}
