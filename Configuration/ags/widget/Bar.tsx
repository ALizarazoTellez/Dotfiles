import { App, Variable, Astal, Gtk, Gdk, GLib, bind } from "astal";

import Tray from "gi://AstalTray";
import Hyprland from "gi://AstalHyprland";
import Battery from "gi://AstalBattery";

export default function Bar(monitor: number) {
  return (
    <window
      name="Bar"
      application={App}
      className="Bar"
      monitor={monitor}
      layer={Astal.Layer.OVERLAY}
      anchor={
        Astal.WindowAnchor.TOP |
        Astal.WindowAnchor.LEFT |
        Astal.WindowAnchor.RIGHT
      }
    >
      <centerbox>
        <Workspaces />
        <Time />
        <box halign={Gtk.Align.END}>
          <SysTray />
          <BatteryLevel />
        </box>
      </centerbox>
    </window>
  );
}

function Workspaces() {
  const hyprland = Hyprland.get_default();

  return (
    <box className="Workspaces">
      {bind(hyprland, "workspaces").as((workspace) =>
        workspace
          .sort((a, b) => a.id - b.id)
          .filter((ws) => ws.id > 0)
          .map((ws) => (
            <button
              className={bind(hyprland, "focusedWorkspace").as((focusedWs) =>
                ws === focusedWs ? "focused" : "",
              )}
              onClicked={() => ws.focus()}
            >
              {ws.id}
            </button>
          )),
      )}
    </box>
  );
}

function Time() {
  const time = Variable<string>("").poll(
    1000,
    () => GLib.DateTime.new_now_local().format("%H:%M")!,
  );

  const date = Variable<string>("").poll(
    1000 * 60 * 60, // Poll every hour.
    () => GLib.DateTime.new_now_local().format("%A, %B %d, %Y")!,
  );

  return (
    <label
      className="Time"
      onDestroy={() => time.drop()}
      label={time()}
      tooltipText={date()}
    />
  );
}

function SysTray() {
  const tray = Tray.get_default();

  return (
    <box>
      {bind(tray, "items").as((items) =>
        items.map((item) => {
          if (item.iconThemePath) App.add_icons(item.iconThemePath);

          const menu = item.create_menu();

          return (
            <button
              tooltipMarkup={bind(item, "tooltipMarkup")}
              onDestroy={() => menu?.destroy()}
              className={"SysTray-item"}
              onClickRelease={(self) => {
                menu?.popup_at_widget(
                  self,
                  Gdk.Gravity.SOUTH,
                  Gdk.Gravity.NORTH,
                  null,
                );
              }}
            >
              <icon gIcon={bind(item, "gicon")} />
            </button>
          );
        }),
      )}
    </box>
  );
}

function BatteryLevel() {
  const bat = Battery.get_default();

  return (
    <box className="Battery" visible={bind(bat, "isPresent")}>
      <icon icon={bind(bat, "batteryIconName")} />
      <label
        label={bind(bat, "percentage").as((p) => ` ${Math.floor(p * 100)}%`)}
      />
    </box>
  );
}
