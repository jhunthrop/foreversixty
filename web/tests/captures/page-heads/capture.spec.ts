// web/tests/captures/page-heads/capture.spec.ts
// The unified page head (design/reviews/2026-11-09-sim-head): every page's head in the states the
// parity gates open (their support openers, so the data is the reviewed fixtures), signed in with
// a realistic character where the page has one, plus the owner's two reported states: /sim with a
// browser run of a spec no parse has measured ("Go Hogs", Windshaper Skyborne Fury Warrior 35), and
// /bis as Felsnap (level 34 Beast Mastery Hunter). Each capture is the viewport (the head and what
// follows it) and a crop of the head band for the contact sheet. Not a test of behaviour.
import { existsSync, mkdirSync, readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test, type Page } from '@playwright/test';
import { prepareParityPage, settleParityPage } from '../../e2e/support/parity';
import { openSimState, openToolState, serveMockData } from '../../e2e/support/sim-parity';
import { openBisLanding } from '../../e2e/support/bis-landing-page';
import { openBis } from '../../e2e/support/bis-page';
import { openHome, type HomeState } from '../../e2e/support/home-page';
import { openCharacter, type CharacterState } from '../../e2e/support/character-page';
import { openGuildState } from '../../e2e/support/guild-parity';
import { GUILD_PARITY_NOW } from '../../e2e/support/guild-parity/clock';
import { openTiersState } from '../../e2e/support/tiers-parity';
import { openReportState } from '../../e2e/support/report-parity';
import { openLogsState } from '../../e2e/support/logs-parity';
import { openRankingsState } from '../../e2e/support/rankings-parity';
import { openGuidesState } from '../../e2e/support/guides-page';
import { openPlanner, PLANNER_CODES } from '../../e2e/support/planner-page';
import { openSetupState } from '../../e2e/support/setup-page';
import { openAccountState, renderAccountMocks } from '../../e2e/support/account-page';
import { routePagesState } from '../../e2e/support/pages-page';
import { signInWith } from '../../e2e/support/selector';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const WEB = path.resolve(HERE, '../../..');
const PHASE = process.env.CAPTURE_PHASE ?? 'after';
const OUT = path.resolve(WEB, '../design/reviews/2026-11-09-sim-head', PHASE);
const HEADS = path.join(OUT, 'heads');
mkdirSync(HEADS, { recursive: true });

/** The repository's build icon folders, newest first: the fixture build publishes only a few
 *  icons, so every icon a page asks for is served from the newest real build carrying it (the
 *  guild parity's own rule), and a capture never shows an empty square the live site would fill. */
const BUILDS = path.resolve(WEB, '../data/builds');
const ICON_DIRS: readonly string[] = existsSync(BUILDS)
  ? readdirSync(BUILDS)
      .sort()
      .reverse()
      .map((build) => path.join(BUILDS, build, 'icons'))
      .filter((dir) => existsSync(dir))
  : [];

async function routeIcons(page: Page): Promise<void> {
  await page.route('**/data/*/icons/*.webp', (route) => {
    const name = new URL(route.request().url()).pathname.split('/').pop() ?? '';
    const file = ICON_DIRS.map((dir) => path.join(dir, name)).find((candidate) => existsSync(candidate));
    return file === undefined
      ? route.continue()
      : route.fulfill({ status: 200, contentType: 'image/webp', body: readFileSync(file) });
  });
}

const parity = (file: string): { states: Record<string, unknown>[] } =>
  JSON.parse(readFileSync(path.resolve(WEB, '../design/mocks/web', file), 'utf8')) as {
    states: Record<string, unknown>[];
  };
const stateOf = <T>(file: string, id: string): T => {
  const found = parity(file).states.find((state) => state.id === id);
  if (found === undefined) throw new Error(`${file}: no state ${id}`);
  return found as T;
};

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'heads' }),
});

