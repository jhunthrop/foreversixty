// web/tests/captures/guild-raids/capture.spec.ts
// Acceptance screenshots for guild raids and sign-ups (design/specs/2026-10-26-guild-raids.md §9,
// the screenshot table, rows 1–16 and 18). Not a test of behaviour: it drives the Raids tab, the
// raid page, the form and the band line into each state, writes PNGs to
// design/reviews/2026-10-26-guild-raids/, and measures each surface against §7's height caps
// (annotations plus heights-<surface>-<width>.json there). Row 17 (the live Discord clients) is
// lane G1b's.
import { readFileSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Locator, type Page } from '@playwright/test';
import { readLayoutShift, trackLayoutShift } from '../../e2e/support/cls';
import { heldRoute } from '../../e2e/support/held-route';
import { gotoHydrated } from '../../e2e/support/hydrated';
import {
  GUILD_URL,
  MEMBER_ALT,
  NOW,
  ONYXIA_ID,
  ONYXIA_SIGNUPS,
  raidUrl,
  storedRaid,
  stripId,
  stubGuildRaids,
  threePlanned,
  type RaidStubOptions,
  type StoredRaid,
} from '../../e2e/support/guild-raids';
import { VIEWER_MEMBER } from '../../../src/fixtures/guild/mock-guild';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const OUT_DIR = path.resolve(HERE, '../../../../design/reviews/2026-10-26-guild-raids');
const MOCK_DIR = path.resolve(HERE, '../../../../design/mocks/renders');
const PAD = 24;
const TOP_CROP_PX = 640;
const VIEWPORT_HEIGHT = 1000;
const SETTLE_MS = 400;
const PHONE_MAX = 767;
const NEW_YORK = 'America/New_York';
const GUILD_ID = 2024;
const BASE_URL = `http://localhost:${process.env.CAPTURE_PORT ?? '4441'}`;
const AWAY = {
  baseURL: BASE_URL,
  timezoneId: NEW_YORK,
  deviceScaleFactor: 1,
  reducedMotion: 'reduce',
} as const;

const W360_390_1440 = [360, 390, 1440];
const PAIR = [390, 1440];
const ALL = [360, 390, 1440, 2000];

const target = (name: string): string => path.join(OUT_DIR, name);
const settle = (page: Page): Promise<void> => page.waitForTimeout(SETTLE_MS);
const nn = (n: number): string => String(n).padStart(2, '0');

// ---- data -----------------------------------------------------------------------------------

/** The member signed In on Zulmara, the alt on Bench (capture 5). */
function memberIn(): StoredRaid {
  const seeds = ONYXIA_SIGNUPS.map((s) => (s.name === MEMBER_ALT ? { ...s, state: 'bench' as const } : s));
  return storedRaid({
    seeds: [...seeds, { name: VIEWER_MEMBER, state: 'in', note: 'Have fire resist gear.' }],
  });
}

/** Neither of the member's characters signed up (capture 6). */
function memberUnsigned(): StoredRaid {
  return storedRaid({ seeds: ONYXIA_SIGNUPS.filter((s) => s.name !== MEMBER_ALT) });
}

/** Ten raids planned, so a repeat no longer fits (capture 13). */
function tenPlanned(): StoredRaid[] {
  return Array.from({ length: 10 }, (_, i) =>
    storedRaid({
      id: 20 + i,
      starts_at: new Date(Date.parse('2026-12-11T02:00:00Z') + i * 7 * 86_400_000).toISOString(),
      seeds: i === 0 ? ONYXIA_SIGNUPS : [],
    }),
  );
}

// ---- drivers --------------------------------------------------------------------------------

interface Open {
  width: number;
  stub: RaidStubOptions;
  at?: Date;
}

async function openRaid(page: Page, { width, stub, at = NOW }: Open, id = ONYXIA_ID): Promise<void> {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await stubGuildRaids(page, stub);
  await page.clock.install({ time: at });
  await gotoHydrated(page, raidUrl(id), 'raid-page');
  await settle(page);
}

async function openRaidsTab(page: Page, { width, stub, at = NOW }: Open): Promise<void> {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await stubGuildRaids(page, stub);
  await page.clock.install({ time: at });
  await gotoHydrated(page, `${GUILD_URL}#raids`, 'guild-raids-tab');
  await settle(page);
}

