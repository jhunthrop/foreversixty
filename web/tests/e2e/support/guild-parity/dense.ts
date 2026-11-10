// web/tests/e2e/support/guild-parity/dense.ts
// The dense guild states (design/specs/2026-11-09-guild-dense.md, mocks design/mocks/ws_guild_dense.py):
// the 40-raider Olympus XXVII five raid nights in, as its officer (Dunmaro) and a member
// (Hollowmend) read it, on the Overview, Roster, Readiness and Loot tabs. Every read answers from
// tests/e2e/fixtures/guild-parity/dense-*.json (design/mocks/guild_dense_fixtures.py), at the
// world's own clock, Sat 9 Jan 2027 4:00 pm CST; anything else falls through to the test guard.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page, Route } from '@playwright/test';
import { expect } from '../test';
import { storePointer } from '../selector';
import type { GuildStateOpener } from '../guild-parity';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../fixtures/guild-parity');
const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';
const GUILD_ID = 2;
/** guild_dense_data.NOW. */
export const DENSE_NOW = new Date('2027-01-09T22:00:00Z');
/** guild_dense_data.SELECTED_DROP. */
const SELECTED_DROP = 17075;

type Viewer = 'member' | 'officer';

const VIEWERS: Readonly<Record<Viewer, { name: string; spec: string; cls: string }>> = {
  member: { name: 'Hollowmend', spec: 'Holy', cls: 'Priest' },
  officer: { name: 'Dunmaro', spec: 'Beast Mastery', cls: 'Hunter' },
};

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'guild-dense' } });

function answers(viewer: Viewer): ReadonlyArray<[RegExp, (url: URL) => string]> {
  return [
    [/^\/home$/, () => `dense-home-${viewer}.json`],
    [/^\/raids$/, () => 'dense-raids.json'],
    [/^\/raids\/45$/, () => `dense-raid-${viewer}.json`],
    [/^\/raids\/44$/, () => 'dense-raid-44.json'],
    [/^\/raids\/43$/, () => 'dense-raid-43.json'],
    [/^\/progression$/, () => 'dense-progression.json'],
    [
      /^\/readiness$/,
      (url) => (url.searchParams.has('raid') ? 'dense-readiness-raid.json' : 'dense-readiness.json'),
    ],
    [/^\/me\/checklist$/, () => 'dense-checklist.json'],
    [/^\/officer\/actions$/, () => 'dense-actions.json'],
    [/^\/attendance\/recent$/, () => 'dense-attendance.json'],
    [/^\/recruit$/, () => 'overview-recruit.json'],
    [/^\/applications$/, () => 'recruit-applications.json'],
    [/^\/loot$/, () => 'dense-loot-table.json'],
    [/^\/loot\/nights$/, () => 'dense-loot-nights.json'],
    [/^\/members\/[^/]+\/profile$/, () => 'dense-profile-member.json'],
    [/^\/loot\/drops\/\d+\/candidates$/, (url) => `dense-loot-drop-${url.pathname.split('/').at(-2)}.json`],
  ];
}

async function answer(route: Route, viewer: Viewer): Promise<void> {
  const url = new URL(route.request().url());
  if (url.pathname === '/v1/me') return route.fulfill(envelope(fixture(`overview-me-${viewer}.json`)));
  if (url.pathname === `/v1${GUILD_PATH.replace('/guild', '/guilds')}`)
    return route.fulfill(envelope(fixture('dense-page.json')));
  const prefix = `/v1/guilds/${GUILD_ID}`;
  if (url.pathname === `${prefix}/loot/interest`)
    return route.fulfill(envelope({ night: url.searchParams.get('night') ?? '', flags: [] }));
  if (url.pathname.startsWith(prefix)) {
    const rest = url.pathname.slice(prefix.length);
    const hit = answers(viewer).find(([pattern]) => pattern.test(rest));
    if (hit !== undefined) return route.fulfill(envelope(fixture(hit[1](url))));
  }
  return route.fallback();
}

async function routeDense(page: Page, viewer: Viewer): Promise<void> {
  const who = VIEWERS[viewer];
  await page.clock.setFixedTime(DENSE_NOW);
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await storePointer(page, {
    source: 'armory',
    ref: `us/pvp/${who.name.toLowerCase()}`,
    label: `${who.name} · ${who.spec} ${who.cls}`,
    classSlug: who.cls.toLowerCase(),
  });
  await page.route('**/v1/**', (route) => answer(route, viewer));
}

async function openTab(page: Page, viewer: Viewer, hash: string, ready: string): Promise<void> {
  await routeDense(page, viewer);
  await page.goto(`${GUILD_PATH}${hash}`);
  await expect(page.getByTestId(ready)).toBeVisible();
  await page.waitForLoadState('networkidle');
}

async function openLoot(page: Page): Promise<void> {
  await openTab(page, 'officer', '#loot', 'guild-loot-fairness');
  await page.getByTestId(`guild-loot-drop-${SELECTED_DROP}`).click();
  await expect(page.locator('[data-testid^="guild-loot-candidate-"]').first()).toBeVisible();
  await page.waitForLoadState('networkidle');
}

export const DENSE_OPENERS: Readonly<Record<string, GuildStateOpener>> = {
  'dense-officer': (page) => openTab(page, 'officer', '', 'guild-week-wednesday'),
  'dense-member': (page) => openTab(page, 'member', '', 'guild-your-week'),
  'dense-roster': (page) => openTab(page, 'officer', '#roster', 'guild-roster-board'),
  'dense-roster-member': (page) => openTab(page, 'member', '#roster', 'guild-roster-board'),
  'dense-readiness': (page) => openTab(page, 'officer', '#readiness', 'guild-readiness-board'),
  'dense-loot': openLoot,
};
