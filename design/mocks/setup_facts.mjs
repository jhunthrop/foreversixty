// The served data and copy the setup mocks render (design/specs/2026-11-07-setup.md):
//   node design/mocks/setup_facts.mjs [out.json]  ->  design/mocks/data/setup-facts.json (or out.json)
// Reads, by bundling the TypeScript modules themselves (esbuild from web/node_modules), never by retyping them:
//  - the copy modules: lib/addon/copy.ts, lib/addon/cta-copy.ts, lib/account/connect-copy.ts,
//    lib/account/character-list-copy.ts, lib/reports/copy.ts (pairing, the companion pointer), lib/reports/my-reports-copy.ts;
//  - the page's own line builders evaluated on the fixtures at a pinned clock: lib/setup/progress.ts (signInLine,
//    addonLine, companionLine, companionPlatformsLine), lib/reports/device-status.ts, lib/account/connect-flow.ts;
//  - data: web/src/data/companion.json, active-build.json, lib/downloads/paths.ts;
//  - fixtures: web/tests/e2e/fixtures/home/me-34.json, me-new.json, web/tests/e2e/fixtures/setup/*.json.
// Three strings still live inside components, not a copy module (gap G6): LogsCompanion.svelte's pairing action,
// failure and code lines are read from that file's source here, verbatim.
import { createRequire } from "node:module";
import { mkdtempSync, readFileSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/setup-facts.json", import.meta.url);
const NOW = "2026-11-06T10:00:00Z";
const src = (p) => JSON.stringify(WEB + p);

const entry = `
import { addonCopy } from ${src("src/lib/addon/copy.ts")};
import { ctaCopy } from ${src("src/lib/addon/cta-copy.ts")};
import { connectCopy } from ${src("src/lib/account/connect-copy.ts")};
import { characterListCopy } from ${src("src/lib/account/character-list-copy.ts")};
import { pairingCopy, logsCompanionCopy } from ${src("src/lib/reports/copy.ts")};
import { yourReportsPageCopy } from ${src("src/lib/reports/my-reports-copy.ts")};
import { signInLine, addonLine, companionLine, companionPlatformsLine } from ${src("src/lib/setup/progress.ts")};
import { deviceStatusLine } from ${src("src/lib/reports/device-status.ts")};
import { displayCode } from ${src("src/lib/account/connect-flow.ts")};
import { relativeTime } from ${src("src/lib/dates.ts")};
import { ADDON_DOWNLOAD_PATH, DATA_DOWNLOAD_PATH, companionDownloadPath } from ${src("src/lib/downloads/paths.ts")};
import companion from ${src("src/data/companion.json")};
import activeBuild from ${src("src/data/active-build.json")};
import me34 from ${src("tests/e2e/fixtures/home/me-34.json")};
import meNew from ${src("tests/e2e/fixtures/home/me-new.json")};
import devices from ${src("tests/e2e/fixtures/setup/devices.json")};
import lookup from ${src("tests/e2e/fixtures/setup/connect-device.json")};
import pairCode from ${src("tests/e2e/fixtures/setup/pairing-code.json")};
const NOW = new Date(${JSON.stringify(NOW)});
const strip = (o) => JSON.parse(JSON.stringify(o, (k, v) => (typeof v === 'function' ? '[fn]' : v)));
// Battle.net brought the character, the addon never did (derived from me-34: its build's source set to blizzard).
const meBnet = { ...me34, characters: me34.characters.map((c) => ({ ...c, build: { ...c.build, source: 'blizzard' } })) };
const lines = (me, devs) => ({ signin: signInLine(me), addon: addonLine(me, NOW), companion: companionLine(me, devs, NOW) });
const ch = me34.characters[0];
const code = displayCode('kq47tm');
export default {
  now: NOW.toISOString(),
  me: { addon: me34, bnet: meBnet, none: meNew },
  devices, lookup, pair_code: pairCode,
  lines: { out: lines(null, []), none: lines(meNew, []), bnet: lines(meBnet, []), addon: lines(me34, []), done: lines(me34, devices) },
  device_status: devices.map((d) => deviceStatusLine(d, NOW)),
  imported: characterListCopy.importedFrom(relativeTime(new Date(me34.bnet_imported_at), NOW)),
  platforms_line: companionPlatformsLine(companion.downloads.map((d) => d.platform)),
  companion,
  download: { label: addonCopy.setupCompanionDownload(companion.downloads[0].platform), href: companionDownloadPath(companion.downloads[0].asset) },
  paths: { addon_zip: ADDON_DOWNLOAD_PATH, data: DATA_DOWNLOAD_PATH },
  data_build: addonCopy.currentDataBuild(activeBuild.build),
  addon_arrived: addonCopy.setupAddonArrived(ch.name, relativeTime(new Date(ch.build.captured_at), NOW)),
  code,
  connect: strip(connectCopy),
  connect_fns: {
    codeLabel: connectCopy.codeLabel(code),
    pendingTitle: connectCopy.pendingTitle(lookup.name, me34.user.battletag),
    signedOutLeadNamed: connectCopy.signedOutLeadNamed(lookup.name),
  },
  addon: strip(addonCopy), cta: ctaCopy, characters: strip(characterListCopy),
  pairing: { ...strip(pairingCopy), success_line: pairingCopy.success(devices[0].name) }, logs_companion: logsCompanionCopy, session_failed: yourReportsPageCopy.sessionFailed,
  wrong_prefix: addonCopy.wrongPrefix('FS2', 'FS1'),
};`;
const dir = mkdtempSync(join(tmpdir(), "setup-facts-"));
const outfile = join(dir, "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm",
  outfile, logLevel: "error", define: { "import.meta.env": "{}" }, loader: { ".json": "json" } });
const facts = (await import(pathToFileURL(outfile).href)).default;
// LogsCompanion.svelte keeps three strings of its own (gap G6): read verbatim from the source.
const lc = readFileSync(join(WEB, "src/components/logs/LogsCompanion.svelte"), "utf8");
const grab = (re) => (re.exec(lc) ?? [])[1] ?? null;
const minutes = Math.round(facts.pair_code.expires_in / 60);
facts.pairing_component = {
  action: grab(/PAIR_ACTION = '([^']+)'/),
  failed: grab(/PAIR_FAILED = '([^']+)'/),
  code_line: (grab(/(Type this into the companion within[\s\S]*?cannot be reused\.)/) ?? "")
    .replace(/\{Math\.round\([^}]+\)\}/, String(minutes)).replace(/\s+/g, " "),
};
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
