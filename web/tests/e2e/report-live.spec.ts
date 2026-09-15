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
  let isClosed = false;
  await page.route(`**${DATA}/report.json`, (route) => {
    reportCalls += 1;
    if (reportCalls >= 3) isClosed = true;
    return route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify(isClosed ? closed : open),
    });
  });
  await page.route(`**${DATA}/fights/3/live.json`, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ...summary, duration_ms: 12_000 }),
    }),
  );
  // 404 while the fight is open, which is the only thing the server can answer: the engine
  // writes summary.json when a fight closes, so an open fight has live.json and nothing
  // else. Stubbing a 200 here -- as this test used to -- hid a report that could not be
  // opened at all during its first pull, because loadFight always asked for the summary
  // and rendered "No report with that id" when it 404ed.
  let summaryCalls = 0;
  await page.route(`**${DATA}/fights/3/summary.json`, (route) => {
    summaryCalls += 1;
    return isClosed
      ? route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(summary) })
      : route.fulfill({ status: 404, contentType: 'text/plain', body: 'not yet' });
  });
  // The percentile endpoint, counted: a whole-fight parse percentile is not a thing a
  // fight still being written has, and the cache key carries the row's rounded dps, which
  // moves every tick -- so one live 25-player pull meant ~25 uncacheable requests every
  // five seconds, per viewer.
  let percentileCalls = 0;
  await page.route('**/v1/rankings/percentile?**', (route) => {
    percentileCalls += 1;
    return route.fulfill(envelope({ percentile: 91.5 }));
  });
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
  // The open fight rendered from live.json, without ever asking for a summary that cannot
  // exist yet.
  await expect(summaryRosterRows(page)).toHaveCount(5);
  await expect(page.getByTestId('report-error')).toHaveCount(0);
  expect(summaryCalls).toBe(0);
  expect(percentileCalls).toBe(0);

  // Two polls at five seconds each, plus slack.
  await expect(page.getByTestId('fight-3-outcome')).toHaveText('Kill', { timeout: 20_000 });
  await expect(page.getByTestId('report-live')).toHaveCount(0);
  expect(reportCalls).toBeGreaterThanOrEqual(3);

  // And the guard is the fight's own state, not a blanket refusal: the moment it closes,
  // the same rows are worth a percentile and ask for one.
  await expect.poll(() => percentileCalls, { timeout: 10_000 }).toBeGreaterThan(0);
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
  // Fight 2's summary is held throughout: it has to still be in flight -- an unconditional
  // cache miss -- when the live response for fight 3 is released, or the race never forms.
  const fight2 = await heldRoute(page, `**${DATA}/fights/2/summary.json`, summary2);
  // live.json is held from the poll's first tick onwards, not from the page's first
  // request: an open fight is rendered from live.json now (there is no summary.json for
  // one), so holding the very first call would just leave the page loading forever.
  let liveCalls = 0;
  let markStarted = (): void => {};
  let release = (): void => {};
  const livePolled = new Promise<void>((resolve) => (markStarted = resolve));
  const gate = new Promise<void>((resolve) => (release = resolve));
  const liveBody = { status: 200, contentType: 'application/json', body: JSON.stringify(summary3) };
  await page.route(`**${DATA}/fights/3/live.json`, async (route) => {
    liveCalls += 1;
    if (liveCalls > 1) {
      markStarted();
      await gate;
    }
    await route.fulfill(liveBody);
  });

  await page.goto('/reports/fixture2live?fight=3');
  await expect(page.getByTestId('report-live')).toBeVisible();
  await expect(summaryRosterRows(page)).toHaveCount(5);

  // Wait for the poll's first tick to actually reach fetchLive for fight 3 before moving.
  await livePolled;

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
  release();
  await answered;
  await page.waitForTimeout(300);

  expect(fight2.callCount()).toBe(1);

  fight2.release();
  await expect(summaryRosterRows(page)).toHaveCount(1);
  await expect(page.getByTestId('report-live')).toHaveCount(0);
  await expect(page.getByTestId('report-fight-error')).toHaveCount(0);
});

