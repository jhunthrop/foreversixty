// web/src/lib/character-selector/session-guild.ts
// The guild the header's Guild panel follows (spec 2026-10-15-grouped-header.md section 2):
// the session's first verified guild, as its name and its page. Pure, so the selector island
// can compute it from the cached `/v1/me` and the header script only applies the result.
import type { Me } from '../account/api';
import { pointerForCharacter } from '../account/main-character';
import { characterSlug, guildHref, parseGuildPath } from '../characters';
import type { CurrentCharacter } from '../current-character';

export interface SessionGuild {
  readonly name: string;
  readonly href: string;
}

/**
 * The guild the header's Roster and Guild settings doors lead to: any guild `/v1/me` lists
 * (a member waiting for an officer to verify them still has a guild page), the main
 * character's when it has one, else the first. `null` signed out or guildless.
 */
export function memberGuildOf(me: Me | null | undefined): SessionGuild | null {
  if (me === null || me === undefined) return null;
  const mainKey = me.main_character_key;
  const mainGuildId =
    mainKey === undefined ? undefined : me.characters?.find((c) => c.key === mainKey)?.guild?.id;
  const guild = me.guilds.find((entry) => entry.id === mainGuildId) ?? me.guilds[0];
  if (guild === undefined) return null;
  return { name: guild.name, href: guildHref(guild.region, guild.ruleset, guild.name) };
}

/** `null` for a signed-out visitor or an account with no verified guild. */
export function sessionGuildOf(me: Me | null | undefined): SessionGuild | null {
  const guild = me?.guilds.find((entry) => entry.verified);
  if (guild === undefined) return null;
  return { name: guild.name, href: guildHref(guild.region, guild.ruleset, guild.name) };
}

/**
 * The character the header shows on a guild page: the account's own character in that guild
 * when it has one (the guild band already names it), else whatever is current. A view only,
 * never written back, and only until the player picks a character on this page: a pointer
 * saved since the page loaded is the player's choice and stands.
 */
export function guildPagePointer(
  me: Me | null | undefined,
  pathname: string,
  pointer: CurrentCharacter | null,
  pageLoadedAtMs: number,
): CurrentCharacter | null {
  if (me === null || me === undefined) return pointer;
  const path = parseGuildPath(pathname);
  if (path === null) return pointer;
  const guild = me.guilds.find(
    (entry) =>
      entry.region === path.region &&
      entry.ruleset === path.ruleset &&
      characterSlug(entry.name) === path.slug,
  );
  if (guild === undefined) return pointer;
  const inGuild = me.characters.filter((character) => character.guild?.id === guild.id);
  if (inGuild.length === 0) return pointer;
  if (pointer !== null) {
    const chosenHere = Date.parse(pointer.savedAt) >= pageLoadedAtMs;
    const pointedName = pointer.label.split(' · ')[0].toLowerCase();
    const alreadyInGuild =
      inGuild.some((character) => character.key === pointer.ref) ||
      inGuild.some((character) => character.name.toLowerCase() === pointedName);
    if (chosenHere || alreadyInGuild) return pointer;
  }
  const own = inGuild.find((character) => character.key === me.main_character_key) ?? inGuild[0];
  return pointerForCharacter(own);
}
