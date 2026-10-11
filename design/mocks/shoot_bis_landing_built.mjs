// design/mocks/shoot_bis_landing_built.mjs
// Acceptance captures of the BUILT /bis landing (spec 2026-10-10-bis-landing.md section 11), taken
// from `astro preview` of a real-data build with `/v1/me` stubbed, at the viewports and states the
// spec names: signed in (synced, no addon sync, stale, level 60, under 20, loading), signed out, at
// 1440, 2000, 1920, 1024, 768, 390 and 360; hover, focus-visible and active on a spec row, a band
// pill and the lead button; a Druid card; the strip with a non-default pill selected.
//
//   (cd web && npm run build && npm run preview -- --port 4422)
//   node design/mocks/shoot_bis_landing_built.mjs [name ...]   -> design/mocks/renders/bis-landing-built-*.png
//   python3 design/mocks/compose_bis_landing_built.py          -> the built-versus-mock composite
//
// The character is the mock's invented example (Obnoxious Yell, Fury Warrior, level 47) wearing the
// first six sourced picks of its band, as in gen_bis_landing.py; nothing here is a real player.
import { createRequire } from "node:module";
import { mkdirSync, readFileSync, writeFileSync } from "node:fs";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const BASE = process.env.BIS_BASE ?? "http://localhost:4422";
const OUT = new URL("./renders/", import.meta.url).pathname;
const REPO = new URL("../../", import.meta.url).pathname;
const BUILD = JSON.parse(readFileSync(`${REPO}web/src/data/active-build.json`, "utf8")).build;
const MIN = 60_000;
const DAY = 86_400_000;
const WORN_EXAMPLE = 6;
const FONT_SETTLE_MS = 2500;

const fury = JSON.parse(readFileSync(`${REPO}data/builds/${BUILD}/bis/warrior-fury.json`, "utf8"));
const bandOf = (band) =>
  fury.bands.find((b) => b.faction === "alliance" && b.band === band && (b.preset ?? "bare") !== "raid") ??
  fury.bands.find((b) => b.faction === "alliance" && b.band === band);
const gearAt = (band, count) =>
  Object.fromEntries(
    bandOf(band).slots
      .filter((s) => typeof s.item_id === "number")
      .slice(0, count)
      .map((s) => [s.slot, s.item_id]),
  );

const OBNOXIOUS = { name: "Obnoxious Yell", class: "Warrior", spec: "Fury", level: 47, faction: "alliance", ageMs: 12 * MIN };
const characterBody = (c, gear) => ({
  ok: true,
  data: {
    user: { id: 1, battletag: "Fixture#1", email: null, role: "user", anonymize: false },
    characters: [
      {
        key: "us/normal/obnoxious-yell",
        region: "us",
        ruleset: "normal",
        name: c.name,
        class: c.class,
        spec: c.spec,
        level: c.level,
        realm: "Living Flame",
        faction: c.faction,
        ...(c.build === false
          ? { build: { source: "blizzard", captured_at: new Date(Date.now() - c.ageMs).toISOString() } }
          : { build: { source: "addon", captured_at: new Date(Date.now() - c.ageMs).toISOString(), gear } }),
      },
    ],
    guilds: [],
  },
  error: null,
  request_id: "r",
});

const SIGNED_IN = { c: OBNOXIOUS, gear: () => gearAt(40, WORN_EXAMPLE) };
const STATES = {
  synced: SIGNED_IN,
  nosync: { c: { ...OBNOXIOUS, build: false }, gear: () => ({}) },
  stale: { c: { ...OBNOXIOUS, ageMs: 19 * DAY }, gear: () => gearAt(40, WORN_EXAMPLE) },
  l60: { c: { ...OBNOXIOUS, level: 60 }, gear: () => gearAt(60, 14) },
  under20: { c: { ...OBNOXIOUS, level: 12 }, gear: () => ({}) },
};

