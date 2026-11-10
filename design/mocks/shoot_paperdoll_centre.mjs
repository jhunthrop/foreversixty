// The paperdoll centre's fill and centring (owner 2026-11-09): on /bis/warrior/fury and the Fury guide's Best in slot doll,
// a slot hovered, the frame at 390 (BiS only: the list form has no centre), 1024, 1440 and 2000, plus a crop of each of the
// centre's four edges (the teal against the knot border's inner line) and the measured insets of the tooltip in the centre.
//   (web) npx astro preview --port 4561   then   node design/mocks/shoot_paperdoll_centre.mjs http://localhost:4561 <out dir> <tag>
import { createRequire } from "node:module";
import { mkdirSync, readFileSync, writeFileSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const BASE = process.argv[2] ?? "http://localhost:4561";
const OUT = (process.argv[3] ?? "design/reviews/2026-11-09-f-content/paperdoll") + "/";
const TAG = process.argv[4] ?? "after";
const EDGE_PX = 70;
// "before": the frame art and centre box main shipped before the fix (pass their folder as the 5th argument: old-1x.webp,
// old-2x.webp from `git show 52748165d:web/public/art/paperdoll-frame@{1x,2x}.webp`).
const OLD_ART = process.argv[5];
const OLD_CENTRE = "left:22.604%;top:21.107%;width:54.679%;height:65.347%";
const SURFACES = [
  ["bis", "/bis/warrior/fury", '[data-testid="bis-doll"]'],
  ["guide", "/guides/warrior/fury", '[data-testid="guide-bis-frame"]'],
];
mkdirSync(OUT, { recursive: true });

const browser = await chromium.launch();
const report = {};
for (const width of [390, 1024, 1440, 2000]) {
  const page = await browser.newPage({ viewport: { width, height: 1300 }, deviceScaleFactor: 1 });
  if (OLD_ART) {
    await page.route(/paperdoll-frame@(1x|2x)\.webp/, (route) => {
      const scale = route.request().url().includes("@2x") ? "2x" : "1x";
      return route.fulfill({ contentType: "image/webp", body: readFileSync(`${OLD_ART}/old-${scale}.webp`) });
    });
  }
  for (const [name, path, frameSel] of SURFACES) {
    await page.goto(`${BASE}${path}`, { waitUntil: "networkidle" });
    const tall = await page.evaluate(() => document.documentElement.scrollHeight);
    await page.setViewportSize({ width, height: Math.min(tall, 16000) });
    const frame = page.locator(frameSel).first();
    const slot = frame.locator('[data-testid="paperdoll-frame-slot-head"]');
    if (!(await slot.isVisible())) {
      if (name === "bis") await frame.screenshot({ path: `${OUT}${TAG}-${name}-${width}.png` });
      await page.setViewportSize({ width, height: 1300 });
      continue;
    }
    await slot.hover();
    await page.waitForTimeout(300);
    if (OLD_ART) {
      await frame.locator('[data-region="paperdoll-centre"]').evaluate((el, style) => el.setAttribute("style", style), OLD_CENTRE);
      await page.waitForTimeout(100);
    }
    await frame.screenshot({ path: `${OUT}${TAG}-${name}-${width}.png` });
    const box = await frame.locator('[data-region="paperdoll-centre"]').boundingBox();
    const tip = await frame.locator('[data-region="paperdoll-centre"] [role="group"]').first().boundingBox();
    const pad = EDGE_PX;
    const crops = {
      left: { x: box.x - pad, y: box.y + box.height / 3, width: pad * 2, height: box.height / 3 },
      right: { x: box.x + box.width - pad, y: box.y + box.height / 3, width: pad * 2, height: box.height / 3 },
      top: { x: box.x + box.width / 2, y: box.y - pad, width: box.width / 2 + pad, height: pad * 2 },
      bottom: { x: box.x, y: box.y + box.height - pad, width: box.width, height: pad * 2 },
    };
    for (const [edge, clip] of Object.entries(crops)) await page.screenshot({ path: `${OUT}${TAG}-${name}-${width}-${edge}.png`, clip });
    report[`${name}@${width}`] = tip && {
      centre: { x: Math.round(box.x), w: Math.round(box.width) },
      tipLeftInset: Math.round(tip.x - box.x),
      tipRightInset: Math.round(box.x + box.width - tip.x - tip.width),
    };
    await page.setViewportSize({ width, height: 1300 });
  }
  await page.close();
}
writeFileSync(`${OUT}${TAG}-insets.json`, JSON.stringify(report, null, 1));
console.log(JSON.stringify(report));
await browser.close();
