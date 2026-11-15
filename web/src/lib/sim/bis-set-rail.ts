// web/src/lib/sim/bis-set-rail.ts
// The /sim landing's "Sim a best-in-slot set" rail (spec 2026-11-06-sim §4.2, ruling S6): one row
// per class with a damage spec the sim can rerun, in the client's class order, each spec linking
// to its level 60 band's "View ranking sim" (the tier list's own link, lib/tiers/tier-load.ts).
//
// 2026-11-15 (SEO lane, owner-approved server render of the tool pages): each spec also carries
// the set's DPS and its top stat weights from the same band, plus its guide and BiS page, so the
// server HTML of /sim reads the figures a player gets from the run. Still no ranking: the specs
// keep the client's class order, with no gap or bar (the tier list owns the ranking). Figures are
// formatted here, at build time, so the island renders the same strings and carries no
// formatter. Read at build time (node fs).
import { MIN_BAND } from '../bis/band-levels';
import { bisBandPath } from '../bis/band-path';
import { bandScaleView } from '../bis/band-scale-view';
import { bandEntry, loadBisFile, readSpecCatalog } from '../bis/load';
import type { ScaleRow } from '../bis/scale-view';
import type { BisBand, BisFile, SpecCatalogEntry } from '../bis/types';
import { oneDecimal } from '../format/number';
import { classRows } from '../planner/reference';
import { tierRankingSimHref } from '../tiers/tier-load';
import { statLabelForSpec } from './weights-display';

/** The band the published ranking sims run: level 60, raid-ready (the tier list's band). */
const RAIL_BAND = 60;
/** The faction whose set the link carries; the race only changes the racial, not the set. */
const RAIL_FACTION = 'alliance';
/** How many of the band's stat weights a spec line names, strongest first. */
const TOP_STATS = 3;

export interface BisSetStat {
  /** "Strength", "Crit". */
  label: string;
  /** The scale factor, two decimals: the top stat reads "1.00". */
  factor: string;
}

export interface BisSetSpec {
  name: string;
  /** The ranking sim of the set: /sim with the set loaded and run. */
  href: string;
  /** The set's DPS, one decimal ("1,053.1"); null when the band publishes none. */
  dps: string | null;
  /** The band's strongest significant stats, best first; empty when unmeasured. */
  stats: BisSetStat[];
  guideHref: string;
  bisHref: string;
}

export interface BisSetClass {
  slug: string;
  name: string;
  specs: BisSetSpec[];
}

export interface BisSetRailDeps {
  catalog: readonly SpecCatalogEntry[];
  classes: readonly { slug: string; name: string }[];
  loadFile: (spec: string) => BisFile | null;
  simHref: (entry: SpecCatalogEntry, band: BisBand) => string | undefined;
}

export function buildBisSetRail(deps: BisSetRailDeps): BisSetClass[] {
  return deps.classes.flatMap((klass) => {
    const specs = deps.catalog
      .filter((entry) => entry.class_slug === klass.slug && entry.role === 'dps')
      .flatMap((entry) => {
        const file = deps.loadFile(entry.spec);
        const band = file === null ? undefined : bandEntry(file, RAIL_BAND, RAIL_FACTION);
        const href = band === undefined ? undefined : deps.simHref(entry, band);
        return band === undefined || href === undefined ? [] : [specLine(entry, band, href)];
      });
    return specs.length === 0 ? [] : [{ slug: klass.slug, name: klass.name, specs }];
  });
}

/** One school's spell damage: on a spec whose spells are all that school it is worth exactly
 *  what spell power is, so it is the same stat twice and would push a real one off the line. */
const SCHOOL_POWER = new Set([
  'arcane_power',
  'fire_power',
  'frost_power',
  'holy_power',
  'nature_power',
  'shadow_power',
]);
const SPELL_POWER = 'spell_power';
/** Two factors closer than this read the same at two decimals. */
const SAME_FACTOR = 0.005;

function duplicatesSpellPower(row: ScaleRow, spellPower: ScaleRow | undefined): boolean {
  if (spellPower === undefined || !SCHOOL_POWER.has(row.stat)) return false;
  return Math.abs(row.scaleFactor - spellPower.scaleFactor) <= Math.max(row.scaleError, SAME_FACTOR);
}

function topStats(band: BisBand, spec: string): BisSetStat[] {
  const rows = bandScaleView(band, spec, 'dps').rows.filter((row) => row.significant);
  const spellPower = rows.find((row) => row.stat === SPELL_POWER);
  return (
    rows
      .filter((row) => !duplicatesSpellPower(row, spellPower))
      .slice(0, TOP_STATS)
      // The bare stat name: the line is per point, as the panel's meta says ("Hit", not "Hit rating").
      .map((row) => ({ label: statLabelForSpec(row.stat, spec), factor: row.scaleFactor.toFixed(2) }))
  );
}

function specLine(entry: SpecCatalogEntry, band: BisBand, href: string): BisSetSpec {
  return {
    name: entry.name,
    href,
    dps: band.set_dps > 0 && band.set_dps_partial !== true ? oneDecimal(band.set_dps) : null,
    stats: topStats(band, entry.spec),
    guideHref: `/guides/${entry.class_slug}/${entry.spec_slug}`,
    bisHref: bisBandPath(entry.class_slug, entry.spec_slug, RAIL_BAND, MIN_BAND),
  };
}

export function loadBisSetRail(build: string): BisSetClass[] {
  return buildBisSetRail({
    catalog: readSpecCatalog(),
    classes: classRows,
    loadFile: (spec) => loadBisFile(spec, build),
    simHref: tierRankingSimHref(build),
  });
}
