// web/src/lib/character-selector/session-guild.ts
// The guild the header's Guild panel follows (spec 2026-10-15-grouped-header.md section 2):
// the session's first verified guild, as its name and its page. Pure, so the selector island
// can compute it from the cached `/v1/me` and the header script only applies the result.
import type { Me } from '../account/api';
import { guildHref } from '../characters';

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
