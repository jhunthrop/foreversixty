// web/tests/e2e/guild-loot.spec.ts
// The Loot tab as the loot council (web layout spec 2026-11-05 §7.2), on the mocked API
// (support/officer-tools.ts stubLootCouncil): an officer picks a drop, votes, awards with MS/OS and a
// reason, reads the record and exports the history; a member flags their own MS/OS and never sees a
// vote; a failed read says so with Try again; nothing scrolls sideways on a phone.
import { expect, test, apiUnreachable, GUILD_OVERVIEW_READS } from './support/test';
import { gotoHydrated } from './support/hydrated';
import { choose } from './support/select';
import { GUILD_ID } from '../../src/fixtures/guild/mock-guild';
import { MOCK_CURRENT_NIGHT, buildMockLootNights } from '../../src/fixtures/guild/mock-loot-council';
import {
  GUILD_URL,
  awardedItemId,
  envelope,
  stubGuildCentre,
  unawardedItem,
  type Viewer,
} from './support/officer-tools';

test.beforeEach(({ page }) => apiUnreachable(page, ...GUILD_OVERVIEW_READS));

/** The member's hunter can use this helm; it is not on the mock night yet. */
const DRAGONSTALKER_HELM = 16939;
const MEMBER_KEY = 'us/pvp/zulmara';

async function openLoot(page: import('@playwright/test').Page, viewer: Viewer): Promise<void> {
  await stubGuildCentre(page, { viewer, officerTools: true, raidGuildTools: true, claimHolder: true });
  await gotoHydrated(page, `${GUILD_URL}#loot`, 'guild-loot-tab');
  await expect(page.getByTestId('guild-loot-drops')).toBeVisible();
}

test.describe('officer', () => {
  test('picks a drop, votes, and awards it with main or off spec and a reason', async ({ page }) => {
    const awards: unknown[] = [];
    await page.route(`**/v1/guilds/${GUILD_ID}/loot/awards`, (route) => {
      awards.push(route.request().postDataJSON());
      return route.fulfill(envelope({ award_id: 77, status: 'awarded' }));
    });
    await openLoot(page, 'officer');
    const night = buildMockLootNights().nights[0];
    await expect(page.getByTestId('guild-loot-night-summary')).toHaveText(
      'Tue 15 Dec · 3 drops, 2 open · 2 recorded by the addon',
    );
    await expect(page.getByTestId('guild-loot-drops').getByText('Open', { exact: true })).toHaveCount(2);

    const { itemId, characterKey } = unawardedItem();
    await page.getByTestId(`guild-loot-drop-${itemId}`).click();
    await expect(page.getByTestId(`guild-loot-drop-${itemId}`)).toHaveAttribute('aria-pressed', 'true');
    const row = page.getByTestId(`guild-loot-candidate-${characterKey}`);
    await expect(row).toContainText('DPS over');
    await expect(row.getByTestId('guild-loot-bis')).toHaveText('BiS for them');

    const second = page.locator('[data-testid^="guild-loot-candidate-"]').nth(1);
    const yes = second.getByTestId('guild-loot-votes-yes');
    await expect(yes).toHaveAttribute('aria-pressed', 'false');
    await yes.click();
    await expect(yes).toHaveAttribute('aria-pressed', 'true');

    await row.getByTestId('guild-loot-award').click();
    const confirm = row.getByTestId('guild-loot-award-confirm');
    await expect(confirm.getByTestId('guild-loot-award-cancel')).toBeFocused();
    await confirm.getByTestId('guild-loot-award-msos-os').click();
    await confirm.getByTestId('guild-loot-award-reason').fill('  Main spec passed  ');
    await confirm.getByTestId('guild-loot-award-yes').click();
    await expect.poll(() => awards.length).toBe(1);
    expect(awards[0]).toEqual({
      night: night.night,
      item_id: itemId,
      character_key: characterKey,
      encounter_id: 1084,
      ms_os: 'os',
      reason: 'Main spec passed',
    });
  });

  test('an awarded drop leads with its record, and the past night keeps its own', async ({ page }) => {
    await openLoot(page, 'officer');
    // The addon recorded this drop: its row names who took it from the loot window, beside the award.
    await expect(page.getByTestId(`guild-loot-drop-${awardedItemId()}`)).toContainText('Looted by');
    await page.getByTestId(`guild-loot-drop-${awardedItemId()}`).click();
    await expect(page.getByTestId('guild-loot-closed')).toBeVisible();
    await expect(page.getByTestId('guild-loot-candidates')).toContainText('3 yes, 0 no of 4');
    await expect(page.getByTestId('guild-loot-candidates').getByTestId('guild-loot-award')).toHaveCount(0);

    // On a phone the bar is one line: Change opens its choices.
    const change = page.locator('#guild-loot-bar-groups').locator('..').locator('[data-control-bar-toggle]');
    if (await change.isVisible()) await change.click();
    await page.getByTestId('control-night').getByRole('radio', { name: 'Tue 8 Dec' }).click();
    await expect(page.getByTestId('guild-loot-night-summary')).toHaveText('Tue 8 Dec · 1 drop, 0 open');
    await expect(page.getByTestId('guild-loot-open-drop')).toHaveCount(0);
    await expect(page.getByTestId('guild-loot-decisions')).toContainText('by Kraggor');
  });

  test('opens a loot table drop on the current night and exports the history in a range', async ({
    page,
  }) => {
    await openLoot(page, 'officer');
    await choose(page.locator('#guild-loot-open-drop'), String(DRAGONSTALKER_HELM));
    await expect(page.getByTestId('guild-loot-candidates')).toContainText(
      "Candidates for Dragonstalker's Helm",
    );

    await page.getByTestId('guild-loot-export-toggle').click();
    await page.getByTestId('guild-loot-export-from').fill('2026-12-01');
    await expect(page.getByTestId('guild-loot-export-csv')).toHaveAttribute(
      'href',
      new RegExp(`/v1/guilds/${GUILD_ID}/export/loot\\.csv\\?from=2026-12-01$`),
    );
    await expect(page.getByTestId('guild-loot-export-json')).toHaveAttribute(
      'href',
      /export\/loot\.json\?from=2026-12-01$/,
    );
  });
});

