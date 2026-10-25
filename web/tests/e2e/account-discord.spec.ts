// web/tests/e2e/account-discord.spec.ts
// The Discord block on the account page (experience spec 2026-10-25-discord-notifications, section
// 9): paste and save, test, event change, the five-failure and lapsed states, remove, the free
// sentence, the Premium row, layout shift through save and remove, and an axe pass per state.
// /v1/me and the three notification routes are stubbed; the fixture webhook URL and token appear
// only in the typed input and the PUT body the stub records.
import { createRequire } from 'node:module';
import { expect, test, type Page } from '@playwright/test';
import { notificationsCopy as copy } from '../../src/lib/account/notifications-copy';
import type { DiscordNotifications } from '../../src/lib/account/api';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import { gotoHydrated } from './support/hydrated';
import {
  discordBlock,
  notificationsKey,
  routeDiscordNotifications,
  type DiscordRouteOptions,
  type DiscordRoutes,
} from './support/discord-notifications';

const WEBHOOK_ID = '1098765432101234567';
const WEBHOOK_TOKEN = 'xT9fixtureTokenNotARealOne_abcdEFGH';
const VALID_URL = `https://discord.com/api/webhooks/${WEBHOOK_ID}/${WEBHOOK_TOKEN}`;
const INVALID_URL = 'https://example.com/api/webhooks/1/tok';
const BOTH_EVENTS = { sim_finished: true, upload_parsed: true };
const SAVED_CLEAR_TIMEOUT_MS = 6000;
const MAX_SHIFT = 0.02;
const PHONE = { width: 390, height: 844 };
const DESKTOP = { width: 1440, height: 1000 };
const TEST_URL_GLOB = '**/v1/notifications/discord/test';

type Discord = DiscordNotifications | null;

interface MeStages {
  entitled?: boolean;
  initial: Discord;
  /** What /v1/me answers once the PUT has been seen. */
  afterPut?: Discord;
  /** What /v1/me answers once a test has been sent. */
  afterTest?: Discord;
}

function envelope(data: unknown) {
  return {
    status: 200,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
  };
}

function currentBlock(stages: MeStages, seen: DiscordRoutes): Discord {
  if (seen.deletes > 0) return null;
  if (seen.tests > 0 && stages.afterTest !== undefined) return stages.afterTest;
  if (seen.puts.length > 0 && stages.afterPut !== undefined) return stages.afterPut;
  return stages.initial;
}

/** Signs a member in; /v1/me follows what the stubbed notification routes have been sent. */
async function signIn(page: Page, stages: MeStages, seen: DiscordRoutes): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
        characters: [],
        guilds: [],
        entitlements: mockEntitlements({ notifications: stages.entitled ?? true }),
        ...notificationsKey(currentBlock(stages, seen)),
      }),
    ),
  );
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
}

async function openAccount(
  page: Page,
  stages: MeStages,
  routeOptions: DiscordRouteOptions = {},
): Promise<DiscordRoutes> {
  const seen = await routeDiscordNotifications(page, routeOptions);
  await signIn(page, stages, seen);
  await gotoHydrated(page, '/account', 'account');
  return seen;
}

const panel = (page: Page) => page.getByTestId('account-discord');
const status = (page: Page) => page.getByTestId('discord-status');
const input = (page: Page) => page.locator('#discord-webhook');
const premiumRow = (page: Page) => page.getByTestId('account-premium-notifications');

async function pasteAndSave(page: Page, url: string): Promise<void> {
  await input(page).fill(url);
  await page.getByTestId('discord-save').click();
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
    await expect(page.getByTestId('discord-destination')).toContainText(
      `${copy.destinationNoun} ${WEBHOOK_ID}`,
    );
    expect(seen.puts).toEqual([{ webhook_url: VALID_URL, events: BOTH_EVENTS }]);
    await expect(status(page)).toContainText(copy.savedFirst);
    await expect(page.getByTestId('discord-test')).toBeFocused();
    const html = await page.content();
    expect(html).not.toContain(WEBHOOK_TOKEN);
    expect(html).not.toContain('discord.com/api/webhooks');
    await expect(premiumRow(page)).toContainText('On · 2 events');
  });

  test('a refusal for want of Premium shows the sentence and a plans link', async ({ page }) => {
    await openAccount(page, { initial: null }, { putFailure: { status: 402, message: 'premium_required' } });
    await pasteAndSave(page, VALID_URL);
    const error = page.getByTestId('discord-webhook-error');
    await expect(error).toContainText(copy.notEntitled);
    await expect(error.getByRole('link', { name: copy.seePlans })).toHaveAttribute('href', '/premium');
    await expect(input(page)).toBeFocused();
  });
});

