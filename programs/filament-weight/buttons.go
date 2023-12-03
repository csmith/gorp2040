package main

type state uint8

const (
	stateMenu state = iota
	stateSpoolWeight
	stateCalibrate
)

var currentState = stateMenu

type menuEntry uint8

const (
	menuEntryZero menuEntry = iota
	menuEntrySpoolWeight
	menuEntryCalibrate
	menuEntryMax
)

func (a menuEntry) Name() string {
	switch a {
	case menuEntryZero:
		return "tare/zero"
	case menuEntrySpoolWeight:
		return "sp weight"
	case menuEntryCalibrate:
		return "calibrate"
	default:
		return "???"
	}
}

var currentMenuEntry = menuEntryZero

func handleActionButton() {
	if currentState == stateMenu {
		if currentMenuEntry == menuEntryZero {
			offset = currentRawValue
			_ = writeOffset(offset)
		} else {
			currentState = state(currentMenuEntry)
		}
	} else {
		_ = writeScale(scale)
		_ = writeSpool(spool)
		currentState = stateMenu
	}
}

func handleLeftButton() {
	if currentState == stateMenu {
		currentMenuEntry = (currentMenuEntry + menuEntryMax - 1) % menuEntryMax
	} else if currentState == stateCalibrate {
		scale = max(1, scale-1)
	} else if currentState == stateSpoolWeight {
		spool = max(0, spool-1)
	}
}

func handleRightButton() {
	if currentState == stateMenu {
		currentMenuEntry = (currentMenuEntry + 1) % menuEntryMax
	} else if currentState == stateCalibrate {
		scale++
	} else if currentState == stateSpoolWeight {
		spool++
	}
}
