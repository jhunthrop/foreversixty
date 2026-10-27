// web/tests/captures/discord-bot/capture.spec.ts
// Acceptance screenshots for the Forever Sixty bot's web surfaces (design/specs/2026-10-27-discord-bot.md
// §9 site captures rows 1-11, the amendment's row 12, and the raid form's Post to Discord box for a
// bot-only guild). Not a test of behaviour: it drives guild settings #discord and the account You
// panel into each state with the API stubbed and writes PNGs to design/reviews/2026-10-27-discord-bot/.
// Row 10 prints CLS (cls-<width>.json). Discord's own messages are rendered by capture-messages.spec.ts.
import { writeFileSync } from 'node:fs';
import path from 'node:path';
import { test, type Locator, type Page } from '@playwright/test';
import { mockEntitlements } from '../../../src/test-support/billing-fixtures';
import { addSessionHint, createHold, seedWarm } from '../../e2e/support/account-reserve';
import { envelope as plainEnvelope } from '../../e2e/support/discord-notifications';
import { GUILD_URL, NOW, stubGuildRaids } from '../../e2e/support/guild-raids';
import {
  CHANNELS,
  SETTINGS_GUILD_PATH,
  botSettings,
  guildChannel,
  stubGuildBot,
  type GuildBotStubOptions,
} from '../../e2e/support/guild-discord-bot';
import { gotoHydrated } from '../../e2e/support/hydrated';
import { readLayoutShift, trackLayoutShift } from '../../e2e/support/cls';
import {
  MOCK_DIR,
  atWidths,
  cropOf,
  fresh,
  settle,
  sideBySide,
  sizeTo,
  targetIn,
  two,
} from '../support/capture-kit';

const target = targetIn('2026-10-27-discord-bot');
const ALL = [360, 390, 1440, 2000];
const THREE = [360, 390, 1440];
const PAIR = [390, 1440];
const PHONE_AND_DESKTOP = [360, 1440];
const RETURN = (flag: string): string => `${SETTINGS_GUILD_PATH}?discord_bot=${flag}#discord`;
const WEBHOOK = guildChannel({ bot_available: true });
const ORPHAN = { channel_id: null, channel_name: null } as const;

const file = (row: number | string, state: string, width: number, kind = ''): string =>
  target(`discord-bot-${typeof row === 'number' ? two(row) : row}-${state}${kind}-${width}.png`);

const section = (page: Page): Locator => page.locator('#discord');
const id = (page: Page, testid: string): Locator => section(page).getByTestId(testid);

async function shoot(page: Page, row: number | string, state: string, width: number, subject?: Locator) {
  await page.screenshot({ path: file(row, state, width), fullPage: true });
  if (subject !== undefined) await cropOf(page, subject, file(row, state, width, '-crop'));
}

async function openSettings(
  page: Page,
  width: number,
  options: GuildBotStubOptions,
  url = `${SETTINGS_GUILD_PATH}#discord`,
) {
  await sizeTo(page, width);
  const seen = await stubGuildBot(page, options);
  await gotoHydrated(page, url, 'guild-settings');
  await section(page).locator('h2').waitFor();
  await settle(page);
  return seen;
}

/** One state on a fresh page per entry, each captured as the section crop plus the full page. */
async function eachState(
  page: Page,
  row: number | string,
  width: number,
  states: [name: string, options: GuildBotStubOptions, url?: string][],
): Promise<void> {
  for (const [index, [name, options, url]] of states.entries()) {
    const on = index === 0 ? page : await fresh(page);
    await openSettings(on, width, options, url);
    await shoot(on, row, name, width, section(on));
  }
}

// ---- 1-2: states 5 and 4 -----------------------------------------------------------------------

atWidths('1 state 5, nothing connected', ALL, async (page, width) => {
  await eachState(page, 1, width, [['state-5', { discord: botSettings(null) }]]);
});

atWidths('2 state 4, webhook with the bot offer', THREE, async (page, width) => {
  await eachState(page, 2, width, [['state-4-offer', { discord: botSettings(null, WEBHOOK) }]]);
});

// ---- 3: the picker ------------------------------------------------------------------------------

