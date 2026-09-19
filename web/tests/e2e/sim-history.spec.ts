import { expect, test, type Page } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';
import { assertNoHorizontalScroll } from './support/phone-scroll';

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

/** Stubs a signed-in player with the two-row fixture above, honouring `kind=`. */
async function stubHistory(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill(envelope({ user: { premium: false }, characters: [] })),
  );
  await page.route('**/v1/sims?mine=1*', (route) => {
    const kind = new URL(route.request().url()).searchParams.get('kind');
    const rows = kind === null ? ROWS : ROWS.filter((row) => row.kind === kind);
    return route.fulfill(envelope({ rows, total: rows.length, page: 1, per_page: 100 }));
  });
}

test('the history lists every kind and filters to one', async ({ page }) => {
  await stubHistory(page);

  await page.goto('/sim');
  await expect(page.getByTestId('sim-history')).toBeVisible();
  await expect(page.getByTestId('sim-history-kind-aaaaaaaaaaaa')).toHaveText(simCopy.kindLabel.run);
  await expect(page.getByTestId('sim-history-kind-bbbbbbbbbbbb')).toHaveText(simCopy.kindLabel.gear);
  await expect(page.getByTestId('sim-history-headline-bbbbbbbbbbbb')).toHaveText('+41 DPS from Viskag');

  await page.getByTestId('sim-history-filter').selectOption('gear');
  await expect(page.getByTestId('sim-history-bbbbbbbbbbbb')).toBeVisible();
  await expect(page.getByTestId('sim-history-aaaaaaaaaaaa')).toHaveCount(0);
});

// sim-phone.spec.ts's own audit never signs in (no /v1/me stub), so it never renders this
// panel -- the three-cell row (kind pill, title, headline) added here needs its own phone
// check rather than borrowing that file's coverage on trust.
test.describe('phone', () => {
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('the three-cell row wraps without a sideways scroll, and every row still clears 44px', async ({
    page,
  }) => {
    await stubHistory(page);
    await page.goto('/sim');
    await expect(page.getByTestId('sim-history')).toBeVisible();

    const width = page.viewportSize()?.width ?? 412;
    await assertNoHorizontalScroll(page, width);

    for (const simId of ['aaaaaaaaaaaa', 'bbbbbbbbbbbb']) {
      const box = await page.getByTestId(`sim-history-${simId}`).boundingBox();
      expect(box?.height ?? 0, simId).toBeGreaterThanOrEqual(44);
    }
  });
});
