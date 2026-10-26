// web/src/lib/account/notifications.ts
// The Discord block's pure state derivation (experience spec 2026-10-25-discord-notifications
// sections 3 and 4.1): which of five states `/v1/me` puts the block in, what the Premium row
// reads, and the inline shape check for a pasted URL. The web never decides the entitlement;
// it reads `entitlements.notifications`.
import type { DiscordNotifications, DiscordSource, Me } from './api';
import { API_BASE_URL } from '../planner/config';
import { discordHintState } from './layout';
import { notificationsCopy, premiumRowCopy } from './notifications-copy';

export type DiscordBlockState =
  | { state: 'off' }
  | { state: 'free' }
  | { state: 'empty' }
  | { state: 'saved'; discord: DiscordNotifications }
  | { state: 'disabled'; discord: DiscordNotifications }
  | { state: 'lapsed'; discord: DiscordNotifications };

export type DiscordBlockKey = DiscordBlockState['state'];

export type WarningLines = ReturnType<typeof notificationsCopy.warningLines>;

const ACCEPTED_HOSTS: ReadonlySet<string> = new Set(['discord.com', 'discordapp.com']);
// /api/webhooks/<number>/<token>, with an optional v6 to v10 segment after /api.
const WEBHOOK_PATH = /^\/api(?:\/v(?:[6-9]|10))?\/webhooks\/\d+\/[\w-]+$/;
const FAILURES_BEFORE_SWITCH_OFF = 5;

/**
 * Which state the block shows, in the spec's precedence: `off` (no `notifications` key: the
 * deployment has the feature off) wins over all, then lapsed.
 */
export function blockState(me: Me): DiscordBlockState {
  if (me.notifications === undefined) return { state: 'off' };
  const discord = me.notifications.discord;
  const entitled = me.entitlements?.notifications === true;
  if (discord !== null) {
    if (!entitled) return { state: 'lapsed', discord };
    return discord.disabled_reason === null ? { state: 'saved', discord } : { state: 'disabled', discord };
  }
  return { state: entitled ? 'empty' : 'free' };
}

/** The Premium event count is the two post events; `upload_private` only widens one of them. */
function enabledEvents(discord: DiscordNotifications): number {
  return Number(discord.events.sim_finished) + Number(discord.events.upload_parsed);
}

/** The row exists unless the deployment has notifications off (no `notifications` key, spec 3.7). */
export function premiumRowVisible(me: Me): boolean {
  return discordHintState(me) !== 'none';
}

/**
 * The Premium row's value: one of the seven spec wordings, non-breaking spaces included. Callers
 * check `premiumRowVisible` first: the `none` case has no row, so this value is never shown.
 */
export function premiumRowValue(me: Me): string {
  const block = blockState(me);
  switch (block.state) {
    case 'off':
    case 'free':
      return premiumRowCopy.off;
    case 'empty':
      return premiumRowCopy.offSetUp;
    case 'lapsed':
      return premiumRowCopy.paused;
    case 'disabled':
      return premiumRowCopy.offNeedsAttention;
    case 'saved': {
      const events = enabledEvents(block.discord);
      return events === 0 ? premiumRowCopy.offNoEvents : premiumRowCopy.on(events);
    }
  }
}

/** The row links to #discord in every state but `free` and `off`, where it is muted text. */
export function premiumRowLinked(me: Me): boolean {
  const { state } = blockState(me);
  return state !== 'free' && state !== 'off';
}

/**
 * The saved panel's warning for 1 to 4 failures in a row; null otherwise (none yet, or already
 * switched off, where the notice replaces it).
 */
export function failureWarning(
  discord: Pick<DiscordNotifications, 'failures' | 'last_error'>,
): WarningLines | null {
  if (discord.failures < 1 || discord.failures >= FAILURES_BEFORE_SWITCH_OFF) return null;
  return notificationsCopy.warningLines(discord.failures, discord.last_error);
}

function parseUrl(input: string): URL | null {
  try {
    return new URL(input);
  } catch {
    return null;
  }
}

/**
 * Mirrors the API's rule: https only, host exactly discord.com or discordapp.com, no userinfo or
 * port, path /api/webhooks/<number>/<token>. A query or fragment is refused, as the API refuses it.
 * The API is the authority; this only spares a round trip.
 */
export function looksLikeWebhookURL(input: string): boolean {
  const trimmed = input.trim();
  const url = parseUrl(trimmed);
  return (
    url !== null &&
    !/[?#]/.test(trimmed) &&
    url.search === '' &&
    url.hash === '' &&
    url.protocol === 'https:' &&
    ACCEPTED_HOSTS.has(url.hostname) &&
    url.username === '' &&
    url.password === '' &&
    url.port === '' &&
    WEBHOOK_PATH.test(url.pathname)
  );
}

/** The inline message for a pasted value, or null when it is fit to send. Never echoes the input. */
export function webhookInputError(input: string): string | null {
  if (input.trim() === '') return notificationsCopy.emptyField;
  return looksLikeWebhookURL(input) ? null : notificationsCopy.invalidUrl;
}

/** Whether this deployment offers the one-click connect (a missing key reads false). */
export function connectAvailable(me: Me): boolean {
  return me.notifications?.connect === true;
}

const CONNECT_PATH = '/v1/notifications/discord/connect';

/** The API's route that answers 302 to Discord's authorise page; absolute, as the website host
 *  (Cloudflare Pages) does not serve it. */
export function connectUrl(apiBase: string = API_BASE_URL): string {
  return `${apiBase}${CONNECT_PATH}`;
}

export type DiscordReturnFlag = 'connected' | 'denied' | 'failed';

const RETURN_FLAGS: readonly DiscordReturnFlag[] = ['connected', 'denied', 'failed'];

/** The `discord` outcome the callback left in the query string; any other value is no flag. */
export function discordReturnFlag(search: string): DiscordReturnFlag | null {
  const value = new URLSearchParams(search).get('discord');
  return RETURN_FLAGS.find((flag) => flag === value) ?? null;
}

/** A webhook saved by an older API has no source: it was pasted. */
export function sourceOf(discord: DiscordNotifications): DiscordSource {
  return discord.source ?? 'pasted';
}

/** The remove question for the webhook's source (spec 4.3). */
export function removeQuestionFor(discord: DiscordNotifications): string {
  return sourceOf(discord) === 'connected'
    ? notificationsCopy.removeQuestionConnected
    : notificationsCopy.removeQuestionPasted;
}

const DIGITS = /^\d+$/;

/** Where `Open in Discord` goes; null until the guild is known. Numeric ids only, so nothing else can reach the URL. */
export function openInDiscordUrl(
  discord: Pick<DiscordNotifications, 'guild_id' | 'channel_id'>,
): string | null {
  const guild = discord.guild_id;
  if (guild === null || guild === undefined || !DIGITS.test(guild) || !DIGITS.test(discord.channel_id))
    return null;
  return `https://discord.com/channels/${guild}/${discord.channel_id}`;
}
