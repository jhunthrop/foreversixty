// design/mocks/shoot_logs_built.mjs
// Acceptance captures of the BUILT /logs landing (spec 2026-10-11-logs.md section 6), taken from
// `astro preview` with the API stubbed: signed in with a paired device at 1440, 1024 and 390 (the
// companion above Upload), signed out at 1440, and a signed-in unpaired capture at 1440 (Upload
// first). The character, its reports, its latest kill and the device are the mock's invented
// examples (gen_logs.py), nothing here is a real player.
//
//   (cd web && npm run build && npm run preview -- --port 4422)
//   node design/mocks/shoot_logs_built.mjs [name ...]   -> design/mocks/renders/logs-built-*.png
//   python3 design/mocks/compose_logs_built.py          -> the built-versus-mock composite
import { createRequire } from "node:module";
import { mkdirSync, readFileSync } from "node:fs";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const BASE = process.env.LOGS_BASE ?? "http://localhost:4422";
const OUT = new URL("./renders/", import.meta.url).pathname;
const REPO = new URL("../../", import.meta.url).pathname;
const SAMPLE_ID = JSON.parse(readFileSync(`${REPO}web/src/data/sample-report.json`, "utf8")).id;
const MIN = 60_000;
const DAY = 86_400_000;
const FONT_SETTLE_MS = 2500;

const ago = (ms) => new Date(Date.now() - ms).toISOString();
const envelope = (data, status = 200) => ({
  status,
  contentType: "application/json",
  body: JSON.stringify({ ok: status < 400, data, error: null, request_id: "r" }),
});

const ME = {
  user: { id: 1, battletag: "Fixture#1", email: null, role: "user", anonymize: false },
  characters: [
    {
      key: "us/normal/obnoxious-yell",
      region: "us",
      ruleset: "normal",
      name: "Obnoxious Yell",
      class: "Warrior",
      spec: "Fury",
      race: "Gnome",
      level: 60,
      realm: "Living Flame",
      faction: "alliance",
      build: { source: "addon", captured_at: ago(12 * MIN) },
    },
  ],
  guilds: [{ id: 5, region: "us", ruleset: "normal", name: "Olympus XXVII", verified: true }],
};

const report = (id, title, zone, daysAgo, fights, kills) => ({
  id,
  title,
  zone,
  status: "complete",
  visibility: "public",
  created_at: ago(daysAgo * DAY + 60_000),
  fight_count: fights,
  kill_count: kills,
});
const REPORTS = [
  report("r1", "Onyxia, wipes", "Onyxia's Lair", 1, 6, 0),
  report("r2", "Onyxia, first kill", "Onyxia's Lair", 3, 4, 1),
  report("r3", "Sanguine Depths run", "Sanguine Depths", 18, 21, 4),
];
const DEVICE = { id: "d1", name: "MacBook Pro", platform: "macOS", created_at: ago(30 * DAY), last_seen_at: ago(4 * MIN) };
const KILL = {
  encounter: "Onyxia",
  encounter_id: 1084,
  difficulty: 1,
  metric: "dps",
  value: 612.4,
  fought_at: ago(3 * DAY + 60_000),
  report_id: "r2",
  fight_index: 3,
  execution_score: null,
};
const SAMPLE_META = {
  id: SAMPLE_ID,
  title: "Sanguine Depths, sample log",
  visibility: "public",
  owner: null,
  zone: "Sanguine Depths",
  status: "complete",
  engine_version: "0.1.0",
  created_at: "2026-09-15T12:55:13Z",
  data_base_url: `/logs-data/reports/${SAMPLE_ID}`,
  players: [],
  fights: Array.from({ length: 12 }, (_, i) => ({ index: i + 1, kind: "encounter", name: "Boss", kill: true, in_progress: false })),
};

const SHOTS = {
  "1440": { width: 1440, signedIn: true, paired: true },
  "unpaired-1440": { width: 1440, signedIn: true, paired: false },
  "signed-out-1440": { width: 1440, signedIn: false },
  "1024": { width: 1024, signedIn: true, paired: true },
  "390": { width: 390, signedIn: true, paired: true },
};

async function shoot(browser, name, shot) {
  const context = await browser.newContext({ viewport: { width: shot.width, height: 900 } });
  await context.route("**/v1/me", (route) => route.fulfill(shot.signedIn ? envelope(ME) : envelope(null, 401)));
  await context.route("**/v1/devices", (route) => route.fulfill(envelope(shot.paired ? [DEVICE] : [])));
  await context.route("**/v1/reports?mine=1**", (route) => route.fulfill(envelope({ rows: REPORTS, total: REPORTS.length, page: 1, per_page: 100 })));
  await context.route("**/v1/guilds/*/reports**", (route) => route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 100 })));
  await context.route("**/v1/reports/recent**", (route) => route.fulfill(envelope({ rows: [] })));
  await context.route(`**/v1/reports/${SAMPLE_ID}`, (route) => route.fulfill(envelope(SAMPLE_META)));
  await context.route("**/v1/characters/us/normal/obnoxious-yell", (route) =>
    route.fulfill(envelope({ player: { key: "us/normal/obnoxious-yell", name: "Obnoxious Yell", class: "Warrior" }, best: [], history: [KILL], builds_seen: [] })),
  );
  if (shot.signedIn) await context.addCookies([{ name: "fs_csrf", value: "token", domain: new URL(BASE).hostname, path: "/" }]);
  const page = await context.newPage();
  await page.goto(`${BASE}/logs`);
  await page.waitForTimeout(FONT_SETTLE_MS);
  const out = `${OUT}logs-built-${name}.png`;
  await page.screenshot({ path: out, fullPage: true });
  await context.close();
  console.log(out);
}

mkdirSync(OUT, { recursive: true });
const names = process.argv.slice(2).length ? process.argv.slice(2) : Object.keys(SHOTS);
const browser = await chromium.launch();
for (const name of names) await shoot(browser, name, SHOTS[name]);
await browser.close();
