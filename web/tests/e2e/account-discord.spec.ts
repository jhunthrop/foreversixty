// web/tests/e2e/account-discord.spec.ts
// The Discord block on the account page (experience spec 2026-10-25-discord-notifications, section
// 9, amended 2026-10-25): paste and save, the webhook name and channel lines, the three event rows
// (the private row's disabled state and three-key save), the steps, test, the five-failure and
// lapsed states, remove, the free sentence, the Premium row, layout shift through save and remove,
// and an axe pass per state.
// /v1/me and the three notification routes are stubbed; the fixture webhook URL and token appear
// only in the typed input and the PUT body the stub records.
import { createRequire } from 'node:module';
import { expect, test, type Page } from '@playwright/test';
import {
  DISCORD_NAME_LIMIT,
  notificationsCopy as copy,
  premiumRowCopy,
} from '../../src/lib/account/notifications-copy';
import { RECENT, routeRecent } from './support/coach-stubs';
import { gotoHydrated } from './support/hydrated';
import {
  CHANNEL_ID,
  CONNECT_URL_GLOB,
  discordBlock,
  E2E_API_BASE,
  GUILD_ID,
  routeDiscordNotifications,
  signInMember,
  type DiscordMeStages,
  type DiscordRouteOptions,
  type DiscordRoutes,
} from './support/discord-notifications';
import { drainShifts, recordShifts, sumShifts } from './support/layout-shift';

const WEBHOOK_ID = '1087654321012345678';
const WEBHOOK_TOKEN = 'xT9fixtureTokenNotARealOne_abcdEFGH';
const VALID_URL = `https://discord.com/api/webhooks/${WEBHOOK_ID}/${WEBHOOK_TOKEN}`;
const INVALID_URL = 'https://example.com/api/webhooks/1/tok';
const BOTH_EVENTS = { sim_finished: true, upload_parsed: true, upload_private: false };
const SAVED_CLEAR_TIMEOUT_MS = 6000;
// Past the 4 s that `Saved.` keeps: a line that is still there has no timer (spec 4.3, first-use fixes).
const PAST_SAVED_CLEAR_MS = 4500;
const MAX_SHIFT = 0.02;
const PHONE = { width: 390, height: 844 };
const DESKTOP = { width: 1440, height: 1000 };

async function openAccount(
  page: Page,
  stages: DiscordMeStages,
  routeOptions: DiscordRouteOptions = {},
  path = '/account',
): Promise<DiscordRoutes> {
  const seen = await routeDiscordNotifications(page, routeOptions);
  await signInMember(page, stages, seen);
  // The coach's last-raid line also reads on this page; without a stub its request hangs on
  // the test server and the line's loading skeleton stays up through the Discord steps.
  await routeRecent(page, RECENT.none.recent);
  await gotoHydrated(page, path, 'account');
  return seen;
}

const panel = (page: Page) => page.getByTestId('account-discord');
const status = (page: Page) => page.getByTestId('discord-status');
const input = (page: Page) => page.locator('#discord-webhook');
const premiumRow = (page: Page) => page.getByTestId('account-premium-notifications');

/** Every localStorage and sessionStorage key and value, as one string. */
async function browserStorage(page: Page): Promise<string> {
  return page.evaluate(() =>
    [localStorage, sessionStorage]
      .flatMap((store) => Object.keys(store).flatMap((key) => [key, store.getItem(key) ?? '']))
      .join('\n'),
  );
}

async function pasteAndSave(page: Page, url: string): Promise<void> {
  await input(page).fill(url);
  await page.getByTestId('discord-save').click();
}

/** Line one of the destination for a webhook name, assembled from the copy module's two parts. */
function namedLine(name: string): string {
  const named = copy.destinationNamed(name);
  return `${named.lead} ${named.name}`;
}

const SAVED = discordBlock();
const FIVE_FAILURES = discordBlock({
  disabled_reason: 'failed five times',
  last_error: 'Unknown Webhook',
  failures: 5,
});