async function openOverview(page: Page, { width, stub, at = NOW }: Open): Promise<void> {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await stubGuildRaids(page, stub);
  await page.clock.install({ time: at });
  await gotoHydrated(page, GUILD_URL, 'guild-overview-standing');
  await page
    .locator('[data-testid="guild-raid-line"], [data-testid="guild-raid-line-none"]')
    .first()
    .waitFor();
  await settle(page);
}

/** The element's page rectangle grown by PAD, at most `maxHeight` tall, kept inside the page. */
async function cropOf(page: Page, element: Locator, file: string, maxHeight = Infinity): Promise<void> {
  const rect = await element.evaluate((node) => {
    const box = node.getBoundingClientRect();
    return {
      x: box.x + window.scrollX,
      y: box.y + window.scrollY,
      width: box.width,
      height: box.height,
      pageWidth: document.documentElement.clientWidth,
    };
  });
  const x = Math.max(0, rect.x - PAD);
  const y = Math.max(0, rect.y - PAD);
  await page.screenshot({
    path: target(file),
    fullPage: true,
    clip: {
      x,
      y,
      width: Math.min(rect.width + PAD * 2, rect.pageWidth - x),
      height: Math.min(rect.height + PAD * 2, maxHeight),
    },
  });
}

/** A full-page capture and a crop of `subject`. */
async function shoot(
  page: Page,
  row: number,
  state: string,
  width: number,
  subject?: Locator,
  maxHeight = TOP_CROP_PX,
): Promise<void> {
  const stem = `guild-raids-${nn(row)}-${state}`;
  await page.screenshot({ path: target(`${stem}-${width}.png`), fullPage: true });
  if (subject !== undefined) await cropOf(page, subject, `${stem}-panel-${width}.png`, maxHeight);
}

function atWidths(title: string, widths: number[], body: (page: Page, width: number) => Promise<void>): void {
  for (const width of widths) {
    test(`${title} @${width}`, ({ page }) => body(page, width));
  }
}

const tab = (page: Page): Locator => page.getByTestId('guild-raids-tab');
/** The guild band: the page header that holds the standing block, not the site header. */
const band = (page: Page): Locator =>
  page.locator('header', { has: page.getByTestId('guild-overview-standing') }).first();
const raidBand = (page: Page): Locator => page.getByTestId('raid-page');

// ---- heights (§7) ---------------------------------------------------------------------------

interface Cap {
  min?: number;
  max: number;
}

const CAPS: Record<string, { phone: Cap; desktop: Cap }> = {
  'planned-row': { phone: { min: 96, max: 120 }, desktop: { min: 64, max: 72 } },
  'strip-row': { phone: { max: 196 }, desktop: { max: 68 } },
  'band-line': { phone: { max: 112 }, desktop: { max: 44 } },
  'roster-row': { phone: { max: 44 }, desktop: { max: 36 } },
  form: { phone: { max: 1180 }, desktop: { max: 760 } },
};

async function heightsOf(scope: Locator): Promise<number[]> {
  return scope.evaluateAll((nodes) =>
    nodes
      .map((node) => Math.round(node.getBoundingClientRect().height * 10) / 10)
      .filter((height) => height > 0),
  );
}

function record(surface: string, width: number, heights: number[], note = ''): void {
  const cap = CAPS[surface][width <= PHONE_MAX ? 'phone' : 'desktop'];
  const within = heights.every((h) => h <= cap.max && (cap.min === undefined || h >= cap.min));
  const result = { surface, width, heights, cap, within, note };
  test.info().annotations.push({
    type: 'height',
    description: `${surface} @${width}: ${heights.join(', ')} px, cap ${cap.min ?? 0}-${cap.max}${within ? '' : ' OVER'}`,
  });
  writeFileSync(target(`heights-${surface}-${width}.json`), `${JSON.stringify(result)}\n`);
}

// ---- 1-4: the Raids tab ---------------------------------------------------------------------

