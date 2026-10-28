// web/tests/captures/log-coach/capture.spec.ts
// Acceptance screenshots for the log coach (design/specs/2026-10-28-log-coach.md section 9,
// rows A1-A11) for the ux-designer and wow-player reviewers. Not a test of behaviour: it drives
// the report row strip, the Coach tab in every state, the writing-to-ready swap (with the logged
// y values), the account line and the Discord messages (rendered from the real handlers'
// payloads.json) with the API stubbed by the coach's own output (src/fixtures/coach) on the worked
// Onyxia pull, writing PNGs, logs and same-scale boards against the mock renders to
// design/reviews/2026-10-27-log-coach/. Widths: 360, 390, 1024, 1440, 1920, 2000.
import { existsSync, mkdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Page } from '@playwright/test';
import { atWidths, settle, sizeTo } from '../support/capture-kit';
import { renderMessage, type Dumped } from '../discord-bot/discord-render';
import {
  GUID,
  PANEL_URL,
  RECENT,
  ROWS_URL,
  STATES,
  routeAccountDevices,
  routeCoach,
  routeRecent,
  signInAs,
  writingThenModel,
  type AccountLineId,
  type CoachStub,
  type StateId,
} from '../../e2e/support/coach-stubs';
import { drainShifts, recordShifts, sumShifts } from '../../e2e/support/layout-shift';
import { PHONE_AND_DESKTOP, WIDTHS, atRest, file, openAt, shoot, writeLog } from './coach-capture-kit';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const WEB = path.resolve(HERE, '../../..');
const BUILDS = path.resolve(WEB, '../data/builds');
const PAYLOADS = path.join(HERE, 'payloads.json');
const POLL_WAIT_MS = 10_000;

/** Mock board per width: the phone, 1024 and 2000 boards exist beside the 1440 one. */
const panelMock = (width: number): string =>
  width <= 390
    ? 'log-coach-panel-phone.png'
    : width === 1024
      ? 'log-coach-panel-1024.png'
      : width >= 1920
        ? 'log-coach-panel-2000.png'
        : 'log-coach-panel.png';
const rowMock = (width: number): string => (width <= 390 ? 'log-coach-row-phone.png' : 'log-coach-row.png');

// ---- A1-A3: the row strip -----------------------------------------------------------------------

atWidths('A1 row strip, Premium owner', WIDTHS, async (page, width) => {
  await signInAs(page, 'premium');
  await routeCoach(page, STATES['ready-model']);
  await openAt(page, ROWS_URL, width);
  await atRest(page);
  await shoot(page, 'a1-row', 'premium', width, rowMock(width));
});

const STRIP_STATES: [string, CoachStub, 'premium' | 'free'][] = [
  ['free', STATES['teaser-several'], 'free'],
  ['free-one', STATES['teaser-one'], 'free'],
  ['free-no-first', STATES['teaser-no-first'], 'free'],
  ['premium-no-first', STATES['raid-buffs-only'], 'premium'],
  ['clean', STATES['no-findings'], 'premium'],
  ['preparing', STATES.preparing, 'premium'],
  ['error', STATES.error, 'premium'],
];

atWidths('A2 row strip states', PHONE_AND_DESKTOP, async (page, width) => {
  for (const [name, stub, viewer] of STRIP_STATES) {
    const fresh = await page.context().newPage();
    await signInAs(fresh, viewer);
    await routeCoach(fresh, stub);
    await openAt(fresh, ROWS_URL, width);
    await atRest(fresh);
    await shoot(
      fresh,
      'a2-strip',
      name,
      width,
      width <= 390 ? 'log-coach-states-phone.png' : 'log-coach-row-states.png',
    );
    await fresh.close();
  }
  const held = await page.context().newPage();
  let release: () => void = () => undefined;
  const gate = new Promise<void>((resolve) => {
    release = resolve;
  });
  await signInAs(held, 'premium');
  await routeCoach(held, { ...STATES['ready-model'], gate });
  await openAt(held, ROWS_URL, width);
  await shoot(held, 'a2-strip', 'loading', width);
  release();
  await held.close();
});