test.describe('paste form', () => {
  test('an invalid host shows the fixed message, sends nothing and echoes nothing', async ({ page }) => {
    const seen = await openAccount(page, { initial: null });
    await pasteAndSave(page, INVALID_URL);
    await expect(page.getByTestId('discord-webhook-error')).toContainText(copy.invalidUrl);
    expect(seen.puts).toHaveLength(0);
    await expect(input(page)).toHaveValue(INVALID_URL);
    expect(await page.content()).not.toContain('example.com/api/webhooks');
  });

  test('an empty field shows the empty message', async ({ page }) => {
    const seen = await openAccount(page, { initial: null });
    await page.getByTestId('discord-save').click();
    await expect(page.getByTestId('discord-webhook-error')).toContainText(copy.emptyField);
    expect(seen.puts).toHaveLength(0);
  });

  test('a valid save sends one PUT, shows the channel and never puts the URL in the page', async ({
    page,
  }) => {
    const seen = await openAccount(page, { initial: null, afterPut: SAVED });
    await pasteAndSave(page, VALID_URL);
    await expect(page.getByTestId('discord-destination')).toContainText(namedLine('Raid pings'));
    await expect(page.getByTestId('discord-destination')).not.toContainText('Channel 1198');
    expect(seen.puts).toEqual([{ webhook_url: VALID_URL, events: BOTH_EVENTS }]);
    await expect(status(page)).toContainText(copy.savedFirst);
    await expect(page.getByTestId('discord-test')).toBeFocused();
    const html = await page.content();
    expect(html).not.toContain(WEBHOOK_TOKEN);
    expect(html).not.toContain(WEBHOOK_ID);
    expect(html).not.toContain('discord.com/api/webhooks');
    expect(await browserStorage(page)).not.toMatch(new RegExp(`${WEBHOOK_TOKEN}|${WEBHOOK_ID}`));
    await expect(premiumRow(page)).toContainText('On · 2 events');
  });

  test('the form shows the three numbered steps above the field', async ({ page }) => {
    await openAccount(page, { initial: null });
    for (const [index, step] of copy.howSteps.entries()) {
      await expect(page.getByTestId(`discord-step-${index + 1}`)).toContainText(step);
    }
    const lastStepTop = (await page.getByTestId('discord-step-3').boundingBox())?.y ?? Infinity;
    const fieldTop = (await input(page).boundingBox())?.y ?? 0;
    expect(lastStepTop).toBeLessThan(fieldTop);
  });

  test('a refusal for want of Premium shows the sentence and a plans link', async ({ page }) => {
    await openAccount(page, { initial: null }, { putFailure: { status: 402, message: 'premium_required' } });
    await pasteAndSave(page, VALID_URL);
    const error = page.getByTestId('discord-webhook-error');
    await expect(error).toContainText(copy.notEntitled);
    await expect(error.getByRole('link', { name: copy.seePlans })).toHaveAttribute('href', '/premium#notifications');
    await expect(input(page)).toBeFocused();
  });
});

