// web/tests/e2e/guild-raids-bot-only.spec.ts
// A guild that took the bot path and has no webhook (discord bot spec §3.1 state 5's primary
// path): the API's raid views report its Discord destination as connected, so the raid form
// offers `Post to Discord` (ticked from the remembered default) instead of the Settings line, and
// a done raid's report offers its Post button (raids spec §3.5 item 8; attendance spec §3.5).
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { AFTER_END, DONE, stubGuildAttendance } from './support/guild-attendance';
import { CHICAGO, GUILD_URL, NOW, ONYXIA_ID, raidUrl, stubGuildRaids, writesTo } from './support/guild-raids';
import { gotoHydrated } from './support/hydrated';

test.use({ timezoneId: CHICAGO });

/** The raid view of a bot-only guild: the bot is active and remembers the box ticked. */
const BOT_ONLY = { discordAvailable: true, discordConnected: true, postDefault: true } as const;

async function openRaidsTab(page: Page): ReturnType<typeof stubGuildRaids> {
  const stub = await stubGuildRaids(page, { viewer: 'officer', ...BOT_ONLY });
  await page.clock.install({ time: NOW });
  await gotoHydrated(page, `${GUILD_URL}#raids`, 'guild-raids-tab');
  return stub;
}

test.describe('a guild with only the bot', () => {
  test('the raid form offers Post to Discord, ticked, and the create carries it', async ({ page }) => {
    const stub = await openRaidsTab(page);
    await page.getByTestId('guild-raids-plan').click();
    await expect(page.getByTestId('raid-form-create')).toBeVisible();
    await expect(page.getByTestId('raid-form-discord')).toBeChecked();
    await expect(page.getByTestId('raid-form-discord-line')).toHaveCount(0);
    await page.getByTestId('raid-form-submit').click();
    await expect(page.getByTestId('raid-form-create')).toHaveCount(0);
    const posts = writesTo(stub, 'POST', /\/raids$/);
    expect(posts).toHaveLength(1);
    expect(posts[0].body).toMatchObject({ post_to_discord: true });
  });

  test('a done raid offers the report Post button, not the Settings line', async ({ page }) => {
    const stub = await stubGuildAttendance(page, { viewer: 'officer', nights: DONE, ...BOT_ONLY });
    await page.clock.install({ time: AFTER_END });
    await gotoHydrated(page, raidUrl(ONYXIA_ID), 'raid-page');
    await expect(page.getByTestId('raid-report-connect')).toHaveCount(0);
    await page.getByTestId('raid-report-post').click();
    await expect(page.getByTestId('raid-report-posted')).toContainText('Posted to Discord at');
    expect(stub.reportReady).toBe(1);
  });
});
