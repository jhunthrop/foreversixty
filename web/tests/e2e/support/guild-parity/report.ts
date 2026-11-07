// web/tests/e2e/support/guild-parity/report.ts
// The raid night report's parity state (`report`, mock web-guild-report.html): Onyxia, Tue 8 Dec,
// read by the officer Dunmaro. Routes the committed fixtures (tests/e2e/fixtures/guild-parity/
// report-*.json, written from design/mocks/guild_report_data.py, the same data the mock draws) and
// serves the guild page's prerendered shell at the raid URL, as the Worker does in production.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, type Page, type Route } from '@playwright/test';
import { characterSlug, guildHref } from '../../../../src/lib/characters';
import type { GuildRaidView } from '../../../../src/lib/guild/api';
import type { GuildPage } from '../../../../src/lib/rankings/api';
import type { GuildStateOpener } from '../guild-parity';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../fixtures/guild-parity');

function fixture<T>(name: string): T {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8')) as T;
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'guild-parity' } });

async function serveGuildShell(route: Route, guildUrl: string): Promise<void> {
  if (route.request().resourceType() !== 'document') return route.fallback();
  const response = await route.fetch({ url: new URL(guildUrl, route.request().url()).toString() });
  return route.fulfill({ response });
}

export const openReport: GuildStateOpener = async (page: Page) => {
  const raid = fixture<GuildRaidView>('report-raid.json');
  const guildPage = fixture<GuildPage>('report-guild.json');
  const { guild } = guildPage;
  const guildUrl = guildHref(guild.region, guild.ruleset, guild.name);
  const raidUrl = `${guildUrl}/raids/${raid.id}`;
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route(`**${raidUrl}`, (route) => serveGuildShell(route, guildUrl));
  await page.route(`**/v1/guilds/${guild.region}/${guild.ruleset}/${characterSlug(guild.name)}`, (route) =>
    route.fulfill(envelope(guildPage)),
  );
  await page.route('**/v1/me', (route) => route.fulfill(envelope(fixture('report-me.json'))));
  await page.route(`**/v1/guilds/${guild.id}/home`, (route) =>
    route.fulfill(envelope(fixture('report-home.json'))),
  );
  await page.route(`**/v1/guilds/${guild.id}/raids/${raid.id}`, (route) => route.fulfill(envelope(raid)));
  await page.goto(raidUrl);
  await expect(page.getByTestId('raid-report')).toBeVisible();
  await expect(page.getByTestId('raid-pull-row')).toHaveCount(raid.report?.pulls.length ?? 0);
};
