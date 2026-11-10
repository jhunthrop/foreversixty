// web/src/lib/guides/build-races.ts
// The build's own class, race and combo tables (classes.json, races.json, combos.json), read
// off data/builds/<build>/ like every other guide input. The planner's reference.ts carries
// the same three files, but from src/data/generated/, which FOREVER_DATA=fixture swaps for
// the two-tree warrior fixture (four classic mage races, no racial sentences); a guide shows
// the build's races whatever the planner tests run against, so it never reads them there.
import type { ClassRow, Combo, RaceRow } from '../planner/types';
import { buildIconExists } from '../data/build-files';
import { dataUrl } from '../planner/load';
import { readBuildJson, requireBuildJson } from './build-files';

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

/** The races whose pairing with `classSlug` is new in Forever (combos.json `new_in_forever`). */
export function newRaceSlugsForClass(build: string, classSlug: string): ReadonlySet<string> {
  const reference = referenceFor(build);
  const classRow = reference.classes.find((row) => row.slug === classSlug);
  if (classRow === undefined) return new Set();
  const ids = new Set(
    reference.combos
      .filter((combo) => combo.class_id === classRow.id && combo.new_in_forever)
      .map((combo) => combo.race_id),
  );
  return new Set(reference.races.filter((race) => ids.has(race.id)).map((race) => race.slug));
}

/** The build's race portraits (`data/builds/<build>/raceicons.json`, race slug to icon stem) as
 *  urls, keeping only those whose art the build ships; empty before the build's first icons run
 *  writes the file, so a race then shows its name alone. */
export function raceIconsFor(build: string): ReadonlyMap<string, string> {
  const file = readBuildJson<Record<string, string>>(build, 'raceicons.json') ?? {};
  return new Map(
    Object.entries(file)
      .filter(([, icon]) => buildIconExists(build, icon))
      .map(([slug, icon]) => [slug, dataUrl(build, `icons/${icon}.webp`)]),
  );
}
