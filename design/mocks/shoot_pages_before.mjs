// Before captures of the live changelog, content (legal and about) pages, 404 and build-unavailable page
// (design/specs/2026-11-08-pages.md section 1): production, signed out, and signed in with the account read
// (api.foreversixty.gg /v1/me) answered from the home fixture me-34. Writes <out>/<name>-<w>.png, -fold.png and
// measure.json (every number section 1 quotes).
//   node design/mocks/shoot_pages_before.mjs <out dir>
import { createRequire } from "node:module";
import { mkdirSync, readFileSync, writeFileSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const W = new URL("../../", import.meta.url).pathname;
const OUT = (process.argv[2] ?? "design/mocks/renders/pages-before") + "/";
mkdirSync(OUT, { recursive: true });
const SITE = "https://foreversixty.gg";
const me34 = JSON.parse(readFileSync(`${W}/web/tests/e2e/fixtures/home/me-34.json`, "utf8"));
const ok = (data) => JSON.stringify({ ok: true, data, error: null, request_id: "r" });
const bad = (m) => JSON.stringify({ ok: false, data: null, error: { message: m }, request_id: "r" });
const cors = { "access-control-allow-origin": SITE, "access-control-allow-credentials": "true", "access-control-allow-headers": "content-type,x-csrf-token", "access-control-allow-methods": "GET,POST,OPTIONS" };
async function stub(page, me) {
  await page.route("https://api.foreversixty.gg/**", async (route) => {
    const req = route.request(); const u = new URL(req.url());
    if (req.method() === "OPTIONS") return route.fulfill({ status: 204, headers: cors });
    const f = (body, status = 200) => route.fulfill({ status, headers: { ...cors, "content-type": "application/json" }, body });
    if (u.pathname === "/v1/me") return me ? f(ok(me)) : f(bad("Sign in"), 401);
    return me ? f(ok([])) : f(bad("Sign in"), 401);
  });
}
const MEASURE = () => {
  const box = (el) => { if (!el) return null; const r = el.getBoundingClientRect(); return { x: Math.round(r.x), y: Math.round(r.y + scrollY), w: Math.round(r.width), h: Math.round(r.height) }; };
  const q = (s) => document.querySelector(s);
  const qa = (s) => [...document.querySelectorAll(s)];
  const vis = (e) => e.offsetParent !== null;
  const h2 = qa("main h2").filter(vis).map((e) => ({ t: e.textContent.trim().slice(0, 40), ...box(e) }));
  const boxes = qa("main *").filter((e) => { const cs = getComputedStyle(e); return vis(e) && parseFloat(cs.borderTopWidth) >= 1 && cs.borderTopStyle !== "none" && parseFloat(cs.borderLeftWidth) >= 1 && cs.borderRadius !== "0px" && e.getBoundingClientRect().height > 60; });
  const fonts = {};
  qa("main *").forEach((e) => { if (!vis(e) || ![...e.childNodes].some((n) => n.nodeType === 3 && n.textContent.trim())) return; const cs = getComputedStyle(e); const k = cs.fontFamily.split(",")[0].replace(/"/g, "") + " " + cs.fontWeight + " " + cs.fontSize + (cs.textTransform === "uppercase" ? " caps" : ""); fonts[k] = (fonts[k] || 0) + 1; });
  const entries = qa("main article").filter(vis).map(box);
  const text = qa("main p, main li").filter(vis).map(box);
  const left = Math.min(...text.map((t) => t.x)), right = Math.max(...text.map((t) => t.x + t.w));
  return {
    height: document.documentElement.scrollHeight, title: box(q("main h1")), titleFont: q("main h1") && getComputedStyle(q("main h1")).fontFamily.split(",")[0],
    titleText: q("main h1")?.textContent.trim(), eyebrow: q("main h1")?.previousElementSibling?.textContent.trim() ?? null,
    main: box(q("main")), textLeft: left, textRight: right, h2, framed: boxes.length, framedBoxes: boxes.slice(0, 6).map(box),
    entries: entries.length, entriesInFold: entries.filter((e) => e.y + e.h <= innerHeight).length, firstEntry: entries[0] ?? null,
    fonts, links: qa("main a").filter(vis).length, imgs: qa("main img").length, paragraphs: qa("main p").filter(vis).length,
  };
};
const b = await chromium.launch();
const measures = {};
async function shot(name, path, me, widths) {
  for (const w of widths) {
    const ctx = await b.newContext({ viewport: { width: w, height: w < 500 ? 844 : 1000 } });
    if (me) await ctx.addCookies([{ name: "fs_csrf", value: "t", domain: "foreversixty.gg", path: "/" }]);
    const page = await ctx.newPage(); await stub(page, me);
    const res = await page.goto(SITE + path, { waitUntil: "networkidle" }); await page.waitForTimeout(1500);
    measures[`${name}-${w}`] = { status: res?.status(), ...(await page.evaluate(MEASURE)) };
    await page.screenshot({ path: `${OUT}${name}-${w}.png`, fullPage: true });
    await page.screenshot({ path: `${OUT}${name}-${w}-fold.png` });
    await ctx.close();
  }
}
const ALL = [390, 1440, 1920, 2000], TWO = [390, 1440];
await shot("changelog", "/changelog", null, ALL);
await shot("changelog-in", "/changelog", me34, TWO);
await shot("privacy", "/privacy", null, ALL);
await shot("terms", "/terms", null, TWO);
await shot("refunds", "/refunds", null, TWO);
await shot("about", "/about", null, ALL);
await shot("sources", "/sources", null, TWO);
await shot("roadmap", "/roadmap", null, TWO);
await shot("404", "/this-page-does-not-exist", null, ALL);
await shot("404-in", "/this-page-does-not-exist", me34, TWO);
await shot("b-unavailable", "/b-unavailable", null, TWO);
await b.close();
writeFileSync(OUT + "measure.json", JSON.stringify(measures, null, 1));
console.log(OUT);
