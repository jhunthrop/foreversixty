// web/tests/e2e/pages.spec.ts
// The changelog and the 404 in the flat system (design/specs/2026-11-08-pages.md §4.1, §4.4, §5):
// the changelog's Show filter (a radio group, `?show=` restored before paint, Coming up and the
// months it empties hidden), New since the last visit, a folded entry opened by its permalink;
// the 404's failed address, Did you mean, the class picker's spec links, the lookup first on a
// dead /character/ address, and its answer signed in. One Plate on the 404, none on the reading
// pages.
import { expect, test } from './support/test';
import { openPagesState, routePagesState } from './support/pages-page';

test.describe('the changelog', () => {
  test('Show is a radio group: All by default, Game and Site rewrite ?show= and filter in place', async ({
    page,
  }) => {
    await page.setViewportSize({ width: 1440, height: 1000 });
    await page.goto('/changelog');
    const show = page.getByRole('radiogroup', { name: 'Show' });
    await expect(show.getByRole('radio', { name: /^All/ })).toHaveAttribute('aria-checked', 'true');
    await expect(page.locator('[data-show="site"]').first()).toBeVisible();
    await show.getByRole('radio', { name: /^Game/ }).click();
    await expect(page).toHaveURL(/\?show=game$/);
    await expect(show.getByRole('radio', { name: /^Game/ })).toHaveAttribute('aria-checked', 'true');
    await expect(page.locator('[data-show="site"]').first()).toBeHidden();
    await expect(page.locator('[data-show="game"]').first()).toBeVisible();
    await expect(page.getByTestId('changelog-coming-up')).toBeVisible();
    await show.getByRole('radio', { name: /^Site/ }).click();
    await expect(page).toHaveURL(/\?show=site$/);
    await expect(page.getByTestId('changelog-coming-up')).toBeHidden();
    await expect(page.locator('[data-show="game"]').first()).toBeHidden();
  });

  test('?show=site loads filtered before paint, its counts per month', async ({ page }) => {
    await page.setViewportSize({ width: 1440, height: 1000 });
    await page.goto('/changelog?show=site');
    await expect(page.locator('html')).toHaveAttribute('data-cl-show', 'site');
    await expect(
      page.getByRole('radiogroup', { name: 'Show' }).getByRole('radio', { name: /^Site/ }),
    ).toHaveAttribute('aria-checked', 'true');
    const november = page.getByTestId('changelog-month-2026-11');
    await expect(november.locator('[data-cl-for="site"]')).toBeVisible();
    await expect(november.locator('[data-cl-for="all"]')).toBeHidden();
  });

  test('every entry from launch on shows its text; a pre-launch entry opens from its title row', async ({
    page,
  }) => {
    await page.goto('/changelog');
    const launch = page.locator('[id="2026-11-04-launch"]');
    await expect(launch).toContainText('3:00 p.m. PST');
    await expect(launch.getByRole('link', { name: 'Blizzard, pre-purchase and beta dates' })).toBeVisible();
    const guild = page.locator('[id="2026-11-05-guild-events"]');
    await expect(guild.getByRole('link', { name: 'Open the guild page ›' })).toHaveAttribute(
      'href',
      '/guild',
    );
    const coach = page.locator('[id="2026-10-28-log-coach"]');
    await expect(coach.getByText('Added the log coach for Premium.', { exact: false })).toBeHidden();
    await coach.locator('summary').click();
    await expect(coach.getByText('Added the log coach for Premium.', { exact: false })).toBeVisible();
  });

  test('a returning visitor sees New on the entries since the last visit, and nothing else changes', async ({
    page,
  }) => {
    await page.addInitScript(() => {
      try {
        localStorage.setItem('fs.changelogSeen', '2026-11-04');
      } catch {
        // a blocked store reads as a first visit
      }
    });
    await page.goto('/changelog');
    await expect(page.locator('[id="2026-11-05-guild-events"] [data-cl-new]')).toBeVisible();
    await expect(page.locator('[id="2026-11-04-launch"] [data-cl-new]')).toBeHidden();
  });

  test('draws no Plate and one h1', async ({ page }) => {
    await page.goto('/changelog');
    await expect(page.locator('[data-region="plate"]')).toHaveCount(0);
    await expect(page.locator('h1')).toHaveCount(1);
  });
});