/** name -> { width, state (null signed out), height?, run?(page) } */
const SHOTS = {
  "1440": { width: 1440, state: "synced" },
  "signed-out-1440": { width: 1440, state: null },
  "nosync-1440": { width: 1440, state: "nosync" },
  "stale-1440": { width: 1440, state: "stale" },
  "l60-1440": { width: 1440, state: "l60" },
  "under20-1440": { width: 1440, state: "under20" },
  "2000": { width: 2000, state: "synced" },
  "1920": { width: 1920, state: "synced" },
  "1024": { width: 1024, state: "synced" },
  "768": { width: 768, state: "synced" },
  "390": { width: 390, state: "synced" },
  "390-signed-out": { width: 390, state: null },
  "360": { width: 360, state: "synced" },
  "lead-loading": { width: 1440, state: "synced", hold: true, clip: "[data-bis-lead]" },
  "strip-selected-50": { width: 1440, state: "synced", run: (p) => p.getByTestId("bis-landing-band-50").click(), clip: "[data-bis-strip]", hrefs: true },
  "druid-card": { width: 1440, state: null, clip: '[data-testid="bis-class-druid"]' },
  "row-hover": { width: 1440, state: null, clip: '[data-testid="bis-class-warrior"]', run: (p) => p.getByTestId("bis-spec-warrior-fury").hover() },
  "row-focus": { width: 1440, state: null, clip: '[data-testid="bis-class-warrior"]', run: (p) => p.getByTestId("bis-spec-warrior-fury").focus(), key: true },
  "row-active": { width: 1440, state: null, clip: '[data-testid="bis-class-warrior"]', run: async (p) => { await p.getByTestId("bis-spec-warrior-fury").hover(); await p.mouse.down(); } },
  "pill-hover": { width: 1440, state: null, clip: "[data-bis-strip]", run: (p) => p.getByTestId("bis-landing-band-30").hover() },
  "pill-focus": { width: 1440, state: null, clip: "[data-bis-strip]", run: (p) => p.getByTestId("bis-landing-band-30").focus(), key: true },
  "pill-active": { width: 1440, state: null, clip: "[data-bis-strip]", run: async (p) => { await p.getByTestId("bis-landing-band-30").hover(); await p.mouse.down(); } },
  "pill-selected-hover": { width: 1440, state: null, clip: "[data-bis-strip]", run: (p) => p.getByTestId("bis-landing-band-20").hover() },
  "cta-hover": { width: 1440, state: "synced", clip: "[data-bis-lead]", run: (p) => p.getByTestId("bis-lead-cta").hover() },
  "cta-focus": { width: 1440, state: "synced", clip: "[data-bis-lead]", run: (p) => p.getByTestId("bis-lead-cta").focus(), key: true },
  "cta-active": { width: 1440, state: "synced", clip: "[data-bis-lead]", run: async (p) => { await p.getByTestId("bis-lead-cta").hover(); await p.mouse.down(); } },
};

async function shoot(browser, name, shot) {
  const context = await browser.newContext({ viewport: { width: shot.width, height: 900 } });
  if (shot.state !== null) {
    const { c, gear } = STATES[shot.state];
    await context.addCookies([{ name: "fs_csrf", value: "token", domain: new URL(BASE).hostname, path: "/" }]);
    let release = () => {};
    const held = shot.hold ? new Promise((resolve) => { release = resolve; }) : null;
    await context.route("**/v1/me", async (route) => {
      if (held) await held;
      await route.fulfill({ status: 200, contentType: "application/json", body: JSON.stringify(characterBody(c, gear())) });
    });
    shot.release = release;
  }
  const page = await context.newPage();
  await page.goto(`${BASE}/bis`);
  await page.waitForTimeout(FONT_SETTLE_MS);
  if (shot.key) await page.keyboard.press("Tab"); // keyboard modality, so :focus-visible shows
  if (shot.run) {
    await shot.run(page);
    await page.waitForTimeout(400);
  }
  const out = `${OUT}bis-landing-built-${name}.png`;
  if (shot.hold) {
    await page.locator(shot.clip).screenshot({ path: out });
    shot.release();
  } else if (shot.clip) {
    await page.locator(shot.clip).screenshot({ path: out });
  } else {
    await page.screenshot({ path: out, fullPage: true });
  }
  if (shot.hrefs) {
    const hrefs = await page.locator("[data-spec-link]").evaluateAll((nodes) => nodes.map((n) => n.getAttribute("href")));
    writeFileSync(`${OUT}bis-landing-built-${name}-hrefs.txt`, `${hrefs.join("\n")}\n`);
  }
  await context.close();
  console.log(out);
}

mkdirSync(OUT, { recursive: true });
const names = process.argv.slice(2).length ? process.argv.slice(2) : Object.keys(SHOTS);
const browser = await chromium.launch();
for (const name of names) await shoot(browser, name, SHOTS[name]);
await browser.close();
