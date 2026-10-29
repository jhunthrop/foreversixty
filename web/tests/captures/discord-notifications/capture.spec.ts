// web/tests/captures/discord-notifications/capture.spec.ts
// Acceptance screenshots for the Discord notifications lane (experience spec
// design/specs/2026-10-25-discord-notifications.md, section 9). Not a test of behaviour: it drives
// the account page, with /v1/me and the notification routes stubbed, into every state and writes
// PNGs to design/reviews/2026-10-25-discord-notifications/. Item 9 (real Discord messages) is made
// by hand; item 13 (side-by-side with the mock) by compose_side_by_side.py in the same folder.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test, type Locator, type Page } from '@playwright/test';
import type { DiscordNotifications } from '../../../src/lib/account/api';
import type { HintState } from '../../../src/lib/account/layout';
import {
  addSessionHint,
  createHold,
  measureHeldLoad,
  seedWarm,
  stubMember,
} from '../../e2e/support/account-reserve';
import { gotoHydrated } from '../../e2e/support/hydrated';
import {
  discordBlock,
  GUILD_ID,
  routeDiscordNotifications,
  signInMember,
  type DiscordMeStages,
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
  stages: DiscordMeStages,
  routeOptions: DiscordRouteOptions = {},
  route = '/account',
): Promise<DiscordRoutes> {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  const seen = await routeDiscordNotifications(page, routeOptions);
  await signInMember(page, stages, seen);
  await gotoHydrated(page, route, 'account');
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
function atWidths(
  title: string,
  widths: number[],
  body: (page: Page, width: number, baseURL: string | undefined) => Promise<void>,
): void {
  for (const width of widths) {
    test(`${title} @${width}`, ({ page, baseURL }) => body(page, width, baseURL));
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
    stages: DiscordMeStages,
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

  const hoverOn = async (page: Page, stages: DiscordMeStages, id: string, name: string): Promise<void> => {
    await openAccount(page, WIDTH_DESKTOP, stages);
    await page.getByTestId(id).hover();
    await settle(page);
    await cropOf(page, panel(page), `discord-hover-${name}-${WIDTH_DESKTOP}.png`);
  };

  test('hover save', ({ page }) => hoverOn(page, { initial: null }, 'discord-save', 'save'));
  test('hover test', ({ page }) => hoverOn(page, { initial: SAVED }, 'discord-test', 'test'));
  test('hover remove', ({ page }) => hoverOn(page, { initial: SAVED }, 'discord-remove', 'remove'));
});

// 10 and 20. The Premium row's seven values; 1024 proves `2 events` never splits.
const ROW_WIDTHS = [390, 1024, WIDTH_DESKTOP];
const ROW_VALUES: { value: string; stages: DiscordMeStages }[] = [
  { value: 'off', stages: { initial: null, entitled: false } },
  { value: 'off-set-up', stages: { initial: null } },
  { value: 'off-no-events', stages: { initial: NO_EVENTS } },
  { value: 'off-needs-attention', stages: { initial: FIVE_FAILURES } },
  { value: 'paused-plan-lapsed', stages: { initial: SAVED, entitled: false } },
  { value: 'on-1-event', stages: { initial: ONE_EVENT } },
  { value: 'on-2-events', stages: { initial: SAVED } },
];

for (const { value, stages } of ROW_VALUES) {
  atWidths(`premium row ${value}`, ROW_WIDTHS, async (page, width) => {
    await openAccount(page, width, stages);
    await cropOf(page, premiumRow(page), `discord-row-${value}-${width}.png`);
  });
}

// 15. The destination: named, unnamed, a name cut at 40 characters, and a name that is only text.
const LONG_NAME = 'Raid pings for the Thursday progression night with the whole guild';
const MARKUP_NAME = '@everyone <b>x</b>';
const DESTINATIONS: { value: string; block: DiscordNotifications }[] = [
  { value: 'named', block: SAVED },
  { value: 'unnamed', block: discordBlock({ name: null }) },
  { value: 'name-long', block: discordBlock({ name: LONG_NAME }) },
  { value: 'name-markup', block: discordBlock({ name: MARKUP_NAME }) },
];

for (const { value, block } of DESTINATIONS) {
  atWidths(`destination ${value}`, BASE_WIDTHS, async (page, width) => {
    await openAccount(page, width, { initial: block });
    await cropOf(page, page.getByTestId('discord-destination'), `discord-destination-${value}-${width}.png`);
    await shoot(page, `saved-${value}`, width);
  });
}

// 14. The paste form's steps and the distance from the input's bottom to the trust line.
atWidths('paste form steps and the dead band', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: null });
  await page.getByTestId('discord-step-3').waitFor();
  await shoot(page, 'paste-steps', width);
});

