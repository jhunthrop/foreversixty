// web/src/lib/guild/discord-state.ts
// Which of the three raids spec §3.8 states the guild's Discord channel is in, and the height each
// one reserves. Kept apart from discord-settings.ts (the section's words and outcomes) because the
// settings page's entry chunk needs these two only, to hold the section's place while its own
// chunk loads; the words stay in that chunk.
import type { GuildSettingsDiscord } from './api';
import {
  BOT_RETURN_PARAM,
  GUILD_BOT_RESERVE,
  WEBHOOK_OFFER_RESERVE,
  guildBotState,
} from './discord-bot-view';

const BOT_INSTALLED = 'installed';

export type GuildDiscordState = 'not-connected' | 'connected' | 'switched-off';

export function guildDiscordState(
  discord: Pick<GuildSettingsDiscord, 'connected' | 'disabled_reason'>,
): GuildDiscordState {
  if (!discord.connected) return 'not-connected';
  return discord.disabled_reason === null ? 'connected' : 'switched-off';
}

/**
 * The section's height at rest per state, phone then desktop, measured on the fixture build
 * (design/reviews/2026-10-26-guild-raids, "Discord section"): the placeholder while the chunk
 * loads, and the `connected` hold after a return from Discord, so nothing under it moves.
 */
export const GUILD_DISCORD_RESERVE: Readonly<Record<GuildDiscordState, string>> = {
  'not-connected': 'min-h-[260px] md:min-h-[224px]',
  connected: 'min-h-[312px] md:min-h-[296px]',
  'switched-off': 'min-h-[411px] md:min-h-[372px]',
};

/** The placeholder's height while the section's chunk loads: a `connected` return holds the
 *  connected height, since the settings read in hand may predate the channel (spec §3.8), and an
 *  `installed` bot return holds the picker's (bot spec §3.1). With the bot on, a webhook keeps
 *  G1's heights plus the offer, and the bot's own states reserve theirs. */
export function guildDiscordReserve(
  discord: Pick<GuildSettingsDiscord, 'connected' | 'disabled_reason' | 'bot_available' | 'bot'>,
  search: string,
): string {
  const params = new URLSearchParams(search);
  const returning = params.get('discord') === 'connected';
  if (discord.bot_available !== true)
    return GUILD_DISCORD_RESERVE[returning ? 'connected' : guildDiscordState(discord)];
  if (params.get(BOT_RETURN_PARAM) === BOT_INSTALLED) return GUILD_BOT_RESERVE.picker;
  if (returning) return WEBHOOK_OFFER_RESERVE.connected;
  const botState = guildBotState(discord, false);
  if (botState !== 'webhook-only') return GUILD_BOT_RESERVE[botState];
  return WEBHOOK_OFFER_RESERVE[guildDiscordState(discord) === 'switched-off' ? 'switched-off' : 'connected'];
}
