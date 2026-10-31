// web/src/lib/account/hero-character.ts
// The one question every signed-in page asks of the current-character pointer: which of the
// account's characters does it name? Account's hero band, the header selector,
// the BiS card, tiers, the guides callout, the planner and the simulator all ask it here, so
// no page can answer differently from the header.
import type { CurrentCharacter } from '../current-character';
import { characterSlug } from '../characters';
import type { Me, MeCharacter } from './api';
import { mainCharacter } from './main-character';

/** "Name · Spec Class" -> "Name"; a label with no name part, or whose name is the class word, is ''. */
export function pointerNameOf(current: CurrentCharacter): string {
  const name = current.label.split(' · ')[0].trim();
  return name.toLowerCase() === current.classSlug.toLowerCase() ? '' : name;
}

/** The account character an export or a pointer names: same name (a key slug counts) and class, case-insensitive. */
export function sameCharacter(character: MeCharacter, name: string, classSlug: string): boolean {
  return (
    characterSlug(character.name) === characterSlug(name) &&
    (character.class ?? '').toLowerCase() === classSlug
  );
}

/**
 * The account character the pointer names, or null: an `armory` pointer by key, an `addon` or
 * `code` pointer by the name its label carries (a label that is only the class names nobody).
 */
export function heroCharacter(
  current: CurrentCharacter | null,
  characters: MeCharacter[],
): MeCharacter | null {
  if (current === null) return null;
  if (current.source === 'armory') {
    return characters.find((character) => character.key === current.ref) ?? null;
  }
  if (current.source !== 'addon' && current.source !== 'code') return null;
  const name = pointerNameOf(current);
  if (name === '') return null;
  return characters.find((character) => sameCharacter(character, name, current.classSlug)) ?? null;
}

/**
 * The one character a signed-in page speaks to, resolved the way the header selector does:
 * the account character the pointer names, or the account main when there is no pointer at
 * all. A pointer that names no account character (a pasted export, a fight, a build) is its
 * own identity and answers null here -- never silently the main, which would put another
 * character's gear on a page whose header says otherwise.
 */
export function selectedCharacter(
  current: CurrentCharacter | null,
  me: Pick<Me, 'characters' | 'main_character_key'>,
): MeCharacter | null {
  if (current === null) return mainCharacter(me.characters, me.main_character_key);
  return heroCharacter(current, me.characters);
}
