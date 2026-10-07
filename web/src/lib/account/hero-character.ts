// web/src/lib/account/hero-character.ts
// /account's hero band (brief 2026-09-22 §B4, loosened §3.1): shown whenever the
// current-character pointer names a listed character -- never a guess -- whether or not
// Battle.net has ever imported a render for it (an addon-only character has no render_url
// and is still a real hero). A pure function of the two things Account.svelte already reads
// (the pointer, the character list) so the matching rule lives in one tested place rather
// than inline in the component.
import type { CurrentCharacter } from '../current-character';
import type { Me, MeCharacter } from './api';
import { mainCharacter } from './main-character';

export function heroCharacter(
  current: CurrentCharacter | null,
  characters: MeCharacter[],
): MeCharacter | null {
  if (current === null || current.source !== 'armory') return null;
  return characters.find((character) => character.key === current.ref) ?? null;
}

/**
 * The one character a signed-in page speaks to: the current-character pointer when it
 * names a listed character (the header's switch list writes exactly that pointer), else
 * the account main. This is the rule the header chip, the planner and the BiS card
 * already follow; a page that reads `characters[0]` instead disagrees with the header the
 * moment the player switches (owner-reported defect 2026-10-07: a warrior selected, the
 * guides callout still said Enhancement Shaman).
 */
export function selectedCharacter(
  current: CurrentCharacter | null,
  me: Pick<Me, 'characters' | 'main_character_key'>,
): MeCharacter | null {
  return heroCharacter(current, me.characters) ?? mainCharacter(me.characters, me.main_character_key);
}
