// Before captures of the live /setup and /connect (design/specs/2026-11-07-setup.md section 0): production, with the
// account reads (api.foreversixty.gg /v1/me, /v1/devices, the connect lookup and approve) answered from the fixtures the
// mocks read, so the signed-in states exist without an account.
//   node design/mocks/shoot_setup_before.mjs <out dir>   then   python3 design/mocks/ws_setup_compare.py <out dir>
import { createRequire } from "node:module";
import { mkdirSync, readFileSync } from "node:fs";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const W = new URL("../../", import.meta.url).pathname;
const OUT = (process.argv[2] ?? "design/mocks/renders/setup-before") + "/";
mkdirSync(OUT, { recursive: true });
const SITE = "https://foreversixty.gg";
const j = (p) => JSON.parse(readFileSync(`${W}/${p}`, "utf8"));
const me34 = j("web/tests/e2e/fixtures/home/me-34.json");
const meNew = j("web/tests/e2e/fixtures/home/me-new.json");
const now = Date.now();
const iso = (ms) => new Date(now - ms).toISOString();
// The live site reads its own clock, so the stubbed account is dated relative to now (the mocks pin 2026-11-06T10:00Z).
const meAddon = { ...me34, characters: me34.characters.map((c) => ({ ...c, build: { ...c.build, captured_at: iso(3600e3) } })) };
const meBnet = { ...me34, characters: me34.characters.map((c) => ({ ...c, build: { ...c.build, source: "blizzard" } })) };
const DEV = [{ id: "d1", name: "JHUNT-PC", platform: "windows/amd64", kind: "companion", created_at: iso(86400e3 * 3), last_seen_at: iso(180e3) }];
const ok = (data) => JSON.stringify({ ok: true, data, error: null, request_id: "r" });
const bad = (m) => JSON.stringify({ ok: false, data: null, error: { message: m }, request_id: "r" });
const cors = { "access-control-allow-origin": SITE, "access-control-allow-credentials": "true", "access-control-allow-headers": "content-type,x-csrf-token", "access-control-allow-methods": "GET,POST,OPTIONS" };
async function stub(page, s) {
  await page.route("https://api.foreversixty.gg/**", async (route) => {
    const req = route.request(); const u = new URL(req.url());
    if (req.method() === "OPTIONS") return route.fulfill({ status: 204, headers: cors });
    const f = (body, status = 200) => route.fulfill({ status, headers: { ...cors, "content-type": "application/json" }, body });
    if (u.pathname === "/v1/me") return s.me ? f(ok(s.me)) : f(bad("Sign in"), 401);
    if (u.pathname === "/v1/devices") return s.me ? f(ok(s.devices ?? [])) : f(bad("Sign in"), 401);
    if (u.pathname.startsWith("/v1/devices/connect/lookup")) return s.lookup === 404 ? f(bad("nf"), 404) : f(ok({ name: "JHUNT-PC", platform: "Windows", expires_in: 540 }));
    if (u.pathname === "/v1/devices/connect/approve") return s.approve === 409 ? f(bad("used"), 409) : s.approve === "abort" ? route.abort() : f(ok({ name: "JHUNT-PC", platform: "Windows" }));
    if (u.pathname === "/v1/devices/connect/deny") return route.fulfill({ status: 204, headers: cors });
    if (u.pathname.startsWith("/v1/reports")) return f(ok({ rows: [], total: 0, page: 1, per_page: 100 }));
    return route.continue();
  });
}
const b = await chromium.launch();
async function shot(name, path, s, widths, after) {
  for (const w of widths) {
    const ctx = await b.newContext({ viewport: { width: w, height: w < 500 ? 844 : 1000 } });
    if (s.me) await ctx.addCookies([{ name: "fs_csrf", value: "t", domain: "foreversixty.gg", path: "/" }]);
    const page = await ctx.newPage(); await stub(page, s);
    await page.goto(SITE + path, { waitUntil: "networkidle" }); await page.waitForTimeout(1200);
    if (after) { await after(page); await page.waitForTimeout(900); }
    await page.screenshot({ path: `${OUT}${name}-${w}.png`, fullPage: true });
    await page.screenshot({ path: `${OUT}${name}-${w}-fold.png` });
    await ctx.close();
  }
}
const ALL = [390, 1440, 1920, 2000], TWO = [390, 1440];
await shot("setup-out", "/setup", {}, ALL);
await shot("setup-new", "/setup", { me: meNew, devices: [] }, TWO);
await shot("setup-bnet", "/setup", { me: meBnet, devices: [] }, TWO);
await shot("setup-done", "/setup", { me: meAddon, devices: DEV }, [390, 1440, 2000]);
await shot("connect-out", "/connect?code=kq47tm", {}, ALL);
await shot("connect-pending", "/connect?code=KQ4-7TM", { me: meNew }, ALL);
await shot("connect-success", "/connect?code=KQ4-7TM", { me: meNew }, TWO, (p) => p.getByTestId("connect-approve").click());
await shot("connect-used", "/connect?code=KQ4-7TM", { me: meNew, approve: 409 }, TWO, (p) => p.getByTestId("connect-approve").click());
await shot("connect-failed", "/connect?code=KQ4-7TM", { me: meNew, approve: "abort" }, TWO, (p) => p.getByTestId("connect-approve").click());
await shot("connect-cancelled", "/connect?code=KQ4-7TM", { me: meNew }, TWO, (p) => p.getByTestId("connect-deny").click());
await shot("connect-expired", "/connect?code=KQ4-7TM", { me: meNew, lookup: 404 }, TWO);
await shot("connect-nocode", "/connect", {}, TWO);
await b.close();
