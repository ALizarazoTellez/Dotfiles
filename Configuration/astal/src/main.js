import { App } from "astal/gtk3";

import { Bar } from "./bar.js";
import { Mode } from "./mode.js";
import { VolumeLevel } from "./volume.js";
import { bind } from "astal";

import GLib from "gi://GLib";

import WirePlumber from "gi://AstalWp";

let astalDir =
  GLib.getenv("XDG_CONFIG_HOME") ||
  GLib.build_filenamev([GLib.getenv("HOME"), ".config", "astal"]);

let bar;
let mode;

App.start({
  instanceName: "status-bar",

  css: GLib.build_filenamev([astalDir, "src", "style.css"]),

  requestHandler: (request, response) => {
    if (request === "show-bar") {
      bar.visible = true;
      return response("show...");
    }

    if (request === "hide-bar") {
      bar.visible = false;
      return response("hide...");
    }

    if (request === "show-mode") {
      mode.visible = true;
      return response("show...");
    }

    if (request === "hide-mode") {
      mode.visible = false;
      return response("hide...");
    }
  },

  main() {
    bar = Bar(0);
    bar.visible = false;

    mode = Mode(0);
    mode.visible = false;

    let volumeLevel = VolumeLevel(0);
    volumeLevel.visible = false;

    let lastShown = null;
    bind(WirePlumber.get_default().audio.defaultSpeaker, "volume").subscribe(
      () => {
        volumeLevel.visible = true;
        lastShown = Date.now();
        setTimeout(() => {
          if (Date.now() - lastShown >= 1000) {
            volumeLevel.visible = false;
          }
        }, 1000);
      },
    );
  },
});
