// web/tests/e2e/character-history.spec.ts
// The History section of the character page (experience spec 2026-10-24-character-history):
// the premium owner's two charts and their lists, the free owner's sentence, the visitor's bare
// card, and the layout-shift check for the hinted owner. /v1/me, the character and the history
// route are stubbed with page.route; the strings asserted are the spec's section 6.
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { characterHistoryFixture } from '../../src/fixtures/character-history';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  ownerMeAnswer,
  routeCharacterHistory,
} from './support/character-fixture';
import { readLayoutShift, trackLayoutShift } from './support/cls';

const MAX_CLS = 0.02;
const PHONE = { width: 390, height: 844 };
const DESKTOP = { width: 1440, height: 1000 };
const HISTORY_URL = `${CHARACTER_URL}/history`;

const SCORE = {
  score: 1830,
  tier: 'epic',
  as_of: '2026-12-11',
  components: { gear: 412, experience: 620, parses: 540, attendance: 180, discipline: 78 },
  experience: { dungeons: 14, raid_kills: 0, alts: 2 },
};

const NOT_FOUND = {
  status: 404,
  contentType: 'application/json',
  body: JSON.stringify({ ok: false, data: null, error: 'not_found', request_id: 'r' }),
};

async function stubCharacter(page: Page, found = true): Promise<void> {
  await page.route(CHARACTER_URL, (route) =>
    found
      ? route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ ...CHARACTER, data: { ...CHARACTER.data, ...SCORE } }),
        })
      : route.fulfill(NOT_FOUND),
  );
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
}

/** Signs the viewer in as the character's owner, with or without the history entitlement. */
async function stubOwner(page: Page, history: boolean): Promise<void> {
  await page
    .context()
    .addCookies([
      { name: 'fs_csrf', value: 'token', url: `http://localhost:${process.env.E2E_PORT ?? '4321'}` },
    ]);
  const me = ownerMeAnswer();
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ...me,
        data: { ...(me.data as Record<string, unknown>), entitlements: mockEntitlements({ history }) },
      }),
    }),
  );
}

/** Every request to the history route, recorded. */
function watchHistoryRequests(page: Page): string[] {
  const seen: string[] = [];
  page.on('request', (request) => {
    if (new URL(request.url()).pathname.endsWith('/history')) seen.push(request.url());
  });
  return seen;
}

async function openPremiumOwner(page: Page, history = characterHistoryFixture()): Promise<void> {
  await page.setViewportSize(DESKTOP);
  await stubOwner(page, true);
  await stubCharacter(page);
  await routeCharacterHistory(page, history);
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('history-score')).toBeVisible();
}

const scoreColumn = (page: Page): Locator => page.getByTestId('history-score');
const parseColumn = (page: Page): Locator => page.getByTestId('history-parse');
const readoutValue = (column: Locator): Locator => column.getByTestId('history-readout-value');
const readoutDetail = (column: Locator): Locator => column.getByTestId('history-readout-detail');

test('the premium owner sees both charts, the score readout and the parse boss pills', async ({ page }) => {
  await openPremiumOwner(page);

  await expect(page.getByTestId('character-history-section')).toBeVisible();
  await expect(page.getByTestId('history-caption')).toHaveText(
    'Includes your private and guild reports. Only you see this section.',
  );
  await expect(scoreColumn(page).getByTestId('history-svg')).toBeVisible();
  await expect(parseColumn(page).getByTestId('history-svg')).toBeVisible();

  await expect(readoutValue(scoreColumn(page))).toHaveText('1830');
  await expect(scoreColumn(page).getByTestId('history-readout-word')).toHaveText('Epic');
  await expect(readoutDetail(scoreColumn(page))).toHaveText('11 Dec 2026 · 2 fights');
  await expect(page.getByTestId('history-score-change')).toHaveText('Up 1090 since 12 Oct');

  const pills = page.getByTestId('history-pill');
  await expect(pills).toHaveCount(2);
  await expect(pills.first()).toHaveAttribute('aria-selected', 'true');
  await expect(pills.first()).toHaveAccessibleName(/Golemagg.*5 kills/);
  await expect(pills.nth(1)).toHaveAccessibleName(/Gehennas.*1 kill$/);
  await expect(readoutValue(parseColumn(page))).toHaveText('88');
  await expect(readoutDetail(parseColumn(page))).toHaveText('11 Dec 2026 · kill 5 of 5');
});

