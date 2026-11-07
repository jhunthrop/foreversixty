// The served data and copy the planner mocks render (design/specs/2026-11-08-planner.md):
//   node design/mocks/planner_facts.mjs [out.json]  ->  design/mocks/data/planner-facts.json (or out.json)
// Runs the planner's own models, bundled with esbuild from web/node_modules, never retyped:
//  - the trees: talents/<class>.json of the active build (data/builds, the committed copy), indexTalents, gridCells,
//    gridSize, connectorsFor, canAddPoint, cellState (lib/planner/*), the spec icons of lib/sim/specs.ts;
//  - a build: decodeFS1 + orderFromRanks (a guide's own code, as /planner?code= applies it), loadFromBandAtLevel (a
//    band build cut at a character's level, as BandCompare's Load does), levelForIndex, levelReached, pointsAvailableAt;
//  - the compare: the band entry of data/builds/<build>/bis/<spec>.json through bandEntryFor, diffAgainstBand,
//    bandDiffStatus, bandForPanel, bandForLevel, bandRangeLabel;
//  - the copy: lib/planner/copy.ts (plannerCopy, plannerHeaderCopy, plannerCharacterCardCopy, bandCompareCopy),
//    lib/addon/copy.ts (import and Copy for the game), lib/addon/cta-copy.ts, lib/planner/rules.ts `messages`,
//    lib/planner/score-unit.ts, lib/sim/planner-copy.ts, lib/planner/share.ts, lib/planner/tree-source.ts,
//    lib/planner/load.ts DATA_LOAD_FAILED, lib/planner/send-state.ts sendActionFor (every send state);
//  - the signed-in character: web/tests/e2e/fixtures/home/me-34.json (Felsnap, a level 34 Dwarf Beast Mastery hunter).
// The live estimate (PlannerDps) runs the engine in the browser; the mocks carry SAMPLE figures (marked in each mock).
import { createRequire } from "node:module";
import { existsSync, mkdtempSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { basename, join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const REPO = new URL("../../", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/planner-facts.json", import.meta.url);
const src = (p) => JSON.stringify(WEB + p);
const { build } = require(join(WEB, "src/data/active-build.json"));
const data = (p) => JSON.stringify(`${REPO}data/builds/${build}/${p}`);
const CLASSES = ["warrior", "hunter", "priest", "mage"];
const BIS = ["warrior-arms", "hunter-survival", "hunter-beast-mastery", "priest-holy", "mage-fire"];

const entry = `
${CLASSES.map((c) => `import t_${c} from ${data(`talents/${c}.json`)};`).join("\n")}
${BIS.map((s) => `import b_${s.replace(/-/g, "_")} from ${data(`bis/${s}.json`)};`).join("\n")}
import classes from ${data("classes.json")};
import races from ${data("races.json")};
import combos from ${data("combos.json")};
import me34 from ${src("tests/e2e/fixtures/home/me-34.json")};
import { decodeFS1, orderFromRanks } from ${src("src/lib/planner/fs1.ts")};
import { indexTalents, canAddPoint, messages } from ${src("src/lib/planner/rules.ts")};
import { gridCells, gridSize } from ${src("src/lib/planner/grid.ts")};
import { connectorsFor, gridViewBox } from ${src("src/lib/planner/connectors.ts")};
import { cellState } from ${src("src/lib/planner/styles.ts")};
import { levelForIndex, levelReached, pointsAvailableAt, ranksByTalent, pointsPerTree } from ${src("src/lib/planner/derive.ts")};
import { diffAgainstBand, bandDiffStatus, bandForPanel, loadFromBandAtLevel } from ${src("src/lib/planner/band-compare.ts")};
import { bandForLevel, bandEntryFor } from ${src("src/lib/bis/hover.ts")};
import { bandRangeLabel } from ${src("src/lib/planner/band-title.ts")};
import { plannerCopy, plannerHeaderCopy, plannerCharacterCardCopy, bandCompareCopy } from ${src("src/lib/planner/copy.ts")};
import { addonCopy } from ${src("src/lib/addon/copy.ts")};
import { ctaCopy } from ${src("src/lib/addon/cta-copy.ts")};
import { plannerScoreCopy, scoreUnitForBuild } from ${src("src/lib/planner/score-unit.ts")};
import { plannerSimCopy } from ${src("src/lib/sim/planner-copy.ts")};
import { RATE_LIMIT_MESSAGE, SAVE_FAILED_MESSAGE, sharedFields } from ${src("src/lib/planner/share.ts")};
import { treeSourceNotice } from ${src("src/lib/planner/tree-source.ts")};
import { DATA_LOAD_FAILED } from ${src("src/lib/planner/load.ts")};
import { READ_ONLY_REASON } from ${src("src/lib/planner/store.svelte.ts")};
import { sendActionFor } from ${src("src/lib/planner/send-state.ts")};
import { specKeyFor } from ${src("src/lib/addon/score.ts")};
import { SPECS } from ${src("src/lib/sim/specs.ts")};
import { specLabel, specDisplayName, siblingSpecNames, classDisplayName } from ${src("src/lib/sim/spec-label.ts")};

const BUILD = ${JSON.stringify(build)};
const TALENTS = { ${CLASSES.map((c) => `${c}: t_${c}`).join(", ")} };
const BISF = { ${BIS.map((s) => `'${s}': b_${s.replace(/-/g, "_")}`).join(", ")} };
const strip = (o) => JSON.parse(JSON.stringify(o, (k, v) => (typeof v === 'function' ? '[fn]' : v)));
const GUIDE = {
  survival: 'FS1:1.60.1.70790:hunter:night-elf:5/005005201/500230230050222151:',
  holy: 'FS1:1.60.1.70790:priest:human:00500503103/33505003030121431/0:',
  fire: 'FS1:1.60.1.70790:mage:human:0501052001000/55500201230133031/003:',
};
const classRow = (slug) => classes.find((c) => c.slug === slug);
const legalRaces = (slug) => { const id = classRow(slug).id; return races.filter((r) => !r.placeholder && combos.some((c) => c.class_id === id && c.race_id === r.id)); };
const raceRow = (slug) => races.find((r) => r.slug === slug);

// One planner state as the page draws it. order: talent ids in spend order; level: the header character's level or null.
function state({ cls, race, order = [], level = null, spec = null }) {
  const index = indexTalents(TALENTS[cls]);
  const ranks = ranksByTalent(order);
  const split = pointsPerTree(index, order);
  const spent = order.length;
  const available = level === null ? 51 : Math.max(pointsAvailableAt(level), spent);
  const shownLevel = level ?? levelReached(order);
  const leads = Math.max(...split) > 0 && split.filter((p) => p === Math.max(...split)).length === 1;
  const specKey = spec ?? specKeyFor(cls, split);
  const faction = raceRow(race).faction === 'horde' ? 'horde' : 'alliance';
  const band = bandForPanel(spent, level, bandForLevel(shownLevel));
  const file = BISF[specKey];
  const entry = file ? bandEntryFor(file, band, faction) : undefined;
  const bandLabel = bandRangeLabel(band);
  let compare = null;
  if (entry) {
    const diff = diffAgainstBand(index, ranks, entry.talents);
    const status = bandDiffStatus(diff);
    compare = { status, differs: [...diff.diffTalentIds], bandRank: Object.fromEntries(diff.bandRankById),
      line: status.kind === 'matches' ? bandCompareCopy.matches : status.kind === 'differs' ? bandCompareCopy.differsSuffix : bandCompareCopy.toTakeSuffix,
      beyond: status.kind === 'toTake' ? '' : (status.beyond > 0 ? bandCompareCopy.beyondGuide(status.beyond) : ''),
      set_dps: plannerScoreCopy.bandFigure(entry.set_dps, entry.score_unit ?? 'dps'), set_caption: bandCompareCopy.setDpsCaption(bandLabel),
      talents: entry.talents };
  }
  const trees = index.trees.map((tree, ti) => {
    const size = gridSize(tree);
    const specRow = SPECS.find((s) => s.class_slug === cls && s.tree_index === ti);
    return { name: tree.name, icon: specRow?.icon ?? tree.icon, background: tree.background, points: split[ti] ?? 0, tiers: size.tiers, columns: size.columns,
      viewBox: gridViewBox(size.tiers, size.columns),
      connectors: connectorsFor(tree, ranks).map((c) => ({ d: c.d, met: c.met })),
      cells: gridCells(tree).map((g) => { const t = g.talent; const rank = ranks.get(t.id) ?? 0;
        const ok = canAddPoint(index, order, t.id).ok;
        return { id: t.id, tier: g.tier, column: g.column, name: t.name, icon: t.icon, rank, max: t.max_rank, state: cellState(rank, t.max_rank, ok),
          now: rank > 0 ? t.ranks[rank - 1].description : '', next: rank < t.max_rank ? t.ranks[rank].description : '',
          refusal: ok ? '' : canAddPoint(index, order, t.id).reason }; }) };
  });
  const points = order.map((id, i) => { const t = index.byId.get(id); const r = order.slice(0, i + 1).filter((x) => x === id).length;
    return { level: levelForIndex(i), name: t.name, icon: t.icon, rank: r, max: t.max_rank }; });
  const unit = scoreUnitForBuild(cls, split);
  return {
    cls, className: classRow(cls).name, title: plannerHeaderCopy.h1(classRow(cls).name), race: raceRow(race).name, race_slug: race, faction,
    races: legalRaces(cls).map((r) => r.name), level: shownLevel, char_level: level, spent, available, left: Math.max(0, available - spent), split,
    spec: leads ? specLabel(specKey) : '', spec_key: specKey, spec_name: specDisplayName(specKey), siblings: siblingSpecNames(specKey),
    meta: [raceRow(race).name, 'Level ' + shownLevel, leads ? specLabel(specKey) : ''].filter(Boolean).join(' · '),
    trees, points, band: { band, label: bandLabel, heading: bandCompareCopy.heading(bandLabel), load: bandCompareCopy.loadButton(bandLabel),
      gear: bandCompareCopy.gearRowLabel(bandLabel), none: entry ? '' : bandCompareCopy.noLevelBuild(shownLevel, specDisplayName(specKey), siblingSpecNames(specKey)) },
    compare, unit, unit_label: plannerScoreCopy.label(unit), unit_show: plannerScoreCopy.show(unit), unit_failed: plannerScoreCopy.failed(unit),
    text_summary: plannerCopy.shareTextSummary(raceRow(race).name, classRow(cls).name, trees.map((t) => ({ name: t.name, points: t.points })), shownLevel),
    point_meta: plannerCopy.pointOrderMeta(spent),
    share_fields: sharedFields({ title: '', gear: {} }, false),
  };
}
const guideOrder = (code) => { const d = decodeFS1(code); const idx = indexTalents(TALENTS[d.build.classSlug]); return orderFromRanks(idx, d.build.treeRanks); };
const surv = guideOrder(GUIDE.survival);
const felsnap = me34.characters[0];
const bmFile = BISF['hunter-beast-mastery'];
const bm30 = bandEntryFor(bmFile, bandForLevel(felsnap.level), 'alliance');
const felsnapOrder = loadFromBandAtLevel(indexTalents(TALENTS.hunter), bm30.talents, felsnap.level);
// At the level cap: the next band's build cut at the character's level, as BandCompare's Load does for a level 34 on the 40 band.
const capOrder = loadFromBandAtLevel(indexTalents(TALENTS.hunter), bandEntryFor(bmFile, 40, 'alliance').talents, felsnap.level);
const invalid = decodeFS1('FS9:nope');
const CH = plannerCharacterCardCopy;
const send = (st, spent = 25, extra = {}) => sendActionFor(felsnap, st, spent, 23, { classSlug: 'hunter', className: 'Hunter', simHref: '/sim', ...extra });

export default {
  build: BUILD,
  classes: classes.map((c) => ({ slug: c.slug, name: c.name })),
  states: {
    empty: { ...state({ cls: 'warrior', race: legalRaces('warrior')[0].slug }), note: null },
    progress: { ...state({ cls: 'hunter', race: 'night-elf', order: surv.order.slice(0, 20) }), note: null },
    guide: { ...state({ cls: 'hunter', race: 'night-elf', order: surv.order }), dropped: surv.dropped.length, origin: CH.originLine('', 'Survival Hunter') },
    holy: { ...state({ cls: 'priest', race: 'human', order: guideOrder(GUIDE.holy).order }), origin: CH.originLine('', 'Holy Priest') },
    fire: { ...state({ cls: 'mage', race: 'human', order: guideOrder(GUIDE.fire).order }), origin: CH.originLine('', 'Fire Mage') },
    you: state({ cls: 'hunter', race: 'dwarf', order: felsnapOrder.order, level: felsnap.level, spec: 'hunter-beast-mastery' }),
    you_empty: state({ cls: 'hunter', race: 'dwarf', order: [], level: felsnap.level, spec: 'hunter-beast-mastery' }),
    you_cap: state({ cls: 'hunter', race: 'dwarf', order: capOrder.order, level: felsnap.level, spec: 'hunter-beast-mastery' }),
    invalid: { ...state({ cls: 'warrior', race: legalRaces('warrior')[0].slug }), note: invalid.ok ? '' : invalid.message },
  },
  felsnap: { name: felsnap.name, level: felsnap.level, race: felsnap.race, spec: felsnap.spec, class: felsnap.class, loaded_line: bandCompareCopy.loadedLine(felsnap.level, felsnapOrder.order.length, felsnapOrder.fullPoints) },
  send: {
    idle: send({ kind: 'idle' }), empty: send({ kind: 'idle' }, 0), sending: send({ kind: 'sending' }), sent: send({ kind: 'sent' }),
    signin: send({ kind: 'sign-in' }), nodevice: send({ kind: 'no-device' }), failed: send({ kind: 'failed' }), refused: send({ kind: 'failed', reason: 'talent 104960 is not in this class.' }),
    other_class: send({ kind: 'idle' }, 51, { classSlug: 'priest', className: 'Priest' }), others: send({ kind: 'idle' }, 51, { owner: '' }),
    confirming: send({ kind: 'confirming' }, 51, { owner: '' }), signed_out: sendActionFor(null, { kind: 'idle' }),
  },
  notes: {
    reconstructed: 'Talents loaded from a character. The order points were spent in is not recorded in game, so this is one way to level into the same tree: the main tree first, each talent as soon as its row opens.',
    band60: plannerCopy.bandLoadedNote('60'), kept: plannerCopy.keptNote(23, 25, 'Felsnap'), gear_only: plannerCopy.noTalentRanksNote,
    no_class: 'That code names a class this planner does not have: monk.', invalid: invalid.ok ? '' : invalid.message,
  },
  no_level_example: bandCompareCopy.noLevelBuild(9, specDisplayName('warrior-arms'), siblingSpecNames('warrior-arms')),
  band_tip: bandCompareCopy.bandRankTooltipLine('60', 5, 5),
  refusals: { tierLocked: messages.tierLocked(10, 'Survival', 5), prereq: messages.prereqMissing('Lacerate', 5, 'Savage Strikes'), maxRank: messages.maxRank('Deadly Aspects', 5), cap: messages.capReached(), noPoints: messages.noPoints('Deadly Aspects') },
  copy: { planner: strip(plannerCopy), header: strip(plannerHeaderCopy), card: strip(CH), band: strip(bandCompareCopy), score: strip(plannerScoreCopy), sim: strip(plannerSimCopy),
    addon: { copyForGame: addonCopy.copyForGame, copiedForGame: addonCopy.copiedForGame, copyFailed: addonCopy.copyFailed, buildCodeLabel: addonCopy.buildCodeLabel, addonCodeHint: addonCopy.addonCodeHint,
      importTitle: addonCopy.importTitle, importPlaceholder: addonCopy.importPlaceholder, importAction: addonCopy.importAction, importNoExport: addonCopy.importNoExport, importGetAddon: addonCopy.importGetAddon,
      importOrderApproximated: addonCopy.importOrderApproximated, importSwitching: addonCopy.importSwitchingClass('Priest'), importOlder: addonCopy.importOlderBuild('1.60.1.70630', BUILD), importDropped: addonCopy.importDropped('Lacerate') },
    cta: ctaCopy, rate_limit: RATE_LIMIT_MESSAGE, save_failed: SAVE_FAILED_MESSAGE, tree_source: treeSourceNotice(BUILD), load_failed: DATA_LOAD_FAILED,
    load_failed_detail: 'Build ' + BUILD + ' did not return the files the planner needs.', read_only: READ_ONLY_REASON,
    sim_done: plannerSimCopy.buildSimDone('612', 'f048630de'), points_to_go: plannerSimCopy.plannerDpsPointsToGo(31) },
};`;
// CI's parity job serves a downloaded dist and never runs sync:data, so src/data/generated is absent there;
// the spec catalogue reads it. Fall back to the fixture the fixture build publishes (setup_facts.mjs's plugin).
const generatedFallback = {
  name: "generated-fallback",
  setup(b) {
    b.onResolve({ filter: /data\/generated\/[^/]+\.json$/ }, (args) => {
      const wanted = join(args.resolveDir, args.path);
      if (existsSync(wanted)) return { path: wanted };
      return { path: join(WEB, "src/fixtures/planner", basename(args.path)) };
    });
  },
};
const dir = mkdtempSync(join(tmpdir(), "planner-facts-"));
const outfile = join(dir, "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm",
  outfile, logLevel: "error", define: { "import.meta.env": "{}" }, loader: { ".json": "json" }, plugins: [generatedFallback] });
const facts = (await import(pathToFileURL(outfile).href)).default;
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
