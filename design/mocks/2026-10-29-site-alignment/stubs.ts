// design/mocks/2026-10-29-site-alignment/stubs.ts
// The API stubs both capture specs share: every /v1 read answers 404 unless stubbed, and a signed-in
// session is a Premium founder with two characters, a paired device and Discord connected.
import type { Page } from '@playwright/test';
import { PREMIUM_OWN_FOUNDER } from '../../../web/src/test-support/subscribe-fixtures';
import { discordBlock, notificationsKey } from '../../../web/tests/e2e/support/discord-notifications';

const API = 'https://api.foreversixty.gg';
const DAY_MS = 86_400_000;

const ok = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

const character = (name: string, cls: string, spec: string, race: string, level: number, daysAgo: number) => ({
  key: `us/normal/${name.toLowerCase().replace(/\s+/g, '-')}`,
  region: 'us',
  ruleset: 'normal',
  name,
  class: cls,
  spec,
  race,
  level,
  realm: 'Living Flame',
  faction: 'alliance',
  build: { source: 'addon', captured_at: new Date(Date.now() - daysAgo * DAY_MS).toISOString() },
});

const ME = {
  ...PREMIUM_OWN_FOUNDER,
  user: { ...PREMIUM_OWN_FOUNDER.user, battletag: 'Hunthrop#1894' },
  characters: [
    character('Obnoxious Yell', 'Warrior', 'Fury', 'Human', 60, 7),
    character('Elyra Duskvale', 'Priest', 'Holy', 'Dwarf', 42, 2),
  ],
  ...notificationsKey(discordBlock(), false),
};

/** Every API read not stubbed below answers 404, so no capture reaches production. */
export async function isolate(page: Page): Promise<void> {
  await page.route(`${API}/**`, (route) =>
    route.fulfill({
      status: 404,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":{"code":"not_found","message":"Not found."},"request_id":"r"}',
    }),
  );
}

export async function signIn(page: Page): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(ok(ME)));
  await page.route('**/v1/devices', (route) =>
    route.fulfill(ok([{ id: 'd1', name: 'Gaming PC', platform: 'Windows', created_at: '2026-10-01T00:00:00Z', last_seen_at: new Date(Date.now() - 3 * 3600_000).toISOString() }])),
  );
  await page.route('**/v1/reports**', (route) => route.fulfill(ok({ rows: [], total: 0, page: 1, per_page: 20 })));
  await page.addInitScript(() => {
    if (window.localStorage.getItem('fs.currentCharacter') !== null) return;
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/obnoxious-yell',
        label: 'Obnoxious Yell · Fury Warrior',
        classSlug: 'warrior',
        savedAt: new Date().toISOString(),
      }),
    );
  });
}

export async function signOut(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill({ status: 401, contentType: 'application/json', body: '{"ok":false,"data":null,"error":null,"request_id":"r"}' }),
  );
}