atWidths('A2 row strip hover and focus-visible', [1440], async (page, width) => {
  await signInAs(page, 'premium');
  await routeCoach(page, STATES['ready-model']);
  await openAt(page, ROWS_URL, width);
  const strip = page.getByTestId('coach-row-line');
  await strip.hover();
  await shoot(page, 'a2-strip', 'hover', width);
  await page.mouse.move(0, 0);
  await strip.focus();
  await page.keyboard.press('Shift+Tab');
  await page.keyboard.press('Tab');
  await shoot(page, 'a2-strip', 'focus', width);
});

atWidths('A3 officer on a member row, a stranger, signed out', PHONE_AND_DESKTOP, async (page, width) => {
  await signInAs(page, 'officer');
  await routeCoach(page, STATES.officer);
  await openAt(page, ROWS_URL, width);
  await page.getByTestId('actor-Player-4184-000000A3').getByRole('button').first().click();
  await settle(page);
  await shoot(page, 'a3-officer', 'opened-row', width);
  for (const kind of ['stranger', 'signedOut'] as const) {
    const other = await page.context().newPage();
    await signInAs(other, kind);
    await routeCoach(other, STATES['ready-model']);
    await openAt(other, ROWS_URL, width);
    await shoot(other, 'a3-officer', kind, width);
    await other.close();
  }
});

// ---- A4-A6: the Coach tab -----------------------------------------------------------------------

atWidths('A4 Coach tab, model review', WIDTHS, async (page, width) => {
  await signInAs(page, 'premium');
  await routeCoach(page, STATES['ready-model']);
  await openAt(page, PANEL_URL(), width);
  await atRest(page);
  await shoot(page, 'a4-panel', 'model', width, panelMock(width));
});

async function yOf(page: Page, testId: string): Promise<number> {
  return page
    .getByTestId(testId)
    .locator('h3')
    .first()
    .evaluate((node) => node.getBoundingClientRect().top + window.scrollY);
}

atWidths('A5 Coach tab, writing then swapped', [390, 1440], async (page, width) => {
  let ready = false;
  await signInAs(page, 'premium');
  await routeCoach(
    page,
    writingThenModel(() => ready),
  );
  await openAt(page, PANEL_URL(), width);
  await atRest(page);
  const measure = async (): Promise<Record<string, number>> => ({
    thenLabelY: await yOf(page, 'coach-then'),
    doingWellLabelY: await yOf(page, 'coach-well'),
    cardBottomY: await page
      .getByTestId('coach-card')
      .evaluate((node) => node.getBoundingClientRect().bottom + window.scrollY),
  });
  await shoot(page, 'a5-swap', 'writing', width, panelMock(width));
  const before = await measure();
  ready = true;
  await page.getByTestId('coach-title').filter({ hasText: 'coach' }).waitFor({ timeout: POLL_WAIT_MS });
  await settle(page);
  const after = await measure();
  await shoot(page, 'a5-swap', 'ready', width, panelMock(width));
  writeLog(`log-coach-a5-y-values-${width}.json`, {
    before,
    after,
    equal: JSON.stringify(before) === JSON.stringify(after),
  });
});

const PANEL_STATES: StateId[] = [
  'ready-rules',
  'busy',
  'preparing',
  'no-findings',
  'not-reviewable',
  'not-reviewable-short',
  'not-reviewable-spec',
  'teaser-several',
  'teaser-one',
  'teaser-empty',
  'officer',
  'starter',
  'healer',
  'raid-buffs-only',
  'error',
  'not-available',
];

atWidths('A6 Coach tab states', PHONE_AND_DESKTOP, async (page, width) => {
  for (const state of PANEL_STATES) {
    const fresh = await page.context().newPage();
    const free = state.startsWith('teaser');
    await signInAs(fresh, free ? 'free' : 'premium', state === 'healer' ? ['sunwick'] : ['baelgrim']);
    await routeCoach(fresh, STATES[state]);
    await openAt(fresh, PANEL_URL(state === 'healer' ? GUID.sunwick : undefined), width);
    await atRest(fresh);
    await shoot(
      fresh,
      'a6-state',
      state,
      width,
      width <= 390 ? 'log-coach-states-phone.png' : 'log-coach-states.png',
    );
    await fresh.close();
  }
  const signedOut = await page.context().newPage();
  await signInAs(signedOut, 'signedOut');
  await openAt(signedOut, PANEL_URL(), width);
  await shoot(signedOut, 'a6-state', 'signed-out', width);
  const night = await page.context().newPage();
  await signInAs(night, 'premium');
  await openAt(night, PANEL_URL().replace('fight=3', 'fight=0'), width);
  await shoot(night, 'a6-state', 'whole-night', width);
  const held = await page.context().newPage();
  let release: () => void = () => undefined;
  const gate = new Promise<void>((resolve) => {
    release = resolve;
  });
  await signInAs(held, 'premium');
  await routeCoach(held, { ...STATES['ready-model'], gate });
  await openAt(held, PANEL_URL(), width);
  await shoot(held, 'a6-state', 'loading', width);
  release();
});

