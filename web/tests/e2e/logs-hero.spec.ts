// web/tests/e2e/logs-hero.spec.ts
// Logs landing spec (2026-10-04) §4.A/§4.A.1/§4.C.2: the header band's hero-selection rule,
// the per-player hook on a signed-in hero, and the canonical sample never double-counted in
// the public feed.
import { expect, test, type Page } from '@playwright/test';
import { meAddonFixture } from '../../src/fixtures/me-addon';
import { routeSampleReportMeta, SAMPLE_REPORT_ID, SAMPLE_REPORT_META } from './support/logs-hero-fixture';

function fulfil(body: unknown, status = 200): { status: number; contentType: string; body: string } {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'e2e' }),
  };
}

async function routeEmptyRecent(page: Page): Promise<void> {
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
}

test('a signed-out visitor sees the canonical sample as the hero, pill and all', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await routeEmptyRecent(page);
  await routeSampleReportMeta(page);

  await page.goto('/logs');
  await expect(page.getByTestId('logs-hero-title')).toContainText(SAMPLE_REPORT_META.title);
  await expect(page.getByTestId('logs-hero-sample-pill')).toHaveText('Sample');
  await expect(page.getByTestId('logs-hero-hook')).toContainText(
    "Open it to see every player's gear next to the planner and the simulator.",
  );
  await expect(page.getByTestId('logs-hero-open')).toHaveAttribute('href', `/reports/${SAMPLE_REPORT_ID}`);
});

test('signed in with zero reports of their own still gets the sample hero, never a blank one', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(meAddonFixture)));
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil([])));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(fulfil({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  await routeEmptyRecent(page);
  await routeSampleReportMeta(page);

  await page.goto('/logs');
  await expect(page.getByTestId('logs-hero-sample-pill')).toHaveText('Sample');
  await expect(page.getByTestId('logs-hero-title')).toContainText(SAMPLE_REPORT_META.title);
});

test("signed in with a report of her own, Zulmara's hero names it and resolves the per-player hook", async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(meAddonFixture)));
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil([])));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(
      fulfil({
        rows: [
          {
            id: 'ownreport01x',
            title: 'Molten Core, week 3',
            zone: 'Molten Core',
            status: 'complete',
            visibility: 'public',
            created_at: '2026-10-03T22:10:00Z',
            fight_count: 9,
            kill_count: 4,
          },
        ],
        total: 1,
        page: 1,
        per_page: 100,
      }),
    ),
  );
  await routeEmptyRecent(page);
  await page.route('**/v1/reports/ownreport01x', (route) =>
    route.fulfill(
      fulfil({
        id: 'ownreport01x',
        title: 'Molten Core, week 3',
        visibility: 'public',
        owner: { id: 1, battletag: 'Fixture#4242' },
        zone: 'Molten Core',
        status: 'complete',
        engine_version: '0.1.0',
        created_at: '2026-10-03T22:10:00Z',
        data_base_url: '/logs-data/reports/ownreport01x',
        players: ['Player-1', 'Player-2'],
        fights: [
          {
            index: 4,
            kind: 'encounter',
            name: 'Lucifron',
            encounter_id: 663,
            difficulty: 1,
            size: 40,
            kill: true,
            in_progress: false,
            start: '2026-10-03T22:30:00Z',
            end: '2026-10-03T22:34:00Z',
            duration_ms: 240_000,
            players: ['Player-1', 'Player-2'],
            deaths: 0,
            npc_kills: 1,
          },
        ],
      }),
    ),
  );
  await page.route('**/fights/4/summary.json**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        engine_version: '0.1.0',
        fight_index: 4,
        duration_ms: 240_000,
        damage_done: [],
        damage_taken: [],
        healing: [],
        healing_taken: [],
        deaths: [],
        auras: [],
        casts: [],
        interrupts: [],
        dispels: [],
        resources: [],
        threat: [],
        combatants: [
          {
            guid: 'Player-1',
            name: 'Zulmara',
            gear: [],
            talents: [],
            consumables: [],
            raid_buffs: [],
            missing_buffs: [],
          },
        ],
        roster: [
          {
            guid: 'Player-1',
            name: 'Zulmara',
            class: 'Hunter',
            role: 'dps',
            active_ms: 240_000,
            activity_pct: 92,
            deaths: 0,
            damage_done: 300_000,
            healing_done: 0,
            damage_taken: 1000,
            dps: 1250.4,
            hps: 0,
            dtps: 4.1,
          },
          {
            guid: 'Player-2',
            name: 'Someone Else',
            class: 'Priest',
            role: 'healer',
            active_ms: 240_000,
            activity_pct: 95,
            deaths: 0,
            damage_done: 10_000,
            healing_done: 400_000,
            damage_taken: 500,
            dps: 41.6,
            hps: 1666.6,
            dtps: 2.1,
          },
        ],
      }),
    }),
  );

  await page.goto('/logs');
  await expect(page.getByTestId('logs-hero-title')).toContainText('Molten Core, week 3');
  await expect(page.getByTestId('logs-hero-sample-pill')).toHaveCount(0);
  const hook = page.getByTestId('logs-hero-hook-line');
  await expect(hook).toBeVisible({ timeout: 10_000 });
  await expect(hook).toContainText('Last kill, Lucifron:');
  await expect(hook).toContainText('Zulmara');
  await expect(hook).toContainText('1,250 DPS');
  await expect(hook.getByRole('link', { name: 'Gear in the planner' })).toHaveAttribute(
    'href',
    /\/planner\?code=/,
  );
  await expect(hook.getByRole('link', { name: 'Sim' })).toHaveAttribute('href', /ownreport01x/);
});