atWidths('1 raids tab officer', ALL, async (page, width) => {
  await openRaidsTab(page, { width, stub: { viewer: 'officer', raids: threePlanned() } });
  await shoot(page, 1, 'tab-officer', width, tab(page));
  record('planned-row', width, await heightsOf(page.locator('[data-testid^="planned-raid-"][href]')));
});

atWidths('2 raids tab member', PAIR, async (page, width) => {
  await openRaidsTab(page, { width, stub: { viewer: 'member', raids: threePlanned() } });
  await shoot(page, 2, 'tab-member', width, tab(page));
});

atWidths('3 raids tab empty planned', PAIR, async (page, width) => {
  for (const viewer of ['member', 'officer'] as const) {
    const on = viewer === 'member' ? page : await page.context().newPage();
    await openRaidsTab(on, { width, stub: { viewer, raids: [] } });
    await shoot(on, 3, `tab-empty-${viewer}`, width, tab(on));
  }
});

atWidths('4 raids tab read-only', [360, 1440], async (page, width) => {
  await openRaidsTab(page, {
    width,
    stub: { viewer: 'officer', officerTools: false, raids: threePlanned() },
  });
  await shoot(page, 4, 'tab-read-only', width, tab(page));
});

// ---- 5-11: the raid page --------------------------------------------------------------------

atWidths('5 raid page member in', ALL, async (page, width) => {
  await openRaid(page, { width, stub: { viewer: 'member', raids: [memberIn()] } });
  await shoot(page, 5, 'page-member-in', width, raidBand(page));
  await cropOf(
    page,
    page.getByTestId('raid-signup-strip'),
    `guild-raids-05-page-member-in-strip-${width}.png`,
  );
  // §3.3: the phone cap (196) counts the whole row with its status line; the desktop 68 is the
  // grid row itself, so the note names both.
  const rows = page.locator('li[data-testid^="signup-row-"]');
  const grids = await heightsOf(rows.locator('> div').first());
  const whole = await heightsOf(rows);
  record(
    'strip-row',
    width,
    width <= PHONE_MAX ? whole : grids,
    `whole rows with status line: ${whole.join(', ')} px; grid rows: ${grids.join(', ')} px`,
  );
  const bare = page.locator('[data-testid="raid-roster-row"]:not(:has(p.break-words))');
  record('roster-row', width, [...new Set(await heightsOf(bare))]);
});

// The viewer in New York, the guild on Chicago: the band's when line carries the guild-time label.
atWidths('5 raid page member in, viewer in New York', PAIR, async (page, width) => {
  const ny = await page.context().browser()!.newContext(AWAY);
  const away = await ny.newPage();
  await openRaid(away, { width, stub: { viewer: 'member', raids: [memberIn()] } });
  await shoot(away, 5, 'page-member-in-new-york', width, raidBand(away));
  await ny.close();
});

atWidths('6 raid page member unsigned', W360_390_1440, async (page, width) => {
  await openRaid(page, { width, stub: { viewer: 'member', raids: [memberUnsigned()] } });
  await shoot(page, 6, 'page-member-unsigned', width, raidBand(page));
});

atWidths('7 raid page locked and the race line', W360_390_1440, async (page, width) => {
  await openRaid(page, { width, stub: { viewer: 'member', raids: [storedRaid({ status: 'locked' })] } });
  await shoot(page, 7, 'page-locked', width, page.getByTestId('raid-signup-strip'));
  const race = await page.context().newPage();
  const stub = await stubGuildRaids(race, { viewer: 'member' });
  await race.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await race.clock.install({ time: NOW });
  await gotoHydrated(race, raidUrl(ONYXIA_ID), 'raid-page');
  stub.refuseNext = { signup: { code: 'raid_locked', status: 409 } };
  await race.getByTestId(`${stripId('segments', VIEWER_MEMBER)}-in`).click();
  await race.getByText('This raid was locked while you were signing up.').waitFor();
  await settle(race);
  await shoot(race, 7, 'page-lock-race', width, race.getByTestId('raid-signup-strip'));
});

