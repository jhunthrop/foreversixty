// Before captures of the live guide pages (design/specs/2026-11-08-guides.md section 1): production, signed out, and
// signed in with the account read (api.foreversixty.gg /v1/me) answered from the home fixtures, so the signed-in
// states exist without an account. Writes <out>/<name>-<w>.png, -fold.png and measure.json (every number section 1 quotes).
//   node design/mocks/shoot_guides_before.mjs <out dir>
import { createRequire } from "node:module";
import { mkdirSync, readFileSync, writeFileSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const W = new URL("../../", import.meta.url).pathname;
const OUT = (process.argv[2] ?? "design/mocks/renders/guides-before") + "/";
mkdirSync(OUT, { recursive: true });
const SITE = "https://foreversixty.gg";
const j = (p) => JSON.parse(readFileSync(`${W}/${p}`, "utf8"));
const me34 = j("web/tests/e2e/fixtures/home/me-34.json");
const me60 = j("web/tests/e2e/fixtures/home/me-60.json");
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
// Everything section 1 measures, in page coordinates.
const MEASURE = () => {
  const box = (el) => { if (!el) return null; const r = el.getBoundingClientRect(); return { x: Math.round(r.x), y: Math.round(r.y + scrollY), w: Math.round(r.width), h: Math.round(r.height) }; };
  const q = (s) => document.querySelector(s);
  const qa = (s) => [...document.querySelectorAll(s)];
  const h2 = qa("main h2, article h2").filter((e) => e.offsetParent !== null).map((e) => ({ t: e.textContent.trim().slice(0, 40), ...box(e) }));
  const boxes = qa("main *").filter((e) => { const cs = getComputedStyle(e); return e.offsetParent !== null && parseFloat(cs.borderTopWidth) >= 1 && cs.borderTopStyle !== "none" && parseFloat(cs.borderLeftWidth) >= 1 && cs.borderRadius !== "0px" && e.getBoundingClientRect().height > 60; });
  const fonts = {};
  qa("main *").forEach((e) => { if (e.offsetParent === null || !e.childNodes.length || ![...e.childNodes].some((n) => n.nodeType === 3 && n.textContent.trim())) return; const cs = getComputedStyle(e); const k = cs.fontFamily.split(",")[0].replace(/"/g, "") + " " + cs.fontWeight + (cs.textTransform === "uppercase" ? " caps" : ""); fonts[k] = (fonts[k] || 0) + 1; });
  const small = qa("main p, main li, main span, main a").filter((e) => e.offsetParent !== null && [...e.childNodes].some((n) => n.nodeType === 3 && n.textContent.trim().length > 20) && parseFloat(getComputedStyle(e).fontSize) < 13).length;
  return {
    height: document.documentElement.scrollHeight, title: box(q("h1")), titleFont: q("h1") && getComputedStyle(q("h1")).fontFamily.split(",")[0],
    rail: box(q('[data-testid="guide-action-rail"]')), toc: box(q("nav[aria-label], .toc, [data-testid*=toc]")), article: box(q("article, .prose")),
    cards: box(q('[data-testid="guide-class-spec-cards"]')), picker: box(q("main ul, main .grid")), h2, framed: boxes.length, imgs: qa("main img").length,
    sectionsInFold: h2.filter((h) => h.y < innerHeight).length, fonts, smallTextNodes: small, links: qa("main a").length,
  };
};
const b = await chromium.launch();
const measures = {};
async function shot(name, path, me, widths) {
  for (const w of widths) {
    const ctx = await b.newContext({ viewport: { width: w, height: w < 500 ? 844 : 1000 } });
    if (me) await ctx.addCookies([{ name: "fs_csrf", value: "t", domain: "foreversixty.gg", path: "/" }]);
    const page = await ctx.newPage(); await stub(page, me);
    await page.goto(SITE + path, { waitUntil: "networkidle" }); await page.waitForTimeout(2200);
    await page.evaluate(async () => { for (let y = 0; y < document.body.scrollHeight; y += 700) { scrollTo(0, y); await new Promise((r) => setTimeout(r, 60)); } scrollTo(0, 0); });
    await page.waitForTimeout(600);
    measures[`${name}-${w}`] = await page.evaluate(MEASURE);
    await page.screenshot({ path: `${OUT}${name}-${w}.png`, fullPage: true });
    await page.screenshot({ path: `${OUT}${name}-${w}-fold.png` });
    await ctx.close();
  }
}
const ALL = [390, 1440, 1920, 2000], TWO = [390, 1440];
await shot("guides", "/guides", null, ALL);
await shot("guides-in", "/guides", me34, TWO);
await shot("guides-hunter", "/guides/hunter", null, ALL);
await shot("guides-hunter-in", "/guides/hunter", me34, TWO);
await shot("guides-priest", "/guides/priest", null, TWO);
await shot("guide-survival", "/guides/hunter/survival", null, ALL);
await shot("guide-survival-in", "/guides/hunter/survival", me60, TWO);
await shot("guide-frost", "/guides/mage/frost", null, [390, 1440, 2000]);
await shot("guide-holy", "/guides/priest/holy", null, [390, 1440, 2000]);
await shot("guide-protection", "/guides/warrior/protection", null, [1440]);
await b.close();
writeFileSync(OUT + "measure.json", JSON.stringify(measures, null, 1));
console.log(OUT);
