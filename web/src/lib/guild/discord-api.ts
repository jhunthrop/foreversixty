// web/src/lib/guild/discord-api.ts
// The guild channel's three routes (raids spec §3.8; api task 10): the connect navigation, the
// test post and the remove. Same transport and error class as api.ts, so a refusal is a
// `GuildApiError` carrying the API's status and `code` (`guild_plan_required`, `claim_contested`).
// Its own module so the settings page's entry chunk does not carry it: only the lazily loaded
// Discord section calls these.
import { API_BASE_URL } from '../planner/config';
import { call, invalidateGuild } from './api';

function discordPath(guildId: number): string {
  return `/v1/guilds/${guildId}/notifications/discord`;
}

/** The API's connect route, which answers 302 to Discord: absolute to the API origin, as the
 *  account's `connectUrl`, because the website host does not serve `/v1`. */
export function guildDiscordConnectUrl(guildId: number, apiBase: string = API_BASE_URL): string {
  return `${apiBase}${discordPath(guildId)}/connect`;
}

/** Posts the guild test embed; resolves on 204. A success also switches a switched-off channel
 *  back on, so the guild's cached reads are dropped. */
export async function testGuildDiscord(guildId: number, apiBase: string = API_BASE_URL): Promise<void> {
  await call<null>(`${discordPath(guildId)}/test`, apiBase, { method: 'POST' }, true);
  invalidateGuild(guildId, apiBase);
}

/** Forgets the channel (and deletes its webhook at Discord, best effort); 204 even when none was
 *  saved. The raid views carry the channel too, so every cached guild read is dropped. */
export async function removeGuildDiscord(guildId: number, apiBase: string = API_BASE_URL): Promise<void> {
  await call<null>(discordPath(guildId), apiBase, { method: 'DELETE' }, true);
  invalidateGuild(guildId, apiBase);
}
