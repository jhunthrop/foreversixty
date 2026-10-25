// THROWAWAY: acceptance captures for the server-runs lane. Never commit.
import { test, type Locator, type Page } from '@playwright/test';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import { ownerMeAnswer } from './support/character-fixture';

const OUT = '../design/reviews/2026-10-24-server-runs';
const WIDTHS = [390, 1440, 2000] as const;
const PORT = process.env.E2E_PORT ?? '4321';
const NOW = Date.now();
const ago = (m: number): string => new Date(NOW - m * 60_000).toISOString();
const json = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});
const envelope = (data: unknown) => json({ data, error: null });

const base = { engine_version: 'e', title: '', kind: 'run', lane: 'server', spec: 'mage-frost' };
const ROWS = [
  { ...base, sim_id: 'run-estimate', state: 'running', dps: 0, created_at: ago(2), estimate_seconds: 180 },
  { ...base, sim_id: 'run-plain', state: 'running', dps: 0, created_at: ago(3) },
  { ...base, sim_id: 'done-unseen', state: 'done', seen: false, dps: 1204, headline: '1,204 DPS', created_at: ago(30), finished_at: ago(14) },
  { ...base, sim_id: 'done-seen', state: 'done', seen: true, dps: 990, headline: "+41 DPS from Vis'kag", kind: 'gear', created_at: ago(60), finished_at: ago(50) },
  { ...base, sim_id: 'failed-run', state: 'error', dps: 0, failure_reason: 'stopped before completion', created_at: ago(130), finished_at: ago(120) },
];

async function cookie(page: Page): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', url: `http://localhost:${PORT}` }]);
}

async function shot(page: Page, name: string, width: number): Promise<void> {
  await page.waitForTimeout(400);
  await page.screenshot({ path: `${OUT}/${name}-${width}.png`, fullPage: true });
}

interface SimStub {
  rows?: unknown[];
  hold?: boolean;
  fail?: boolean;
  rerun?: number;
  holdRerun?: Promise<void>;
}

async function stubSim(page: Page, opts: SimStub = {}): Promise<void> {
  await cookie(page);
  await page.route('**/v1/me', (r) =>
    r.fulfill(
      envelope({
        user: { premium: true },
        characters: [],
        guilds: [],
        entitlements: mockEntitlements({ server_sims: true }),
      }),
    ),
  );
  await page.route('**/v1/sims?mine=1*', async (r) => {
    const lane = new URL(r.request().url()).searchParams.get('lane');
    if (lane !== 'server') return r.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 100 }));
    if (opts.hold) return new Promise<void>(() => {});
    if (opts.fail) return r.fulfill(json({ data: null, error: { code: 'x', message: 'no' } }, 500));
    const rows = opts.rows ?? [];
    return r.fulfill(envelope({ rows, total: rows.length, page: 1, per_page: 100 }));
  });
  await page.route('**/v1/sims/*/seen', (r) => r.fulfill({ status: 204, body: '' }));
  await page.route('**/v1/sims/*/rerun', async (r) => {
    if (opts.holdRerun) await opts.holdRerun;
    const status = opts.rerun ?? 202;
    if (status !== 202) return r.fulfill(json({ data: null, error: { code: 'x', message: 'no' } }, status));
    return r.fulfill({ ...envelope({ sim_id: 'new-run' }), status: 202 });
  });
}

const finished = (n: number) =>
  Array.from({ length: n }, (_, i) => ({
    sim_id: `sim-${n - i}`,
    spec: 'mage-frost',
    title: 'Frost night',
    dps: 1204 + 40 * (n - i),
    finished_at: '2026-10-24T11:00:00Z',
  }));

async function stubHome(page: Page, fin: unknown[]): Promise<void> {
  await cookie(page);
  const me = ownerMeAnswer();
  const data = me.data as Record<string, unknown>;
  await page.route('**/v1/me', (r) =>
    r.fulfill(
      json({ ...me, data: { ...data, entitlements: mockEntitlements({ server_sims: true }), finished_sims: fin } }),
    ),
  );
  await page.route('**/v1/sims/*/seen', (r) => r.fulfill({ status: 204, body: '' }));
}

async function keyboardFocus(page: Page, target: Locator): Promise<void> {
  await page.keyboard.press('Shift');
  await target.focus();
  await page.evaluate(() => void document.body.offsetHeight);
}

const openAt = async (page: Page, w: number, path: string): Promise<void> => {
  await page.setViewportSize({ width: w, height: w === 390 ? 844 : 1000 });
  await page.goto(path);
};

for (const w of WIDTHS) {
  test(`sim states ${w}`, async ({ page }) => {
    await stubSim(page, { rows: ROWS });
    await openAt(page, w, '/sim');
    await page.getByTestId('server-run-failed-run').waitFor();
    await shot(page, 'sim-payer-states', w);
  });
  test(`sim empty ${w}`, async ({ page }) => {
    await stubSim(page, { rows: [] });
    await openAt(page, w, '/sim');
    await page.getByTestId('server-runs-empty').waitFor();
    await shot(page, 'sim-empty', w);
  });
  test(`home line one ${w}`, async ({ page }) => {
    await stubHome(page, finished(1));
    await openAt(page, w, '/');
    await page.getByTestId('finished-sim-line').waitFor();
    await shot(page, 'home-line-one', w);
  });
  test(`home line three + dismissed ${w}`, async ({ page }) => {
    await stubHome(page, finished(3));
    await openAt(page, w, '/');
    await page.getByTestId('finished-sim-line').waitFor();
    await shot(page, 'home-line-three', w);
    await page.getByTestId('finished-sim-dismiss').click();
    await page.getByTestId('finished-sim-line').waitFor({ state: 'detached' });
    await shot(page, 'home-line-dismissed', w);
  });
}

