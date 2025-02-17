#!/usr/bin/env fish

gjs -m ~/.config/astal/astal.js &
poweralertd &
hypridle &
wayland-pipewire-idle-inhibit &
hyprpaper &
~/.config/hypr/support wallpaper-service &
playerctld &
