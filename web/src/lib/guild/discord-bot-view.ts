// web/src/lib/guild/discord-bot-view.ts
// Which of the bot spec §3.1 states the guild's Discord section is in, and the height each one
// reserves. Kept free of every copy import because discord-state.ts (the settings page's entry
// chunk, which holds the section's place while its own chunk loads) imports it; the words,
// returns and picker helpers live in discord-bot-state.ts and discord-bot-copy.ts.
import type { GuildBot, GuildSettingsDiscord } from './api';

export type GuildBotState = 'picker' | 'installed' | 'unusable' | 'removed' | 'webhook-only' | 'nothing';

/** The state with the bot it was read from: the bot is present exactly for the three bot states. */
export type GuildBotView =
  | { state: 'webhook-only' | 'nothing'; bot: null }
  | { state: 'picker' | 'installed' | 'unusable' | 'removed'; bot: GuildBot };

type BotFacts = Pick<GuildSettingsDiscord, 'bot_available' | 'bot' | 'connected'>;

/** §3.1's precedence, with state 2b (`unusable`) below removed and the picker, above installed. With the bot off for the deployment only the webhook states can occur. */
export function guildBotView(discord: BotFacts, changingChannel: boolean): GuildBotView {
  const bot = discord.bot_available === true ? (discord.bot ?? null) : null;
  if (bot === null) return { state: discord.connected ? 'webhook-only' : 'nothing', bot: null };
  if (bot.removed_in_discord) return { state: 'removed', bot };
  if (bot.channel_id === null || changingChannel) return { state: 'picker', bot };
  if (bot.channel_unusable != null) return { state: 'unusable', bot };
  return { state: 'installed', bot };
}

export const guildBotState = (discord: BotFacts, changingChannel: boolean): GuildBotState =>
  guildBotView(discord, changingChannel).state;

/**
 * Section heights per state, phone then desktop. `picker` is the §3.1 cap, held on the install
 * return until the read answers. The other three are estimates: they are measured on the
 * fixture build at branch level, with the §9 captures, and replace these then.
 * A webhook-only section keeps G1's two heights plus the offer (WEBHOOK_OFFER_RESERVE).
 */
export const GUILD_BOT_RESERVE: Readonly<Record<Exclude<GuildBotState, 'webhook-only'>, string>> = {
  picker: 'min-h-[462px] md:min-h-[392px]',
  installed: 'min-h-[392px] md:min-h-[312px]',
  unusable: 'min-h-[382px] md:min-h-[302px]',
  removed: 'min-h-[362px] md:min-h-[292px]',
  nothing: 'min-h-[402px] md:min-h-[312px]',
};

/** G1's connected (250 / 230 on md) and switched-off (349 / 310) heights plus the offer's 100, and (site alignment) the Card head's 62 px every table here carries. */
export const WEBHOOK_OFFER_RESERVE = {
  connected: 'min-h-[412px] md:min-h-[392px]',
  'switched-off': 'min-h-[511px] md:min-h-[472px]',
} as const;

/** The query parameter the install callback leaves (`?discord_bot=`), its own name so it never
 *  collides with the webhook's `discord=`. */
export const BOT_RETURN_PARAM = 'discord_bot';
