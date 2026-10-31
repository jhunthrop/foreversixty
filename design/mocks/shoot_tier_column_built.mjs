// design/mocks/shoot_tier_column_built.mjs
// Acceptance captures of the BUILT /tiers page for the tier column fix (spec
// design/specs/2026-10-31-tier-column.md), taken from `astro preview` on the fixture dist, signed in
// as a Fury Warrior so the "Your spec" row sits inside a tier group. The site is dark-only.
//
//   (cd web && FOREVER_DATA=fixture PUBLIC_API_BASE_URL=http://127.0.0.1:9 npm run build \
//      && npm run preview -- --port 4352)
//   TIERS_BASE=http://localhost:4352 TIERS_OUT=design/reviews/2026-10-31-tier-column \
//      node design/mocks/shoot_tier_column_built.mjs [prefix]
//        -> <out>/<prefix>-<width>.png (full page) and <prefix>-<width>-list.png (the list panel)
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";
import { resolve } from "node:path";

const require = createRequire(
  new URL("../../web/package.json", import.meta.url),
);
const { chromium } = require("@playwright/test");

const BASE = process.env.TIERS_BASE ?? "http://localhost:4352";
const OUT = resolve(
  process.env.TIERS_OUT ?? "design/reviews/2026-10-31-tier-column",
);
const PREFIX = process.argv[2] ?? "built";
const WIDTHS = [360, 1024, 1099, 1100, 1440, 2000];
const FONT_SETTLE_MS = 1500;

const envelope = (data) => ({
  status: 200,
  contentType: "application/json",
  body: JSON.stringify({ ok: true, data, error: null, request_id: "r" }),
});
const ME = {
  user: {
    id: 1,
    battletag: "Fixture#1",
    email: null,
    role: "user",
    anonymize: false,
  },
  characters: [
    {
      key: "us/normal/tester",
      region: "us",
      ruleset: "normal",
      name: "Tester",
      level: 60,
      class: "Warrior",
      spec: "Fury",
    },
  ],
  guilds: [],
};

mkdirSync(OUT, { recursive: true });
const browser = await chromium.launch();
for (const width of WIDTHS) {
  const context = await browser.newContext({
    viewport: { width, height: 900 },
    deviceScaleFactor: 2,
  });
  await context.addCookies([
    { name: "fs_csrf", value: "token", domain: "localhost", path: "/" },
  ]);
  const page = await context.newPage();
  await page.route("**/v1/me", (route) => route.fulfill(envelope(ME)));
  await page.goto(`${BASE}/tiers`);
  await page.locator(".tier-row.is-you").first().waitFor();
  await page.waitForTimeout(FONT_SETTLE_MS);
  await page.screenshot({
    path: `${OUT}/${PREFIX}-${width}.png`,
    fullPage: true,
  });
  await page
    .locator(
      '[data-testid="tier-faction-panel-alliance"] [data-testid="tier-list"]',
    )
    .screenshot({
      path: `${OUT}/${PREFIX}-${width}-list.png`,
    });
  if (width === 1440) await shootStates(page);
  await context.close();
}
await browser.close();

// Spec section 6: hover on a plain row and on the you-row, and the tank and healer lists.
async function shootStates(page) {
  const list = page.locator(
    '[data-testid="tier-faction-panel-alliance"] [data-testid="tier-list"]',
  );
  await page
    .locator(
      '[data-testid="tier-faction-panel-alliance"] .tier-row:not(.is-you)',
    )
    .first()
    .hover();
  await list.screenshot({ path: `${OUT}/${PREFIX}-1440-hover.png` });
  await page.locator(".tier-row.is-you").first().hover();
  await list.screenshot({ path: `${OUT}/${PREFIX}-1440-hover-you.png` });
  for (const role of ["tank", "healer"]) {
    await page.goto(`${BASE}/tiers/${role}`);
    await page.waitForTimeout(FONT_SETTLE_MS);
    await list.screenshot({ path: `${OUT}/${PREFIX}-1440-${role}.png` });
  }
}
