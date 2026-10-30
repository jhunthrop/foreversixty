import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test, type BrowserContext, type Locator, type Page } from '@playwright/test';
import { meAddonFixture } from '../../src/fixtures/me-addon';
import {
  OBNOXIOUS,
  armoryPointer,
  keyOf,
  signInWith,
  storePointer,
  type FixtureCharacter,
} from './support/selector';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

// §3.B.2/§3.B.3's comparison needs a real BiS file and a real item table for the fixture's
// own classes -- routed straight from `data/builds/<build>/`, same as the real pipeline
// publishes, rather than whatever `FOREVER_DATA` happened to seed `public/data/<build>/`
// with for this run (the committed fallback fixture at `src/data/fixtures/bis/` carries
// different item ids entirely, which would make `meAddonFixture`'s own worn-gear ids match
// nothing). The same precedent other e2e specs already use for a planner data file
// (`sim-unsimulated-spec.spec.ts`'s own `stubDruidTalents`).
const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../..');
const GEAR_BUILD = '1.60.1.70009';

function realDataFile(relativePath: string): string {
  return readFileSync(path.join(REPO_ROOT, 'data/builds', GEAR_BUILD, relativePath), 'utf8');
}

async function stubRealGearData(page: Page): Promise<void> {
  for (const relativePath of [
    'bis/hunter-marksmanship.json',
    'items/hunter.json',
    'bis/mage-frost.json',
    'items/mage.json',
  ]) {
    await page.route(`**/data/*/${relativePath}`, (route) =>
      route.fulfill({ status: 200, contentType: 'application/json', body: realDataFile(relativePath) }),
    );
  }
}

const NO_SIMS_PAGE = { rows: [], total: 0, page: 1, per_page: 20 };
const NO_RATING = null;

/** Stubs every fetch the signed-in hero and its cards make, beyond `/v1/me` itself, so a
 *  test that does not care about sims/ratings/talents still settles instead of hanging on
 *  a real network call (the real API's CORS policy refuses `localhost` outright, so an
 *  un-stubbed request never resolves at all in a local run). */
async function stubHeroExtras(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/sims**', (route) => route.fulfill(fulfil(NO_SIMS_PAGE)));
  await page.route('**/rating', (route) => route.fulfill(fulfil(NO_RATING)));
  await page.route('**/v1/characters/**', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: { message: 'none' }, request_id: 'r' }, 404)),
  );
}

test('the home page offers Battle.net sign-in when signed out', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/');
  const signedOut = page.getByTestId('home-signed-out');
  await expect(signedOut).toBeVisible();
  // Home rebuild spec §1 / tenet 14: the one sentence this page is allowed to say about
  // itself, replacing the old pitch headline.
  await expect(signedOut.getByRole('heading', { level: 1 })).toHaveText('Play your class better.');
  await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=%2Faccount%3Fsigned_in%3D1$/,
  );
  await expect(page.getByTestId('home-account-panel')).toHaveCount(0);
  // The nine-class picker (§3.A.1): one row, in web/src/data/classes.json's own order.
  const picker = page.getByTestId('home-class-picker');
  await expect(picker.getByTestId(/^home-class-picker-/)).toHaveCount(9);
  await expect(picker.getByTestId('home-class-picker-warrior')).toContainText('Warrior');
  // §3.A.4: the signed-out "Best in slot by class" row is present; the signed-in-only
  // regions are not (display:none, per the session hint).
  await expect(page.getByTestId('home-class-picker-large')).toBeVisible();
  await expect(page.getByTestId('home-upgrades')).toBeHidden();
  await expect(page.getByTestId('home-another-class')).toBeHidden();
});

test('a signed-in visitor with zero characters still sees the Battle.net sign-in CTA, not a blank hero', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toHaveCount(0);
  const signedOut = page.getByTestId('home-signed-out');
  await expect(signedOut).toBeVisible();
  await expect(signedOut).not.toHaveAttribute('inert');
  await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toBeVisible();
});

