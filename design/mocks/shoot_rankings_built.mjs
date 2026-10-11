// design/mocks/shoot_rankings_built.mjs
// Acceptance captures of the BUILT /rankings page (spec 2026-10-11-rankings.md section 8), taken
// from `astro preview` with the API stubbed to the mock's example world (gen_rankings.py: Obnoxious
// Yell, 14th of 212 Fury Warriors on Onyxia; the pinned state uses Thornwick Pell at 204th), at
// 1440 signed in and out, 2000, 1024 and 390, plus the phone's open filters and the pinned state.
//
//   (cd web && FOREVER_DATA=fixture npm run build && npm run preview -- --port 4351)
//   RANKINGS_BASE=http://localhost:4351 node design/mocks/shoot_rankings_built.mjs [name ...]
//        -> design/mocks/renders/rankings-built-*.png
//   python3 design/mocks/compose_rankings_built.py     -> the built-versus-mock composite
//
// Every figure here is the mock's invented example; nothing is a real player.
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const BASE = process.env.RANKINGS_BASE ?? "http://localhost:4351";
const OUT = new URL("./renders/", import.meta.url).pathname;
const FONT_SETTLE_MS = 2500;
const MIN = 60_000;
const DAY = 86_400_000;
// The page's clock is fixed to the mock's world (Dec 13, phase 1), so "Phase 1", "Updated Dec 13" and
// the relative dates read as they do in gen_rankings.py whatever day the capture runs.
const NOW = new Date("2026-12-13T12:00:00Z");
const OWN_KEY = "us/normal/obnoxious-yell";
const OWN_RANK = 14;
const TOTAL = 212;
const ONYXIA = { id: 9001, name: "Onyxia", slug: "onyxia" };

// rank, name, class, guild, size, value, executed %, days ago, split, length (ms). Gen_rankings.py's board.
const NAMES = ["Thorgrim Ashveil", "Valkyra", "Brannoch", "Kaelith Rowe", "Rhoswen", "Dunmar", "Tessiel Vane", "Ironvow", "Gaelwyn", "Hrafnild", "Morvain Holt", "Sellar", "Orsolya", "Obnoxious Yell", "Bladewick", "Draveth"];
const GUILDS = ["The Last Watch", "Stonehearth Vanguard", "", "Olympus XXVII", "The Last Watch", "", "Ember Court", "Olympus XXVII", "", "Stonehearth Vanguard", "", "Ember Court", "", "Olympus XXVII", "", "The Last Watch"];
const VALUES = [842, 821, 805, 790, 771, 755, 741, 728, 690, 672, 655, 641, 622, 612, 605, 598];
const EXEC = [97, 95, 96, 93, 94, 91, 92, 90, 89, 88, 90, 86, 87, 94, 85, 84];
const DAYS = [1, 2, 4, 3, 1, 5, 2, 6, 9, 4, 11, 2, 8, 3, 12, 5];
const REGION = ["us", "eu", "us", "us", "us", "eu", "us", "us", "us", "eu", "us", "us", "eu", "us", "us", "us"];
const SPLIT = ["31/20/0", "31/20/0", "30/21/0", "31/20/0", "31/20/0", "31/20/0", "30/21/0", "31/20/0", "31/20/0", "31/20/0", "30/21/0", "31/20/0", "31/20/0", "31/20/0", "30/21/0", "31/20/0"];
const LENGTHS = ["2:05", "2:11", "1:52", "2:18", "2:24", "3:40", "2:31", "2:58", "2:40", "2:47", "3:12", "2:52", "4:05", "2:41", "2:49", "3:20"];
const slug = (name) => name.toLowerCase().replace(/\s+/g, "-");
const lengthMs = (text) => { const [m, s] = text.split(":").map(Number); return (m * 60 + s) * 1000; };

const row = (seed, rank) => ({
  rank,
  player: { key: seed.key, name: seed.name, class: seed.cls, spec: seed.spec },
  ...(seed.guild ? { guild: { name: seed.guild, ruleset: "normal", region: seed.region } } : {}),
  value: seed.value, size: 40, fought_at: new Date(NOW.getTime() - seed.days * DAY - 10 * MIN).toISOString(),
  duration_ms: lengthMs(seed.length), talent_split: seed.split, build_id: `b${rank}`, trinkets: [], buff_count: 8,
  report_id: `report${rank}`, fight_index: 1, state: "ok", execution_score: seed.exec / 100,
});
const warriors = NAMES.map((name, i) => ({
  name, cls: "Warrior", spec: "Fury", key: name === "Obnoxious Yell" ? OWN_KEY : `us/normal/${slug(name)}`,
  guild: GUILDS[i], region: REGION[i], value: VALUES[i], exec: EXEC[i], days: DAYS[i], split: SPLIT[i], length: LENGTHS[i],
}));
const mixed = [
  { name: "Reglitch", cls: "Priest", spec: "Holy", guild: "", region: "us", value: 905, exec: 96, days: 1, split: "0/31/20", length: "2:09" },
  { name: "Mishvamp", cls: "Shaman", spec: "Enhancement", guild: "Olympus XXVII", region: "us", value: 872, exec: 95, days: 2, split: "31/0/20", length: "2:14" },
  { name: "Thorgrim Ashveil", cls: "Warrior", spec: "Fury", guild: "The Last Watch", region: "us", value: 842, exec: 97, days: 1, split: "31/20/0", length: "2:05" },
  { name: "Valkyra", cls: "Warrior", spec: "Fury", guild: "Stonehearth Vanguard", region: "eu", value: 821, exec: 95, days: 2, split: "31/20/0", length: "2:11" },
  { name: "Quillfeather", cls: "Rogue", spec: "Combat", guild: "", region: "us", value: 830, exec: 90, days: 3, split: "31/20/0", length: "1:58" },
  { name: "Hanabanana", cls: "Priest", spec: "Shadow", guild: "Ember Court", region: "eu", value: 801, exec: 91, days: 4, split: "0/31/20", length: "2:20" },
].sort((a, b) => b.value - a.value).map((seed, i) => ({ ...seed, key: `${seed.region}/normal/${slug(seed.name)}` }));
const far = { name: "Thornwick Pell", cls: "Warrior", spec: "Fury", key: "us/normal/thornwick-pell", guild: "Ember Court", region: "us", value: 410, exec: 71, days: 20, split: "31/20/0", length: "3:58" };

