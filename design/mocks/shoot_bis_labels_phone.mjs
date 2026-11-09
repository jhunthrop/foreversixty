// The 390 before/after of the 2026-11-09 source lines, on the live page: the parity mock (ws_bis.py) is a 1440
// mock and squeezes its rows to one word per line at 390, so the phone captures come from production with the
// changed lines swapped in place (the words web/src/lib/bis/source-line.ts prints for the new loot.json fields).
//   node design/mocks/shoot_bis_labels_phone.mjs <out dir>
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const OUT = (process.argv[2] ?? "design/reviews/2026-11-09-feedback-ai-look/labels") + "/";
mkdirSync(OUT, { recursive: true });
const PAGE = "https://foreversixty.gg/bis/hunter/marksmanship?band=60";
const AFTER = {
  "Vendor: Pix Xizzix": "Vendor: Pix Xizzix, Booty Bay · 100 Tarnished Undermine Reals",
  "Theramore Expeditionary Force reputation · Exalted": "Theramore Expeditionary Force reputation · Exalted · The Barrens",
  "Vendor: Creeg Bothunk": "Darkspear Raiders reputation · Exalted · The Barrens",
};
const swap = (after) => {
  const walker = document.createTreeWalker(document.body, NodeFilter.SHOW_TEXT);
  let count = 0;
  for (let node = walker.nextNode(); node; node = walker.nextNode()) {
    const text = node.textContent.trim();
    if (after[text] !== undefined) {
      node.textContent = after[text];
      count += 1;
    }
  }
  return count;
};
const browser = await chromium.launch();
const page = await browser.newPage({ viewport: { width: 390, height: 900 }, deviceScaleFactor: 2 });
for (const faction of ["alliance", "horde"]) {
  await page.goto(`${PAGE}&faction=${faction}`, { waitUntil: "networkidle" });
  await page.waitForTimeout(1500);
  const gear = page.locator('[aria-label="Gear"]').first();
  await gear.screenshot({ path: `${OUT}${faction}-before-390-gear.png` });
  const swapped = await page.evaluate(swap, AFTER);
  await gear.screenshot({ path: `${OUT}${faction}-after-390-gear.png` });
  console.log(`${faction}: ${swapped} lines swapped`);
}
await browser.close();