test.describe('member', () => {
  test('flags their own main or off spec and never sees a vote', async ({ page }) => {
    const flags: unknown[] = [];
    await openLoot(page, 'member');
    // Registered after the stubs, so it answers the flag before their echo does.
    await page.route(`**/v1/guilds/${GUILD_ID}/loot/interest`, (route) => {
      if (route.request().method() !== 'PUT') return route.fallback();
      const body = route.request().postDataJSON() as Record<string, unknown>;
      flags.push(body);
      return route.fulfill(envelope(body));
    });
    await expect(page.getByTestId('guild-loot-tab')).not.toContainText(' yes, ');
    await expect(page.getByTestId('guild-loot-export-toggle')).toHaveCount(0);

    await choose(page.locator('#guild-loot-open-drop'), String(DRAGONSTALKER_HELM));
    const flag = page.getByTestId(`guild-loot-flag-${MEMBER_KEY}-choice`);
    await expect(flag.getByRole('radio', { name: 'None' })).toHaveAttribute('aria-checked', 'true');
    await flag.getByRole('radio', { name: 'MS' }).click();
    await expect.poll(() => flags.length).toBe(1);
    expect(flags[0]).toEqual({
      night: MOCK_CURRENT_NIGHT,
      item_id: DRAGONSTALKER_HELM,
      character_key: MEMBER_KEY,
      ms_os: 'ms',
    });
    await expect(flag.getByRole('radio', { name: 'MS' })).toHaveAttribute('aria-checked', 'true');
    await expect(page.locator('[aria-label^="Vote "]')).toHaveCount(0);
  });
});

test('a failed read says Loot did not load, with Try again', async ({ page }) => {
  await stubGuildCentre(page, { viewer: 'officer', officerTools: true, raidGuildTools: true });
  await page.route(`**/v1/guilds/${GUILD_ID}/loot/nights*`, (route) =>
    route.fulfill({
      status: 500,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":null}',
    }),
  );
  await gotoHydrated(page, `${GUILD_URL}#loot`, 'guild-loot-tab');
  await expect(page.getByTestId('guild-loot-error')).toContainText('Loot did not load.');
  await expect(page.getByTestId('guild-loot-error').getByRole('button')).toHaveText('Try again');
});

test('at 390 the drops come first and nothing scrolls sideways', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 844 });
  await openLoot(page, 'officer');
  await page.getByTestId(`guild-loot-drop-${unawardedItem().itemId}`).click();
  await expect(page.getByTestId('guild-loot-candidates')).toBeVisible();
  const drops = await page.getByTestId('guild-loot-drops').boundingBox();
  const rail = await page.getByTestId('guild-loot-candidates').boundingBox();
  expect(rail!.y).toBeGreaterThan(drops!.y);
  expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(390);
});
