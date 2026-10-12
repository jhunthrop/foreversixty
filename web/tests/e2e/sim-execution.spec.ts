import { expect, test, type Page } from '@playwright/test';
import { gotoHydrated } from './support/hydrated';
import { simCopy } from '../../src/lib/sim/copy';

/** The rankings page reads GET /v1/rankings; the suite stubs it, the way rankings.spec.ts does. */
function page1(rows: unknown[]) {
  return {
    ok: true,
    data: { rows, total: rows.length, page: 1, per_page: 100, updated_at: '2026-09-14T00:00:00Z' },
    error: null,
    request_id: 'r',
  };
}

const base = {
  rank: 1,
  player: { key: 'us/forever/thrallgar', name: 'Thrallgar-Whitemane', class: 'Warrior', spec: 'Fury' },
  value: 1180,
  size: 40,
  fought_at: '2026-09-13T22:00:00Z',
  duration_ms: 181000,
  talent_split: '0/31/20',
  trinkets: [],
  buff_count: 12,
  report_id: 'fixture2abcd',
  fight_index: 2,
  state: 'ok',
};

/**
 * L8.5 ruling 6: these three failed on the mobile project (a Pixel 7, 412px) since the Rankings
 * redesign (295def25), which puts the filters behind a summary line's toggle below 600px and
 * drops the Execution column there (RankingRow's TABLE_UP): waiting for `filter-metric` to be
 * visible timed out. The page is now opened on the filters' own container, which every width
 * shows; the sort test opens the toggle first, as rankings.spec.ts's openFilters does.
 */
async function openRankings(page: Page): Promise<void> {
  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-filters');
  const toggle = page.getByTestId('rankings-filters-toggle');
  if (await toggle.isVisible()) await toggle.click();
  await expect(page.getByTestId('filter-metric')).toBeVisible();
}

test.describe('the execution column', () => {
  // The column is not drawn under 600px (the redesign's phone row); its two cell tests are
  // desktop only. The sort order is a filter, so it runs on both.
  test('shows the score and opens compare mode on that fight', async ({ page }) => {
    test.skip(
      test.info().project.name !== 'desktop',
      'no Execution column under 600px (RankingRow TABLE_UP)',
    );
    await page.route('**/v1/rankings?**', (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(page1([{ ...base, execution_score: 0.92 }])),
      }),
    );
    await openRankings(page);
    const cell = page.getByTestId('ranking-execution').first();
    await expect(cell).toHaveText('92%');
    await expect(cell).toHaveAttribute('href', '/sim?source=fight&ref=fixture2abcd%3A2&mode=compare');
  });

  test('says why a fight has no score instead of leaving the cell blank', async ({ page }) => {
    test.skip(
      test.info().project.name !== 'desktop',
      'no Execution column under 600px (RankingRow TABLE_UP)',
    );
    await page.route('**/v1/rankings?**', (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(page1([{ ...base, execution_score: null }])),
      }),
    );
    await openRankings(page);
    const cell = page.getByTestId('ranking-execution').first();
    await expect(cell).toHaveText('—');
    await expect(cell).toHaveAttribute('title', simCopy.executionUnscored);
  });

  test('offers Execution as a sort order and puts it in the URL', async ({ page }) => {
    await page.route('**/v1/rankings?**', (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(page1([{ ...base, execution_score: 0.92 }])),
      }),
    );
    await openRankings(page);
    await page.getByTestId('filter-metric').selectOption('execution');
    await expect(page).toHaveURL(/metric=execution/);
  });
});

/**
 * "Best per encounter" rows are `CharacterFight`, the same shape "Every ranked fight" rows
 * are, and already carry `execution_score`/`report_id`/`fight_index` -- fix round 1 closes
 * the gap the review flagged: this table gets the column too, per the brief's "beside the
 * parse percentile on every ranked fight".
 */
function character(best: unknown[]) {
  return {
    ok: true,
    data: {
      character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', class: 'Priest' },
      best,
      history: [],
      builds_seen: [],
    },
    error: null,
    request_id: 'r',
  };
}

const bestRow = {
  encounter: 'Warden Kelthas',
  encounter_id: 9001,
  difficulty: 8,
  metric: 'hps',
  value: 1840,
  percentile: 96.2,
  spec: 'Discipline',
  fought_at: '2026-12-09T22:10:00Z',
  report_id: 'fixture2abcd',
  fight_index: 3,
};

test.describe('the execution column on the character page', () => {
  test('shows the score on "Best per encounter" and opens compare mode on that fight', async ({ page }) => {
    await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(character([{ ...bestRow, execution_score: 0.92 }])),
      }),
    );
    await page.goto('/character/us/hardcore/elyra-duskvale');
    const cell = page.getByTestId('character-best-execution').first();
    await expect(cell).toHaveText('92%');
    await expect(cell).toHaveAttribute('href', '/sim?source=fight&ref=fixture2abcd%3A3&mode=compare');
  });

  test('says why a "Best per encounter" fight has no score instead of leaving the cell blank', async ({
    page,
  }) => {
    await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(character([{ ...bestRow, execution_score: null }])),
      }),
    );
    await page.goto('/character/us/hardcore/elyra-duskvale');
    const cell = page.getByTestId('character-best-execution').first();
    await expect(cell).toHaveText('—');
    // L8.4 fix round 5 item 7: the profile names the dash "Not scored yet"
    // (lib/character-profile.ts's characterTableCopy.executionUnscored, which this spec cannot
    // import: the module pulls in the generated BiS data).
    await expect(cell).toHaveAttribute('title', 'Not scored yet');
    await expect(cell).toHaveAttribute('aria-label', 'Not scored yet');
  });
});
