// web/tests/e2e/support/guild-parity/raid.ts
// Lane B of the guild parity gate: the raid planner for Thursday's Onyxia's Lair as its officer
// (Dunmaro) reads it, mock web-guild-raid.html. Every read is a committed fixture
// (tests/e2e/fixtures/guild-parity/raid-*.json, written by design/mocks/guild_fixtures.py from
// guild_raid.py, the same view ws_guild2.raid_planner draws).
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, type Page } from '@playwright/test';
import { serveRaidShell } from '../guild-raids';
import { GUILD_PLANNER_READS, apiUnreachable } from '../test';
import type { GuildStateOpener } from '../guild-parity';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../fixtures/guild-parity');

interface RaidFixture {
  id: number;
}
interface GuildFixture {
  guild: { id: number; region: string; ruleset: string };
}

function fixture<T>(name: string): T {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8')) as T;
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'guild-parity' } });

export async function openRaidPlanner(page: Page): Promise<void> {
  const raid = fixture<RaidFixture>('raid-onyxia.json');
  const page_ = fixture<GuildFixture>('raid-guild.json');
  const { id, region, ruleset } = page_.guild;
  const guildPath = `/guild/${region}/${ruleset}/olympus-xxvii`;
  await apiUnreachable(page, ...GUILD_PLANNER_READS);
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route(`**${guildPath}/raids/*`, serveRaidShell);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(fixture('raid-me.json'))));
  await page.route(`**/v1/guilds/${region}/${ruleset}/olympus-xxvii`, (route) =>
    route.fulfill(envelope(page_)),
  );
  await page.route(`**/v1/guilds/${id}/home`, (route) => route.fulfill(envelope(fixture('raid-home.json'))));
  await page.route(`**/v1/guilds/${id}/raids/${raid.id}`, (route) => route.fulfill(envelope(raid)));
  await page.goto(`${guildPath}/raids/${raid.id}`);
  await expect(page.getByTestId('raid-signups')).toBeVisible();
  await expect(page.getByTestId('raid-officer-row')).toBeVisible();
}

export const RAID_OPENERS: Readonly<Record<string, GuildStateOpener>> = { raid: openRaidPlanner };
