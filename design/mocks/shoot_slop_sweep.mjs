// Captures the 2026-11-09 slop-sweep pages at 1440, before (production) and after (a local preview).
// The character page is the owner's view of a level 60 with one talent point unspent and a graded
// spec: /v1/me, the character and /v1/home are stubbed so both builds draw the same head.
//   node design/mocks/shoot_slop_sweep.mjs https://foreversixty.gg design/reviews/2026-11-09-slop-sweep/before
//   (web) npm run preview -- --port 4517   then   node design/mocks/shoot_slop_sweep.mjs http://localhost:4517 design/reviews/2026-11-09-slop-sweep/after
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const BASE = process.argv[2] ?? "http://localhost:4517";
const OUT = `${process.argv[3] ?? "design/reviews/2026-11-09-slop-sweep/after"}/`;
const WIDTH = 1440;
const HEIGHT = 1000;
const SETTLE_MS = 800;
mkdirSync(OUT, { recursive: true });

/** [name, path]: the pages with the most changed copy, plus the character page. */
const PAGES = (process.env.SLOP_PAGES ?? "").split(",").filter(Boolean).map((entry) => entry.split("="));

const KEY = "us/hardcore/elyra-duskvale";
const envelope = (data) => ({
  status: 200,
  contentType: "application/json",
  body: JSON.stringify({ ok: true, data, error: null, request_id: "r" }),
});

async function stubOwner(page) {
  const { createJiti } = require("jiti");
  const fixture = new URL("../../web/tests/e2e/support/character-fixture.ts", import.meta.url).pathname;
  const { CHARACTER, GEAR, GEAR_ITEMS, routeGearIcons } = await createJiti(import.meta.url).import(fixture);
  await page.route("**/v1/me", (route) =>
    route.fulfill(
      envelope({
        user: { id: 1, battletag: "F#1", email: null, role: "user", anonymize: false },
        characters: [
          {
            key: KEY,
            region: "us",
            ruleset: "hardcore",
            name: "Elyra Duskvale",
            class: "Priest",
            level: 60,
            build: { source: "addon", captured_at: "2026-11-08T21:40:00Z", talents: { points: [15, 30, 5] } },
          },
        ],
        guilds: [],
      }),
    ),
  );
  await page.route("**/v1/home?*", (route) =>
    route.fulfill(
      envelope({
        import_sim: null,
        movers: null,
        guild_sims: null,
        spec_week: { spec: "priest-holy", band: 60, faction: "alliance", role: "healer", grade: "D", last_week_grade: null, place: 4, of: 9 },
      }),
    ),
  );
  await page.route(`**/v1/characters/${KEY}`, (route) => route.fulfill(envelope({ ...CHARACTER.data, ...GEAR })));
  await page.route("**/data/*/items/priest.json", (route) =>
    route.fulfill({ status: 200, contentType: "application/json", body: JSON.stringify(GEAR_ITEMS) }),
  );
  await routeGearIcons(page);
}

const browser = await chromium.launch();
const page = await browser.newPage({ viewport: { width: WIDTH, height: HEIGHT }, deviceScaleFactor: 1 });
await stubOwner(page);
for (const [name, path] of [["character", `/character/${KEY}`], ...PAGES]) {
  await page.goto(`${BASE}${path}`, { waitUntil: "networkidle" });
  await page.waitForTimeout(SETTLE_MS);
  await page.screenshot({ path: `${OUT}${name}-${WIDTH}.png`, fullPage: true });
}
await browser.close();
console.log(`wrote ${PAGES.length + 1} captures to ${OUT}`);