test('a signed-in panel shows the shared header, the two next-action cards, and hides the signed-out door', async ({
  page,
  context,
}) => {
  await stubHeroExtras(page);
  // The signed-out/signed-in region swap (Best in slot by class vs. Your upgrades/Another
  // class) reads the session cookie's readable half before first paint (Base.astro), not
  // the mocked /v1/me response -- the same pre-paint hint `data-session-hide` already
  // relies on elsewhere on this page.
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [
            {
              key: 'us/normal/zulmara',
              region: 'us',
              ruleset: 'normal',
              name: 'Zulmara',
              class: 'hunter',
              race: 'Troll',
              realm: 'Skyborne',
              level: 24,
              faction: 'horde',
              guild: { id: 9, name: 'Sample Guild', verified: true },
              build: { source: 'addon', captured_at: new Date(Date.now() - 4 * 60_000).toISOString() },
            },
          ],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/zulmara',
        label: 'Zulmara · Hunter',
        classSlug: 'hunter',
        savedAt: new Date().toISOString(),
      }),
    );
  });
  await page.goto('/');
  const panel = page.getByTestId('home-account-panel');
  await expect(panel).toBeVisible();
  await expect(panel.getByRole('heading', { name: 'Zulmara', level: 1 })).toBeVisible();
  // The shared header's descriptor (character header spec 2026-10-11): race and class, the
  // level and the faction; no guild and no realm clause.
  const descriptor = panel.getByTestId('character-header-descriptor');
  await expect(descriptor).toContainText('Troll Hunter');
  await expect(descriptor).toContainText('Level 24');
  await expect(descriptor).toContainText('Horde');
  await expect(descriptor).not.toContainText('Sample Guild');
  // No gear on file (this build carries none): the sync line goes and the answer is one
  // sentence with the link that fixes it -- never a zero.
  await expect(panel.getByTestId('character-header-sync')).toHaveCount(0);
  await expect(panel.getByTestId('character-header-no-gear')).toContainText(
    'Add the addon to see which slots you can upgrade.',
  );
  await expect(panel.getByTestId('character-header-no-gear').getByRole('link')).toHaveAttribute(
    'href',
    '/setup',
  );
  await expect(panel.getByTestId('character-header-action')).toHaveText('See your upgrades →');
  // Talents: this fixture's Zulmara carries no `spec` yet, so the card shows the honest "Pick
  // a spec" state -- never a fabricated figure (§3.B.2). Best in slot is the header's answer.
  await expect(panel.getByTestId('home-hero-card-bis')).toHaveCount(0);
  await expect(panel.getByTestId('home-hero-card-talents')).toContainText('Pick a spec');
  // Simulator: the one card with a real, live figure -- no saved sim for this fixture, so
  // the existing empty-state copy/action shows.
  await expect(panel.getByTestId('home-hero-card-sim')).toContainText('No sim yet.');
  await expect(panel.getByTestId('home-hero-card-sim-run')).toHaveText('Run a sim');
  await expect(panel.getByTestId('home-hero-card-sim-run').locator('[data-glyph="chevron"]')).toHaveCount(1);
  await expect(page.locator('#home-signed-out')).toHaveAttribute('inert', '');
  await expect(page.locator('#home-signed-out')).toHaveAttribute('aria-hidden', 'true');
  // §3.A.4/§3.B.3/§3.B.5: the signed-out class-picker row is gone, replaced by the
  // signed-in-only "Your upgrades" (no spec yet) and "Another class" regions.
  await expect(page.getByTestId('home-class-picker-large')).toBeHidden();
  await expect(page.getByTestId('home-upgrades')).toBeVisible();
  await expect(page.getByTestId('home-upgrades-no-spec')).toBeVisible();
  await expect(page.getByTestId('home-another-class').getByTestId(/^class-crest-/)).toHaveCount(9);
});

const FROSTBYTE: FixtureCharacter = {
  name: 'Frostbyte',
  class: 'Mage',
  spec: 'Frost',
  level: 42,
  faction: 'horde',
  buildDaysAgo: 3,
};
/** Same class, spec, level, faction and realm as OBNOXIOUS: only the name differs, so a
 *  switch between the two can move nothing but what the name itself does. */