test.describe('saved webhook', () => {
  test('a test that is sent', async ({ page }) => {
    const seen = await openAccount(page, { initial: SAVED });
    await page.getByTestId('discord-test').click();
    await expect(status(page)).toContainText(copy.testSent);
    expect(seen.tests).toBe(1);
  });

  test('Sent stays until the next action, and a press in the panel clears it', async ({ page }) => {
    await openAccount(page, { initial: SAVED });
    await page.getByTestId('discord-test').click();
    await expect(status(page)).toContainText(copy.testSent);
    await page.waitForTimeout(PAST_SAVED_CLEAR_MS);
    await expect(status(page)).toContainText(copy.testSent);
    await page.getByTestId('discord-remove').click();
    await expect(status(page)).toHaveCount(0);
  });

  test('a test Discord refuses shows its words as an alert', async ({ page }) => {
    const seen = await openAccount(
      page,
      { initial: SAVED },
      { testFailure: { status: 502, message: 'Missing Permissions' } },
    );
    await page.getByTestId('discord-test').click();
    await expect(status(page)).toContainText(copy.testFailed('Missing Permissions'));
    await expect(status(page).getByRole('alert')).toHaveClass(/text-wipe/);
    expect(seen.tests).toBe(1);
  });

  test('a test sent too soon shows the wait sentence', async ({ page }) => {
    await openAccount(page, { initial: SAVED }, { testFailure: { status: 429, message: 'rate_limited' } });
    await page.getByTestId('discord-test').click();
    await expect(status(page)).toContainText(copy.testRateLimited);
  });

  test('turning an event off sends the events alone and says Saved.', async ({ page }) => {
    const seen = await openAccount(page, {
      initial: SAVED,
      afterPut: discordBlock({ events: { sim_finished: true, upload_parsed: false, upload_private: false } }),
    });
    await page.getByTestId('discord-event-upload').uncheck();
    await expect(status(page)).toHaveText(copy.savedEvents);
    expect(seen.puts).toEqual([
      { events: { sim_finished: true, upload_parsed: false, upload_private: false } },
    ]);
    expect(JSON.stringify(seen.puts)).not.toContain('webhook_url');
    await expect(premiumRow(page)).toContainText('On · 1 event');
    await expect(status(page)).toHaveCount(0, { timeout: SAVED_CLEAR_TIMEOUT_MS });
  });

  test('one to four failures show the warning with Discord’s last error', async ({ page }) => {
    await openAccount(page, { initial: discordBlock({ failures: 2, last_error: 'Unknown Webhook' }) });
    const { failures, lastError } = copy.warningLines(2, 'Unknown Webhook');
    await expect(panel(page)).toContainText(failures);
    await expect(panel(page)).toContainText(lastError ?? '');
  });

  test('the destination reads the webhook name and never the channel id', async ({ page }) => {
    await openAccount(page, { initial: SAVED });
    const destination = page.getByTestId('discord-destination');
    await expect(destination).toContainText(namedLine('Raid pings'));
    await expect(destination).not.toContainText('Channel 1198');
    await expect(destination).not.toContainText(CHANNEL_ID);
  });

  test('a webhook without a name reads the unnamed line', async ({ page }) => {
    await openAccount(page, { initial: discordBlock({ name: null }) });
    const destination = page.getByTestId('discord-destination');
    await expect(destination).toContainText(copy.destinationChannel);
    await expect(destination).not.toContainText('Posts through');
    await expect(destination).not.toContainText('Channel 1198');
  });

  test('a name with markup and markdown is plain text, cut at 40 characters', async ({ page }) => {
    const name = `<b>bold</b> [x](https://evil.example) **${'n'.repeat(60)}`;
    await openAccount(page, { initial: discordBlock({ name }) });
    const destination = page.getByTestId('discord-destination');
    await expect(destination).toContainText(namedLine(name));
    await expect(destination).toContainText('<b>bold</b> [x](https://evil.example)');
    await expect(destination).not.toContainText('n'.repeat(60));
    expect(await destination.locator('b, a, img').count()).toBe(0);
    const shown = copy.destinationNamed(name).name;
    expect([...shown.slice(1, -1)]).toHaveLength(DISCORD_NAME_LIMIT);
    expect(shown.endsWith('…”')).toBe(true);
  });

  test('the private row is disabled, value kept, while the log box is off', async ({ page }) => {
    const events = { sim_finished: true, upload_parsed: false, upload_private: true };
    const seen = await openAccount(page, { initial: discordBlock({ events }) });
    const privateBox = page.getByTestId('discord-event-private');
    await expect(privateBox).toBeDisabled();
    await expect(privateBox).toBeChecked();
    await expect(page.getByTestId('discord-event-upload')).toBeEnabled();
    await expect(panel(page)).toContainText(copy.eventUploadPrivate);
    await expect(panel(page)).toContainText(copy.eventsNote);
    expect(seen.puts).toHaveLength(0);
  });

  test('switching the log box on enables the private row with its value kept', async ({ page }) => {
    const events = { sim_finished: true, upload_parsed: false, upload_private: true };
    const seen = await openAccount(page, {
      initial: discordBlock({ events }),
      afterPut: discordBlock({ events: { ...events, upload_parsed: true } }),
    });
    await page.getByTestId('discord-event-upload').check();
    await expect(page.getByTestId('discord-event-private')).toBeEnabled();
    await expect(page.getByTestId('discord-event-private')).toBeChecked();
    expect(seen.puts).toEqual([{ events: { ...events, upload_parsed: true } }]);
  });

  test('turning private logs on sends all three event keys and nothing else', async ({ page }) => {
    const privateOn = { sim_finished: true, upload_parsed: true, upload_private: true };
    const seen = await openAccount(page, {
      initial: SAVED,
      afterPut: discordBlock({ events: privateOn }),
    });
    await expect(page.getByTestId('discord-event-private')).not.toBeChecked();
    await page.getByTestId('discord-event-private').check();
    await expect(status(page)).toHaveText(copy.savedEvents);
    expect(seen.puts).toEqual([{ events: privateOn }]);
    await expect(page.getByTestId('discord-event-private')).toBeChecked();
  });

  test('replacing opens the paste form on the input and Escape closes it', async ({ page }) => {
    await openAccount(page, { initial: SAVED });
    await page.getByTestId('discord-replace').click();
    await expect(input(page)).toBeFocused();
    await expect(page.getByTestId('discord-cancel')).toBeVisible();
    await page.keyboard.press('Escape');
    await expect(page.getByTestId('discord-destination')).toBeVisible();
  });

  test('Remove asks first, then deletes once and brings the paste form back', async ({ page }) => {
    const seen = await openAccount(page, { initial: SAVED });
    await page.getByTestId('discord-remove').click();
    await expect(panel(page)).toContainText(copy.removeQuestionPasted);
    await expect(page.getByTestId('discord-remove-cancel')).toBeFocused();
    await expect(page.getByTestId('discord-replace')).toHaveCount(0);
    await page.getByTestId('discord-remove-cancel').click();
    await expect(page.getByTestId('discord-replace')).toBeVisible();
    expect(seen.deletes).toBe(0);
    await page.getByTestId('discord-remove').click();
    await page.getByTestId('discord-remove-confirm').click();
    await expect(input(page)).toBeVisible();
    expect(seen.deletes).toBe(1);
    await expect(input(page)).toBeFocused();
    await expect(panel(page)).toContainText(copy.removed);
    await expect(premiumRow(page).getByText(premiumRowCopy.offSetUp, { exact: true })).toBeVisible();
  });
});

