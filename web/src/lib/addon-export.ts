// Whether the API's newest source for a character is an addon export -- the one signal
// this site can turn into a working `/sim?source=addon&ref=` or `/planner?code=` link
// today. See the ruling in docs/superpowers/plans/2026-09-21-op-handoffs.md, Task 4: this
// makes the same read-only call the sim's own signed-in landing state makes
// (fetchSimInput, lib/sim/api.ts) and applies the same check fromStoredCharacter does
// (lib/sim/sources.ts) -- never a guess, never a second copy of that logic.
import { FS1_PREFIX, decodeFS1 } from './planner/fs1';
import type { CharacterPath } from './characters';
import { fetchSimInput } from './sim/api';

export interface AddonExportLookup {
  path: CharacterPath;
  /** The FS1 string, present only when the API's newest source for this character is the addon. */
  code: string | null;
}

export async function lookupAddonExport(path: CharacterPath, apiBase?: string): Promise<AddonExportLookup> {
  try {
    const input = await fetchSimInput(path, apiBase);
    const code =
      (input.source === 'addon' || input.source === 'blizzard') &&
      typeof input.gear === 'string' &&
      input.gear.startsWith(`${FS1_PREFIX}:`)
        ? input.gear
        : null;
    return { path, code };
  } catch {
    return { path, code: null };
  }
}

/**
 * Whether an export string carries anything the addon read from the character: at least one
 * worn item, or at least one spent talent point. A logout before the world finished loading
 * writes an export with neither, and that is "not read yet", never a character with an empty
 * bag of slots and no talents (W.3). A string that does not decode has read nothing either.
 */
export function exportHasGear(code: string | null | undefined): boolean {
  if (typeof code !== 'string' || code === '') return false;
  const decoded = decodeFS1(code);
  if (!decoded.ok) return false;
  const { gearSlots, treeRanks } = decoded.build;
  return gearSlots.length > 0 || treeRanks.some((tree) => tree.some((rank) => rank > 0));
}

/** The same rule over the API's gear panel: a slot with a worn item. */
export function slotsHaveGear(slots: readonly { worn: unknown }[]): boolean {
  return slots.some((slot) => slot.worn !== null);
}

/** The same rule over a /v1/me character's build: worn gear, or talent points spent. */
export function buildHasGear(
  build: { gear?: Record<string, number>; talents?: { points: number[] } } | undefined,
): boolean {
  if (build === undefined) return false;
  return (
    Object.keys(build.gear ?? {}).length > 0 || (build.talents?.points ?? []).some((points) => points > 0)
  );
}

/** The same rule over a sim character built from an export: gear in a slot, or points spent. */
export function simCharacterHasGear(character: {
  gear: Record<string, unknown>;
  point_order: unknown[];
}): boolean {
  return Object.keys(character.gear).length > 0 || character.point_order.length > 0;
}
