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

/**
 * Home rebuild spec §3.B.1's signed-in hero descriptor: "Level 24 Troll Marksmanship
 * Hunter" in the mock -- a DIFFERENT word order from `characterDescriptor`'s own "Race
 * Class · Level N · Realm (Ruleset REGION)" line, and mixed with markup
 * (`HomeAccountPanel.svelte` inserts a `FactionMark` icon and the guild tag around this
 * text, plus the faction word in the faction's own colour) that a single plain-text
 * function cannot produce anyway. Kept as its own small helper, not an extension of
 * `characterDescriptor` (which three other surfaces -- Account.svelte's hero band,
 * CharacterList.svelte and sim/LandingState.svelte -- render as-is today, unrelated to this
 * lane and out of scope to reshape): reusing its own building blocks
 * (`classDisplayName`/`rulesetLabel` are for the OTHER function's own realm/ruleset clause,
 * not needed here) keeps this honest and small rather than overloading one function with
 * two unrelated output shapes behind a flag.
 *
 * The mock's line also names a spec ("Marksmanship Hunter"): `MeCharacter.spec` (see that
 * field's own TODO) carries it once `GET /v1/me` ships it, rendered between race and class
 * ("Level 24 Troll Marksmanship Hunter"); until then the field is undefined and this omits
 * the clause rather than guess (tenet 8) -- "Level 24 Troll Hunter", never a fabricated
 * "Marksmanship" or a placeholder like "Unknown spec".
 */
export function homeHeroLevelRaceClassLine(character: MeCharacter): string {
  const className = character.class === undefined ? undefined : classDisplayName(character.class);
  const parts = [
    character.level === undefined ? undefined : characterListCopy.levelPrefix(character.level),
    character.race,
    character.spec,
    className,
  ].filter((part): part is string => part !== undefined);
  return parts.join(' ');
}

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