test.describe('the 404', () => {
  test('signed out: the failed address, the line, the class picker as the one Plate, then the lookup', async ({
    page,
  }) => {
    await openPagesState(page, '/this-page-does-not-exist', 'out');
    await expect(page.getByRole('heading', { level: 1 })).toHaveText('No page here');
    await expect(page.getByTestId('nf-path')).toHaveText('/this-page-does-not-exist');
    await expect(page.getByTestId('nf-suggest')).toBeHidden();
    await expect(page.locator('[data-region="plate"]')).toHaveCount(1);
    await expect(page.getByTestId('nf-answer')).toBeHidden();
    const hunter = page.getByTestId('not-found-class-hunter');
    await expect(hunter.getByRole('link', { name: 'Hunter' })).toHaveAttribute('href', '/guides/hunter');
    await expect(hunter.getByRole('link', { name: 'Beast Mastery' })).toHaveAttribute(
      'href',
      '/guides/hunter/beast-mastery',
    );
    const picker = (await page.getByTestId('not-found-classes').boundingBox())!;
    const lookup = (await page.getByTestId('home-character-search').boundingBox())!;
    expect(lookup.y).toBeGreaterThan(picker.y);
  });

  test('a mistyped guide address suggests the guide', async ({ page }) => {
    await openPagesState(page, '/guides/hunter/beastmastery', 'out');
    const suggestion = page.getByTestId('nf-suggest');
    await expect(suggestion).toBeVisible();
    await expect(suggestion).toContainText('Did you mean');
    await expect(suggestion.getByRole('link', { name: 'Beast Mastery Hunter guide ›' })).toHaveAttribute(
      'href',
      '/guides/hunter/beast-mastery',
    );
  });

  test('signed out on a dead /character/ address, the lookup comes first', async ({ page }) => {
    await openPagesState(page, '/character/us/normal', 'out');
    const picker = (await page.getByTestId('not-found-classes').boundingBox())!;
    const lookup = (await page.getByTestId('home-character-search').boundingBox())!;
    expect(lookup.y).toBeLessThan(picker.y);
  });

  test('signed in: your character, your guide, your best in slot, then the smaller Your character', async ({
    page,
  }) => {
    await page.setViewportSize({ width: 1440, height: 1000 });
    await openPagesState(page, '/this-page-does-not-exist', 'in');
    const answer = page.getByTestId('nf-answer');
    await expect(answer).toHaveAttribute('data-state', 'ready');
    await expect(answer).toContainText('Felsnap');
    const links = answer.locator('[data-field="links"] a');
    await expect(links).toHaveText([
      'Your guide: Beast Mastery Hunter ›',
      'Your best in slot ›',
      'Your character ›',
    ]);
    await expect(links.nth(1)).toHaveAttribute('href', '/bis/hunter/beast-mastery?band=30');
    await expect(links.nth(2)).toHaveAttribute('href', '/character/us/normal/felsnap');
    await expect(links.nth(2)).toHaveAttribute('data-size', 'small');
  });

  test('the account read failing says so, with Try again', async ({ page }) => {
    await openPagesState(page, '/this-page-does-not-exist', 'error');
    const answer = page.getByTestId('nf-answer');
    await expect(answer).toHaveAttribute('data-state', 'error');
    await expect(answer.getByRole('alert')).toHaveText('We could not load your character.');
    await expect(answer.getByRole('button', { name: 'Try again' })).toBeVisible();
  });

  test('signed in with no character: Get set up', async ({ page }) => {
    await routePagesState(page, 'none');
    await page.goto('/this-page-does-not-exist');
    const answer = page.getByTestId('nf-answer');
    await expect(answer).toHaveAttribute('data-state', 'none');
    await expect(answer.getByRole('link', { name: 'Get set up ›' })).toHaveAttribute('href', '/setup');
  });
});
