package main

import "log"

type TargetMode int

const (
	TargetModeMaxLevel TargetMode = iota
)

type Icon struct {
	Target int
	Icon   rune
}

type Icons struct {
	TargetMode TargetMode
	Icons      []Icon
}

func (icons Icons) Icon(target int) rune {
	for i := len(icons.Icons) - 1; i != 0; i-- {
		icon := icons.Icons[i]

		if icon.Target <= target {
			return icon.Icon
		}
	}
	log.Printf("Missing icon for target: %d", target)

	return ''
}

var BatteryIconsCharging = Icons{
	TargetMode: TargetModeMaxLevel,
	Icons: []Icon{
		{10, '󰢜'},
		{20, '󰂆'},
		{30, '󰂇'},
		{40, '󰂈'},
		{50, '󰢝'},
		{60, '󰂉'},
		{70, '󰢞'},
		{80, '󰂊'},
		{90, '󰂋'},
		{100, '󰂅'},
	},
}

var BatteryIconsDischarging = Icons{
	TargetMode: TargetModeMaxLevel,
	Icons: []Icon{
		{10, '󰁺'},
		{20, '󰁻'},
		{30, '󰁼'},
		{40, '󰁽'},
		{50, '󰁾'},
		{60, '󰁿'},
		{70, '󰂀'},
		{80, '󰂁'},
		{90, '󰂂'},
		{100, '󱈑'},
	},
}

var VolumeIcons = Icons{
	TargetMode: TargetModeMaxLevel,
	Icons: []Icon{
		{0, '󰝟'},
		{35, '󰕿'},
		{70, '󰖀'},
		{150, '󰕾'},
	},
}