test('the canonical sample never appears twice: filtered out of the public feed by id', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await routeSampleReportMeta(page);
  await page.route('**/v1/reports/recent**', (route) =>
    route.fulfill(
      fulfil({
        rows: [
          {
            id: SAMPLE_REPORT_ID,
            title: SAMPLE_REPORT_META.title,
            created_at: SAMPLE_REPORT_META.created_at,
            fight_count: 2,
            kill_count: 1,
          },
        ],
      }),
    ),
  );

  await page.goto('/logs');
  const panel = page.getByTestId('recent-reports');
  await panel.scrollIntoViewIfNeeded();
  await expect(page.getByTestId('recent-reports-empty')).toBeVisible();
  await expect(panel.getByRole('link', { name: SAMPLE_REPORT_META.title })).toHaveCount(0);
});

test.describe('no layout shift on hydration (spec §8)', () => {
  for (const width of [360, 1280]) {
    test(`the hero's loading, failed and ready states share one reserved height at ${width}px`, async ({
      page,
    }) => {
      await page.setViewportSize({ width, height: 1000 });

      // Failed: the network call never answers.
      await page.route('**/v1/me', (route) => route.abort('failed'));
      await routeEmptyRecent(page);
      await page.goto('/logs');
      await page.getByTestId('logs-hero-error').waitFor();
      const failedHeight = (await page.getByTestId('logs-hero').boundingBox())?.height ?? 0;

      // Ready: the sample hero resolves.
      await page.unroute('**/v1/me');
      await page.route('**/v1/me', (route) =>
        route.fulfill({
          status: 401,
          contentType: 'application/json',
          body: JSON.stringify({ ok: false, data: null, error: null, request_id: 'e2e' }),
        }),
      );
      await routeSampleReportMeta(page);
      await page.goto('/logs');
      await page.getByTestId('logs-hero-description').waitFor();
      const readyHeight = (await page.getByTestId('logs-hero').boundingBox())?.height ?? 0;

      expect(readyHeight, `ready (${readyHeight}) vs failed (${failedHeight}) heights at ${width}px`).toBe(
        failedHeight,
      );
    });
  }
});