const TWO_LINE_NAME: FixtureCharacter = { ...OBNOXIOUS, name: 'Sir Obnoxious Yellington' };

const selectorTrigger = (page: Page) => page.getByTestId('character-selector');
const chooseInHeader = async (page: Page, character: FixtureCharacter): Promise<void> => {
  await selectorTrigger(page).click();
  await page
    .getByTestId('selector-panel')
    .getByTestId(`selector-row-${keyOf(character)}`)
    .getByRole('button')
    .first()
    .click();
};

test('the header says "Your character" with one or several characters and has no switcher of its own', async ({
  page,
}) => {
  await stubHeroExtras(page);
  await signInWith(page, [{ ...OBNOXIOUS, race: 'Gnome' }, FROSTBYTE]);
  await storePointer(page, armoryPointer(OBNOXIOUS));
  await page.goto('/');
  const panel = page.getByTestId('home-account-panel');
  await expect(panel.getByRole('heading', { name: 'Obnoxious Yell', level: 1 })).toBeVisible();
  const header = panel.getByTestId('character-header');
  await expect(header.locator('.label').first()).toHaveText('Your character');
  await expect(header.getByTestId('character-header-descriptor')).toContainText('Gnome Fury Warrior');
  // No hint line, no switcher: the selector in the nav bar is the one way to change character.
  await expect(panel.getByText('Change character')).toHaveCount(0);
  await expect(page.getByText('Switch character')).toHaveCount(0);
  await expect(page.getByText('Add one')).toHaveCount(0);
  await expect(page.getByTestId('home-switch-character-panel')).toHaveCount(0);
  await expect(header.getByRole('button')).toHaveCount(0);
  await expect(page.getByTestId('home-launch-block')).toBeHidden();
  // The flat faction logo, 16 px, beside the faction word.
  const emblem = header.locator('.ch-faction img');
  await expect(emblem).toHaveAttribute('src', '/icons/hd/faction/alliance-logo-64.webp');
  await expect(emblem).toHaveAttribute('width', '16');

  // The header selector is the one way to change character; the header follows it in place.
  await page.evaluate(() => {
    (window as unknown as { __kept: boolean }).__kept = true;
  });
  await chooseInHeader(page, FROSTBYTE);
  await expect(panel.getByRole('heading', { level: 1 })).toHaveText('Frostbyte');
  await expect(header.locator('.label').first()).toHaveText('Your character');
  expect(await page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept)).toBe(true);
});

test('on a phone a header switch from a one-line to a two-line name does not move the upgrades table', async ({
  page,
}) => {
  await page.setViewportSize({ width: 390, height: 844 });
  await stubHeroExtras(page);
  await signInWith(page, [OBNOXIOUS, TWO_LINE_NAME]);
  await storePointer(page, armoryPointer(OBNOXIOUS));
  await page.goto('/');
  const panel = page.getByTestId('home-account-panel');
  const heading = panel.getByRole('heading', { level: 1 });
  await expect(heading).toHaveText('Obnoxious Yell');
  await expect(panel.getByTestId('home-hero-cards-slot')).toBeVisible();
  const table = page.getByTestId('home-upgrades');
  const oneLineTop = (await table.boundingBox())!.y;

  await chooseInHeader(page, TWO_LINE_NAME);
  await expect(heading).toHaveText('Sir Obnoxious Yellington');
  expect((await table.boundingBox())!.y).toBe(oneLineTop);
});

test('a name longer than two lines is cut with an ellipsis, not allowed to grow the header', async ({
  page,
}) => {
  await page.setViewportSize({ width: 1440, height: 900 });
  await stubHeroExtras(page);
  const longName: FixtureCharacter = { ...OBNOXIOUS, name: 'Sir Obnoxious Yellington the Unreasonably Long' };
  await signInWith(page, [longName]);
  await storePointer(page, armoryPointer(longName));
  await page.goto('/');
  const name = page.getByTestId('home-account-panel').getByRole('heading', { level: 1 });
  await expect(name).toBeVisible();
  // Two 24 px lines at a 1.2 line height; the whole name stays in the title.
  expect((await name.boundingBox())!.height).toBeLessThan(60);
  await expect(name).toHaveAttribute('title', longName.name);
});

