// web/tests/captures/guild-attendance/capture.spec.ts
// Acceptance screenshots for attendance and the raid-night report
// (design/specs/2026-10-27-guild-attendance.md §9, the screenshot table, rows 1–11 and 13). Not a
// test of behaviour: it drives the raid page, the picker, the roster, the readiness board and the
// account's guild line into each state and writes PNGs to design/reviews/2026-10-27-guild-attendance/.
// Row 11 measures CLS on a done raid with a warm reserve, and the poll's announced done swap on its
// own (cls-<width>.json). Row 12 (the live Discord clients) needs a staging channel.
import { readFileSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Locator, type Page } from '@playwright/test';
import { readLayoutShift, trackLayoutShift } from '../../e2e/support/cls';
import {
  AFTER_END,
  AFTER_WINDOW,
  DONE,
  bareReport,
  fixtureCandidates,
  memberNoShow,
  rosterAttendance,
  storedNight,
  stubGuildAttendance,
  type AttendanceStub,
  type AttendanceStubOptions,
} from '../../e2e/support/guild-attendance';
import { GUILD_URL, NOW, ONYXIA_ID, raidUrl } from '../../e2e/support/guild-raids';
import { heldRoute } from '../../e2e/support/held-route';
import { gotoHydrated } from '../../e2e/support/hydrated';
import { GUILD, GUILD_ID } from '../../../src/fixtures/guild/mock-guild';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const OUT_DIR = path.resolve(HERE, '../../../../design/reviews/2026-10-27-guild-attendance');
const MOCK_DIR = path.resolve(HERE, '../../../../design/mocks/renders');
const PAD = 24;
const TOP_CROP_PX = 900;
const VIEWPORT_HEIGHT = 1000;
const SETTLE_MS = 400;
const POLL_MS = 10_000;
const NEW_YORK = 'America/New_York';
const BASE_URL = `http://localhost:${process.env.CAPTURE_PORT ?? '4451'}`;
const AWAY = {
  baseURL: BASE_URL,
  timezoneId: NEW_YORK,
  deviceScaleFactor: 1,
  reducedMotion: 'reduce',
} as const;

const W3 = [360, 390, 1440];
const PAIR = [390, 1440];
const ALL = [360, 390, 1440, 2000];

const target = (name: string): string => path.join(OUT_DIR, name);
const settle = (page: Page): Promise<void> => page.waitForTimeout(SETTLE_MS);
const nn = (n: number): string => String(n).padStart(2, '0');

// ---- drivers --------------------------------------------------------------------------------

async function openRaid(
  page: Page,
  width: number,
  stub: AttendanceStubOptions,
  at = AFTER_END,
): Promise<AttendanceStub> {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  const made = await stubGuildAttendance(page, stub);
  await page.clock.install({ time: at });
  await gotoHydrated(page, raidUrl(ONYXIA_ID), 'raid-page');
  await settle(page);
  return made;
}

async function openTab(page: Page, width: number, stub: AttendanceStubOptions, hash: string, ready: string) {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  const made = await stubGuildAttendance(page, stub);
  await page.clock.install({ time: NOW });
  await gotoHydrated(page, `${GUILD_URL}#${hash}`, ready);
  await settle(page);
  return made;
}

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

/** A full-page capture and, with a subject, a crop of it: `guild-attendance-<nn>-<state>-<width>`. */
async function shoot(
  page: Page,
  row: number,
  state: string,
  width: number,
  subject?: Locator,
): Promise<void> {
  const stem = `guild-attendance-${nn(row)}-${state}`;
  await page.screenshot({ path: target(`${stem}-${width}.png`), fullPage: true });
  if (subject !== undefined) await cropOf(page, subject, `${stem}-panel-${width}.png`, TOP_CROP_PX);
}

function atWidths(title: string, widths: number[], body: (page: Page, width: number) => Promise<void>): void {
  for (const width of widths) {
    test(`${title} @${width}`, ({ page }) => body(page, width));
  }
}

const band = (page: Page): Locator => page.locator('[data-testid="raid-page"] > *').nth(1);
const fresh = (page: Page): Promise<Page> => page.context().newPage();

// ---- 1-3: the done and ended raid pages -------------------------------------------------------