test.describe('saved webhook', () => {
  test('a test that is sent, then a test Discord refuses', async ({ page }) => {
    const seen = await openAccount(page, { initial: SAVED });
    await page.getByTestId('discord-test').click();
    await expect(status(page)).toContainText(copy.testSent);
    expect(seen.tests).toBe(1);
    await page.unroute(TEST_URL_GLOB);
    await page.route(TEST_URL_GLOB, (route) => {
      seen.tests += 1;
      return route.fulfill({
        status: 502,
        contentType: 'application/json',
        body: JSON.stringify({
          ok: false,
          data: null,
          error: { message: 'Missing Permissions' },
          request_id: 'r',
        }),
      });
    });
    await page.getByTestId('discord-test').click();
    await expect(status(page)).toContainText(copy.testFailed('Missing Permissions'));
    await expect(status(page).getByRole('alert')).toHaveClass(/text-wipe/);
    expect(seen.tests).toBe(2);
  });

  test('turning an event off sends the events alone and says Saved.', async ({ page }) => {
    const seen = await openAccount(page, {
      initial: SAVED,
      afterPut: discordBlock({ events: { sim_finished: true, upload_parsed: false } }),
    });
    await page.getByTestId('discord-event-upload').uncheck();
    await expect(status(page)).toHaveText(copy.savedEvents);
    expect(seen.puts).toEqual([{ events: { sim_finished: true, upload_parsed: false } }]);
    expect(JSON.stringify(seen.puts)).not.toContain('webhook_url');
    await expect(premiumRow(page)).toContainText('On · 1 event');
    await expect(status(page)).toHaveText('', { timeout: SAVED_CLEAR_TIMEOUT_MS });
  });

  test('one to four failures show the warning with Discord’s last error', async ({ page }) => {
    await openAccount(page, { initial: discordBlock({ failures: 2, last_error: 'Unknown Webhook' }) });
    await expect(panel(page)).toContainText(copy.warning(2, 'Unknown Webhook'));
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
    await expect(panel(page)).toContainText(copy.removeQuestion);
    await expect(page.getByTestId('discord-remove-cancel')).toBeFocused();
    await page.getByTestId('discord-remove-cancel').click();
    expect(seen.deletes).toBe(0);
    await page.getByTestId('discord-remove').click();
    await page.getByTestId('discord-remove-confirm').click();
    await expect(input(page)).toBeVisible();
    expect(seen.deletes).toBe(1);
    await expect(input(page)).toBeFocused();
    await expect(panel(page)).toContainText(copy.removed);
    await expect(premiumRow(page)).toContainText('Off');
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
    await expect(panel(page)).toContainText(copy.switchedOffHelper);
    await expect(premiumRow(page)).toContainText('Off · failed five times');
    await page.getByRole('button', { name: copy.testAndSwitchOn }).click();
    await expect(status(page)).toContainText(copy.testSentSwitchedOn);
    await expect(page.getByTestId('discord-notice')).toHaveCount(0);
    expect(seen.tests).toBe(1);
    await expect(premiumRow(page)).toContainText('On · 2 events');
  });

  test('lapsed: the why, read-only boxes, Remove and nothing that posts', async ({ page }) => {
    await openAccount(page, { initial: SAVED, entitled: false });
    const notice = page.getByTestId('discord-notice');
    await expect(notice).toContainText(copy.lapsedTitle);
    await expect(notice).toContainText(copy.lapsedBody);
    await expect(page.getByTestId('discord-event-sim')).toBeDisabled();
    await expect(page.getByTestId('discord-event-upload')).toBeDisabled();
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
    await expect(free.getByRole('link', { name: copy.seePlans })).toHaveAttribute('href', '/premium');
    await expect(panel(page)).toHaveCount(0);
    await expect(premiumRow(page)).toContainText('Off');
    await expect(page.getByTestId('account-premium-notifications-link')).toHaveCount(0);
  });
});

test.describe('Premium row', () => {
  test('reads On · 2 events and links to the block for a saved webhook', async ({ page }) => {
    await openAccount(page, { initial: SAVED });
    await expect(premiumRow(page)).toContainText('Discord notifications');
    const link = page.getByTestId('account-premium-notifications-link');
    await expect(link).toHaveText('On · 2 events');
    await expect(link).toHaveAttribute('href', '#discord');
  });
});

interface LayoutShiftEntry extends PerformanceEntry {
  value: number;
  hadRecentInput: boolean;
}

declare global {
  interface Window {
    __shifts: number[];
    __skeletonSeen: boolean;
  }
}

async function recordShifts(page: Page): Promise<void> {
  await page.addInitScript(() => {
    window.__shifts = [];
    new PerformanceObserver((list) => {
      for (const entry of list.getEntries() as LayoutShiftEntry[]) {
        if (!entry.hadRecentInput) window.__shifts.push(entry.value);
      }
    }).observe({ type: 'layout-shift', buffered: true });
  });
}

/** The shifts reported since the last call, once the page has had two frames and a beat to settle. */
async function drainShifts(page: Page): Promise<number[]> {
  return page.evaluate(
    () =>
      new Promise<number[]>((resolve) => {
        requestAnimationFrame(() =>
          requestAnimationFrame(() => {
            setTimeout(() => resolve(window.__shifts.splice(0)), 300);
          }),
        );
      }),
  );
}

/** From now on, notes whether the account page ever draws a loading skeleton again. */
async function watchForSkeleton(page: Page): Promise<void> {
  await page.evaluate(() => {
    window.__skeletonSeen = false;
    new MutationObserver(() => {
      if (document.querySelector('[data-testid$="-skeleton"]') !== null) window.__skeletonSeen = true;
    }).observe(document.body, { childList: true, subtree: true });
  });
}

const sum = (values: number[]): number => values.reduce((total, value) => total + value, 0);

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
        description: `${viewport.width}px load ${sum(load)} save ${sum(save)} test ${sum(testing)} remove ${sum(remove)}`,
      });
      // The load shift is the page's own skeleton giving way to the account panels (the same
      // 0.45 at 390 for a free member); it is printed above, the member-caused ones are asserted.
      expect(await page.evaluate(() => window.__skeletonSeen)).toBe(false);
      expect(sum(testing)).toBeLessThanOrEqual(MAX_SHIFT);
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

test.describe('axe', () => {
  test.use({ reducedMotion: 'reduce' });
  const states: { name: string; stages: MeStages }[] = [
    { name: 'free', stages: { initial: null, entitled: false } },
    { name: 'paste form', stages: { initial: null } },
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
