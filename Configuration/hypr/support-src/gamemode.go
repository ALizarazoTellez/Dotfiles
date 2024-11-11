package main

import (
	"bytes"
	"log"
	"os/exec"
)

func ToggleGamemode() {
	if isGamemodeActive() {
		log.Println("Gamemode is active.")
		if err := exec.Command("hyprctl", "reload").Run(); err != nil {
			log.Println("Ignored error:", err)
		}

		return
	}

	log.Println("Gamemode is not active.")
	const code = `
         keyword animations:enabled 0;
         keyword decoration:drop_shadow 0;
         keyword decoration:blur:enabled 0;
         keyword general:gaps_in 0;
         keyword general:gaps_out 0;
         keyword general:border_size 1;
         keyword decoration:rounding 0`

	if err := exec.Command("hyprctl", "--batch", code).Run(); err != nil {
		log.Println("Ignored error:", err)
	}
}

func isGamemodeActive() bool {
	out, err := exec.Command("hyprctl", "getoption", "animations:enabled").Output()
	if err != nil {
		log.Println("Ignored error:", err)
		return false
	}

	// TODO: Use JSON output and parse it.
	return !bytes.Contains(out, []byte("1"))
}
