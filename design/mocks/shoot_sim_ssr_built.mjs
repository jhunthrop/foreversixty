// design/mocks/shoot_sim_ssr_built.mjs
// Acceptance captures of /sim for the server-rendered landing (SEO lane 2026-11-15): at 1440 and
// 390, the server HTML before the island runs (its script blocked), after hydration, and with
// JavaScript off. It also logs the layout shift the hydration causes and the server word count.
//
//   (cd web && npm run build && npm run preview -- --port 4360)
//   SIM_BASE=http://localhost:4360 SIM_OUT=design/reviews/2026-11-15-sim-ssr \
//      node design/mocks/shoot_sim_ssr_built.mjs [prefix]
//        -> <out>/<prefix>-<width>-{server,hydrated,nojs}.png and <prefix>-cls.json
import { createRequire } from "node:module";
import { mkdirSync, writeFileSync } from "node:fs";
import { resolve } from "node:path";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const BASE = process.env.SIM_BASE ?? "http://localhost:4360";
const OUT = resolve(process.env.SIM_OUT ?? "design/reviews/2026-11-15-sim-ssr");
const PATH = process.env.SIM_PATH ?? "/sim";
const PREFIX = process.argv[2] ?? "after";
const WIDTHS = [1440, 390];
const FONT_SETTLE_MS = 1500;
const HYDRATE_SETTLE_MS = 2500;

const CLS_PROBE = `
  window.__cls = 0;
  new PerformanceObserver((list) => {
    for (const entry of list.getEntries()) if (!entry.hadRecentInput) window.__cls += entry.value;
  }).observe({ type: 'layout-shift', buffered: true });
`;

async function shoot(browser, width, state) {
  const context = await browser.newContext({
    viewport: { width, height: 900 },
    deviceScaleFactor: width < 768 ? 3 : 2,
    javaScriptEnabled: state !== "nojs",
  });
  const page = await context.newPage();
  // Signed out: the landing is the state a crawler and a cold visitor get.
  await page.route("**/v1/**", (route) => route.fulfill({ status: 401, body: "{}" }));
  if (state === "server") await page.route("**/sim-island.js", (route) => route.abort());
  if (state === "hydrated") await page.addInitScript(CLS_PROBE);
  await page.goto(`${BASE}${PATH}`);
  await page.waitForTimeout(state === "hydrated" ? HYDRATE_SETTLE_MS : FONT_SETTLE_MS);
  await page.screenshot({ path: `${OUT}/${PREFIX}-${width}-${state}.png`, fullPage: true });
  const height = await page.evaluate(() => document.documentElement.scrollHeight);
  const cls = state === "hydrated" ? await page.evaluate(() => window.__cls) : null;
  await context.close();
  return { width, state, height, cls };
}

mkdirSync(OUT, { recursive: true });
const browser = await chromium.launch();
const results = [];
for (const width of WIDTHS) {
  for (const state of ["server", "hydrated", "nojs"]) results.push(await shoot(browser, width, state));
}
await browser.close();
writeFileSync(`${OUT}/${PREFIX}-cls.json`, `${JSON.stringify(results, null, 2)}\n`);
console.log(JSON.stringify(results));
