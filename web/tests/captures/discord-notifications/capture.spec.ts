// web/tests/captures/discord-notifications/capture.spec.ts
// Acceptance screenshots for the Discord notifications lane (experience spec
// design/specs/2026-10-25-discord-notifications.md, section 9). Not a test of behaviour: it drives
// the account page, with /v1/me and the notification routes stubbed, into every state and writes
// PNGs to design/reviews/2026-10-25-discord-notifications/. Item 9 (real Discord messages) is made
// by hand; item 13 (side-by-side with the mock) by compose_side_by_side.py in the same folder.
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Locator, type Page } from '@playwright/test';
import type { DiscordNotifications } from '../../../src/lib/account/api';
import { mockEntitlements } from '../../../src/test-support/billing-fixtures';
import { gotoHydrated } from '../../e2e/support/hydrated';
import {
  discordBlock,
  notificationsKey,
  routeDiscordNotifications,
  type DiscordRouteOptions,
  type DiscordRoutes,
} from '../../e2e/support/discord-notifications';

const OUT_DIR = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../../design/reviews/2026-10-25-discord-notifications',
);
const PAD = 24;
const VIEWPORT_HEIGHT = 1000;
const SETTLE_MS = 400;
const WIDTH_DESKTOP = 1440;
const BASE_WIDTHS = [390, WIDTH_DESKTOP];
const NARROW_WIDTHS = [360, ...BASE_WIDTHS];
const ALL_WIDTHS = [360, 390, 768, 1024, WIDTH_DESKTOP, 2000];

const VALID_URL = 'https://discord.com/api/webhooks/1087654321012345678/xT9fixtureTokenNotARealOne_abcdEFGH';
const INVALID_URL = 'https://example.com/api/webhooks/1/tok';
const LONG_ERROR =
  'Invalid Webhook Token: the token in this webhook URL was changed, or the webhook was deleted from the channel settings by someone with Manage Webhooks, so Discord no longer accepts it';

const SAVED = discordBlock();
const WARNING = discordBlock({ failures: 2, last_error: 'Unknown Webhook' });
const WARNING_LONG = discordBlock({ failures: 4, last_error: LONG_ERROR });
const FIVE_FAILURES = discordBlock({
  disabled_reason: 'failed five times',
  last_error: 'Unknown Webhook',
  failures: 5,
});
const FIVE_FAILURES_LONG = discordBlock({
  disabled_reason: 'failed five times',
  last_error: LONG_ERROR,
  failures: 5,
});
const ONE_EVENT = discordBlock({
  events: { sim_finished: true, upload_parsed: false, upload_private: false },
});
const NO_EVENTS = discordBlock({
  events: { sim_finished: false, upload_parsed: false, upload_private: false },
});

interface Stages {
  entitled?: boolean;
  featureOff?: boolean;
  initial: DiscordNotifications | null;
  afterPut?: DiscordNotifications | null;
  afterTest?: DiscordNotifications | null;
  /** Holds /v1/me until this resolves (the pre-paint skeleton). */
  holdMe?: Promise<void>;
}

function envelope(data: unknown) {
  return {
    status: 200,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
  };
}

function currentBlock(stages: Stages, seen: DiscordRoutes): DiscordNotifications | null {
  if (seen.deletes > 0) return null;
  if (seen.tests > 0 && stages.afterTest !== undefined) return stages.afterTest;
  if (seen.puts.length > 0 && stages.afterPut !== undefined) return stages.afterPut;
  return stages.initial;
}

async function signIn(page: Page, stages: Stages, seen: DiscordRoutes): Promise<void> {
  await page.route('**/v1/me', async (route) => {
    if (stages.holdMe !== undefined) await stages.holdMe;
    await route.fulfill(
      envelope({
        user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
        characters: [],
        guilds: [],
        entitlements: mockEntitlements({ notifications: stages.entitled ?? true }),
        ...(stages.featureOff === true ? {} : notificationsKey(currentBlock(stages, seen))),
      }),
    );
  });
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
}

/** A promise a spec resolves by hand, to hold a stubbed route until the frame is captured. */
function gate(): { wait: Promise<void>; release: () => void } {
  let release = (): void => undefined;
  const wait = new Promise<void>((resolve) => {
    release = resolve;
  });
  return { wait, release };
}

