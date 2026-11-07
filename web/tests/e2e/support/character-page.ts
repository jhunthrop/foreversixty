// web/tests/e2e/support/character-page.ts
// The character page's states from committed fixtures (tests/e2e/fixtures/character/ and
// fixtures/home, written by design/mocks/character_fixtures.py): the character answer, its
// rating, the owner's /v1/me, the name search, the trimmed BiS and item files and the icons. The
// parity gate opens every state through it (design/specs/2026-11-06-character.md section 3).
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, type Page } from '@playwright/test';
import { ACTIVE_BUILD } from './active-build';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const FIXTURES = path.resolve(HERE, '../fixtures/character');
const HOME = path.resolve(HERE, '../fixtures/home');
const ICONS = path.resolve(HERE, '../../../../data/builds', ACTIVE_BUILD, 'icons');
/** The fixtures' own "today", the home fixtures' too: the sync line's age reads against it. */
export const CHARACTER_NOW = new Date('2026-11-06T18:00:00Z');

export interface CharacterState {
  id: string;
  path: string;
  /** The character answer's fixture, or a status (404 unknown, 500 failed), or `pending`. */
  character: string | number;
  rating?: string;
  /** The owner's /v1/me; absent, the viewer is signed out. */
  me?: string;
  search?: string;
  /** The character's activity (GET /v1/feed?character=); absent, no rows. */
  feed?: string;
  /** The owner's /v1/home answer and their guild's home (the standing line). */
  home?: string;
  guild?: string;
}

function fixture(name: string): unknown {
  const own = path.join(FIXTURES, name);
  return JSON.parse(readFileSync(existsSync(own) ? own : path.join(HOME, name), 'utf8'));
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'character' } });
const failure = (status: number) => ({
  status,
  json: { ok: false, data: null, error: { message: 'failed' }, request_id: 'character' },
});

async function routeData(page: Page): Promise<void> {
  for (const [file, route] of [
    ['bis-priest-holy.json', 'bis/priest-holy.json'],
    ['bis-hunter-beast-mastery.json', 'bis/hunter-beast-mastery.json'],
    ['items-priest.json', 'items/priest.json'],
    ['items-hunter.json', 'items/hunter.json'],
  ]) {
    await page.route(`**/data/*/${route}`, (r) => r.fulfill({ json: fixture(file) }));
  }
  await page.route('**/data/*/icons/*.webp', (route) => {
    const file = path.join(ICONS, new URL(route.request().url()).pathname.split('/').pop() ?? '');
    return existsSync(file)
      ? route.fulfill({ status: 200, contentType: 'image/webp', body: readFileSync(file) })
      : route.fulfill({ status: 404, body: '' });
  });
}

/** Routes everything the state reads. */
export async function routeCharacter(page: Page, state: CharacterState): Promise<void> {
  await page.clock.setFixedTime(CHARACTER_NOW);
  const api = `**/v1/characters${state.path.replace('/character', '')}`;
  if (state.me !== undefined) {
    await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
    const me = fixture(state.me);
    await page.route('**/v1/me', (route) => route.fulfill(envelope(me)));
  }
  const { character } = state;
  await page.route(api, (route) =>
    character === 'pending'
      ? new Promise<void>(() => {})
      : typeof character === 'number'
        ? route.fulfill(failure(character))
        : route.fulfill(envelope(fixture(character))),
  );
  if (state.rating !== undefined) {
    const rating = fixture(state.rating);
    await page.route(`${api}/rating`, (route) => route.fulfill(envelope(rating)));
  }
  const feed = state.feed === undefined ? { rows: [], next_cursor: null } : fixture(state.feed);
  await page.route('**/v1/feed?*', (route) => route.fulfill(envelope(feed)));
  if (state.home !== undefined) {
    const home = fixture(state.home);
    await page.route('**/v1/home?*', (route) => route.fulfill(envelope(home)));
  }
  if (state.guild !== undefined) {
    const guild = fixture(state.guild);
    await page.route('**/v1/guilds/*/home', (route) => route.fulfill(envelope(guild)));
  }
  const rows = state.search === undefined ? [] : (fixture(state.search) as { rows: unknown[] }).rows;
  await page.route('**/v1/characters/search*', (route) => route.fulfill(envelope({ rows })));
  await routeData(page);
}

/** Opens the state and waits for it to settle: the island hydrated and its reads answered. */
export async function openCharacter(page: Page, state: CharacterState): Promise<void> {
  await routeCharacter(page, state);
  await page.goto(state.path);
  await expect(page.locator('astro-island:not([ssr]) [data-region="head"]')).toHaveCount(1);
  if (state.character === 'pending') return;
  await expect(page.locator('[aria-busy="true"]')).toHaveCount(0);
  await page.waitForLoadState('networkidle');
}
