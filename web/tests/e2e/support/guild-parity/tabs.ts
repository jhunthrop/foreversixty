// web/tests/e2e/support/guild-parity/tabs.ts
// The guild parity state for the Progression tab (GUILD-ALL lane), as the officer (Dunmaro) on the
// guild page's hash tab. The Overview's routing answers every read (support/guild-parity/overview.ts);
// the Roster and Readiness states are the dense ones (support/guild-parity/dense.ts).
import type { Page } from '@playwright/test';
import { expect } from '../test';
import type { GuildStateOpener } from '../guild-parity';
import { routeGuild } from './overview';

const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';

async function openTab(page: Page, hash: string, ready: string): Promise<void> {
  await routeGuild(page, 'officer');
  await page.goto(`${GUILD_PATH}#${hash}`);
  await expect(page.getByTestId(ready)).toBeVisible();
  await page.waitForLoadState('networkidle');
}

export const TABS_OPENERS: Readonly<Record<string, GuildStateOpener>> = {
  progression: (page) => openTab(page, 'progression', 'guild-progression-encounter'),
};