// A private or guild report is refused at /logs-data/, so the island reads it through a
// signed base url from GET /v1/reports/{id}/access -- which expires in ten minutes, while
// a report page is left open for hours. Nothing re-asked for one: after ten minutes,
// selecting an uncached fight told the person who owns the report that it was not public.
test('an expired signed base is re-signed and the same request retried', async ({ page }) => {
  const report = JSON.parse(await readFile(path.join(FIXTURES, 'report.json'), 'utf8'));
  const summary3 = JSON.parse(await readFile(path.join(FIXTURES, 'fights/3/summary.json'), 'utf8'));
  const summary2 = JSON.parse(await readFile(path.join(FIXTURES, 'fights/2/summary.json'), 'utf8'));
  const summary1 = JSON.parse(await readFile(path.join(FIXTURES, 'fights/1/summary.json'), 'utf8'));
  const meta = JSON.parse(await readFile(path.join(FIXTURES, 'meta.json'), 'utf8'));

  // One base per signature, so a request carries the signature it was made with in its
  // own path and the stub can refuse the stale one the way R2 refuses an expired one.
  const signedBase = (signature: number): string => `/signed/${signature}/reports/fixture2live`;
  let issued = 0;
  const live = new Set<number>();
  let accessCalls = 0;
  await page.route('**/v1/reports/fixture2live/access', (route) => {
    accessCalls += 1;
    issued += 1;
    live.add(issued);
    return route.fulfill(envelope({ data_base_url: signedBase(issued) }));
  });
  await page.route('**/v1/reports/fixture2live', (route) =>
    route.fulfill(
      envelope({
        ...meta,
        id: 'fixture2live',
        visibility: 'private',
        status: 'complete',
        data_base_url: '',
        fights: report.fights,
      }),
    ),
  );

  // Every fight the test selects has a summary here. The stub answers 404 for anything
  // else, and a 404 is a real "No report with that id" alert: the final step selects
  // fight 1, and without its summary the test only passed when the assertion ran before
  // the 404 landed, which it did on a fast machine and never in CI.
  const bodies: Record<string, unknown> = {
    'report.json': { ...report, report_id: 'fixture2live' },
    'fights/1/summary.json': summary1,
    'fights/2/summary.json': summary2,
    'fights/3/summary.json': summary3,
  };
  await page.route('**/signed/*/reports/fixture2live/**', (route) => {
    const url = new URL(route.request().url());
    const signature = Number(/\/signed\/(\d+)\//.exec(url.pathname)?.[1]);
    if (!live.has(signature)) {
      return route.fulfill({ status: 403, contentType: 'text/plain', body: 'expired' });
    }
    const file = url.pathname.split(`${signedBase(signature)}/`)[1];
    const body = bodies[file];
    return body === undefined
      ? route.fulfill({ status: 404, contentType: 'text/plain', body: 'no' })
      : route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(body) });
  });

  await page.goto('/reports/fixture2live?fight=3');
  await expect(summaryRosterRows(page)).toHaveCount(5);
  expect(accessCalls).toBe(1);

  // Ten minutes pass. The page is still open and the base it holds no longer signs.
  live.delete(1);

  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-2').click();

  // Fight 2 renders from a freshly signed base rather than "That report is not public".
  await expect(summaryRosterRows(page)).toHaveCount(1);
  await expect(page.getByTestId('report-fight-error')).toHaveCount(0);
  expect(accessCalls).toBe(2);

  // And the fresh base is kept: going back to a fight that is no longer cached does not
  // re-sign a second time.
  await page.getByTestId('fight-1').click();
  await expect(page.getByTestId('fight-1')).toHaveAttribute('aria-current', 'true');
  await expect(page.getByTestId('report-fight-error')).toHaveCount(0);
  expect(accessCalls).toBe(2);
});

// The poll's in_progress branch assigned `summary` and returned without clearing `error`,
// and Effect 3 cannot clear it in this shape: coming back to the fight already on screen
// leaves `summary.fight_index` matching the selection, so Effect 3 returns early and never
// runs its success handler. The alert from the fight the visitor bounced off then sat
// under the header for as long as the live fight stayed open.
test('a stale fight error is cleared by the live poll', async ({ page }) => {
  const report = JSON.parse(await readFile(path.join(FIXTURES, 'report.json'), 'utf8'));
  const summary3 = JSON.parse(await readFile(path.join(FIXTURES, 'fights/3/summary.json'), 'utf8'));
  const meta = JSON.parse(await readFile(path.join(FIXTURES, 'meta.json'), 'utf8'));

  const open = {
    ...report,
    report_id: 'fixture2live',
    fights: [report.fights[0], report.fights[1], { ...report.fights[2], in_progress: true, kill: false }],
  };

  await page.route(`**${DATA}/report.json`, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(open) }),
  );
  await page.route(`**${DATA}/fights/3/live.json`, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(summary3) }),
  );
  // Fight 2 never loads, which is what puts the alert on screen in the first place.
  await page.route(`**${DATA}/fights/2/summary.json`, (route) =>
    route.fulfill({ status: 500, contentType: 'text/plain', body: 'no' }),
  );
  await page.route('**/v1/reports/fixture2live', (route) =>
    route.fulfill(
      envelope({ ...meta, id: 'fixture2live', status: 'live', data_base_url: DATA, fights: open.fights }),
    ),
  );

  await page.goto('/reports/fixture2live?fight=3');
  await expect(page.getByTestId('report-live')).toBeVisible();

  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-2').click();
  await expect(page.getByTestId('report-fight-error')).toBeVisible();

  await page.getByTestId('fight-3').click();
  await expect(page.getByTestId('fight-3')).toHaveAttribute('aria-current', 'true');

  // One poll tick, plus slack: the fight on screen is live and being updated, so an alert
  // about a different fight must not outlive the next update.
  await expect(page.getByTestId('report-fight-error')).toHaveCount(0, { timeout: 15_000 });
  await expect(summaryRosterRows(page)).toHaveCount(5);
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
