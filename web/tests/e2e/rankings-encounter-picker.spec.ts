// web/tests/e2e/rankings-encounter-picker.spec.ts
// A bare /rankings opens on an encounter (the character's most recent kill, else the first the API
// lists) with the rest one filter away; the picker is for when nothing is ranked at all, where it
// says so plainly and offers the Guilds board instead.
import { expect, test } from './support/test';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

test('bare /rankings opens on the first listed encounter and offers the others as a filter', async ({
  page,
}) => {
  await page.route('**/v1/encounters', (route) =>
    route.fulfill(
      envelope({
        rows: [
          { id: 9001, name: 'Warden Kelthas', slug: 'warden-kelthas' },
          { id: 9002, name: 'Skolex the Insatiable', slug: 'skolex-the-insatiable' },
        ],
      }),
    ),
  );
  const asked: string[] = [];
  await page.route('**/v1/rankings?**', (route) => {
    asked.push(new URL(route.request().url()).searchParams.get('encounter') ?? '');
    return route.fulfill(
      envelope({ rows: [], total: 0, page: 1, per_page: 100, updated_at: '2026-12-09T00:00:00Z' }),
    );
  });

  await page.goto('/rankings');

  await expect(page.getByTestId('rankings-empty')).toBeVisible();
  await expect(page).toHaveURL(/\/rankings\/warden-kelthas$/);
  await expect(page.getByTestId('encounter-picker')).toHaveCount(0);
  await expect(page.getByTestId('filter-encounter').locator('option')).toHaveText([
    'Warden Kelthas',
    'Skolex the Insatiable',
  ]);
  expect(asked).toEqual(['warden-kelthas']);
});

test('with no encounter data yet, the Characters board says so and offers the Guilds board', async ({
  page,
}) => {
  await page.route('**/v1/encounters', (route) => route.fulfill(envelope({ rows: [] })));
  await page.route('**/v1/rankings/guilds?**', (route) =>
    route.fulfill(
      envelope({
        rows: [
          {
            rank: 1,
            guild: { name: 'The Last Watch', ruleset: 'hardcore', region: 'us' },
            value: 12,
            fought_at: '2026-12-09T00:00:00Z',
            report_id: 'r1',
          },
        ],
      }),
    ),
  );

  await page.goto('/rankings');

  await expect(page.getByTestId('rankings-no-encounters')).toHaveText('No encounters have been ranked yet.');
  await expect(page.getByTestId('filter-encounter')).toBeDisabled();
  await expect(page.getByTestId('filter-encounter').locator('option')).toHaveText(['No encounters yet']);
  const guildsButton = page.getByTestId('rankings-try-guilds');
  await expect(guildsButton).toContainText('Open the Guilds board');

  await guildsButton.click();
  await expect(page.getByTestId('board-guild')).toHaveAttribute('aria-selected', 'true');
  await expect(page.getByTestId('guild-rows')).toContainText('The Last Watch');
});

test('the Guilds board also picks an encounter for speed and execution, not for progress', async ({
  page,
}) => {
  await page.route('**/v1/encounters', (route) => route.fulfill(envelope({ rows: [] })));

  await page.goto('/rankings?board=guild&kind=speed');
  await expect(page.getByTestId('rankings-no-encounters')).toBeVisible();
  const progressButton = page.getByTestId('rankings-try-guilds');
  await expect(progressButton).toContainText('Open the Progress board');
});
