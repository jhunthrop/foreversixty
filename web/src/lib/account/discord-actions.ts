// web/src/lib/account/discord-actions.ts
// The Discord block's four flows (experience spec 2026-10-25-discord-notifications section 4.3):
// each runs the API call and returns the words to draw and where to draw them. Only the copy
// module's sentences come back; the API's raw message is shown only as the Discord text inside
// `testFailed`, and the typed webhook URL is never part of any result.
import { fetchMeFresh } from './api';
import type { DiscordEvents, DiscordNotifications, Me } from './api';
import {
  deleteDiscordWebhook,
  NotificationsApiError,
  patchMeDiscord,
  saveDiscordWebhook,
  testDiscordWebhook,
} from './notifications-api';
import { webhookInputError } from './notifications';
import { notificationsCopy } from './notifications-copy';

/** `field` is the reserved line under the input, `plans` adds the See plans link, `status` the status line. */
export type FailureTarget = 'field' | 'plans' | 'status';

export type ActionResult =
  { ok: true; message: string } | { ok: false; target: FailureTarget; message: string };

const failed = (target: FailureTarget, message: string): ActionResult => ({ ok: false, target, message });
const succeeded = (message: string): ActionResult => ({ ok: true, message });

/**
 * Puts the server's answer into the cached `/v1/me` now, then re-reads it in the background.
 * `fetchMeFresh` (not `refetchMe`) is the silent one: `refetchMe` drops the entry first, which
 * shows the page's loading state.
 */
function applyToMe(me: Me, discord: DiscordNotifications | null): void {
  patchMeDiscord(me, discord);
  void fetchMeFresh().catch(() => {});
}

/** A test that Discord accepted resets the failure count and the switched-off reason (spec section 10). */
function clearedFailures(discord: DiscordNotifications): DiscordNotifications {
  return { ...discord, failures: 0, disabled_reason: null, last_error: null };
}

function kindOf(error: unknown): NotificationsApiError {
  return error instanceof NotificationsApiError
    ? error
    : new NotificationsApiError('failed', notificationsCopy.saveFailed, 0);
}

function saveFailure(error: unknown): ActionResult {
  switch (kindOf(error).kind) {
    case 'invalid':
    case 'discord':
      return failed('field', notificationsCopy.invalidUrl);
    case 'not-entitled':
      return failed('plans', notificationsCopy.notEntitled);
    default:
      return failed('status', notificationsCopy.saveFailed);
  }
}

function testFailure(error: unknown): ActionResult {
  const { kind, message } = kindOf(error);
  switch (kind) {
    case 'discord':
      return failed('status', notificationsCopy.testFailed(message));
    case 'rate-limited':
      return failed('status', notificationsCopy.testRateLimited);
    case 'not-entitled':
      return failed('plans', notificationsCopy.notEntitled);
    default:
      return failed('status', notificationsCopy.testOurError);
  }
}

/** A new or replacement webhook: checked here first, so a bad paste costs no round trip. */
export async function saveWebhook(me: Me, webhookUrl: string, events: DiscordEvents): Promise<ActionResult> {
  const invalid = webhookInputError(webhookUrl);
  if (invalid !== null) return failed('field', invalid);
  try {
    applyToMe(me, await saveDiscordWebhook({ webhook_url: webhookUrl.trim(), events }));
    return succeeded(notificationsCopy.savedFirst);
  } catch (error) {
    return saveFailure(error);
  }
}

/** An events-only change: the page never holds the stored URL, so none is sent. */
export async function saveEvents(me: Me, events: DiscordEvents): Promise<ActionResult> {
  try {
    applyToMe(me, await saveDiscordWebhook({ events }));
    return succeeded(notificationsCopy.savedEvents);
  } catch {
    return failed('status', notificationsCopy.saveFailed);
  }
}

/** The test post; on a switched-off webhook a success also switches it back on. */
export async function sendTest(me: Me, wasSwitchedOff: boolean): Promise<ActionResult> {
  try {
    await testDiscordWebhook();
    const saved = me.notifications?.discord ?? null;
    if (saved !== null && (saved.failures > 0 || saved.disabled_reason !== null))
      applyToMe(me, clearedFailures(saved));
    return succeeded(wasSwitchedOff ? notificationsCopy.testSentSwitchedOn : notificationsCopy.testSent);
  } catch (error) {
    return testFailure(error);
  }
}

export async function removeWebhook(me: Me): Promise<ActionResult> {
  try {
    await deleteDiscordWebhook();
    applyToMe(me, null);
    return succeeded(notificationsCopy.removed);
  } catch {
    return failed('status', notificationsCopy.removeFailed);
  }
}
