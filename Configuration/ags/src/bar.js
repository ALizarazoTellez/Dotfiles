import { Astal, Widget } from "astal/gtk3";
import { Variable, GLib } from "astal";

export function Bar(monitor) {
  const { TOP, LEFT, RIGHT } = Astal.WindowAnchor;

  return new Widget.Window(
    { monitor: monitor, anchor: TOP | LEFT | RIGHT },

    new Widget.CenterBox(
      {},
      new Widget.Label({ label: "Workspace Name" }),
      Time(),
      new Widget.Label({ label: "Wifi, Battery & System Tray" }),
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
