// Renders the source-line before/after mocks (ws_bis_labels.py) at 1440 and 390: the full page and the gear block.
//   python3 design/mocks/ws_bis_labels.py && node design/mocks/shoot_bis_labels.mjs <out dir>
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const HERE = new URL("./web/", import.meta.url).pathname;
const OUT = (process.argv[2] ?? "design/reviews/2026-11-09-feedback-ai-look/labels") + "/";
mkdirSync(OUT, { recursive: true });
const STATES = ["alliance-before", "alliance-after", "horde-before", "horde-after"];
const WIDTHS = [1440, 390];
const browser = await chromium.launch();
for (const width of WIDTHS) {
  const page = await browser.newPage({ viewport: { width, height: 1000 }, deviceScaleFactor: 2 });
  for (const state of STATES) {
    await page.goto(`file://${HERE}web-bis-labels-${state}.html`);
    await page.waitForLoadState("networkidle");
    await page.screenshot({ path: `${OUT}${state}-${width}.png`, fullPage: true });
    const gear = page.locator('[aria-label="Gear"], .plate').first();
    await gear.screenshot({ path: `${OUT}${state}-${width}-gear.png` });
  }
  await page.close();
}
await browser.close();
console.log(`wrote ${STATES.length * WIDTHS.length * 2} captures to ${OUT}`);
