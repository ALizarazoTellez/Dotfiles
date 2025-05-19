import { Astal, App, Widget, Gtk } from "astal/gtk3";
import { Variable, GLib, bind } from "astal";

import Hyprland from "gi://AstalHyprland";
import Tray from "gi://AstalTray";
import Battery from "gi://AstalBattery";

export function Mode(monitor) {
  const { LEFT, TOP, RIGHT, BOTTOM } = Astal.WindowAnchor;

  return new Widget.Window(
    {
      name: "Mode",
      setup: (self) => App.add_window(self),
      monitor: monitor,
      anchor: LEFT | TOP | RIGHT | BOTTOM,
      layer: Astal.Layer.OVERLAY,
      className: "Mode",
    },
    new Widget.Box({ css: `border: 3px solid #31748f;` }),
  );
}
