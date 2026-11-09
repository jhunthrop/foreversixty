// web/src/lib/character/your-characters.ts
// /characters' "Your characters" rail (directions mock b2-characters): each of the account's
// characters as its name in the class colour, "Level 60 Fury Warrior · <Guild> · Item level 61"
// under it, and its page.
import type { MeCharacter } from '../account/api';
import { characterListCopy } from '../account/character-list-copy';
import { characterSpecName } from '../character-spec';
import { characterHref } from '../characters';
import { gearCopy } from '../gear/copy';
import { classDisplayName } from '../sim/spec-label';
import { classTone, type TitleTone } from '../ui/primitives';

export interface YourCharacterRow {
  key: string;
  name: string;
  /** The name's class colour; the strong colour before the Battle.net import names the class. */
  tone: TitleTone;
  /** The class slug for the crest ('' before the import names it). */
  classSlug: string;
  /** "Level 60 Fury Warrior"; the parts the import has not learned are left out. */
  line: string;
  href: string;
}

export function yourCharacterRows(characters: readonly MeCharacter[]): YourCharacterRow[] {
  return characters.map((character) => {
    const slug = character.class === undefined ? '' : character.class.toLowerCase();
    const line = [
      character.level === undefined ? undefined : characterListCopy.levelPrefix(character.level),
      characterSpecName(character),
      slug === '' ? undefined : classDisplayName(slug),
    ]
      .filter((part): part is string => part !== undefined && part !== '')
      .join(' ');
    const extra = [
      character.guild === undefined ? undefined : `<${character.guild.name}>`,
      character.item_level === undefined ? undefined : gearCopy.paperdoll.itemLevelLine(character.item_level),
    ].filter((part): part is string => part !== undefined);
    return {
      key: character.key,
      name: character.name,
      tone: classTone(slug),
      classSlug: slug,
      line: [line, ...extra].filter((part) => part !== '').join(' · '),
      href: characterHref(character.region, character.ruleset, character.name),
    };
  });
}