test.describe('switched off, lapsed, free', () => {
  test('five failures: the notice, the last error and one button that switches it back on', async ({
    page,
  }) => {
    const seen = await openAccount(page, { initial: FIVE_FAILURES, afterTest: SAVED });
    const notice = page.getByTestId('discord-notice');
    await expect(notice).toContainText(copy.switchedOff);
    await expect(notice).toContainText(copy.lastError('Unknown Webhook'));
    await expect(notice).toContainText(copy.switchedOffFix);
    await expect(page.getByTestId('discord-replace')).toContainText(copy.replace);
    await expect(page.getByRole('button', { name: copy.testAndSwitchOn })).toBeVisible();
    await expect(premiumRow(page)).toContainText(premiumRowCopy.offNeedsAttention);
    await expect(premiumRow(page)).not.toContainText('failed five times');
    await page.getByRole('button', { name: copy.testAndSwitchOn }).click();
    await expect(status(page)).toContainText(copy.testSentSwitchedOn);
    await expect(page.getByTestId('discord-notice')).toHaveCount(0);
    expect(seen.tests).toBe(1);
    await expect(premiumRow(page)).toContainText('On · 2 events');
  });

  test('five failures: Use a different webhook leaves while the remove question is open', async ({
    page,
  }) => {
    await openAccount(page, { initial: FIVE_FAILURES });
    await expect(page.getByTestId('discord-replace')).toBeVisible();
    await page.getByTestId('discord-remove').click();
    await expect(panel(page)).toContainText(copy.removeQuestionPasted);
    await expect(page.getByTestId('discord-replace')).toHaveCount(0);
    await page.getByTestId('discord-remove-cancel').click();
    await expect(page.getByTestId('discord-replace')).toBeVisible();
  });

  test('lapsed: the why, read-only boxes, Remove and nothing that posts', async ({ page }) => {
    await openAccount(page, { initial: SAVED, entitled: false });
    const notice = page.getByTestId('discord-notice');
    await expect(notice).toContainText(copy.lapsedTitle);
    await expect(notice).toContainText(copy.lapsedBody);
    await expect(page.getByTestId('discord-event-sim')).toBeDisabled();
    await expect(page.getByTestId('discord-event-upload')).toBeDisabled();
    await expect(page.getByTestId('discord-event-private')).toBeDisabled();
    await expect(page.getByTestId('discord-remove')).toBeVisible();
    await expect(page.getByTestId('discord-test')).toHaveCount(0);
    await expect(page.getByTestId('discord-replace')).toHaveCount(0);
    await expect(premiumRow(page)).toContainText('Paused · plan lapsed');
  });

  test('lapsed, then removed: focus lands on the plans link of the free sentence', async ({ page }) => {
    await openAccount(page, { initial: SAVED, entitled: false });
    await page.getByTestId('discord-remove').click();
    await page.getByTestId('discord-remove-confirm').click();
    await expect(page.getByTestId('discord-free-link')).toBeFocused();
  });

  test('a free member gets one sentence and a plans link, not the panel', async ({ page }) => {
    await openAccount(page, { initial: null, entitled: false });
    const free = page.getByTestId('discord-free');
    await expect(free).toContainText(copy.freeSentence);
    await expect(free.getByRole('link', { name: copy.seePlans })).toHaveAttribute('href', '/premium#notifications');
    await expect(page.getByTestId('discord-state-tag')).toHaveCount(0);
    await expect(page.getByTestId('account-discord').locator('form')).toHaveCount(0);
    await expect(premiumRow(page).getByText(premiumRowCopy.off, { exact: true })).toBeVisible();
    await expect(page.getByTestId('account-premium-notifications-link')).toHaveCount(0);
  });
});

