// web/tests/e2e/support/guild-parity/settings.ts
// The guild parity states for the Settings tab and the guild settings page (GUILD-ALL lane):
// `settings`, the officer's Settings tab on the guild page (a founder Raid guild plan billed to
// Dunmaro, one seat given, a connected Discord channel); `settings-bot`, the same tab with the bot
// installed in the guild's Discord server instead of the webhook; `settings-page`, the standalone
// /settings route for the same guild before a plan (the plan chooser). Every read is a committed
// fixture (design/mocks/guild_settings_fixtures.py and the Overview's own files); anything else
// falls through to the test fixture's guard, which fails the test.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page, Route } from '@playwright/test';
import { expect } from '../test';
import { serveRaidShell } from '../guild-raids';
import { storePointer } from '../selector';
import type { GuildStateOpener } from '../guild-parity';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../fixtures/guild-parity');
const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';
const GUILD_ID = 2;

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'guild-parity' } });

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

/** The guild API's answers by path (after `/v1/guilds/<id>`); `billing` names the plan's read,
 *  `settings` the settings read (the webhook channel, or the bot installed). */
function guildAnswers(billing: string, settings: string): ReadonlyArray<[RegExp, string]> {
  return [
    [/^\/home$/, 'overview-home-officer.json'],
    [/^\/raids\/\d+$/, 'overview-raid-officer.json'],
    [/^\/settings$/, settings],
    [/^\/billing$/, billing],
  ];
}

async function answer(route: Route, billing: string, settings: string): Promise<void> {
  const url = new URL(route.request().url());
  if (url.pathname === '/v1/me') return route.fulfill(envelope(fixture('overview-me-officer.json')));
  if (url.pathname === `/v1${GUILD_PATH.replace('/guild', '/guilds')}`)
    return route.fulfill(envelope(fixture('overview-page.json')));
  const prefix = `/v1/guilds/${GUILD_ID}`;
  if (url.pathname.startsWith(prefix)) {
    const rest = url.pathname.slice(prefix.length);
    const hit = guildAnswers(billing, settings).find(([pattern]) => pattern.test(rest));
    if (hit !== undefined) return route.fulfill(envelope(fixture(hit[1])));
  }
  return route.fallback();
}

/** Signs in as the Overview's officer (Dunmaro, its header character) and routes every read. */
async function routeSettings(
  page: Page,
  billing: string,
  settings = 'settings-settings.json',
): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await storePointer(page, {
    source: 'armory',
    ref: 'us/pvp/dunmaro',
    label: 'Dunmaro · Beast Mastery Hunter',
    classSlug: 'hunter',
  });
  await page.route('**/v1/**', (route) => answer(route, billing, settings));
}

async function settled(page: Page): Promise<void> {
  await expect(page.getByTestId('guild-billing-ready')).toBeVisible();
  await expect(page.getByTestId('guild-settings-records')).toBeVisible();
  await page.waitForLoadState('networkidle');
}

async function openSettingsTab(page: Page, settings = 'settings-settings.json'): Promise<void> {
  await routeSettings(page, 'settings-billing.json', settings);
  await page.goto(`${GUILD_PATH}#settings`);
  await expect(page.getByTestId('guild-settings-discord')).toBeVisible();
  await expect(page.getByTestId('guild-crest-settings')).toBeVisible();
  await settled(page);
}

async function openSettingsPage(page: Page): Promise<void> {
  await routeSettings(page, 'settings-billing-none.json');
  // The fixture build prerenders one guild's /settings; this guild's route is served the same shell.
  await page.route(`**${GUILD_PATH}/settings`, serveRaidShell);
  await page.goto(`${GUILD_PATH}/settings`);
  await expect(page.getByTestId('guild-billing-chooser')).toBeVisible();
  await settled(page);
}

export const SETTINGS_OPENERS: Readonly<Record<string, GuildStateOpener>> = {
  settings: (page) => openSettingsTab(page),
  // The same tab with the Forever Sixty bot installed and no webhook (bot spec 3.1, installed).
  'settings-bot': async (page) => {
    await openSettingsTab(page, 'settings-settings-bot.json');
    await expect(page.getByTestId('bot-destination')).toBeVisible();
  },
  'settings-page': openSettingsPage,
};
