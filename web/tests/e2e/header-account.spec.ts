// web/tests/e2e/header-account.spec.ts
// The header's account control (design/specs/2026-10-28-polish.md section 1): the character
// selector's panel gains the account head and the Your profile, Account, Plan and Sign out rows;
// signed out it leads with Sign in with Battle.net. Open and close, the keyboard, 44 px rows,
// the sheet fitting 360 px and the bar not moving when the panel opens.
import { expect, test, type Page } from '@playwright/test';
import { OBNOXIOUS, armoryPointer, signInWith, signedOut, storePointer } from './support/selector';

const FOUNDER_BILLING = {
  plan: 'premium',
  status: 'active',
  current_period_end: '2027-10-27T00:00:00Z',
  cancel_at_period_end: false,
  founder: true,
  interval: 'yearly',
  amount_cents: 5000,
};
const PREMIUM_ME = {
  user: { battletag: 'Hunthrop#1894' },
  entitlements: {
    server_sims: true,
    retention: true,
    multi_compare: true,
    history: true,
    notifications: true,
    officer_views: false,
    roster_check: true,
    log_coach: true,
    supporter_mark: true,
    billing: FOUNDER_BILLING,
  },
};
const MIN_TARGET = 44;
const ROW_MIN = 52;

const control = (page: Page) => page.getByTestId('character-selector');
const panel = (page: Page) => page.getByTestId('selector-panel');
const rows = (page: Page) => panel(page).getByTestId('selector-account-rows');

async function signedInPage(page: Page, me: Record<string, unknown> = {}): Promise<void> {
  await signInWith(page, [OBNOXIOUS], [], me);
  await storePointer(page, armoryPointer(OBNOXIOUS));
  await page.goto('/guides');
  await expect(control(page)).toBeVisible();
}

test.describe('the account panel, signed in', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  test('opens with the head, the characters and the four account rows, in that order', async ({ page }) => {
    await signedInPage(page, PREMIUM_ME);
    await control(page).click();
    await expect(panel(page)).toBeVisible();
    await expect(panel(page).getByTestId('selector-account-name')).toHaveText('Hunthrop#1894');
    await expect(panel(page).getByTestId('selector-account-head')).toContainText('Signed in with Battle.net');
    await expect(panel(page).getByTestId('selector-premium-tag')).toHaveText('Premium');
    const labels = await rows(page).locator('.csel-arow-title').allTextContents();
    expect(labels).toEqual(['Your profile', 'Account', 'Plan', 'Sign out']);
    await expect(rows(page).getByTestId('selector-row-profile')).toHaveAttribute(
      'href',
      '/character/us/normal/obnoxious-yell',
    );
    await expect(rows(page).getByTestId('selector-row-account')).toHaveAttribute('href', '/account');
    await expect(rows(page).getByTestId('selector-row-plan')).toHaveAttribute('href', '/account#plan');
    await expect(rows(page).getByTestId('selector-row-plan')).toContainText('Founder');
    await expect(rows(page).getByTestId('selector-row-plan')).toContainText(/Premium · renews/);
    const head = (await panel(page).getByTestId('selector-account-head').boundingBox())!;
    const list = (await panel(page).getByTestId('selector-list').boundingBox())!;
    const account = (await rows(page).boundingBox())!;
    expect(head.y).toBeLessThan(list.y);
    expect(list.y).toBeLessThan(account.y);
  });

  test('a free account reads Free on Plan and goes to /premium, with no Premium tag', async ({ page }) => {
    await signedInPage(page);
    await control(page).click();
    await expect(panel(page).getByTestId('selector-premium-tag')).toHaveCount(0);
    await expect(rows(page).getByTestId('selector-row-plan')).toContainText('Free');
    await expect(rows(page).getByTestId('selector-row-plan')).toHaveAttribute('href', '/premium');
  });

  test('closes on Escape with focus back on the control, and on a second click', async ({ page }) => {
    await signedInPage(page);
    await control(page).click();
    await expect(panel(page)).toBeVisible();
    await page.keyboard.press('Escape');
    await expect(panel(page)).toBeHidden();
    await expect(control(page)).toBeFocused();
    await control(page).click();
    await expect(panel(page)).toBeVisible();
    await control(page).click();
    await expect(panel(page)).toBeHidden();
  });

  test('Enter opens it, Tab walks Paste then each account row, Enter on a row follows it', async ({
    page,
  }) => {
    await signedInPage(page);
    await control(page).focus();
    await page.keyboard.press('Enter');
    await expect(panel(page)).toBeVisible();
    await expect(panel(page).locator('.csel-row-main[aria-current="true"]')).toBeFocused();
    await page.keyboard.press('Tab');
    await expect(panel(page).getByTestId('selector-paste')).toBeFocused();
    await page.keyboard.press('Tab');
    await expect(rows(page).getByTestId('selector-row-profile')).toBeFocused();
    await page.keyboard.press('Tab');
    await expect(rows(page).getByTestId('selector-row-account')).toBeFocused();
    await page.keyboard.press('Enter');
    await expect(page).toHaveURL(/\/account$/);
  });

  test('every account row is at least 52 px tall and a 44 px target', async ({ page }) => {
    await signedInPage(page);
    await control(page).click();
    for (const row of await rows(page).locator('.csel-arow').all()) {
      const box = (await row.boundingBox())!;
      expect(box.height).toBeGreaterThanOrEqual(ROW_MIN);
      expect(box.height).toBeGreaterThanOrEqual(MIN_TARGET);
    }
  });

  test('opening the panel shifts nothing in the bar', async ({ page }) => {
    await signedInPage(page);
    const before = await page.getByRole('banner').boundingBox();
    const controlBefore = await control(page).boundingBox();
    await control(page).click();
    await expect(panel(page)).toBeVisible();
    expect(await page.getByRole('banner').boundingBox()).toEqual(before);
    expect(await control(page).boundingBox()).toEqual(controlBefore);
  });

  test('the panel hangs from the control, never past the 1344 column at 2000', async ({ page }) => {
    await page.setViewportSize({ width: 2000, height: 900 });
    await signedInPage(page);
    await control(page).click();
    const box = (await panel(page).boundingBox())!;
    const bar = (await control(page).boundingBox())!;
    expect(Math.abs(box.x + box.width - (bar.x + bar.width))).toBeLessThanOrEqual(1);
    expect(box.width).toBe(380);
  });

  test('Sign out calls the session endpoint and reloads signed out', async ({ page }) => {
    await signedInPage(page);
    let deleted = false;
    await page.route('**/v1/sessions', (route) => {
      deleted = route.request().method() === 'DELETE';
      return route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: '{"ok":true,"data":null,"error":null,"request_id":"r"}',
      });
    });
    await control(page).click();
    await Promise.all([page.waitForNavigation(), rows(page).getByTestId('selector-row-sign-out').click()]);
    expect(deleted).toBe(true);
    await expect(page).toHaveURL(/\/guides$/);
  });

  test('a failed sign-out says so under the row and the row works again', async ({ page }) => {
    await signedInPage(page);
    await page.route('**/v1/sessions', (route) =>
      route.fulfill({
        status: 500,
        contentType: 'application/json',
        body: '{"ok":false,"data":null,"error":"boom","request_id":"r"}',
      }),
    );
    await control(page).click();
    await rows(page).getByTestId('selector-row-sign-out').click();
    await expect(panel(page).getByTestId('selector-sign-out-failed')).toHaveText(
      'Could not sign out. Try again.',
    );
    await expect(rows(page).getByTestId('selector-row-sign-out')).toBeEnabled();
  });

  test('has no Discord link in the bar', async ({ page }) => {
    await signedInPage(page);
    await expect(page.getByRole('banner').getByRole('link', { name: 'Discord' })).toHaveCount(0);
  });
});

