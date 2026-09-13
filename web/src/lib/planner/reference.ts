// web/src/lib/planner/reference.ts
// Typed access to the three files scripts/sync-data.mjs copies into src/data/generated/.
// The planner island fetches the same data from /data/<build>/ at runtime; the classes page
// imports it at build time, because it is a static page and the data never changes between
// builds.
import combosJson from '../../data/generated/combos.json';
import classesJson from '../../data/generated/classes.json';
import racesJson from '../../data/generated/races.json';
import type { ClassRow, Combo, RaceRow } from './types';

export const classRows = classesJson as ClassRow[];
export const raceRows = racesJson as RaceRow[];
export const comboRows = combosJson as Combo[];

export function comboFor(raceId: number, classId: number): Combo | null {
  return comboRows.find((combo) => combo.race_id === raceId && combo.class_id === classId) ?? null;
}

/** The races that can be this class, in the order races.json lists them. */
export function racesForClass(classId: number): RaceRow[] {
  return raceRows.filter((race) => comboFor(race.id, classId) !== null);
}

export function plannerHref(classSlug: string, raceSlug?: string): string {
  return raceSlug ? `/planner?class=${classSlug}&race=${raceSlug}` : `/planner?class=${classSlug}`;
}