// ---- A8: the account line -----------------------------------------------------------------------

const ACCOUNT_STATES = Object.keys(RECENT) as AccountLineId[];

atWidths('A8 account line, every state', PHONE_AND_DESKTOP, async (page, width) => {
  for (const id of ACCOUNT_STATES) {
    const fresh = await page.context().newPage();
    await signInAs(fresh, RECENT[id].viewer);
    await routeAccountDevices(fresh);
    await routeRecent(fresh, RECENT[id].recent);
    await openAt(fresh, '/account', width);
    await atRest(fresh);
    await shoot(
      fresh,
      'a8-account',
      id,
      width,
      width <= 390 ? 'log-coach-account-phone.png' : 'log-coach-account.png',
    );
    await fresh.close();
  }
});

atWidths('A8 account line, loading', PHONE_AND_DESKTOP, async (page, width) => {
  let release: () => void = () => undefined;
  const gate = new Promise<void>((resolve) => {
    release = resolve;
  });
  await signInAs(page, 'premium');
  await routeAccountDevices(page);
  await routeRecent(page, RECENT.premium.recent, gate);
  await openAt(page, '/account', width);
  await shoot(page, 'a8-account', 'loading', width);
  release();
});

// ---- 360 wide: the strip, the account line and the phone fold (spec 1 and 9) ---------------------

atWidths('360: the strip, the account line, and the tab right after it opens', [360], async (page, width) => {
  await signInAs(page, 'premium');
  await routeCoach(page, STATES['ready-model']);
  await routeAccountDevices(page);
  await routeRecent(page, RECENT.premium.recent);
  await openAt(page, ROWS_URL, width);
  await atRest(page);
  await shoot(page, 'a1-row', 'premium', width);
  const account = await page.context().newPage();
  await signInAs(account, 'free');
  await routeAccountDevices(account);
  await routeRecent(account, RECENT.free.recent);
  await openAt(account, '/account', width);
  await shoot(account, 'a8-account', 'free', width);
  await account.close();
  await page.setViewportSize({ width, height: 640 });
  await page.getByTestId('coach-row-line').click();
  await page.getByTestId('coach-fix-first').waitFor();
  await settle(page);
  mkdirSync(path.dirname(file('x', 'x', width)), { recursive: true });
  await page.screenshot({ path: file('phone-fold', 'tab-opened', width), fullPage: false });
});

// ---- A3: the More menu, by viewer ----------------------------------------------------------------

atWidths('A3 More menu open: owner, stranger, signed out', [1440], async (page, width) => {
  for (const kind of ['premium', 'stranger', 'signedOut'] as const) {
    const fresh = await page.context().newPage();
    await signInAs(fresh, kind);
    await routeCoach(fresh, STATES['ready-model']);
    await openAt(fresh, ROWS_URL, width);
    await fresh.getByTestId('tab-more').click();
    await settle(fresh);
    await shoot(fresh, 'a3-more', kind, width);
    await fresh.close();
  }
});

// ---- A7: the layout-shift log -------------------------------------------------------------------

