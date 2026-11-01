// web/src/lib/app-data/guide-file.ts
// /data/{build}/guides/{class}-{spec}.json (Wave 3 contract W3-C1): one spec guide as the apps'
// guide reader draws it. Every block is composed from the same lib/guides functions the web's
// guide page (`pages/guides/[class]/[spec].astro`) calls, so the app and the page never differ:
// the overview split, the key talents, the rotation lists, the sim-ordered stat pills, the hit
// cap line and the race table. The Markdown keeps the prose; numbers come from the build data.
import classes from '../../data/classes.json';
import { bandEntry, readSpecCatalog } from '../bis/load';
import { presetLabelFor } from '../bis/presets';
import { bandTopLevel, loadRotations, rotationEntryFor, rotationLinesFor } from '../bis/rotation-view';
import { specTabIcon } from '../bis/spec-icon';
import type { SpecCatalogEntry } from '../bis/types';
import { decodeBuildTrees, type DecodedTreeView } from '../guides/build-tree';
import { band60Weights, factionForFirstRace, rankedBisFile } from '../guides/class-dps';
import { pageUpdatedAt } from '../guides/guide-header';
import { guideHitCapLine, schoolHitTalent } from '../guides/hit-cap-line';
import { keyTalentsFor } from '../guides/key-talents';
import { plainText, splitOverview } from '../guides/overview';
import { raceTableFor, specKindOf } from '../guides/race-table';
import { railStatRows } from '../guides/rail-stats';
import { loadApl, rotationListsFrom, specCooldowns, spellLookupFor, type RotationItem } from '../guides/rotation-lists';
import { splitSpecSections, type SpecSection } from '../guides/sections';
import type { GuideSource } from './guide-content';

/** The level every guide block answers for: the raid build, the rotation ranks and the BiS link. */
export const GUIDE_LEVEL = 60;

export type GuideSectionId = 'overview' | 'talents' | 'rotation' | 'stats' | 'gear' | 'races';
export type GuideRole = 'dps' | 'healer' | 'tank';

export interface AppGuideRace {
  slug: string;
  name: string;
  faction: string;
}

export interface AppGuideKeyTalent {
  name: string;
  icon: string | null;
  /** Points the guide's build spends in it; null when the build has no talent of that name. */
  rank: number | null;
  max_rank: number | null;
  /** The game's own text for the talent's top rank. */
  text: string | null;
  /** The guide's one line on why, sentence case, no final stop; null when it would only repeat `text`. */
  why: string | null;
}

export interface AppGuideBuild {
  /** The FS1 code the planner loads. */
  code: string;
  /** Points per tree, in the trees' left-to-right order. */
  totals: number[];
  tree_names: string[];
}

export interface AppGuideSpell {
  spell_id: number | null;
  name: string;
  /** The top rank learned by level 60; null for a step with no rank. */
  rank: number | null;
  icon: string | null;
  /** One clause, sentence case, no final stop. */
  clause: string;
}

export interface AppGuideRacial {
  name: string;
  effect: string;
}

export interface AppGuideRaceRow extends AppGuideRace {
  recommended: boolean;
  racials: AppGuideRacial[];
}

export interface AppGuideSourceLink {
  label: string;
  url: string;
  kind: string;
}

export interface AppGuideFile {
  build: string;
  /** `shaman-elemental`, the BiS file's own spec id. */
  spec: string;
  class: string;
  spec_slug: string;
  class_name: string;
  spec_name: string;
  role: GuideRole;
  icon: string | null;
  title: string;
  /** `2026-10-31`: the later of the guide's own date and its BiS file's (the web's stamp). */
  updated: string;
  summary: string;
  /** The `In a raid` block's Role line ("Ranged spell damage"), when the guide has one. */
  role_line: string | null;
  recommended_races: AppGuideRace[];
  /** In the order the web's pills draw them: the sim's per-point order, else the guide's list. */
  stat_priority: string[];
  hit_cap: string | null;
  overview: { strengths: string[]; weaknesses: string[]; raid: string[] };
  talents: {
    key: AppGuideKeyTalent[];
    leveling_build: AppGuideBuild | null;
    /** Null when the raid build is the leveling build. */
    raid_build: AppGuideBuild | null;
  };
  rotation: { opener: AppGuideSpell[]; priority: AppGuideSpell[]; cooldowns: AppGuideSpell[] } | null;
  races: AppGuideRaceRow[];
  /** The BiS list the Gear row opens: the level 60 band and the file's own preset label. */
  gear: { spec: string; band: number; preset_label: string } | null;
  sources: AppGuideSourceLink[];
  /** The sections this guide has, in reading order; the section strip draws exactly these. */
  sections: GuideSectionId[];
}

