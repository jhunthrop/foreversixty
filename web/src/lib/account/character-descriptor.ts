// web/src/lib/account/character-descriptor.ts
// The muted descriptor line the design system's Character row shows under a name --
// "Night Elf Hunter · Level 25 · Living Flame (PvP US)" (brief 2026-09-22 §B3) -- and the
// class-initial fallback square shown when a character has no avatar_url. One place for
// both so CharacterList.svelte's rows and Account.svelte's hero band read the exact same
// text and letter rather than two hand-rolled copies drifting apart.
import { classColorVar } from '../report/format';
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

/** The class-coloured fallback square's letter and colour, for a character with no
 *  avatar_url -- the class's first letter (design/DESIGN-SYSTEM.md's Character row) in
 *  the class's own colour, or the site's own text colour for a character with no class on
 *  file yet. */
export function classSquare(character: MeCharacter): { letter: string; color: string } {
  const label = character.class === undefined ? character.name : classDisplayName(character.class);
  return { letter: label.charAt(0).toUpperCase(), color: classColorVar(character.class) };
}
