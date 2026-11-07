// Before captures of the live /premium, /premium/checkout, /account and /b/:id
// (design/specs/2026-11-08-premium-account.md section 0): production, with the account reads
// (api.foreversixty.gg /v1/me, /v1/devices, /v1/billing/checkout) answered from the fixtures the mocks read, so the
// signed-in states exist without an account. /b/:id is the API's own page: a missing id is fetched for real (the API's
// 404 page); a valid build is the API template (api/internal/site/page.go) filled with a Survival hunter record and the
// live planner island, as web/tests/e2e/shared-build.spec.ts stubs it. Writes measure.json (page height, h1, framed blocks).
//   node design/mocks/shoot_premium_before.mjs <out dir>   then   python3 design/mocks/ws_premium.py compare <out dir>
import { createRequire } from "node:module";
import { mkdirSync, readFileSync, writeFileSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const W = new URL("../../", import.meta.url).pathname;
const OUT = (process.argv[2] ?? "design/mocks/renders/premium-before") + "/";
mkdirSync(OUT, { recursive: true });
const SITE = "https://foreversixty.gg";
const j = (p) => JSON.parse(readFileSync(`${W}/${p}`, "utf8"));
const me34 = j("web/tests/e2e/fixtures/home/me-34.json");
const me60 = j("web/tests/e2e/fixtures/home/me-60.json");
const meNew = j("web/tests/e2e/fixtures/home/me-new.json");
const now = Date.now();
const iso = (ms) => new Date(now + ms).toISOString();
const ENT = { personal_portal: true, server_sims: true, retention: true, multi_compare: true, history: true, notifications: true, officer_views: false, roster_check: false, log_coach: true, supporter_mark: true };
const mePremium = { ...me60, entitlements: { ...ENT, billing: { source: "stripe", plan: "premium", status: "active", current_period_end: iso(86400e3 * 300), cancel_at_period_end: false, founder: true, interval: "yearly", amount_cents: 4000 } },
  server_sims: { runs_this_month: 4 }, discord_link_available: true, discord_link: { display_name: "thandrel" }, notifications: { connect: true, discord: null } };
const DEV = [{ id: "d1", name: "JHUNT-PC", platform: "windows/amd64", kind: "companion", created_at: iso(-86400e3 * 3), last_seen_at: iso(-180e3) }];
const ok = (data) => JSON.stringify({ ok: true, data, error: null, request_id: "r" });
const bad = (m) => JSON.stringify({ ok: false, data: null, error: { message: m }, request_id: "r" });
const cors = { "access-control-allow-origin": SITE, "access-control-allow-credentials": "true", "access-control-allow-headers": "content-type,x-csrf-token", "access-control-allow-methods": "GET,POST,OPTIONS" };
// The shared build: the Survival hunter of the planner mocks (planner-facts.json, state `guide`), as a record.
const PF = j("design/mocks/data/planner-facts.json");
const guide = PF.states.guide;
const idOf = Object.fromEntries(guide.trees.flatMap((t) => t.cells.map((c) => [c.name, c.id])));
const RECORD = { id: "k7Qm2xa", class_id: 3, race_id: 4, tree_version: PF.build, point_order: guide.points.map((p) => idOf[p.name]), gear: {}, title: "Survival raid build", created_at: iso(-86400e3 * 2), views: 41 };
const apiPage = readFileSync(`${W}/api/internal/site/testdata/chrome.html`, "utf8");
const header = apiPage.match(/<header[\s\S]*?<\/header>/)?.[0] ?? "";
const footer = apiPage.match(/<footer[\s\S]*?<\/footer>/)?.[0] ?? "";
const buildHtml = `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1">
<title>Survival raid build · 5/13/33 · Forever Sixty</title><link rel="stylesheet" href="${SITE}/planner-island.css"></head>
<body class="min-h-screen flex flex-col bg-bg">${header}<main id="main"><div id="planner" data-build='${JSON.stringify(RECORD).replace(/'/g, "&#39;")}' data-tree-version="${PF.build}"></div></main>${footer}
<script type="module" src="${SITE}/planner-island.js"></script></body></html>`;

async function stub(page, s) {
  await page.route("https://api.foreversixty.gg/**", async (route) => {
    const req = route.request(); const u = new URL(req.url());
    if (req.method() === "OPTIONS") return route.fulfill({ status: 204, headers: cors });
    const f = (body, status = 200) => route.fulfill({ status, headers: { ...cors, "content-type": "application/json" }, body });
    if (u.pathname === "/v1/me") return s.me ? f(ok(s.me)) : f(bad("Sign in"), 401);
    if (u.pathname === "/v1/devices") return s.me ? f(ok(s.devices ?? [])) : f(bad("Sign in"), 401);
    if (u.pathname === "/v1/billing/checkout") return s.checkout === "fail" ? f(bad("boom"), 500) : f(ok({ checkout_url: `${SITE}/premium?stub=1` }));
    if (u.pathname.startsWith("/v1/reports")) return f(ok({ rows: [], total: 0, page: 1, per_page: 100 }));
    return route.continue();
  });
  if (s.shared) await page.route(`${SITE}/b/k7Qm2xa`, (route) => route.fulfill({ status: 200, contentType: "text/html", body: buildHtml }));
}
const measures = {};
async function measure(page) {
  return page.evaluate(() => {
    const main = document.querySelector("main") ?? document.body;
    const top = (el) => (el ? Math.round(el.getBoundingClientRect().top + scrollY) : null);
    const framed = [...main.querySelectorAll("*")].filter((el) => {
      const cs = getComputedStyle(el); const r = el.getBoundingClientRect();
      return r.height > 60 && parseFloat(cs.borderTopWidth) >= 1 && cs.borderTopStyle !== "none" && parseFloat(cs.borderLeftWidth) >= 1 && cs.borderLeftStyle !== "none";
    }).length;
    const btn = [...main.querySelectorAll("a,button")].find((b) => /subscribe|sign in|see plans|fork|copy for/i.test(b.textContent ?? "") && b.getBoundingClientRect().height > 0);
    return { height: document.documentElement.scrollHeight, mainWidth: Math.round(main.getBoundingClientRect().width), h1: top(document.querySelector("h1")), h1Text: document.querySelector("h1")?.textContent?.trim() ?? null,
      framed, firstAction: btn ? { text: btn.textContent.trim().slice(0, 48), y: top(btn) } : null, scrollWidth: document.documentElement.scrollWidth };
  });
}
const b = await chromium.launch();
async function shot(name, path, s, widths, after) {
  for (const w of widths) {
    const ctx = await b.newContext({ viewport: { width: w, height: w < 500 ? 844 : 1000 } });
    if (s.me) await ctx.addCookies([{ name: "fs_csrf", value: "t", domain: "foreversixty.gg", path: "/" }]);
    if (s.hint) await ctx.addInitScript((v) => localStorage.setItem("fs.premiumSource", v), s.hint);
    const page = await ctx.newPage(); await stub(page, s);
    await page.goto(SITE + path, { waitUntil: "networkidle" }).catch(() => null); await page.waitForTimeout(1500);
    if (after) { await after(page); await page.waitForTimeout(900); }
    await page.screenshot({ path: `${OUT}${name}-${w}.png`, fullPage: true });
    await page.screenshot({ path: `${OUT}${name}-${w}-fold.png` });
    measures[`${name}-${w}`] = await measure(page);
    await ctx.close();
  }
}
const ALL = [390, 1440, 1920, 2000], TWO = [390, 1440];
await shot("premium-out", "/premium", {}, ALL);
await shot("premium-own", "/premium", { me: mePremium, hint: "own" }, TWO);
await shot("premium-canceled", "/premium?canceled=1", {}, [1440]);
await shot("checkout-success", "/premium/checkout?plan=premium&interval=yearly&status=success", { me: mePremium }, TWO);
await shot("checkout-error", "/premium/checkout?plan=premium&interval=yearly", { me: me34, checkout: "fail" }, [1440]);
await shot("checkout-guild", "/premium/checkout?plan=guild&interval=yearly", { me: me60 }, TWO);
await shot("account-out", "/account", {}, ALL);
await shot("account-new", "/account", { me: meNew, devices: [] }, TWO);
await shot("account-34", "/account", { me: me34, devices: [] }, ALL);
await shot("account-premium", "/account", { me: mePremium, devices: DEV }, [390, 1440, 2000]);
await shot("b-shared", "/b/k7Qm2xa", { shared: true }, ALL);
await shot("b-missing", "/b/zz9zz9zz", {}, TWO);
writeFileSync(`${OUT}measure.json`, JSON.stringify(measures, null, 1) + "\n");
await b.close();
