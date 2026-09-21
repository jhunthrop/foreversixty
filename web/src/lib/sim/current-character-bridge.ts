// web/src/lib/sim/current-character-bridge.ts
// The one place a SimCharacter becomes a CurrentCharacter pointer. Kept separate from
// current-character.ts (source-kind agnostic, no sim-lib import) and from sources.ts (keeps
// that file's own header claim -- "the four/five ways a character reaches the simulator" --
// free of a second concern).
import { writeCurrent, type CurrentCharacter, type CurrentCharacterSource } from '../current-character';
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
}
