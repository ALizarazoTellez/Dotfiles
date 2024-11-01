package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	RegisterHandler("activespecial", func(s string) error {
		data := strings.Split(s, ",")
		if len(data) != 2 {
			panic(fmt.Sprintf("Invalid layout: %#v", data))
		}

		// Data{Workspace, Monitor}.
		if data[0] == "special:scratchpad" {
			if err := exec.Command("ags", "-m", "show bar").Run(); err != nil {
				return err
			}

			return nil
		}

		if err := exec.Command("ags", "-m", "hide bar").Run(); err != nil {
			return err
		}

		return nil
	})

	select {}
}
