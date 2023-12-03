package main

import (
	"fmt"
	"github.com/csmith/gorp2040/drivers/sh1106"
	"image/color"
	"strconv"
	"tinygo.org/x/drivers"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/proggy"
)

var oled = sh1106.NewI2C(oledBus)

func configureDisplay() {
	oled.Configure(sh1106.Config{
		Width:  128,
		Height: 64,
	})
}

var black = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var white = color.RGBA{R: 255, G: 255, B: 255, A: 255}

func renderScreen() {
	// Clear the screen
	_ = tinydraw.FilledRectangle(&oled, 0, 0, 128, 64, black)

	// Draw a spool with a rough representation of the weight
	tinydraw.Circle(&oled, 30, 32, 30, white)
	r := 5 + max(1, min(25, int16(0.025*float64(currentWeight))))
	tinydraw.FilledCircle(&oled, 30, 32, r, white)
	tinydraw.FilledCircle(&oled, 30, 32, 5, black)
	tinydraw.Line(&oled, 30-r, 32, 30-r, 64, white)

	// Status text
	writeLineCentered(&oled, &freesans.Bold12pt7b, 64, 128, 20, strconv.Itoa(currentWeight)+"g", white)
	writeLineCentered(&oled, &proggy.TinySZ8pt7b, 64, 128, 38, strconv.Itoa(spool)+"g spool", white)
	writeLineCentered(&oled, &proggy.TinySZ8pt7b, 52, 128, 61, "<"+currentMenuEntry.Name()+">", white)

	if err := oled.Display(); err != nil {
		fmt.Printf("Failed to display: %v\n", err)
	}
}

func writeLineCentered(display drivers.Displayer, font tinyfont.Fonter, x1, x2, y int16, str string, c color.RGBA) {
	width, _ := tinyfont.LineWidth(font, str)
	offset := ((x2 - x1) - int16(width)) / 2
	tinyfont.WriteLine(display, font, x1+offset, y, str, c)
}
