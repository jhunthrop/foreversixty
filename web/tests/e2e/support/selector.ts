// web/tests/e2e/support/selector.ts
// Fixtures for the nav character selector specs: a `/v1/me` answer with the characters a test
// names, a stored pointer, and a signed-out answer. Names and classes are examples, the same
// way the design mock's are.
import type { Page } from '@playwright/test';

export interface FixtureCharacter {
  name: string;
  class: string;
  spec: string;
  level?: number;
  realm?: string;
  faction?: 'alliance' | 'horde';
  ruleset?: string;
  /** Days since the newest build was captured; omit for a character with no build. */
  buildDaysAgo?: number;
  buildSource?: 'addon' | 'blizzard';
  syncError?: string | null;
}

export const OBNOXIOUS: FixtureCharacter = {
  name: 'Obnoxious Yell',
  class: 'Warrior',
  spec: 'Fury',
  level: 60,
  realm: 'Living Flame',
  faction: 'alliance',
  buildDaysAgo: 0.01,
};

const DAY_MS = 86_400_000;

export function keyOf(character: FixtureCharacter): string {
  return `us/normal/${character.name.toLowerCase().replace(/\s+/g, '-')}`;
}

function meCharacter(character: FixtureCharacter): Record<string, unknown> {
  return {
    key: keyOf(character),
    region: 'us',
    ruleset: character.ruleset ?? 'normal',
    name: character.name,
    class: character.class,
    spec: character.spec,
    level: character.level ?? 60,
    realm: character.realm ?? 'Living Flame',
    faction: character.faction ?? 'alliance',
    ...(character.buildDaysAgo === undefined
      ? {}
      : {
          build: {
            source: character.buildSource ?? 'addon',
            captured_at: new Date(Date.now() - character.buildDaysAgo * DAY_MS).toISOString(),
            ...(character.syncError === undefined ? {} : { sync_error: character.syncError }),
          },
        }),
  };
}

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

/** A signed-in session whose `/v1/me` lists `characters`. */
export async function signInWith(page: Page, characters: FixtureCharacter[]): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
        characters: characters.map(meCharacter),
        guilds: [],
      }),
    ),
  );
}

/** `/v1/me` answers 401: nobody is signed in (or the session ended). */
export async function signedOut(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
    }),
  );
}

export interface StoredPointer {
  source: 'armory' | 'code' | 'addon';
  ref: string;
  label: string;
  classSlug: string;
}

/** Writes `fs.currentCharacter` before any page script runs. */
export async function storePointer(page: Page, pointer: StoredPointer): Promise<void> {
  // Only when none is stored yet: the script runs on every navigation, and a test that
  // chooses another character must keep its choice across the next page load.
  await page.addInitScript((value) => {
    if (window.localStorage.getItem('fs.currentCharacter') !== null) return;
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({ ...value, savedAt: new Date().toISOString() }),
    );
  }, pointer);
}

export function armoryPointer(character: FixtureCharacter): StoredPointer {
  return {
    source: 'armory',
    ref: keyOf(character),
    label: `${character.name} · ${character.spec} ${character.class}`,
    classSlug: character.class.toLowerCase(),
  };
}
