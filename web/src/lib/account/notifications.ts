// web/src/lib/account/notifications.ts
// The Discord block's pure state derivation (experience spec 2026-10-25-discord-notifications
// sections 3 and 4.1): which of five states `/v1/me` puts the block in, what the Premium row
// reads, and the inline shape check for a pasted URL. The web never decides the entitlement;
// it reads `entitlements.notifications`.
import type { DiscordNotifications, Me } from './api';
import { notificationsCopy, premiumRowCopy } from './notifications-copy';

export type DiscordBlockState =
  | { state: 'free' }
  | { state: 'empty' }
  | { state: 'saved'; discord: DiscordNotifications }
  | { state: 'disabled'; discord: DiscordNotifications }
  | { state: 'lapsed'; discord: DiscordNotifications };

export type DiscordBlockKey = DiscordBlockState['state'];

const ACCEPTED_HOSTS: ReadonlySet<string> = new Set(['discord.com', 'discordapp.com']);
// /api/webhooks/<number>/<token>, with an optional v6 to v10 segment after /api.
const WEBHOOK_PATH = /^\/api(?:\/v(?:[6-9]|10))?\/webhooks\/\d+\/[\w-]+$/;
const FAILURES_BEFORE_SWITCH_OFF = 5;

/** Which of the five states the block shows, in the spec's precedence: lapsed wins over all. */
export function blockState(me: Me): DiscordBlockState {
  const discord = me.notifications?.discord ?? null;
  const entitled = me.entitlements?.notifications === true;
  if (discord !== null) {
    if (!entitled) return { state: 'lapsed', discord };
    return discord.disabled_reason === null ? { state: 'saved', discord } : { state: 'disabled', discord };
  }
  return { state: entitled ? 'empty' : 'free' };
}

function enabledEvents(discord: DiscordNotifications): number {
  return Number(discord.events.sim_finished) + Number(discord.events.upload_parsed);
}

/** The Premium row's value: one of the five spec wordings (seven strings with the event counts). */
export function premiumRowValue(me: Me): string {
  const block = blockState(me);
  switch (block.state) {
    case 'free':
      return premiumRowCopy.off;
    case 'empty':
      return premiumRowCopy.offSetUp;
    case 'lapsed':
      return premiumRowCopy.paused;
    case 'disabled':
      return premiumRowCopy.offFailed;
    case 'saved': {
      const events = enabledEvents(block.discord);
      return events === 0 ? premiumRowCopy.offNoEvents : premiumRowCopy.on(events);
    }
  }
}

/** The row links to #discord in every state but `free`, where it is muted text. */
export function premiumRowLinked(me: Me): boolean {
  return blockState(me).state !== 'free';
}

/**
 * The saved panel's warning for 1 to 4 failures in a row; null otherwise (none yet, or already
 * switched off, where the notice replaces it).
 */
export function failureWarning(discord: DiscordNotifications): string | null {
  if (discord.failures < 1 || discord.failures >= FAILURES_BEFORE_SWITCH_OFF) return null;
  return notificationsCopy.warning(discord.failures, discord.last_error);
}

/** "channel 1098765432101234567" is drawn from the id; null while the API has none. */
export function channelLabel(channelId: string | null): string | null {
  return channelId === null || channelId === '' ? null : `${notificationsCopy.destinationNoun} ${channelId}`;
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
