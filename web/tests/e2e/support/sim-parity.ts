// web/tests/e2e/support/sim-parity.ts
// The simulator's parity states (parity-sim.spec.ts, design/mocks/web/sim-parity.json), each the
// mock's data on the built page: the fixture run (web/src/fixtures/sim/result-full.json) and its
// character. The character reaches /sim the way a cold visitor's does: the landing rail's link to
// the published Fury set, the mock's spec at 60 with every slot filled. The session is
// stubbed (one character, three saved sims, as the mock); the live result is that character's
// server run, answered with the fixture itself (it ran on our servers, as the mock's label says).
import { existsSync, mkdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, type Page } from '@playwright/test';
import { writeSimNames } from '../../../scripts/sync-data.mjs';

export type SimParityState = 'landing-out' | 'landing-in' | 'result' | 'saved';

const FIXTURE = JSON.parse(
  readFileSync(
    path.join(import.meta.dirname, '..', '..', '..', 'src', 'fixtures', 'sim', 'result-full.json'),
    'utf8',
  ),
) as { request: { spec: string } };

const WEB_ROOT = path.join(import.meta.dirname, '..', '..', '..');
const BUILD = (
  JSON.parse(readFileSync(path.join(WEB_ROOT, 'src', 'data', 'active-build.json'), 'utf8')) as {
    build: string;
  }
).build;
/** The build the mocks read (ws_sim.py: data/builds/<build>), and the names the site derives from it. */
const REAL_BUILD_DIR = path.join(WEB_ROOT, '..', 'data', 'builds', BUILD);
const DERIVED_DIR = path.join(WEB_ROOT, 'test-results', 'sim-parity-data', BUILD);

/** The saved run the mocks draw, prerendered under FOREVER_DATA=fixture. */
export const SAVED_SIM_PATH = '/sim/simfullfixt1';
const SERVER_RUN_ID = 'parityrun001';

function envelope(data: unknown): { status: number; contentType: string; body: string } {
  return {
    status: 200,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'req-parity' }),
  };
}

/** The mock's "Your sims": three saved runs of the same spec. */
const HISTORY = [
  { sim_id: 'aaaaaaaaaaaa', title: 'Raid buffs', dps: 899.6, created_at: '2026-11-03T20:00:00Z' },
  { sim_id: 'bbbbbbbbbbbb', title: 'No buffs', dps: 402.3, created_at: '2026-11-01T20:00:00Z' },
  { sim_id: 'cccccccccccc', title: 'Raid buffs', dps: 881.2, created_at: '2026-10-29T20:00:00Z' },
].map((row) => ({ ...row, spec: FIXTURE.request.spec, engine_version: 'bf72769', kind: 'run' }));

async function stubSession(page: Page, serverSims: boolean): Promise<void> {
  const port = process.env.E2E_PORT ?? '4321';
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', url: `http://localhost:${port}` }]);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        user: {
          id: 7,
          battletag: 'Thandrel#1234',
          email: null,
          role: 'user',
          anonymize: false,
          premium: serverSims,
        },
        characters: [],
        guilds: [],
        entitlements: { server_sims: serverSims },
      }),
    ),
  );
  await page.route('**/v1/sims?mine=1*', (route) =>
    route.fulfill(envelope({ rows: HISTORY, total: HISTORY.length, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/sims/server*', (route) => route.fulfill(envelope({ rows: [] })));
}

/**
 * /sim with a level 60 Fury warrior's full raid set: the landing rail's own "Fury ›" link (the
 * published BiS set, the mock's spec), without its `run=1`, so the page stops at "No sim yet".
 */
async function openFuryWarrior(page: Page): Promise<void> {
  await page.goto('/sim');
  const href = await page
    .locator('[data-region="section-sim-a-best-in-slot-set"] a', { hasText: 'Fury' })
    .getAttribute('href');
  expect(href).not.toBeNull();
  // The mock's Troll: the rail carries the alliance set, so the race in the code is swapped.
  await page.goto((href ?? '').replace(/[&?]run=1\b/, '').replace(/(%3A|:)night-elf(%3A|:)/, '$1troll$2'));
  await expect(page.getByTestId('sim-run-button')).toBeEnabled({ timeout: 15_000 });
  await expect(page.locator('[data-region="slots"]')).toBeVisible();
}

async function stubServerRun(page: Page): Promise<void> {
  const result = readFileSync(
    path.join(import.meta.dirname, '..', '..', '..', 'src', 'fixtures', 'sim', 'result-full.json'),
    'utf8',
  );
  await page.route('**/v1/sims/run', (route) => route.fulfill(envelope({ sim_id: SERVER_RUN_ID })));
  await page.route(`**/v1/sims/${SERVER_RUN_ID}/progress`, (route) =>
    route.fulfill(envelope({ state: 'done', iterations_done: 10_000 })),
  );
  await page.route(`**/v1/sims/${SERVER_RUN_ID}`, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: `{"ok":true,"data":${result},"error":null,"request_id":"req-parity"}`,
    }),
  );
}

/** Every lazy section of a report drawn: scroll it all into view, then back to the top. */
async function drawWholeReport(page: Page): Promise<void> {
  await page.evaluate(async () => {
    for (let y = 0; y < document.body.scrollHeight; y += 600) {
      window.scrollTo(0, y);
      await new Promise((resolve) => setTimeout(resolve, 60));
    }
    window.scrollTo(0, 0);
  });
  await expect(page.locator('[aria-busy="true"]')).toHaveCount(0, { timeout: 15_000 });
}

let derived: Promise<void> | undefined;

/**
 * The mock's data on the page: every /data/<build>/ file the fixture build serves is answered from
 * the committed build the mocks are generated from (the fixture build's planner data is a small
 * stand-in: seven warrior items, a fixture talent tree), and the simnames files the site derives
 * at sync time are derived from it the same way (scripts/sync-data.mjs `writeSimNames`).
 */
async function serveMockData(page: Page): Promise<void> {
  derived ??= (async () => {
    mkdirSync(DERIVED_DIR, { recursive: true });
    await writeSimNames(REAL_BUILD_DIR, DERIVED_DIR);
  })();
  await derived;
  await page.route(`**/data/${BUILD}/**`, async (route) => {
    const rel = decodeURIComponent(new URL(route.request().url()).pathname).replace(`/data/${BUILD}/`, '');
    const file = [path.join(DERIVED_DIR, rel), path.join(REAL_BUILD_DIR, rel)].find((candidate) =>
      existsSync(candidate),
    );
    if (file === undefined) return route.continue();
    return route.fulfill({ path: file });
  });
}

export async function openSimState(page: Page, state: SimParityState): Promise<void> {
  await serveMockData(page);
  if (state === 'landing-out') {
    await page.goto('/sim');
    await expect(page.getByTestId('sim-view')).toBeVisible();
    await expect(page.locator('[data-region="section-sim-a-best-in-slot-set"]')).toBeVisible();
    return;
  }
  if (state === 'saved') {
    await page.goto(SAVED_SIM_PATH);
    await expect(page.getByTestId('sim-view')).toBeVisible();
    await drawWholeReport(page);
    return;
  }
  await stubSession(page, state === 'result');
  if (state === 'result') await stubServerRun(page);
  await openFuryWarrior(page);
  await expect(page.getByTestId('sim-history')).toBeVisible();
  if (state === 'result') {
    await page.getByTestId('sim-settings-more').click();
    await page.getByTestId('sim-server-run').click();
    await expect(page.getByTestId('sim-answer')).toHaveAttribute('data-state', 'result', { timeout: 15_000 });
    await drawWholeReport(page);
  }
}