atWidths('8 raid page cancelled, ended, done', W360_390_1440, async (page, width) => {
  const states: [string, Open][] = [
    ['cancelled', { width, stub: { viewer: 'member', raids: [storedRaid({ status: 'cancelled' })] } }],
    ['ended', { width, stub: { viewer: 'member' }, at: new Date('2026-12-11T06:00:00Z') }],
    [
      'done',
      {
        width,
        stub: { viewer: 'member', raids: [storedRaid({ status: 'done' })] },
        at: new Date('2026-12-11T06:00:00Z'),
      },
    ],
  ];
  for (const [index, [state, open]] of states.entries()) {
    const on = index === 0 ? page : await page.context().newPage();
    await openRaid(on, open);
    await shoot(on, 8, `page-${state}`, width, raidBand(on));
  }
});

atWidths('9 raid page officer', [390, 1440, 2000], async (page, width) => {
  await openRaid(page, { width, stub: { viewer: 'officer' } });
  await shoot(page, 9, 'page-officer', width, raidBand(page));
  await cropOf(page, page.getByTestId('raid-roster'), `guild-raids-09-page-officer-roster-${width}.png`);
  if (width === 1440) {
    await page.setViewportSize({ width: 1024, height: VIEWPORT_HEIGHT });
    await settle(page);
    await cropOf(page, page.getByTestId('raid-roster'), `guild-raids-09-page-officer-roster-1024.png`);
    await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
    await settle(page);
  }
  await page.getByTestId('raid-officer-cancel').click();
  await page.getByTestId('raid-cancel-keep').waitFor();
  await settle(page);
  await cropOf(
    page,
    page.locator('[data-testid="raid-page"] > *').first(),
    `guild-raids-09-page-officer-confirm-${width}.png`,
  );
});

atWidths('10 raid page read-only', PAIR, async (page, width) => {
  await openRaid(page, { width, stub: { viewer: 'officer', officerTools: false } });
  await shoot(page, 10, 'page-read-only', width, raidBand(page));
});

atWidths('11 raid page visitor and unverified', W360_390_1440, async (page, width) => {
  for (const viewer of ['visitor', 'unverified'] as const) {
    const on = viewer === 'visitor' ? page : await page.context().newPage();
    await on.setViewportSize({ width, height: VIEWPORT_HEIGHT });
    await stubGuildRaids(on, { viewer });
    await gotoHydrated(on, raidUrl(ONYXIA_ID), 'raid-not-visible');
    await settle(on);
    await shoot(on, 11, `page-${viewer}`, width, on.getByTestId('raid-not-visible'));
  }
});

// ---- 12: loading reserves beside ready, CLS printed ----------------------------------------

atWidths('12 loading beside ready', PAIR, async (page, width) => {
  const surfaces = [
    {
      name: 'page',
      url: raidUrl(ONYXIA_ID),
      hold: `**/v1/guilds/${GUILD_ID}/raids/${ONYXIA_ID}`,
      ready: 'raid-page',
      loading: 'raid-page-loading',
    },
    {
      name: 'tab',
      url: `${GUILD_URL}#raids`,
      hold: `**/v1/guilds/${GUILD_ID}/raids`,
      ready: 'guild-planned',
      loading: 'guild-planned-reserve',
    },
  ];
  const cls: Record<string, number> = {};
  for (const [index, s] of surfaces.entries()) {
    const on = index === 0 ? page : await page.context().newPage();
    await on.setViewportSize({ width, height: VIEWPORT_HEIGHT });
    await stubGuildRaids(on, { viewer: 'member', raids: threePlanned() });
    await on.clock.install({ time: NOW });
    // A real warm visit first: the page itself stores the strip rows, the planned rows and the
    // guild band's member height that the measured visit then reserves. The read cache (`fs.q.*`)
    // is then dropped, so the measured visit loads every read cold behind those reserves.
    await gotoHydrated(on, s.url, s.ready);
    await settle(on);
    await on.evaluate(() => {
      for (const key of Object.keys(localStorage)) if (key.startsWith('fs.q.')) localStorage.removeItem(key);
    });
    const held = await heldRoute(on, s.hold, (route) => route.fallback());
    await trackLayoutShift(on);
    await on.reload();
    await held.started;
    await on.getByTestId(s.loading).waitFor();
    await settle(on);
    await on.screenshot({ path: target(`guild-raids-12-${s.name}-loading-${width}.png`), fullPage: true });
    held.release();
    await on.getByTestId(s.ready).waitFor();
    await settle(on);
    await on.screenshot({ path: target(`guild-raids-12-${s.name}-ready-${width}.png`), fullPage: true });
    cls[s.name] = Math.round((await readLayoutShift(on)) * 10_000) / 10_000;
  }
  test
    .info()
    .annotations.push({ type: 'cls', description: `@${width}: page ${cls.page}, tab ${cls.tab} (cap 0.02)` });
  writeFileSync(target(`cls-${width}.json`), `${JSON.stringify({ width, ...cls, cap: 0.02 })}\n`);
});

