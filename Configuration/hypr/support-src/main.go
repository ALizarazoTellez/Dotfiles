package main

import (
	"os/exec"
)

func main() {
	RegisterHandler("createworkspace", func(workspace string) error {
		if workspace != "special:scratchpad" {
			return nil
		}

		if err := exec.Command("ags", "-m", "show bar").Run(); err != nil {
			return err
		}

		return nil
	})

	RegisterHandler("destroyworkspace", func(workspace string) error {
		if workspace != "special:scratchpad" {
			return nil
		}

		if err := exec.Command("ags", "-m", "hide bar").Run(); err != nil {
			return err
		}

		return nil
	})

	select {}
}