atWidths('3 picker', THREE, async (page, width) => {
  const picker = botSettings({ ...ORPHAN });
  await eachState(page, 3, width, [
    [
      'install-return',
      { discord: botSettings(null), after: picker, channels: CHANNELS },
      RETURN('installed'),
    ],
    ['open-list', { discord: picker }],
    ['empty', { discord: picker, channels: [] }],
    ['error', { discord: picker, channels: 'error' }],
  ]);
  const loading = await fresh(page);
  await sizeTo(loading, width);
  await stubGuildBot(loading, { discord: picker });
  let release: () => void = () => undefined;
  const held = new Promise<void>((resolve) => (release = resolve));
  await loading.route('**/discord/bot/channels', async (route) => {
    await held;
    await route.fallback();
  });
  await gotoHydrated(loading, `${SETTINGS_GUILD_PATH}#discord`, 'guild-settings');
  await id(loading, 'bot-channel').waitFor();
  await settle(loading);
  await shoot(loading, 3, 'loading', width, section(loading));
  release();
  const race = await fresh(page);
  await openSettings(race, width, { discord: picker, unusableSaves: 1 });
  await id(race, 'bot-channel-save').click();
  await race.getByRole('alert').first().waitFor();
  await settle(race);
  await shoot(race, 3, 'race', width, section(race));
});

// ---- 4: installed -------------------------------------------------------------------------------

atWidths('4 installed', ALL, async (page, width) => {
  await openSettings(page, width, { discord: botSettings({ ...ORPHAN }) });
  await id(page, 'bot-channel-save').click();
  await id(page, 'bot-destination').waitFor();
  await settle(page);
  await shoot(page, 4, 'just-saved', width, section(page));
  await eachState(await fresh(page), 4, width, [
    ['warning-3-failures', { discord: botSettings({ failures: 3, last_error: 'Missing Access' }) }],
  ]);
  const remove = await fresh(page);
  await openSettings(remove, width, { discord: botSettings({}, WEBHOOK) });
  await id(remove, 'bot-remove').click();
  await id(remove, 'bot-remove-question').waitFor();
  await settle(remove);
  await shoot(remove, 4, 'remove-question', width, section(remove));
});

// ---- 5: install returns -------------------------------------------------------------------------

atWidths('5 install returns', THREE, async (page, width) => {
  await eachState(page, 5, width, [
    [
      'taken',
      { discord: botSettings(null, { bot_refusal: { guild_name: 'Mograine Raiders' } }) },
      RETURN('taken'),
    ],
    ['taken-expired', { discord: botSettings(null) }, RETURN('taken')],
    ['denied', { discord: botSettings(null) }, RETURN('denied')],
    ['failed', { discord: botSettings(null) }, RETURN('failed')],
  ]);
});

// ---- 6-7: removed in Discord, read-only and frozen ---------------------------------------------

atWidths('6 removed in Discord', PAIR, async (page, width) => {
  await eachState(page, 6, width, [
    ['with-webhook', { discord: botSettings({ removed_in_discord: true }, WEBHOOK) }],
    ['without-webhook', { discord: botSettings({ removed_in_discord: true }) }],
  ]);
});

atWidths('7 read-only and frozen', PHONE_AND_DESKTOP, async (page, width) => {
  await eachState(page, 7, width, [
    ['read-only-state-5', { discord: botSettings(null), officerTools: false }],
    ['read-only-installed', { discord: botSettings({}), officerTools: false }],
    ['frozen-installed', { discord: botSettings({}), frozen: true }],
  ]);
});

// ---- 8: the account You panel ------------------------------------------------------------------

const DISPLAY_NAME = 'zulmara';
const GUILD_NAME = 'Olympus XXVII';

interface LinkState {
  linked?: boolean;
  suggested?: string | null;
  hold?: () => Promise<void>;
}

function meBody(state: LinkState): unknown {
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ notifications: true }),
    discord_link_available: true,
    discord_link: state.linked === true ? { display_name: DISPLAY_NAME } : null,
    discord_link_suggested: state.linked === true ? null : (state.suggested ?? null),
  };
}

async function stubAccount(page: Page, state: LinkState): Promise<void> {
  await page.route('**/v1/me', async (route) => {
    await state.hold?.();
    await route.fulfill(plainEnvelope(meBody(state)));
  });
  await page.route('**/v1/reports**', (route) =>
    route.fulfill(plainEnvelope({ rows: [], total: 0, page: 1, per_page: 20 })),
  );
  await page.route('**/v1/devices', async (route) => {
    await state.hold?.();
    await route.fulfill(plainEnvelope([]));
  });
  await page.route('**/v1/discord/link', (route) => route.fulfill({ status: 204 }));
}

