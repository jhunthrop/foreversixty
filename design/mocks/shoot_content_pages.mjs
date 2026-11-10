// Captures the content pages lane P15 moves to round 5 (about, the legal pages, sources, the changelog, the 404 and
// the build-unavailable page) from a running preview, full page at 390, 1024, 1440 and 2000, and the round-5 mocks
// (design/mocks/directions/html/b4-*.html) at the same widths for the side-by-side.
//   (web) npx astro preview --port 4561   then   node design/mocks/shoot_content_pages.mjs http://localhost:4561 <out dir> [page ...]
import { createRequire } from "node:module";
import { existsSync, mkdirSync } from "node:fs";
import { fileURLToPath } from "node:url";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const BASE = process.argv[2] ?? "http://localhost:4561";
const OUT = (process.argv[3] ?? "design/reviews/2026-11-09-f-content") + "/";
const ONLY = process.argv.slice(4);
const MOCKS = fileURLToPath(new URL("directions/html/", import.meta.url));
const PAGES = [
  ["about", "/about", "b4-about"],
  ["privacy", "/privacy", "b4-privacy"],
  ["terms", "/terms", "b4-terms"],
  ["refunds", "/refunds", "b4-refunds"],
  ["sources", "/sources", "b4-sources"],
  ["changelog", "/changelog", "b4-changelog"],
  ["404", "/guides/hunter/beast-mastry", "b4-404"],
  ["unavailable", "/b-unavailable", "b4-unavailable"],
];
const WIDTHS = [390, 1024, 1440, 2000];
mkdirSync(OUT, { recursive: true });

const browser = await chromium.launch();
for (const width of WIDTHS) {
  const page = await browser.newPage({ viewport: { width, height: width < 500 ? 844 : 900 } });
  for (const [name, path, mock] of PAGES.filter(([n]) => !ONLY.length || ONLY.includes(n))) {
    await page.goto(`${BASE}${path}`, { waitUntil: "networkidle" });
    await page.evaluate(() => document.fonts.ready);
    await page.waitForTimeout(250);
    await page.screenshot({ path: `${OUT}build-${name}-${width}.png`, fullPage: true });
    const file = `${MOCKS}${mock}.html`;
    if (!existsSync(file)) continue;
    await page.goto(`file://${file}`, { waitUntil: "load" });
    await page.evaluate(() => document.fonts.ready);
    await page.waitForTimeout(250);
    await page.screenshot({ path: `${OUT}mock-${name}-${width}.png`, fullPage: true });
  }
  await page.close();
}
await browser.close();
