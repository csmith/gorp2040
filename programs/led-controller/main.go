package main

import (
	"machine"
	"time"
	"tinygo.org/x/drivers/ws2812"
)

const leds = 60

func main() {
	pin := machine.GPIO29
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	strip := ws2812.New(pin)

	stars := []float64{0, 0, 0, 0, 0}
	deltas := []float64{1, 0.5, 0.25, 0.125, 0.0625}
	colours := [][]byte{
		{32, 0, 0},
		{0, 32, 0},
		{0, 0, 32},
		{32, 32, 0},
		{32, 0, 32},
		{0, 32, 32},
	}

	for {
		for i := range stars {
			stars[i] += deltas[i]
			if stars[i] >= float64(leds-1) || stars[i] < 0 {
				deltas[i] *= -1
				stars[i] += 2 * deltas[i]
			}
		}

		for i := 0; i < leds; i++ {
			for j := range stars {
				if int(stars[j]) == i {
					writeColors(&strip, colours[j][0], colours[j][1], colours[j][2])
					break
				}
			}

			writeColors(&strip, 0, 0, 0)
		}

		time.Sleep(time.Millisecond * 50)
	}
}

func writeColors(strip *ws2812.Device, r, g, b byte) {
	strip.WriteByte(g)
	strip.WriteByte(r)
	strip.WriteByte(b)
	strip.WriteByte(0)
}
