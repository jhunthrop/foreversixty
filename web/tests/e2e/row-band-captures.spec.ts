// web/tests/e2e/row-band-captures.spec.ts
// The row-band sweep's review captures (design/specs/2026-11-01-row-bands.md): every page that
// draws a row list or table, full page, at the widths in ROW_CAPTURE_WIDTHS (default 1440, 2000
// and 390), written to ROW_CAPTURES (a directory) with ROW_CAPTURE_PREFIX (before/after). They are
// pictures for the reviewers, not assertions.
//   ROW_CAPTURES=../design/reviews/2026-11-01-row-bands ROW_CAPTURE_PREFIX=after \
//     npx playwright test row-band-captures.spec.ts --project=desktop
import { writeFileSync } from 'node:fs';
import { join } from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import {
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockGuildPage,
  buildMockHome,
  buildMockLoot,
  buildMockProgression,
  buildMockRaids,
  buildMockReadiness,
} from '../../src/fixtures/guild/mock-guild';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  GEAR,
  GEAR_ITEMS,
  routeGearIcons,
} from './support/character-fixture';
import { serveDuckdbRuntime } from './support/duckdb-runtime';
import { AFTER_END, DONE, stubGuildAttendance } from './support/guild-attendance';
import { CHICAGO, NOW, ONYXIA_ID, raidUrl, stubGuildRaids, threePlanned } from './support/guild-raids';
import { stubRankingsApi } from './support/rankings-fixture';
import { OBNOXIOUS, armoryPointer, signInWith, storePointer } from './support/selector';

const OUT = process.env.ROW_CAPTURES ?? '';
const PREFIX = process.env.ROW_CAPTURE_PREFIX ?? 'built';
const WIDTHS = (process.env.ROW_CAPTURE_WIDTHS ?? '1440,2000,390').split(',').map(Number);
const SETTLE_MS = 1200;

test.skip(OUT === '', 'review captures run on request');
test.skip(() => test.info().project.name !== 'desktop', 'one project draws every width');

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-capture' }),
  };
}

function guildMe(viewer: string, rank: 'leader' | 'member') {
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
    characters: [
      {
        key: `us/pvp/${viewer.toLowerCase()}`,
        region: 'us',
        ruleset: 'pvp',
        name: viewer,
        class: 'warrior',
      },
    ],
    guilds: [
      {
        id: GUILD_ID,
        region: 'us',
        ruleset: 'pvp',
        name: 'Olympus XXVII',
        rank,
        verified: true,
        can: rank === 'leader' ? ['approve', 'remove', 'settings', 'award', 'billing', 'release'] : [],
      },
    ],
  };
}

