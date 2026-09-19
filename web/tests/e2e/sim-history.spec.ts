import { expect, test } from '@playwright/test';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ data, error: null }),
});

const ROWS = [
  {
    sim_id: 'aaaaaaaaaaaa',
    spec: 'warrior-fury',
    dps: 1204,
    engine_version: 'edc0c8e9a',
    created_at: '2026-09-14T10:02:00Z',
    title: 'Raid-buffed, 3:00, single target',
    kind: 'run',
    headline: '1,204 DPS',
  },
  {
    sim_id: 'bbbbbbbbbbbb',
    spec: 'warrior-fury',
    dps: 1245,
    engine_version: 'edc0c8e9a',
    created_at: '2026-09-15T10:02:00Z',
    title: '',
    kind: 'gear',
    headline: '+41 DPS from Viskag',
  },
];

test('the history lists every kind and filters to one', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(envelope({ user: { premium: false }, characters: [] })),
  );
  await page.route('**/v1/sims?mine=1*', (route) => {
    const kind = new URL(route.request().url()).searchParams.get('kind');
    const rows = kind === null ? ROWS : ROWS.filter((row) => row.kind === kind);
    return route.fulfill(envelope({ rows, total: rows.length, page: 1, per_page: 100 }));
  });

  await page.goto('/sim');
  await expect(page.getByTestId('sim-history')).toBeVisible();
  await expect(page.getByTestId('sim-history-kind-aaaaaaaaaaaa')).toHaveText('Sim');
  await expect(page.getByTestId('sim-history-kind-bbbbbbbbbbbb')).toHaveText('Top Gear');
  await expect(page.getByTestId('sim-history-headline-bbbbbbbbbbbb')).toHaveText('+41 DPS from Viskag');

  await page.getByTestId('sim-history-filter').selectOption('gear');
  await expect(page.getByTestId('sim-history-bbbbbbbbbbbb')).toBeVisible();
  await expect(page.getByTestId('sim-history-aaaaaaaaaaaa')).toHaveCount(0);
});