const SAVED_SIM = {
  sim_id: 's1',
  spec: 'hunter-marksmanship',
  dps: 26.4,
  engine_version: '1',
  created_at: '2026-10-09T00:00:00Z',
  title: '',
  kind: 'run',
};

async function openHomeWithSavedSim(page: Page, context: BrowserContext, sim: object): Promise<Locator> {
  await stubHeroExtras(page);
  await stubRealGearData(page);
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: true, data: meAddonFixture, error: null, request_id: 'r' })),
  );
  await page.route('**/v1/sims**', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: { rows: [sim], total: 1, page: 1, per_page: 1 },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.goto('/');
  return page.getByTestId('home-account-panel').getByTestId('home-hero-card-sim');
}

test('the Simulator card pairs a saved sim of the band own setup with the band best in slot, a unit on both', async ({
  page,
  context,
}) => {
  const card = await openHomeWithSavedSim(page, context, {
    ...SAVED_SIM,
    band: 20,
    faction: 'horde',
    preset: 'bare',
  });
  await expect(card.getByTestId('home-hero-card-sim-value')).toHaveText('26 DPS now');
  await expect(card).toContainText(/\d+\.\d DPS at best in slot for your level range/, { timeout: 10_000 });
});

test('the Simulator card shows the band best in slot alone when the saved sim is not comparable', async ({
  page,
  context,
}) => {
  const card = await openHomeWithSavedSim(page, context, SAVED_SIM);
  await expect(card.getByTestId('home-hero-card-sim-value')).toHaveText(/^\d+\.\d DPS$/, {
    timeout: 10_000,
  });
  await expect(card).toContainText('at best in slot for your level range');
  await expect(card).toContainText('Run');
  await expect(card).not.toContainText('DPS now');
});

test('a returning signed-in visitor sees the hub from the session snapshot before /v1/me answers', async ({
  page,
  context,
}) => {
  const body = {
    ok: true,
    data: {
      user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
      characters: [
        {
          key: 'us/normal/kiloz',
          region: 'us',
          ruleset: 'normal',
          name: 'Kiloz',
          class: 'warrior',
          level: 60,
        },
      ],
      guilds: [],
    },
    error: null,
    request_id: 'r',
  };
  await stubHeroExtras(page);
  // The readable half of the session cookie pair is what makes the snapshot trusted.
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(body)));
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toBeVisible();
  const stored = await page.evaluate(() =>
    Object.keys(window.localStorage)
      .filter((k) => k.startsWith('fs.q.'))
      .map((k) => window.localStorage.getItem(k) ?? '')
      .join('\n'),
  );
  expect(stored).toContain('Kiloz');

  // Second load: /v1/me is held for five seconds, yet the hub renders at once from the
  // snapshot, and the signed-out block never shows (Base.astro's session hint).
  await page.unroute('**/v1/me');
  await page.route('**/v1/me', async (route) => {
    await new Promise((resolve) => setTimeout(resolve, 5000));
    await route.fulfill(fulfil(body));
  });
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toBeVisible({ timeout: 3000 });
  await expect(page.getByTestId('home-signed-out')).toBeHidden();
});

