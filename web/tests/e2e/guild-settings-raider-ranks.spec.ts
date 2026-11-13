// web/tests/e2e/guild-settings-raider-ranks.spec.ts
// The Settings tab's raider ranks row: the saved ranks ticked, a tick and Save send the PATCH with
// the shared spinner and "Saving…", "Saved." after, and clearing every box sends an empty list.
import { apiUnreachable, expect, test, GUILD_OVERVIEW_READS } from './support/test';
import { gotoHydrated } from './support/hydrated';
import { RAIDER_RANKS_URL, stubRaiderRanks } from './support/guild-raider-ranks';

test.beforeEach(({ page }) => apiUnreachable(page, ...GUILD_OVERVIEW_READS, '**/v1/guilds/*/billing'));

const check = (rank: number) => `guild-raider-rank-${rank}-check`;

test('an officer ticks the raider ranks and saves them', async ({ page }) => {
  let release = (): void => {};
  const hold = new Promise<void>((resolve) => (release = resolve));
  const bodies = await stubRaiderRanks(page, { saved: [2], hold });
  await gotoHydrated(page, RAIDER_RANKS_URL, 'guild-raider-ranks-list');

  await expect(page.getByTestId(check(2))).toBeChecked();
  await expect(page.getByTestId(check(3))).not.toBeChecked();
  const summary = page.getByTestId('guild-raider-ranks-summary');
  await expect(summary).toHaveText('14 characters, 13 at 60');
  await expect(page.getByTestId('guild-raider-ranks-save')).toBeDisabled();

  await page.getByTestId(check(3)).check();
  await page.getByTestId(check(4)).check();
  // The line follows the boxes before the save.
  await expect(summary).toHaveText('27 characters, 22 at 60');
  await page.getByTestId('guild-raider-ranks-save').click();
  await expect(page.getByTestId('guild-raider-ranks-save')).toContainText('Saving…');
  await expect(page.getByTestId(check(5))).toBeDisabled();
  release();

  await expect(page.getByTestId('guild-raider-ranks-saved')).toHaveText('Saved.');
  expect(bodies).toEqual([{ raider_rank_indexes: [2, 3, 4] }]);
  await expect(page.getByTestId('guild-raider-ranks-save')).toBeDisabled();

  // The next change clears Saved.
  await page.getByTestId(check(4)).uncheck();
  await expect(page.getByTestId('guild-raider-ranks-saved')).toHaveCount(0);
});

test('clearing every box goes back to every level 60', async ({ page }) => {
  const bodies = await stubRaiderRanks(page, { saved: [2, 3] });
  await gotoHydrated(page, RAIDER_RANKS_URL, 'guild-raider-ranks-list');
  await page.getByTestId(check(2)).uncheck();
  await page.getByTestId(check(3)).uncheck();
  await expect(page.getByTestId('guild-raider-ranks-summary')).toHaveText(
    'Not set. The addon counts every level 60 as a raider.',
  );
  await page.getByTestId('guild-raider-ranks-save').click();
  await expect(page.getByTestId('guild-raider-ranks-saved')).toBeVisible();
  expect(bodies).toEqual([{ raider_rank_indexes: [] }]);
  await expect(page.getByTestId(check(2))).not.toBeChecked();
});

test('a contested claim locks the boxes', async ({ page }) => {
  await stubRaiderRanks(page, { saved: [2], frozen: true });
  await gotoHydrated(page, RAIDER_RANKS_URL, 'guild-raider-ranks-list');
  await expect(page.getByTestId(check(2))).toBeDisabled();
  await expect(page.getByTestId('guild-raider-ranks-save')).toBeDisabled();
});