// ---- 13: the form ---------------------------------------------------------------------------

atWidths('13 form', W360_390_1440, async (page, width) => {
  await openRaidsTab(page, { width, stub: { viewer: 'officer', discordConnected: true } });
  await page.getByTestId('guild-raids-plan').click();
  const form = page.getByTestId('raid-form-create');
  await form.waitFor();
  await settle(page);
  await shoot(page, 13, 'form-default', width, form, Infinity);
  record('form', width, await heightsOf(form));

  await page.getByTestId('raid-form-date').fill('');
  await page.getByTestId('raid-form-submit').click();
  await page.getByTestId('raid-form-error-date').waitFor();
  await settle(page);
  await shoot(page, 13, 'form-errors', width, form, Infinity);

  const ny = await page.context().browser()!.newContext(AWAY);
  const away = await ny.newPage();
  await openRaidsTab(away, { width, stub: { viewer: 'officer', discordConnected: true } });
  await away.getByTestId('guild-raids-plan').click();
  await away.getByTestId('raid-form-for-you').waitFor();
  await settle(away);
  await shoot(away, 13, 'form-for-you', width, away.getByTestId('raid-form-create'), Infinity);
  await ny.close();

  const capped = await page.context().newPage();
  await openRaidsTab(capped, { width, stub: { viewer: 'officer', raids: tenPlanned() } });
  await capped.getByTestId('guild-raids-plan').click();
  await capped.getByTestId('raid-form-repeat-helper').waitFor();
  await settle(capped);
  await shoot(capped, 13, 'form-cap-no-webhook', width, capped.getByTestId('raid-form-create'), Infinity);
});

// ---- 14: the band line ----------------------------------------------------------------------

atWidths('14 band line', ALL, async (page, width) => {
  const states: [string, RaidStubOptions][] = [
    ['in', { viewer: 'member', raids: [memberIn()] }],
    ['unsigned', { viewer: 'member', raids: [memberUnsigned()] }],
    ['none-member', { viewer: 'member', raids: [] }],
    ['none-officer', { viewer: 'officer', raids: [] }],
  ];
  for (const [index, [state, stub]] of states.entries()) {
    const on = index === 0 ? page : await page.context().newPage();
    await openOverview(on, { width, stub });
    await shoot(on, 14, `band-${state}`, width, band(on));
    if (state === 'in') record('band-line', width, await heightsOf(on.getByTestId('guild-raid-line')));
  }
});

// ---- 15: focus-visible and hover ------------------------------------------------------------

async function focusShot(page: Page, control: Locator, crop: Locator, state: string, width: number) {
  await page.keyboard.press('Shift');
  await control.focus();
  await settle(page);
  await cropOf(page, crop, `guild-raids-15-focus-${state}-${width}.png`);
}