// 16. The three event rows: third on, third off, second off (third disabled), lapsed (all disabled).
const EVENT_ROWS: { value: string; stages: DiscordMeStages }[] = [
  {
    value: 'private-on',
    stages: {
      initial: discordBlock({ events: { sim_finished: true, upload_parsed: true, upload_private: true } }),
    },
  },
  { value: 'private-off', stages: { initial: SAVED } },
  {
    value: 'logs-off-private-disabled',
    stages: {
      initial: discordBlock({ events: { sim_finished: true, upload_parsed: false, upload_private: true } }),
    },
  },
  { value: 'lapsed-all-disabled', stages: { initial: SAVED, entitled: false } },
];

for (const { value, stages } of EVENT_ROWS) {
  atWidths(`event rows ${value}`, BASE_WIDTHS, async (page, width) => {
    await openAccount(page, width, stages);
    await shoot(page, `events-${value}`, width);
  });
}

// 17. The warning: one failure, two failures, four failures with a 140-character error.
const WARNINGS: { value: string; block: DiscordNotifications }[] = [
  { value: 'one', block: discordBlock({ failures: 1, last_error: 'Unknown Webhook' }) },
  { value: 'two', block: WARNING },
  { value: 'four-long', block: WARNING_LONG },
];

for (const { value, block } of WARNINGS) {
  atWidths(`warning ${value}`, BASE_WIDTHS, async (page, width) => {
    await openAccount(page, width, { initial: block });
    await shoot(page, `warning-${value}`, width);
  });
}

// 18 and 19. The four-line switched-off notice with its heights, and the remove question with no replace link.
atWidths('switched off notice, remove question open', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: FIVE_FAILURES });
  await page.getByTestId('discord-remove').click();
  await page.getByTestId('discord-remove-cancel').waitFor();
  await settle(page);
  await shoot(page, 'switched-off-remove-confirm', width);
});

for (const [name, block] of [
  ['switched-off', FIVE_FAILURES],
  ['switched-off-long', FIVE_FAILURES_LONG],
] as const) {
  atWidths(`panel height ${name}`, NARROW_WIDTHS, async (page, width) => {
    await openAccount(page, width, { initial: block });
    const height = (await panel(page).boundingBox())?.height ?? 0;
    test.info().annotations.push({ type: 'height', description: `${name} @${width}: ${height} px` });
  });
}

// 21. The loading reserve: the loading frame beside the ready page for each `data-discord` state, warm
// (the cached /v1/me and the stored heights) at both widths, and cold (nothing stored) once per width.
const RESERVE_STATES: { state: HintState; stages: Omit<DiscordMeStages, 'hold'> }[] = [
  { state: 'none', stages: { initial: null, featureOff: true } },
  { state: 'free', stages: { initial: null, entitled: false } },
  { state: 'connect', stages: { initial: null, connect: true } },
  { state: 'form', stages: { initial: null } },
  { state: 'saved', stages: { initial: SAVED } },
  { state: 'off', stages: { initial: FIVE_FAILURES } },
  { state: 'lapsed', stages: { initial: SAVED, entitled: false } },
];

async function captureLoadingReserve(
  page: Page,
  baseURL: string | undefined,
  label: string,
  width: number,
  stages: Omit<DiscordMeStages, 'hold'>,
  warm: boolean,
): Promise<void> {
  const hold = createHold();
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await addSessionHint(page, baseURL);
  await stubMember(page, stages, hold);
  if (warm) await seedWarm(page);
  const measure = await measureHeldLoad(page, hold, async () => {
    await full(page, `loading-${label}`, width);
  });
  await settle(page);
  await full(page, `loading-${label}-ready`, width);
  test.info().annotations.push({
    type: warm ? 'warm' : 'cold',
    description: `${label} @${width}: cls ${measure.cls.toFixed(4)}, footer ${measure.footerLoading} -> ${measure.footerReady}`,
  });
}

