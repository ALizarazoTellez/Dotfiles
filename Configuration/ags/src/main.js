import { App } from "astal/gtk3";

import { Bar } from "./bar.js";

let bar;

App.start({
  instanceName: "status-bar",

  requestHandler: (request, response) => {
    if (request === "show") {
      bar.visible = true;
      return response("show...");
    }

    if (request === "hide") {
      bar.visible = false;
      return response("hide...");
    }
  },

  main() {
    bar = Bar(0);
    bar.visible = false;
  },
});
