import * as esbuild from "esbuild";

await esbuild.build({
  bundle: true,
  entryPoints: ["src/main.js"],
  outfile: "astal.js",

  platform: "neutral",
  external: ["gi://*", "console", "system"],
  alias: {
    astal: "/usr/share/astal/gjs",
    "astal/*": "/usr/share/astal/gjs/src/*",
  },
});