test('the step buttons and the arrow keys move the selection and the marked rows', async ({ page }) => {
  await openPremiumOwner(page);
  const score = scoreColumn(page);
  const marked = score.locator('[data-marked="true"]');
  const later = score.getByTestId('history-step-later');
  const earlier = score.getByTestId('history-step-earlier');

  await expect(marked).toHaveCount(2);
  await expect(later).toHaveAttribute('aria-disabled', 'true');

  await earlier.click();
  await expect(readoutValue(score)).toHaveText('1700');
  await expect(readoutDetail(score)).toHaveText('8 Dec 2026 · 1 fight');
  await expect(marked).toHaveCount(1);

  await later.click();
  await expect(readoutValue(score)).toHaveText('1830');

  await score.getByTestId('history-plot').focus();
  await page.keyboard.press('ArrowLeft');
  await expect(readoutValue(score)).toHaveText('1700');
  await page.keyboard.press('Home');
  await expect(readoutValue(score)).toHaveText('740');
  await expect(earlier).toHaveAttribute('aria-disabled', 'true');
  await page.keyboard.press('End');
  await expect(readoutValue(score)).toHaveText('1830');

  const parse = parseColumn(page);
  await parse.getByTestId('history-plot').focus();
  await page.keyboard.press('ArrowLeft');
  await expect(readoutValue(parse)).toHaveText('82');
  await expect(readoutDetail(parse)).toHaveText('8 Dec 2026 · kill 4 of 5');
  await expect(parse.locator('[data-marked="true"]')).toHaveCount(1);
});

test('each of the three gap reasons is named on the chart', async ({ page }) => {
  await openPremiumOwner(
    page,
    characterHistoryFixture({
      gaps: [
        { from: '2026-10-26', to: '2026-11-17', reason: 'No rated fights' },
        { from: '2026-11-24', to: '2026-12-01', reason: 'Report flagged' },
        { from: '2026-12-01', to: '2026-12-08', reason: 'Score not kept' },
      ],
    }),
  );
  await expect(page.getByTestId('history-gap')).toHaveCount(3);
  // A gap too narrow for its band names its reason under the chart instead (spec 3.6).
  const named = [
    ...(await page.getByTestId('history-gap-label').allTextContents()),
    ...(await page.getByTestId('history-gap-note').allTextContents()),
  ]
    .join(' ')
    .replace(/\s+/g, ' ');
  for (const reason of ['No rated fights', 'Report flagged', 'Score not kept']) {
    expect(named).toContain(reason);
  }
});

test('a boss with one kill is a point and its line; the other boss returns the chart', async ({ page }) => {
  await openPremiumOwner(page);
  const parse = parseColumn(page);

  await page.getByTestId('history-pill').nth(1).click();
  await expect(page.getByTestId('history-pill').nth(1)).toHaveAttribute('aria-selected', 'true');
  await expect(page.getByTestId('history-parse-one-kill')).toHaveText('One kill so far. A line needs two.');
  await expect(parse.getByTestId('history-svg')).toHaveCount(0);
  await expect(
    parse.getByTestId('history-parse-fights-rows').locator('li:not([data-testid="character-table-header"])'),
  ).toHaveCount(1);

  await page.getByTestId('history-pill').first().click();
  await expect(parse.getByTestId('history-svg')).toBeVisible();
  await expect(readoutValue(parse)).toHaveText('88');
});

test('Show all expands the list in place, Show fewer collapses it, and a private fight says so', async ({
  page,
}) => {
  await openPremiumOwner(page);
  const rows = page
    .getByTestId('history-score-fights-rows')
    .locator('li:not([data-testid="character-table-header"])');
  const toggle = page.getByTestId('history-score-fights-toggle');

  await expect(rows).toHaveCount(8);
  await expect(toggle).toHaveText('Show all 9 fights');
  await toggle.click();
  await expect(rows).toHaveCount(9);
  await expect(toggle).toHaveText('Show fewer');
  await toggle.click();
  await expect(rows).toHaveCount(8);
  await expect(toggle).toHaveText('Show all 9 fights');

  const notes = await page
    .getByTestId('history-score-fights-rows')
    .getByTestId('character-fight-note')
    .allTextContents();
  expect(notes).toContain('Private');
  expect(notes).toContain('Guild');
});

test('the free owner gets the card and one sentence, and the history is never requested', async ({
  page,
}) => {
  await page.setViewportSize(DESKTOP);
  const requests = watchHistoryRequests(page);
  await stubOwner(page, false);
  await stubCharacter(page);
  await routeCharacterHistory(page);
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('forever-score')).toBeVisible();
  await expect(page.getByTestId('history-free-sentence')).toHaveText(
    'History charts draw your score and parses over time. See plans.',
  );
  await expect(page.getByTestId('history-see-plans')).toHaveAttribute('href', '/premium#history');
  await expect(page.getByTestId('character-history-section')).toHaveCount(0);
  await page.waitForTimeout(500);
  expect(requests).toEqual([]);
});

