// web/src/lib/character-header/descriptor.ts
// The header's descriptor: "Gnome Fury Warrior · Level 60 · [logo] Alliance". Race only when
// the API has sent it, never a placeholder; a bare class when the spec is not known. The two
// halves are separate fields because a narrow container splits the line in two, so a separator
// can never start or end a line.
import type { MeCharacter } from '../account/api';
import type { Faction } from '../faction-mark';
import { classSlugFromName } from '../report/tree-sizes';
import { characterSpecName } from '../character-spec';
import { classDisplayName } from '../sim/spec-label';
import { characterHeaderCopy } from './copy';

export interface DescriptorLines {
  /** "Gnome Fury Warrior"; empty when the character has no class either. */
  identity: string;
  /** "Level 60"; null when the level is not known. */
  level: string | null;
  faction: Faction | null;
}

export function descriptorLines(character: MeCharacter): DescriptorLines {
  const className =
    character.class === undefined ? undefined : classDisplayName(classSlugFromName(character.class));
  const identity = [character.race, characterSpecName(character), className]
    .filter((part): part is string => part !== undefined && part !== '')
    .join(' ');
  return {
    identity,
    level: character.level === undefined ? null : `${characterHeaderCopy.levelWord} ${character.level}`,
    faction: character.faction ?? null,
  };
}
