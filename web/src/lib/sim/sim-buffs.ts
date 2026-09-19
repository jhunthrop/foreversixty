// web/src/lib/sim/sim-buffs.ts
// data/builds/<build>/simbuffs.json: the display name and icon for every IDS.md buff,
// debuff, world buff and consumable id.
//
// The engine's ids are snake case and legible on their own -- "flask_of_supreme_power" --
// so a missing row is a de-underscored id rather than a blank, and a build that ships no
// file loses the icons and nothing else.
import { dataUrl, fetchJson, DataLoadError } from '../planner/load';

export interface SimBuffRow {
  name: string;
  icon: string;
}

export interface SimBuffFile {
  entries: Record<string, SimBuffRow>;
}

const EMPTY: SimBuffFile = { entries: {} };

/**
 * Optional file, matching `loadSets`'s convention in `planner/load.ts`: a 404 means the
 * build ships no such file and resolves to the empty file; every other failure is a broken
 * build and is rethrown.
 */
export async function loadSimBuffs(build: string): Promise<SimBuffFile> {
  try {
    return await fetchJson<SimBuffFile>(dataUrl(build, 'simbuffs.json'));
  } catch (error) {
    if (error instanceof DataLoadError && error.status === 404) return EMPTY;
    throw error;
  }
}

export function buffName(file: SimBuffFile, id: string): string {
  return file.entries[id]?.name ?? id.replaceAll('_', ' ');
}

/** "" when the file has no row: the caller draws no icon rather than a broken one. */
export function buffIcon(file: SimBuffFile, id: string): string {
  return file.entries[id]?.icon ?? '';
}
