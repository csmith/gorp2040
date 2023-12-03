package main

import (
	"fmt"
	"github.com/csmith/gorp2040/drivers/hx711"
)

var (
	offset int
	scale  int
	spool  int

	currentRawValue int
	currentWeight   int
)

func loadWeightSettings() {
	offset = readOffset()
	scale = readScale()
	spool = readSpool()
	println("Loaded weight settings: offset=", offset, " scale=", scale, " spool=", spool)
}

var amp = hx711.New(ampClock, ampData)

func monitorAmp(data chan<- event) {
	if err := amp.Wait(func() {
		data <- eventLoadCellReady
	}); err != nil {
		fmt.Printf("Failed to read from amp: %v\n", err)
	}
}

func readWeight() {
	currentRawValue = amp.Read(hx711.ChannelA_Gain128)
	currentWeight = (currentRawValue-offset)/scale - spool
}