for (const w of [1440, 390] as const) {
  test(`sim loading ${w}`, async ({ page }) => {
    await stubSim(page, { hold: true });
    await openAt(page, w, '/sim');
    await page.getByTestId('server-runs-reserve').or(page.getByTestId('server-runs')).first().waitFor();
    await shot(page, 'sim-loading', w);
  });
  test(`sim error ${w}`, async ({ page }) => {
    await stubSim(page, { fail: true });
    await openAt(page, w, '/sim');
    await page.getByRole('button', { name: 'Try again' }).waitFor();
    await shot(page, 'sim-error', w);
  });
  test(`failed row states ${w}`, async ({ page }) => {
    let release: () => void = () => {};
    const gate = new Promise<void>((res) => {
      release = res;
    });
    await stubSim(page, { rows: ROWS, holdRerun: gate });
    await openAt(page, w, '/sim');
    const btn = page.getByTestId('server-run-again-failed-run');
    await btn.waitFor();
    await shot(page, 'sim-failed-runagain-rest', w);
    if (w === 1440) {
      await btn.hover();
      await shot(page, 'sim-failed-runagain-hover', w);
      await page.mouse.move(0, 0);
    }
    await keyboardFocus(page, btn);
    await shot(page, 'sim-failed-runagain-focus', w);
    await btn.click();
    await shot(page, 'sim-failed-runagain-pending', w);
    release();
    await page.getByText('Started', { exact: true }).first().waitFor();
    await shot(page, 'sim-failed-runagain-started', w);
  });
  test(`failed row could-not-start ${w}`, async ({ page }) => {
    await stubSim(page, { rows: ROWS, rerun: 500 });
    await openAt(page, w, '/sim');
    await page.getByTestId('server-run-again-failed-run').click();
    await page.getByText('Could not start it again').waitFor();
    await shot(page, 'sim-failed-runagain-couldnotstart', w);
  });
  test(`done row ${w}`, async ({ page }) => {
    await stubSim(page, { rows: ROWS });
    await openAt(page, w, '/sim');
    const row = page.getByTestId('server-run-done-unseen');
    await row.waitFor();
    if (w === 1440) {
      await row.hover();
      await shot(page, 'sim-done-row-hover', w);
      await page.mouse.move(0, 0);
    }
    await keyboardFocus(page, row);
    await shot(page, 'sim-done-row-focus', w);
  });
  test(`other pages with line ${w}`, async ({ page }) => {
    test.skip(w !== 1440);
    await stubHome(page, finished(1));
    for (const [path, name] of [
      ['/bis', 'bis-line'],
      ['/rankings', 'rankings-line'],
      ['/logs', 'logs-line'],
    ] as const) {
      await openAt(page, w, path);
      await page.getByTestId('finished-sim-line').waitFor();
      await shot(page, name, w);
    }
  });
  test(`band focus ${w}`, async ({ page }) => {
    test.skip(w !== 390);
    await stubHome(page, finished(1));
    await openAt(page, w, '/');
    await page.getByTestId('finished-sim-line').waitFor();
    await keyboardFocus(page, page.getByTestId('finished-sim-link'));
    await shot(page, 'home-line-focus-link', w);
    await keyboardFocus(page, page.getByTestId('finished-sim-dismiss'));
    await shot(page, 'home-line-focus-dismiss', w);
  });
  test(`account ${w}`, async ({ page }) => {
    const variants: Array<[string, boolean, { runs_this_month: number } | undefined]> = [
      ['account-runs-4', true, { runs_this_month: 4 }],
      ['account-runs-1', true, { runs_this_month: 1 }],
      ['account-runs-none', true, { runs_this_month: 0 }],
      ['account-runs-nocount', true, undefined],
      ['account-off', false, undefined],
    ];
    await page.route('**/v1/devices', (r) => r.fulfill(json({ ok: true, data: [], error: null, request_id: 'r' })));
    for (const [name, on, serverSims] of variants) {
      await page.unroute('**/v1/me');
      await page.route('**/v1/me', (r) =>
        r.fulfill(
          json({
            ok: true,
            data: {
              user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
              characters: [
                { key: 'us/hardcore/elyra-duskvale', region: 'us', ruleset: 'hardcore', name: 'Elyra Duskvale', class: 'Priest', level: 60 },
              ],
              guilds: [],
              entitlements: mockEntitlements(on ? { server_sims: true, supporter_mark: true } : {}),
              server_sims: serverSims,
            },
            error: null,
            request_id: 'r',
          }),
        ),
      );
      await openAt(page, w, '/account');
      await page.getByTestId('account').waitFor();
      await shot(page, name, w);
    }
  });
}

import { readLayoutShift, trackLayoutShift } from './support/cls';
for (const w of [390, 1440] as const) {
  test(`cls list ${w}`, async ({ page }) => {
    await stubSim(page, { rows: ROWS });
    await openAt(page, w, '/sim');
    await page.getByTestId('server-run-failed-run').waitFor();
    await trackLayoutShift(page);
    await page.reload();
    await page.getByTestId('server-run-failed-run').waitFor();
    console.log(`CLS list ${w}: ${await readLayoutShift(page)}`);
  });
  test(`cls line ${w}`, async ({ page }) => {
    await stubHome(page, finished(1));
    await openAt(page, w, '/');
    await page.getByTestId('finished-sim-line').waitFor();
    await trackLayoutShift(page);
    await page.reload();
    await page.getByTestId('finished-sim-line').waitFor();
    console.log(`CLS line ${w}: ${await readLayoutShift(page)}`);
  });
}
