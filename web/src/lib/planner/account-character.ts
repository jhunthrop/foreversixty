// web/src/lib/planner/account-character.ts
// An account character as the planner's own build: its talents, gear, race and level, read from
// the account (`/v1/me`'s `characters[].build`) and, when that lacks the talent ranks, from the
// character's sim-input export. Both end as an FS1 string -- the one format the planner already
// loads a build from -- so a character opens through exactly the path a `?code=` takes.
import { parseCharacterPath } from '../characters';
import type { MeCharacter } from '../account/api';
import { FS1_PREFIX, encodeFS1V2 } from './fs1';
import type { Gear } from './types';

/** "Night Elf" -> "night-elf": the planner's race slug for the account's race name. */
export function raceSlugOf(race: string | undefined): string {
  return (race ?? '').trim().toLowerCase().replace(/\s+/g, '-');
}

/** One digit per talent cell, in tab order: "503200000" -> [5, 0, 3, 2, 0, ...]. */
function ranksOfTree(tree: string): number[] {
  return [...tree].map((digit) => {
    const rank = Number.parseInt(digit, 36);
    return Number.isNaN(rank) ? 0 : rank;
  });
}

/**
 * The character's build as an FS1 code, from what `/v1/me` already carries; null when the
 * account holds no talent ranks for it (a Battle.net-only row, an export that predates them).
 */
export function exportCodeFromAccount(character: MeCharacter, fallbackBuild: string): string | null {
  const build = character.build;
  const classSlug = character.class?.toLowerCase();
  if (build?.talents === undefined || classSlug === undefined || classSlug === '') return null;
  // Ranks the addon never read are not a build: the sim-input export is the next place to look.
  if (build.talents.points.every((points) => points === 0)) return null;
  return encodeFS1V2({
    dataBuild: build.data_build ?? fallbackBuild,
    classSlug,
    raceSlug: raceSlugOf(character.race),
    treeRanks: build.talents.trees.map(ranksOfTree),
    gear: (build.gear ?? {}) as Gear,
    level: build.level ?? character.level,
    character: { name: character.name, realm: character.realm ?? '' },
  });
}

/** The character's newest export string from its sim-input, or null when the site holds none. */
export async function exportCodeFromSimInput(key: string, apiBase: string): Promise<string | null> {
  const path = parseCharacterPath(`/character/${key}`);
  if (path === null) return null;
  try {
    const { fetchSimInput } = await import('../sim/api');
    const input = await fetchSimInput(path, apiBase);
    return typeof input.gear === 'string' && input.gear.startsWith(`${FS1_PREFIX}:`) ? input.gear : null;
  } catch {
    return null;
  }
}

/** The code the planner opens an account character with: the account's own ranks first. */
export async function accountCharacterCode(
  character: MeCharacter,
  fallbackBuild: string,
  apiBase: string,
): Promise<string | null> {
  return (
    exportCodeFromAccount(character, fallbackBuild) ?? (await exportCodeFromSimInput(character.key, apiBase))
  );
}
