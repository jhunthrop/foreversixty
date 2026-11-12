// Captures the creator program's pages from a built site for the persona reviews:
//   cd web && npx astro preview --port 4399   (after `npm run build` with FOREVER_DATA=fixture)
//   node design/reviews/2026-11-11-creators/capture.mjs [http://localhost:4399]
// -> design/reviews/2026-11-11-creators/build/<state>-<width>.png, full page, at 1440 and 390.
// Every API call is answered from the e2e fixtures (web/tests/e2e/fixtures/creators, the owner from
// fixtures/home/me-60.json), the way web/tests/e2e/creators.spec.ts serves them, so no server is
// touched and the states are the spec's: the public page signed out; apply empty, filled, pending,
// approved and declined; the creator's stats; the owner's creators and statements, the payouts-off
// notice and the Approve confirm. Playwright is found through web/.
import { createRequire } from "node:module";
import { mkdirSync, readFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const ROOT = join(HERE, "../../..");
const require = createRequire(join(ROOT, "web/package.json"));
const { chromium } = require("@playwright/test");
const BASE = process.argv[2] ?? "http://localhost:4399";
const OUT = join(HERE, "build");
const FIX = join(ROOT, "web/tests/e2e/fixtures");
const WIDTHS = [1440, 390];
mkdirSync(OUT, { recursive: true });

const read = (name) => JSON.parse(readFileSync(join(FIX, name), "utf8"));
const owner = read("home/me-60.json");
const envelope = (data, status = 200) => ({
  status,
  contentType: "application/json",
  body: JSON.stringify({
    ok: status < 400,
    data: status < 400 ? data : null,
    error: status < 400 ? null : { code: "x", message: "x" },
    request_id: "c",
  }),
});

async function signedIn(page, extra = {}, role = "admin") {
  await page.context().addCookies([{ name: "fs_csrf", value: "tok", domain: "localhost", path: "/" }]);
  await page.route("**/v1/me", (route) => route.fulfill(envelope({ ...owner, user: { ...owner.user, role }, ...extra })));
}
const application = (page, data) => page.route("**/v1/creators/application", (route) => route.fulfill(envelope(data)));
const noApplication = { status: null, created_at: null, creator: false, reapply_on: null };

async function fillForm(page) {
  await page.getByTestId("apply-channel").fill("https://twitch.tv/grimbrew");
  await page.getByRole("radio", { name: "Twitch" }).check();
  await page.getByTestId("apply-audience").fill("1200");
  await page
    .getByTestId("apply-about")
    .fill("Three nights a week of Forever raiding on a Fury warrior, plus a Sunday leveling stream for new players.");
  await page.getByTestId("apply-email").fill("grimbrew@example.com");
  await page.getByTestId("apply-disclose").locator("input").check();
}

const STATES = [
  { name: "creators", url: "/creators", ready: "creators-section-apply", setup: (page) => page.route("**/v1/me", (r) => r.fulfill(envelope(null, 401))) },
  { name: "apply-empty", url: "/creators/apply", ready: "apply-form", setup: async (page) => { await signedIn(page, {}, "member"); await application(page, noApplication); } },
  { name: "apply-filled", url: "/creators/apply", ready: "apply-form", setup: async (page) => { await signedIn(page, {}, "member"); await application(page, noApplication); }, then: fillForm },
  { name: "apply-pending", url: "/creators/apply", ready: "apply-status-pending", setup: async (page) => { await signedIn(page, {}, "member"); await application(page, { status: "pending", created_at: "2026-12-08T16:20:00Z", creator: false, reapply_on: null }); } },
  { name: "apply-approved", url: "/creators/apply", ready: "apply-status-approved", setup: async (page) => { await signedIn(page, {}, "member"); await application(page, { status: "approved", created_at: "2026-12-08T16:20:00Z", creator: false, reapply_on: null }); } },
  { name: "apply-declined", url: "/creators/apply", ready: "apply-status-declined", setup: async (page) => { await signedIn(page, {}, "member"); await application(page, { status: "declined", created_at: "2026-12-08T16:20:00Z", creator: false, reapply_on: "2027-01-08" }); } },
  { name: "stats", url: "/creators/stats", ready: "stats-figures", setup: async (page) => { await signedIn(page, { creator: true }, "member"); await page.route("**/v1/creators/me", (r) => r.fulfill(envelope(read("creators/me.json")))); } },
  {
    name: "admin-creators", url: "/admin/creators", ready: "creator-row-grimbrew",
    setup: async (page) => {
      await signedIn(page, { creator_requests: 2 });
      await page.route("**/v1/admin/creators", (r) => r.fulfill(envelope({ creators: read("creators/creators.json"), stripe_enabled: true, payouts_enabled: true, connect_denied: false })));
      await page.route("**/v1/admin/creators/requests", (r) => r.fulfill(envelope({ requests: read("creators/requests.json") })));
    },
  },
  {
    name: "admin-creators-edit", url: "/admin/creators", ready: "creator-row-grimbrew",
    setup: async (page) => {
      await signedIn(page);
      await page.route("**/v1/admin/creators", (r) => r.fulfill(envelope({ creators: read("creators/creators.json"), stripe_enabled: true, payouts_enabled: true, connect_denied: false })));
      await page.route("**/v1/admin/creators/requests", (r) => r.fulfill(envelope({ requests: [] })));
    },
    then: (page) => page.getByTestId("creator-edit-grimbrew").click(),
  },
  {
    name: "admin-payouts", url: "/admin/creators/payouts", ready: "payout-row-1",
    setup: async (page) => { await signedIn(page); await page.route("**/v1/admin/creators/statements**", (r) => r.fulfill(envelope(read("creators/statements.json")))); },
  },
  {
    name: "admin-payouts-off", url: "/admin/creators/payouts", ready: "payout-row-1",
    setup: async (page) => { await signedIn(page); await page.route("**/v1/admin/creators/statements**", (r) => r.fulfill(envelope({ ...read("creators/statements.json"), payouts_enabled: false }))); },
  },
  {
    name: "admin-payouts-confirm", url: "/admin/creators/payouts", ready: "payout-row-1",
    setup: async (page) => {
      await signedIn(page);
      await page.route("**/v1/admin/creators/statements**", (r) => r.fulfill(envelope(read("creators/statements.json"))));
      await page.route("**/v1/admin/creators/balance", (r) => r.fulfill(envelope({ available_cents: 50000 })));
    },
    then: async (page) => { await page.getByTestId("payout-approve-1").click(); await page.getByTestId("pay-balance").waitFor(); },
  },
];

const only = process.argv.slice(3);
const browser = await chromium.launch();
for (const state of STATES) {
  if (only.length > 0 && !only.includes(state.name)) continue;
  for (const width of WIDTHS) {
    const context = await browser.newContext({ viewport: { width, height: width < 500 ? 844 : 1000 }, reducedMotion: "reduce" });
    const page = await context.newPage();
    await state.setup(page);
    await page.goto(`${BASE}${state.url}`, { waitUntil: "networkidle" });
    await page.getByTestId(state.ready).first().waitFor({ timeout: 15000 });
    if (state.then) await state.then(page);
    await page.evaluate(() => document.fonts.ready);
    await page.waitForTimeout(400);
    await page.screenshot({ path: join(OUT, `${state.name}-${width}.png`), fullPage: true });
    await context.close();
    console.log(`${state.name}-${width}.png`);
  }
}
await browser.close();
