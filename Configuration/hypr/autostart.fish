#!/usr/bin/env fish

ags run &
poweralertd &
hypridle &
wayland-pipewire-idle-inhibit &
hyprpaper &
~/.config/hypr/support wallpaper-service &
playerctld &
