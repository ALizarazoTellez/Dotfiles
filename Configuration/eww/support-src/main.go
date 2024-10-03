package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Requires at least one arg!")
	}

	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	switch cmd := os.Args[1]; cmd {
	case "volume":
		cmdVolume(args)
	case "battery":
		cmdBattery(args)
	default:
		log.Fatalf("Invalid command: %q", cmd)
	}
}

type iconRange []struct {
	max  int
	icon rune
}

func (ir iconRange) icon(level int) rune {
	for _, icon := range ir {
		if level <= icon.max {
			return icon.icon
		}
	}

	return '*'
}

func cmdVolume([]string) {
	rawVolume, err := exec.Command("wpctl", "get-volume", "@DEFAULT_AUDIO_SINK@").Output()
	if err != nil {
		log.Fatal(err)
	}

	rawVolume = rawVolume[len("Volume: ") : len(rawVolume)-1]
	floatVolume, err := strconv.ParseFloat(string(rawVolume), 32)
	if err != nil {
		log.Fatal(err)
	}
	volume := int(floatVolume * 100)

	icons := iconRange{
		{0, '󰝟'}, {35, '󰕿'}, {70, '󰖀'}, {200, '󰕾'},
	}
	fmt.Printf("%c %d%%", icons.icon(volume), volume)
}

func cmdBattery([]string) {
	rawCapacity, err := os.ReadFile("/sys/class/power_supply/BAT1/capacity")
	if err != nil {
		log.Fatal(err)
	}

	capacity, err := strconv.Atoi(string(rawCapacity[:len(rawCapacity)-1]))
	if err != nil {
		log.Fatal(err)
	}

	rawStatus, err := os.ReadFile("/sys/class/power_supply/BAT1/status")
	if err != nil {
		rawStatus = []byte("Discharging")
		log.Print(err)
	}
	status := string(rawStatus[:len(rawStatus)-1])

	log.Println("Capacity:", capacity)
	log.Println("Status:", status)

	dischargingIcons := iconRange{
		{10, '󰁺'}, {20, '󰁻'}, {30, '󰁼'}, {40, '󰁽'}, {50, '󰁾'}, {60, '󰁿'}, {70, '󰂀'}, {80, '󰂁'}, {90, '󰂂'}, {9999, '󱈑'},
	}
	chargingIcons := iconRange{
		{10, '󰢜'}, {20, '󰂆'}, {30, '󰂇'}, {40, '󰂈'}, {50, '󰢝'}, {60, '󰂉'}, {70, '󰢞'}, {80, '󰂊'}, {90, '󰂋'}, {9999, '󰂅'},
	}

	icons := dischargingIcons
	if status == "Charging" {
		icons = chargingIcons
	}

	fmt.Printf("%c %d%%", icons.icon(capacity), capacity)
}
