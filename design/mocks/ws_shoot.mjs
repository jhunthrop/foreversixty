// Renders a web-layout mock at the given widths: node design/mocks/ws_shoot.mjs web-bis 1440,1920,2000,390 [height]
// -> design/mocks/renders/web/<name>-<w>.png (full page) and <name>-<w>-fold.png (first screen). Playwright is found through web/.
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const [name, widths, vh] = process.argv.slice(2);
const OUT = new URL("./renders/web/", import.meta.url).pathname;
const SRC = new URL("./web/", import.meta.url).pathname;
mkdirSync(OUT, { recursive: true });
const b = await chromium.launch();
for (const w of widths.split(",").map(Number)) {
  const pg = await (await b.newContext({ viewport: { width: w, height: Number(vh) || (w < 500 ? 844 : 1000) } })).newPage();
  await pg.goto(`file://${SRC}${name}.html`, { waitUntil: "networkidle" });
  await pg.waitForTimeout(400);
  await pg.screenshot({ path: `${OUT}${name}-${w}.png`, fullPage: true });
  await pg.screenshot({ path: `${OUT}${name}-${w}-fold.png`, fullPage: false });
}
await b.close();
