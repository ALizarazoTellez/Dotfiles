package main

import (
	"fmt"
	"log"
)

func cmdInfo(args []string) {
	if len(args) != 1 {
		log.Fatal("info: requires one arg: volume, battery")
	}

	switch args[0] {
	case "battery":
		capacity, status := BatteryInfo()
		fmt.Println("Capacity:", capacity)
		fmt.Println("Status:", status)

		icons := BatteryIconsCharging
		if status != BatteryStatusCharging {
			icons = BatteryIconsDischarging
		}

		fmt.Printf("Icon: %c\n", icons.Icon(capacity))

	case "volume":
		volume := VolumeInfo()
		fmt.Println("Volume:", volume)
		fmt.Printf("Icon: %c\n", VolumeIcons.Icon(volume))
	}
}

func cmdStyled(args []string) {
	if len(args) != 1 {
		log.Fatal("info: requires one arg: volume, battery")
	}

	switch args[0] {
	case "battery":
		capacity, status := BatteryInfo()
		icons := BatteryIconsCharging
		if status != BatteryStatusCharging {
			icons = BatteryIconsDischarging
		}

		fmt.Printf("%c %d%%", icons.Icon(capacity), capacity)

	case "volume":
		volume := VolumeInfo()
		fmt.Printf("%c %d%%", VolumeIcons.Icon(volume), volume)
	}
}
