// web/tests/e2e/rankings-encounter-picker.spec.ts
// Spec section 4: the Characters board needs an encounter picker on bare /rankings rather
// than demanding an encounter the visitor cannot choose, and with no encounter data yet it
// says so plainly and offers the Guilds board instead.
import { expect, test } from '@playwright/test';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

test('bare /rankings offers a picker instead of asking the API for an encounter it does not have', async ({
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
  let rankingsCalled = false;
  await page.route('**/v1/rankings?**', (route) => {
    rankingsCalled = true;
    return route.fulfill(
      envelope({ rows: [], total: 0, page: 1, per_page: 100, updated_at: '2026-12-09T00:00:00Z' }),
    );
  });

  await page.goto('/rankings');

  const picker = page.getByTestId('encounter-picker');
  await expect(picker.getByRole('link', { name: 'Warden Kelthas' })).toHaveAttribute(
    'href',
    '/rankings/warden-kelthas',
  );
  await expect(picker.getByRole('link', { name: 'Skolex the Insatiable' })).toBeVisible();
  expect(rankingsCalled).toBe(false);
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
  const guildsButton = page.getByTestId('rankings-try-guilds');
  await expect(guildsButton).toHaveText('Guilds');

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
  await expect(progressButton).toHaveText('Progress');
});