test('a visitor gets the card alone: no section, no sentence, no history request', async ({ page }) => {
  await page.setViewportSize(DESKTOP);
  const requests = watchHistoryRequests(page);
  await stubCharacter(page);
  await routeCharacterHistory(page);
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('forever-score')).toBeVisible();
  await expect(page.getByTestId('character-history-section')).toHaveCount(0);
  await expect(page.getByTestId('history-free-sentence')).toHaveCount(0);
  await page.waitForTimeout(500);
  expect(requests).toEqual([]);
});

test('an anonymised character is the 404 page with no history anywhere', async ({ page }) => {
  await page.setViewportSize(DESKTOP);
  const requests = watchHistoryRequests(page);
  await stubCharacter(page, false);
  await routeCharacterHistory(page);
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-not-found-line')).toBeVisible();
  await expect(page.getByTestId('character-history-section')).toHaveCount(0);
  await expect(page.getByTestId('history-free-sentence')).toHaveCount(0);
  await page.waitForTimeout(500);
  expect(requests).toEqual([]);
});

test('an anonymised character is the 404 page for its entitled owner too', async ({ page }) => {
  await page.setViewportSize(DESKTOP);
  const requests = watchHistoryRequests(page);
  await stubOwner(page, true);
  await stubCharacter(page, false);
  await routeCharacterHistory(page);
  await page.goto(CHARACTER_PATH);

  // The owner's /v1/me lists the character, so the 404 is their own unlogged card.
  await expect(page.getByTestId('character-unlogged')).toBeVisible();
  await expect(page.getByTestId('character-history-section')).toHaveCount(0);
  await expect(page.getByTestId('history-free-sentence')).toHaveCount(0);
  await page.waitForTimeout(500);
  expect(requests).toEqual([]);
});

test('#history lands on the section for the entitled owner', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 600 });
  await stubOwner(page, true);
  await stubCharacter(page);
  await routeCharacterHistory(page);
  await page.goto(`${CHARACTER_PATH}#history`);

  await expect(page.getByTestId('history-score')).toBeVisible();
  await expect(page.getByTestId('character-history-section')).toBeInViewport();
  await expect.poll(() => page.evaluate(() => window.scrollY)).toBeGreaterThan(0);
});

test('a failed read shows the error with Try again, which reads again', async ({ page }) => {
  await page.setViewportSize(DESKTOP);
  await stubOwner(page, true);
  await stubCharacter(page);
  let calls = 0;
  await page.route(HISTORY_URL, (route) => {
    calls += 1;
    return calls === 1
      ? route.fulfill({
          status: 500,
          contentType: 'application/json',
          body: JSON.stringify({ ok: false, data: null, error: { message: 'history failed' } }),
        })
      : route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ ok: true, data: characterHistoryFixture(), error: null }),
        });
  });
  await page.goto(CHARACTER_PATH);

  const error = page.getByTestId('history-error');
  await expect(error).toContainText('Your history did not load.');
  await expect(page.getByTestId('forever-score')).toBeVisible();
  await error.getByRole('button', { name: 'Try again' }).click();
  await expect(scoreColumn(page)).toBeVisible();
  await expect(page.getByTestId('history-error')).toHaveCount(0);
  expect(calls).toBe(2);
});

for (const viewport of [PHONE, DESKTOP]) {
  test(`the hinted owner's section loads without layout shift at ${viewport.width}px`, async ({ page }) => {
    await page.setViewportSize(viewport);
    await stubOwner(page, true);
    await stubCharacter(page);
    await routeCharacterHistory(page);
    // The first visit stores the /v1/me snapshot and the section's height; the second is hinted.
    await page.goto(CHARACTER_PATH);
    await expect(scoreColumn(page)).toBeVisible();
    await expect(parseColumn(page).getByTestId('history-svg')).toBeVisible();
    await page.waitForTimeout(500);

    await trackLayoutShift(page);
    await page.reload();
    await expect(scoreColumn(page)).toBeVisible();
    await expect(parseColumn(page).getByTestId('history-svg')).toBeVisible();
    const cls = await readLayoutShift(page);
    test.info().annotations.push({ type: 'cls', description: `${viewport.width}px: ${cls}` });
    expect(cls).toBeLessThanOrEqual(MAX_CLS);
  });
}