atWidths('1 done raid, officer', ALL, async (page, width) => {
  await openRaid(page, width, { viewer: 'officer', nights: DONE, discordConnected: true });
  await shoot(page, 1, 'done-officer', width, band(page));
  await cropOf(page, page.getByTestId('raid-report'), `guild-attendance-01-done-officer-report-${width}.png`);
});

atWidths('2 done raid, member', W3, async (page, width) => {
  await openRaid(page, width, { viewer: 'member', nights: DONE });
  await shoot(page, 2, 'member-present', width, page.getByTestId('raid-your-night'));
  const noShow = await fresh(page);
  await openRaid(noShow, width, { viewer: 'member', nights: memberNoShow() });
  await shoot(noShow, 2, 'member-no-show', width, noShow.getByTestId('raid-your-night'));
});

atWidths('3 ended, no log', W3, async (page, width) => {
  const states: [string, AttendanceStubOptions, Date][] = [
    ['officer-window', { viewer: 'officer' }, AFTER_END],
    ['officer-after', { viewer: 'officer' }, AFTER_WINDOW],
    ['member-window', { viewer: 'member' }, AFTER_END],
    ['member-after', { viewer: 'member' }, AFTER_WINDOW],
  ];
  for (const [index, [state, stub, at]] of states.entries()) {
    const on = index === 0 ? page : await fresh(page);
    await openRaid(on, width, stub, at);
    await shoot(on, 3, `ended-${state}`, width, band(on));
  }
});

// ---- 4-5: the picker and the clear confirm -----------------------------------------------------

atWidths('4 picker', W3, async (page, width) => {
  const states: [string, AttendanceStubOptions['candidates']][] = [
    ['closest', fixtureCandidates()],
    ['empty', { zone_known: true, rows: [] }],
    ['unknown-zone', fixtureCandidates({ zone_known: false })],
  ];
  for (const [index, [state, candidates]] of states.entries()) {
    const on = index === 0 ? page : await fresh(page);
    await openRaid(on, width, { viewer: 'officer', candidates });
    await on.getByTestId('raid-match-open').click();
    await on.locator('[data-testid="raid-match-rows"], [data-testid="raid-match-empty"]').first().waitFor();
    await settle(on);
    await shoot(on, 4, `picker-${state}`, width, on.getByTestId('raid-match-picker'));
  }
});

atWidths('5 clear confirm', PAIR, async (page, width) => {
  await openRaid(page, width, { viewer: 'officer', nights: DONE });
  await page.getByTestId('raid-not-this-log').focus();
  await page.keyboard.press('Enter');
  await page.getByTestId('raid-clear-keep').waitFor();
  await settle(page);
  await shoot(page, 5, 'clear-confirm', width, page.getByTestId('raid-match-line'));
});

// ---- 6-8: attendance, the bare report, read-only -----------------------------------------------

atWidths('6 attendance', ALL, async (page, width) => {
  await openRaid(page, width, { viewer: 'officer', nights: DONE });
  await page.getByLabel('Excused, Mendwhisper').check();
  await page.getByTestId('raid-attendance-summary').getByText('1 excused').waitFor();
  await settle(page);
  await shoot(page, 6, 'attendance', width, page.getByTestId('raid-attendance'));
});

atWidths('7 bare report', PAIR, async (page, width) => {
  await openRaid(page, width, {
    viewer: 'officer',
    nights: { [ONYXIA_ID]: storedNight({ report: bareReport() }) },
  });
  await shoot(page, 7, 'bare-report', width, page.getByTestId('raid-report'));
});

atWidths('8 read-only', PAIR, async (page, width) => {
  await openRaid(page, width, {
    viewer: 'officer',
    officerTools: false,
    nights: DONE,
    discordConnected: true,
  });
  await shoot(page, 8, 'read-only', width, band(page));
  await cropOf(
    page,
    page.getByTestId('raid-attendance'),
    `guild-attendance-08-read-only-attendance-${width}.png`,
    TOP_CROP_PX,
  );
});

// ---- 8b: Notes and Sign-ups opened on a done raid, read-only (WP 6) --------------------------