for (const { state, stages } of RESERVE_STATES) {
  atWidths(`loading reserve warm ${state}`, BASE_WIDTHS, (page, width, baseURL) =>
    captureLoadingReserve(page, baseURL, `warm-${state}`, width, stages, true),
  );
}

atWidths('loading reserve cold', BASE_WIDTHS, (page, width, baseURL) =>
  captureLoadingReserve(page, baseURL, 'cold', width, { initial: null }, false),
);

// 22. Gold checkboxes: the Discord row checked, disabled and focus-visible, and the rankings filter.
atWidths('gold checkbox, Discord row', [WIDTH_DESKTOP], async (page, width) => {
  await openAccount(page, width, {
    initial: discordBlock({ events: { sim_finished: true, upload_parsed: false, upload_private: true } }),
  });
  await cropOf(page, panel(page), `discord-checkbox-checked-disabled-${width}.png`);
  await focusWithKeyboard(page, page.getByTestId('discord-event-sim'));
  await cropOf(page, panel(page), `discord-checkbox-focus-${width}.png`);
});

// 24 to 28. The connect lane (experience spec section 3.2, 4.3 and 9, amended 2026-10-25).
const CONNECT_ON: DiscordMeStages = { initial: null, connect: true };
const CONNECTED = discordBlock({ source: 'connected', name: null, guild_id: GUILD_ID });
const CONNECT_CAP_PX: Readonly<Record<number, number>> = { 360: 340, 390: 340, [WIDTH_DESKTOP]: 320 };
async function annotateHeight(page: Page, what: string): Promise<void> {
  const height = (await panel(page).boundingBox())?.height ?? 0;
  test.info().annotations.push({ type: 'height', description: `${what}: ${height} px` });
}
const connectControl = (page: Page): Locator => page.getByTestId('discord-connect');

// 24. The Connect state: default at 360, 390 and 1440 with the height against the cap, then
// hover, focus-visible and `Opening Discord…` at 1440.
atWidths('connect state', NARROW_WIDTHS, async (page, width) => {
  await openAccount(page, width, CONNECT_ON);
  await shoot(page, 'connect', width);
  const height = (await panel(page).boundingBox())?.height ?? 0;
  test.info().annotations.push({
    type: 'height',
    description: `connect @${width}: ${height} px, cap ${CONNECT_CAP_PX[width]}`,
  });
});

test('connect hover @1440', async ({ page }) => {
  await openAccount(page, WIDTH_DESKTOP, CONNECT_ON);
  await connectControl(page).hover();
  await settle(page);
  await cropOf(page, panel(page), `discord-connect-hover-panel-${WIDTH_DESKTOP}.png`);
});

test('connect focus-visible @1440', async ({ page }) => {
  await openAccount(page, WIDTH_DESKTOP, CONNECT_ON);
  await focusWithKeyboard(page, connectControl(page));
  await cropOf(page, panel(page), `discord-connect-focus-panel-${WIDTH_DESKTOP}.png`);
});

atWidths('connect opening Discord', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, CONNECT_ON);
  // A window listener runs after the component's own click handler and cancels the navigation, so
  // the link stays on its `Opening Discord…` frame for the capture.
  await page.evaluate(() => window.addEventListener('click', (event) => event.preventDefault()));
  await connectControl(page).click();
  await page.getByText('Opening Discord…').waitFor();
  await shoot(page, 'connect-opening', width);
});

// 25. The paste view with connect on (the back link under Save), and paste-first with connect off.
atWidths('paste view with connect on', NARROW_WIDTHS, async (page, width) => {
  await openAccount(page, width, CONNECT_ON);
  await page.getByTestId('discord-paste-instead').click();
  await input(page).waitFor();
  await settle(page);
  await shoot(page, 'connect-paste-view', width);
  await annotateHeight(page, `paste view with back link @${width}`);
});

atWidths('paste first with connect off', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: null, connect: false });
  await shoot(page, 'connect-off-paste-first', width);
});

// 26. The three return outcomes, each from `/account?discord=<flag>#discord`; the URL is asserted
// clean (a PNG cannot show the address bar), and a reload shows no status line.
const RETURNS: { flag: string; stages: DiscordMeStages }[] = [
  { flag: 'connected', stages: { initial: CONNECTED, connect: true } },
  { flag: 'denied', stages: CONNECT_ON },
  { flag: 'failed', stages: CONNECT_ON },
];