atWidths('A7 layout-shift log for the strip, the panel and the swap', [1280], async (page, width) => {
  const log: Record<string, { shifts: number[]; sum: number }> = {};
  const record = async (name: string, run: (p: Page) => Promise<void>): Promise<void> => {
    const fresh = await page.context().newPage();
    await recordShifts(fresh);
    await run(fresh);
    const shifts = await drainShifts(fresh);
    log[name] = { shifts, sum: sumShifts(shifts) };
    await fresh.close();
  };
  await record('A1 strip (Damage Done, Premium owner)', async (p) => {
    await signInAs(p, 'premium');
    await routeCoach(p, STATES['ready-model']);
    await openAt(p, ROWS_URL, width);
    await p.getByTestId('coach-row-line').waitFor();
  });
  await record('A4 panel (model review)', async (p) => {
    await signInAs(p, 'premium');
    await routeCoach(p, STATES['ready-model']);
    await openAt(p, PANEL_URL(), width);
    await p.getByTestId('coach-fix-first').waitFor();
  });
  await record('A5 writing then swapped', async (p) => {
    let ready = false;
    await signInAs(p, 'premium');
    await routeCoach(
      p,
      writingThenModel(() => ready),
    );
    await openAt(p, PANEL_URL(), width);
    await p.getByTestId('coach-fix-first').waitFor();
    ready = true;
    await p.getByTestId('coach-title').filter({ hasText: 'coach' }).waitFor({ timeout: POLL_WAIT_MS });
  });
  await record('A8 account line (Premium)', async (p) => {
    await signInAs(p, 'premium');
    await routeAccountDevices(p);
    await routeRecent(p, RECENT.premium.recent);
    await openAt(p, '/account', width);
    await p.getByTestId('coach-last-raid').and(p.locator('[data-state="ready"]')).waitFor();
  });
  writeLog(`log-coach-a7-layout-shift-${width}.json`, {
    budget: 0.05,
    runs: log,
    under_budget: Object.values(log).every((run) => run.sum < 0.05),
  });
});

// ---- A9: keyboard run ---------------------------------------------------------------------------

atWidths('A9 keyboard: row to strip to panel heading, focus rings', [1440], async (page, width) => {
  await signInAs(page, 'free');
  await routeCoach(page, STATES['teaser-several']);
  await openAt(page, ROWS_URL, width);
  await page.getByTestId('coach-row-open').focus();
  await shoot(page, 'a9-keyboard', 'strip-focus', width);
  await page.keyboard.press('Tab');
  await shoot(page, 'a9-keyboard', 'premium-link-focus', width);
  const premium = await page.context().newPage();
  await signInAs(premium, 'premium');
  await routeCoach(premium, STATES['ready-model']);
  await openAt(premium, ROWS_URL, width);
  await premium.getByTestId('coach-row-line').focus();
  await premium.keyboard.press('Enter');
  await settle(premium);
  await shoot(premium, 'a9-keyboard', 'heading-focus', width);
  await premium.getByTestId('coach-chips').first().locator('a').first().focus();
  await shoot(premium, 'a9-keyboard', 'chip-focus', width);
});

// ---- A11: Discord, rendered from the real handlers' payloads ------------------------------------

/** The active build's icons answer for a build the dump names that the repo does not hold. */
const ACTIVE_BUILD = (
  JSON.parse(readFileSync(path.join(WEB, 'src/data/active-build.json'), 'utf8')) as {
    build: string;
  }
).build;

function localFile(url: URL): string {
  const icon = /^\/data\/([^/]+)\/(icons\/.+)$/.exec(url.pathname);
  if (icon === null) return path.join(WEB, 'public', url.pathname);
  const build = existsSync(path.join(BUILDS, icon[1]!)) ? icon[1]! : ACTIVE_BUILD;
  return path.join(BUILDS, build, icon[2]!);
}

const CLIENTS = [
  { name: 'desktop', width: 640 },
  { name: 'phone', width: 390 },
];

const dumped: Dumped[] = existsSync(PAYLOADS) ? (JSON.parse(readFileSync(PAYLOADS, 'utf8')) as Dumped[]) : [];

test('A11 payloads.json exists (go test -run Dump with COACH_DUMP_PAYLOADS=1 writes it)', () => {
  test.skip(dumped.length === 0, 'payloads.json absent: dump it from the Go handlers first');
});

for (const entry of dumped) {
  for (const client of CLIENTS) {
    test(`A11 discord ${entry.id} @${client.name}`, async ({ page }) => {
      // Every embed image, whatever host the dump names (the site, or a loot fixture's test host).
      await page.route(/^https:\/\/[^/]+\/data\//, (route) => {
        const local = localFile(new URL(route.request().url()));
        return existsSync(local)
          ? route.fulfill({ path: local, contentType: 'image/webp' })
          : route.fulfill({ status: 404, body: '' });
      });
      await sizeTo(page, client.width, 600);
      await page.setContent(renderMessage(entry, { width: client.width, selectOpen: false }));
      await settle(page);
      await shoot(page, 'a11-discord', `${entry.id}-${client.name}`, client.width);
    });
  }
}