async function stubGuild(page: Page, rank: 'leader' | 'member'): Promise<void> {
  await page.route('**/v1/guilds/us/pvp/olympus-xxvii', (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  const viewer = rank === 'leader' ? VIEWER_OFFICER : VIEWER_MEMBER;
  await page.route('**/v1/me', (route) => route.fulfill(envelope(guildMe(viewer, rank))));
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
    route.fulfill(envelope(buildMockHome(viewer))),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/settings`, (route) => route.abort('connectionrefused'));
  await page.route(`**/v1/guilds/${GUILD_ID}/billing`, (route) => route.abort('connectionrefused'));
  await page.route(`**/v1/guilds/${GUILD_ID}/raids*`, (route) => route.fulfill(envelope(buildMockRaids())));
  await page.route(`**/v1/guilds/${GUILD_ID}/progression`, (route) =>
    route.fulfill(envelope(buildMockProgression())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/readiness`, (route) =>
    route.fulfill(envelope(buildMockReadiness())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/loot*`, (route) => route.fulfill(envelope(buildMockLoot())));
}

async function stubCharacter(page: Page): Promise<void> {
  await page.route(CHARACTER_URL, (route) => route.fulfill(envelope({ ...CHARACTER.data, ...GEAR })));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(envelope(null, 404)));
  await page.route('**/data/*/items/priest.json', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GEAR_ITEMS) }),
  );
  await routeGearIcons(page);
}

function reportRow(id: string, title: string, hoursAgo: number) {
  return {
    id,
    title,
    zone: 'Sanguine Depths',
    status: 'ready',
    visibility: 'public',
    created_at: new Date(Date.now() - hoursAgo * 3_600_000).toISOString(),
    fight_count: 12,
    kill_count: 2,
    boss_count: 2,
    logging_character: null,
  };
}

interface Shot {
  name: string;
  path: string;
  setup?: (page: Page) => Promise<unknown>;
  /** The clock the page reads (a raid page's times). */
  at?: Date;
  ready?: (page: Page) => Promise<void>;
  /** A state to show before the shot: a hovered row, an addressed fight. */
  act?: (page: Page) => Promise<void>;
  /** Test ids of the lists the sweep changed: each is shot on its own as `<shot>-<width>-<id>.png`. */
  focus?: string[];
}

/** Points at the first row matching `selector`, so its hover band is in the shot. */
const hoverFirst =
  (selector: string) =>
  async (page: Page): Promise<void> => {
    await page.locator(selector).first().hover();
  };

const GUILD_TABS = ['overview', 'roster', 'raids', 'progression', 'readiness', 'loot', 'settings'] as const;
const GUILD_TAB_FOCUS: Partial<Record<(typeof GUILD_TABS)[number], string[]>> = {
  roster: ['guild-roster-card'],
  readiness: ['guild-readiness-rows'],
};

const SHOTS: Shot[] = [
  ...GUILD_TABS.map((tab): Shot => ({
    name: `guild-${tab}`,
    path: `/guild/us/pvp/olympus-xxvii#${tab}`,
    setup: (page) => stubGuild(page, tab === 'settings' ? 'leader' : 'member'),
    ready: async (page) => {
      await expect(page.getByTestId(`guild-tab-${tab}`)).toBeVisible();
    },
    focus: GUILD_TAB_FOCUS[tab],
  })),
  {
    name: 'guild-raids-planned',
    path: '/guild/us/pvp/olympus-xxvii#raids',
    at: NOW,
    setup: (page) => stubGuildRaids(page, { viewer: 'member', raids: threePlanned() }),
    ready: async (page) => {
      await expect(page.getByTestId('guild-raids-tab')).toBeVisible();
    },
    act: hoverFirst('[data-testid="guild-planned-rows"] a'),
    focus: ['guild-planned-rows'],
  },
  {
    name: 'raid-planned',
    path: raidUrl(ONYXIA_ID),
    at: NOW,
    setup: (page) => stubGuildRaids(page, { viewer: 'member' }),
    ready: async (page) => {
      await expect(page.getByTestId('raid-page')).toBeVisible();
    },
    focus: ['raid-roster'],
  },
  {
    name: 'raid-report',
    path: raidUrl(ONYXIA_ID),
    at: AFTER_END,
    setup: (page) => stubGuildAttendance(page, { viewer: 'officer', nights: DONE }),
    ready: async (page) => {
      await expect(page.getByTestId('raid-report')).toBeVisible();
    },
    act: hoverFirst('[data-testid="raid-kill-row"]'),
    focus: ['raid-kill-order', 'raid-deaths', 'raid-loot', 'raid-consumables', 'raid-attendance'],
  },
  {
    name: 'rankings',
    path: '/rankings',
    setup: async (page) => {
      await signInWith(page, [OBNOXIOUS]);
      await storePointer(page, armoryPointer(OBNOXIOUS));
      await stubRankingsApi(page);
    },
    ready: async (page) => {
      await expect(page.getByTestId('ranking-1')).toBeVisible();
    },
  },
  {
    name: 'character',
    path: CHARACTER_PATH,
    setup: stubCharacter,
    ready: async (page) => {
      await expect(page.getByTestId('forever-score-card')).toBeVisible();
    },
    act: async (page) => {
      const id = await page.locator('li[id^="fight-"]').first().getAttribute('id');
      if (id !== null) await page.evaluate((hash) => (window.location.hash = hash), id);
    },
    focus: ['character-history'],
  },
  {
    name: 'report',
    path: '/reports/fixture2abcd?fight=1&tab=damage-done',
    setup: serveDuckdbRuntime,
  },
  {
    name: 'report-overview',
    path: '/reports/fixture2abcd',
    setup: serveDuckdbRuntime,
  },
  { name: 'logs', path: '/logs' },
  {
    name: 'reports',
    path: '/reports',
    setup: async (page) => {
      await page.route('**/v1/me', (route) => route.fulfill(envelope(guildMe(VIEWER_MEMBER, 'member'))));
      await page.route('**/v1/reports?mine=1**', (route) =>
        route.fulfill(
          envelope({
            rows: [
              reportRow('r1', 'Molten Core', 2),
              reportRow('r2', 'Onyxia', 30),
              reportRow('r3', 'Zul', 80),
            ],
            total: 3,
            page: 1,
            per_page: 100,
          }),
        ),
      );
    },
  },
  { name: 'home', path: '/' },
  { name: 'bis', path: '/bis/warrior/fury' },
  { name: 'bis-landing', path: '/bis' },
  { name: 'tiers', path: '/tiers' },
  { name: 'guides', path: '/guides/warrior/fury' },
  { name: 'premium', path: '/premium' },
  { name: 'changelog', path: '/changelog' },
  {
    name: 'account',
    path: '/account',
    setup: async (page) => {
      await page.route('**/v1/me', (route) => route.fulfill(envelope(guildMe(VIEWER_MEMBER, 'member'))));
    },
  },
];

