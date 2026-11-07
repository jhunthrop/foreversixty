// web/tests/e2e/support/guild-parity/raids.ts
// The guild parity states for the Raids tab and the raid page (GUILD-ALL lane): `raids`, the
// officer's Raids tab (Dunmaro: two planned raids and the logged night); `raids-plan`, the same tab
// with Plan an event's form open; `assignments`, the raid planner with the Off tank slot's picker
// open. The guild page reads the Overview's answers (design/mocks/guild_overview_fixtures.py) and
// the Raids tab its own (design/mocks/guild_raids_fixtures.py); anything else falls through to the
// test fixture's guard, which fails the test.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page, Route } from '@playwright/test';
import { expect } from '../test';
import { storePointer } from '../selector';
import type { GuildStateOpener } from '../guild-parity';
import { openRaidPlanner } from './raid';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../fixtures/guild-parity');
const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';
const GUILD_ID = 2;
/** The slot whose picker the assignments state opens (guild_raid.onyxia_template). */
const EDITED_SLOT = 'off_tank';

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'guild-parity' } });

/** The guild API's answers by path (after `/v1/guilds/<id>`), as the officer reads them. */
const GUILD_ANSWERS: ReadonlyArray<[RegExp, string]> = [
  [/^\/home$/, 'overview-home-officer.json'],
  [/^\/raids$/, 'raids-page.json'],
  [/^\/raids\/\d+$/, 'overview-raid-officer.json'],
  [/^\/progression$/, 'overview-progression.json'],
];

async function answer(route: Route): Promise<void> {
  const url = new URL(route.request().url());
  if (url.pathname === '/v1/me') return route.fulfill(envelope(fixture('overview-me-officer.json')));
  if (url.pathname === `/v1${GUILD_PATH.replace('/guild', '/guilds')}`)
    return route.fulfill(envelope(fixture('overview-page.json')));
  const prefix = `/v1/guilds/${GUILD_ID}`;
  if (url.pathname.startsWith(prefix)) {
    const rest = url.pathname.slice(prefix.length);
    const hit = GUILD_ANSWERS.find(([pattern]) => pattern.test(rest));
    if (hit !== undefined) return route.fulfill(envelope(fixture(hit[1])));
  }
  return route.fallback();
}

async function openRaidsTab(page: Page, withForm: boolean): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await storePointer(page, {
    source: 'armory',
    ref: 'us/pvp/dunmaro',
    label: 'Dunmaro · Beast Mastery Hunter',
    classSlug: 'hunter',
  });
  await page.route('**/v1/**', answer);
  await page.goto(`${GUILD_PATH}#raids`);
  await expect(page.getByTestId('guild-planned-rows')).toBeVisible();
  await expect(page.getByTestId('guild-raids-rows')).toBeVisible();
  if (withForm) {
    await page.getByTestId('guild-raids-plan').click();
    await expect(page.getByTestId('raid-form-create')).toBeVisible();
  }
  await page.waitForLoadState('networkidle');
}

async function openAssignments(page: Page): Promise<void> {
  await openRaidPlanner(page);
  await page.getByTestId(`assign-edit-${EDITED_SLOT}`).click();
  await expect(page.getByTestId(`assign-picker-${EDITED_SLOT}`)).toBeVisible();
}

export const RAIDS_OPENERS: Readonly<Record<string, GuildStateOpener>> = {
  raids: (page) => openRaidsTab(page, false),
  'raids-plan': (page) => openRaidsTab(page, true),
  assignments: openAssignments,
};