test("a signed-in panel with a real gear/talent export sees the two upgrade figures, Talents and Your upgrades, matching the SignedIn mock's own state", async ({
  page,
  context,
}) => {
  await stubHeroExtras(page);
  await stubRealGearData(page);
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: true, data: meAddonFixture, error: null, request_id: 'r' })),
  );
  await page.goto('/');

  const panel = page.getByTestId('home-account-panel');
  await expect(panel.getByRole('heading', { name: 'Zulmara', level: 1 })).toBeVisible();
  const descriptor = panel.getByTestId('character-header-descriptor');
  await expect(descriptor).toContainText('Troll Marksmanship Hunter');
  await expect(descriptor).toContainText('Level 24');
  await expect(descriptor).toContainText('Horde');
  await expect(panel.getByTestId('character-header-sync')).toContainText('Addon');

  // The answer (character header spec 2026-10-11 section 4): the upgrades are counted against
  // the committed data/builds/<build>/bis/hunter-marksmanship.json band 20 (horde), which the
  // nightly re-ranks, so the test pins the relationships between the figures and not the
  // numbers: the first figure is the number of rows "Your upgrades" lists, the second is the
  // gain as a DPS figure at the same size, and the total never mixes the full-sim `score` with
  // a `scoreItem`-only estimate (fix round 1 item A.1).
  await expect(panel.getByTestId('character-header-answer-label')).toHaveText(
    'Upgrades for your level range',
  );
  const countFigure = panel.getByTestId('character-header-figure-1');
  await expect(countFigure).toHaveText(/^\d+ upgrades?$/, { timeout: 10_000 });
  const upgradeCount = Number.parseInt((await countFigure.textContent()) ?? '', 10);
  expect(upgradeCount).toBeGreaterThan(0);
  await expect(panel.getByTestId('character-header-figure-2')).toHaveText(/^\+\d+\.\d DPS$/);
  const sizes = await panel
    .locator('.ch-figure')
    .evaluateAll((nodes) => nodes.map((node) => getComputedStyle(node).fontSize));
  expect(new Set(sizes).size).toBe(1);
  await expect(panel.getByTestId('character-header-note')).toContainText(
    'Best in slot for 20 to 29, together',
  );

  // Talents (§3.B.2): some points differ from the band's own build.
  const talentsCard = panel.getByTestId('home-hero-card-talents');
  await expect(talentsCard.getByTestId('home-hero-card-talents-value')).toHaveText('Unoptimized');
  await expect(talentsCard).toContainText(/\d+ of the \d+ points in the 20 to 29 build differ/);
  // The count is the band's own points that differ, so it can never exceed the band's total.
  const [, differ, total] = /(\d+) of the (\d+) points/.exec((await talentsCard.textContent()) ?? '') ?? [];
  expect(Number(differ)).toBeGreaterThan(0);
  expect(Number(differ)).toBeLessThanOrEqual(Number(total));

  // "Your upgrades" (§3.B.3): one row per slot the fixture's own worn item differs from the
  // band's pick, each reading the band's own sim-verified gain.
  await expect(page.getByTestId('home-upgrades-list')).toBeVisible();
  await expect(page.locator('[data-testid^="home-upgrade-row-"]')).toHaveCount(upgradeCount);
  await expect(page.locator('[data-testid^="home-upgrade-gain-"]').first()).toHaveText(/^\+\d+\.\d DPS$/);
  const alreadyBis = page.getByTestId('home-upgrades-already-bis');
  await expect(alreadyBis).toContainText('Already best in slot:');
  await expect(alreadyBis).toContainText('more slot');
  // Fix round 1 item A.2: every already-best-in-slot name carries its rarity colour and the
  // shared item tooltip, the same as every other item name on this page.
  await expect(alreadyBis.locator('[data-testid^="item-hover-"]').first()).toBeVisible();

  // The worn side of a row names the item and nothing else: no "you wear this" filler.
  await expect(page.getByTestId('home-upgrades-list')).not.toContainText('you wear this');
  await expect(page.getByTestId('home-switch-character-panel')).toHaveCount(0);
});

// The first layout settles inside this window (the hero's font swap moves it once, the same with or
// without the reorder); anything after it would be the island or a reserve moving the page.
const FIRST_PAINT_SETTLE_MS = 300;

