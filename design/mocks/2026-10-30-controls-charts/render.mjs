// Screenshots every board in renders/ (or the names given), as the 2026-10-29 premium-entry renderer does:
// "phone" in the name renders at 390, "2000" at 2000, everything else at 1440; full page, fonts loaded.
import { readdirSync } from "node:fs";
import { join, dirname } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";
import { createRequire } from "node:module";

const here = dirname(fileURLToPath(import.meta.url));
const require = createRequire(join(here, "../../../web/package.json"));
const { chromium } = require("playwright");
const renders = join(here, "renders");
const names = process.argv.slice(2).length ? process.argv.slice(2) : readdirSync(renders).filter((f) => f.endsWith(".html")).map((f) => f.slice(0, -5));
const browser = await chromium.launch();
for (const name of names) {
  const width = name.includes("phone") ? 390 : name.includes("2000") ? 2000 : 1440;
  const page = await browser.newPage({ viewport: { width, height: 900 } });
  await page.goto(pathToFileURL(join(renders, `${name}.html`)).href, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(400);
  await page.screenshot({ path: join(renders, `${name}.png`), fullPage: true });
  await page.close();
  console.log(name, width);
}
await browser.close();
