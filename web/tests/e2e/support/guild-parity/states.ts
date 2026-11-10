// web/tests/e2e/support/guild-parity/states.ts
// The Overview states the persona round left out (GUILD-ALL lane): `before-raids` (an officer at
// level 23 while nothing is logged, the owner's screenshot), `low-level` (a level 23 member, before
// the raids open), `no-raid` (the raids are open, nothing planned or logged) and `public` (a
// signed-in visitor in no guild). Every read is answered from tests/e2e/fixtures/guild-parity
// (design/mocks/guild_states_fixtures.py, the data ws_guild_states.py draws the mocks from); the
// pre-raid and open states re-pin the clock to the mock's own date.
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
/** guild_states_fixtures.PRE_NOW and OPEN_NOW. */
const PRE_NOW = new Date('2026-11-07T22:00:00Z');
const OPEN_NOW = new Date('2026-12-15T22:00:00Z');

interface StateSpec {
  /** The header character: name, spec and class words, class slug. */
  who: { name: string; spec: string; cls: string };
  /** The `states-me-*`, `states-home-*` and feed and checklist file suffixes; null: no membership. */
  me: string;
  home: string | null;
  feed: string;
  checklist: string;
  now: Date | null;
}

const STATES: Readonly<Record<string, StateSpec>> = {
  'before-raids': {
    who: { name: 'Dunmaro', spec: 'Beast Mastery', cls: 'Hunter' },
    me: 'officer-pre',
    home: 'officer-pre',
    feed: 'pre',
    checklist: 'pre',
    now: PRE_NOW,
  },
  'low-level': {
    who: { name: 'Nightlatch', spec: 'Subtlety', cls: 'Rogue' },
    me: 'member-pre',
    home: 'member-pre',
    feed: 'pre',
    checklist: 'pre',
    now: PRE_NOW,
  },
  'no-raid': {
    who: { name: 'Hollowmend', spec: 'Holy', cls: 'Priest' },
    me: 'member-open',
    home: 'member-open',
    feed: 'open',
    checklist: 'open',
    now: OPEN_NOW,
  },
  public: {
    who: { name: 'Felsnap', spec: 'Fury', cls: 'Warrior' },
    me: 'visitor',
    home: null,
    feed: 'pre',
    checklist: 'pre',
    now: null,
  },
};

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'guild-parity' } });

/** The guild API's answers by path (after `/v1/guilds/<id>`). The visitor's raids and progression
 *  are the member Overview's (one logged night); the members' are an empty tier. */
function guildAnswers(s: StateSpec): ReadonlyArray<[RegExp, string]> {
  const visitor = s.home === null;
  return [
    [/^\/home$/, `states-home-${s.home ?? ''}.json`],
    [/^\/raids$/, visitor ? 'overview-raids.json' : 'states-raids.json'],
    [/^\/progression$/, visitor ? 'overview-progression.json' : 'states-progression.json'],
    [/^\/readiness$/, 'states-readiness.json'],
    [/^\/me\/checklist$/, `states-checklist-${s.checklist}.json`],
    [/^\/officer\/actions$/, 'states-actions.json'],
    [/^\/attendance\/recent$/, 'states-attendance.json'],
    [/^\/recruit$/, 'overview-recruit.json'],
    [/^\/loot\/nights$/, 'states-loot-nights.json'],
    [/^\/applications$/, 'states-applications.json'],
  ];
}

async function answer(route: Route, s: StateSpec): Promise<void> {
  const url = new URL(route.request().url());
  if (url.pathname === '/v1/me') return route.fulfill(envelope(fixture(`states-me-${s.me}.json`)));
  if (url.pathname === '/v1/feed') return route.fulfill(envelope(fixture(`states-feed-${s.feed}.json`)));
  if (url.pathname === `/v1${GUILD_PATH.replace('/guild', '/guilds')}`)
    return route.fulfill(envelope(fixture('overview-page.json')));
  const prefix = `/v1/guilds/${GUILD_ID}`;
  if (/^\/members\/[^/]+\/profile$/.test(url.pathname.slice(prefix.length)))
    return route.fulfill({
      status: 404,
      json: {
        ok: false,
        data: null,
        error: { code: 'not_found', message: 'no profile' },
        request_id: 'guild-parity',
      },
    });
  if (url.pathname.startsWith(prefix)) {
    const rest = url.pathname.slice(prefix.length);
    const hit = guildAnswers(s).find(([pattern]) => pattern.test(rest));
    if (hit !== undefined) return route.fulfill(envelope(fixture(hit[1])));
  }
  return route.fallback();
}

async function openState(page: Page, id: string): Promise<void> {
  const s = STATES[id];
  if (s.now !== null) await page.clock.setFixedTime(s.now);
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await storePointer(page, {
    source: 'armory',
    ref: `us/pvp/${s.who.name.toLowerCase()}`,
    label: `${s.who.name} · ${s.who.spec} ${s.who.cls}`,
    classSlug: s.who.cls.toLowerCase(),
  });
  await page.route('**/v1/**', (route) => answer(route, s));
  await page.goto(GUILD_PATH);
  // An officer's rail is the cover and attendance alone (batch 3); a member and a visitor read the last raid.
  if (!s.me.startsWith('officer')) await expect(page.getByTestId('guild-overview-last-raid')).toBeVisible();
  if (s.home === null) {
    await expect(page.getByTestId('guild-overview-recruit')).toBeVisible();
    await expect(page.getByTestId('guild-roster-gate')).toBeVisible();
  } else {
    await expect(page.getByTestId('guild-week-one-sentence')).toBeVisible();
    await expect(page.getByTestId('guild-overview-activity-row').first()).toBeVisible();
  }
  await page.waitForLoadState('networkidle');
}

export const STATES_OPENERS: Readonly<Record<string, GuildStateOpener>> = Object.fromEntries(
  Object.keys(STATES).map((id) => [id, (page: Page) => openState(page, id)]),
);
