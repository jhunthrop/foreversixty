// web/tests/e2e/tab-strip.spec.ts
// The shared TabStrip in a real browser (site alignment spec 2026-10-29 D3), on its first two
// homes: the guild page's in-page tabs (keyboard, roving tabindex, aria-selected) and the
// simulator's link tabs (aria-current). Phone scroll and fades: guild-phone and sim-tabs specs.
import { apiUnreachable, expect, test } from './support/test';
import { stubGuildAttendance } from './support/guild-attendance';
import { GUILD_URL } from './support/guild-raids';

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/guilds/*/settings'));

test.describe('guild tabs, keyboard', () => {
  test.beforeEach(async ({ page }) => {
    await stubGuildAttendance(page, { viewer: 'officer', officerTools: true });
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-tab-overview')).toHaveAttribute('aria-selected', 'true');
  });

  test('only the current tab is a tab stop', async ({ page }) => {
    const tabs = page.getByTestId('guild-tabs').getByRole('tab');
    const count = await tabs.count();
    expect(count).toBeGreaterThan(2);
    for (let index = 0; index < count; index += 1) {
      await expect(tabs.nth(index)).toHaveAttribute('tabindex', index === 0 ? '0' : '-1');
    }
  });

  test('arrows move, select and route; Home and End jump; ArrowLeft wraps', async ({ page }) => {
    await page.getByTestId('guild-tab-overview').focus();
    await page.keyboard.press('ArrowRight');
    const roster = page.getByTestId('guild-tab-roster');
    await expect(roster).toBeFocused();
    await expect(roster).toHaveAttribute('aria-selected', 'true');
    await expect(roster).toHaveAttribute('tabindex', '0');
    await expect(page.getByTestId('guild-tab-overview')).toHaveAttribute('aria-selected', 'false');
    await expect(page).toHaveURL(/#roster$/);

    await page.keyboard.press('End');
    await expect(page.getByTestId('guild-tab-settings')).toBeFocused();
    await expect(page.getByTestId('guild-tab-settings')).toHaveAttribute('aria-selected', 'true');

    await page.keyboard.press('Home');
    await expect(page.getByTestId('guild-tab-overview')).toBeFocused();

    await page.keyboard.press('ArrowLeft');
    await expect(page.getByTestId('guild-tab-settings')).toBeFocused();
    await expect(page).toHaveURL(/#settings$/);
  });

  test('Tab leaves the strip from the current tab', async ({ page }) => {
    await page.getByTestId('guild-tab-overview').focus();
    await page.keyboard.press('Tab');
    await expect(page.getByTestId('guild-tabs').getByRole('tab').nth(1)).not.toBeFocused();
  });
});

test('simulator tabs are links, the current one aria-current', async ({ page }) => {
  await page.goto('/sim/gear');
  const nav = page.getByRole('navigation', { name: /simulator/i });
  await expect(nav.getByRole('tab')).toHaveCount(0);
  await expect(page.getByTestId('sim-nav-tab-gear')).toHaveAttribute('aria-current', 'page');
  await expect(page.getByTestId('sim-nav-tab-quick-sim')).not.toHaveAttribute('aria-current', 'page');
});