const board = ({ signedIn, pinned }) => {
  if (!signedIn) return { rows: mixed.map((seed, i) => row(seed, i + 1)), total: mixed.length, deduped: true };
  const rows = warriors.map((seed, i) => row(seed, i + 1));
  const own = pinned ? { rank: 204, row: row({ ...far, key: OWN_KEY, name: "Obnoxious Yell" }, 204) } : { rank: OWN_RANK, row: rows[OWN_RANK - 1] };
  return { rows: pinned ? rows.filter((r) => r.player.key !== OWN_KEY) : rows, total: TOTAL, deduped: true, character_rank: own };
};

const meBody = () => ({
  ok: true,
  data: {
    user: { id: 1, battletag: "Fixture#1", email: null, role: "user", anonymize: false },
    characters: [{
      key: OWN_KEY, region: "us", ruleset: "normal", name: "Obnoxious Yell", class: "Warrior", spec: "Fury", race: "Gnome",
      level: 60, realm: "Living Flame", faction: "alliance",
      build: { source: "addon", captured_at: new Date(NOW.getTime() - 12 * MIN).toISOString(), gear: { head: 1 } },
    }],
    guilds: [],
  },
  error: null, request_id: "r",
});
const envelope = (data) => ({ status: 200, contentType: "application/json", body: JSON.stringify({ ok: true, data, error: null, request_id: "r" }) });

/** name -> { width, signedIn, pinned?, open? (phone filters), clip? } */
const SHOTS = {
  "1440": { width: 1440, signedIn: true },
  "signed-out-1440": { width: 1440, signedIn: false },
  "2000": { width: 2000, signedIn: true },
  "1024": { width: 1024, signedIn: true },
  "390": { width: 390, signedIn: true },
  "390-open": { width: 390, signedIn: true, open: true },
  "390-signed-out": { width: 390, signedIn: false },
  "pinned-1440": { width: 1440, signedIn: true, pinned: true },
};

async function shoot(browser, name, shot) {
  const context = await browser.newContext({ viewport: { width: shot.width, height: 900 } });
  if (shot.signedIn) {
    await context.addCookies([{ name: "fs_csrf", value: "token", domain: new URL(BASE).hostname, path: "/" }]);
    await context.route("**/v1/me", (route) => route.fulfill({ status: 200, contentType: "application/json", body: JSON.stringify(meBody()) }));
    await context.addInitScript((pointer) => {
      window.localStorage.setItem("fs.currentCharacter", JSON.stringify({ ...pointer, savedAt: new Date().toISOString() }));
    }, { source: "armory", ref: OWN_KEY, label: "Obnoxious Yell · Fury Warrior", classSlug: "warrior" });
  } else {
    await context.route("**/v1/me", (route) => route.fulfill({ status: 401, contentType: "application/json", body: '{"ok":false,"data":null,"error":null,"request_id":"r"}' }));
  }
  await context.route("**/v1/encounters", (route) => route.fulfill(envelope({ rows: [ONYXIA, { id: 9002, name: "Warden Kelthas", slug: "warden-kelthas" }] })));
  await context.route("**/v1/characters/us/normal/obnoxious-yell", (route) => route.fulfill(envelope({
    character: { name: "Obnoxious Yell", region: "us", ruleset: "normal", class: "Warrior" }, best: [], builds_seen: [],
    history: [{ encounter: "Onyxia", encounter_id: ONYXIA.id, difficulty: 8, kill: true, metric: "dps", value: 612, fought_at: new Date(NOW.getTime() - 3 * DAY).toISOString(), report_id: "report14", fight_index: 1, execution_score: 0.94 }],
  })));
  await context.route("**/v1/rankings?**", (route) => route.fulfill(envelope({
    ...board({ signedIn: shot.signedIn, pinned: shot.pinned === true }), page: 1, per_page: 100, updated_at: "2026-12-13T09:00:00Z",
  })));
  const page = await context.newPage();
  await page.clock.install({ time: NOW });
  await page.goto(`${BASE}/rankings`);
  await page.getByTestId("ranking-1").waitFor();
  await page.waitForTimeout(FONT_SETTLE_MS);
  if (shot.open) {
    await page.getByTestId("rankings-filters-toggle").click();
    await page.waitForTimeout(300);
  }
  const out = `${OUT}rankings-built-${name}.png`;
  await page.screenshot({ path: out, fullPage: true });
  await context.close();
  console.log(out);
}

mkdirSync(OUT, { recursive: true });
const names = process.argv.slice(2).length ? process.argv.slice(2) : Object.keys(SHOTS);
const browser = await chromium.launch();
for (const name of names) await shoot(browser, name, SHOTS[name]);
await browser.close();
