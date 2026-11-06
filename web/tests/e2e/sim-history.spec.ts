// Your sims in the /sim rail (spec 2026-11-06-sim §4.3, §4.5, §5 "Rail rows"): the signed-in
// player's newest saved sims as open link rows -- title, the kind (or a run's spec) and the
// date, a run's DPS at the right -- under a character. The kind filter and Best per spec
// retired; the account page keeps the full list ("All ›").
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { KIND_TITLES } from '../../src/lib/sim/copy';
import { assertNoHorizontalScroll } from './support/phone-scroll';
import { pasteAndRun } from './support/sim-page';

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
    route.fulfill(envelope({ user: { premium: false }, characters: [], guilds: [] })),
  );
  await page.route('**/v1/sims?mine=1*', (route) => {
    const kind = new URL(route.request().url()).searchParams.get('kind');
    const rows = kind === null ? ROWS : ROWS.filter((row) => row.kind === kind);
    return route.fulfill(envelope({ rows, total: rows.length, page: 1, per_page: 100 }));
  });
}

test('Your sims lists every kind, each row its kind or spec, its date and a link to it', async ({ page }) => {
  await stubHistory(page);
  await pasteAndRun(page);

  const history = page.getByTestId('sim-history');
  await history.scrollIntoViewIfNeeded();
  await expect(history).toBeVisible();
  const run = page.getByTestId('sim-history-aaaaaaaaaaaa');
  await expect(run).toContainText('Raid-buffed, 3:00, single target');
  await expect(run).toContainText('Fury · 14 Sep');
  await expect(run).toContainText('1,204');
  await expect(run).toHaveAttribute('href', '/sim/aaaaaaaaaaaa');
  // An untitled Top Gear result reads its spec as the title and its kind under it.
  const gear = page.getByTestId('sim-history-bbbbbbbbbbbb');
  await expect(gear).toContainText(`${KIND_TITLES.gear} · 15 Sep`);
  await expect(gear).toHaveAttribute('href', '/sim/bbbbbbbbbbbb');
  await expect(history.getByRole('link', { name: /All/ })).toHaveAttribute('href', '/account#sims');
  await expect(page.getByTestId('sim-history-filter')).toHaveCount(0);
});

// sim-phone.spec.ts's own audit never signs in (no /v1/me stub), so it never renders this
// panel -- the three-cell row (kind pill, title, headline) added here needs its own phone
// check rather than borrowing that file's coverage on trust.
test.describe('phone', () => {
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('a row wraps without a sideways scroll, and every row still clears 44px', async ({ page }) => {
    await stubHistory(page);
    await pasteAndRun(page);
    await page.getByTestId('sim-history').scrollIntoViewIfNeeded();
    await expect(page.getByTestId('sim-history')).toBeVisible();

    const width = page.viewportSize()?.width ?? 412;
    await assertNoHorizontalScroll(page, width);

    for (const simId of ['aaaaaaaaaaaa', 'bbbbbbbbbbbb']) {
      const box = await page.getByTestId(`sim-history-${simId}`).boundingBox();
      expect(box?.height ?? 0, simId).toBeGreaterThanOrEqual(44);
    }
  });
});
