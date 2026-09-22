// web/src/lib/account/hero-character.ts
// /account's hero band (brief 2026-09-22 §B4): shown only when the current-character
// pointer names a listed character that has a render_url -- never a guess. A pure function
// of the two things Account.svelte already reads (the pointer, the character list) so the
// matching rule lives in one tested place rather than inline in the component.
import type { CurrentCharacter } from '../current-character';
import type { MeCharacter } from './api';

export function heroCharacter(
  current: CurrentCharacter | null,
  characters: MeCharacter[],
): MeCharacter | null {
  if (current === null || current.source !== 'armory') return null;
  const match = characters.find((character) => character.key === current.ref);
  return match?.render_url === undefined ? null : match;
}
