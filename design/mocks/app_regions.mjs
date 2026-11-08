// The native parity reference (design/parity.md, "Native"): for every app mock, each region's box and inventory, in
// points / dp from the top of the content (the bottom of the navigation bar), at the device width.
//   node design/mocks/app_regions.mjs           -> design/mocks/app/regions/<platform>-<screen>.json
//   node design/mocks/app_regions.mjs --check   -> exit 1 naming each region file that differs from a fresh render
// Run `python3 design/mocks/ws_app.py` first (the web-derived mocks are generated, not committed).
// The committed files are rendered on the Linux runner the web parity job checks them on
// (the web parity job uploads them as app-regions when its check fails); a Mac lays some
// mocks out differently.
// PLAYWRIGHT_FROM: another package.json whose node_modules carries @playwright/test.
import { createRequire } from "node:module";
import { mkdirSync, readdirSync, readFileSync, writeFileSync, existsSync } from "node:fs";
import { routeFonts } from "./app_fonts.mjs";

const require = createRequire(process.env.PLAYWRIGHT_FROM ?? new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const SRC = new URL("./app/", import.meta.url).pathname;
const OUT = new URL("./app/regions/", import.meta.url).pathname;
const DEVICE = { ios: { width: 393, height: 852 }, android: { width: 412, height: 915 } };
const check = process.argv.includes("--check");
mkdirSync(OUT, { recursive: true });

// What the inventory counts per region. The native region modifier reports the same kinds from the semantics tree:
// row (a Row, FeedRow, ItemRow or paperdoll cell), button, segment (a radio in a Segmented), field, image (icon or crest).
const KINDS = {
  row: ".row, .frow, .cell, .irow, .rr, .sr",
  button: ".btn, button",
  segment: "[role=radio]",
  field: ".field, input, .sel",
  image: "img",
};

async function measure(page) {
  return page.evaluate((kinds) => {
    const main = document.querySelector("main");
    const top = main.getBoundingClientRect().top;
    const out = {};
    for (const el of document.querySelectorAll("main [data-region]")) {
      const b = el.getBoundingClientRect();
      if (b.width === 0 || b.height === 0) continue;
      const name = el.getAttribute("data-region");
      const inv = {};
      for (const [k, sel] of Object.entries(kinds)) {
        const n = [...el.querySelectorAll(sel)].filter((x) => x.getBoundingClientRect().height > 0).length;
        if (n) inv[k] = n;
      }
      out[name] = { box: [b.x, b.y - top, b.width, b.height].map(Math.round), inventory: inv };
    }
    return out;
  }, KINDS);
}

const files = readdirSync(SRC).filter((f) => f.endsWith(".html"));
const browser = await chromium.launch();
const stale = [];
for (const file of files) {
  const platform = file.startsWith("android-") ? "android" : "ios";
  const context = await browser.newContext({ viewport: DEVICE[platform], reducedMotion: "reduce" });
  await routeFonts(context);
  const page = await context.newPage();
  await page.goto(`file://${SRC}${file}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready);
  const regions = await measure(page);
  const json = JSON.stringify({ device: DEVICE[platform], regions }, null, 1) + "\n";
  const target = `${OUT}${file.replace(/\.html$/, ".json")}`;
  if (check) {
    if (!existsSync(target) || readFileSync(target, "utf8") !== json) stale.push(target);
  } else {
    writeFileSync(target, json);
  }
  await page.context().close();
}
await browser.close();
if (check && stale.length) {
  console.error(`stale native region files (run python3 design/mocks/ws_app.py && node design/mocks/app_regions.mjs):\n${stale.join("\n")}`);
  process.exit(1);
}
console.log(`${files.length} region files ${check ? "checked" : "written"} in ${OUT}`);
