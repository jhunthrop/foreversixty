// web/src/lib/guild/discord-bot-api.ts
// The guild bot's routes (discord bot spec §9): the install navigation, the channel list, the
// channel save, the test post and the remove. Same transport and error class as api.ts, so a
// refusal is a `GuildApiError` carrying the API's status and `code` (`channel_unusable`,
// `guild_plan_required`). Each write drops the guild's cached reads, as discord-api.ts does.
import { API_BASE_URL } from '../planner/config';
import {
  call,
  invalidateGuild,
  normalizeGuildDiscord,
  type GuildBotChannel,
  type GuildSettingsDiscord,
} from './api';

function botPath(guildId: number): string {
  return `/v1/guilds/${guildId}/discord/bot`;
}

/** The API's install route, which answers 302 to Discord's authorise page: absolute to the API
 *  origin, because the website host does not serve `/v1`. */
export function guildBotInstallUrl(guildId: number, apiBase: string = API_BASE_URL): string {
  return `${apiBase}${botPath(guildId)}/install`;
}

/** The server's text channels the bot can post in, in Discord's order. Never cached: the picker
 *  must show what the server holds now. */
export async function fetchBotChannels(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildBotChannel[]> {
  const data = await call<{ channels: GuildBotChannel[] }>(`${botPath(guildId)}/channels`, apiBase);
  return data.channels;
}

/** Saves the raid channel; answers the settings `discord` object as it now reads. */
export async function saveBotChannel(
  guildId: number,
  channelId: string,
  apiBase: string = API_BASE_URL,
): Promise<GuildSettingsDiscord> {
  const data = await call<GuildSettingsDiscord>(`${botPath(guildId)}/channel`, apiBase, {
    method: 'PUT',
    body: { channel_id: channelId },
  });
  invalidateGuild(guildId, apiBase);
  return normalizeGuildDiscord(data);
}

/** Posts the guild test embed through the bot; resolves on 204 and clears the failures. */
export async function testGuildBot(guildId: number, apiBase: string = API_BASE_URL): Promise<void> {
  await call<null>(`${botPath(guildId)}/test`, apiBase, { method: 'POST' }, true);
  invalidateGuild(guildId, apiBase);
}

/** Makes the bot leave the server; 204. */
export async function removeGuildBot(guildId: number, apiBase: string = API_BASE_URL): Promise<void> {
  await call<null>(botPath(guildId), apiBase, { method: 'DELETE' }, true);
  invalidateGuild(guildId, apiBase);
}
