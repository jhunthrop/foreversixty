// web/tests/e2e/support/guild-parity/tabs.ts
// The guild parity states for the Roster, Progression and Readiness tabs (GUILD-ALL lane), each as
// the officer (Dunmaro) on the guild page's hash tab. The Overview's routing answers every read
// (support/guild-parity/overview.ts); the Roster's home carries a waiting list
// (design/mocks/guild_tabs_fixtures.py).
import type { Page } from '@playwright/test';
import { expect } from '../test';
import type { GuildStateOpener } from '../guild-parity';
import { routeGuild } from './overview';

const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';

async function openTab(page: Page, hash: string, ready: string, home?: string): Promise<void> {
  await routeGuild(page, 'officer', home);
  await page.goto(`${GUILD_PATH}#${hash}`);
  await expect(page.getByTestId(ready)).toBeVisible();
  await page.waitForLoadState('networkidle');
}

export const TABS_OPENERS: Readonly<Record<string, GuildStateOpener>> = {
  roster: (page) => openTab(page, 'roster', 'guild-roster-waiting', 'tabs-home-officer.json'),
  progression: (page) => openTab(page, 'progression', 'guild-progression-encounter'),
  readiness: (page) => openTab(page, 'readiness', 'guild-readiness-rows'),
};
