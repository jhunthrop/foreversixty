// The served data and copy the guide mocks render (design/specs/2026-11-08-guides.md):
//   node design/mocks/guides_facts.mjs [out.json]  ->  design/mocks/data/guides-facts.json (or out.json)
// Runs the guide pages' own models, bundled with esbuild from web/node_modules, never retyped:
//  - the content: web/src/content/guides/**/*.md (frontmatter through `yaml`, the body as Astro's entry.body);
//  - /guides: the class and spec catalogue, pageUpdatedAt over every guide (index.astro's own fold);
//  - /guides/<class>: splitClassLanding, cardLineFor, renderGuideMarkdown ([class]/index.astro's calls);
//  - /guides/<class>/<spec>: every call [spec].astro and its components make (splitSpecSections, splitOverview,
//    decodeBuildTrees, guideBuildBlocks, diffBuilds, keyTalentsFor, guideRotation, markRaidBuildRows, withTipKeys,
//    band60Weights, railStatRows, rankedStatLabels, guideHitCapLine, specGuideData, raceTableFor, recommendedRaceLines,
//    consumableLines, guideLadderSteps, linkProse over renderGuideMarkdown, headerMeta, dataSourceLine);
//  - the signed-in answer: descriptorLines (lib/character-header) on the home fixtures, tierCopy.answer's states;
//  - the copy module: lib/guides/copy.ts (guidesCopy), lib/tiers/tier-copy.ts (answer states).
// Every lib module resolves the repo through import.meta.url; the bundle defines it as lib/guides/build-files.ts's own.
import { createRequire } from "node:module";
import { mkdtempSync, writeFileSync, readFileSync, readdirSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const YAML = require("yaml");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/guides-facts.json", import.meta.url);
const src = (p) => JSON.stringify(WEB + p);

// The three mocked specs (a melee, a caster, a healer; Survival is melee in Forever) and the partial-data spec
// the coverage scan below finds. Three class landings follow them.
const SPECS = ["hunter/survival", "mage/fire", "priest/holy"];
const CLASSES = ["hunter", "mage", "priest"];

// The collection: what getCollection('guides') hands the pages.
const DIR = join(WEB, "src/content/guides");
const entries = [];
for (const cls of readdirSync(DIR, { withFileTypes: true }).filter((d) => d.isDirectory())) {
  for (const f of readdirSync(join(DIR, cls.name)).filter((f) => f.endsWith(".md"))) {
    const raw = readFileSync(join(DIR, cls.name, f), "utf8");
    const m = /^---\n([\s\S]*?)\n---\n?([\s\S]*)$/.exec(raw);
    const data = YAML.parse(m[1]);
    data.updated = new Date(data.updated);
    entries.push({ id: `${cls.name}/${f.replace(/\.md$/, "")}`.replace(/\/index$/, ""), data, body: m[2] });
  }
}

const entry = `
import classes from ${src("src/data/classes.json")};
import activeBuild from ${src("src/data/active-build.json")};
import me34 from ${src("tests/e2e/fixtures/home/me-34.json")};
import me60 from ${src("tests/e2e/fixtures/home/me-60.json")};
import { readSpecCatalog } from ${src("src/lib/bis/load.ts")};
import { bandTopLevel } from ${src("src/lib/bis/rotation-view.ts")};
import { latest } from ${src("src/lib/dates.ts")};
import { guidesCopy } from ${src("src/lib/guides/copy.ts")};
import { pillLabelFor } from ${src("src/lib/sources.ts")};
import { addonCopy } from ${src("src/lib/addon/copy.ts")};
import { tiersCopy as tierCopy } from ${src("src/lib/tiers/tier-copy.ts")};
import { descriptorLines } from ${src("src/lib/character-header/descriptor.ts")};
import { decodeBuildTrees } from ${src("src/lib/guides/build-tree.ts")};
import { describedBuildTrees, guideRotation } from ${src("src/lib/guides/guide-rotation.ts")};
import { band60Weights, rankedBisFile } from ${src("src/lib/guides/class-dps.ts")};
import { renderGuideMarkdown } from ${src("src/lib/guides/markdown.ts")};
import { railStatRows, rankedStatLabels, railStatValueText } from ${src("src/lib/guides/rail-stats.ts")};
import { splitSpecSections, slugify, SPEC_SECTIONS } from ${src("src/lib/guides/sections.ts")};
import { diffBuilds, buildDiffLine } from ${src("src/lib/guides/build-diff.ts")};
import { guideBuildBlocks } from ${src("src/lib/guides/guide-builds.ts")};
import { loadBuildHref, simBuildHref } from ${src("src/lib/guides/build-links.ts")};
import { proseNames, specGuideData, withTipKeys } from ${src("src/lib/guides/guide-tips.ts")};
import { linkProse } from ${src("src/lib/guides/prose-tips.ts")};
import { dayMonth, pageUpdatedAt } from ${src("src/lib/guides/guide-header.ts")};
import { guideLadderSteps } from ${src("src/lib/guides/leveling-builds.ts")};
import { keyTalentsFor, withoutTalentBullets } from ${src("src/lib/guides/key-talents.ts")};
import { splitOverview } from ${src("src/lib/guides/overview.ts")};
import { raceTableFor, specKindOf, recommendedRaceLines } from ${src("src/lib/guides/race-table.ts")};
import { consumableLines } from ${src("src/lib/guides/gear-extras.ts")};
import { guideExpertiseCapLine, guideHitCapLine } from ${src("src/lib/guides/hit-cap-line.ts")};
import { bandTalentHits } from ${src("src/lib/guides/talent-hit.ts")};
import { CLASS_LANDING_PANELS, cardLineFor, splitClassLanding } from ${src("src/lib/guides/class-landing.ts")};
import { litTalentsOf, markRaidBuildRows, raidBuildOnlyTalents, talentNamesOf } from ${src("src/lib/guides/rotation-lists.ts")};
import { glanceRotation, glanceStatRows, markCheapRanks, nextTalentPoint, racesByFaction, specRowTag, topRankAt60 } from ${src("src/lib/guides/glance.ts")};
import { buildRaces } from ${src("src/lib/guides/build-races.ts")};

const build = activeBuild.build;
const catalog = readSpecCatalog();
const strip = (o) => JSON.parse(JSON.stringify(o, (k, v) => (typeof v === 'function' ? '[fn]' : v instanceof Map ? Object.fromEntries(v) : v instanceof Set ? [...v] : v)));
const specsOf = (cls) => catalog.filter((s) => s.class_slug === cls).sort((a, b) => a.tree_index - b.tree_index);

export async function run(entries, SPECS, CLASSES) {
  // ---- /guides (index.astro)
  const updated = latest(entries.map((g) => { const spec = g.data.spec === undefined ? undefined : g.data.classSlug + '-' + g.data.spec;
    return pageUpdatedAt(g.data.updated, spec === undefined ? undefined : rankedBisFile(build, spec)?.generated_at); }));
  const index = { updated: dayMonth(updated), classes: classes.map((c) => ({ slug: c.slug, name: c.name, color: c.color,
    specs: specsOf(c.slug).map((s) => { const role = ['dps', 'healer', 'tank'].includes(s.role) ? s.role : 'dps';
      return { slug: s.spec_slug, name: s.name, icon: s.icon, role, tag: specRowTag(c.slug, c.slug + '-' + s.spec_slug, role) ?? null }; }) })) };

  // ---- /guides/<class> ([class]/index.astro)
  const landings = {};
  for (const cls of CLASSES) {
    const e = entries.find((g) => g.id === cls);
    const landing = splitClassLanding(e.body, 'guides/' + cls + '/index.md');
    const sections = [];
    for (const s of landing.sections) sections.push({ heading: s.heading, id: slugify(s.heading), panel: CLASS_LANDING_PANELS.includes(s.heading),
      html: (await renderGuideMarkdown(s.markdown)).replace(/<p><strong>([^<]+)<\\/strong><\\/p>/g, '<p class="glabel">$1</p>') });
    const specGuides = entries.filter((g) => g.data.classSlug === cls && g.data.spec !== undefined);
    const cards = specsOf(cls).flatMap((s) => { const g = specGuides.find((x) => x.data.spec === s.spec_slug); if (!g || !g.data.role) return [];
      return [{ slug: s.spec_slug, name: s.name, role: g.data.role, tag: specRowTag(cls, cls + '-' + s.spec_slug, g.data.role) ?? null, icon: s.icon, line: cardLineFor(g.body, 'guides/' + g.id + '.md') }]; });
    landings[cls] = { title: classes.find((c) => c.slug === cls).name, identity: landing.identity, sections, cards, sources: e.data.sources,
      updated: dayMonth(e.data.updated) };
  }

  // ---- /guides/<class>/<spec> ([spec].astro, in its own order)
  const guides = {};
  for (const id of SPECS) guides[id] = await specPage(entries.find((g) => g.id === id));

  // ---- coverage: which spec lacks which block (the partial-data state)
  const coverage = {};
  for (const g of entries.filter((x) => x.data.spec !== undefined)) {
    const [cls, spec] = g.id.split('/');
    const full = cls + '-' + spec;
    const trees = g.data.build ? decodeBuildTrees(build, g.data.build) ?? [] : [];
    const raidTrees = g.data.raidBuild && g.data.raidBuild !== g.data.build ? decodeBuildTrees(build, g.data.raidBuild) : undefined;
    const rot = guideRotation(build, cls, full, g.data.role ?? 'dps', describedBuildTrees(trees, raidTrees));
    const b60 = band60Weights(build, full, cls, g.data.recommendedRaces);
    const { bis, extras } = specGuideData(build, cls, full, g.data.recommendedRaces);
    coverage[g.id] = { role: g.data.role, raid_build: raidTrees !== undefined, rotation: rot ? { opener: rot.opener.length, priority: rot.priority.length, cooldowns: rot.cooldowns.length } : null,
      stat_rows: b60 ? railStatRows(g.data.statPriority, b60.weights, full, b60.haste).length : 0, bis: bis ? bis.bands.length : 0, extras: extras ? extras.enchants.length + '/' + extras.consumables.length : null,
      ranked: rankedBisFile(build, full) !== undefined };
  }

  // ---- the signed-in answer's lines, from the fixtures (a Survival hunter is me-34 with its spec set: derived, marked)
  // DERIVED (marked): a level 34 Survival hunter is me-34 with its points in the Survival tree; a level 60 Fire mage is
  // me-60 with its class, race and points set (its name stays the fixture's). No fixture carries either.
  const pts = (c, points) => ({ ...c.build, talents: { ...c.build.talents, points } });
  const meSurvival = { ...me34, characters: me34.characters.map((c) => ({ ...c, spec: 'Survival', build: pts(c, [0, 0, 23]) })) };
  const meSurvival60 = { ...me60, characters: me60.characters.map((c) => ({ ...c, class: 'Hunter', spec: 'Survival', race: 'Night Elf', build: pts(c, [5, 11, 35]) })) };
  const meMage = { ...me60, characters: me60.characters.map((c) => ({ ...c, class: 'Mage', spec: 'Fire', race: 'Human', build: pts(c, [14, 34, 3]) })) };
  const nextOf = (g, level) => { const n = nextTalentPoint(g.leveling.ladder, level); return n === undefined ? null : guidesCopy.nextPointAt(n.level, n.name); };
  const who = (me) => { const c = me.characters[0]; const d = descriptorLines(c); return { name: c.name, class: c.class, spec: c.spec, level: c.level, race: c.race, faction: c.faction, identity: d.identity, level_line: d.level, talents_at: guidesCopy.talentsAtLevel(c.level) }; };
  return { build, copy: strip(guidesCopy), tier_answer: strip(tierCopy.answer), sections: [...SPEC_SECTIONS].map((s) => ({ title: s, id: slugify(s) })),
    index, landings, guides, coverage, me: { hunter34: who(me34), survival34: { ...who(meSurvival), next: nextOf(guides['hunter/survival'], 34) }, survival60: who(meSurvival60), warrior60: who(me60), mage60: who(meMage) },
    source_labels: Object.fromEntries(['blizzard', 'datamined', 'community', 'site'].map((k) => [k, pillLabelFor(k)])), setup_title: addonCopy.setupPageTitle,
    fns: { signedInCallout: { bm: guidesCopy.signedInCallout('Beast Mastery', 'Hunter'), survival: guidesCopy.signedInCallout('Survival', 'Hunter'), fury: guidesCopy.signedInCallout('Fury', 'Warrior') },
      dataSourceLine: guidesCopy.dataSourceLine(build), unknown: tierCopy.answer.unknown('Felsnap'),
      levelRangeLabel: Object.fromEntries([20, 30, 40, 50, 60].map((b) => [b, guidesCopy.levelRangeLabel(b)])) } };
}

async function specPage(entry) {
  const sections = splitSpecSections(entry.body ?? '');
  const proseOnly = (s) => { const [, ...rest] = (sections.get(s) ?? '').split('\\n'); return rest.join('\\n').replace(/^\\n+/, ''); };
  const fac = (f) => (f === 'alliance' ? 'Alliance' : 'Horde');
  const buildCode = entry.data.build, classSlug = entry.data.classSlug, specSlug = entry.id.split('/')[1];
  const fullSpecId = classSlug + '-' + specSlug, role = entry.data.role ?? 'dps';
  const className = classes.find((c) => c.slug === classSlug)?.name ?? classSlug;
  const { bis: guideBis, extras: gearExtras, bundle: tipBundle } = specGuideData(build, classSlug, fullSpecId, entry.data.recommendedRaces);
  const buildTrees = buildCode === undefined ? [] : (decodeBuildTrees(build, buildCode) ?? []);
  const raidCode = entry.data.raidBuild;
  const raidTrees = raidCode === undefined || raidCode === buildCode ? undefined : decodeBuildTrees(build, raidCode);
  const proseTargets = proseNames(build, classSlug, raidTrees ?? buildTrees, tipBundle);
  const prose = async (md) => linkProse(await renderGuideMarkdown(md), proseTargets);
  const overview = splitOverview(sections.get('Overview') ?? '');
  const tabs = specsOf(classSlug).map((s) => ({ slug: s.spec_slug, name: s.name, icon: s.icon, active: s.spec_slug === specSlug }));
  const specName = tabs.find((t) => t.active)?.name ?? specSlug;
  const raidDiffLine = raidTrees === undefined || buildTrees.length === 0 ? undefined : buildDiffLine(diffBuilds(buildTrees, raidTrees));
  const keyTalents = keyTalentsFor(proseOnly('Talents and builds'), buildTrees.length > 0 ? buildTrees : undefined, raidTrees);
  const describedTrees = describedBuildTrees(buildTrees, raidTrees);
  const describedCode = raidTrees === undefined ? buildCode : raidCode;
  const unkeyed = guideRotation(build, classSlug, fullSpecId, role, describedTrees);
  const raidOnly = raidTrees === undefined ? new Set() : raidBuildOnlyTalents(talentNamesOf(build, classSlug), litTalentsOf(buildTrees), litTalentsOf(raidTrees));
  const rotation = unkeyed === undefined ? undefined : markCheapRanks(withTipKeys(markRaidBuildRows(unkeyed, raidOnly), tipBundle), (name) => topRankAt60(build, classSlug, name), proseOnly('Rotation and priority'));
  // Each row's rank as the page prints it ("Rank 7 (cheap rank)").
  const ranked_ = (items) => items.map((x) => (x.rank === undefined ? x : { ...x, rankText: guidesCopy.rankLabel(x.rank, x.cheapRank === true) }));
  if (rotation !== undefined) for (const k of ['opener', 'priority', 'cooldowns']) rotation[k] = ranked_(rotation[k]);
  const glance = rotation === undefined ? { rows: [] } : glanceRotation(rotation.priority, buildRaces(build).map((r) => r.name));
  const ranked = rankedBisFile(build, fullSpecId);
  const band60 = band60Weights(build, fullSpecId, classSlug, entry.data.recommendedRaces);
  const statRows = band60 === undefined ? [] : railStatRows(entry.data.statPriority, band60.weights, fullSpecId, band60.haste);
  // GuideStatTable.astro's own fold
  const talentHits = band60 === undefined ? [] : bandTalentHits(build, classSlug, band60.talents);
  const hitCap = guideHitCapLine(band60?.hitToCap, talentHits);
  const expertiseCap = guideExpertiseCapLine(band60?.expertiseToCap, talentHits);
  const unitRow = statRows.find((r) => r.reading === 'value' && r.perPercent === undefined && r.perPoint !== undefined);
  const rated = statRows.filter((r) => r.reading === 'value' && r.ratingFactor !== undefined);
  const statTable = { anchor: statRows.find((r) => r.reading === 'value' && r.perPercent === undefined)?.label, hitCap, expertiseCap, role: band60?.role ?? 'dps',
    rows: statRows.map((r) => ({ ...r, valueText: railStatValueText(r) })),
    unitLine: unitRow ? guidesCopy.unitLine(unitRow.label.toLowerCase(), unitRow.perPoint >= 10 ? unitRow.perPoint.toFixed(1) : unitRow.perPoint.toFixed(2), band60?.role ?? 'dps', unitRow.ratingFactor !== undefined) : undefined,
    ratingNote: rated.length ? guidesCopy.ratingNote(rated.map((r) => r.label).join(' and '), rated.length > 1, rated.map((r) => guidesCopy.ratingConversion(r.label.toLowerCase(), r.ratingFactor)).join(' and ')) : undefined,
    perPointHead: guidesCopy.perPointColumn(band60?.role ?? 'dps', rated.length > 0), weightHead: guidesCopy.weightColumn(statRows.find((r) => r.reading === 'value' && r.perPercent === undefined)?.label ?? ''),
    notOnItems: statRows.filter((r) => r.reading === 'not-on-items').map((r) => guidesCopy.notOnItemsNote(r.label, r.perPercent, band60?.role ?? 'dps')),
    insignificant: statRows.filter((r) => r.reading === 'not-significant').map((r) => r.label), noGain: statRows.filter((r) => r.reading === 'no-gain').map((r) => r.label) };
  const raceRows = raceTableFor(build, classSlug, [...entry.data.recommendedRaces, ...(entry.data.alsoRecommendedRaces ?? [])], specKindOf(classSlug, fullSpecId, role));
  const racesProse = proseOnly('Races');
  const picks = recommendedRaceLines(raceRows, racesProse);
  const raceLines = picks.map((p) => guidesCopy.recommendedRacesLine(raceRows.find((r) => r.slug === p.slug)?.name ?? p.slug, fac(p.faction),
    picks.filter((o) => o.faction === p.faction).length === 1 ? 'the' : 'a', p.names.join(' and ')));
  const treeIndex = specsOf(classSlug).find((s) => s.spec_slug === specSlug)?.tree_index;
  const steps = buildCode === undefined || treeIndex === undefined ? [] : guideLadderSteps(build, classSlug, fullSpecId, treeIndex, buildCode);
  // The tree at a level: the step's talent string in the guide's own FS1 code (what the leveling tree draws at that level).
  const atLevel = (lvl) => { const s = steps.find((x) => x.level === lvl); if (!s) return undefined; const p = buildCode.split(':'); p[4] = s.talents.replace(/-/g, '/');
    return { level: lvl, name: s.name, rank: s.rank, point: guidesCopy.levelingPoint(lvl, s.name, s.rank), spent: guidesCopy.levelingPointsSpent(lvl - 9, 60 - lvl), trees: decodeBuildTrees(build, p.join(':')) }; };
  const LEVEL_STRIP = /^### Levels 20 to 60\\s*$/m;
  const lev = proseOnly('Leveling'); const mm = LEVEL_STRIP.exec(lev);
  const [early, stripHead, late] = mm === null ? ['', '', lev] : [lev.slice(0, mm.index), mm[0], lev.slice(mm.index + mm[0].length)];
  const blocks = buildCode === undefined ? [] : guideBuildBlocks(buildCode, raidCode).map((b) => ({ ...strip(b), trees: decodeBuildTrees(build, b.code), load: loadBuildHref(b.code), sim: simBuildHref(b.code) }));
  return {
    id: entry.id, classSlug, specSlug, className, specName, role, title: specName + ' ' + className, description: entry.data.description,
    summary: overview.summary !== '' ? overview.summary : entry.data.description,
    meta: guidesCopy.headerMeta(role, dayMonth(pageUpdatedAt(entry.data.updated, ranked?.generated_at))), tabs,
    band: { isRaidBuild: raidTrees !== undefined, code: describedCode, trees: describedTrees, rotationBand: String(bandTopLevel(60)),
      rotationTitle: role === 'healer' ? guidesCopy.healingPriorityRailTitle : guidesCopy.rotationRailTitle,
      rotation: glance.rows, filler: glance.filler ? guidesCopy.glanceFiller(glance.filler.name) : null,
      stats: glanceStatRows(statRows).map((r) => ({ ...r, valueText: railStatValueText(r) })), hitCap, unitLine: statTable.unitLine },
    overview: { strengths: await prose(overview.strengths), weaknesses: await prose(overview.weaknesses), raid: await prose(overview.raid) },
    talents: { blocks, raidDiffLine, keyTalents: strip(keyTalents.map((k) => ({ ...k, talents: k.talents.map((t) => ({ ...t, rankText: t.buildRanks
      ? [guidesCopy.keyTalentBuildRank('leveling', t.buildRanks.leveling, t.maxRank), guidesCopy.keyTalentBuildRank('raid', t.buildRanks.raid, t.maxRank)]
      : t.rank === undefined ? [] : [t.rank + '/' + t.maxRank] })) }))), prose: await prose(withoutTalentBullets(proseOnly('Talents and builds'))) },
    rotation: { lists: rotation ?? null, forRaidBuild: rotation !== undefined && raidTrees !== undefined, prose: await prose(proseOnly('Rotation and priority')) },
    stats: { pills: rankedStatLabels(statRows, entry.data.statPriority), table: statTable, prose: await prose(proseOnly('Stat priority')) },
    bis: { data: guideBis ? strip({ ...guideBis, panels: guideBis.panels.map((p) => ({ ...p, caption: guidesCopy.bisCaption(p.raceName, fac(p.faction)) })) }) : null, href: '/bis/' + classSlug + '/' + specSlug, prose: await prose(proseOnly('Best in slot')) },
    extras: { data: gearExtras ? { ...strip(gearExtras), lines: consumableLines(gearExtras.consumables) } : null, prose: await prose(proseOnly('Enchants and consumables')) },
    races: { rows: strip(raceRows), groups: strip(racesByFaction(raceRows)), lines: raceLines, prose: await prose(racesProse) },
    professions: await prose(proseOnly('Professions')),
    leveling: { early: await prose(early), stripHeading: stripHead.replace(/^#+\\s*/, '').trim(), late: await prose(late), steps: steps.length, ladder: steps.map((x) => ({ level: x.level, name: x.name })), at: { 34: atLevel(34), 60: atLevel(60) } },
    sources: entry.data.sources, statPriority: entry.data.statPriority,
  };
}`;
const dir = mkdtempSync(join(tmpdir(), "guides-facts-"));
const outfile = join(dir, "facts.mjs");
const metaUrl = pathToFileURL(WEB + "src/lib/guides/build-files.ts").href;
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm", outfile,
  logLevel: "error", loader: { ".json": "json" }, external: ["@astrojs/markdown-satteri", "@playwright/test"],
  define: { "import.meta.url": JSON.stringify(metaUrl), "import.meta.env": "{}" } });
// The one external (the markdown processor, a native package) resolves from web/node_modules.
const bundled = readFileSync(outfile, "utf8").replace(/from "@astrojs\/markdown-satteri"/g, `from ${JSON.stringify(pathToFileURL(require.resolve("@astrojs/markdown-satteri")).href)}`);
writeFileSync(outfile, bundled);
const { run } = await import(pathToFileURL(outfile).href);
const facts = await run(entries, SPECS, CLASSES);
// Strings that live in markup, not a copy module (gap G5): read from the files themselves, verbatim.
const literal = (file, re) => { const m = re.exec(readFileSync(join(WEB, file), "utf8")); if (!m) throw new Error(file + ": " + re); return m[1]; };
facts.markup = {
  statHead: literal("src/components/guides/GuideStatTable.astro", />(Stat)</),
  indexTitle: literal("src/pages/guides/index.astro", /<Base\s+title="([^"]+)"/),
  indexDescription: literal("src/pages/guides/index.astro", /description="([^"]+)"/),
};
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