atWidths('15 focus and hover', PAIR, async (page, width) => {
  await openRaidsTab(page, { width, stub: { viewer: 'officer', raids: threePlanned() } });
  const row = page.locator('[data-testid^="planned-raid-"][href]').first();
  await focusShot(page, row, page.getByTestId('guild-planned'), 'planned-row', width);
  if (width === 1440) {
    await page.locator('body').click({ position: { x: 1, y: 1 } });
    await row.hover();
    await settle(page);
    await cropOf(page, page.getByTestId('guild-planned'), `guild-raids-15-hover-planned-row-${width}.png`);
  }
  await page.getByTestId('guild-raids-plan').click();
  await page.getByTestId('raid-form-create').waitFor();
  for (const field of ['zone', 'date', 'timezone', 'submit']) {
    await focusShot(
      page,
      page.getByTestId(`raid-form-${field}`),
      page.getByTestId('raid-form-create'),
      `form-${field}`,
      width,
    );
  }

  const raid = await page.context().newPage();
  await openRaid(raid, { width, stub: { viewer: 'officer' } });
  const strip = raid.getByTestId('raid-signup-strip');
  const officer = 'Kraggor';
  await focusShot(raid, raid.getByTestId(`${stripId('segments', officer)}-in`), strip, 'segment', width);
  await focusShot(raid, raid.getByTestId(stripId('role', officer)), strip, 'role', width);
  await focusShot(raid, raid.getByTestId(stripId('note', officer)), strip, 'note', width);
  await focusShot(
    raid,
    raid.getByTestId('raid-officer-lock'),
    raid.getByTestId('raid-officer-row'),
    'officer-button',
    width,
  );
  if (width === 1440) {
    await raid.locator('body').click({ position: { x: 1, y: 1 } });
    await raid.getByTestId('raid-officer-lock').hover();
    await settle(raid);
    await cropOf(
      raid,
      raid.getByTestId('raid-officer-row'),
      `guild-raids-15-hover-officer-button-${width}.png`,
    );
  }
});

// ---- 16: copy failed ------------------------------------------------------------------------

atWidths('16 copy failed', PAIR, async (page, width) => {
  await page.addInitScript(() => {
    Object.defineProperty(navigator, 'clipboard', {
      configurable: true,
      value: { writeText: () => Promise.reject(new Error('denied')) },
    });
  });
  await openRaid(page, { width, stub: { viewer: 'officer' } });
  await page.getByTestId('raid-officer-copy').click();
  await page.getByTestId('raid-copy-text').waitFor();
  await settle(page);
  await shoot(page, 16, 'copy-failed', width, page.getByTestId('raid-officer-row'));
});

// ---- 18: side by side with the mock boards --------------------------------------------------

async function sideBySide(
  page: Page,
  mock: string,
  built: string,
  file: string,
  width: number,
): Promise<void> {
  const data = (p: string): string => `data:image/png;base64,${readFileSync(p).toString('base64')}`;
  await page.setViewportSize({ width: width * 2 + 48, height: VIEWPORT_HEIGHT });
  await page.setContent(
    `<body style="margin:0;background:#000;color:#ccc;font:14px sans-serif;display:flex;gap:48px;align-items:flex-start">
      <figure style="margin:0;width:${width}px"><figcaption>Mock</figcaption><img src="${data(mock)}" style="width:${width}px"></figure>
      <figure style="margin:0;width:${width}px"><figcaption>Built (fixture)</figcaption><img src="${data(built)}" style="width:${width}px"></figure>
    </body>`,
  );
  await page.screenshot({ path: target(file), fullPage: true });
}

for (const width of [1024, 1440, 2000]) {
  test(`18 side by side @${width}`, async ({ browser }) => {
    const boards: [string, string, (page: Page) => Promise<void>][] = [
      [
        'guild-raids-tab',
        'tab',
        (p) => openRaidsTab(p, { width, stub: { viewer: 'officer', raids: threePlanned() } }),
      ],
      ['guild-raids-page', 'page', (p) => openRaid(p, { width, stub: { viewer: 'officer' } })],
    ];
    for (const [board, name, open] of boards) {
      // The mock board is drawn from Eastern time, so the built page is too.
      const context = await browser.newContext(AWAY);
      const mockPage = await context.newPage();
      await mockPage.setViewportSize({ width, height: VIEWPORT_HEIGHT });
      await mockPage.goto(`file://${path.join(MOCK_DIR, `${board}.html`)}`);
      await mockPage.waitForLoadState('networkidle');
      const mockPng = target(`guild-raids-18-mock-${name}-${width}.png`);
      await mockPage.screenshot({ path: mockPng, fullPage: true });
      const builtPage = await context.newPage();
      await open(builtPage);
      const builtPng = target(`guild-raids-18-built-${name}-${width}.png`);
      await builtPage.screenshot({ path: builtPng, fullPage: true });
      const composer = await context.newPage();
      await sideBySide(
        composer,
        mockPng,
        builtPng,
        `guild-raids-18-side-by-side-${name}-${width}.png`,
        width,
      );
      await context.close();
    }
  });
}
