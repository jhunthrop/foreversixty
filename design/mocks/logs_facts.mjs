// The served data and copy the logs mocks render (design/specs/2026-11-07-logs-reports.md):
//   node design/mocks/logs_facts.mjs  ->  design/mocks/data/logs-facts.json
// Reads, by bundling the TypeScript modules themselves (esbuild from web/node_modules), never by retyping them:
//  - the e2e rankings stubs (web/tests/e2e/support/rankings-fixture.ts): the encounters, the signed-out board
//    (every class), the signed-in Fury Warrior board with the character's place, and the character page, i.e. what
//    a parity spec for /rankings serves;
//  - the page copy modules (lib/reports, lib/report, lib/rankings, lib/score, lib/rating) and the score's tier
//    ranges and part ceilings.
// The report's own figures are read by ws_report.py straight from web/src/fixtures/report and
// web/public/sample-report; the coach review from web/src/fixtures/coach.
import { createRequire } from "node:module";
import { mkdtempSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = new URL("./data/logs-facts.json", import.meta.url);

const entry = `
import { boardFor, characterPage, ONYXIA, KELTHAS, OWN_KEY } from ${JSON.stringify(WEB + "tests/e2e/support/rankings-fixture.ts")};
import { recentReportsCopy, logsBandCopy, logsCompanionCopy, pairingCopy } from ${JSON.stringify(WEB + "src/lib/reports/copy.ts")};
import { myReportsCopy, yourReportsPageCopy } from ${JSON.stringify(WEB + "src/lib/reports/my-reports-copy.ts")};
import { reportCopy } from ${JSON.stringify(WEB + "src/lib/report/copy.ts")};
import { rankingsCopy, rankingsEmptyCopy, encounterPickerCopy } from ${JSON.stringify(WEB + "src/lib/rankings/copy.ts")};
import { scoreCopy } from ${JSON.stringify(WEB + "src/lib/score/copy.ts")};
import { ratingCopy, componentLabel } from ${JSON.stringify(WEB + "src/lib/rating/copy.ts")};
import { COMPONENT_MAX, TIER_FROM, tierRanges } from ${JSON.stringify(WEB + "src/lib/score/score.ts")};
const strip = (o) => JSON.parse(JSON.stringify(o, (k, v) => (typeof v === 'function' ? '[fn]' : v)));
const board = (q, o) => boardFor(new URL('https://api.test/v1/rankings?' + q), o);
const PARTS = ['output', 'survival', 'mechanics', 'utility', 'preparation', 'activity'];
export default {
  encounters: [ONYXIA, KELTHAS],
  board_signed_out: board('encounter=warden-kelthas&metric=dps'),
  board_you: board('encounter=onyxia&class=warrior&spec=Fury&ruleset=normal&metric=dps&distinct=character&character=' + OWN_KEY),
  board_you_off_page: board('encounter=onyxia&class=warrior&spec=Fury&ruleset=normal&metric=dps&distinct=character&character=' + OWN_KEY, { ownOffPage: true }),
  character_page: characterPage('Obnoxious Yell', 'Warrior', ONYXIA.id),
  own_key: OWN_KEY,
  copy: strip({ recentReportsCopy, logsBandCopy, logsCompanionCopy, pairingCopy, myReportsCopy, yourReportsPageCopy, reportCopy,
    rankingsCopy, rankingsEmptyCopy, encounterPickerCopy, scoreCopy, ratingCopy }),
  fns: {
    answerUnit: rankingsCopy.answerUnit(212, 'Fury Warrior', true),
    answerNote: rankingsCopy.answerNote('Onyxia', 'Normal ruleset'),
    answerUnranked: rankingsCopy.answerUnranked('Onyxia'),
    rating: {
      labels: Object.fromEntries(PARTS.map((p) => [p, componentLabel(p)])),
      definitions: Object.fromEntries(PARTS.map((p) => [p, ratingCopy.componentDefinition(p)])),
      excludedMechanics: ratingCopy.excludedReason(componentLabel('mechanics'), 'no_mechanics_table'),
      headingElyra: ratingCopy.headingFor('Elyra Duskvale'),
      percentileBasisHolyPriest: ratingCopy.percentileBasis('Holy', 'Priest', 'healer', 0, 40),
    },
    score: {
      upTo: Object.fromEntries(Object.entries(COMPONENT_MAX).map(([k, v]) => [k, scoreCopy.upTo(v)])),
      mostToGainRest: scoreCopy.mostToGainRest(88),
      samplePill: scoreCopy.samplePill,
      tierShort: Object.fromEntries(Object.entries(TIER_FROM).map(([t, from], i, all) => [t, scoreCopy.tierShort(from, i + 1 < all.length ? all[i + 1][1] - 1 : null)])),
    },
    showing: rankingsCopy.showingNote('Fury Warriors', 'Fury Warriors'),
  },
  score: { component_max: COMPONENT_MAX, tiers: tierRanges() },
};`;
const dir = mkdtempSync(join(tmpdir(), "logs-facts-"));
const outfile = join(dir, "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm",
  outfile, logLevel: "error", external: ["@playwright/test"] });
const facts = (await import(pathToFileURL(outfile).href)).default;
// The stub's dates are relative to now (Date.now() - rank days): pin them to the stub's own updated_at so the
// mock is stable between runs; the date column then reads as the stub intends ("n days ago" from 13 Dec).
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
