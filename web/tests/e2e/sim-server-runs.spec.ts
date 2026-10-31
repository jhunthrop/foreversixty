import type { Page, Route } from '@playwright/test';
import { expect, test } from './support/test';
import { serverRunsCopy } from '../../src/lib/sim/server-runs-copy';
import { readLayoutShift, trackLayoutShift } from './support/cls';

// "Your server runs" on /sim (experience spec 2026-10-24 sections 3 and 4), with the API
// stubbed to the contract in section 10.

const MAX_CLS = 0.02;
const NOW = Date.now();
const minutesAgo = (minutes: number): string => new Date(NOW - minutes * 60_000).toISOString();

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ data, error: null }),
});

const ENTITLEMENTS = {
  server_sims: true,
  retention: false,
  multi_compare: false,
  history: false,
  notifications: false,
  officer_views: false,
  roster_check: false,
  supporter_mark: false,
};

const base = { engine_version: 'e', title: '', kind: 'run', lane: 'server', spec: 'mage-frost' };

const RUNNING_WITH_ESTIMATE = {
  ...base,
  sim_id: 'run-estimate',
  state: 'running',
  dps: 0,
  created_at: minutesAgo(2),
  estimate_seconds: 180,
};
const RUNNING_NO_ESTIMATE = {
  ...base,
  sim_id: 'run-plain',
  state: 'running',
  dps: 0,
  created_at: minutesAgo(3),
};
const DONE_UNSEEN = {
  ...base,
  sim_id: 'done-unseen',
  state: 'done',
  seen: false,
  dps: 1204,
  headline: '1,204 DPS',
  created_at: minutesAgo(30),
  finished_at: minutesAgo(14),
};
const DONE_SEEN = {
  ...base,
  sim_id: 'done-seen',
  state: 'done',
  seen: true,
  dps: 990,
  headline: '+41 DPS from Viskag',
  kind: 'gear',
  created_at: minutesAgo(60),
  finished_at: minutesAgo(50),
};
const FAILED = {
  ...base,
  sim_id: 'failed-run',
  state: 'error',
  dps: 0,
  failure_reason: 'stopped before completion',
  created_at: minutesAgo(130),
  finished_at: minutesAgo(120),
};

interface Stub {
  rows: unknown[];
  rerunStatus?: number;
  rerunRequests: string[];
}

async function stubSignedIn(page: Page, server: boolean): Promise<void> {
  await page
    .context()
    .addCookies([
      { name: 'fs_csrf', value: 'token', url: `http://localhost:${process.env.E2E_PORT ?? '4321'}` },
    ]);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        user: { premium: server },
        characters: [],
        guilds: [],
        entitlements: { ...ENTITLEMENTS, server_sims: server },
      }),
    ),
  );
}

async function stubServerRuns(page: Page, rows: unknown[], rerunStatus = 202): Promise<Stub> {
  const stub: Stub = { rows, rerunStatus, rerunRequests: [] };
  await page.route('**/v1/sims?mine=1*', (route: Route) => {
    const lane = new URL(route.request().url()).searchParams.get('lane');
    const list = lane === 'server' ? stub.rows : [];
    return route.fulfill(envelope({ rows: list, total: list.length, page: 1, per_page: 100 }));
  });
  await page.route('**/v1/sims/*/rerun', (route) => {
    stub.rerunRequests.push(route.request().url());
    if (stub.rerunStatus !== 202) {
      return route.fulfill({
        status: stub.rerunStatus,
        contentType: 'application/json',
        body: JSON.stringify({ data: null, error: { code: 'plan', message: 'no' } }),
      });
    }
    return route.fulfill({ ...envelope({ sim_id: 'new-run' }), status: 202 });
  });
  await page.route('**/v1/sims/*/seen', (route) => route.fulfill({ status: 204, body: '' }));
  return stub;
}

test('a payer sees every row state with the spec words', async ({ page }) => {
  await stubSignedIn(page, true);
  await stubServerRuns(page, [RUNNING_WITH_ESTIMATE, RUNNING_NO_ESTIMATE, DONE_UNSEEN, DONE_SEEN, FAILED]);
  await page.goto('/sim');

  await expect(page.getByTestId('server-runs')).toBeVisible();
  await expect(page.getByRole('heading', { name: serverRunsCopy.heading })).toBeVisible();

  const estimate = page.getByTestId('server-run-run-estimate');
  await expect(estimate).toContainText('Running · takes about 3 minutes');
  await expect(estimate).toContainText('Started 2 minutes ago');
  await expect(estimate.locator('a')).toHaveCount(0);

  const plain = page.getByTestId('server-run-run-plain');
  await expect(plain).toContainText('Running');
  await expect(plain).not.toContainText('takes about');

  const unseen = page.getByTestId('server-run-done-unseen');
  await expect(unseen).toContainText('1,204 DPS');
  await expect(unseen).toContainText('Finished 14 minutes ago');
  await expect(unseen).toContainText(serverRunsCopy.unreadSr);
  await expect(unseen).toHaveAttribute('href', '/sim/done-unseen');
  await expect(unseen).toHaveClass(/server-run-unseen/);

  const seen = page.getByTestId('server-run-done-seen');
  await expect(seen).toContainText('+41 DPS from Viskag');
  await expect(seen).not.toContainText(serverRunsCopy.unreadSr);
  await expect(seen).toHaveAttribute('href', '/sim/done-seen');

  const failed = page.getByTestId('server-run-failed-run');
  await expect(failed).toContainText('Failed: stopped before completion');
  await expect(failed).toContainText('Ended 2 hours ago');
  await expect(failed.locator('a')).toHaveCount(0);
});

