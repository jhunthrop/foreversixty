// web/tests/e2e/report-live.spec.ts
import { readFile } from 'node:fs/promises';
import path from 'node:path';
import { expect, test } from '@playwright/test';

const FIXTURES = path.resolve(process.cwd(), 'src/fixtures/report');
const DATA = '/logs-data/reports/fixture2live';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

test('a live fight polls, then settles when it closes', async ({ page }) => {
  const report = JSON.parse(await readFile(path.join(FIXTURES, 'report.json'), 'utf8'));
  const summary = JSON.parse(await readFile(path.join(FIXTURES, 'fights/3/summary.json'), 'utf8'));
  const meta = JSON.parse(await readFile(path.join(FIXTURES, 'meta.json'), 'utf8'));

  const open = {
    ...report,
    report_id: 'fixture2live',
    fights: [{ ...report.fights[2], in_progress: true, kill: false, duration_ms: 12_000 }],
  };
  const closed = { ...open, fights: [{ ...report.fights[2], in_progress: false }] };

  let reportCalls = 0;
  await page.route(`**${DATA}/report.json`, (route) => {
    reportCalls += 1;
    return route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify(reportCalls < 3 ? open : closed),
    });
  });
  await page.route(`**${DATA}/fights/3/live.json`, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ...summary, duration_ms: 12_000 }),
    }),
  );
  await page.route(`**${DATA}/fights/3/summary.json`, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(summary) }),
  );
  await page.route('**/v1/reports/fixture2live', (route) =>
    route.fulfill(
      envelope({
        ...meta,
        id: 'fixture2live',
        status: 'live',
        data_base_url: DATA,
        fights: open.fights,
      }),
    ),
  );

  await page.goto('/reports/fixture2live');

  await expect(page.getByTestId('report-live')).toBeVisible();
  await expect(page.getByTestId('fight-3-outcome')).toHaveText('Live');

  // Two polls at five seconds each, plus slack.
  await expect(page.getByTestId('fight-3-outcome')).toHaveText('Kill', { timeout: 20_000 });
  await expect(page.getByTestId('report-live')).toHaveCount(0);
  expect(reportCalls).toBeGreaterThanOrEqual(3);
});

test('a closed report never polls', async ({ page }) => {
  let calls = 0;
  await page.route('**/logs-data/reports/fixture2abcd/report.json', (route) => {
    calls += 1;
    return route.continue();
  });
  await page.goto('/reports/fixture2abcd');
  await expect(page.getByTestId('report-title')).toBeVisible();
  await page.waitForTimeout(12_000);
  expect(calls).toBe(1);
});

// Holds every matching request open until release() is called, then fulfils each with
// `body` -- the same shape report-tabs.spec.ts's heldRoute() uses for the fight-switch
// races, extended with a call counter: the poll test below needs to prove a *duplicate*
// request happens (or does not), not just that a response lands late.
async function heldRoute(
  page: import('@playwright/test').Page,
  url: string,
  body: unknown,
): Promise<{ started: Promise<void>; release: () => void; callCount: () => number }> {
  let count = 0;
  let markStarted = (): void => {};
  let release = (): void => {};
  const started = new Promise<void>((resolve) => (markStarted = resolve));
  const gate = new Promise<void>((resolve) => (release = resolve));

  await page.route(url, async (route) => {
    count += 1;
    markStarted();
    await gate;
    await route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(body) });
  });

  return { started, release: () => release(), callCount: () => count };
}

function summaryRosterRows(page: import('@playwright/test').Page): import('@playwright/test').Locator {
  return page.getByTestId('summary-tab').locator('li[data-testid^="roster-"]');
}

