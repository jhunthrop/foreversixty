// The served data and copy the planner mocks render (design/specs/2026-11-08-planner.md):
//   node design/mocks/planner_facts.mjs [out.json]  ->  design/mocks/data/planner-facts.json (or out.json)
// Runs the planner's own models, bundled with esbuild from web/node_modules, never retyped:
//  - the trees: talents/<class>.json of the active build (data/builds, the committed copy), indexTalents, gridCells,
//    gridSize, connectorsFor, canAddPoint (lib/planner/*), cellState (lib/ui/talent-tree), the spec icons of lib/sim/specs.ts;
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
import { existsSync, mkdtempSync, readdirSync, readFileSync, writeFileSync } from "node:fs";
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
// Every spec guide's published builds (the guides' frontmatter `build` and `raidBuild`), as planner.astro reads them:
// the rail's Guide builds and, for a character, the guide's order past its build (lib/planner/guide-builds.ts, point-order.ts).
const GUIDES_DIR = join(WEB, "src/content/guides");
const front = (text, key) => text.match(new RegExp(`^${key}: '([^']+)'`, "m"))?.[1];
const GUIDE_BUILDS = readdirSync(GUIDES_DIR, { withFileTypes: true }).filter((d) => d.isDirectory()).map((d) => d.name).flatMap((classSlug) =>
  readdirSync(join(GUIDES_DIR, classSlug)).filter((f) => f.endsWith(".md") && f !== "index.md").flatMap((f) => {
    const text = readFileSync(join(GUIDES_DIR, classSlug, f), "utf8");
    const specSlug = f.replace(/\.md$/, "");
    return [["leveling", front(text, "build")], ["raid", front(text, "raidBuild")]]
      .filter(([, code]) => code !== undefined).map(([kind, code]) => ({ classSlug, specSlug, kind, code }));
  }));

