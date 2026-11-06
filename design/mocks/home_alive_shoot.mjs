// Captures the three home-alive mocks: node design/mocks/home_alive_shoot.mjs
// -> design/reviews/2026-11-06-home-alive/home-alive-<a|b|c>-<w>.png (first screen at 1440x1100 and 2000x1200, and full page).
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const OUT = new URL("../reviews/2026-11-06-home-alive/", import.meta.url).pathname;
const SRC = new URL("./web/", import.meta.url).pathname;
const SIZES = [[1440, 1100], [2000, 1200]];
mkdirSync(OUT, { recursive: true });
const b = await chromium.launch();
for (const k of ["a", "b", "c"]) {
  for (const [w, h] of SIZES) {
    const pg = await (await b.newContext({ viewport: { width: w, height: h } })).newPage();
    await pg.goto(`file://${SRC}web-home-alive-${k}.html`, { waitUntil: "networkidle" });
    await pg.waitForTimeout(1400); // count-up and slide-in settle
    await pg.screenshot({ path: `${OUT}home-alive-${k}-${w}.png` });
    await pg.screenshot({ path: `${OUT}home-alive-${k}-${w}-full.png`, fullPage: true });
  }
}
await b.close();
