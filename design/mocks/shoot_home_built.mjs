// design/mocks/shoot_home_built.mjs
// Acceptance captures of the BUILT signed-in home page (spec 2026-10-10-home-signed-in-panel.md
// section 9), taken from `astro preview` of a FOREVER_DATA=fixture build with `/v1/me` stubbed
// and the real best-in-slot and item files routed in from data/builds/<active build>/.
//
//   (cd web && FOREVER_DATA=fixture npm run build && npm run preview -- --port 4377)
//   HOME_BASE=http://localhost:4377 node design/mocks/shoot_home_built.mjs [name ...]
//                                                  -> design/mocks/renders/home-panel-built-*.png
//   python3 design/mocks/compose_home_built.py     -> the built-versus-mock composite at 1440
//
// The characters are the invented examples gen_home_panel.py uses. A worn item is the pick's best
// sim-verified runner-up where one costs at least MIN_GAIN, else the pick itself, the same rule the
// mock applies, so the table is the build's own comparison. Playwright is found through web/.
import { createRequire } from "node:module";
import { mkdirSync, readFileSync } from "node:fs";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const BASE = process.env.HOME_BASE ?? "http://localhost:4377";
const OUT = new URL("./renders/", import.meta.url).pathname;
const ROOT = new URL("../../", import.meta.url).pathname;
const BUILD = JSON.parse(readFileSync(`${ROOT}web/src/data/active-build.json`, "utf8")).build;
const MIN_GAIN = 0.095;
const MIN = 60_000;
const HOUR = 3_600_000;

const character = (name, cls, spec, level, faction, realm, guild, ageMs) => ({
  name,
  cls,
  spec,
  level,
  faction,
  realm,
  guild,
  ageMs,
});
const OBNOX = character("Obnoxious Yell", "warrior", "fury", 27, "alliance", "Living Flame", "Midnight Oil", 12 * MIN);
const TWO_LINE = { ...OBNOX, name: "Sir Obnoxious Yellington" };
const FROST = character("Frostbyte", "mage", "frost", 28, "horde", "Stonehearth", undefined, 12 * MIN);
const TREE = character("Treewalker", "druid", "balance", 24, "alliance", "Living Flame", undefined, 456 * HOUR);
const OAK = character("Oakheart", "paladin", "protection", 25, "alliance", "Living Flame", undefined, 12 * MIN);

const keyOf = (c) => `us/normal/${c.name.toLowerCase().replace(/\s+/g, "-")}`;
const titleCase = (word) => word.charAt(0).toUpperCase() + word.slice(1);
const specName = (c) => c.spec.split("-").map(titleCase).join(" ");

const readJson = (relative) => JSON.parse(readFileSync(`${ROOT}data/builds/${BUILD}/${relative}`, "utf8"));

function bandOf(c) {
  const file = readJson(`bis/${c.cls}-${c.spec}.json`);
  const level = Math.floor(Math.max(20, Math.min(60, c.level)) / 10) * 10;
  const candidates = file.bands.filter((b) => b.band === level && b.faction === c.faction);
  return (
    candidates.find((b) => b.preset === "raid") ??
    candidates.find((b) => b.preset === "bare") ??
    candidates[0]
  );
}

function wornGear(c) {
  const gear = {};
  for (const slot of bandOf(c).slots) {
    if (slot.item_id === undefined) continue;
    const runnerUp = (slot.alternatives ?? []).find((a) => (a.dps_delta ?? 0) <= -MIN_GAIN);
    gear[slot.slot] = runnerUp?.item_id ?? slot.item_id;
  }
  return gear;
}

const meCharacter = (c) => ({
  key: keyOf(c),
  region: "us",
  ruleset: "normal",
  name: c.name,
  class: titleCase(c.cls),
  spec: specName(c),
  race: bandOf(c).race.split("-").map(titleCase).join(" "),
  realm: c.realm,
  level: c.level,
  faction: c.faction,
  ...(c.guild === undefined ? {} : { guild: { id: 9, name: c.guild, verified: true } }),
  build: {
    source: "addon",
    captured_at: new Date(Date.now() - c.ageMs).toISOString(),
    level: c.level,
    data_build: BUILD,
    gear: wornGear(c),
  },
});

