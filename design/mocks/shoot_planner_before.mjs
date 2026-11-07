// Before captures of the live /planner (design/specs/2026-11-08-planner.md section 1): production, with the account read
// (api.foreversixty.gg /v1/me) answered from the home fixtures so the signed-in state exists without an account.
// Writes <out>/<state>-<w>.png (full page), -fold.png (first screen) and measure.json (the boxes the diagnosis quotes).
//   node design/mocks/shoot_planner_before.mjs <out dir>   then   python3 design/mocks/ws_planner.py compare <out dir>
import { createRequire } from "node:module";
import { mkdirSync, readFileSync, writeFileSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const W = new URL("../../", import.meta.url).pathname;
const OUT = (process.argv[2] ?? "design/mocks/renders/planner-before") + "/";
mkdirSync(OUT, { recursive: true });
const SITE = "https://foreversixty.gg";
const j = (p) => JSON.parse(readFileSync(`${W}/${p}`, "utf8"));
const me34 = j("web/tests/e2e/fixtures/home/me-34.json");
const ok = (data) => JSON.stringify({ ok: true, data, error: null, request_id: "r" });
const bad = (m) => JSON.stringify({ ok: false, data: null, error: { message: m }, request_id: "r" });
const cors = { "access-control-allow-origin": SITE, "access-control-allow-credentials": "true", "access-control-allow-headers": "content-type,x-csrf-token", "access-control-allow-methods": "GET,POST,OPTIONS" };
const SURVIVAL = "FS1:1.60.1.70790:hunter:night-elf:5/005005201/500230230050222151:";
async function stub(page, s) {
  await page.route("https://api.foreversixty.gg/**", async (route) => {
    const req = route.request(); const u = new URL(req.url());
    if (req.method() === "OPTIONS") return route.fulfill({ status: 204, headers: cors });
    const f = (body, status = 200) => route.fulfill({ status, headers: { ...cors, "content-type": "application/json" }, body });
    if (u.pathname === "/v1/me") return s.me ? f(ok(s.me)) : f(bad("Sign in"), 401);
    if (u.pathname.startsWith("/v1/characters")) return f(bad("nf"), 404);
    return route.continue();
  });
}
// What the diagnosis measures: the page height, the boxes of the parts a player looks for, and every bordered block.
const measure = () => {
  const box = (sel) => { const el = document.querySelector(sel); if (!el) return null; const r = el.getBoundingClientRect(); return { x: Math.round(r.x), y: Math.round(r.y + scrollY), w: Math.round(r.width), h: Math.round(r.height) }; };
  const main = document.querySelector("main");
  const framed = [...main.querySelectorAll("*")].filter((el) => { const cs = getComputedStyle(el); const r = el.getBoundingClientRect();
    return r.height > 60 && r.width > 120 && cs.borderTopWidth !== "0px" && cs.borderTopStyle !== "none" && cs.borderLeftWidth !== "0px"; }).length;
  const visible = (sel) => [...document.querySelectorAll(sel)].filter((el) => el.getBoundingClientRect().width > 0);
  return {
    height: document.documentElement.scrollHeight, scrollWidth: document.documentElement.scrollWidth,
    h1: box("h1"), head: box('[data-testid="planner-header-band"]'), facts: box('[data-testid="planner-facts"]'),
    firstTree: box('[data-testid="tree-columns"] [role="tabpanel"]:not(.hidden)'), treeColumns: box('[data-testid="tree-columns"]'),
    firstCell: box('[data-testid^="talent-"]'), tabs: box('[data-testid="tree-tablist"]'),
    band: box('[data-testid="band-compare"]'), order: box('[data-testid="order-strip"]'), share: box('[data-testid="planner-share-section"]'),
    importBox: box('[data-testid="import-box"]'), sendBtn: box('[data-testid="character-header"] button, [data-testid="character-header"] a'),
    simLink: box('[data-testid="planner-sim-link"]'), copyGame: box('[data-testid="copy-addon-code"]'), footer: box("footer"),
    framed, images: visible("main img").length, buttons: visible("main button").length,
  };
};
const b = await chromium.launch();
const results = {};
async function shot(name, path, s, widths, after) {
  for (const w of widths) {
    const ctx = await b.newContext({ viewport: { width: w, height: w < 500 ? 844 : 1000 } });
    if (s.me) await ctx.addCookies([{ name: "fs_csrf", value: "t", domain: "foreversixty.gg", path: "/" }]);
    const page = await ctx.newPage(); await stub(page, s);
    await page.goto(SITE + path, { waitUntil: "networkidle" }); await page.waitForTimeout(1800);
    if (after) { await after(page); await page.waitForTimeout(900); }
    results[`${name}-${w}`] = await page.evaluate(measure);
    await page.screenshot({ path: `${OUT}${name}-${w}.png`, fullPage: true });
    await page.screenshot({ path: `${OUT}${name}-${w}-fold.png` });
    await ctx.close();
  }
}
const ALL = [390, 1440, 1920, 2000], TWO = [390, 1440];
await shot("planner-out", "/planner", {}, ALL);
await shot("planner-guide", "/planner?code=" + encodeURIComponent(SURVIVAL), {}, ALL);
await shot("planner-in", "/planner", { me: me34 }, [390, 1440, 2000]);
await shot("planner-invalid", "/planner?code=FS9%3Anope", {}, TWO);
await shot("planner-priest", "/planner?class=priest", {}, TWO);
writeFileSync(`${OUT}measure.json`, JSON.stringify(results, null, 1) + "\n");
await b.close();
