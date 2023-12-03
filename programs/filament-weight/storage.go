package main

import (
	"fmt"
	"machine"
	"os"
	"strconv"
	"tinygo.org/x/tinyfs/littlefs"
)

var (
	blockDevice = machine.Flash
	filesystem  = littlefs.New(blockDevice)
	storageInit = false
)

func initFileSystem() {
	if !storageInit {
		filesystem.Configure(&littlefs.Config{
			CacheSize:     512,
			LookaheadSize: 512,
			BlockCycles:   100,
		})

		err := filesystem.Mount()
		if err != nil {
			fmt.Printf("Failed to mount filesystem: %v\n", err)
			err = filesystem.Format()
			if err != nil {
				fmt.Printf("Failed to format filesystem: %v\n", err)
			}
			err = filesystem.Mount()
			if err != nil {
				fmt.Printf("Failed to re-mount filesystem: %v\n", err)
			}
		}

		println("File system init complete")
		storageInit = true
	}
}

func readOffset() int {
	return readFile("/offset", 0)
}

func readScale() int {
	return readFile("/scale", 1)
}

func readSpool() int {
	return readFile("/spool", 100)
}

func writeOffset(value int) error {
	return writeFile("/offset", value)
}

func writeScale(value int) error {
	return writeFile("/scale", value)
}

func writeSpool(value int) error {
	return writeFile("/spool", value)
}

func readFile(file string, fallback int) int {
	initFileSystem()
	f, err := filesystem.Open(file)
	if err != nil {
		fmt.Printf("Failed to open %s: %v\n", file, err)
		return fallback
	}

	defer f.Close()

	b := make([]byte, 16)
	n, err := f.Read(b)
	if err != nil {
		fmt.Printf("Failed to read %s: %v\n", file, err)
		return fallback
	}

	fmt.Printf("Read data from %s: %s\n", file, string(b[:n]))
	v, err := strconv.Atoi(string(b[:n]))
	if err != nil {
		fmt.Printf("Failed to parse %s: %v\n", file, err)
		return fallback
	}
	return v
}

func writeFile(file string, value int) error {
	initFileSystem()
	f, err := filesystem.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		fmt.Printf("Failed to open %s: %v\n", file, err)
		return err
	}

	defer f.Close()

	_, err = f.Write([]byte(strconv.Itoa(value)))
	if err != nil {
		fmt.Printf("Failed to write %s: %v\n", file, err)
	} else {
		fmt.Printf("Wrote data to %s: %d\n", file, value)
	}
	return err
}
