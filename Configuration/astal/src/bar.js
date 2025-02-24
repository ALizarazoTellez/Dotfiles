import { Astal, App, Widget, Gtk } from "astal/gtk3";
import { Variable, GLib, bind } from "astal";

import Hyprland from "gi://AstalHyprland";
import Tray from "gi://AstalTray";
import Battery from "gi://AstalBattery";

export function Bar(monitor) {
  const { TOP, LEFT, RIGHT } = Astal.WindowAnchor;

  return new Widget.Window(
    {
      name: "StatusBar",
      setup: (self) => App.add_window(self),
      monitor: monitor,
      anchor: TOP | LEFT | RIGHT,
      layer: Astal.Layer.OVERLAY,
      className: "Bar",
    },

    new Widget.CenterBox(
      {
        className: "Bar-CenterBox",
      },
      Workspace(),
      Time(),
      SystemStatus(),
    ),
  );
}

function Workspace() {
  const hyprland = Hyprland.get_default();

  return new Widget.Box(
    {
      className: "Bar-CenterBox-Workspace",
    },
    bind(hyprland, "focusedWorkspace").as((workspace) => workspace.name),
  );
}

function Time() {
  const time = Variable("").poll(1000, () =>
    GLib.DateTime.new_now_local().format("%H:%M"),
  );
  const date = Variable("").poll(1000, () =>
    GLib.DateTime.new_now_local().format("%A, %d %B %Y"),
  );

  return new Widget.Label({
    className: "Bar-CenterBox-Time",
    label: time((v) => v),
    tooltipText: date((v) => v),
    onDestroy: () => {
      time.drop();
      date.drop();
    },
  });
}

function SystemStatus() {
  return new Widget.Box(
    { halign: Gtk.Align.END, className: "Bar-CenterBox-SystemStatus" },
    SystemTray(),
    BatteryStatus(),
  );
}

function SystemTray() {
  const tray = Tray.get_default();

  return new Widget.Box(
    { className: "Bar-CenterBox-SystemStatus-SystemTray" },
    bind(tray, "items").as((items) =>
      items.map(
        (item) =>
          new Widget.MenuButton(
            {
              tooltipMarkup: bind(item, "tooltipMarkup"),
              usePopover: false,
              actionGroup: bind(item, "actionGroup").as((ag) => [
                "dbusmenu",
                ag,
              ]),
              menuModel: bind(item, "menuModel"),
              className: "Bar-CenterBox-SystemStatus-SystemTray-MenuButton",
            },
            new Widget.Icon({
              gicon: bind(item, "gicon"),
              className:
                "Bar-CenterBox-SystemStatus-SystemTray-MenuButton-Icon",
            }),
          ),
      ),
    ),
  );
}

function BatteryStatus() {
  const battery = Battery.get_default();

  return new Widget.Box(
    { className: "Bar-CenterBox-SystemStatus-BatteryStatus" },
    new Widget.Icon({ icon: bind(battery, "batteryIconName") }),
    new Widget.Label({
      label: bind(battery, "percentage").as((p) => ` ${Math.floor(p * 100)}%`),
    }),
  );
}
