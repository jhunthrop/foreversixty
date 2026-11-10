// design/mocks/shoot_report_tail.mjs
// Captures of the BUILT report tail (batch H lane h-report-tail): the Events, Queries, Dispels,
// Interrupts and Timeline tabs of /reports/sample, the whole night (fight=all), and /ratings, at
// 1440, 2000 and 390, full page. The API is stubbed signed out, so nothing here is a real account.
//
//   (cd web && npm run build && npx astro preview --port 4431)
//   node design/mocks/shoot_report_tail.mjs [surface ...] -> design/reviews/2026-11-09-h-report-tail/<surface>-<width>.png
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const BASE = process.env.TAIL_BASE ?? "http://localhost:4431";
const OUT = new URL("../reviews/2026-11-09-h-report-tail/", import.meta.url).pathname;
const RUNTIME = new URL("../../web/build/duckdb-runtime/", import.meta.url).pathname;
const WIDTHS = (process.env.TAIL_WIDTHS ?? "1440,2000,390").split(",").map(Number);
const SETTLE_MS = 2500;
const QUERY_SETTLE_MS = 9000;

const SURFACES = {
  events: { path: "/reports/sample?tab=events" },
  queries: { path: "/reports/sample?tab=queries", settle: QUERY_SETTLE_MS, run: true },
  dispels: { path: "/reports/sample?tab=dispels" },
  interrupts: { path: "/reports/sample?tab=interrupts" },
  timeline: { path: "/reports/sample?tab=timeline" },
  night: { path: "/reports/sample?fight=all" },
  "night-damage": { path: "/reports/sample?fight=all&tab=damage-done" },
  ratings: { path: "/ratings" },
};

const unauthorized = { status: 401, contentType: "application/json", body: JSON.stringify({ ok: false, data: null, error: "signed out" }) };

async function shoot(browser, name, surface, width) {
  const context = await browser.newContext({ viewport: { width, height: 900 } });
  await context.route("**/v1/me", (route) => route.fulfill(unauthorized));
  // The Worker serves the DuckDB engine from R2; `astro preview` has no Worker, so the staged copy
  // (npm run sync:duckdb -> web/build/duckdb-runtime/) stands in.
  await context.route("**/duckdb-runtime/**", (route) =>
    route.fulfill({ path: `${RUNTIME}${new URL(route.request().url()).pathname.replace("/duckdb-runtime/", "")}`, contentType: "application/wasm" }),
  );
  const page = await context.newPage();
  await page.goto(`${BASE}${surface.path}`);
  await page.waitForTimeout(SETTLE_MS);
  if (surface.run) {
    const run = page.getByRole("button", { name: /^Run/ }).first();
    if (await run.isVisible().catch(() => false)) await run.click();
  }
  await page.waitForTimeout(surface.settle ?? SETTLE_MS);
  const out = `${OUT}${name}-${width}.png`;
  await page.screenshot({ path: out, fullPage: true });
  await context.close();
  console.log(out);
}

mkdirSync(OUT, { recursive: true });
const names = process.argv.slice(2).length ? process.argv.slice(2) : Object.keys(SURFACES);
const browser = await chromium.launch();
for (const name of names) for (const width of WIDTHS) await shoot(browser, name, SURFACES[name], width);
await browser.close();
