// Guide talent trees in their lanes (owner 2026-11-09: the Leveling lane's three trees ran out of the panel at 2560) and
// the guide top under the spec tabs (the struck "<Spec> is the" lede). For each guide and width: the Leveling and
// Talents sections, whether any tree cell sits outside its lane (printed), and the page top down to At a glance.
//   (web) npx astro preview --port 4561   then   node design/mocks/shoot_guide_trees.mjs http://localhost:4561 <out dir> <tag>
import { createRequire } from "node:module";
import { mkdirSync, writeFileSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const BASE = process.argv[2] ?? "http://localhost:4561";
const OUT = (process.argv[3] ?? "design/reviews/2026-11-09-f-content/trees") + "/";
const TAG = process.argv[4] ?? "after";
const GUIDES = ["warrior/protection", "hunter/beast-mastery", "priest/holy", "druid/restoration"];
const WIDTHS = [390, 1024, 1440, 2000, 2560];
mkdirSync(OUT, { recursive: true });

const overflow = (id) => {
  const sec = document.querySelector(`#${id}`);
  const lane = sec?.querySelector("[data-reading-lane]") ?? sec?.querySelector(".ui-rd-lane");
  if (!lane) return null;
  const box = lane.getBoundingClientRect();
  const panel = sec.closest(".ui-reading-panel")?.getBoundingClientRect() ?? box;
  let worst = 0;
  for (const el of lane.querySelectorAll("*")) {
    const r = el.getBoundingClientRect();
    if (r.width === 0) continue;
    worst = Math.max(worst, r.right - Math.min(box.right, panel.right));
  }
  return { laneW: Math.round(box.width), past: Math.round(worst) };
};

const browser = await chromium.launch();
const report = {};
for (const width of WIDTHS) {
  const page = await browser.newPage({ viewport: { width, height: 1000 }, deviceScaleFactor: 1 });
  for (const guide of GUIDES) {
    await page.goto(`${BASE}/guides/${guide}`, { waitUntil: "networkidle" });
    const tall = await page.evaluate(() => document.documentElement.scrollHeight);
    await page.setViewportSize({ width, height: Math.min(tall, 16000) });
    await page.evaluate(() => window.scrollTo(0, 0));
    await page.waitForTimeout(400);
    const slug = guide.replace("/", "-");
    report[`${guide}@${width}`] = {
      leveling: await page.evaluate(overflow, "leveling"),
      talents: await page.evaluate(overflow, "talents-and-builds"),
    };
    for (const id of ["leveling", "talents-and-builds"]) {
      await page.locator(`#${id}`).screenshot({ path: `${OUT}${TAG}-${slug}-${id}-${width}.png` });
    }
    // The top: the band and the spec tabs down to the At a glance plate's head.
    const glance = await page.locator('[data-testid="guide-glance"]').boundingBox();
    await page.screenshot({ path: `${OUT}${TAG}-${slug}-top-${width}.png`, clip: { x: 0, y: 0, width, height: Math.max(400, Math.round((glance?.y ?? 400) + 120)) } });
    await page.setViewportSize({ width, height: 1000 });
  }
  await page.close();
}
writeFileSync(`${OUT}${TAG}-overflow.json`, JSON.stringify(report, null, 1));
console.log(JSON.stringify(report));
await browser.close();
