// Home build vs mock at any width: node design/mocks/ws_home_shoot.mjs <base url> <out dir> <state> <widths>
// Opens the mock (design/mocks/web/web-home-<state>.html) and the built page with the same fixtures the parity
// spec routes (web/tests/e2e/fixtures/home/), and writes <out>/home-<state>-<w>.png: mock | build, same scale.
import { createRequire } from "node:module";
import { mkdirSync, readFileSync } from "node:fs";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const [base, out, state, widths] = process.argv.slice(2);
const FIX = new URL("../../web/tests/e2e/fixtures/home/", import.meta.url);
const STATES = JSON.parse(readFileSync(new URL("./web/home-parity.json", import.meta.url), "utf8")).states;
const s = STATES.find((x) => x.id === state);
const fx = (n) => JSON.parse(readFileSync(new URL(n, FIX), "utf8"));
const env = (data) => ({ json: { ok: true, data, error: null } });
mkdirSync(out, { recursive: true });
const b = await chromium.launch();
for (const w of widths.split(",").map(Number)) {
  const h = w < 500 ? 844 : 1000;
  const mock = await (await b.newContext({ viewport: { width: w, height: h } })).newPage();
  await mock.goto(new URL(`./web/${s.mock}`, import.meta.url).href, { waitUntil: "networkidle" });
  const ctx = await b.newContext({ viewport: { width: w, height: h } });
  const page = await ctx.newPage();
  await page.route("**/v1/**", (r) => r.fulfill({ status: 401, json: { data: null } }));
  if (s.me) {
    await page.clock.setFixedTime(new Date("2026-11-06T18:00:00Z"));
    await ctx.addCookies([{ name: "fs_csrf", value: "t", domain: new URL(base).hostname, path: "/" }]);
    await page.route("**/v1/me", (r) => (s.me === "pending" ? new Promise(() => {}) : r.fulfill(env(fx(s.me)))));
    await page.route("**/v1/sims?**", (r) => r.fulfill(env(s.sims ? fx(s.sims) : { rows: [] })));
    await page.route("**/v1/reports?mine=1**", (r) => r.fulfill(env(s.reports ? fx(s.reports) : { rows: [] })));
    if (s.guild) await page.route("**/v1/guilds/*/home", (r) => r.fulfill(env(fx(s.guild))));
    for (const [spec, cls] of [["hunter-beast-mastery", "hunter"], ["warrior-fury", "warrior"]]) {
      await page.route(`**/data/*/bis/${spec}.json`, (r) => (s.failBis ? r.fulfill({ status: 500, body: "" }) : r.fulfill({ json: fx(`bis-${spec}.json`) })));
      await page.route(`**/data/*/items/${cls}.json`, (r) => r.fulfill({ json: fx(`items-${cls}.json`) }));
    }
  }
  await page.goto(base + "/", { waitUntil: "networkidle" });
  await page.waitForTimeout(800);
  const [m, p] = [await mock.screenshot({ fullPage: true }), await page.screenshot({ fullPage: true })];
  const board = await b.newPage({ viewport: { width: 2 * w + 36, height: 800 } });
  const fig = (png, label) => `<figure style="margin:0"><figcaption style="font:700 16px/28px sans-serif;color:#e5b955">${label}</figcaption><img style="display:block;width:${w}px" src="data:image/png;base64,${png.toString("base64")}"></figure>`;
  await board.setContent(`<body style="margin:0;padding:12px;background:#1c1c1e;display:flex;gap:12px;align-items:flex-start">${fig(m, `MOCK home-${state} ${w}`)}${fig(p, `BUILD home-${state} ${w}`)}</body>`);
  await board.screenshot({ path: `${out}/home-${state}-${w}.png`, fullPage: true });
  await board.close(); await mock.close(); await ctx.close();
}
await b.close();
