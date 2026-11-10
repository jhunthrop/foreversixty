// web/src/lib/setup/answer.ts
// /setup's head answer once the account has a character (setup spec U2, rulings: "Level 34 Beast
// Mastery Hunter"): the header selector's character, never switched from here.
import type { Me } from '../account/api';
import { selectedCharacter } from '../account/hero-character';
import { addonCopy } from '../addon/copy';
import { bisHref } from '../character-profile';
import { characterHref } from '../characters';
import type { CurrentCharacter } from '../current-character';

export interface SetupAnswer {
  name: string;
  /** Lower-case class slug for the crest and the class colour; null when Battle.net has not said. */
  classSlug: string | null;
  levelLine: string;
  open: { label: string; href: string };
  bis: { label: string; href: string } | null;
}

export function setupAnswer(me: Me | null, current: CurrentCharacter | null): SetupAnswer | null {
  if (me === null || me.characters.length === 0) return null;
  const character = selectedCharacter(current, me);
  if (character === null) return null;
  const cls = character.class ?? '';
  const level = character.level ?? character.build?.level;
  const levelLine =
    typeof level === 'number'
      ? addonCopy.setupLevelLine(level, character.spec ?? '', cls)
      : `${character.spec ?? ''} ${cls}`.trim();
  const bis = bisHref(character.class, character.spec ?? null, level, character.race);
  return {
    name: character.name,
    classSlug: cls === '' ? null : cls.toLowerCase(),
    levelLine,
    open: {
      label: addonCopy.setupOpenCharacter(character.name),
      href: characterHref(character.region, character.ruleset, character.name),
    },
    bis: bis === null ? null : { label: addonCopy.setupBis, href: bis },
  };
}
