import { Astal, App, Widget } from "astal/gtk3";

import { bind } from "astal";

import WirePlumber from "gi://AstalWp";

export function VolumeLevel(monitor) {
  const { BOTTOM } = Astal.WindowAnchor;

  const speaker = WirePlumber.get_default().audio.defaultSpeaker;

  return new Widget.Window(
    {
      name: "VolumeLevel",
      setup: (self) => App.add_window(self),
      monitor: monitor,
      anchor: BOTTOM,
    },

    new Widget.Label({
      label: bind(speaker, "volume").as(
        (v) => `Volume: ${Math.floor(v * 100)}%`,
      ),
    }),
  );
}
