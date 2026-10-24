// web/src/lib/guides/build-races.ts
// The build's own class, race and combo tables (classes.json, races.json, combos.json), read
// off data/builds/<build>/ like every other guide input. The planner's reference.ts carries
// the same three files, but from src/data/generated/, which FOREVER_DATA=fixture swaps for
// the two-tree warrior fixture (four classic mage races, no racial sentences); a guide shows
// the build's races whatever the planner tests run against, so it never reads them there.
import type { ClassRow, Combo, RaceRow } from '../planner/types';
import { requireBuildJson } from './build-files';

interface BuildReference {
  classes: readonly ClassRow[];
  races: readonly RaceRow[];
  combos: readonly Combo[];
}

const loaded = new Map<string, BuildReference>();

/** The three tables of `build`, read once per build for the whole page build. */
function referenceFor(build: string): BuildReference {
  const cached = loaded.get(build);
  if (cached !== undefined) return cached;
  const reference: BuildReference = {
    classes: requireBuildJson<ClassRow[]>(build, 'classes.json'),
    races: requireBuildJson<RaceRow[]>(build, 'races.json'),
    combos: requireBuildJson<Combo[]>(build, 'combos.json'),
  };
  loaded.set(build, reference);
  return reference;
}

/** Every race the build knows, in races.json's order. */
export function buildRaces(build: string): readonly RaceRow[] {
  return referenceFor(build).races;
}

/** One race by slug; `undefined` when the build has no such race. */
export function buildRace(build: string, slug: string): RaceRow | undefined {
  return referenceFor(build).races.find((race) => race.slug === slug);
}

/** The races that can be `classSlug`, in races.json's order; none for a class the build does
 *  not know. */
export function buildRacesForClass(build: string, classSlug: string): RaceRow[] {
  const reference = referenceFor(build);
  const classRow = reference.classes.find((row) => row.slug === classSlug);
  if (classRow === undefined) return [];
  return reference.races.filter((race) =>
    reference.combos.some((combo) => combo.race_id === race.id && combo.class_id === classRow.id),
  );
}
