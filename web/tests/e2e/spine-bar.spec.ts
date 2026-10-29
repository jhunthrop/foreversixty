// web/tests/e2e/spine-bar.spec.ts
// Spec 2026-09-25 section 9, lane B (/logs left the spine's pages with the 2026-10-11 logs spec): one proof per door that it opens on the current
// character, plus one proving the bar's signed-out line. current-character.spec.ts,
// sim-landing.spec.ts and planner.spec.ts each already prove one door's own bootstrap
// (restore-from-pointer) in depth; this file proves the *bar itself* -- that each door's
// href actually carries the pointer, on every page it is supposed to mount on.
import { expect, test } from '@playwright/test';
import { ACTIVE_BUILD } from './support/active-build';

// An export that names its character: the only kind that becomes the current one.
const FURY = `FS1:${ACTIVE_BUILD}:warrior:orc:0/5530515/0:|who=Simfury:Realm`;

async function pasteAndGoTo(page: import('@playwright/test').Page, path: string) {
  await page.goto('/setup');
  await page.getByTestId('setup-try-export').locator('summary').click();
  await page.getByTestId('addon-paste-code').fill(FURY);
  await page.getByTestId('addon-paste-submit').click();
  await page.goto(path);
}

test.describe('the spine bar opens every door on the current character', () => {
  test("Plan and Sim carry the pasted export on /sim's bar", async ({ page }) => {
    await pasteAndGoTo(page, '/sim');
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

  test("Rankings for Warrior is on /sim's bar once a warrior export is loaded", async ({ page }) => {
    await pasteAndGoTo(page, '/sim');
    // The door carries both its labels in the DOM ("Rankings" below md, the full one from md),
    // so match the full one rather than the exact text.
    await expect(page.getByTestId('current-character-bar-rankings')).toContainText('Rankings for Warrior');
    await expect(page.getByTestId('current-character-bar-rankings')).toHaveAttribute(
      'href',
      '/rankings?class=warrior',
    );
  });

  // Rankings rebuild spec (2026-10-11): the page's one character switch is the shared header, so
  // the spine bar leaves it; the bar's own "Rankings for Warrior" door still opens the class's
  // board through the URL.
  test("/rankings has no spine bar of its own, and the bar's class door filters the board", async ({
    page,
  }) => {
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
    await pasteAndGoTo(page, '/rankings/warden-kelthas?class=warrior');
    await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
    await expect.poll(() => requestedClass).toBe('warrior');
  });

  test("the bar mounts on /sim, above the page's own chip", async ({ page }) => {
    await pasteAndGoTo(page, '/sim');
    await expect(page.getByTestId('current-character-bar')).toBeVisible();
  });

  // Planner rebuild spec §3: the spine band is replaced on this one page by the new
  // shared character header -- the pointer bridge itself is unchanged (`decidePlannerLoad`
  // still restores a bare /planner from the pasted export), so this proves the restore
  // rather than a chip that no longer renders here.
  test('/planner restores the pasted export on a bare load, with no spine chip of its own', async ({
    page,
  }) => {
    await pasteAndGoTo(page, '/planner');
    await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
    await expect(page.getByLabel('Class')).toHaveValue('warrior');
    // "5530515" base-36 digits sum to 24 points in Fury, nothing dropped (the same fixture
    // string current-character.spec.ts's own restore test and sim-saved.spec.ts's fixture
    // result, "0-5530515-", both already rely on).
    await expect(page.getByTestId('tree-points-164')).toHaveText('24');
  });
});

// The tool pages (/sim/gear and its siblings) say it once, as /sim does: the bare signed-out
// sentence and its underlined links are left out, and the Your characters card keeps the buttons.
test('a tool page leaves the bare sign-in line out; its account card carries the buttons', async ({
  page,
}) => {
  await page.goto('/sim/gear');
  await expect(page.getByTestId('sim-signin').first()).toBeVisible();
  await expect(page.getByTestId('current-character-bar-signed-out')).toHaveCount(0);
});

test('the simulator leaves the bare sign-in line out; its account card carries the buttons', async ({
  page,
}) => {
  await page.goto('/sim');
  await expect(page.getByTestId('sim-signin')).toBeVisible();
  await expect(page.getByTestId('current-character-bar-signed-out')).toHaveCount(0);
});
