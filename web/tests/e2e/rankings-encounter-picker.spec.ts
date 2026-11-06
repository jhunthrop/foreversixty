// web/tests/e2e/rankings-encounter-picker.spec.ts
// Signed out, a bare /rankings is the encounter list (design/specs/2026-11-07-logs-reports.md
// section 7, decision K4): "Choose an encounter", one row per ranked encounter, each a link to its
// board. Signed in it opens the encounter the character last killed (rankings.spec.ts). When
// nothing is ranked at all the list says so plainly and offers the Guilds board instead.
import { expect, test } from './support/test';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

test('signed out, bare /rankings lists the ranked encounters, each a link to its board', async ({ page }) => {
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
    asked.push(route.request().url());
    return route.fulfill(
      envelope({ rows: [], total: 0, page: 1, per_page: 100, updated_at: '2026-12-09T00:00:00Z' }),
    );
  });

  await page.goto('/rankings');

  const list = page.getByTestId('rankings-encounters');
  await expect(list).toBeVisible();
  await expect(page).toHaveURL(/\/rankings$/);
  await expect(page.getByTestId('page-title')).toHaveText('Rankings');
  await expect(page.getByTestId('encounter-picker').locator('h2')).toHaveText('Choose an encounter');
  await expect(page.getByTestId('encounter-picker')).toContainText('2 encounters ranked');
  const links = list.getByRole('link');
  await expect(links).toHaveCount(2);
  await expect(links.nth(0)).toContainText('Warden Kelthas');
  await expect(links.nth(0)).toHaveAttribute('href', '/rankings/warden-kelthas');
  await expect(links.nth(1)).toContainText('Skolex the Insatiable');
  await expect(links.nth(1)).toHaveAttribute('href', '/rankings/skolex-the-insatiable');
  // The list is the page: no bar, no board, and no board was asked for.
  await expect(page.getByTestId('rankings-filters')).toHaveCount(0);
  await expect(page.getByTestId('ranking-board')).toHaveCount(0);
  await expect(page.getByTestId('rankings-legend')).toBeVisible();
  expect(asked).toEqual([]);

  // The fixture build prerenders the one board Playwright can open without a Worker.
  await links.nth(0).click();
  await expect(page).toHaveURL(/\/rankings\/warden-kelthas$/);
  await expect(page.getByTestId('rankings-empty')).toBeVisible();
  await expect(page.getByTestId('control-encounter')).toHaveValue('warden-kelthas');
  await expect(page.getByTestId('control-encounter').locator('option')).toHaveText([
    'Warden Kelthas',
    'Skolex the Insatiable',
  ]);
});

test('with no encounter data yet, the list says so and offers the Guilds board', async ({ page }) => {
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

  const none = page.getByTestId('rankings-no-encounters');
  await expect(none).toContainText('No encounters have been ranked yet.');
  await expect(none).toContainText(
    'The Guilds board ranks guild progress across every encounter and needs none.',
  );
  await expect(page.getByTestId('rankings-encounters')).toHaveCount(0);
  const guildsButton = page.getByTestId('rankings-try-guilds');
  await expect(guildsButton).toContainText('Open the Guilds board');

  await guildsButton.click();
  await expect(page.getByTestId('control-board-guild')).toHaveAttribute('aria-checked', 'true');
  await expect(page.getByTestId('guild-rows')).toContainText('The Last Watch');
  await expect(page).toHaveURL(/board=guild/);
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
