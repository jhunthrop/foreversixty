// web/src/lib/sim/current-character-bridge.ts
// The one place a SimCharacter becomes a CurrentCharacter pointer. Kept separate from
// current-character.ts (source-kind agnostic, no sim-lib import) and from sources.ts (keeps
// that file's own header claim -- "the four/five ways a character reaches the simulator" --
// free of a second concern).
import { adoptHeldCharacter } from '../account/adopt-held';
import {
  announcePageLoadedCharacter,
  writeCurrent,
  type CurrentCharacter,
  type CurrentCharacterSource,
} from '../current-character';
import type { SimCharacter } from './character';
import { specLabel } from './spec-label';

export function recordCurrentCharacter(
  character: SimCharacter,
  source: CurrentCharacterSource,
  ref: string,
  storage?: Storage,
): void {
  const pointer: CurrentCharacter = {
    source,
    ref,
    // A character with no name of its own (a nameless build) is its spec and class, never a
    // class word standing in for a name.
    label:
      character.name === ''
        ? `${specLabel(character.spec)} · build`
        : `${character.name} · ${specLabel(character.spec)}`,
    classSlug: character.class_slug,
    savedAt: new Date().toISOString(),
  };
  writeCurrent(pointer, storage);
  // Task 8: every other writeCurrent call site (Account.svelte, the header character
  // selector, AddonPasteBox.svelte) dispatches this immediately after, so a sibling island
  // on the same page -- the header selector -- learns of the new pointer without a full
  // reload.
  announcePageLoadedCharacter();
}

/**
 * What an FS1 export that was loaded (not chosen) does to the pointer. The pointer changes only
 * when the player chooses, so a load never overwrites the one they have:
 *  - a nameless export never becomes the pointer;
 *  - an export naming a character the account holds makes that character current (an 'armory'
 *    pointer), unless it already is;
 *  - any other named export becomes an 'addon' pointer only when the player pasted it
 *    ('paste'); one that arrived in a link ('link') -- a guild member's gear, another player's
 *    character page -- is looked at and leaves the pointer alone.
 */
export async function recordLoadedExport(
  character: SimCharacter,
  code: string,
  policy: 'paste' | 'link',
  storage?: Storage,
): Promise<void> {
  if (character.name === '') return;
  if ((await adoptHeldCharacter(character.name, character.class_slug, storage)) !== null) return;
  if (policy === 'paste') recordCurrentCharacter(character, 'addon', code, storage);
}