test.describe('deployment with notifications off', () => {
  test('a payer sees no block and no Discord row', async ({ page }) => {
    await openAccount(page, { initial: null, featureOff: true });
    await expect(page.getByTestId('account')).toBeVisible();
    await expect(premiumRow(page)).toHaveCount(0);
    await expect(page.getByTestId('account-premium-notifications-link')).toHaveCount(0);
    await expect(page.locator('#discord')).toHaveCount(0);
    await expect(panel(page)).toHaveCount(0);
    await expect(page.getByTestId('discord-free')).toHaveCount(0);
  });
});

test.describe('Premium row', () => {
  test('reads On · 2 events and links to the block for a saved webhook', async ({ page }) => {
    await openAccount(page, { initial: SAVED });
    await expect(premiumRow(page)).toContainText('Discord notifications');
    const link = page.getByTestId('account-premium-notifications-link');
    await expect(link).toContainText('On · 2 events');
    await expect(link).toHaveAttribute('href', '#discord');
  });
});

declare global {
  interface Window {
    __skeletonSeen: boolean;
  }
}

/** From now on, notes whether the account page ever draws a loading skeleton again. */
// Flags a skeleton that APPEARS after the watcher starts: a flash during save or remove. One
// already on the page when the watcher starts (a line still loading) is not a flash.
async function watchForSkeleton(page: Page): Promise<void> {
  await page.evaluate(() => {
    window.__skeletonSeen = false;
    const present = new Set(document.querySelectorAll('[data-testid$="-skeleton"]'));
    new MutationObserver(() => {
      for (const node of document.querySelectorAll('[data-testid$="-skeleton"]')) {
        if (!present.has(node)) window.__skeletonSeen = true;
      }
    }).observe(document.body, { childList: true, subtree: true });
  });
}

