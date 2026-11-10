// The served data and copy the rankings and ratings mocks render (ws_rankings.py, ws_ratings.py; directions b1-rankings*):
//   node design/mocks/rankings_facts.mjs [out.json]  ->  design/mocks/data/rankings-facts.json (or out.json)
// Bundles the e2e rankings stubs (web/tests/e2e/support/rankings-fixture.ts: the ranked encounters, the signed-out
// board, the signed-in Fury Warrior board with the character's place, the guild Speed board) and the page's copy
// (web/src/lib/rankings/copy.ts), with the board's own view helpers (the percentile, the rank colour, the date), so
// the mock never retypes a rule the page has.
import { createRequire } from "node:module";
import { mkdtempSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/rankings-facts.json", import.meta.url);

const entry = `
import { boardFor, guildSpeedRows, ONYXIA, KELTHAS, OWN_KEY } from ${JSON.stringify(WEB + "tests/e2e/support/rankings-fixture.ts")};
import { rankingsCopy, rankingsEmptyCopy, encounterPickerCopy, homeTopGuildsCopy } from ${JSON.stringify(WEB + "src/lib/rankings/copy.ts")};
import { rankRowView } from ${JSON.stringify(WEB + "src/lib/rankings/rank-row-view.ts")};
import { guildRowView } from ${JSON.stringify(WEB + "src/lib/rankings/guild-row-view.ts")};
import { topParses, specLeaders } from ${JSON.stringify(WEB + "src/lib/rankings/landing.ts")};
import { RANK_LEGEND, showRowCrests } from ${JSON.stringify(WEB + "src/lib/rankings/row-view.ts")};
import { ordinal } from ${JSON.stringify(WEB + "src/lib/report/format.ts")};
import { rankColor } from ${JSON.stringify(WEB + "src/lib/rankings/row-view.ts")};
import { scoreCopy } from ${JSON.stringify(WEB + "src/lib/score/copy.ts")};
import { ratingCopy } from ${JSON.stringify(WEB + "src/lib/rating/copy.ts")};
import { SAMPLE_SCORE, SAMPLE_TREND, SAMPLE_WEAKEST, WEIGHTS, ratingsPageCopy } from ${JSON.stringify(WEB + "src/lib/score/ratings-page.ts")};
import { COMPONENT_MAX, tierRanges } from ${JSON.stringify(WEB + "src/lib/score/score.ts")};
const strip = (o) => JSON.parse(JSON.stringify(o, (k, v) => (typeof v === 'function' ? '[fn]' : v)));
const board = (q, o) => boardFor(new URL('https://api.test/v1/rankings?' + q), o);
const YOU_Q = 'encounter=onyxia&class=warrior&spec=Fury&ruleset=normal&metric=dps&distinct=character&character=' + OWN_KEY;
const view = (b) => {
  const shape = { total: b.total, top: b.rows[0]?.value ?? 0, unit: 'DPS', multiClass: showRowCrests(b.rows) };
  const rows = b.rows.map((r) => ({ key: r.player.key, ...rankRowView(r, shape) }));
  const pinned = b.character_rank && !b.rows.some((r) => r.player.key === b.character_rank.row.player.key)
    ? rankRowView({ ...b.character_rank.row, rank: b.character_rank.rank }, shape) : null;
  return { total: b.total, per_page: b.per_page, rows, pinned };
};
// Your rank as lib/rankings/own-rank.ts ownRankAnswer reads it, from the same copy and colour helpers. own-rank.ts itself
// reaches the planner's generated reference (through the cohort name), which the parity job's checkout does not build.
const answerOf = (b) => b.character_rank
  ? { kind: 'figures', figures: [{ value: ordinal(b.character_rank.rank), unit: rankingsCopy.answerUnit(b.total, 'Fury Warrior', true),
      color: rankColor(b.character_rank.rank, b.total) }], note: { text: rankingsCopy.answerNote('Onyxia', 'Normal ruleset') } }
  : { kind: 'sentence', text: rankingsCopy.answerUnranked('Onyxia') };
const guilds = guildSpeedRows();
export default {
  encounters: [ONYXIA, KELTHAS],
  board_signed_out: view(board('encounter=warden-kelthas&metric=dps')),
  board_you: view(board(YOU_Q)),
  board_you_off_page: view(board(YOU_Q, { ownOffPage: true })),
  board_you_unranked: view(board(YOU_Q, { unranked: true })),
  answer_you: answerOf(board(YOU_Q)),
  answer_off: answerOf(board(YOU_Q, { ownOffPage: true })),
  answer_unranked: answerOf(board(YOU_Q, { unranked: true })),
  // The landing (RankingsLanding.svelte): each encounter's top three and the spec leaders, from the stub's every-class board
  // (the stub answers every encounter and both metrics with it).
  landing: {
    top: Object.fromEntries([ONYXIA, KELTHAS].map((e) => [e.slug, topParses(board('encounter=' + e.slug + '&metric=dps'))])),
    leaders: specLeaders(board('encounter=onyxia&metric=dps'), board('encounter=onyxia&metric=hps'), (c, s) => '/rankings/onyxia?class=' + c + '&spec=' + s),
  },
  guilds: guilds.map((g) => guildRowView(g, { kind: 'speed', best: guilds[0].value })),
  own_key: OWN_KEY,
  legend: RANK_LEGEND,
  copy: strip({ rankingsCopy, rankingsEmptyCopy, encounterPickerCopy, homeTopGuildsCopy, scoreCopy, ratingCopy }),
  ratings: {
    sample: SAMPLE_SCORE, trend: SAMPLE_TREND, weakest: SAMPLE_WEAKEST, weights: WEIGHTS, copy: ratingsPageCopy,
    component_max: COMPONENT_MAX, tiers: tierRanges(),
    up_to: Object.fromEntries(Object.entries(COMPONENT_MAX).map(([k, v]) => [k, scoreCopy.upTo(v)])),
    most_to_gain_rest: Object.fromEntries(Object.entries(COMPONENT_MAX).map(([k, v]) => [k, scoreCopy.mostToGainRest(v - SAMPLE_SCORE.components[k])])),
  },
  fns: {
    where: Object.fromEntries([ONYXIA, KELTHAS].map((e) => [e.slug, encounterPickerCopy.where(e.zone, e.size)])),
    ranked: Object.fromEntries([ONYXIA, KELTHAS].map((e) => [e.slug, encounterPickerCopy.ranked(e.kills)])),
    boardEmpty: rankingsEmptyCopy.board('Fury Warrior', 'Warden Kelthas'),
    noneYet: encounterPickerCopy.noneYet,
    leadersHeading: encounterPickerCopy.leadersHeading,
    damageHead: encounterPickerCopy.damageHead,
    healingHead: encounterPickerCopy.healingHead,
  },
};`;
const dir = mkdtempSync(join(tmpdir(), "rankings-facts-"));
const outfile = join(dir, "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm",
  outfile, logLevel: "error", external: ["@playwright/test"], define: { "import.meta.env": "{}" } });
const facts = (await import(pathToFileURL(outfile).href)).default;
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
