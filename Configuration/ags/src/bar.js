import { Astal, Widget } from "astal/gtk3";

export function Bar(monitor) {
  const { TOP, LEFT, RIGHT } = Astal.WindowAnchor;

  return new Widget.Window(
    { monitor: monitor, anchor: TOP | LEFT | RIGHT },

    new Widget.CenterBox(
      {},
      new Widget.Label({ label: "Workspace Name" }),
      new Widget.Label({ label: "Time" }),
      new Widget.Label({ label: "Wifi, Battery & System Tray" }),
    ),
  );
}
