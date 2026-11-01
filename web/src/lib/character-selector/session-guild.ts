// web/src/lib/character-selector/session-guild.ts
// The guild the header's Guild panel follows (spec 2026-10-15-grouped-header.md section 2):
// the current character's guild, as its name and its page. Pure, so the selector island can
// compute it from the cached `/v1/me` and the header script only applies the result.
import type { Me } from '../account/api';
import { selectedCharacter } from '../account/hero-character';
import type { CurrentCharacter } from '../current-character';
import { guildHref } from '../characters';

export interface SessionGuild {
  readonly name: string;
  readonly href: string;
}

/**
 * The guild the header's Guild menu, the guild landing and the guild page follow: the current
 * character's own (the pointer's resolution, `selectedCharacter`; the main with no pointer),
 * verified or not, read from the character's guild line in `/v1/me` -- `me.guilds` may not list
 * it. `null` signed out, for a guildless current character and for a pointer that names no
 * account character: the menu then offers the guildless doors (owner, 2026-11-01: the menu
 * never falls back to the main's guild).
 */
export function currentGuildOf(
  me: Me | null | undefined,
  pointer: CurrentCharacter | null,
): SessionGuild | null {
  if (me === null || me === undefined) return null;
  const character = selectedCharacter(pointer, {
    characters: me.characters ?? [],
    main_character_key: me.main_character_key,
  });
  const guild = character?.guild;
  if (character === null || guild === undefined) return null;
  return { name: guild.name, href: guildHref(character.region, character.ruleset, guild.name) };
}
