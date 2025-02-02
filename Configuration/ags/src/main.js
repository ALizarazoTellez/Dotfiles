import { App } from "astal/gtk3";

import { Bar } from "./bar.js";

App.start({
  instanceName: "top-bar",
  main() {
    Bar(0);
  },
});
