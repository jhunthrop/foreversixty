import { expect, test } from '@playwright/test';
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

test.describe('the execution column', () => {
  test('shows the score and opens compare mode on that fight', async ({ page }) => {
    await page.route('**/v1/rankings?**', (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(page1([{ ...base, execution_score: 0.92 }])),
      }),
    );
    await page.goto('/rankings/warden-kelthas');
    const cell = page.getByTestId('ranking-execution').first();
    await expect(cell).toHaveText('92%');
    await expect(cell).toHaveAttribute('href', '/sim?source=fight&ref=fixture2abcd%3A2&mode=compare');
  });

  test('says why a fight has no score instead of leaving the cell blank', async ({ page }) => {
    await page.route('**/v1/rankings?**', (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(page1([{ ...base, execution_score: null }])),
      }),
    );
    await page.goto('/rankings/warden-kelthas');
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
    await page.goto('/rankings/warden-kelthas');
    await page.getByTestId('filter-metric').selectOption('execution');
    await expect(page).toHaveURL(/metric=execution/);
  });
});
