// web/src/lib/guild/landing-mine.ts
// The guild landing's `Your guilds` list (guild round 3): every guild any of the account's
// characters is in, from `/v1/me` `guilds`, so a player whose current character is an alt outside
// a guild still finds their raid guild at the top of /guild. Pure, so the rows are pinned without
// a page.
import type { Me } from '../account/api';
import { guildHref, isRegion, isRuleset, rulesetLabel } from '../characters';
import { guildRosterModelCopy } from './copy';

export interface MyGuildRow {
  key: string;
  name: string;
  href: string;
  /** `US PvP`: where the guild lives, the way the lookup's own selects write it. */
  where: string;
  /** The account's rank there (`Officer`), as the roster writes it; null when /v1/me sends none. */
  rank: string | null;
}

const RANKS: Readonly<Record<string, string>> = guildRosterModelCopy.ranks;

/** The account's guilds in the order `/v1/me` sends them; none signed out. */
export function myGuildRows(me: Me | null): MyGuildRow[] {
  return (me?.guilds ?? []).map((guild) => ({
    key: `${guild.region}-${guild.ruleset}-${guild.id}`,
    name: guild.name,
    href: guildHref(guild.region, guild.ruleset, guild.name),
    where: `${guild.region.toUpperCase()} ${rulesetLabel(guild.ruleset)}`,
    rank: guild.rank === undefined ? null : (RANKS[guild.rank] ?? null),
  }));
}

/** The lookup's region and ruleset for a signed-in player: their first guild's, else none. */
export function lookupDefaults(me: Me | null): { region: string; ruleset: string } | null {
  const first = me?.guilds.find((guild) => isRegion(guild.region) && isRuleset(guild.ruleset));
  return first === undefined ? null : { region: first.region, ruleset: first.ruleset };
}

/** The guild whose claim the account holds, for the landing's foot line: a claim holder is not
 *  asked "Run a guild?" (guild loop round 3). Null when the account holds none. */
export function heldClaimRow(me: Me | null): MyGuildRow | null {
  const index = (me?.guilds ?? []).findIndex((guild) => guild.claim?.yours === true);
  return index === -1 ? null : (myGuildRows(me)[index] ?? null);
}