atWidths('8b notes and sign-ups opened', PAIR, async (page, width) => {
  await openRaid(page, width, { viewer: 'officer', nights: DONE });
  await page.getByTestId('raid-notes-disclosure-summary').click();
  await page.getByTestId('raid-roster-disclosure-summary').click();
  await page.getByTestId('raid-roster').waitFor();
  await settle(page);
  await shoot(page, 8, 'sections-open', width, page.getByTestId('raid-notes-disclosure'));
  await cropOf(
    page,
    page.getByTestId('raid-roster-disclosure'),
    `guild-attendance-08-sections-open-roster-${width}.png`,
    TOP_CROP_PX,
  );
});

// ---- N1: the guild views with no band keep their top spacing (UX N1) ----------------------------

/** The fixture build's prerendered guild shells (lib/report/shell-paths.ts, which reads the site
 *  config and so is not imported here). */
const FIXTURE_GUILD = '/guild/us/hardcore/the-last-watch';
const FIXTURE_INVITE = '/guild/invite/fixtureinvitetoken1';

const BANDLESS: [name: string, url: string, ready: string][] = [
  ['landing', '/guild', 'guild-landing'],
  ['claim', `${FIXTURE_GUILD}/claim`, 'guild-view'],
  ['settings', `${FIXTURE_GUILD}/settings`, 'guild-view'],
  ['join-bad-token', FIXTURE_INVITE, 'guild-view'],
  ['not-an-address', '/guild/not-a-guild-address', 'guild-missing'],
];

atWidths('N1 band-less guild views', PAIR, async (page, width) => {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  // Every route answers as a signed-in officer with no claim or invite to show: the views draw
  // their refusals, which is what has no band.
  await page.route('**/v1/**', (route) =>
    route.fulfill({
      status: 404,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":{"code":"not_found","message":"Not found."}}',
    }),
  );
  // The fixture build prerenders no unknown guild address; the Worker serves the guild shell.
  await page.route('**/guild/not-a-guild-address', async (route) => {
    const shell = new URL(FIXTURE_GUILD, route.request().url()).toString();
    return route.fulfill({ response: await route.fetch({ url: shell }) });
  });
  for (const [name, url, ready] of BANDLESS) {
    await page.goto(url);
    await page.getByTestId(ready).first().waitFor();
    await settle(page);
    await page.screenshot({ path: target(`guild-attendance-N1-${name}-${width}.png`) });
  }
});

// ---- 9-10: the member record and readiness ------------------------------------------------------

/** `/v1/me` for the account page: the member's two characters with their guild line in raid mode. */
function accountMe() {
  const guild = { id: GUILD_ID, name: GUILD.name, rank: 'member', verified: true };
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
    characters: [
      // The same read as the roster chip (WP 4): rosterAttendance's last 8.
      ['Zulmara', 'hunter'],
      ['Duskstrider', 'rogue'],
    ].map(([name, cls]) => ({
      key: `us/pvp/${String(name).toLowerCase()}`,
      region: 'us',
      ruleset: 'pvp',
      name,
      class: cls,
      guild: { ...guild, attendance: rosterAttendance(name).last8 },
    })),
    guilds: [{ ...guild, region: GUILD.region, ruleset: GUILD.ruleset, consent: 'gear' }],
  };
}

atWidths('9 roster chip and guild line', W3, async (page, width) => {
  for (const [index, source] of (['raids', 'reports'] as const).entries()) {
    const on = index === 0 ? page : await fresh(page);
    await openTab(on, width, { viewer: 'member', attendanceSource: source }, 'roster', 'guild-roster-tab');
    await shoot(on, 9, `roster-${source}`, width, on.getByTestId('guild-roster-tab'));
  }
  const account = await fresh(page);
  await account.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await account.route('**/v1/me', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: accountMe(), error: null }),
    }),
  );
  await account.route('**/v1/devices', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: '{"ok":true,"data":[],"error":null}',
    }),
  );
  await account.goto('/account');
  await account.getByTestId('character-guild-came').first().waitFor();
  await settle(account);
  await shoot(
    account,
    9,
    'guild-line',
    width,
    account.getByTestId('character-guild-line').first().locator('xpath=../..'),
  );
});

atWidths('10 readiness next raid', W3, async (page, width) => {
  await openTab(page, width, { viewer: 'officer' }, 'readiness', 'guild-readiness-next-raid');
  await shoot(page, 10, 'readiness-on', width, page.getByTestId('guild-readiness-tab'));
  await page.getByTestId('guild-readiness-next-raid').click();
  await settle(page);
  await shoot(page, 10, 'readiness-off', width, page.getByTestId('guild-readiness-tab'));
});

