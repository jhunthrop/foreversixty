// web/tests/captures/h-edges/capture.spec.ts
// Batch H lane h-edges: review captures, not a test of behaviour. Each surface at 390, 1024, 1440
// and 2000 on the parity fixtures (the guild states through their parity openers), each beside its
// mock where one is rendered (design/mocks/renders/web/<mock>-<width>.png from ws_shoot.mjs, and the
// batch 4 directions mock b4-guilds for the guild landing). The S11 system captures (the shell and
// the design sheet at 390 to 2560) go to design/reviews/2026-11-09-redesign/system/.
import { existsSync, mkdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Page } from '@playwright/test';
import { prepareParityPage, settleParityPage } from '../../e2e/support/parity';
import { openGuildState } from '../../e2e/support/guild-parity';
import { GUILD_PARITY_NOW } from '../../e2e/support/guild-parity/clock';
import { CHARACTER_PATH } from '../../e2e/support/character-fixture';
import { sideBySide, sizeTo } from '../support/capture-kit';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const REPO = path.resolve(HERE, '../../../..');
const PHASE = process.env.CAPTURE_PHASE ?? 'after';
const OUT = path.join(REPO, 'design/reviews/2026-11-09-h-edges', PHASE);
const SYSTEM = path.join(REPO, 'design/reviews/2026-11-09-redesign/system');
const RENDERS = path.join(REPO, 'design/mocks/renders/web');
const DIRECTIONS = path.join(REPO, 'design/mocks/directions/png');
mkdirSync(OUT, { recursive: true });
mkdirSync(SYSTEM, { recursive: true });

const WIDTHS = [390, 1024, 1440, 2000] as const;
const SYSTEM_WIDTHS = [390, 1024, 1440, 2000, 2560] as const;
const PHONE_HEIGHT = 844;
const DESKTOP_HEIGHT = 1000;

const heightFor = (width: number): number => (width < 768 ? PHONE_HEIGHT : DESKTOP_HEIGHT);

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'h-edges' }),
});

/** The mock a guild state is reviewed against at one width: the directions mock where the batch
 *  draws that width, else the parity mock's render. */
function mockFor(mock: string, width: number, directions?: string): string | null {
  const drawn = directions === undefined ? null : path.join(DIRECTIONS, `${directions}-${width}.png`);
  if (drawn !== null && existsSync(drawn)) return drawn;
  const render = path.join(RENDERS, `${mock}-${width}.png`);
  return existsSync(render) ? render : null;
}

async function shoot(page: Page, name: string, width: number, mock: string | null): Promise<void> {
  const built = path.join(OUT, `${name}-${width}.png`);
  await page.screenshot({ path: built, fullPage: true });
  if (mock === null) return;
  const board = await page.context().newPage();
  await sideBySide(board, mock, built, path.join(OUT, `side-by-side-${name}-${width}.png`), width);
  await board.close();
}

interface GuildSurface {
  state: string;
  mock: string;
  directions?: string;
}

const GUILD: readonly GuildSurface[] = [
  { state: 'landing-in', mock: 'web-guild-landing-in', directions: 'b4-guilds' },
  { state: 'landing-out', mock: 'web-guild-landing-out' },
  { state: 'join', mock: 'web-guild-join' },
  { state: 'profile', mock: 'web-guild-profile' },
  { state: 'report', mock: 'web-guild-report' },
];

for (const surface of GUILD) {
  for (const width of WIDTHS) {
    test(`guild ${surface.state} @${width}`, async ({ page }) => {
      await sizeTo(page, width, heightFor(width));
      await page.clock.setFixedTime(GUILD_PARITY_NOW);
      await prepareParityPage(page);
      await openGuildState(page, surface.state);
      await settleParityPage(page);
      await shoot(page, surface.state, width, mockFor(surface.mock, width, surface.directions));
    });
  }
}

const me = (role: string) => ({
  user: { id: 78, battletag: 'Owner#1894', email: null, role, anonymize: false },
  characters: [],
  guilds: [],
});

