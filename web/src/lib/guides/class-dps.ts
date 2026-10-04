// web/src/lib/guides/class-dps.ts
// The class landing's spec-card DPS line (rebuild spec §4.B) and the Leveling band strip's
// own faction pick (§4.D) share one rule, stated once here rather than twice: the guide's
// own first `recommendedRaces` entry decides which faction's band data to read. Never a
// literal DPS number or a hardcoded race/faction table in this module (round-1 mock review
// finding 1's own defect, closed by reading real reference data instead).
import { existsSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { bandEntry, loadBisFile } from '../bis/load';
import { classRows, racesForClass } from '../planner/reference';
import type { BisFile, Faction } from '../bis/types';

function repoRoot(): string {
  return path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
}

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
  return existsSync(path.join(repoRoot(), 'data/builds', build, 'bis', `${spec}.json`));
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
  return bandEntry(file, 60, faction)?.set_dps;
}
