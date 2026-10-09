// web/src/lib/guides/class-dps.ts
// The class landing's spec-card DPS line (rebuild spec §4.B) and the Leveling band strip's
// own faction pick (§4.D) share one rule, stated once here rather than twice: the guide's
// own first `recommendedRaces` entry decides which faction's band data to read. Never a
// literal DPS number or a hardcoded race/faction table in this module (round-1 mock review
// finding 1's own defect, closed by reading real reference data instead).
import { existsSync } from 'node:fs';
import { bandEntry, loadBisFile, realBisPath } from '../bis/load';
import { presetLabelFor } from '../bis/presets';
import { headlineDpsOf } from '../bis/tank-view';
import { classRows, racesForClass } from '../planner/reference';
import type { BisFile, BisExpertiseToCap, BisHitToCap, BisStatWeight, Faction } from '../bis/types';

/**
 * Whether the real pipeline has published a BiS file for this spec -- `loadBisFile` (used
 * below for the actual numbers once this is true) also falls back to a committed test
 * fixture (`src/data/fixtures/bis/<spec>.json`, `load.ts`'s own header comment: "until [the
 * real file] lands") when the pipeline has not published one yet. That fallback exists so
 * `/bis` always has something to render during development; a class landing card showing a
 * fixture's own placeholder DPS figure next to two specs' real, sim-ranked numbers would be
 * exactly the fabricated-precision the design review's "never a zero or dash standing in
 * for missing data" rule exists to prevent, so this card's own eligibility check reads the
 * real file's existence directly, never through the fixture-falling-back loader.
 */
export function hasRealBisFile(build: string, spec: string): boolean {
  return existsSync(realBisPath(build, spec));
}

/**
 * `loadBisFile`, but `undefined` whenever the pipeline has not published a real file yet
 * (see `hasRealBisFile`'s own doc comment) -- the spec guide rail's Stat priority card and
 * the Leveling band strip (§4.C.3/§4.D) share this same "real data only" rule with the
 * class landing's DPS line, so none of this rebuild's three new numeric surfaces ever reads
 * a committed test fixture as if it were a shipped ranking.
 */
export function rankedBisFile(build: string, spec: string): BisFile | undefined {
  if (!hasRealBisFile(build, spec)) return undefined;
  return loadBisFile(spec, build) ?? undefined;
}

/**
 * The faction of a guide's own first `recommendedRaces` entry -- `undefined` when the guide
 * names no recommended race, or names one this class cannot legally play (should not occur
 * on real content, but never guessed at).
 */
export function factionForFirstRace(
  classSlug: string,
  recommendedRaces: readonly string[],
): Faction | undefined {
  const firstRace = recommendedRaces[0];
  if (firstRace === undefined) return undefined;
  const classRow = classRows.find((row) => row.slug === classSlug);
  if (classRow === undefined) return undefined;
  const race = racesForClass(classRow.id).find((candidate) => candidate.slug === firstRace);
  if (race === undefined) return undefined;
  return race.faction === 'alliance' || race.faction === 'horde' ? race.faction : undefined;
}

/**
 * The class landing's own Set DPS line (spec §4.B): this spec's band-60 `set_dps`, at the
 * faction `factionForFirstRace` resolves -- `undefined` for a spec with no ranked BiS file
 * (Protection) or no band-60 row at that faction, never a zero or a dash standing in for
 * missing data (tenet 8).
 */
export function classLandingSetDps(
  build: string,
  /** The spec's full id (`specs.json`'s own `spec` field, e.g. `"warrior-fury"`) -- the
   *  BiS file's own name on disk, not the bare spec slug (`"fury"`). */
  spec: string,
  classSlug: string,
  recommendedRaces: readonly string[],
): number | undefined {
  const file = rankedBisFile(build, spec);
  if (file === undefined) return undefined;
  const faction = factionForFirstRace(classSlug, recommendedRaces);
  if (faction === undefined) return undefined;
  const band = bandEntry(file, 60, faction);
  return band === undefined ? undefined : headlineDpsOf(band);
}

/** This spec's band-60 stat weights, plus the file-level timestamp they were published
 *  with -- everything `railStatRows` (rail-stats.ts) and its own "Updated" caption need. */
export interface Band60Weights {
  weights: readonly BisStatWeight[];
  hasteScaleFactor: number | null;
  hitToCap: BisHitToCap | null;
  expertiseToCap: BisExpertiseToCap | null;
  generatedAt: string;
  /** Which preset these weights were measured under, as the file labels it. */
  presetLabel: string;
}

/**
 * This spec's band-60 per-point stat weights, at the faction `factionForFirstRace` resolves
 * -- the one load path the Stat priority rail card and `GuideStatTable.astro` (live-numbers
 * lane, 2026-10-06) share, so a guide's own numbers never diverge by reading the band
 * through two different call sequences. `undefined` for a spec with no ranked BiS file or
 * no band-60 row at that faction, same as `classLandingSetDps`'s own "real data only" rule.
 */
export function band60Weights(
  build: string,
  spec: string,
  classSlug: string,
  recommendedRaces: readonly string[],
): Band60Weights | undefined {
  const file = rankedBisFile(build, spec);
  if (file === undefined) return undefined;
  const faction = factionForFirstRace(classSlug, recommendedRaces);
  if (faction === undefined) return undefined;
  const band = bandEntry(file, 60, faction);
  if (band === undefined) return undefined;
  return {
    weights: band.weights,
    hasteScaleFactor: band.haste_scale_factor ?? null,
    hitToCap: band.hit_to_cap ?? null,
    expertiseToCap: band.expertise_to_cap ?? null,
    generatedAt: file.generated_at,
    presetLabel: presetLabelFor(file, band),
  };
}
