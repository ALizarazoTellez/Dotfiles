package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

type BatteryStatus string

const (
	BatteryStatusCharging    BatteryStatus = "Charging"
	BatteryStatusDischarging BatteryStatus = "Discharging"
)

func BatteryInfo() (capacity int, status BatteryStatus) {
	rawCapacity, err := os.ReadFile("/sys/class/power_supply/BAT1/capacity")
	if err != nil {
		log.Fatal(err)
	}

	// Remove line break at the end, and convert it to [int].
	capacity, err = strconv.Atoi(string(rawCapacity[:len(rawCapacity)-1]))
	if err != nil {
		log.Fatal(err)
	}

	rawStatus, err := os.ReadFile("/sys/class/power_supply/BAT1/status")
	if err != nil {
		log.Fatal(err)
	}

	// Remove line break at the end.
	status = BatteryStatus(rawStatus[:len(rawStatus)-1])

	return capacity, status
}

func VolumeInfo() (volume int) {
	rawVolume, err := exec.Command("wpctl", "get-volume", "@DEFAULT_AUDIO_SINK@").Output()
	if err != nil {
		log.Fatal(err)
	}

	// Remove the `Volume: ` prefix and the line break at the end.
	rawVolume = rawVolume[len("Volume: ") : len(rawVolume)-1]

	// WirePlumber reports the volume as something like `0.56`.
	floatVolume, err := strconv.ParseFloat(string(rawVolume), 32)
	if err != nil {
		log.Fatal(err)
	}

	return int(floatVolume * 100)
}