test.describe('the account panel, signed out', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  test('Sign in with Battle.net is the one filled button, over an email link, paste and the addon', async ({
    page,
  }) => {
    await signedOut(page);
    await page.goto('/guides');
    await control(page).click();
    await expect(panel(page).getByTestId('selector-sign-in-bnet')).toBeVisible();
    await expect(panel(page).getByTestId('selector-sign-in-email')).toBeVisible();
    await expect(panel(page).getByTestId('selector-paste')).toBeVisible();
    await expect(panel(page).getByTestId('selector-addon')).toBeVisible();
    for (const id of ['selector-sign-in-bnet', 'selector-sign-in-email']) {
      expect((await panel(page).getByTestId(id).boundingBox())!.height).toBeGreaterThanOrEqual(MIN_TARGET);
    }
  });
});

test.describe('the account sheet on a phone', () => {
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');
  test.use({ viewport: { width: 360, height: 800 } });

  test('fits 360 px with every row a 44 px target, no sideways scroll', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS], [], PREMIUM_ME);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await expect(control(page)).toBeVisible();
    await control(page).click();
    const sheet = (await panel(page).boundingBox())!;
    expect(sheet.x).toBeGreaterThanOrEqual(0);
    expect(sheet.x + sheet.width).toBeLessThanOrEqual(360);
    const targets = panel(page).locator('.csel-arow, .csel-row-main, .csel-action, .csel-close');
    for (const target of await targets.all()) {
      const box = (await target.boundingBox())!;
      expect(box.height).toBeGreaterThanOrEqual(MIN_TARGET);
      expect(box.x + box.width).toBeLessThanOrEqual(360);
    }
    await expect(rows(page).locator('.csel-arow')).toHaveCount(4);
    const scroll = await page.evaluate(() => document.documentElement.scrollWidth <= window.innerWidth);
    expect(scroll).toBe(true);
  });

  test('signed out, the bar button is the gold person glyph named Sign in, and the sheet leads with the buttons', async ({
    page,
  }) => {
    await signedOut(page);
    await page.goto('/guides');
    await expect(control(page)).toHaveAttribute('aria-label', /Sign in/);
    await expect(control(page).locator('.csel-text')).toBeHidden();
    await control(page).click();
    await expect(panel(page).getByTestId('selector-sign-in-bnet')).toBeVisible();
  });
});
