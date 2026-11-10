// web/tests/e2e/support/guild-parity/directory.ts
// The guild parity states for the guild landing (/guild), join and claim (GUILD-ALL lane):
// `landing-out` (signed out), `landing-in` (an alt with no guild, the account in two guilds),
// `join` (Olympus XXVII's invite link, signed in with no guild) and `claim` (Dunmaro holds the
// claim). Each opener routes the committed fixtures (design/mocks/guild_directory_fixtures.py)
// and waits until the page shows its data; anything else falls through to the test fixture's
// guard, which fails the test.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page, Route } from '@playwright/test';
import { expect } from '../test';
import { serveShell } from '../guild-recruit';
import { storePointer } from '../selector';
import type { GuildStateOpener } from '../guild-parity';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../fixtures/guild-parity');
/** guild_directory_fixtures.INVITE_TOKEN: the one invite path the fixture build prerenders. */
const INVITE_TOKEN = 'fixtureinvitetoken1';
const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';
const GUILD_ID = 2;

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'guild-parity' } });
const SIGNED_OUT = {
  status: 401,
  contentType: 'application/json',
  body: '{"ok":false,"data":null,"error":null,"request_id":"guild-parity"}',
};

/** The reads one state answers, by path; `me` null is signed out. */
interface Answers {
  me: string | null;
  extra?: Readonly<Record<string, string>>;
}

async function route(page: Page, answers: Answers): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/**', (r: Route) => {
    const url = new URL(r.request().url());
    if (url.pathname === '/v1/me')
      return answers.me === null ? r.fulfill(SIGNED_OUT) : r.fulfill(envelope(fixture(answers.me)));
    if (url.pathname === '/v1/rankings/guilds')
      return r.fulfill(envelope(fixture('directory-rankings.json')));
    const file = answers.extra?.[url.pathname];
    if (file !== undefined) return r.fulfill(envelope(fixture(file)));
    return r.fallback();
  });
}

async function openLanding(page: Page, signedIn: boolean): Promise<void> {
  if (signedIn) {
    await storePointer(page, {
      source: 'armory',
      ref: 'us/pvp/briarwick',
      label: 'Briarwick · Druid',
      classSlug: 'druid',
    });
  }
  await route(page, { me: signedIn ? 'directory-me-in.json' : null });
  await page.goto('/guild');
  await expect(page.getByTestId('home-top-guilds')).toBeVisible();
  if (signedIn) await expect(page.getByTestId('guild-landing-mine')).toBeVisible();
  else await expect(page.getByTestId('guild-landing-sign-in')).toBeVisible();
  await page.waitForLoadState('networkidle');
}

async function openJoin(page: Page): Promise<void> {
  await storePointer(page, {
    source: 'armory',
    ref: 'us/pvp/ashenveil',
    label: 'Ashenveil · Paladin',
    classSlug: 'paladin',
  });
  await route(page, {
    me: 'directory-me-join.json',
    extra: { [`/v1/guild-invites/${INVITE_TOKEN}`]: 'directory-invite.json' },
  });
  await page.goto(`/guild/invite/${INVITE_TOKEN}`);
  await expect(page.getByTestId('guild-join-button')).toBeVisible();
  await page.waitForLoadState('networkidle');
}

async function openClaim(page: Page): Promise<void> {
  await storePointer(page, {
    source: 'armory',
    ref: 'us/pvp/dunmaro',
    label: 'Dunmaro · Beast Mastery Hunter',
    classSlug: 'hunter',
  });
  await route(page, {
    me: 'directory-me-claim.json',
    extra: {
      [`/v1${GUILD_PATH.replace('/guild', '/guilds')}`]: 'overview-page.json',
      [`/v1/guilds/${GUILD_ID}/settings`]: 'directory-claim-settings.json',
      [`/v1/guilds/${GUILD_ID}/home`]: 'overview-home-officer.json',
      [`/v1/guilds/${GUILD_ID}/setup`]: 'directory-claim-setup.json',
    },
  });
  // The fixture build prerenders one claim path (The Last Watch's): this one is the guild shell.
  await page.route(`**${GUILD_PATH}/claim`, serveShell);
  await page.goto(`${GUILD_PATH}/claim`);
  await expect(page.getByTestId('guild-claim-next')).toBeVisible();
  await page.waitForLoadState('networkidle');
}

export const DIRECTORY_OPENERS: Readonly<Record<string, GuildStateOpener>> = {
  'landing-out': (page) => openLanding(page, false),
  'landing-in': (page) => openLanding(page, true),
  join: openJoin,
  claim: openClaim,
};
