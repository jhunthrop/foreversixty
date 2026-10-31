// Screenshots every board in renders/ (or the names given) at its own width, fonts loaded, at 1.5x.
// Run from the repo root after gen_wave3.py: node design/mocks/2026-10-31-mobile-wave3/render.mjs
import { readdirSync, existsSync } from "node:fs";
import { join, dirname } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";
import { createRequire } from "node:module";

const here = dirname(fileURLToPath(import.meta.url));
// A worktree has no web/node_modules; the main checkout (FOREVER_WEB) holds playwright.
const webRoot = existsSync(join(here, "../../../web/node_modules")) ? join(here, "../../../web") : (process.env.FOREVER_WEB ?? "/Users/jh/code/forever/web");
const require = createRequire(join(webRoot, "package.json"));
const { chromium } = require("playwright");
const renders = join(here, "renders");
const names = process.argv.slice(2).length ? process.argv.slice(2) : readdirSync(renders).filter((f) => f.endsWith(".html")).map((f) => f.slice(0, -5));
const browser = await chromium.launch();
for (const name of names) {
  const page = await browser.newPage({ viewport: { width: 4000, height: 1000 }, deviceScaleFactor: 1.5 });
  await page.goto(pathToFileURL(join(renders, `${name}.html`)).href, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready);
  const box = await page.evaluate(() => { const b = document.querySelector(".board").getBoundingClientRect(); return { w: Math.ceil(b.width), h: Math.ceil(b.height) }; });
  await page.setViewportSize({ width: box.w, height: box.h });
  await page.waitForTimeout(300);
  await page.screenshot({ path: join(renders, `${name}.png`), fullPage: true });
  await page.close();
  console.log(name, box.w, box.h);
}
await browser.close();
