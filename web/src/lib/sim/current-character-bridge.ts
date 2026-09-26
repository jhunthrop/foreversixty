// web/src/lib/sim/current-character-bridge.ts
// The one place a SimCharacter becomes a CurrentCharacter pointer. Kept separate from
// current-character.ts (source-kind agnostic, no sim-lib import) and from sources.ts (keeps
// that file's own header claim -- "the four/five ways a character reaches the simulator" --
// free of a second concern).
import {
  CURRENT_CHARACTER_CHANGED,
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
    label: `${character.name} · ${specLabel(character.spec)}`,
    classSlug: character.class_slug,
    savedAt: new Date().toISOString(),
  };
  writeCurrent(pointer, storage);
  // Task 8: every other writeCurrent call site (Account.svelte, AccountMenu.svelte,
  // AddonPasteBox.svelte, CurrentCharacterBar.svelte's own Switch) dispatches this
  // immediately after, so a sibling island on the same page -- CurrentCharacterBar's
  // `spine` mode, now mounted alongside SimView.svelte and ToolsView.svelte -- learns of
  // the new pointer without a full reload. This bridge was the one caller that did not,
  // so a character loaded via addon paste on /sim or a bulk tool page left the spine bar
  // showing its signed-out line rather than the freshly loaded character's doors.
  if (typeof window !== 'undefined') window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
}