async function openAccount(page: Page, width: number, state: LinkState, url = '/account#you') {
  await sizeTo(page, width);
  await stubAccount(page, state);
  await gotoHydrated(page, url, 'account');
  await settle(page);
}

const you = (page: Page): Locator => page.locator('#you');

atWidths('8 account You panel', ALL, async (page, width) => {
  const states: [string, LinkState, string?][] = [
    ['unlinked', {}],
    ['suggested', { suggested: GUILD_NAME }],
    ['linked', { linked: true }],
    ['return-linked', { linked: true }, '/account?discord_link=linked#you'],
    ['return-denied', {}, '/account?discord_link=denied#you'],
    ['return-failed', {}, '/account?discord_link=failed#you'],
    ['conflict', {}, '/account?discord_link=conflict#you'],
  ];
  for (const [index, [name, state, url]] of states.entries()) {
    const on = index === 0 ? page : await fresh(page);
    await openAccount(on, width, state, url);
    await shoot(on, 8, name, width, you(on));
  }
  const question = await fresh(page);
  await openAccount(question, width, { linked: true });
  await question.getByTestId('discord-unlink').click();
  await question.getByTestId('discord-unlink-question').waitFor();
  await settle(question);
  await shoot(question, 8, 'unlink-question', width, you(question));
  const signedOut = await fresh(page);
  await sizeTo(signedOut, width);
  await signedOut.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":{"code":"unauthorized","message":"Sign in."},"request_id":"r"}',
    }),
  );
  await signedOut.goto('/account');
  await settle(signedOut);
  await shoot(signedOut, 8, 'signed-out', width);
});

// ---- 9: focus-visible and hover -----------------------------------------------------------------

async function focusEach(page: Page, width: number, scope: Locator, controls: string[], state: string) {
  for (const control of controls) {
    await page.keyboard.press('Shift');
    await scope.getByTestId(control).focus();
    await settle(page);
    await cropOf(page, scope, file(9, `${state}-focus-${control}`, width, '-crop'));
  }
}

atWidths('9 focus-visible and hover', PAIR, async (page, width) => {
  await openSettings(page, width, { discord: botSettings(null, WEBHOOK) });
  await focusEach(page, width, section(page), ['bot-add', 'discord-test'], 'offer');
  const installed = await fresh(page);
  await openSettings(installed, width, { discord: botSettings({}) });
  await focusEach(installed, width, section(installed), ['bot-test', 'bot-change', 'bot-remove', 'bot-open'], 'installed');
  const account = await fresh(page);
  await openAccount(account, width, { suggested: GUILD_NAME });
  await focusEach(account, width, you(account), ['discord-link-start'], 'account');
  if (width === 1440) {
    for (const [on, scope, control, state] of [
      [page, section(page), 'bot-add', 'offer'],
      [installed, section(installed), 'bot-test', 'installed'],
      [account, you(account), 'discord-link-start', 'account'],
    ] as const) {
      await on.mouse.move(0, 0);
      await scope.getByTestId(control).hover();
      await settle(on);
      await cropOf(on, scope, file(9, `${state}-hover-${control}`, width, '-crop'));
    }
  }
});

// ---- 10: loading beside ready, CLS --------------------------------------------------------------

atWidths('10 loading beside ready', PAIR, async (page, width) => {
  const cls: Record<string, number> = {};
  const state = botSettings({}, WEBHOOK);
  await sizeTo(page, width);
  await trackLayoutShift(page);
  await stubGuildBot(page, { discord: state });
  let open = (): void => undefined;
  let gate = Promise.resolve();
  await page.route('**/v1/guilds/501/settings', async (route) => {
    await gate;
    await route.fallback();
  });
  await gotoHydrated(page, `${SETTINGS_GUILD_PATH}#discord`, 'guild-settings');
  await section(page).locator('h2').waitFor();
  gate = new Promise<void>((resolve) => (open = resolve));
  await page.reload();
  await settle(page);
  await shoot(page, 10, 'settings-loading', width);
  open();
  await section(page).locator('h2').waitFor();
  await settle(page);
  await shoot(page, 10, 'settings-ready', width);
  cls.settings = Math.round((await readLayoutShift(page)) * 10_000) / 10_000;

  const account = await fresh(page);
  await sizeTo(account, width);
  const hold = createHold();
  await addSessionHint(account, process.env.CAPTURE_BASE_URL ?? `http://localhost:${process.env.CAPTURE_PORT ?? '4481'}`);
  await trackLayoutShift(account);
  await stubAccount(account, { linked: true, hold: hold.wait });
  await seedWarm(account);
  hold.arm();
  await account.goto('/account', { waitUntil: 'commit' });
  await account.getByTestId('account-left-skeleton').waitFor();
  await settle(account);
  await shoot(account, 10, 'account-loading', width);
  await account.evaluate(() => (window.__cls = 0));
  hold.release();
  await account.getByTestId('account').waitFor();
  await account.getByTestId('account-left-skeleton').waitFor({ state: 'detached' });
  await settle(account);
  await shoot(account, 10, 'account-ready', width);
  cls.account = Math.round((await readLayoutShift(account)) * 10_000) / 10_000;
  test.info().annotations.push({ type: 'cls', description: `@${width}: ${JSON.stringify(cls)}` });
  writeFileSync(target(`cls-${width}.json`), `${JSON.stringify({ width, ...cls, cap: 0.02 })}\n`);
});