for (const shot of SHOTS) {
  test(`capture ${shot.name}`, async ({ browser }) => {
    for (const width of WIDTHS) {
      const context = await browser.newContext({
        viewport: { width, height: 900 },
        deviceScaleFactor: 1,
        timezoneId: CHICAGO,
      });
      const page = await context.newPage();
      await shot.setup?.(page);
      if (shot.at !== undefined) await page.clock.install({ time: shot.at });
      await page.goto(shot.path);
      await shot.ready?.(page);
      await page.waitForLoadState('networkidle').catch(() => undefined);
      await page.waitForTimeout(SETTLE_MS);
      await shot.act?.(page);
      await page.screenshot({ path: join(OUT, `${PREFIX}-${shot.name}-${width}.png`), fullPage: true });
      // Where each changed list sits on the full page, so the before and after shots crop alike.
      const boxes: Record<string, unknown> = {};
      for (const id of shot.focus ?? []) {
        boxes[id] = await page
          .getByTestId(id)
          .first()
          .evaluate((el) => {
            const r = el.getBoundingClientRect();
            return {
              x: r.left + window.scrollX,
              y: r.top + window.scrollY,
              width: r.width,
              height: r.height,
            };
          });
      }
      if (shot.focus !== undefined) {
        writeFileSync(join(OUT, `${PREFIX}-${shot.name}-${width}.focus.json`), JSON.stringify(boxes));
      }
      if (process.env.ROW_AUDIT === '1') {
        const bands = await audit(page);
        writeFileSync(
          join(OUT, `${PREFIX}-${shot.name}-${width}.audit.json`),
          JSON.stringify(bands, null, 1),
        );
      }
      await context.close();
    }
  });
}

/** Every row-like box that paints a band (now or on hover) inside a card, inset from the card's
 *  edges, with its geometry: an inset band needs its own padding and radius. */
function audit(page: Page): Promise<unknown[]> {
  return page.evaluate(() => {
    const found: unknown[] = [];
    for (const el of Array.from(document.querySelectorAll<HTMLElement>('body *'))) {
      const cs = getComputedStyle(el);
      const tinted = cs.backgroundColor.startsWith('rgba(') && cs.backgroundColor !== 'rgba(0, 0, 0, 0)';
      const cls = typeof el.className === 'string' ? el.className : '';
      const hover = /hover:bg-|bg-row-hover|hover:bg/.test(cls);
      if (!tinted && !hover) continue;
      const r = el.getBoundingClientRect();
      if (r.width < 200 || r.height < 24 || r.height > 400) continue;
      const card = el.parentElement?.closest('.card-surface, .card-surface-quiet, [class*="rounded-panel"]');
      if (!card) continue;
      const c = card.getBoundingClientRect();
      const inset = Math.round(r.left - c.left);
      const insetRight = Math.round(c.right - r.right);
      if (inset <= 2 && insetRight <= 2) continue;
      const kids = Array.from(el.querySelectorAll('*'))
        .map((k) => k.getBoundingClientRect())
        .filter((k) => k.width > 0);
      const minLeft = Math.min(...kids.map((k) => k.left), r.right);
      found.push({
        tag: el.tagName,
        testid: el.dataset.testid ?? null,
        cls: cls.slice(0, 160),
        bg: cs.backgroundColor,
        hover,
        inset,
        insetRight,
        padL: cs.paddingLeft,
        padR: cs.paddingRight,
        radius: cs.borderTopLeftRadius,
        contentGapLeft: Math.round(minLeft - r.left),
        y: Math.round(r.top + window.scrollY),
      });
    }
    return found;
  });
}
