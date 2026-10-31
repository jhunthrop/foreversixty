// web/tests/e2e/coach-history.spec.ts
// The character page's coach history (design/specs/2026-10-28-log-coach.md sections 3.5, 4.3 and
// 10.3): twenty 56 px rows from the history route, each a link to that pull's Coach tab, then
// `Show older` (aria-busy while it reads), the error text, the empty and free states, and the
// Coach tab's `#coach` landing. Every row is the coach's own output (src/fixtures/coach/history-*).
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { coachCopy } from '../../src/lib/coach/copy';
import { historyRow } from '../../src/lib/coach/history-list';
import { normalizeHistory } from '../../src/lib/coach/normalize';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import { CHARACTER_PATH, CHARACTER_URL, ownerMeAnswer, stubCharacter } from './support/character-fixture';
import { HISTORY_PAGES, routeCoachHistory, type HistoryStub } from './support/coach-stubs';
import { envelope } from './support/guild-billing';
import { expectNoAxeViolations } from './support/officer-tools';
import { plantSessionHint } from './support/subscribe-stubs';
import { shortControls } from './support/short-controls';

const ROW_PX = 56;
const OLDER_PX = 44;
const SEPARATOR = ' · ';
const list = (page: Page) => page.getByTestId('coach-history-list');
const rows = (page: Page) => page.getByTestId('coach-history-row');
const FIRST = normalizeHistory(HISTORY_PAGES.first);
const OLDER = normalizeHistory(HISTORY_PAGES.older);

/** Signs the character's owner in, Premium (`log_coach`) or free, and serves the page. */
async function openOwner(page: Page, premium: boolean, history: HistoryStub, hash = ''): Promise<string[]> {
  await plantSessionHint(page);
  const me = ownerMeAnswer();
  const data = {
    ...(me.data as Record<string, unknown>),
    entitlements: mockEntitlements({ log_coach: premium }),
  };
  await page.route('**/v1/me', (route) => route.fulfill(envelope(data)));
  await stubCharacter(page);
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(envelope({ code: 'not_found' }, 404)));
  const asked: string[] = [];
  page.on('request', (request) => {
    if (request.url().includes('/coach/history')) asked.push(request.url());
  });
  await routeCoachHistory(page, history);
  await page.goto(`${CHARACTER_PATH}${hash}`);
  await expect(page.getByTestId('character')).toBeVisible();
  return asked;
}

test('twenty rows, each the review line, a link to its Coach tab, and 56 px tall', async ({ page }) => {
  await openOwner(page, true, { first: HISTORY_PAGES.first });
  await expect(rows(page)).toHaveCount(20);
  for (const [index, item] of FIRST.reviews.entries()) {
    const row = historyRow(item);
    const link = rows(page).nth(index);
    await expect(link).toHaveAttribute('href', row.href);
    await expect(link.getByTestId('coach-history-title')).toHaveText(row.title);
    await expect(link).toContainText([...row.lead, row.title].join(SEPARATOR));
    await expect(link.getByTestId('coach-history-wipe')).toHaveCount(row.wipe ? 1 : 0);
    await expect(link.getByTestId('coach-history-rules')).toHaveCount(row.logOnly ? 1 : 0);
  }
  const box = await rows(page).first().boundingBox();
  expect(box?.height ?? 0).toBeGreaterThanOrEqual(ROW_PX);
  await expect(list(page).getByTestId('coach-history-wipe').first()).toHaveText(coachCopy.historyWipe);
  await expect(list(page).getByTestId('coach-history-rules').first()).toHaveText(coachCopy.historyRules);
  await expect(list(page).getByTestId('coach-history-rules').first()).toHaveAttribute(
    'title',
    coachCopy.historyRulesTitle,
  );
  await expect(list(page).getByText(coachCopy.historyNone).first()).toBeVisible();
  expect((await list(page).textContent()) ?? '').not.toMatch(/!|heal more/i);
});

test('Show older is aria-busy while it reads, then appends the last page and goes', async ({ page }) => {
  let release = (): void => {};
  const olderGate = new Promise<void>((resolve) => (release = resolve));
  const asked = await openOwner(page, true, { first: HISTORY_PAGES.first, olderGate });
  const older = page.getByTestId('coach-history-older');
  await expect(older).toHaveText(coachCopy.historyOlder);
  expect((await older.boundingBox())?.height ?? 0).toBeGreaterThanOrEqual(OLDER_PX);
  await older.click();
  await expect(older).toHaveAttribute('aria-busy', 'true');
  await expect(older).toBeDisabled();
  release();
  await expect(rows(page)).toHaveCount(FIRST.reviews.length + OLDER.reviews.length);
  await expect(older).toHaveCount(0);
  expect(asked.at(-1)).toContain(`cursor=${FIRST.next_cursor ?? ''}`);
});

test('a failed older page keeps the rows and says so under the list', async ({ page }) => {
  await openOwner(page, true, { first: HISTORY_PAGES.first, older: 500 });
  await page.getByTestId('coach-history-older').click();
  await expect(page.getByTestId('coach-history-older-error')).toHaveText(coachCopy.historyOlderError);
  await expect(rows(page)).toHaveCount(20);
  await expect(page.getByTestId('coach-history-older')).toBeEnabled();
});

test('a failed first page shows the error with a retry', async ({ page }) => {
  await openOwner(page, true, { first: 500 });
  await expect(page.getByTestId('coach-history-error')).toContainText(coachCopy.error);
  await page.unroute('**/v1/characters/*/*/*/coach/history*');
  await routeCoachHistory(page, { first: HISTORY_PAGES.first });
  await page.getByTestId('coach-history-error-retry').click();
  await expect(rows(page)).toHaveCount(20);
});

test('no history: the section is absent', async ({ page }) => {
  const asked = await openOwner(page, true, { first: { reviews: [], next_cursor: null } });
  await expect.poll(() => asked.length).toBe(1);
  await expect(list(page)).toHaveCount(0);
});

test('a free owner gets no list and no request', async ({ page }) => {
  const asked = await openOwner(page, false, { first: HISTORY_PAGES.first });
  await expect(page.getByTestId('character-name')).toBeVisible();
  await expect(list(page)).toHaveCount(0);
  expect(asked).toEqual([]);
});

test("the Coach tab's #coach link lands on the list", async ({ page }) => {
  await openOwner(page, true, { first: HISTORY_PAGES.first }, '#coach');
  await expect(rows(page)).toHaveCount(20);
  await expect(list(page)).toBeInViewport();
});

test('axe finds nothing, and no control is under 44 px at 360', async ({ page }) => {
  await page.setViewportSize({ width: 360, height: 800 });
  await openOwner(page, true, { first: HISTORY_PAGES.first });
  await expect(rows(page)).toHaveCount(20);
  await expectNoAxeViolations(page, '[data-testid="coach-history-list"]');
  expect(await shortControls(list(page))).toEqual([]);
});