test.describe('layout shift through save and remove', () => {
  for (const viewport of [PHONE, DESKTOP]) {
    test(`${viewport.width}px: nothing shifts beyond the panel's own height change`, async ({ page }) => {
      await recordShifts(page);
      await page.setViewportSize(viewport);
      await openAccount(page, { initial: null, afterPut: SAVED });
      const load = await drainShifts(page);
      await watchForSkeleton(page);
      await pasteAndSave(page, VALID_URL);
      await expect(page.getByTestId('discord-destination')).toBeVisible();
      const save = await drainShifts(page);
      await page.getByTestId('discord-test').click();
      await expect(status(page)).toContainText(copy.testSent);
      const testing = await drainShifts(page);
      await page.getByTestId('discord-remove').click();
      await page.getByTestId('discord-remove-confirm').click();
      await expect(input(page)).toBeVisible();
      const remove = await drainShifts(page);
      test.info().annotations.push({
        type: 'cls',
        description: `${viewport.width}px load ${sumShifts(load)} save ${sumShifts(save)} test ${sumShifts(testing)} remove ${sumShifts(remove)}`,
      });
      // The load shift is the page's own skeleton giving way to the account panels (the same
      // 0.45 at 390 for a free member); it is printed above, the member-caused ones are asserted.
      expect(await page.evaluate(() => window.__skeletonSeen)).toBe(false);
      expect(sumShifts(testing)).toBeLessThanOrEqual(MAX_SHIFT);
      for (const value of [...save, ...remove]) expect(value).toBeLessThanOrEqual(MAX_SHIFT);
    });
  }
});

interface AxeResults {
  violations: { id: string; nodes: { target: unknown }[] }[];
}

const AXE_SOURCE = createRequire(import.meta.url).resolve('axe-core/axe.min.js');

async function expectNoAxeViolations(page: Page): Promise<void> {
  await page.evaluate(() => Promise.all(document.getAnimations().map((animation) => animation.finished)));
  await page.addScriptTag({ path: AXE_SOURCE });
  const results = await page.evaluate(async (): Promise<AxeResults> => {
    const axe = (window as unknown as { axe: { run: (context: string) => Promise<AxeResults> } }).axe;
    return axe.run('[data-testid="account"]');
  });
  expect(
    results.violations.map((violation) => ({
      id: violation.id,
      targets: violation.nodes.map((node) => node.target),
    })),
  ).toEqual([]);
}

const CONNECTED = discordBlock({ source: 'connected', name: null, guild_id: GUILD_ID });
const ME_PATH = /\/v1\/me$/;
const SETTLE_MS = 750;

/** Visits /account with `search`, returning how many `/v1/me` reads the page made once settled. */
async function meReadsAfterVisit(page: Page, search: string): Promise<number> {
  let reads = 0;
  page.on('request', (request) => {
    if (ME_PATH.test(new URL(request.url()).pathname)) reads += 1;
  });
  await openAccount(page, { initial: CONNECTED, connect: true }, {}, `/account${search}#discord`);
  await expect(page.getByTestId('discord-test')).toBeVisible();
  await page.waitForTimeout(SETTLE_MS);
  return reads;
}

const CONNECT_HREF = CONNECT_URL_GLOB;
const CONNECT_ON = { initial: null, connect: true } satisfies DiscordMeStages;

