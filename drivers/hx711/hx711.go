package hx711

import (
	"machine"
	"time"
)

type Channel uint8

const (
	ChannelA_Gain128 Channel = 1
	ChannelB_Gain32  Channel = 2
	ChannelA_Gain64  Channel = 3
)

var (
	clockTime = 200 * time.Nanosecond
)

type HX711 struct {
	clock machine.Pin
	data  machine.Pin
}

func New(clock machine.Pin, data machine.Pin) *HX711 {
	return &HX711{
		clock: clock,
		data:  data,
	}
}

func (h *HX711) tick() {
	h.clock.High()
	time.Sleep(clockTime)
	h.clock.Low()
	time.Sleep(clockTime)
}

func (h *HX711) read() int {
	// Read the 24-bit value
	value := 0
	for i := 0; i < 24; i++ {
		h.tick()
		value = value << 1
		if h.data.Get() {
			value |= 1
		}
	}

	// Deal with negative two's complement numbers
	if value > 0x7fffff {
		value -= 0x1000000
	}

	return value
}

func (h *HX711) selectChannel(channel Channel) {
	for i := 0; i < int(channel); i++ {
		h.tick()
	}
}

func (h *HX711) Ready() bool {
	return !h.data.Get()
}

func (h *HX711) Read(nextChannel Channel) int {
	value := h.read()
	h.selectChannel(nextChannel)
	return value
}

func (h *HX711) Wait(receiver func()) error {
	return h.data.SetInterrupt(machine.PinFalling, func(machine.Pin) {
		_ = h.data.SetInterrupt(machine.PinFalling, nil)
		receiver()
	})
}
