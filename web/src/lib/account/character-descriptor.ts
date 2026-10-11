// web/src/lib/account/character-descriptor.ts
// The muted descriptor line the design system's Character row shows under a name --
// "Night Elf Hunter · Level 25 · Living Flame (PvP US)" (brief 2026-09-22 §B3). One place so
// CharacterList.svelte's rows and Account.svelte's hero band read the exact same text rather
// than two hand-rolled copies drifting apart. (The character's picture is always the ringed
// class crest, CharacterPortrait.svelte; the letter square and class-icon URL that lived here
// are gone, design/DESIGN-SYSTEM.md tenet 7.)
import { classDisplayName } from '../sim/spec-label';
import { rulesetLabel } from '../characters';
import type { MeCharacter } from './api';
import { characterListCopy } from './character-list-copy';

/** "Night Elf Hunter · Level 25 · Living Flame (PvP US)"; every part but the ruleset/
 *  region location is omitted when the character has never been through the Battle.net
 *  import. */
export function characterDescriptor(character: MeCharacter): string {
  const className = character.class === undefined ? undefined : classDisplayName(character.class);
  const raceClass = [character.race, className]
    .filter((part): part is string => part !== undefined)
    .join(' ');
  const location =
    character.realm === undefined
      ? locationLabel(character)
      : `${character.realm} (${locationLabel(character)})`;
  const segments = [
    raceClass,
    character.level === undefined ? '' : characterListCopy.levelPrefix(character.level),
    location,
  ];
  return segments.filter((segment) => segment !== '').join(' · ');
}

function locationLabel(character: MeCharacter): string {
  return `${rulesetLabel(character.ruleset)} ${character.region.toUpperCase()}`;
}
