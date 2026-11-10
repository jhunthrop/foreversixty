// The guide sections' left-column audit and captures (owner 2026-11-09: no dead space beside a tall lane).
// For each guide and width it measures every ReadingSection: the prose column's content height, the lane's height
// and the left-column gap (lane minus prose, in px, two-column widths only), prints a table and writes JSON; then
// it captures the Best in slot section (with a slot hovered) and any section whose gap passes the threshold.
//   (web) npx astro preview --port 4561   then   node design/mocks/shoot_guide_sections.mjs http://localhost:4561 <out dir> [tag]
import { createRequire } from "node:module";
import { mkdirSync, writeFileSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const BASE = process.argv[2] ?? "http://localhost:4561";
const OUT = (process.argv[3] ?? "design/reviews/2026-11-09-f-content/guides") + "/";
const TAG = process.argv[4] ?? "after";
const GAP_PX = 200;
const GUIDES = ["warrior/fury", "hunter/beast-mastery", "priest/holy", "rogue/combat", "mage/frost", "shaman/enhancement"];
const CAPTURED = new Set(["warrior/fury", "hunter/beast-mastery", "priest/holy"]);
const WIDTHS = [1440, 2000, 390];
mkdirSync(OUT, { recursive: true });

const measure = () =>
  [...document.querySelectorAll(".ui-rsec")].map((sec) => {
    const lane = sec.querySelector(".ui-rd-lane");
    const prose = sec.querySelector(".ui-rd-prose");
    const box = (el) => (el ? el.getBoundingClientRect() : null);
    const kids = prose ? [...prose.querySelectorAll(":scope > *, :scope > astro-island > *")] : [];
    const bottoms = kids.map((k) => k.getBoundingClientRect()).filter((r) => r.height > 0).map((r) => r.bottom);
    const pTop = box(prose)?.top ?? 0;
    const proseH = bottoms.length ? Math.max(...bottoms) - pTop : 0;
    const laneH = box(lane)?.height ?? 0;
    const side = lane && prose && Math.abs(box(lane).top - box(prose).top) < 4;
    return { id: sec.id, lane: Math.round(laneH), prose: Math.round(proseH), gap: side ? Math.round(laneH - proseH) : 0 };
  });

const browser = await chromium.launch();
const report = {};
for (const width of WIDTHS) {
  const page = await browser.newPage({ viewport: { width, height: 1000 }, deviceScaleFactor: 1 });
  for (const guide of GUIDES) {
    await page.goto(`${BASE}/guides/${guide}`, { waitUntil: "networkidle" });
    await page.evaluate(() => document.querySelectorAll("[loading=lazy]").forEach((img) => (img.loading = "eager")));
    // "before" (shot on the build before lib/ui/reading-wrap.ts): the page as main draws it, the doll removed.
    if (TAG === "before") await page.evaluate(() => document.querySelector('[data-testid="guide-bis-doll"]')?.closest("astro-island")?.remove());
    const rows = await page.evaluate(measure);
    report[`${guide}@${width}`] = rows;
    for (const r of rows) if (width !== 390) console.log(`${TAG} ${width} ${guide.padEnd(22)} ${r.id.padEnd(28)} lane ${String(r.lane).padStart(5)} prose ${String(r.prose).padStart(5)} gap ${String(r.gap).padStart(5)}${r.gap > GAP_PX ? "  <-- GAP" : ""}`);
    if (!CAPTURED.has(guide) || width === 2000 && guide !== "warrior/fury") continue;
    const slug = guide.replace("/", "-");
    // The whole page in one tall viewport: nothing scrolls, so the sticky spec bar never lies over a section.
    const tall = await page.evaluate(() => document.documentElement.scrollHeight);
    await page.setViewportSize({ width, height: Math.min(tall, 16000) });
    await page.evaluate(() => window.scrollTo(0, 0));
    const slot = page.locator('[data-testid="guide-bis-frame"] [data-testid="paperdoll-frame-slot-head"]');
    if (await slot.isVisible()) await slot.hover();
    await page.waitForTimeout(300);
    await page.locator("#best-in-slot").screenshot({ path: `${OUT}${TAG}-${slug}-bis-${width}.png` });
    for (const r of rows) {
      const shot = width === 390 ? r.id === "talents-and-builds" : r.gap > GAP_PX || TAG === "after";
      if (shot && r.id !== "best-in-slot") await page.locator(`#${r.id}`).screenshot({ path: `${OUT}${TAG}-${slug}-${r.id}-${width}.png` });
    }
    await page.setViewportSize({ width, height: 1000 });
    // A pinned short lane (reading-wrap's data-lane-sticky) only shows while scrolling: the viewport partway down Leveling.
    if (width === 1440 && TAG === "after") {
      await page.setViewportSize({ width, height: 900 });
      await page.evaluate(() => {
        const sec = document.querySelector("#leveling");
        window.scrollTo(0, sec.getBoundingClientRect().top + window.scrollY + 380);
      });
      await page.waitForTimeout(300);
      await page.screenshot({ path: `${OUT}pinned-${slug}-leveling-1440.png` });
      await page.setViewportSize({ width, height: 1000 });
    }
  }
  await page.close();
}
writeFileSync(`${OUT}${TAG}-gaps.json`, JSON.stringify(report, null, 1));
await browser.close();
