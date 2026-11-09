// web/tests/e2e/rankings-encounter-picker.spec.ts
// Signed out, a bare /rankings is the encounter list (design/specs/2026-11-07-logs-reports.md
// section 7, decision K4): "Choose an encounter", one row per encounter with ranked kills (GET
// /v1/rankings/encounters, most kills first), its zone, raid size and kill count, each a link to its
// board. Signed in it opens the encounter the character last killed (rankings.spec.ts). When nothing
// is ranked at all the list says so, says what fills it, and offers Upload a log.
import { expect, test } from './support/test';
import { stubRankedEncounters } from './support/rankings-fixture';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

const KELTHAS = { id: 9001, name: 'Warden Kelthas', slug: 'warden-kelthas', zone: null, size: 5, kills: 3 };
const SKOLEX = {
  id: 9002,
  name: 'Skolex the Insatiable',
  slug: 'skolex-the-insatiable',
  zone: 'Excavation Site: Wetlands',
  size: 5,
  kills: 1,
};

test('signed out, bare /rankings lists the ranked encounters, each a link to its board', async ({ page }) => {
  await stubRankedEncounters(page, [KELTHAS, SKOLEX]);
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
  await expect(page.locator('main header')).toContainText('2 encounters ranked');
  await expect(page.getByTestId('encounter-picker').locator('h2')).toHaveText('Choose an encounter');
  const links = list.getByRole('link');
  await expect(links).toHaveCount(2);
  await expect(links.nth(0)).toContainText('Warden Kelthas');
  await expect(links.nth(0)).toContainText('5-player');
  await expect(links.nth(0)).toContainText('3 ranked');
  await expect(links.nth(0)).toHaveAttribute('href', '/rankings/warden-kelthas');
  await expect(links.nth(1)).toContainText('Excavation Site: Wetlands · 5-player');
  await expect(links.nth(1)).toHaveAttribute('href', '/rankings/skolex-the-insatiable');
  // The list is the page: the bar holds the board tabs alone, no board was asked for.
  await expect(page.getByTestId('control-encounter')).toHaveCount(0);
  await expect(page.getByTestId('ranking-board')).toHaveCount(0);
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
  // The empty board points at the bosses that do have kills.
  await expect(page.getByTestId('rankings-elsewhere').getByRole('link')).toHaveText(/Skolex the Insatiable/);
});

test('with nothing ranked yet, the list says so, says what fills it, and offers Upload a log', async ({
  page,
}) => {
  await stubRankedEncounters(page, []);

  await page.goto('/rankings');

  const none = page.getByTestId('rankings-no-encounters');
  await expect(none).toContainText('No boss kills are ranked yet.');
  await expect(none).toContainText('A boss gets its board when the first log of its kill is uploaded.');
  await expect(none.getByRole('link', { name: 'Upload a log' })).toHaveAttribute('href', '/logs');
  await expect(page.getByTestId('rankings-encounters')).toHaveCount(0);
});

test('the Guilds Progress board needs no encounter and offers every encounter at once', async ({
  page,
  isMobile,
}) => {
  test.skip(isMobile, 'the board tabs sit on the desktop bar line; the phone opens them from its summary');
  await stubRankedEncounters(page, [KELTHAS]);
  await page.route('**/v1/rankings/guilds?**', (route) =>
    route.fulfill(
      envelope({
        rows: [
          {
            rank: 1,
            guild: { name: 'The Last Watch', ruleset: 'hardcore', region: 'us' },
            value: 2,
            fought_at: '2026-12-09T00:00:00Z',
          },
        ],
      }),
    ),
  );

  await page.goto('/rankings');
  await expect(page.getByTestId('rankings-encounters')).toBeVisible();
  await page.getByTestId('control-board-guild').click();

  await expect(page).toHaveURL(/board=guild/);
  await expect(page.getByTestId('guild-rows')).toContainText('The Last Watch');
  await expect(page.getByTestId('guild-rows')).toContainText('2 bosses');
  await expect(page.getByTestId('control-encounter')).toHaveValue('');
  await expect(page.getByTestId('control-encounter').locator('option').first()).toHaveText('Every encounter');
});