const link = (slug: string, target: string, label: string, f: readonly number[]) => ({
  slug,
  target,
  label,
  created_at: '2026-11-02T00:00:00Z',
  clicks: f[0],
  bot_clicks: f[1],
  sign_ins: f[2],
  linked: f[3],
  premium: f[4],
});

const LINKS = [
  link('reddit', '/', 'Reddit', [412, 38, 61, 44, 9]),
  link('reddit-sim', '/sim', 'Reddit: sim post', [870, 31, 97, 80, 14]),
  link('discord', '/', 'Discord', [340, 512, 71, 66, 12]),
  link('curseforge', '/', 'CurseForge', [58, 2, 9, 9, 2]),
  link('wago', '/', 'Wago', [0, 0, 0, 0, 0]),
];

const TESTERS = [
  { email: 'new@example.com', platform: 'ios', created_at: '2026-11-08T12:00:00Z' },
  { email: 'droid@example.com', platform: 'android', created_at: '2026-11-07T09:00:00Z' },
  { email: 'old@example.com', platform: 'ios', created_at: '2026-11-01T23:30:00Z' },
];

const ADMIN: readonly { name: string; path: string; api: string; body: unknown; ready: string }[] = [
  {
    name: 'admin-links',
    path: '/admin/links',
    api: '**/v1/admin/golinks',
    body: { links: LINKS },
    ready: 'links-row-reddit-sim',
  },
  {
    name: 'admin-testers',
    path: '/admin/testers',
    api: '**/v1/admin/app-testers',
    body: { testers: TESTERS },
    ready: 'testers-grid-row',
  },
];

for (const admin of ADMIN) {
  for (const width of WIDTHS) {
    test(`${admin.name} @${width}`, async ({ page }) => {
      await sizeTo(page, width, heightFor(width));
      await prepareParityPage(page);
      await page.context().addCookies([{ name: 'fs_csrf', value: 'tok', domain: 'localhost', path: '/' }]);
      await page.route('**/v1/me', (route) => route.fulfill(envelope(me('admin'))));
      await page.route(admin.api, (route) => route.fulfill(envelope(admin.body)));
      await page.goto(admin.path);
      await page.getByTestId(admin.ready).first().waitFor();
      await settleParityPage(page);
      await shoot(page, admin.name, width, mockFor(`web-${admin.name}`, width));
    });
  }
}

test.describe('with JavaScript off', () => {
  test.use({ javaScriptEnabled: false });
  for (const width of WIDTHS) {
    test(`character no-script @${width}`, async ({ page }) => {
      await sizeTo(page, width, heightFor(width));
      await prepareParityPage(page);
      await page.goto(CHARACTER_PATH);
      await page.getByTestId('character-noscript').waitFor();
      await page.evaluate(() => document.fonts.ready);
      await page.screenshot({ path: path.join(OUT, `character-noscript-${width}.png`), fullPage: true });
    });
  }
});

/** The shell on a quiet page (the about page: the band, the bar, a reading body, the footer) and
 *  the design sheet, at every system width; at 390 the menu sheet open too. */
for (const width of SYSTEM_WIDTHS) {
  test(`system shell and sheet @${width}`, async ({ page }) => {
    await sizeTo(page, width, heightFor(width));
    await prepareParityPage(page);
    await page.route('**/v1/me', (route) =>
      route.fulfill({ status: 401, contentType: 'application/json', body: '{"ok":false}' }),
    );
    await page.goto('/about');
    await settleParityPage(page);
    await page.screenshot({ path: path.join(SYSTEM, `shell-${width}.png`) });
    await page
      .locator('footer')
      .last()
      .screenshot({ path: path.join(SYSTEM, `shell-footer-${width}.png`) });
    if (width < 1024) {
      await page.getByTestId('menu-button').click();
      await page.waitForTimeout(300);
      await page.screenshot({ path: path.join(SYSTEM, `shell-menu-${width}.png`) });
    }
    await page.goto('/design-sheet');
    await settleParityPage(page);
    await page.screenshot({ path: path.join(SYSTEM, `sheet-${width}.png`), fullPage: true });
  });
}