/** Holds every request matching `glob` until released, then hands it to the stub registered earlier. */
async function holdRoute(page: Page, glob: string, wait: Promise<void>): Promise<void> {
  await page.route(glob, async (route) => {
    await wait;
    await route.fallback();
  });
}

const settle = (page: Page): Promise<void> => page.waitForTimeout(SETTLE_MS);
const panel = (page: Page): Locator => page.getByTestId('account-discord');
const input = (page: Page): Locator => page.locator('#discord-webhook');
const premiumRow = (page: Page): Locator => page.getByTestId('account-premium-notifications');
const target = (name: string): string => path.join(OUT_DIR, name);

async function openAccount(
  page: Page,
  width: number,
  stages: Stages,
  routeOptions: DiscordRouteOptions = {},
): Promise<DiscordRoutes> {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  const seen = await routeDiscordNotifications(page, routeOptions);
  await signIn(page, stages, seen);
  await gotoHydrated(page, '/account', 'account');
  await settle(page);
  return seen;
}

async function full(page: Page, name: string, width: number): Promise<void> {
  await page.screenshot({ path: target(`discord-${name}-${width}.png`), fullPage: true });
}

/** The element's page rectangle grown by PAD and kept inside the page. */
async function cropOf(page: Page, element: Locator, file: string): Promise<void> {
  const rect = await element.evaluate((node) => {
    const box = node.getBoundingClientRect();
    return {
      x: box.x + window.scrollX,
      y: box.y + window.scrollY,
      width: box.width,
      height: box.height,
      pageWidth: document.documentElement.clientWidth,
    };
  });
  const x = Math.max(0, rect.x - PAD);
  const y = Math.max(0, rect.y - PAD);
  await page.screenshot({
    path: target(file),
    fullPage: true,
    clip: {
      x,
      y,
      width: Math.min(rect.width + PAD * 2, rect.pageWidth - x),
      height: rect.height + PAD * 2,
    },
  });
}

/** A full-page capture and a padded crop of `subject` (the Discord panel by default). */
async function shoot(page: Page, name: string, width: number, subject: Locator = panel(page)): Promise<void> {
  await full(page, name, width);
  await cropOf(page, subject, `discord-${name}-panel-${width}.png`);
}

/** One test per width, each on a fresh page. */
function atWidths(title: string, widths: number[], body: (page: Page, width: number) => Promise<void>): void {
  for (const width of widths) {
    test(`${title} @${width}`, ({ page }) => body(page, width));
  }
}

/** A key press first, so `:focus-visible` matches on the programmatic focus that follows. */
async function focusWithKeyboard(page: Page, control: Locator): Promise<void> {
  await page.keyboard.press('Shift');
  await control.focus();
  await settle(page);
}

// 1 and 11. The rail with a saved webhook, at every named width.
atWidths('saved', ALL_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: SAVED });
  await shoot(page, 'saved', width);
});

// 2. The paste form: empty and focused, the invalid-URL message, saving, the Premium refusal.
atWidths('paste form empty and focused', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: null });
  await focusWithKeyboard(page, input(page));
  await shoot(page, 'paste-empty', width);
});

atWidths('paste form invalid URL', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: null });
  await input(page).fill(INVALID_URL);
  await page.getByTestId('discord-save').click();
  await page.getByTestId('discord-webhook-error').waitFor();
  await settle(page);
  await shoot(page, 'paste-invalid', width);
});

atWidths('paste form saving', BASE_WIDTHS, async (page, width) => {
  const held = gate();
  await openAccount(page, width, { initial: null, afterPut: SAVED });
  // Registered after the stub, so it runs first and holds the PUT until the frame is taken.
  await holdRoute(page, '**/v1/notifications/discord', held.wait);
  await input(page).fill(VALID_URL);
  await page.getByTestId('discord-save').click();
  await page.getByText('Saving…').waitFor();
  await shoot(page, 'saving', width);
  held.release();
});

