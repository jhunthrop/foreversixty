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

/** `null` for a signed-out visitor or an account with no verified guild. */
export function sessionGuildOf(me: Me | null | undefined): SessionGuild | null {
  const guild = me?.guilds.find((entry) => entry.verified);
  if (guild === undefined) return null;
  return { name: guild.name, href: guildHref(guild.region, guild.ruleset, guild.name) };
}
