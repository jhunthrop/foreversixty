// web/src/lib/bis/hover.ts
// The planner gear slot's "Best in slot at <band>" popover: band math, the per-spec file
// fetch (with an in-memory cache so a spec is fetched at most once per page), and the
// previous-band diff that drives the "new at this band" marker and the "was: …" line.
//
// Deliberately does not import web/src/lib/bis/load.ts: that module reads data/curated/ and
// data/builds/ straight off disk with node:fs for Astro's build-time /bis pages, which never
// runs in a browser bundle. The planner island is a browser bundle (design step 2: it must
// not bundle every spec's BiS JSON at all), so this file re-derives the handful of pure
// lookups load.ts also has (band levels, a band's entry) against the same BisFile shape
// (./types) rather than pulling node:fs into the island.
import { specRow } from '../sim/spec-label';
import type { BisBand, BisFile, BisSlot, Faction } from './types';

/** Bands run 20..60 in steps of 10 (sim/cmd/leveling-bis's own defaultBandsFlag). */
export const MIN_BAND = 20;
export const MAX_BAND = 60;
const BAND_STEP = 10;

/**
 * The band a character level falls into: rounded down to the nearest step, clamped to the
 * range the BiS files cover. A level below 20 still gets the level 20
 * band -- the earliest gear list is the honest answer for "not leveled yet" -- and a level
 * at or above 60 gets the level 60 band.
 */
export function bandForLevel(level: number): number {
  const clamped = Math.min(MAX_BAND, Math.max(MIN_BAND, level));
  return Math.floor(clamped / BAND_STEP) * BAND_STEP;
}

export function bisAssetUrl(build: string, spec: string): string {
  return `/data/${build}/bis/${spec}.json`;
}

/** The bands a file covers, ascending, deduplicated across factions -- load.ts's own
 *  `bandLevels`, copied rather than imported (see the file header). */
export function bandLevelsFor(file: BisFile): number[] {
  return [...new Set(file.bands.map((band) => band.band))].sort((a, b) => a - b);
}

/** load.ts's own `bandEntry`, copied for the same reason. */
export function bandEntryFor(file: BisFile, band: number, faction: Faction): BisBand | undefined {
  return file.bands.find((entry) => entry.band === band && entry.faction === faction);
}

export interface SlotBandPick {
  band: number;
  pick: BisSlot | undefined;
}

export interface SlotHoverDiff {
  band: number;
  faction: Faction;
  /** This band's pick for the slot, or undefined when the band has no known source for it. */
  pick: BisSlot | undefined;
  /** The next lower band this file has data for, or undefined at the file's lowest band. */
  previous: SlotBandPick | undefined;
  /** True once a lower band exists and its pick differs from this one (including "had none,
   *  now has one"). Always false at the file's lowest band -- there is nothing to be new
   *  relative to. */
  isNewAtBand: boolean;
}

/**
 * The single-slot diff the popover shows: this band's pick, the previous band's pick for
 * the same slot (for "was: …"), and whether this band's pick is new relative to it. Mirrors
 * the pipeline's own per-band `new_at_band` diff (types.ts's doc comment on `BisBand`), but
 * computed for one slot on demand rather than parsed back out of that field's `"<slot>:
 * <item name>"` strings.
 *
 * The previous band is `band - BAND_STEP` by arithmetic, not a walk back through whichever
 * bands the file happens to carry: every band from MIN_BAND..MAX_BAND is ranked for both
 * factions by contract, so a band a faction's file is missing (which the real pipeline
 * never produces) reads as "nothing known at the previous band" rather than reaching
 * further back for one that exists -- a smaller, cheaper rule than that reach would need
 * (this ships in the planner island, which has its own tight gzipped budget; see
 * scripts/check-island-size.mjs), and one the real data can never actually exercise.
 */
export function slotHoverDiff(file: BisFile, band: number, faction: Faction, slot: string): SlotHoverDiff {
  const pick = bandEntryFor(file, band, faction)?.slots.find((row) => row.slot === slot);
  if (band <= MIN_BAND) {
    return { band, faction, pick, previous: undefined, isNewAtBand: false };
  }
  const previousBand = band - BAND_STEP;
  const previousPick = bandEntryFor(file, previousBand, faction)?.slots.find((row) => row.slot === slot);
  const isNewAtBand = pick !== undefined && pick.item_id !== previousPick?.item_id;
  return { band, faction, pick, previous: { band: previousBand, pick: previousPick }, isNewAtBand };
}

/**
 * `/bis/<class>/<spec>?faction=<faction>#band-<faction>-<band>` -- the same route and
 * in-page anchor src/pages/bis/[class]/[spec].astro renders (its `#panel-${faction}` CSS
 * toggle plus each band's `#band-${faction}-${band}` id), built from the spec key the
 * planner already computes (addon/score.ts's specKeyFor) rather than a class/spec slug pair
 * the caller would otherwise have to carry alongside it. An unrecognised spec key links to
 * the index rather than a broken route.
 */
export function bisPageHref(spec: string, faction: Faction, band: number): string {
  const row = specRow(spec);
  if (row === null) return '/bis';
  return `/bis/${row.class_slug}/${row.spec_slug}?faction=${faction}#band-${faction}-${band}`;
}

// --- fetch + in-memory cache -------------------------------------------------------------

const cache = new Map<string, Promise<BisFile | null>>();

async function fetchBisFileUncached(build: string, spec: string): Promise<BisFile | null> {
  let response: Response;
  try {
    response = await fetch(bisAssetUrl(build, spec), { headers: { accept: 'application/json' } });
  } catch {
    return null; // offline, blocked, or otherwise unreachable: the popover's own empty
    // state already reads as "nothing here," which is true either way from a hover card.
  }
  if (!response.ok) return null; // 404 (no file published for this spec yet) or a broken
  // build; a hover popover is not the place for a retry/error banner either way.
  try {
    return (await response.json()) as BisFile;
  } catch {
    return null;
  }
}

/**
 * A spec's BiS file, fetched at most once per (build, spec) for the life of the page and
 * cached in memory -- the planner island never bundles this data (design step 2), so
 * hovering a spec nobody has hovered yet costs exactly one small fetch, and every later
 * hover of the same spec (a different slot, a second hover) costs none. Concurrent hovers of
 * the same spec before the first fetch resolves share the one in-flight request.
 */
export function fetchBisFile(build: string, spec: string): Promise<BisFile | null> {
  const key = `${build}::${spec}`;
  let pending = cache.get(key);
  if (pending === undefined) {
    pending = fetchBisFileUncached(build, spec);
    cache.set(key, pending);
  }
  return pending;
}

/** Tests only: forgets every cached fetch, so one test's fixture never answers the next. */
export function resetBisHoverCache(): void {
  cache.clear();
}