/** The owner's character from the report: a Horde Fury Warrior at 35. */
const GO_HOGS = {
  name: 'Go Hogs',
  class: 'Warrior',
  spec: 'Fury',
  level: 35,
  realm: 'Living Flame',
  faction: 'horde' as const,
  race: 'Windshaper Skyborne',
};
/** /sim signed in as Go Hogs, a browser run finished, the spec measured by no parse ("Not yet"). */
async function simBrowserNotYet(page: Page): Promise<void> {
  await signInWith(page, [GO_HOGS]);
  await page.route('**/v1/specs**', (route) => route.fulfill(envelope({ specs: [] })));
  await page.route('**/v1/sims?mine=1*', (route) =>
    route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/sims/server*', (route) => route.fulfill(envelope({ rows: [] })));
  await serveMockData(page);
  // The published Fury set from the landing rail, worn by Go Hogs: a Windshaper Skyborne at 35.
  await page.goto('/sim');
  const href = await page
    .locator('[data-region="section-sim-a-best-in-slot-set"] a', { hasText: 'Fury' })
    .getAttribute('href');
  const url = new URL(href ?? '', 'http://localhost');
  // Go Hogs at 35: the published Fury set worn, a level-35 Fury spend (24 points) and the level.
  const parts = (url.searchParams.get('code') ?? '').split(':');
  parts[3] = 'windshaper-skyborne';
  parts[4] = '0/5530515/0';
  const code = parts.join(':').replace('|level=60', '|level=35');
  url.searchParams.set('code', `${code}|who=${GO_HOGS.name}:Forever`);
  url.searchParams.set('source', 'addon');
  await page.goto(`${url.pathname}${url.search}`);
  await expect(page.getByTestId('sim-answer')).toHaveAttribute('data-state', 'result', { timeout: 30_000 });
  await expect(page.getByTestId('spec-fidelity-note')).toBeVisible({ timeout: 15_000 });
}

/** /sim's server run of the parity fixture, the spec validated. */
async function simServerCalibrated(page: Page): Promise<void> {
  await page.route('**/v1/specs**', (route) =>
    route.fulfill(
      envelope({
        specs: [
          {
            spec: 'warrior-fury',
            state: 'validated',
            median_gap: 0.021,
            parses: 50,
            worst_actions: [],
            engine_version: 'bf72769',
            updated_at: '2026-11-08T10:00:00Z',
          },
        ],
      }),
    ),
  );
  await openSimState(page, 'result');
}

interface Shot {
  id: string;
  open: (page: Page) => Promise<void>;
  widths?: readonly number[];
  /** Opened at 1440 and then narrowed: the state's opener drives desktop-only controls. */
  openWide?: boolean;
}

const ALL = [1440, 390] as const;
const WITH_WIDE = [1440, 2000, 390] as const;
const pagesIn =
  (pathname: string) =>
  async (page: Page): Promise<void> => {
    await routePagesState(page, 'in');
    await page.goto(pathname);
  };
const pagesOut =
  (pathname: string) =>
  async (page: Page): Promise<void> => {
    await routePagesState(page, 'out');
    await page.goto(pathname);
  };

const SHOTS: readonly Shot[] = [
  { id: 'sim-browser-notyet', open: simBrowserNotYet, widths: WITH_WIDE },
  { id: 'sim-server-calibrated', open: simServerCalibrated, widths: WITH_WIDE, openWide: true },
  { id: 'sim-landing-out', open: (page) => openSimState(page, 'landing-out'), widths: WITH_WIDE },
  { id: 'sim-landing-in', open: (page) => openSimState(page, 'landing-in') },
  { id: 'sim-saved', open: (page) => openSimState(page, 'saved') },
  { id: 'sim-gear', open: (page) => openToolState(page, 'gear') },
  { id: 'sim-specs', open: (page) => openToolState(page, 'specs') },
  { id: 'bis-landing-34', open: (page) => openBisLanding(page, '/bis', '34'), widths: WITH_WIDE },
  { id: 'bis-landing-out', open: (page) => openBisLanding(page, '/bis', 'out') },
  { id: 'bis-spec', open: (page) => openBis(page, '/bis/hunter/marksmanship/60') },
  { id: 'home-out', open: (page) => openHome(page, stateOf<HomeState>('home-parity.json', 'out')) },
  { id: 'home-34', open: (page) => openHome(page, stateOf<HomeState>('home-parity.json', '34')) },
  {
    id: 'character-owner-34',
    open: (page) => openCharacter(page, stateOf<CharacterState>('character-parity.json', 'owner-34')),
  },
  {
    id: 'character-visitor-60',
    open: (page) => openCharacter(page, stateOf<CharacterState>('character-parity.json', 'visitor-60')),
  },
  { id: 'guild-member', open: (page) => guild(page, 'member') },
  { id: 'guild-landing-in', open: (page) => guild(page, 'landing-in') },
  { id: 'guild-raid', open: (page) => guild(page, 'raid') },
  { id: 'guild-report', open: (page) => guild(page, 'report') },
  { id: 'guild-profile', open: (page) => guild(page, 'profile') },
  { id: 'guild-recruit', open: (page) => guild(page, 'recruit') },
  { id: 'guild-claim', open: (page) => guild(page, 'claim') },
  { id: 'guild-join', open: (page) => guild(page, 'join') },
  { id: 'guild-settings-page', open: (page) => guild(page, 'settings-page') },
  { id: 'tiers-dps-you', open: (page) => openTiersState(page, 'dps-you') },
  {
    id: 'report-summary',
    open: (page) => openReportState(page, 'report-summary', '/reports/sample?tab=summary'),
  },
  { id: 'logs-you', open: (page) => openLogsState(page, 'logs-you') },
  { id: 'reports-list', open: (page) => openLogsState(page, 'reports-list') },
  { id: 'rankings-you', open: (page) => openRankingsState(page, 'rankings-you') },
  { id: 'rankings-board', open: (page) => openRankingsState(page, 'rankings-board') },
  { id: 'guides-hunter34', open: (page) => openGuidesState(page, '/guides', 'hunter34') },
  { id: 'guides-class-warrior-out', open: (page) => openGuidesState(page, '/guides/warrior', 'out') },
  { id: 'guides-class-druid-out', open: (page) => openGuidesState(page, '/guides/druid', 'out') },
  { id: 'guides-class-hunter34', open: (page) => openGuidesState(page, '/guides/hunter', 'hunter34') },
  {
    id: 'guides-spec-hunter34',
    open: (page) => openGuidesState(page, '/guides/hunter/beast-mastery', 'hunter34'),
  },
  {
    id: 'planner-felsnap',
    open: (page) =>
      openPlanner(page, `/planner?code=${encodeURIComponent(PLANNER_CODES.survival)}`, 'felsnap'),
  },
  {
    id: 'account-in',
    // The account read settled (the parity fixture's account), never the loading reserve.
    open: async (page) => {
      renderAccountMocks(['web-account']);
      await openAccountState(page, 'account');
    },
  },
  { id: 'setup-done', open: (page) => openSetupState(page, 'done') },
  { id: 'premium-in', open: pagesIn('/premium') },
  { id: 'characters-in', open: pagesIn('/characters') },
  { id: 'changelog-in', open: pagesIn('/changelog') },
  { id: 'privacy-out', open: pagesOut('/privacy') },
  { id: 'connect-in', open: pagesIn('/connect') },
  { id: 'login-out', open: pagesOut('/login') },
  { id: 'not-found-in', open: pagesIn('/this-page-does-not-exist') },
];

async function guild(page: Page, id: string): Promise<void> {
  await page.clock.setFixedTime(GUILD_PARITY_NOW);
  await openGuildState(page, id);
}

/** The head band's crop: from the top of the page to 24 px under the head (or its tab row). */
async function headCrop(page: Page, file: string): Promise<void> {
  const head = page.locator('header[data-region="head"], [data-region="head"]').first();
  const box = (await head.count()) > 0 ? await head.boundingBox() : null;
  const bottom = box === null ? 320 : Math.ceil(box.y + box.height + 24);
  const width = page.viewportSize()?.width ?? 1440;
  await page.screenshot({ path: file, clip: { x: 0, y: 0, width, height: Math.max(bottom, 120) } });
}

for (const shot of SHOTS) {
  for (const width of shot.widths ?? ALL) {
    test(`${shot.id} @${width}`, async ({ page }) => {
      const height = width < 600 ? 844 : 900;
      await page.setViewportSize(shot.openWide === true ? { width: 1440, height: 900 } : { width, height });
      await prepareParityPage(page);
      await routeIcons(page);
      await shot.open(page);
      await page.setViewportSize({ width, height });
      await settleParityPage(page);
      await page.evaluate(() => window.scrollTo(0, 0));
      await page.screenshot({ path: path.join(OUT, `${shot.id}-${width}.png`) });
      await headCrop(page, path.join(HEADS, `${shot.id}-${width}.png`));
    });
  }
}