// ---- 11: side by side with the mock boards ------------------------------------------------------

test('11 side by side @1440', async ({ browser }) => {
  const context = await browser.newContext({ reducedMotion: 'reduce', timezoneId: 'America/Chicago' });
  const boards: [board: string, name: string, open: (page: Page) => Promise<void>][] = [
    ['discord-bot-settings', 'settings', (page) => openSettings(page, 1440, { discord: botSettings({}, WEBHOOK) }).then(() => undefined)],
    ['discord-bot-account', 'account', (page) => openAccount(page, 1440, { suggested: GUILD_NAME })],
  ];
  for (const [board, name, open] of boards) {
    const mockPage = await context.newPage();
    await sizeTo(mockPage, 1440);
    await mockPage.goto(`file://${path.join(MOCK_DIR, `${board}.html`)}`);
    await mockPage.waitForLoadState('networkidle');
    const mockPng = file(11, `mock-${name}`, 1440);
    await mockPage.screenshot({ path: mockPng, fullPage: true });
    const built = await context.newPage();
    await open(built);
    const builtPng = file(11, `built-${name}`, 1440);
    await built.screenshot({ path: builtPng, fullPage: true });
    await sideBySide(await context.newPage(), mockPng, builtPng, file(11, `side-by-side-${name}`, 1440), 1440);
  }
  await context.close();
});

// ---- 12: channel can't be used ------------------------------------------------------------------

const DELETED = { reason: 'deleted', missing: [] } as const;
const PERMISSIONS = {
  reason: 'permissions',
  missing: ['Send Messages', 'Embed Links'],
} as const;

atWidths('12 channel cannot be used', THREE, async (page, width) => {
  const causes: [string, GuildBotCause, boolean][] = [
    ['deleted-webhook', DELETED, true],
    ['permissions-no-webhook', PERMISSIONS, false],
    ['deleted-no-webhook', DELETED, false],
    ['permissions-webhook', PERMISSIONS, true],
  ];
  for (const [index, [name, cause, withWebhook]] of causes.entries()) {
    const on = index === 0 ? page : await fresh(page);
    const discord = botSettings(
      { channel_unusable: { ...cause, missing: [...cause.missing] } },
      withWebhook ? WEBHOOK : {},
    );
    await openSettings(on, width, { discord });
    await shoot(on, 12, `${name}-notice`, width, section(on));
    if (index > 1) continue;
    await id(on, 'bot-choose-channel').click();
    await id(on, 'bot-channel').waitFor();
    await settle(on);
    await shoot(on, 12, `${name}-picker-open`, width, section(on));
  }
});

type GuildBotCause = { reason: 'deleted' | 'permissions'; missing: readonly string[] };

// ---- the raid form for a bot-only guild ----------------------------------------------------------

atWidths('raid form, bot-only Post to Discord', THREE, async (page, width) => {
  await sizeTo(page, width);
  await stubGuildRaids(page, {
    viewer: 'officer',
    discordAvailable: true,
    discordConnected: true,
    postDefault: true,
  });
  await page.clock.install({ time: NOW });
  await gotoHydrated(page, `${GUILD_URL}#raids`, 'guild-raids-tab');
  await page.getByTestId('guild-raids-plan').click();
  await page.getByTestId('raid-form-create').waitFor();
  await settle(page);
  await shoot(page, 'raid-form', 'post-to-discord-bot-only', width, page.getByTestId('raid-form-create'));
});