atWidths('paste form after the Premium refusal', BASE_WIDTHS, async (page, width) => {
  await openAccount(
    page,
    width,
    { initial: null },
    { putFailure: { status: 402, message: 'premium_required' } },
  );
  await input(page).fill(VALID_URL);
  await page.getByTestId('discord-save').click();
  await page.getByTestId('discord-webhook-error').waitFor();
  await settle(page);
  await shoot(page, 'paste-not-entitled', width);
});

// 3. Saved: just saved, sending, sent, Discord error, rate limit, event change, warnings.
atWidths('just saved', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: null, afterPut: SAVED });
  await input(page).fill(VALID_URL);
  await page.getByTestId('discord-save').click();
  await page.getByTestId('discord-destination').waitFor();
  await settle(page);
  await shoot(page, 'saved-just-now', width);
});

atWidths('sending', BASE_WIDTHS, async (page, width) => {
  const held = gate();
  await openAccount(page, width, { initial: SAVED });
  // Registered after the stub, so it runs first and holds the POST until the frame is taken.
  await holdRoute(page, '**/v1/notifications/discord/test', held.wait);
  await page.getByTestId('discord-test').click();
  await page.getByText('Sending…').waitFor();
  await shoot(page, 'sending', width);
  held.release();
});

atWidths('test sent', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: SAVED });
  await page.getByTestId('discord-test').click();
  await page.getByText('Sent. Check the channel.').waitFor();
  await shoot(page, 'test-sent', width);
});

atWidths('test failed', BASE_WIDTHS, async (page, width) => {
  await openAccount(
    page,
    width,
    { initial: SAVED },
    { testFailure: { status: 502, message: 'Missing Permissions' } },
  );
  await page.getByTestId('discord-test').click();
  await page.getByText('Discord said:').waitFor();
  await shoot(page, 'test-failed', width);
});

atWidths('test rate limited', BASE_WIDTHS, async (page, width) => {
  await openAccount(
    page,
    width,
    { initial: SAVED },
    { testFailure: { status: 429, message: 'rate_limited' } },
  );
  await page.getByTestId('discord-test').click();
  await page.getByText('Wait a minute').waitFor();
  await shoot(page, 'test-rate-limited', width);
});

atWidths('event turned off', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: SAVED, afterPut: ONE_EVENT });
  await page.getByTestId('discord-event-upload').uncheck();
  await page.getByTestId('discord-status').getByText('Saved.').waitFor();
  await shoot(page, 'events-saved', width);
});

atWidths('one to four failures', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: WARNING });
  await shoot(page, 'warning', width);
});

atWidths('four failures, a long Discord error', NARROW_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: WARNING_LONG });
  await shoot(page, 'warning-long', width);
});

// 4. Switched off after five failures; after Test and switch on.
atWidths('switched off', NARROW_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: FIVE_FAILURES });
  await shoot(page, 'switched-off', width);
});

atWidths('switched off, a long Discord error', NARROW_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: FIVE_FAILURES_LONG });
  await shoot(page, 'switched-off-long', width);
});

atWidths('switched back on', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: FIVE_FAILURES, afterTest: SAVED });
  await page.getByRole('button', { name: 'Test and switch on' }).click();
  await page.getByText('Sent. Posts are on again.').waitFor();
  await shoot(page, 'switched-on-again', width);
});

// 5. Plan lapsed.
atWidths('plan lapsed', NARROW_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: SAVED, entitled: false });
  await shoot(page, 'lapsed', width);
});

// 6. Remove confirm, removed, and the replace form.
atWidths('remove confirm', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: SAVED });
  await page.getByTestId('discord-remove').click();
  await page.getByTestId('discord-remove-cancel').waitFor();
  await settle(page);
  await shoot(page, 'remove-confirm', width);
});

atWidths('removed', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: SAVED });
  await page.getByTestId('discord-remove').click();
  await page.getByTestId('discord-remove-confirm').click();
  await page.getByText('Removed.').waitFor();
  await shoot(page, 'removed', width);
});

atWidths('replace form', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: SAVED });
  await page.getByTestId('discord-replace').click();
  await input(page).waitFor();
  await settle(page);
  await shoot(page, 'replace', width);
});

// 7. Free member, and a deployment with notifications off.
atWidths('free member', ALL_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: null, entitled: false });
  await shoot(page, 'free', width, page.getByTestId('discord-free'));
});