/** `2026-10-31T…` → `2026-10-31`. */
export function isoDay(date: Date): string {
  return date.toISOString().slice(0, 10);
}

/** The `- ` bullets of a Markdown block, each as plain text. */
export function bulletLines(markdown: string): string[] {
  return markdown
    .split(/\r?\n/)
    .filter((line) => /^\s*[-*]\s+/.test(line))
    .map((line) => plainText(line.replace(/^\s*[-*]\s+/, '')))
    .filter((line) => line !== '');
}

const ROLE_LINE = /^Role:\s*/i;

/** The raid block's `**Role:** …` bullet, without its label. */
function roleLineOf(raid: readonly string[]): string | null {
  const line = raid.find((entry) => ROLE_LINE.test(entry));
  return line === undefined ? null : line.replace(ROLE_LINE, '');
}

function className(slug: string): string {
  const found = classes.find((row) => row.slug === slug);
  if (found === undefined) throw new Error(`guide file: no class ${slug} in classes.json`);
  return found.name;
}

function buildOf(code: string | undefined, trees: readonly DecodedTreeView[] | undefined): AppGuideBuild | null {
  if (code === undefined || trees === undefined || trees.length === 0) return null;
  return { code, totals: trees.map((tree) => tree.points), tree_names: trees.map((tree) => tree.name) };
}

function spellOf(item: RotationItem): AppGuideSpell {
  return {
    spell_id: item.spellId ?? null,
    name: item.name,
    rank: item.rank ?? null,
    icon: item.icon ?? null,
    clause: item.clause,
  };
}

function topRankText(trees: readonly DecodedTreeView[], name: string): string | null {
  for (const tree of trees) {
    const talent = tree.talents.find((candidate) => candidate.name === name);
    const text = talent?.ranks[talent.ranks.length - 1]?.description;
    if (text !== undefined && text !== '') return text;
  }
  return null;
}

/** The web's fallback "why" is the talent's own rank text; the app prints that once. */
function repeatsText(why: string, text: string | null): boolean {
  const bare = (value: string): string => value.replace(/[.\s]+$/, '').replace(/\s+/g, ' ');
  return text !== null && bare(why) === bare(text);
}

function keyTalents(prose: string, trees: readonly DecodedTreeView[]): AppGuideKeyTalent[] {
  return keyTalentsFor(prose, trees.length > 0 ? trees : undefined).flatMap((entry) =>
    entry.talents.map((ref) => ({
      name: ref.name,
      icon: ref.icon ?? null,
      rank: ref.rank ?? null,
      max_rank: ref.maxRank ?? null,
      text: topRankText(trees, ref.name),
      why: repeatsText(entry.why, topRankText(trees, ref.name)) ? null : entry.why,
    })),
  );
}

interface RotationContext {
  build: string;
  classSlug: string;
  spec: string;
  litTalents: ReadonlySet<string>;
}

function rotationOf(ctx: RotationContext): AppGuideFile['rotation'] {
  const entry = rotationEntryFor(loadRotations(ctx.build, ctx.spec) ?? [], bandTopLevel(GUIDE_LEVEL));
  if (entry === undefined) return null;
  const facts = spellLookupFor(ctx.build, ctx.classSlug);
  const lists = rotationListsFrom(
    rotationLinesFor(entry, ctx.build, ctx.classSlug, ctx.spec),
    loadApl(ctx.spec),
    facts,
    specCooldowns(ctx.build, ctx.classSlug, ctx.litTalents, facts),
  );
  return { opener: lists.opener.map(spellOf), priority: lists.priority.map(spellOf), cooldowns: lists.cooldowns.map(spellOf) };
}

function gearOf(build: string, spec: string, classSlug: string, races: readonly string[]): AppGuideFile['gear'] {
  const file = rankedBisFile(build, spec);
  const faction = factionForFirstRace(classSlug, races);
  if (file === undefined || faction === undefined) return null;
  const band = bandEntry(file, GUIDE_LEVEL, faction);
  return band === undefined ? null : { spec, band: GUIDE_LEVEL, preset_label: presetLabelFor(file, band) };
}

