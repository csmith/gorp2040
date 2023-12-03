# TinyGo libraries and programs for RP2040-based boards

## Drivers

* hx711 - driver for HX711 load cell amplifiers. Supports switching gain after
  every reading (for if you want to alternate between channels), and using
  interrupts to wait for data to be ready.
* sh1106 - driver for SH1106 OLED displays. Copied from the official TinyGo
  repository with a few tweaks to make it work with the driver interface for
  RP2040s.

## Programs

* filament-weight - reads a HX711 load cell amplifier and displays the weight
  on a SH1106 OLED display. Useful for weighing filament for 3D printers.
* led-controller - test for controlling a WS2812 LED strip. 