const entry = `
${CLASSES.map((c) => `import t_${c} from ${data(`talents/${c}.json`)};`).join("\n")}
${BIS.map((s) => `import b_${s.replace(/-/g, "_")} from ${data(`bis/${s}.json`)};`).join("\n")}
import classes from ${data("classes.json")};
import races from ${data("races.json")};
import combos from ${data("combos.json")};
import me34 from ${src("tests/e2e/fixtures/home/me-34.json")};
import { deltaText, guideBuildRows, guideDiff, ranksTrees } from ${src("src/lib/planner/guide-builds.ts")};
import { aheadOf, pointOrderRows } from ${src("src/lib/planner/point-order.ts")};
import { decodeFS1, orderFromRanks } from ${src("src/lib/planner/fs1.ts")};
import { indexTalents, canAddPoint, messages } from ${src("src/lib/planner/rules.ts")};
import { gridCells, gridSize } from ${src("src/lib/planner/grid.ts")};
import { connectorsFor, gridViewBox } from ${src("src/lib/planner/connectors.ts")};
import { cellState } from ${src("src/lib/ui/talent-tree.ts")};
import { levelForIndex, levelReached, pointsAvailableAt, ranksByTalent, pointsPerTree } from ${src("src/lib/planner/derive.ts")};
import { diffAgainstBand, bandDiffStatus, bandForPanel, loadFromBandAtLevel, bandDiffNames, bandPointTotal, bandCoverage } from ${src("src/lib/planner/band-compare.ts")};
import { ladderPrefixFor } from ${src("src/lib/planner/ladder-load.ts")};
import { guideLadderOrder } from ${src("src/lib/guides/leveling-builds.ts")};
import { bandForLevel, bandEntryFor } from ${src("src/lib/bis/hover.ts")};
import { bandRangeLabel } from ${src("src/lib/planner/band-title.ts")};
import { plannerCopy, plannerHeaderCopy, plannerCharacterCardCopy, bandCompareCopy } from ${src("src/lib/planner/copy.ts")};
import { addonCopy } from ${src("src/lib/addon/copy.ts")};
import { ctaCopy } from ${src("src/lib/addon/cta-copy.ts")};
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
const GUIDE_BUILDS = ${JSON.stringify(GUIDE_BUILDS)};
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
// Build vs guide for one guide row (GuideBuilds.svelte): empty on the build on screen or with nothing spent.
function guideRowDiff(trees, currentTrees, index) {
  if (currentTrees === null || trees === currentTrees) return { off: '', delta: '', missing: '' };
  const names = index.trees.map((tree) => tree.talents.map((talent) => talent.name));
  const diff = guideDiff(trees, currentTrees, names);
  return {
    off: plannerCopy.guideBuildOff(diff.off),
    delta: deltaText(diff.delta),
    missing: diff.missing === null ? '' : plannerCopy.guideBuildMissing(diff.missing),
  };
}

function state({ cls, race, order = [], level = null, spec = null }) {
  const index = indexTalents(TALENTS[cls]);
  const ranks = ranksByTalent(order);
  const split = pointsPerTree(index, order);
  const spent = order.length;
  // Bug B2: a character's points stop at its level; points past it are planned ahead (P4), never spendable.
  const available = level === null ? 51 : pointsAvailableAt(level);
  const planned = level === null ? 0 : Math.max(0, spent - available);
  const atCap = spent >= available;
  const shownLevel = level ?? levelReached(order);
  // The head answer (PlannerAnswer): the points spent of those there are, and any planned past the level.
  const pointsUnit = plannerCopy.pointsOfUnit(available);
  const plannedLine = planned > 0 ? plannerCopy.plannedAhead(planned) : '';
  const plannedRanks = new Map();
  order.forEach((id, i) => { if (level !== null && i >= available) plannedRanks.set(id, (plannedRanks.get(id) ?? 0) + 1); });
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
    const names = bandDiffNames(index, ranks, entry.talents);
    const covers = bandCoverage(status, bandPointTotal(index, entry.talents), level === null ? null : available);
    const line = covers !== null ? bandCompareCopy.coversFewer(covers.covered, covers.total)
      : status.kind === 'matches' ? bandCompareCopy.matches(specDisplayName(specKey))
      : status.kind === 'differs' ? bandCompareCopy.differsNamed(status.count, names.mine, names.theirs, bandLabel)
      : bandCompareCopy.toTakeSuffix;
    // kind: matches (green word) | sentence (a whole sentence: covers or a named difference) | count (gold count, then the line)
    const kind = covers !== null || status.kind === 'differs' ? 'sentence' : status.kind === 'matches' ? 'matches' : 'count';
    compare = { status, kind, differs: [...diff.diffTalentIds], bandRank: Object.fromEntries(diff.bandRankById), line,
      beyond: status.kind === 'toTake' || covers !== null ? '' : (status.beyond > 0 ? bandCompareCopy.beyondGuide(status.beyond) : ''),
      talents: entry.talents };
  }
  const trees = index.trees.map((tree, ti) => {
    const size = gridSize(tree);
    const specRow = SPECS.find((s) => s.class_slug === cls && s.tree_index === ti);
    return { name: tree.name, icon: specRow?.icon ?? tree.icon, background: tree.background, points: split[ti] ?? 0, tiers: size.tiers, columns: size.columns,
      viewBox: gridViewBox(size.tiers, size.columns),
      connectors: connectorsFor(tree, ranks).map((c) => ({ d: c.d, met: c.met })),
      cells: gridCells(tree).map((g) => { const t = g.talent; const rank = ranks.get(t.id) ?? 0;
        const check = canAddPoint(index, order, t.id);
        // P3: with no points left at the character's level the unspent cells read locked, and say why.
        const capped = check.ok && level !== null && atCap;
        return { id: t.id, tier: g.tier, column: g.column, name: t.name, icon: t.icon, rank, max: t.max_rank, state: cellState(rank, t.max_rank, check.ok && !capped),
          planned: plannedRanks.get(t.id) ?? 0,
          now: rank > 0 ? t.ranks[rank - 1].description : '', next: rank < t.max_rank ? t.ranks[rank].description : '',
          refusal: capped ? plannerCopy.noPointsLeft(level) : check.ok ? '' : check.reason }; }) };
  });
  // Point order as PlannerOrder draws it: for a character, the spec guide's order past its build (aheadOf).
  const specRow = SPECS.find((s) => s.spec === specKey);
  const guideBuild = GUIDE_BUILDS.find((g) => g.kind === 'leveling' && g.classSlug + '-' + g.specSlug === specKey);
  // An empty build previews the spec guide's order (PlannerOrder.svelte, h-audit 2026-11-09).
  const ladder = (level !== null || order.length === 0) && guideBuild && specRow ? guideLadderOrder(BUILD, cls, specKey, specRow.tree_index, guideBuild.code) : [];
  const aheadPoints = ladder.length > 0 ? aheadOf(index, order, ladder) : [];
  const previewLine = order.length === 0 && aheadPoints.length > 0 ? plannerCopy.previewLine(specRow.name) : '';
  const orderView = pointOrderRows(index, order, level, me34.characters[0].name, (icon) => icon,
    aheadPoints.length > 0 ? { specName: specRow.name, points: aheadPoints } : null);
  const currentTrees = order.length > 0 ? ranksTrees(index.trees.map((tree) => tree.talents.map((t) => ranks.get(t.id) ?? 0))) : null;
  const guideBuilds = guideBuildRows(GUIDE_BUILDS, cls).map((row) => ({
    ...row,
    current: row.trees === currentTrees,
    // Build vs guide (GuideBuilds.svelte), once a point is spent: points off, the split's change, the first talent lacking.
    ...guideRowDiff(row.trees, currentTrees, index),
  }));
  const points = order.map((id, i) => { const t = index.byId.get(id); const r = order.slice(0, i + 1).filter((x) => x === id).length;
    return { level: levelForIndex(i), name: t.name, icon: t.icon, rank: r, max: t.max_rank, planned: level !== null && i >= available, tip: plannerCopy.pointAt(levelForIndex(i), t.name, r) }; });
  return {
    cls, className: classRow(cls).name, title: plannerHeaderCopy.h1(classRow(cls).name), race: raceRow(race).name, race_slug: race, faction,
    races: legalRaces(cls).map((r) => r.name), level: shownLevel, char_level: level, spent, available, left: Math.max(0, available - spent), planned, points_unit: pointsUnit, planned_line: plannedLine, split,
    planned_cell: level === null ? '' : plannerCopy.plannedCell(level),
    order_rows: [10, 20, 30, 40, 50, 60].map((lo) => plannerCopy.decade(lo)),
    next_point: spent < 51 && (level === null || spent < available) ? plannerCopy.nextPoint(levelForIndex(spent)) : '',
    level_line: level !== null && spent >= available ? plannerCopy.levelLine(me34.characters[0].name, level) : '',
    spec: leads ? specLabel(specKey) : '', spec_key: specKey, spec_name: specDisplayName(specKey), siblings: siblingSpecNames(specKey),
    meta: [raceRow(race).name, 'Level ' + shownLevel, leads ? specLabel(specKey) : ''].filter(Boolean).join(' · '),
    trees, points, band: { band, label: bandLabel, heading: bandCompareCopy.heading(bandLabel), load: bandCompareCopy.loadButton(bandLabel),
      none: entry ? '' : bandCompareCopy.noLevelBuild(shownLevel, specDisplayName(specKey), siblingSpecNames(specKey)) },
    compare,
    text_summary: plannerCopy.shareTextSummary(raceRow(race).name, classRow(cls).name, trees.map((t) => ({ name: t.name, points: t.points })), shownLevel),
    order_view: orderView, preview_line: previewLine, guide_builds: guideBuilds,
    share_fields: sharedFields({ title: '', gear: {} }, false),
  };
}
const guideOrder = (code) => { const d = decodeFS1(code); const idx = indexTalents(TALENTS[d.build.classSlug]); return orderFromRanks(idx, d.build.treeRanks); };
// The Survival guide's own levelling order (the guide's Leveling strip, /data/leveling-plans.json's ladder): a guide's
// link that names its spec loads the build in it, when the ladder reaches exactly the guide's ranks (ladderPrefixFor).
const survIndex = indexTalents(TALENTS.hunter);
const survLadder = guideLadderOrder(BUILD, 'hunter', 'hunter-survival', 2, GUIDE.survival);
const survPrefix = ladderPrefixFor(survIndex, decodeFS1(GUIDE.survival).build.treeRanks, survLadder);
const surv = survPrefix !== undefined ? { order: survPrefix, dropped: [] } : guideOrder(GUIDE.survival);
const felsnap = me34.characters[0];
const bmFile = BISF['hunter-beast-mastery'];
const bm30 = bandEntryFor(bmFile, bandForLevel(felsnap.level), 'alliance');
const felsnapOrder = loadFromBandAtLevel(indexTalents(TALENTS.hunter), bm30.talents, felsnap.level);
// At the level cap, then two clicks past it (P4): the next band's build cut at level 36, so points 26 and 27 are planned ahead.
const capOrder = loadFromBandAtLevel(indexTalents(TALENTS.hunter), bandEntryFor(bmFile, 40, 'alliance').talents, felsnap.level + 2);
const invalid = decodeFS1('FS9:nope');
// Every refusal line the states sheet shows, from canAddPoint on real builds and the hunter's own talents (ruling MUST 2).
function refusals() {
  const progress = surv.order.slice(0, 20);
  const survival = survIndex.trees[2];
  const firstRefusal = (order, pred) => { for (const tree of [survival]) for (const t of tree.talents) { const c = canAddPoint(survIndex, order, t.id); if (!c.ok && pred(c.reason)) return c.reason; } return ''; };
  const withPrereq = survival.talents.find((t) => t.prereq_talent_id !== null && survIndex.byId.get(t.prereq_talent_id)?.max_rank > 1);
  const maxed = ranksByTalent(surv.order);
  const maxedTalent = [...maxed].map(([id, r]) => survIndex.byId.get(id)).find((t) => maxed.get(t.id) === t.max_rank && t.max_rank > 1);
  const unspent = survIndex.trees.flatMap((tree) => tree.talents).find((t) => !maxed.has(t.id));
  return {
    tierLocked: firstRefusal(progress, (r) => r.startsWith('Needs ')),
    prereq: messages.prereqMissing(withPrereq.name, withPrereq.prereq_rank, survIndex.byId.get(withPrereq.prereq_talent_id).name),
    maxRank: messages.maxRank(maxedTalent.name, maxedTalent.max_rank),
    cap: messages.capReached(),
    noPoints: messages.noPoints(unspent.name),
    noPointsLeft: plannerCopy.noPointsLeft(felsnap.level),
  };
}
const CH = plannerCharacterCardCopy;
const send = (st, spent = 25, extra = {}) => sendActionFor(felsnap, st, spent, 23, { classSlug: 'hunter', className: 'Hunter', simHref: '/sim', ...extra });

export default {
  build: BUILD,
  classes: classes.map((c) => ({ slug: c.slug, name: c.name })),
  states: {
    empty: { ...state({ cls: 'warrior', race: legalRaces('warrior')[0].slug }), note: null },
    progress: { ...state({ cls: 'hunter', race: 'night-elf', order: surv.order.slice(0, 20) }), note: null },
    // A class chosen in the bar from the bare Warrior page: selectClass keeps the race when it is legal for the class.
    hunter_empty: state({ cls: 'hunter', race: legalRaces('hunter').some((r) => r.slug === legalRaces('warrior')[0].slug) ? legalRaces('warrior')[0].slug : legalRaces('hunter')[0].slug }),
    guide: { ...state({ cls: 'hunter', race: 'night-elf', order: surv.order }), dropped: surv.dropped.length, origin: CH.originLine('', 'Survival Hunter') },
    holy: { ...state({ cls: 'priest', race: 'human', order: guideOrder(GUIDE.holy).order }), origin: CH.originLine('', 'Holy Priest') },
    fire: { ...state({ cls: 'mage', race: 'human', order: guideOrder(GUIDE.fire).order }), origin: CH.originLine('', 'Fire Mage') },
    you: state({ cls: 'hunter', race: 'dwarf', order: felsnapOrder.order, level: felsnap.level, spec: 'hunter-beast-mastery' }),
    you_empty: state({ cls: 'hunter', race: 'dwarf', order: [], level: felsnap.level, spec: 'hunter-beast-mastery' }),
    you_cap: state({ cls: 'hunter', race: 'dwarf', order: capOrder.order, level: felsnap.level, spec: 'hunter-beast-mastery' }),
    invalid: { ...state({ cls: 'warrior', race: legalRaces('warrior')[0].slug }), note: invalid.ok ? '' : invalid.message },
  },
  // The committed signed-in fixture with real ranks (gap G5, web/tests/e2e/fixtures/planner/me-34-talents.json): Felsnap's
  // own 25 points (the 40 band cut at 34), one digit per talent in tab order, and the two points the cap mock plans past them.
  felsnap_cap: { trees: indexTalents(TALENTS.hunter).trees.map((tree) => { const r = ranksByTalent(capOrder.order.slice(0, 25)); return tree.talents.map((t) => (r.get(t.id) ?? 0).toString(36)).join(''); }),
    points: pointsPerTree(indexTalents(TALENTS.hunter), capOrder.order.slice(0, 25)),
    planned: capOrder.order.slice(25).map((id) => indexTalents(TALENTS.hunter).byId.get(id).name) },
  felsnap: { name: felsnap.name, level: felsnap.level, race: felsnap.race, spec: felsnap.spec, class: felsnap.class, loaded_line: bandCompareCopy.loadedLine(felsnap.level, felsnapOrder.order.length, felsnapOrder.fullPoints) },
  send: {
    idle: send({ kind: 'idle' }), empty: send({ kind: 'idle' }, 0), sending: send({ kind: 'sending' }), sent: send({ kind: 'sent' }),
    signin: send({ kind: 'sign-in' }), nodevice: send({ kind: 'no-device' }), failed: send({ kind: 'failed' }), refused: send({ kind: 'failed', reason: 'talent 104960 is not in this class.' }),
    other_class: { ...send({ kind: 'idle' }, 51, { classSlug: 'priest', className: 'Priest' }), otherClass: true }, others: send({ kind: 'idle' }, 51, { owner: '' }),
    confirming: send({ kind: 'confirming' }, 51, { owner: '' }), signed_out: sendActionFor(null, { kind: 'idle' }),
  },
  notes: {
    // A clean character load and a guide's link carry no note (h-audit 2026-11-09): the header and the Loaded tag say it.
    reconstructed: '', linked: plannerCopy.linkedNote, linked_short: plannerCopy.linkedNoteShort,
    // The guide's link names its spec: spent in the guide's own levelling order when the ladder reaches it, else a plain link.
    guide_reached: survPrefix !== undefined,
    band60: plannerCopy.bandLoadedNote('60'), kept: plannerCopy.keptNote(23, 25, 'Felsnap'), gear_only: plannerCopy.noTalentRanksNote,
    no_class: plannerCopy.noClassNote('monk'), invalid: invalid.ok ? '' : invalid.message,
  },
  band_tip: bandCompareCopy.bandRankTooltipLine('60', 5, 5),
  refusals: refusals(),
  copy: { planner: strip(plannerCopy), header: strip(plannerHeaderCopy), card: strip(CH), band: strip(bandCompareCopy), sim: strip(plannerSimCopy),
    addon: { copyForGame: addonCopy.copyForGame, copiedForGame: addonCopy.copiedForGame, copyFailed: addonCopy.copyFailed, buildCodeLabel: addonCopy.buildCodeLabel, addonCodeHint: addonCopy.addonCodeHint,
      importTitle: addonCopy.importTitle, importPlaceholder: addonCopy.importPlaceholder, importAction: addonCopy.importAction, importNoExport: addonCopy.importNoExport, importGetAddon: addonCopy.importGetAddon,
      importOrderApproximated: addonCopy.importOrderApproximated, importSwitching: addonCopy.importSwitchingClass('Priest'), importOlder: addonCopy.importOlderBuild('1.60.1.70630', BUILD), importDropped: addonCopy.importDropped('Lacerate') },
    cta: ctaCopy, rate_limit: RATE_LIMIT_MESSAGE, save_failed: SAVE_FAILED_MESSAGE, tree_source: treeSourceNotice(BUILD), load_failed: DATA_LOAD_FAILED,
    load_failed_detail: plannerCopy.loadFailedDetail(BUILD), read_only: READ_ONLY_REASON,
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
// The guide ladder (lib/guides/leveling-builds.ts) finds the repo through import.meta.url: the bundle gives it the lib's own.
const metaUrl = pathToFileURL(join(WEB, "src/lib/data/build-files.ts")).href;
const dir = mkdtempSync(join(tmpdir(), "planner-facts-"));
const outfile = join(dir, "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm",
  outfile, logLevel: "error", define: { "import.meta.url": JSON.stringify(metaUrl), "import.meta.env": "{}" }, loader: { ".json": "json" }, plugins: [generatedFallback] });
const facts = (await import(pathToFileURL(outfile).href)).default;
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