// ---- 11: loading beside ready, CLS -------------------------------------------------------------

/** Loading beside ready (UX 8): the y of the tabs, and of the first region under the band. */
const LOADING_EDGES = {
  tabsTop: '[data-testid="raid-page-loading"] [data-testid="guild-tabs"]',
  bandBottom: '[data-testid="raid-page-loading"] > :nth-child(2)',
  yourNightBottom: '[data-testid="raid-your-night-reserve"]',
  reportTop: '[data-testid="raid-report-reserve"]',
};
const READY_EDGES = {
  tabsTop: '[data-testid="raid-page"] [data-testid="guild-tabs"]',
  bandBottom: '[data-testid="raid-page"] > :nth-child(2)',
  yourNightBottom: '[data-testid="raid-your-night"]',
  reportTop: '[data-testid="raid-report"]',
};
type Edges = Record<keyof typeof LOADING_EDGES, number | null>;

async function edgesOf(page: Page, selectors: typeof LOADING_EDGES): Promise<Edges> {
  return page.evaluate((sel) => {
    const y = (css: string, edge: 'top' | 'bottom'): number | null => {
      const node = document.querySelector(css);
      if (node === null) return null;
      const box = node.getBoundingClientRect();
      return Math.round((edge === 'top' ? box.top : box.bottom) + window.scrollY);
    };
    return {
      tabsTop: y(sel.tabsTop, 'top'),
      bandBottom: y(sel.bandBottom, 'bottom'),
      yourNightBottom: y(sel.yourNightBottom, 'bottom'),
      reportTop: y(sel.reportTop, 'top'),
    };
  }, selectors);
}

function deltas(loading: Edges, ready: Edges): Record<string, number | null> {
  const keys = Object.keys(loading) as (keyof Edges)[];
  return Object.fromEntries(
    keys.map((key) => {
      const a = loading[key];
      const b = ready[key];
      return [key, a === null || b === null ? null : b - a];
    }),
  );
}

async function warmCls(
  page: Page,
  width: number,
  stub: AttendanceStubOptions,
  viewer: string,
): Promise<number> {
  await openRaid(page, width, stub);
  // The warm visit stored the strip rows and the done flag; the read cache is dropped, so the
  // measured visit loads the raid cold behind those reserves.
  await page.evaluate(() => {
    for (const key of Object.keys(localStorage)) if (key.startsWith('fs.q.')) localStorage.removeItem(key);
  });
  const held = await heldRoute(page, `**/v1/guilds/${GUILD_ID}/raids/${ONYXIA_ID}`, (route) =>
    route.fallback(),
  );
  await trackLayoutShift(page);
  await page.reload();
  await held.started;
  await page.getByTestId('raid-page-loading').waitFor();
  await settle(page);
  await page.screenshot({
    path: target(`guild-attendance-11-loading-${viewer}-${width}.png`),
    fullPage: true,
  });
  const loading = await edgesOf(page, LOADING_EDGES);
  held.release();
  await page.getByTestId('raid-page').waitFor();
  await settle(page);
  const ready = await edgesOf(page, READY_EDGES);
  writeFileSync(
    target(`heights-reserve-${viewer}-${width}.json`),
    `${JSON.stringify({ viewer, width, loading, ready, delta: deltas(loading, ready) })}\n`,
  );
  await page.screenshot({ path: target(`guild-attendance-11-ready-${viewer}-${width}.png`), fullPage: true });
  return Math.round((await readLayoutShift(page)) * 10_000) / 10_000;
}

/** The poll's announced swap, measured on its own: from a settled ended page to the done one. */
async function pollCls(page: Page, width: number): Promise<number> {
  await trackLayoutShift(page);
  const stub = await openRaid(page, width, { viewer: 'officer' });
  await page.waitForTimeout(500);
  await page.evaluate(() => (window.__cls = 0));
  stub.matchFromElsewhere(ONYXIA_ID);
  await page.clock.runFor(POLL_MS);
  await page.getByTestId('raid-night-line').waitFor();
  await settle(page);
  return Math.round((await readLayoutShift(page)) * 10_000) / 10_000;
}