test.describe('the connect state (connect lane)', () => {
  test('connect on: the gold link to the connect route, what Discord asks, no paste form', async ({
    page,
  }) => {
    await openAccount(page, CONNECT_ON);
    const connect = page.getByTestId('discord-connect');
    await expect(connect).toHaveText(copy.connect);
    await expect(connect).toHaveAttribute('href', CONNECT_HREF);
    expect(await connect.getAttribute('href')).toMatch(new RegExp(`^${E2E_API_BASE}/`));
    await expect(panel(page)).toContainText(copy.connectExplanation);
    await expect(page.getByTestId('discord-paste-instead')).toHaveText(copy.pasteInstead);
    await expect(input(page)).toHaveCount(0);
    await expect(page.getByTestId('discord-save')).toHaveCount(0);
  });

  test('the paste disclosure reveals the form with the way back under Save, and back again', async ({
    page,
  }) => {
    await openAccount(page, CONNECT_ON);
    await page.getByTestId('discord-paste-instead').click();
    await expect(input(page)).toBeVisible();
    const back = page.getByTestId('discord-connect-instead');
    await expect(back).toHaveText(copy.connectInstead);
    const saveBottom = await page.getByTestId('discord-save').boundingBox();
    const backTop = await back.boundingBox();
    expect((backTop?.y ?? 0) >= (saveBottom?.y ?? Infinity) + (saveBottom?.height ?? 0) - 1).toBe(true);
    await back.click();
    await expect(page.getByTestId('discord-connect')).toBeVisible();
    await expect(input(page)).toHaveCount(0);
  });

  test('connect off: the paste form alone, with no trace of the connect controls', async ({ page }) => {
    await openAccount(page, { initial: null, connect: false });
    await expect(input(page)).toBeVisible();
    for (const id of ['discord-connect', 'discord-connect-instead', 'discord-paste-instead']) {
      await expect(page.getByTestId(id)).toHaveCount(0);
    }
    expect(await page.content()).not.toContain(CONNECT_HREF);
  });

  test('denied: the status sentence, focus on the connect control, the flag gone from the address', async ({
    page,
  }) => {
    await openAccount(page, CONNECT_ON, {}, '/account?discord=denied#discord');
    await expect(panel(page).getByRole('status')).toHaveText(copy.returnDenied);
    await expect(page.getByTestId('discord-connect')).toBeFocused();
    await expect(page).toHaveURL(/\/account#discord$/);
    await page.reload();
    await expect(page.getByTestId('discord-connect')).toBeVisible();
    await expect(panel(page).getByRole('status')).toHaveCount(0);
  });
});

test.describe('a connected channel (connect lane)', () => {
  test('the destination, the Open in Discord link, the reconnect link with its note', async ({ page }) => {
    await openAccount(page, { initial: CONNECTED, connect: true });
    const destination = page.getByTestId('discord-destination');
    await expect(destination).toContainText(copy.destinationChannel);
    await expect(destination).not.toContainText('Channel 1198');
    await expect(destination).not.toContainText('Posts through');
    const open = destination.getByTestId('discord-open');
    await expect(open).toHaveText(copy.openInDiscord);
    await expect(open).toHaveAttribute('href', `https://discord.com/channels/${GUILD_ID}/${CHANNEL_ID}`);
    await expect(open).toHaveAttribute('target', '_blank');
    const reconnect = page.getByTestId('discord-replace');
    await expect(reconnect).toContainText(copy.reconnect);
    await expect(reconnect).toHaveAttribute('href', CONNECT_HREF);
    await expect(page.getByTestId('discord-replace-note')).toHaveText(copy.reconnectNote);
    const html = await page.content();
    expect(html).not.toContain(WEBHOOK_TOKEN);
    expect(html).not.toContain(WEBHOOK_ID);
    expect(html).not.toContain('discord.com/api/webhooks');
  });

  test('a connected webhook without a known server shows no Open in Discord link', async ({ page }) => {
    await openAccount(page, {
      initial: discordBlock({ source: 'connected', guild_id: null }),
      connect: true,
    });
    await expect(page.getByTestId('discord-destination')).toBeVisible();
    await expect(page.getByTestId('discord-open')).toHaveCount(0);
  });

  test('Remove asks the connected question, and deleting brings the connect state back', async ({ page }) => {
    const seen = await openAccount(page, { initial: CONNECTED, connect: true });
    await page.getByTestId('discord-remove').click();
    await expect(panel(page)).toContainText(copy.removeQuestionConnected);
    await expect(panel(page)).not.toContainText(copy.removeQuestionPasted);
    await expect(page.getByTestId('discord-replace')).toHaveCount(0);
    await page.getByTestId('discord-remove-confirm').click();
    await expect(page.getByTestId('discord-connect')).toBeVisible();
    expect(seen.deletes).toBe(1);
  });

  test('Remove asks the pasted question with connect on', async ({ page }) => {
    await openAccount(page, { initial: SAVED, connect: true });
    await page.getByTestId('discord-remove').click();
    await expect(panel(page)).toContainText(copy.removeQuestionPasted);
    await expect(panel(page)).not.toContainText(copy.removeQuestionConnected);
  });

  test('replacing a pasted webhook with connect on opens the connect view with Cancel', async ({ page }) => {
    await openAccount(page, { initial: SAVED, connect: true });
    await page.getByTestId('discord-replace').click();
    await expect(page.getByTestId('discord-connect')).toHaveAttribute('href', CONNECT_HREF);
    await expect(page.getByTestId('discord-cancel')).toBeVisible();
    await page.getByTestId('discord-cancel').click();
    await expect(page.getByTestId('discord-destination')).toBeVisible();
  });
});

test.describe('the return from Discord (connect lane)', () => {
  test('failed: the alert, focus on the connect control, the flag gone from the address', async ({
    page,
  }) => {
    await openAccount(page, { initial: null, connect: true }, {}, '/account?discord=failed#discord');
    await expect(page.getByRole('alert')).toHaveText(copy.returnFailed);
    await expect(page.getByTestId('discord-connect')).toBeFocused();
    await expect(page).toHaveURL(/\/account#discord$/);
    await page.reload();
    await expect(page.getByTestId('discord-connect')).toBeVisible();
    await expect(page.getByRole('alert')).toHaveCount(0);
  });

  test('connected: the status, focus on Send test message, one extra /v1/me read', async ({
    page,
    browser,
  }) => {
    const flagged = await meReadsAfterVisit(page, '?discord=connected');
    await expect(panel(page).getByRole('status')).toHaveText(copy.returnConnected);
    await expect(page.getByTestId('discord-test')).toBeFocused();
    await expect(page).toHaveURL(/\/account#discord$/);
    const plain = await meReadsAfterVisit(await browser.newPage(), '');
    expect(flagged).toBe(plain + 1);
  });
});

test.describe('axe', () => {
  // Axe samples colours mid-transition otherwise; reduced motion makes the panels settle at once.
  test.use({ reducedMotion: 'reduce' });
  const states: { name: string; stages: DiscordMeStages }[] = [
    { name: 'free', stages: { initial: null, entitled: false } },
    { name: 'paste form', stages: { initial: null } },
    { name: 'connect', stages: CONNECT_ON },
    { name: 'saved connected', stages: { initial: CONNECTED, connect: true } },
    { name: 'saved', stages: { initial: SAVED } },
    { name: 'switched off', stages: { initial: FIVE_FAILURES } },
    { name: 'lapsed', stages: { initial: SAVED, entitled: false } },
  ];
  for (const { name, stages } of states) {
    test(`${name} has no violations`, async ({ page }) => {
      await openAccount(page, stages);
      await expectNoAxeViolations(page);
    });
  }

  test('the invalid-URL error has no violations', async ({ page }) => {
    await openAccount(page, { initial: null });
    await pasteAndSave(page, INVALID_URL);
    await expect(page.getByTestId('discord-webhook-error')).toContainText(copy.invalidUrl);
    await expectNoAxeViolations(page);
  });
});