function sectionsOf(guide: Omit<AppGuideFile, 'sections'>): GuideSectionId[] {
  const present: Record<GuideSectionId, boolean> = {
    overview: guide.summary !== '' || guide.overview.strengths.length + guide.overview.weaknesses.length > 0,
    talents: guide.talents.key.length > 0 || guide.talents.leveling_build !== null,
    rotation: guide.rotation !== null && guide.rotation.priority.length > 0,
    stats: guide.stat_priority.length > 0,
    gear: guide.gear !== null,
    races: guide.races.length > 0,
  };
  return (Object.keys(present) as GuideSectionId[]).filter((id) => present[id]);
}

function catalogEntry(catalog: readonly SpecCatalogEntry[], spec: string): SpecCatalogEntry {
  const entry = catalog.find((row) => row.spec === spec);
  if (entry === undefined) throw new Error(`guide file: ${spec} is not in data/curated/specs.json`);
  return entry;
}

/** The app's guide file for one spec guide, on `build`. */
export function buildGuideFile(
  build: string,
  source: GuideSource,
  catalog: readonly SpecCatalogEntry[] = readSpecCatalog(),
): AppGuideFile {
  const { data } = source;
  const [classSlug, specSlug] = source.id.split('/') as [string, string];
  const spec = `${classSlug}-${specSlug}`;
  const entry = catalogEntry(catalog, spec);
  const role: GuideRole = data.role ?? 'dps';
  const sections = splitSpecSections(source.body);
  const prose = (section: SpecSection): string => (sections.get(section) ?? '').split('\n').slice(1).join('\n');

  const levelingTrees = data.build === undefined ? [] : (decodeBuildTrees(build, data.build) ?? []);
  const raidTrees =
    data.raidBuild === undefined || data.raidBuild === data.build ? undefined : decodeBuildTrees(build, data.raidBuild);
  const overview = splitOverview(sections.get('Overview') ?? '');
  const raid = bulletLines(overview.raid);
  const weights = band60Weights(build, spec, classSlug, data.recommendedRaces);
  const statRows =
    weights === undefined ? [] : railStatRows(data.statPriority, weights.weights, spec, weights.hasteScaleFactor);
  const raceRows = raceTableFor(build, classSlug, data.recommendedRaces, specKindOf(classSlug, spec, role));
  const litTalents = new Set(
    [...levelingTrees, ...(raidTrees ?? [])].flatMap((tree) => tree.talents.filter((t) => t.rank > 0).map((t) => t.name)),
  );

  const guide: Omit<AppGuideFile, 'sections'> = {
    build,
    spec,
    class: classSlug,
    spec_slug: specSlug,
    class_name: className(classSlug),
    spec_name: entry.name,
    role,
    icon: specTabIcon(build, classSlug, entry.tree_index) ?? null,
    title: data.title,
    updated: isoDay(pageUpdatedAt(data.updated, rankedBisFile(build, spec)?.generated_at)),
    summary: overview.summary !== '' ? overview.summary : (data.description ?? ''),
    role_line: roleLineOf(raid),
    recommended_races: data.recommendedRaces.flatMap((slug) => {
      const row = raceRows.find((candidate) => candidate.slug === slug);
      return row === undefined ? [] : [{ slug, name: row.name, faction: row.faction }];
    }),
    stat_priority: statRows.length > 0 ? statRows.map((row) => row.label) : [...data.statPriority],
    hit_cap: guideHitCapLine(weights?.hitToCap, schoolHitTalent(raidTrees ?? levelingTrees))?.text ?? null,
    overview: {
      strengths: bulletLines(overview.strengths),
      weaknesses: bulletLines(overview.weaknesses),
      raid: raid.filter((line) => !ROLE_LINE.test(line)),
    },
    talents: {
      key: keyTalents(prose('Talents and builds'), levelingTrees),
      leveling_build: buildOf(data.build, levelingTrees),
      raid_build: buildOf(data.raidBuild, raidTrees),
    },
    rotation: rotationOf({ build, classSlug, spec, litTalents }),
    races: raceRows.map((row) => ({ ...row, racials: [...row.racials] })),
    gear: gearOf(build, spec, classSlug, data.recommendedRaces),
    sources: data.sources.map((link) => ({ label: link.label, url: link.url, kind: link.kind })),
  };
  return { ...guide, sections: sectionsOf(guide) };
}