atWidths('11 loading beside ready', PAIR, async (page, width) => {
  const cls = {
    officer: await warmCls(page, width, { viewer: 'officer', nights: DONE }, 'officer'),
    member: await warmCls(await fresh(page), width, { viewer: 'member', nights: DONE }, 'member'),
    poll_swap: await pollCls(await fresh(page), width),
  };
  test.info().annotations.push({
    type: 'cls',
    description: `@${width}: officer ${cls.officer}, member ${cls.member}, poll swap ${cls.poll_swap} (cap 0.02 on the warm loads)`,
  });
  writeFileSync(target(`cls-${width}.json`), `${JSON.stringify({ width, ...cls, cap: 0.02 })}\n`);
});

// ---- 1b: the first screen against §1 (UX 6) ----------------------------------------------------

const FIRST_SCREENS = [
  {
    name: 'officer',
    width: 1440,
    height: 800,
    guide: 620,
    stub: { viewer: 'officer', nights: DONE } as const,
    subject: '[data-testid="raid-kill-row"]',
    edge: 'bottom' as const,
  },
  {
    name: 'member',
    width: 360,
    height: 640,
    guide: 620,
    stub: { viewer: 'member', nights: DONE } as const,
    subject: '[data-testid="raid-your-night-sentence"]',
    edge: 'top' as const,
  },
];

for (const screen of FIRST_SCREENS) {
  test(`1b first screen ${screen.name} @${screen.width}`, async ({ page }) => {
    await page.setViewportSize({ width: screen.width, height: screen.height });
    await stubGuildAttendance(page, screen.stub);
    await page.clock.install({ time: AFTER_END });
    await gotoHydrated(page, raidUrl(ONYXIA_ID), 'raid-page');
    await settle(page);
    const y = await page
      .locator(screen.subject)
      .first()
      .evaluate((node, edge) => {
        const box = node.getBoundingClientRect();
        return Math.round(edge === 'top' ? box.top : box.bottom);
      }, screen.edge);
    await page.evaluate((guide) => {
      const line = document.createElement('div');
      line.style.cssText = `position:fixed;left:0;right:0;top:${guide}px;height:1px;background:#ff00ff;z-index:9999`;
      document.body.append(line);
    }, screen.guide);
    await page.screenshot({
      path: target(`guild-attendance-01b-first-screen-${screen.name}-${screen.width}.png`),
    });
    writeFileSync(
      target(`first-screen-${screen.name}-${screen.width}.json`),
      `${JSON.stringify({ ...screen, stub: undefined, y, within: y <= screen.guide })}\n`,
    );
  });
}

// ---- 13: side by side with the mock boards ------------------------------------------------------

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

const BOARDS: [board: string, name: string, stub: AttendanceStubOptions, boardWidth?: number][] = [
  ['guild-attendance-page', 'officer', { viewer: 'officer', nights: DONE, discordConnected: true }],
  ['guild-attendance-page-phone', 'member', { viewer: 'member', nights: DONE }, 390],
];

for (const width of [1024, 1440, 2000]) {
  test(`13 side by side @${width}`, async ({ browser }) => {
    for (const [board, name, stub, boardWidth] of BOARDS) {
      const shotWidth = boardWidth ?? width;
      if (boardWidth !== undefined && width !== 1440) continue;
      // The mock boards are drawn from Eastern time, so the built page is too.
      const context = await browser.newContext(AWAY);
      const mockPage = await context.newPage();
      await mockPage.setViewportSize({ width: shotWidth, height: VIEWPORT_HEIGHT });
      const boardFile = shotWidth === 1024 ? `${board}-1024` : board;
      await mockPage.goto(`file://${path.join(MOCK_DIR, `${boardFile}.html`)}`);
      await mockPage.waitForLoadState('networkidle');
      const mockPng = target(`guild-attendance-13-mock-${name}-${shotWidth}.png`);
      await mockPage.screenshot({ path: mockPng, fullPage: true });
      const builtPage = await context.newPage();
      await openRaid(builtPage, shotWidth, stub);
      const builtPng = target(`guild-attendance-13-built-${name}-${shotWidth}.png`);
      await builtPage.screenshot({ path: builtPng, fullPage: true });
      await sideBySide(
        await context.newPage(),
        mockPng,
        builtPng,
        `guild-attendance-13-side-by-side-${name}-${shotWidth}.png`,
        shotWidth,
      );
      await context.close();
    }
  });
}