// Owner 2026-10-29: signed in, "Your upgrades" sits directly under the character cards, then the
// lookup, then the rest; the lookup is moved there before first paint, so nothing shifts.
for (const width of [390, 1440]) {
  test(`signed in at ${width}: dates line, cards, upgrades, lookup, then the rest, nothing moving after the first layout`, async ({
    page,
  }) => {
    await page.setViewportSize({ width, height: 900 });
    await stubHeroExtras(page);
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.addInitScript((afterMs) => {
      const late: number[] = [];
      (window as unknown as { __lateShifts: number[] }).__lateShifts = late;
      new PerformanceObserver((list) => {
        for (const entry of list.getEntries() as unknown as { value: number; startTime: number }[]) {
          if (entry.startTime > afterMs) late.push(entry.value);
        }
      }).observe({ type: 'layout-shift', buffered: true });
    }, FIRST_PAINT_SETTLE_MS);
    await page.goto('/');
    await expect(page.getByTestId('home-upgrades')).toBeVisible();
    const top = async (testId: string): Promise<number> => (await page.getByTestId(testId).boundingBox())!.y;
    const cards = await top('home-account-block');
    const line = await top('home-dates-line');
    const upgrades = await top('home-upgrades');
    const lookup = await top('home-lookup-cell');
    const outcomes = await top('home-outcome-planner');
    expect(line).toBeGreaterThan(cards);
    expect(upgrades).toBeGreaterThan(line);
    expect(lookup).toBeGreaterThan(upgrades);
    expect(outcomes).toBeGreaterThan(lookup);
    await expect(page.getByTestId('home-lookup-slot').getByTestId('home-lookup-cell')).toHaveCount(1);
    await page.waitForTimeout(500);
    const late = await page.evaluate(() => (window as unknown as { __lateShifts: number[] }).__lateShifts);
    expect(late, 'nothing moves once the first layout has settled (hydration, reserves)').toEqual([]);
  });
}

test('signed out the lookup stays in the hero, above Best in slot by class', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 900 });
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/');
  await expect(page.getByTestId('home-lookup-slot')).toBeHidden();
  const lookup = (await page.getByTestId('home-lookup-cell').boundingBox())!.y;
  const classes = (await page.getByTestId('home-class-picker-large').boundingBox())!.y;
  expect(lookup).toBeLessThan(classes);
});

// Fix round 1: no empty band above "Latest" -- the signed-in Get set up card used to hold its height
// invisibly until its island ran, and signed out the island's 44 px sat under the card.
const MAX_BAND_ABOVE_LATEST_PX = 64;
for (const width of [390, 1440]) {
  test(`signed in at ${width}: Latest follows the panels above by at most ${MAX_BAND_ABOVE_LATEST_PX}px`, async ({
    page,
  }) => {
    await page.setViewportSize({ width, height: 900 });
    await stubHeroExtras(page);
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/');
    await expect(page.getByTestId('home-upgrades')).toBeVisible();
    await page.waitForLoadState('networkidle');
    const gap = await page.evaluate(() => {
      const bottomOf = (testId: string): number =>
        document.querySelector(`[data-testid="${testId}"]`)!.getBoundingClientRect().bottom;
      // Latest's title sits in its card's head now, so the card's top is what follows the panels above.
      const latest = [...document.querySelectorAll('h2')]
        .find((h) => h.textContent === 'Latest')!
        .closest('section')!;
      return latest.getBoundingClientRect().top - bottomOf('home-around-the-site');
    });
    expect(gap).toBeLessThanOrEqual(MAX_BAND_ABOVE_LATEST_PX);
    await expect(page.locator('#home-get-set-up-signed-out')).toBeHidden();
  });

  test(`signed out at ${width}: Latest follows the Get set up card by at most ${MAX_BAND_ABOVE_LATEST_PX}px`, async ({
    page,
  }) => {
    await page.setViewportSize({ width, height: 900 });
    await page.route('**/v1/me', (route) =>
      route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
    );
    await page.goto('/');
    const card = (await page.getByTestId('home-get-set-up-signed-out').boundingBox())!;
    const latest = (await page
      .getByRole('heading', { name: 'Latest' })
      .locator('xpath=ancestor::section')
      .boundingBox())!;
    expect(latest.y - (card.y + card.height)).toBeLessThanOrEqual(MAX_BAND_ABOVE_LATEST_PX);
  });
}
