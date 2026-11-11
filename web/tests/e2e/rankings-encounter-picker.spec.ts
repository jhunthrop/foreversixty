// web/tests/e2e/rankings-encounter-picker.spec.ts
// Signed out, a bare /rankings is the encounter list (design/specs/2026-11-07-logs-reports.md
// section 7, decision K4): "Encounters", one row per encounter with ranked kills (GET
// /v1/rankings/encounters, most kills first), its zone, raid size and kill count, each a link to its
// board. Signed in it opens the encounter the character last killed (rankings.spec.ts). When nothing
// is ranked at all the list says so, says what fills it, and offers Upload a log.
import { expect, test } from './support/test';
import { stubRankedEncounters, stubRankingsApi } from './support/rankings-fixture';
import { signedOut } from './support/selector';

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
  await expect(page.getByTestId('encounter-picker').locator('h2')).toHaveText('Encounters');
  const links = list.getByTestId('rankings-encounter');
  await expect(links).toHaveCount(2);
  await expect(links.nth(0)).toContainText('Warden Kelthas');
  await expect(links.nth(0)).toContainText('5-player');
  await expect(links.nth(0)).toContainText('3 kills');
  await expect(links.nth(0)).toHaveAttribute('href', '/rankings/warden-kelthas');
  await expect(links.nth(1)).toContainText('Excavation Site: Wetlands · 5-player');
  await expect(links.nth(1)).toHaveAttribute('href', '/rankings/skolex-the-insatiable');
  // The bar holds the board tabs alone. Each encounter's board is read for its top parses (every
  // class, each character's best); an empty one says so in one line.
  await expect(page.getByTestId('control-encounter')).toHaveCount(0);
  await expect(page.getByTestId('ranking-board')).toHaveCount(0);
  await expect(page.getByTestId('rankings-top-warden-kelthas')).toHaveText(
    'No ranked kill on this board yet.',
  );
  expect(
    asked.some((url) => url.includes('encounter=warden-kelthas') && url.includes('distinct=character')),
  ).toBe(true);

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
  await page.route('**/v1/rankings?**', (route) =>
    route.fulfill(
      envelope({ rows: [], total: 0, page: 1, per_page: 100, updated_at: '2026-12-09T00:00:00Z' }),
    ),
  );
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

test("the landing shows each encounter's top three and the spec leaders on the raid boss", async ({
  page,
}) => {
  await signedOut(page);
  await stubRankingsApi(page);

  await page.goto('/rankings');

  const top = page.getByTestId('rankings-top-onyxia');
  await expect(top.locator('li')).toHaveCount(3);
  await expect(top.locator('li').first()).toContainText('Vashrel');
  await expect(top.locator('li').first()).toContainText('Combat Rogue');
  // Each boss reads its own board: Warden Kelthas' top three are not Onyxia's.
  await expect(page.getByTestId('rankings-top-warden-kelthas').locator('li').first()).toContainText(
    'Reglitch',
  );
  await expect(top.getByTestId('ranking-value').first()).toHaveAttribute('href', '/reports/report1?fight=1');
  await expect(page.locator('main header')).toContainText('Updated 13 Dec');
  const leaders = page.getByTestId('rankings-leaders');
  await expect(leaders.locator('h2')).toHaveText('Leaders by spec');
  await expect(leaders).toContainText('Onyxia');
  // Damage by DPS, healing by HPS, each spec once, numbered within its list, the spec leading the row.
  const damage = page.getByTestId('rankings-leaders-damage');
  await expect(damage.locator('li')).toHaveCount(8);
  await expect(damage.getByTestId('rank-rank')).toHaveText(['1', '2', '3', '4', '5', '6', '7', '8']);
  await expect(damage.getByTestId('rank-name').first()).toHaveText('Combat Rogue');
  await expect(damage.getByTestId('rank-name').first()).toHaveAttribute(
    'href',
    /\/rankings\/onyxia\?class=rogue&spec=Combat/,
  );
  const healing = page.getByTestId('rankings-leaders-healing');
  await expect(healing.getByTestId('rank-name')).toHaveText([
    'Holy Priest',
    'Restoration Shaman',
    'Holy Paladin',
    'Restoration Druid',
  ]);
});
