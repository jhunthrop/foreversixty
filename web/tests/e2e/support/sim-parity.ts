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
import { closeRunLine, openRunLine } from './sim-page';

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
/** The mocks' signed-in character (design/mocks/ws_sim.py `YOU`). */
const MOCK_CHARACTER = 'Thandrel';

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
].map((row) => ({
  ...row,
  spec: FIXTURE.request.spec,
  engine_version: 'bf72769',
  kind: 'run',
  state: 'done',
  seen: true,
  finished_at: row.created_at,
}));

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
/**
 * Waits until the first page has stored its /v1/me answer, so the character's page (the second
 * load) always reads it before paint: Base.astro's head script then sets data-server-sims for a
 * payer and the header hides its Premium door from that hint (data-premium-nav="hide"), as a
 * returning reader sees it. Without the wait the second load raced the store, and the chrome
 * capture flipped between "hide" and the after-paint "hold".
 */
async function settleCachedMe(page: Page): Promise<void> {
  await page.waitForFunction(() => {
    const key = document.querySelector<HTMLElement>('[data-me-storage-key]')?.dataset.meStorageKey;
    return key !== undefined && localStorage.getItem(key) !== null;
  });
}

async function openFuryWarrior(page: Page): Promise<void> {
  await page.goto('/sim');
  await settleCachedMe(page);
  const href = await page
    .locator('[data-region="section-sim-a-best-in-slot-set"] a', { hasText: 'Fury' })
    .getAttribute('href');
  expect(href).not.toBeNull();
  // The mock's character: the rail carries the alliance set, so the race is swapped to the mock's
  // Troll, and the code names its character as the addon's export does (`who=`), so the title is
  // the character's name (ruling S1) and the source a linked addon export.
  const url = new URL(href ?? '', 'http://localhost');
  const code = (url.searchParams.get('code') ?? '').replace(':night-elf:', ':troll:');
  url.searchParams.set('code', `${code}|who=${MOCK_CHARACTER}:Forever`);
  url.searchParams.set('source', 'addon');
  url.searchParams.delete('run');
  await page.goto(`${url.pathname}${url.search}`);
  await expect(page.getByTestId('sim-run-button')).toBeEnabled({ timeout: 15_000 });
  // The Gear section's worn doll (SimGearSection on the shared frame; the list form under 423 px).
  await expect(page.getByTestId('sim-gear-doll')).toBeVisible();
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
export async function serveMockData(page: Page): Promise<void> {
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
    // The fixture ran 10,000 runs: High, so the bar and the result read one run count (M3). On a
    // phone the run line folds behind "Change": open it first (a no-op at 1440).
    await openRunLine(page);
    await page.getByTestId('control-precision').selectOption('high', { force: true });
    // A change folds the phone run line again: reopen it for More settings.
    await openRunLine(page);
    await page.getByTestId('sim-settings-more').click();
    await page.getByTestId('sim-server-run').click();
    await expect(page.getByTestId('sim-answer')).toHaveAttribute('data-state', 'result', { timeout: 15_000 });
    // A phone run left the run line open; a player reads the result under its one line.
    await closeRunLine(page);
    await drawWholeReport(page);
  }
}

/** The tool pages' parity states (parity-sim-tools.spec.ts, design/mocks/web/simtools-parity.json). */
export type ToolParityState = 'gear' | 'gear-empty' | 'talents' | 'drops' | 'weights' | 'specs';

/**
 * The tool mocks' character (design/mocks/ws_simtools.py `CODE`): a level 60 Orc Fury Warrior
 * entered by hand, wearing Lionheart Helm and Arcanite Reaper, handed to the tool by `?code=` as
 * a tab carries it. Two worn items keep every candidate the mock draws readable from the build's
 * item file alone.
 */
export const TOOL_CHARACTER = (build: string): string =>
  `FS1:${build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784|level=60`;

async function openToolWithFury(page: Page, tool: string): Promise<void> {
  const params = new URLSearchParams({ code: TOOL_CHARACTER(BUILD) });
  await page.goto(`/sim/${tool}?${params.toString()}`);
  await expect(page.getByTestId('sim-character')).toBeVisible({ timeout: 15_000 });
  await expect(page.getByTestId('sim-run-bulk')).toBeVisible();
}

/** The mocks draw a desktop's run line: a CI runner's few cores would add the low-core cap note. */
const MOCK_CORES = 8;

export async function openToolState(page: Page, state: ToolParityState): Promise<void> {
  await page.addInitScript((cores) => {
    Object.defineProperty(navigator, 'hardwareConcurrency', { get: () => cores });
  }, MOCK_CORES);
  await serveMockData(page);
  if (state === 'specs') {
    // No parse has been measured yet: every damage spec reads "Not yet", as production does.
    await page.route('**/v1/specs**', (route) => route.fulfill(envelope({ specs: [] })));
    await page.goto('/sim/specs');
    await expect(page.getByTestId('spec-warrior-fury')).toBeVisible({ timeout: 15_000 });
    return;
  }
  if (state === 'gear-empty') {
    await page.goto('/sim/gear');
    await expect(page.getByTestId('sim-addon-input')).toBeVisible({ timeout: 15_000 });
    return;
  }
  // Signed out, the saved-builds read is out of reach: the talent builds say so and go on.
  await page.route('**/v1/builds?*', (route) => route.abort('connectionrefused'));
  await openToolWithFury(page, state);
  const ready: Record<Exclude<ToolParityState, 'specs' | 'gear-empty'>, string> = {
    gear: 'sim-slot-grid',
    talents: 'sim-talent-candidates',
    drops: 'sim-source-picker',
    weights: 'sim-weights-pick',
  };
  await expect(page.getByTestId(ready[state])).toBeVisible({ timeout: 15_000 });
  await expect(page.locator('[aria-busy="true"]')).toHaveCount(0, { timeout: 15_000 });
}
