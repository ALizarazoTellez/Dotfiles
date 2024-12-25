package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		log.Print("Registering handlers...")
		registerHandlers()
		return
	}

	if os.Args[1] == "toggle-gamemode" {
		log.Print("Toggling gamemode...")
		ToggleGamemode()
	}

	if os.Args[1] == "wallpaper-service" {
		log.Print("Starting wallpaper service...")

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Could not get the user home directory: %q", err)
		}
		wallpapersDir := filepath.Join(homeDir, "Pictures", "Wallpapers", "LaptopCurrent")

		dir, err := os.Open(wallpapersDir)
		if err != nil {
			log.Printf("Wallpapers directory: %q", wallpapersDir)
			log.Fatalf("Could not open the wallpapers directory: %q", err)
		}

		wallpapers, err := dir.Readdirnames(-1)
		if err != nil && len(wallpapers) == 0 {
			log.Fatalf("Unknown error while listing wallpapers: %q", err)
		}
		if err != nil {
			log.Printf("No handled error: %q", err)
		}

		dir.Close()

		var lastLoaded string
		for {
			for _, wallpaper := range wallpapers {
				wallpaper = filepath.Join(wallpapersDir, wallpaper)
				log.Printf("Setting wallpaper: %q", wallpaper)
				if output, err := exec.Command("hyprctl", "hyprpaper", "preload", wallpaper).CombinedOutput(); err != nil {
					log.Printf("No handled error: %q", err)
					log.Printf("%s", output)
					time.Sleep(time.Second)
					continue
				}

				if output, err := exec.Command("hyprctl", "hyprpaper", "wallpaper", ","+wallpaper).CombinedOutput(); err != nil {
					log.Printf("No handled error: %q", err)
					log.Printf("%s", output)
					time.Sleep(time.Second)
					continue
				}

				if lastLoaded != "" {
					if output, err := exec.Command("hyprctl", "hyprpaper", "unload", lastLoaded).CombinedOutput(); err != nil {
						log.Printf("No handled error: %q", err)
						log.Printf("%s", output)
						continue
					}
				}
				lastLoaded = wallpaper

				time.Sleep(time.Second * 60 * 15) // 15 minutes.
			}
		}
	}
}

func registerHandlers() {
	RegisterHandler("activespecial", func(s string) error {
		data := strings.Split(s, ",")
		if len(data) != 2 {
			panic(fmt.Sprintf("Invalid layout: %#v", data))
		}

		// Data{Workspace, Monitor}.
		if data[0] == "special:scratchpad" {
			if err := exec.Command("ags", "request", "show bar").Run(); err != nil {
				return err
			}

			return nil
		}

		if err := exec.Command("ags", "request", "hide bar").Run(); err != nil {
			return err
		}

		return nil
	})

	select {}
}