for (const { flag, stages } of RETURNS) {
  atWidths(`return ${flag}`, BASE_WIDTHS, async (page, width) => {
    await openAccount(page, width, stages, {}, `/account?discord=${flag}#discord`);
    await expect(page).toHaveURL(/\/account#discord$/);
    if (flag === 'connected') await expect(page.getByTestId('discord-test')).toBeFocused();
    await shoot(page, `return-${flag}`, width);
    await page.reload();
    await settle(page);
    await expect(panel(page).getByRole('status')).toHaveCount(0);
    await shoot(page, `return-${flag}-reloaded`, width);
  });
}

// 27. Saved connected: the destination with the inline link, the reconnect link and its note at 360
// and 390, and a saved pasted webhook beside it.
atWidths('saved connected', NARROW_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: CONNECTED, connect: true });
  await shoot(page, 'connected-saved', width);
  await annotateHeight(page, `saved connected @${width}`);
  await cropOf(page, page.getByTestId('discord-destination'), `discord-connected-destination-${width}.png`);
});

// The switched-off connected webhook: the notice holds `Connect a different channel` and its note.
const CONNECTED_SWITCHED_OFF = discordBlock({
  source: 'connected',
  name: null,
  guild_id: GUILD_ID,
  disabled_reason: 'failed five times',
  last_error: 'Unknown Webhook',
  failures: 5,
});

atWidths('switched off connected', NARROW_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: CONNECTED_SWITCHED_OFF, connect: true });
  await shoot(page, 'connected-switched-off', width);
  await annotateHeight(page, `switched-off connected @${width}`);
});

atWidths('saved pasted beside connected', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: SAVED, connect: true });
  await shoot(page, 'connected-pasted-beside', width);
});

// 28. Both remove questions, the replace view of a pasted webhook with connect on, and the reconnect view.
for (const [name, block] of [
  ['connected', CONNECTED],
  ['pasted', SAVED],
] as const) {
  atWidths(`remove confirm ${name}, connect on`, BASE_WIDTHS, async (page, width) => {
    await openAccount(page, width, { initial: block, connect: true });
    await page.getByTestId('discord-remove').click();
    await page.getByTestId('discord-remove-cancel').waitFor();
    await settle(page);
    await shoot(page, `connect-remove-confirm-${name}`, width);
  });
}

atWidths('replace view of a pasted webhook, connect on', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: SAVED, connect: true });
  await page.getByTestId('discord-replace').click();
  await page.getByTestId('discord-cancel').waitFor();
  await settle(page);
  await shoot(page, 'connect-replace-view', width);
});

atWidths('reconnect view', BASE_WIDTHS, async (page, width) => {
  await openAccount(page, width, { initial: CONNECTED, connect: true });
  await shoot(page, 'connect-reconnect', width);
});

// 31. The built site carries neither Discord's logo nor its blurple. Not a PNG: the account page and
// every script and stylesheet in the build's `_astro` folder (the hydrated Account island among them)
// are read from disk and searched; a build without an Account chunk fails the test.
const ASTRO_DIR = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../dist/_astro');
const LOGO_NAME = /discord[-_ ]?(logo|mark)|clyde/i;
const BLURPLE = /(?:#|0x)5865f2\b/i;
const BUNDLE_FILE = /\.(?:js|css)$/;
const ACCOUNT_CHUNK = /^Account\./;

test('the built account page and its bundle carry no Discord logo and no blurple', async ({ request }) => {
  const files = readdirSync(ASTRO_DIR).filter((file) => BUNDLE_FILE.test(file));
  expect(files.some((file) => ACCOUNT_CHUNK.test(file))).toBe(true);
  const page = await (await request.get('/account')).text();
  const sources = [
    ['/account', page],
    ...files.map((file): [string, string] => [file, readFileSync(path.join(ASTRO_DIR, file), 'utf8')]),
  ] as [string, string][];
  const offenders = sources.filter(
    ([name, body]) => LOGO_NAME.test(name) || LOGO_NAME.test(body) || BLURPLE.test(body),
  );
  expect(offenders.map(([name]) => name)).toEqual([]);
});