test('Run again posts the rerun and the button reads Started', async ({ page }) => {
  await stubSignedIn(page, true);
  const stub = await stubServerRuns(page, [FAILED]);
  await page.goto('/sim');

  const button = page.getByTestId('server-run-again-failed-run');
  await button.click();
  await expect(button).toHaveText(serverRunsCopy.runAgainDone);
  await expect(button).toBeDisabled();
  expect(stub.rerunRequests).toHaveLength(1);
  expect(stub.rerunRequests[0]).toContain('/v1/sims/failed-run/rerun');
});

test('a lapsed plan answers Run again with the Premium line and See plans', async ({ page }) => {
  await stubSignedIn(page, true);
  await stubServerRuns(page, [FAILED], 402);
  await page.goto('/sim');

  await page.getByTestId('server-run-again-failed-run').click();
  const row = page.getByTestId('server-run-failed-run');
  await expect(row).toContainText(serverRunsCopy.planLapsed);
  await expect(row.getByRole('link', { name: serverRunsCopy.seePlans })).toHaveAttribute(
    'href',
    '/premium#server-sims',
  );
  await expect(page.getByTestId('server-run-again-failed-run')).toBeEnabled();
});

test('another Run again failure shows the failure line, and the stored reserve grows with it', async ({
  page,
}) => {
  await stubSignedIn(page, true);
  await stubServerRuns(page, [FAILED], 500);
  await page.goto('/sim');
  const reserve = (): Promise<number> =>
    page.evaluate(() => Number(localStorage.getItem('fs.serverRuns.reserve')));
  await expect(page.getByTestId('server-run-failed-run')).toBeVisible();
  await expect.poll(reserve).toBeGreaterThan(0);
  const before = await reserve();

  await page.getByTestId('server-run-again-failed-run').click();
  await expect(page.getByTestId('server-run-failed-run')).toContainText(serverRunsCopy.runAgainFailed);
  await expect.poll(reserve).toBeGreaterThan(before);
});

for (const route of ['/sim/simfixtureab', '/sim/specs']) {
  test(`a payer on ${route} never reads the server list`, async ({ page }) => {
    await stubSignedIn(page, true);
    const laneRequests: string[] = [];
    await stubServerRuns(page, [DONE_UNSEEN]);
    page.on('request', (request) => {
      if (new URL(request.url()).searchParams.get('lane') === 'server') laneRequests.push(request.url());
    });
    await page.goto(route);
    await expect(page.getByTestId('server-runs')).toHaveCount(0);
    await page.waitForTimeout(1000);
    expect(laneRequests).toEqual([]);
  });
}

test('a free member sees no list, no heading and no empty state', async ({ page }) => {
  await stubSignedIn(page, false);
  await stubServerRuns(page, [DONE_UNSEEN]);
  await page.goto('/sim');

  await expect(page.getByTestId('sim-view')).toBeVisible();
  await expect(page.getByTestId('server-runs')).toHaveCount(0);
  await expect(page.getByText(serverRunsCopy.heading)).toHaveCount(0);
  await expect(page.getByTestId('server-runs-reserve')).toHaveCount(0);
});

test('a payer with no runs sees the one-line empty state', async ({ page }) => {
  await stubSignedIn(page, true);
  await stubServerRuns(page, []);
  await page.goto('/sim');

  await expect(page.getByTestId('server-runs-empty')).toContainText(serverRunsCopy.empty);
});

test('a run that finishes under a visible tab is marked seen at once', async ({ page }) => {
  await stubSignedIn(page, true);
  const stub = await stubServerRuns(page, [RUNNING_WITH_ESTIMATE]);
  const seenCalls: string[] = [];
  await page.route('**/v1/sims/*/seen', (route) => {
    seenCalls.push(route.request().url());
    return route.fulfill({ status: 204, body: '' });
  });
  await page.clock.install();
  await page.goto('/sim');
  await expect(page.getByTestId('server-run-run-estimate')).toContainText('Running');

  stub.rows = [{ ...DONE_UNSEEN, sim_id: 'run-estimate' }];
  await page.clock.runFor(10_500);

  const row = page.getByTestId('server-run-run-estimate');
  await expect(row).toContainText('1,204 DPS');
  await expect(row).not.toContainText(serverRunsCopy.unreadSr);
  expect(seenCalls.some((url) => url.includes('/v1/sims/run-estimate/seen'))).toBe(true);
});

for (const viewport of [
  { width: 390, height: 844 },
  { width: 1440, height: 900 },
]) {
  test(`the hinted payer's list loads without layout shift at ${viewport.width}px`, async ({ page }) => {
    await page.setViewportSize(viewport);
    await stubSignedIn(page, true);
    await stubServerRuns(page, [RUNNING_WITH_ESTIMATE, DONE_UNSEEN, FAILED]);
    // The first visit stores the /v1/me snapshot and the list's height; the second is the hinted one.
    await page.goto('/sim');
    await expect(page.getByTestId('server-run-failed-run')).toBeVisible();

    await trackLayoutShift(page);
    await page.reload();
    await expect(page.getByTestId('server-run-failed-run')).toBeVisible();
    expect(await readLayoutShift(page)).toBeLessThanOrEqual(MAX_CLS);
  });
}
