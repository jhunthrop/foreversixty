// web/src/lib/account/notifications-api.ts
// The three calls the account page's Discord block makes (experience spec
// 2026-10-25-discord-notifications section 10). None of them ever carries the webhook URL back:
// the API answers the block that `/v1/me` carries. Each success drops the cached `/v1/me`, so
// the Premium row and the block re-read the state the server now holds.
import { API_BASE_URL } from '../planner/config';
import { AccountError, invalidateMe, requestEnvelope } from './api';
import type { DiscordEvents, DiscordNotifications } from './api';
import { notificationsCopy, shortenDiscordText } from './notifications-copy';

export type { DiscordEvents, DiscordNotifications } from './api';

const DISCORD_PATH = '/v1/notifications/discord';
const TEST_PATH = `${DISCORD_PATH}/test`;

const STATUS_INVALID = [400, 422];
const STATUS_NOT_ENTITLED = 402;
const STATUS_RATE_LIMITED = 429;
const STATUS_DISCORD_REFUSED = 502;

/** What went wrong, in the terms the block draws: it picks its own words per kind. */
export type NotificationsErrorKind = 'invalid' | 'not-entitled' | 'rate-limited' | 'discord' | 'failed';

export class NotificationsApiError extends Error {
  constructor(
    readonly kind: NotificationsErrorKind,
    /** The API's message; for `discord` the sanitised Discord text, at most 140 characters. */
    message: string,
    readonly status: number,
  ) {
    super(message);
    this.name = 'NotificationsApiError';
  }
}

export interface SaveDiscordRequest {
  /** Omitted for an events-only change: the page never holds the stored URL. */
  webhook_url?: string;
  events: DiscordEvents;
}

function kindOf(status: number): NotificationsErrorKind {
  if (STATUS_INVALID.includes(status)) return 'invalid';
  if (status === STATUS_NOT_ENTITLED) return 'not-entitled';
  if (status === STATUS_RATE_LIMITED) return 'rate-limited';
  if (status === STATUS_DISCORD_REFUSED) return 'discord';
  return 'failed';
}

function toNotificationsError(error: unknown): NotificationsApiError {
  if (error instanceof NotificationsApiError) return error;
  if (!(error instanceof AccountError))
    return new NotificationsApiError('failed', notificationsCopy.saveFailed, 0);
  const kind = kindOf(error.status);
  const message = kind === 'discord' ? shortenDiscordText(error.message) : error.message;
  return new NotificationsApiError(kind, message, error.status);
}

async function send<T>(path: string, method: string, apiBase: string, body?: unknown): Promise<T | null> {
  try {
    const { data } = await requestEnvelope<T>(path, apiBase, { method, body });
    invalidateMe(apiBase);
    return data;
  } catch (error) {
    throw toNotificationsError(error);
  }
}

/** Saves a new or replacement webhook, or (without `webhook_url`) only the chosen events. */
export async function saveDiscordWebhook(
  request: SaveDiscordRequest,
  apiBase: string = API_BASE_URL,
): Promise<DiscordNotifications> {
  const data = await send<DiscordNotifications>(DISCORD_PATH, 'PUT', apiBase, request);
  if (data === null) throw new NotificationsApiError('failed', notificationsCopy.saveFailed, 200);
  return data;
}

/** Posts the test message; resolves on 204 and rejects with kind `discord` when Discord refuses. */
export async function testDiscordWebhook(apiBase: string = API_BASE_URL): Promise<void> {
  await send<null>(TEST_PATH, 'POST', apiBase);
}

export async function deleteDiscordWebhook(apiBase: string = API_BASE_URL): Promise<void> {
  await send<null>(DISCORD_PATH, 'DELETE', apiBase);
}
