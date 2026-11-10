// web/tests/e2e/support/guild-parity/overview.ts
// The guild's routing as a member (Hollowmend, Holy Priest) or an officer (Dunmaro, Beast Mastery
// Hunter), which the Progression tab state opens on (guild-parity/tabs.ts); the Overview, Roster,
// Readiness and Loot states are the dense ones (guild-parity/dense.ts).
// Every read the page makes is answered from tests/e2e/fixtures/guild-parity (written by
// design/mocks/guild_overview_fixtures.py, the data ws_guild2.py draws the mocks from); anything
// else falls through to the test fixture's guard, which fails the test.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page, Route } from '@playwright/test';
import { storePointer } from '../selector';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../fixtures/guild-parity');
const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';
const GUILD_ID = 2;

export type Viewer = 'member' | 'officer';

const VIEWERS: Readonly<Record<Viewer, { name: string; spec: string; cls: string }>> = {
  member: { name: 'Hollowmend', spec: 'Holy', cls: 'Priest' },
  officer: { name: 'Dunmaro', spec: 'Beast Mastery', cls: 'Hunter' },
};

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'guild-parity' } });

/** The guild API's answers by path (after `/v1/guilds/<id>`), for one viewer; `home` swaps the
 *  home's file (the Roster state's has a waiting list). */
function guildAnswers(viewer: Viewer, home: string): ReadonlyArray<[RegExp, (url: URL) => string]> {
  return [
    [/^\/home$/, () => home],
    [/^\/raids$/, () => 'overview-raids.json'],
    [/^\/raids\/\d+$/, () => `overview-raid-${viewer}.json`],
    [/^\/progression$/, () => 'overview-progression.json'],
    [
      /^\/readiness$/,
      (url) => (url.searchParams.has('raid') ? 'overview-readiness-raid.json' : 'overview-readiness.json'),
    ],
    [/^\/me\/checklist$/, () => 'overview-checklist.json'],
    [/^\/officer\/actions$/, () => 'overview-actions.json'],
    [/^\/attendance\/recent$/, () => 'overview-attendance.json'],
    [/^\/recruit$/, () => 'overview-recruit.json'],
    [/^\/loot$/, () => 'loot-table.json'],
    [/^\/loot\/nights$/, () => 'loot-nights.json'],
    [/^\/loot\/drops\/\d+\/candidates$/, () => 'loot-drop.json'],
  ];
}

async function answer(route: Route, viewer: Viewer, home: string): Promise<void> {
  const url = new URL(route.request().url());
  if (url.pathname === '/v1/me') return route.fulfill(envelope(fixture(`overview-me-${viewer}.json`)));
  if (url.pathname === `/v1${GUILD_PATH.replace('/guild', '/guilds')}`)
    return route.fulfill(envelope(fixture('overview-page.json')));
  const prefix = `/v1/guilds/${GUILD_ID}`;
  if (url.pathname === `${prefix}/loot/interest`)
    return route.fulfill(envelope({ night: url.searchParams.get('night') ?? '', flags: [] }));
  if (url.pathname.startsWith(prefix)) {
    const rest = url.pathname.slice(prefix.length);
    const hit = guildAnswers(viewer, home).find(([pattern]) => pattern.test(rest));
    if (hit !== undefined) return route.fulfill(envelope(fixture(hit[1](url))));
  }
  return route.fallback();
}

/** Signs in as the mock's viewer (its header character) and routes every guild read. */
export async function routeGuild(
  page: Page,
  viewer: Viewer,
  home = `overview-home-${viewer}.json`,
): Promise<void> {
  const who = VIEWERS[viewer];
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await storePointer(page, {
    source: 'armory',
    ref: `us/pvp/${who.name.toLowerCase()}`,
    label: `${who.name} · ${who.spec} ${who.cls}`,
    classSlug: who.cls.toLowerCase(),
  });
  await page.route('**/v1/**', (route) => answer(route, viewer, home));
}