const envelope = (data) => ({
  status: 200,
  contentType: "application/json",
  body: JSON.stringify({ ok: true, data, error: null, request_id: "r" }),
});

async function stubApi(page, characters) {
  await page.context().addCookies([{ name: "fs_csrf", value: "token", domain: "localhost", path: "/" }]);
  await page.route("**/v1/me", (route) =>
    route.fulfill(
      envelope({
        user: { id: 1, battletag: "Fixture#1", email: null, role: "user", anonymize: false },
        characters: characters.map(meCharacter),
        guilds: [],
      }),
    ),
  );
  await page.route("**/v1/sims**", (route) =>
    route.fulfill(
      envelope({
        rows: [{ sim_id: "s1", spec: `${characters[0].cls}-${characters[0].spec}`, dps: 26.4, engine_version: "1", created_at: "2026-10-09T00:00:00Z", title: "", kind: "run" }],
        total: 1,
        page: 1,
        per_page: 1,
      }),
    ),
  );
  await page.route("**/rating", (route) => route.fulfill(envelope(null)));
  for (const kind of ["bis", "items"]) {
    await page.route(`**/data/*/${kind}/*.json`, (route) => {
      const file = new URL(route.request().url()).pathname.split("/").slice(-2).join("/");
      try {
        route.fulfill({ status: 200, contentType: "application/json", body: readFileSync(`${ROOT}data/builds/${BUILD}/${file}`, "utf8") });
      } catch {
        route.continue();
      }
    });
  }
}

async function shoot(browser, name, { width, height = 900, characters, current }) {
  const context = await browser.newContext({ viewport: { width, height } });
  const page = await context.newPage();
  await stubApi(page, characters);
  const pointer = {
    source: "armory",
    ref: keyOf(current),
    label: `${current.name} · ${specName(current)} ${titleCase(current.cls)}`,
    classSlug: current.cls,
    savedAt: new Date().toISOString(),
  };
  await page.addInitScript((value) => window.localStorage.setItem("fs.currentCharacter", JSON.stringify(value)), pointer);
  await page.goto(`${BASE}/`);
  await page.getByTestId("home-hero-card-bis-value").waitFor({ timeout: 20_000 });
  await page.getByTestId("home-upgrades-list").waitFor({ timeout: 20_000 });
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(600);
  await page.screenshot({ path: `${OUT}home-panel-built-${name}.png`, fullPage: true });
  const table = await page.getByTestId("home-upgrades").boundingBox();
  const hero = await page.getByTestId("home-account-panel").boundingBox();
  console.log(`${name}: hero height=${Math.round(hero.height)}`);
  console.log(`${name}: upgrades section top y=${Math.round(table.y)}`);
  await context.close();
}

const SHOTS = {
  "1440": { width: 1440, characters: [OBNOX, FROST, TREE, OAK], current: OBNOX },
  "1440-one": { width: 1440, characters: [OBNOX], current: OBNOX },
  "1440-twoline": { width: 1440, characters: [TWO_LINE, FROST, TREE, OAK], current: TWO_LINE },
  "2000": { width: 2000, characters: [OBNOX, FROST, TREE, OAK], current: OBNOX },
  "1280": { width: 1280, characters: [OBNOX, FROST, TREE, OAK], current: OBNOX },
  "1024": { width: 1024, characters: [OBNOX, FROST, TREE, OAK], current: OBNOX },
  "390": { width: 390, height: 844, characters: [OBNOX, FROST, TREE, OAK], current: OBNOX },
};

mkdirSync(OUT, { recursive: true });
const wanted = process.argv.slice(2);
const browser = await chromium.launch();
for (const [name, options] of Object.entries(SHOTS)) {
  if (wanted.length === 0 || wanted.includes(name)) await shoot(browser, name, options);
}
await browser.close();