// The poll can be mid-fetchLive for the fight on screen at the exact moment the visitor
// switches to a different one -- picking a fight does not cancel a poll tick already in
// flight, the same as it does not cancel a fight-selection fetch (report-tabs.spec.ts's
// two interleaving tests).
//
// A naive assertion on the settled DOM ("does the wrong roster ever show?") cannot catch a
// missing guard here: ReportView.svelte's Effect 3 also depends on `summary?.fight_index`,
// so the moment an unguarded write paints fight 3's data under fight 2's selection, Effect
// 3 notices the mismatch and immediately re-calls loadFight(2) to correct it -- often from
// its own cache, within the same reactive flush, faster than anything Playwright's polling
// assertions can observe. That self-heal is real and welcome, but it means the *visible*
// symptom of a missing guard is not "the wrong roster sticks"; it is "a wasted, duplicate
// network request for the fight the visitor is on", fired by Effect 3's correction. This
// test forces that duplicate into the open by holding fight 2's own summary fetch open
// too, so it is still uncached (a guaranteed cache miss) at the moment the stale live
// response lands and Effect 3 re-fires.
test('a stale live response does not trigger a second fetch for the fight the visitor switched to', async ({
  page,
}) => {
  const report = JSON.parse(await readFile(path.join(FIXTURES, 'report.json'), 'utf8'));
  const summary3 = JSON.parse(await readFile(path.join(FIXTURES, 'fights/3/summary.json'), 'utf8'));
  const summary2 = JSON.parse(await readFile(path.join(FIXTURES, 'fights/2/summary.json'), 'utf8'));
  const meta = JSON.parse(await readFile(path.join(FIXTURES, 'meta.json'), 'utf8'));

  // report.json never changes shape in this test: fight 3 stays in_progress throughout, so
  // the poll keeps running (and keeps re-fetching live.json) regardless of which fight is
  // selected on screen -- selecting a fight does not touch the poll's own effect.
  const open = {
    ...report,
    report_id: 'fixture2live',
    fights: [report.fights[0], report.fights[1], { ...report.fights[2], in_progress: true, kill: false }],
  };

  await page.route(`**${DATA}/report.json`, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(open) }),
  );
  await page.route(`**${DATA}/fights/3/summary.json`, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(summary3) }),
  );
  await page.route('**/v1/reports/fixture2live', (route) =>
    route.fulfill(
      envelope({ ...meta, id: 'fixture2live', status: 'live', data_base_url: DATA, fights: open.fights }),
    ),
  );
  // Both held: fight 2's own summary fetch has to still be in flight -- an unconditional
  // cache miss -- when the live response for fight 3 is released, or the race never forms.
  const live = await heldRoute(page, `**${DATA}/fights/3/live.json`, summary3);
  const fight2 = await heldRoute(page, `**${DATA}/fights/2/summary.json`, summary2);

  await page.goto('/reports/fixture2live?fight=3');
  await expect(page.getByTestId('report-live')).toBeVisible();
  await expect(summaryRosterRows(page)).toHaveCount(5);

  // Wait for the poll's first tick to actually reach fetchLive for fight 3 before moving.
  await live.started;

  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-2').click();
  await expect(page.getByTestId('fight-2')).toHaveAttribute('aria-current', 'true');
  // Fight 2's own summary request has started (held, uncached) but not yet resolved.
  await fight2.started;
  expect(fight2.callCount()).toBe(1);

  // The stale live response for fight 3 lands now, while fight 2's own fetch is still
  // outstanding: the guard's exact job. Without it, this write's new object reference
  // reruns Effect 3, which finds `summary.fight_index` (3) no longer matches `state.fight`
  // (2) and calls loadFight(2) again -- a second, wasted fights/2/summary.json request,
  // since the first has not cached yet.
  const answered = page.waitForResponse((response) => response.url().includes('/fights/3/live.json'));
  live.release();
  await answered;
  await page.waitForTimeout(300);

  expect(fight2.callCount()).toBe(1);

  fight2.release();
  await expect(summaryRosterRows(page)).toHaveCount(1);
  await expect(page.getByTestId('report-live')).toHaveCount(0);
  await expect(page.getByTestId('report-fight-error')).toHaveCount(0);
});

test('the timelines and events views render from the summary', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&view=timelines');
  await expect(page.getByTestId('timelines')).toContainText('Baelgrim');
  await expect(page.getByTestId('lane-Player-4184-000000A4')).toBeVisible();

  await page.goto('/reports/fixture2abcd?fight=3&view=events');
  await expect(page.getByTestId('event-list')).toContainText('Thalgrit died');
  await page.getByTestId('events-search').fill('frostbolt');
  await expect(page.getByTestId('event-list')).toContainText('Frostbolt');
  await expect(page.getByTestId('event-list')).not.toContainText('died');
});
