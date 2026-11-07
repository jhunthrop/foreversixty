// web/tests/e2e/coach-account.spec.ts
// The account page's "last raid" coach line (design/specs/2026-10-28-log-coach.md section 3.3; premium
// spec 2026-11-08 7.2 and G5): the Characters Plate's foot row in its four states with a raid, nothing
// at all with no raid or on a failed read, and the privacy page's Log coach section.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
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
import { expectNoAxeViolations } from './support/officer-tools';
import { shortControls } from './support/short-controls';

const ZONE = 'America/Chicago';
const SEPARATOR = ' · ';
test.use({ timezoneId: ZONE });

const line = (page: Page) => page.getByTestId('coach-last-raid');

async function openAccount(page: Page, id: AccountLineId): Promise<void> {
  const { viewer, recent } = RECENT[id];
  await signInAs(page, viewer);
  await routeAccountDevices(page);
  await routeRecent(page, recent);
  await gotoHydrated(page, '/account', 'account-characters');
}

/** The sentence the line should read, from the same module the page uses. */
function expectedSentence(id: AccountLineId): string {
  const { viewer, recent } = RECENT[id];
  if (typeof recent === 'number') return coachCopy.lastRaidError;
  const view = lastRaidLine(normalizeRecent(recent), viewer === 'premium', ZONE);
  return view.segments.map((segment) => segment.text).join(SEPARATOR);
}

const readable: AccountLineId[] = ['premium', 'premium-one', 'free', 'pending', 'clean'];

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

for (const id of ['none', 'error'] as const) {
  test(`${id === 'none' ? 'no raid in 14 days' : 'a failed read'} draws no line at all`, async ({ page }) => {
    const read = page.waitForResponse('**/v1/me/coach/recent');
    await openAccount(page, id);
    await read;
    await expect(page.getByTestId('account-character-row').first()).toBeVisible();
    await expect(line(page)).toHaveCount(0);
  });
}

test("the line is the Characters Plate's foot row, under the character rows", async ({ page }) => {
  await openAccount(page, 'premium');
  await expect(line(page)).toHaveAttribute('data-state', 'ready');
  const inPlate = page
    .getByTestId('account-characters')
    .locator('[data-region="plate"] [data-testid="coach-last-raid"]');
  await expect(inPlate).toBeVisible();
  const rowBottom = await page
    .getByTestId('account-character-row')
    .last()
    .evaluate((el) => el.getBoundingClientRect().bottom);
  const lineTop = await line(page).evaluate((el) => el.getBoundingClientRect().top);
  expect(lineTop).toBeGreaterThanOrEqual(rowBottom - 1);
});

test('axe finds nothing, and no control is under 44 px at 360', async ({ page }) => {
  await page.setViewportSize({ width: 360, height: 800 });
  await openAccount(page, 'premium');
  await expect(line(page)).toHaveAttribute('data-state', 'ready');
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
