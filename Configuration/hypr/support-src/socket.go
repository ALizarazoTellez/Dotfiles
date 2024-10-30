package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var HyprlandSocket2 = filepath.Join(
	os.Getenv("XDG_RUNTIME_DIR"),
	"hypr",
	os.Getenv("HYPRLAND_INSTANCE_SIGNATURE"),
	".socket2.sock",
)

func RegisterHandler(event string, handler func(string) error) error {
	conn, err := net.Dial("unix", HyprlandSocket2)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(conn)

	go func() {
		for scanner.Scan() {
			e, arguments, _ := strings.Cut(scanner.Text(), ">>")

			if e != event {
				continue
			}

			if err := handler(arguments); err != nil {
				log.Printf("Error in handler for %q: %s\n", event, err)
				return
			}
		}

		if err := scanner.Err(); err != nil {
			log.Println("Error reading Hyprland Socket2:", err)
		}
	}()

	return nil
}
