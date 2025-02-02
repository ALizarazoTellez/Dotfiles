import { Astal, Widget, Gtk } from "astal/gtk3";
import { Variable, GLib, bind } from "astal";

import Tray from "gi://AstalTray";

export function Bar(monitor) {
  const { TOP, LEFT, RIGHT } = Astal.WindowAnchor;

  return new Widget.Window(
    { monitor: monitor, anchor: TOP | LEFT | RIGHT },

    new Widget.CenterBox(
      {},
      new Widget.Label({ label: "Workspace Name" }),
      Time(),
      SystemStatus(),
    ),
  );
}

function Time() {
  const time = Variable("").poll(1000, () =>
    GLib.DateTime.new_now_local().format("%H:%M"),
  );

  return new Widget.Label({
    label: time((v) => v),
    onDestroy: () => {
      time.drop();
    },
  });
}

function SystemStatus() {
  const tray = Tray.get_default();

  return new Widget.Box(
    { halign: Gtk.Align.END },
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
            },
            new Widget.Icon({ gicon: bind(item, "gicon") }),
          ),
      ),
    ),
  );
}