atWidths('deployment with notifications off', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: null, featureOff: true });
  await full(page, 'feature-off', width);
});

// 8. Focus-visible on every control, and hover on the three buttons, at 1440 (panel crops).
test.describe('focus and hover at 1440', () => {
  const focusOn = async (
    page: Page,
    stages: Stages,
    control: (page: Page) => Locator,
    name: string,
    subject: (page: Page) => Locator = panel,
  ): Promise<void> => {
    await openAccount(page, WIDTH_DESKTOP, stages);
    await focusWithKeyboard(page, control(page));
    await cropOf(page, subject(page), `discord-focus-${name}-${WIDTH_DESKTOP}.png`);
  };
  const byId = (id: string) => (page: Page) => page.getByTestId(id);

  test('input', ({ page }) => focusOn(page, { initial: null }, input, 'input'));
  test('save', ({ page }) => focusOn(page, { initial: null }, byId('discord-save'), 'save'));
  test('paste form checkbox sim', ({ page }) =>
    focusOn(page, { initial: null }, byId('discord-event-sim'), 'checkbox-sim-paste'));
  test('checkbox sim', ({ page }) =>
    focusOn(page, { initial: SAVED }, byId('discord-event-sim'), 'checkbox-sim'));
  test('checkbox upload', ({ page }) =>
    focusOn(page, { initial: SAVED }, byId('discord-event-upload'), 'checkbox-upload'));
  test('test', ({ page }) => focusOn(page, { initial: SAVED }, byId('discord-test'), 'test'));
  test('remove', ({ page }) => focusOn(page, { initial: SAVED }, byId('discord-remove'), 'remove'));
  test('replace', ({ page }) => focusOn(page, { initial: SAVED }, byId('discord-replace'), 'replace'));
  test('see plans', ({ page }) =>
    focusOn(
      page,
      { initial: null, entitled: false },
      byId('discord-free-link'),
      'see-plans',
      byId('discord-free'),
    ));

  const hoverOn = async (page: Page, stages: Stages, id: string, name: string): Promise<void> => {
    await openAccount(page, WIDTH_DESKTOP, stages);
    await page.getByTestId(id).hover();
    await settle(page);
    await cropOf(page, panel(page), `discord-hover-${name}-${WIDTH_DESKTOP}.png`);
  };

  test('hover save', ({ page }) => hoverOn(page, { initial: null }, 'discord-save', 'save'));
  test('hover test', ({ page }) => hoverOn(page, { initial: SAVED }, 'discord-test', 'test'));
  test('hover remove', ({ page }) => hoverOn(page, { initial: SAVED }, 'discord-remove', 'remove'));
});

// 10. The Premium row's seven values.
const ROW_VALUES: { value: string; stages: Stages }[] = [
  { value: 'off', stages: { initial: null, entitled: false } },
  { value: 'off-set-up', stages: { initial: null } },
  { value: 'off-no-events', stages: { initial: NO_EVENTS } },
  { value: 'off-failed-five-times', stages: { initial: FIVE_FAILURES } },
  { value: 'paused-plan-lapsed', stages: { initial: SAVED, entitled: false } },
  { value: 'on-1-event', stages: { initial: ONE_EVENT } },
  { value: 'on-2-events', stages: { initial: SAVED } },
];

for (const { value, stages } of ROW_VALUES) {
  atWidths(`premium row ${value}`, BASE_WIDTHS, async (page, width) => {
    await openAccount(page, width, stages);
    await cropOf(page, premiumRow(page), `discord-row-${value}-${width}.png`);
  });
}

// The account page's loading skeleton before /v1/me answers: the reserve gap the lane deferred.
atWidths('account loading skeleton', BASE_WIDTHS, async (page, width) => {
  const held = gate();
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  // The readable half of the session cookie is the hint that makes the page draw its skeletons.
  await page.context().addCookies([{ name: 'fs_csrf', value: 'capture', url: 'http://localhost:4411' }]);
  await signIn(page, { initial: SAVED, holdMe: held.wait }, { puts: [], tests: 0, deletes: 0 });
  await page.goto('/account');
  await page.getByTestId('account-characters-skeleton').waitFor();
  await settle(page);
  await full(page, 'skeleton', width);
  held.release();
});
