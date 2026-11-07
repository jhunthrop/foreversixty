// web/tests/e2e/support/guild-parity/report.ts
// Lane C of the guild parity gate: the raid night report (state `report`, mock web-guild-report.html),
// Tuesday's Onyxia's Lair as the officer Dunmaro reads it. The raid view is
// tests/e2e/fixtures/guild-parity/report-raid.json (design/mocks/guild_report_data.py, the same view
// ws_guild2.raid_report draws); the viewer, guild and home reads are the planner's own raid-*.json
// (guild_raid.py: the same officer on the same guild).
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, type Page } from '@playwright/test';
import type { GuildRaidView } from '../../../../src/lib/guild/api';
import type { GuildPage } from '../../../../src/lib/rankings/api';
import { serveRaidShell } from '../guild-raids';
import type { GuildStateOpener } from '../guild-parity';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../fixtures/guild-parity');

function fixture<T>(name: string): T {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8')) as T;
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'guild-parity' } });

export const openReport: GuildStateOpener = async (page: Page) => {
  const raid = fixture<GuildRaidView>('report-raid.json');
  const guildPage = fixture<GuildPage>('raid-guild.json');
  const { id, region, ruleset } = guildPage.guild;
  const guildPath = `/guild/${region}/${ruleset}/olympus-xxvii`;
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route(`**${guildPath}/raids/*`, serveRaidShell);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(fixture('raid-me.json'))));
  await page.route(`**/v1/guilds/${region}/${ruleset}/olympus-xxvii`, (route) =>
    route.fulfill(envelope(guildPage)),
  );
  await page.route(`**/v1/guilds/${id}/home`, (route) => route.fulfill(envelope(fixture('raid-home.json'))));
  await page.route(`**/v1/guilds/${id}/raids/${raid.id}`, (route) => route.fulfill(envelope(raid)));
  await page.goto(`${guildPath}/raids/${raid.id}`);
  await expect(page.getByTestId('raid-report')).toBeVisible();
  await expect(page.getByTestId('raid-pull-row')).toHaveCount(raid.report?.pulls.length ?? 0);
};
