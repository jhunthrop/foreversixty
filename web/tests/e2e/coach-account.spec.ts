// web/tests/e2e/coach-account.spec.ts
// The account page's "last raid" coach line (design/specs/2026-10-28-log-coach.md section 3.3):
// its five states plus the error, the shared reserve (no shift), and the privacy page's Log coach
// section, shown once the plans page stops holding the coach line and absent while it holds it.
import { expect, test, type Page } from '@playwright/test';
import { coachCopy } from '../../src/lib/coach/copy';
import { coachPrivacySection } from '../../src/lib/coach/privacy-section';
import { lastRaidLine } from '../../src/lib/coach/last-raid';
import { normalizeRecent } from '../../src/lib/coach/normalize';
import {
  RECENT,
  routeAccountDevices,
  routeRecent,
  signInAs,
  type AccountLineId,
} from './support/coach-stubs';
import { gotoHydrated } from './support/hydrated';
import { drainShifts, recordShifts, sumShifts } from './support/layout-shift';
import { expectNoAxeViolations } from './support/officer-tools';
import { shortControls } from './support/short-controls';

const ZONE = 'America/Chicago';
const MAX_ACCOUNT_CLS = 0.02;
const SEPARATOR = ' · ';
test.use({ timezoneId: ZONE });

const line = (page: Page) => page.getByTestId('coach-last-raid');

async function openAccount(page: Page, id: AccountLineId): Promise<void> {
  const { viewer, recent } = RECENT[id];
  await signInAs(page, viewer);
  await routeAccountDevices(page);
  await routeRecent(page, recent);
  await recordShifts(page);
  await gotoHydrated(page, '/account', 'coach-last-raid');
}

/** The sentence the line should read, from the same module the page uses. */
function expectedSentence(id: AccountLineId): string {
  const { viewer, recent } = RECENT[id];
  if (typeof recent === 'number') return coachCopy.lastRaidError;
  const view = lastRaidLine(normalizeRecent(recent), viewer === 'premium', ZONE);
  return view.segments.map((segment) => segment.text).join(SEPARATOR);
}

const readable: AccountLineId[] = ['premium', 'premium-one', 'free', 'pending', 'clean', 'none'];

for (const id of readable) {
  test(`the ${id} line reads its sentence`, async ({ page }) => {
    await openAccount(page, id);
    await expect(line(page)).toHaveAttribute('data-state', 'ready');
    await expect(line(page)).toHaveText(expectedSentence(id).replace(/\s+/g, ' '));
    expect((await line(page).textContent()) ?? '').not.toMatch(/!|heal more/i);
  });
}

test('the Premium line opens the review; the free line goes to /premium', async ({ page }) => {
  await openAccount(page, 'premium');
  await expect(page.getByTestId('coach-last-raid-link')).toHaveAttribute(
    'href',
    /\/reports\/fixture2abcd\?fight=3&tab=coach/,
  );
  const free = await page.context().newPage();
  await openAccount(free, 'free');
  await expect(free.getByTestId('coach-last-raid-link')).toHaveAttribute('href', '/premium#coach');
});

test('no raid in 14 days shows the muted none line, never hidden', async ({ page }) => {
  await openAccount(page, 'none');
  await expect(line(page)).toHaveText(coachCopy.lastRaidNone);
  await expect(page.getByTestId('coach-last-raid-link')).toHaveCount(0);
});

test('a failed read shows the error line and a retry that reads again', async ({ page }) => {
  await openAccount(page, 'error');
  await expect(line(page)).toContainText(coachCopy.lastRaidError);
  await page.unroute('**/v1/me/coach/recent');
  await routeRecent(page, RECENT.premium.recent);
  await page.getByTestId('coach-last-raid-retry').click();
  await expect(line(page)).toHaveAttribute('data-state', 'ready');
});

test('the line is the first child of the Characters panel and shifts nothing', async ({ page }) => {
  await openAccount(page, 'premium');
  await expect(line(page)).toHaveAttribute('data-state', 'ready');
  const first = page.getByTestId('account-characters').locator('[data-testid="coach-last-raid"]');
  await expect(first).toBeVisible();
  expect(sumShifts(await drainShifts(page))).toBeLessThanOrEqual(MAX_ACCOUNT_CLS);
});

test('axe finds nothing, and no control is under 44 px at 360', async ({ page }) => {
  await page.setViewportSize({ width: 360, height: 800 });
  await openAccount(page, 'error');
  await expect(line(page)).toContainText(coachCopy.lastRaidError);
  await expectNoAxeViolations(page, '[data-testid="coach-last-raid"]');
  expect(await shortControls(line(page))).toEqual([]);
});

test('the privacy page carries the Log coach section and its links', async ({ page }) => {
  const section = coachPrivacySection();
  await page.goto('/privacy');
  await expect(page.getByRole('heading', { name: 'Log coach' })).toBeVisible();
  for (const paragraph of section.paragraphs) await expect(page.getByText(paragraph)).toBeVisible();
  for (const link of section.links)
    await expect(page.getByRole('link', { name: link.label })).toHaveAttribute('href', link.href);
});
