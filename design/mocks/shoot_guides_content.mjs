// Captures the guide sections the 2026-11-09 content pass touches, from a running preview, at 1440 and 390:
// the Hunter class landing (Races), Survival (Stat priority, Races, Leveling) and one guide per class's
// Stat priority section, so a reviewer reads the rendered prose rather than the diff.
//   (web) npm run preview -- --port 4517   then   node design/mocks/shoot_guides_content.mjs http://localhost:4517 <out dir>
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const BASE = process.argv[2] ?? "http://localhost:4517";
const OUT = (process.argv[3] ?? "design/reviews/2026-11-09-guides-content/after") + "/";
mkdirSync(OUT, { recursive: true });
const SHOTS = [
  ["hunter-index", "/guides/hunter", null],
  ["hunter-survival-stats", "/guides/hunter/survival", "#stat-priority"],
  ["hunter-survival-races", "/guides/hunter/survival", "#races"],
  ["hunter-survival-leveling", "/guides/hunter/survival", "#leveling"],
  ["warrior-fury-races", "/guides/warrior/fury", "#races"],
  ["mage-frost-races", "/guides/mage/frost", "#races"],
  ["rogue-subtlety-stats", "/guides/rogue/subtlety", "#stat-priority"],
  ["druid-restoration-stats", "/guides/druid/restoration", "#stat-priority"],
];
const browser = await chromium.launch();
for (const width of [1440, 390]) {
  const page = await browser.newPage({ viewport: { width, height: 1000 }, deviceScaleFactor: 1 });
  for (const [name, path, anchor] of SHOTS) {
    await page.goto(`${BASE}${path}`, { waitUntil: "networkidle" });
    const target = anchor ? page.locator(anchor).first() : page.locator("main").first();
    await target.screenshot({ path: `${OUT}${name}-${width}.png` });
  }
  await page.close();
}
await browser.close();
console.log(`wrote ${SHOTS.length * 2} captures to ${OUT}`);
