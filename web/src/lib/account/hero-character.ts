// web/src/lib/account/hero-character.ts
// /account's hero band (brief 2026-09-22 §B4, loosened §3.1): shown whenever the
// current-character pointer names a listed character -- never a guess -- whether or not
// Battle.net has ever imported a render for it (an addon-only character has no render_url
// and is still a real hero). A pure function of the two things Account.svelte already reads
// (the pointer, the character list) so the matching rule lives in one tested place rather
// than inline in the component.
import type { CurrentCharacter } from '../current-character';
import type { MeCharacter } from './api';

export function heroCharacter(
  current: CurrentCharacter | null,
  characters: MeCharacter[],
): MeCharacter | null {
  if (current === null || current.source !== 'armory') return null;
  return characters.find((character) => character.key === current.ref) ?? null;
}
