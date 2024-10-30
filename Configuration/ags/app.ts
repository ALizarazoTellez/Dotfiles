import { App, Gtk } from "astal";
import style from "inline:./style.css";
import Bar from "./widget/Bar";

let bar: Gtk.Widget;

App.start({
  css: style,
  requestHandler(request: string, res: (response: any) => void) {
    if (request == "show bar") {
      bar.visible = true;
      res("visible true");
    } else if (request == "hide bar") {
      bar.visible = false;
      res("visible false");
    }

    res("ok");
  },
  main() {
    bar = Bar(0);
    bar.visible = false;
  },
});
