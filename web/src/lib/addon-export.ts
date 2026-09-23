// Whether the API's newest source for a character is an addon export -- the one signal
// this site can turn into a working `/sim?source=addon&ref=` or `/planner?code=` link
// today. See the ruling in docs/superpowers/plans/2026-09-21-op-handoffs.md, Task 4: this
// makes the same read-only call the sim's own signed-in landing state makes
// (fetchSimInput, lib/sim/api.ts) and applies the same check fromStoredCharacter does
// (lib/sim/sources.ts) -- never a guess, never a second copy of that logic.
import { FS1_PREFIX } from './planner/fs1';
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
