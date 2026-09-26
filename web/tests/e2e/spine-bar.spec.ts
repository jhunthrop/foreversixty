// web/tests/e2e/spine-bar.spec.ts
// Spec 2026-09-25 section 9, lane B: one proof per door that it opens on the current
// character, plus one proving the bar's signed-out line. current-character.spec.ts,
// sim-landing.spec.ts and planner.spec.ts each already prove one door's own bootstrap
// (restore-from-pointer) in depth; this file proves the *bar itself* -- that each door's
// href actually carries the pointer, on every page it is supposed to mount on.
import { expect, test } from '@playwright/test';
import { ACTIVE_BUILD } from './support/active-build';

const FURY = `FS1:${ACTIVE_BUILD}:warrior:orc:0/5530515/0:`;

async function pasteAndGoTo(page: import('@playwright/test').Page, path: string) {
  await page.goto('/setup');
  await page.getByTestId('addon-paste-code').fill(FURY);
  await page.getByTestId('addon-paste-submit').click();
  await page.goto(path);
}

test.describe('the spine bar opens every door on the current character', () => {
  test("Plan and Sim carry the pasted export on /logs' bar", async ({ page }) => {
    await pasteAndGoTo(page, '/logs');
    await expect(page.getByTestId('current-character-bar')).toBeVisible();
    await expect(page.getByTestId('current-character-bar-plan')).toHaveAttribute(
      'href',
      `/planner?code=${encodeURIComponent(FURY)}`,
    );
    await expect(page.getByTestId('current-character-bar-sim')).toHaveAttribute(
      'href',
      expect.stringContaining(encodeURIComponent(FURY)),
    );
  });

  test("Rankings for Warrior is on /logs' bar once a warrior export is loaded", async ({ page }) => {
    await pasteAndGoTo(page, '/logs');
    await expect(page.getByTestId('current-character-bar-rankings')).toHaveText('Rankings for Warrior');
    await expect(page.getByTestId('current-character-bar-rankings')).toHaveAttribute(
      'href',
      '/rankings?class=warrior',
    );
  });

  test('the bar mounts on /rankings and pre-fills the class filter from the pointer', async ({ page }) => {
    // Rankings.svelte has no visible <select> for `class` (it is URL/link-driven only, per
    // Rankings.svelte's own template read during planning) -- the request itself is the
    // proof the prefill reached the fetch, not a visible control.
    let requestedClass: string | null = null;
    await page.route('**/v1/rankings?*', (route) => {
      requestedClass = new URL(route.request().url()).searchParams.get('class');
      return route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          data: { rows: [], total: 0, per_page: 25, updated_at: '2026-09-25T00:00:00Z' },
          error: null,
        }),
      });
    });
    await pasteAndGoTo(page, '/rankings/warden-kelthas');
    await expect(page.getByTestId('current-character-bar')).toBeVisible();
    await expect.poll(() => requestedClass).toBe('warrior');
  });

  test("the bar mounts on /planner and /sim, above each page's own chip", async ({ page }) => {
    await pasteAndGoTo(page, '/planner');
    await expect(page.getByTestId('current-character-bar')).toBeVisible();
    await pasteAndGoTo(page, '/sim');
    await expect(page.getByTestId('current-character-bar')).toBeVisible();
  });
});

// Fixed after this spec first found it broken (see the ledger, "Task 10 follow-up"): the
// signed-out branch no longer sits inside `.chip-slot` (CurrentCharacterBar.svelte) -- that
// class is hidden pre-paint whenever neither `data-pointer` nor `data-session` is set, which
// is exactly the one visitor shape this line greets, so hiding it made the line permanently
// unreachable rather than merely not pre-reserving space for it.
test('the bar shows the signed-out, no-pointer line with two links, on a fresh visit', async ({ page }) => {
  await page.goto('/logs');
  const bar = page.getByTestId('current-character-bar-signed-out');
  await expect(bar).toBeVisible();
  await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toHaveAttribute('href', '/login');
  await expect(page.getByRole('link', { name: 'Paste an export' })).toHaveAttribute('href', '/setup#paste');
});
