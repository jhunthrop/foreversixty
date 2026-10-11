// web/tests/e2e/bis-landing.spec.ts
// The /bis landing (design/specs/2026-10-10-bis-landing.md): signed out, the class grid's first row
// is above the fold at 1440, the strip defaults to 20 to 29 and a pill rewrites every spec link;
// signed in, the lead shows the selected character's band, count and one button, the strip never
// moves it, and the header selector moves the lead, the "you" pill and the "Yours" row in place;
// on a phone the pills are 44 px tall with no wrapped label.
import { expect, test, type Page } from '@playwright/test';
import { bisBand, readBisFile, readSpecCatalog } from './support/bis-file';
import {
  OBNOXIOUS,
  armoryPointer,
  keyOf,
  signInWith,
  signedOut,
  storePointer,
  type FixtureCharacter,
} from './support/selector';

const DESKTOP = { width: 1440, height: 900 };
const PHONE = { width: 390, height: 844 };
const BANDS = [20, 30, 40, 50, 60];
const MIN_TOUCH_TARGET_PX = 44;

const WARRIOR_SPEC = 'warrior-fury';
const warriorBand = bisBand(WARRIOR_SPEC, 'alliance', 40, 'bare');
const sourcedPicks = (warriorBand.slots ?? []).filter((slot) => typeof slot.item_id === 'number');
const unsourcedSlots = (warriorBand.slots ?? []).filter((slot) => typeof slot.item_id !== 'number');
const WORN_COUNT = 4;
const wornGear = Object.fromEntries(
  sourcedPicks.slice(0, WORN_COUNT).map((slot) => [slot.slot, slot.item_id as number]),
);

const FURY: FixtureCharacter = { ...OBNOXIOUS, level: 47, gear: wornGear };
const HUNTER: FixtureCharacter = {
  name: 'Quickshot',
  class: 'Hunter',
  spec: 'Marksmanship',
  level: 33,
  faction: 'horde',
  buildDaysAgo: 1,
};

const specLink = (page: Page, classSlug: string, specSlug: string) =>
  page.getByTestId(`bis-spec-${classSlug}-${specSlug}`);
const pill = (page: Page, band: number) => page.getByTestId(`bis-landing-band-${band}`);
const lead = (page: Page) => page.getByTestId('bis-lead');
const selector = (page: Page) => page.getByTestId('character-selector');

/** The client fetches each spec's published file; serve the one the build was made from. */
async function stubPublishedBis(page: Page, specs: string[]): Promise<void> {
  for (const spec of specs) {
    await page.route(`**/data/*/bis/${spec}.json`, (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify(readBisFile(spec)),
      }),
    );
  }
}

async function openSignedIn(
  page: Page,
  characters: FixtureCharacter[],
  specs = [WARRIOR_SPEC],
): Promise<void> {
  await signInWith(page, characters);
  await storePointer(page, armoryPointer(characters[0]!));
  await stubPublishedBis(page, specs);
  await page.goto('/bis');
  await expect(page.getByTestId('bis-lead-name')).toHaveText(characters[0]!.name);
}

async function choose(page: Page, character: FixtureCharacter): Promise<void> {
  await page.evaluate(() => {
    (window as unknown as { __kept: boolean }).__kept = true;
  });
  await selector(page).click();
  await page
    .getByTestId('selector-panel')
    .getByTestId(`selector-row-${keyOf(character)}`)
    .getByRole('button')
    .first()
    .click();
  await expect(page.getByTestId('selector-name')).toHaveText(character.name);
}

test.describe('signed out, desktop', () => {
  test.use({ viewport: DESKTOP });

  test.beforeEach(async ({ page }) => {
    await signedOut(page);
    await page.goto('/bis');
  });

  test('the first row of the class grid is above the fold, in game order', async ({ page }) => {
    const first = ['warrior', 'paladin', 'hunter'];
    for (const slug of first) {
      const box = await page.getByTestId(`bis-class-${slug}`).boundingBox();
      expect(box?.y ?? Infinity).toBeLessThan(DESKTOP.height);
      expect((box?.y ?? 0) + (box?.height ?? Infinity)).toBeLessThanOrEqual(DESKTOP.height);
    }
    const order = await page
      .locator('[data-testid^="bis-class-"]:not([data-testid="bis-class-grid"])')
      .evaluateAll((nodes) => nodes.map((node) => node.getAttribute('data-testid')));
    expect(order).toEqual(
      ['warrior', 'paladin', 'hunter', 'rogue', 'priest', 'shaman', 'mage', 'warlock', 'druid'].map(
        (slug) => `bis-class-${slug}`,
      ),
    );
  });

  test('the hero says its one line, the tier list pointer and one Updated stamp', async ({ page }) => {
    await expect(page.getByRole('heading', { level: 1 })).toHaveText('Best in slot');
    await expect(
      page.getByText('The best gear you can wear at your level, for every spec. Open yours.'),
    ).toBeVisible();
    await expect(page.getByTestId('bis-index-tier-link')).toHaveAttribute('href', '/tiers');
    await expect(page.getByTestId('bis-landing').getByText(/^Updated /)).toHaveCount(1);
    await expect(page.getByTestId('bis-grid-heading')).toHaveText('Pick a class and spec', {
      useInnerText: true,
    });
  });

  test('the lead is absent and the strip defaults to 20 to 29', async ({ page }) => {
    await expect(lead(page)).toBeHidden();
    await expect(pill(page, 20)).toHaveAttribute('aria-current', 'true');
    await expect(page.getByTestId('bis-strip-head')).toHaveText('Spec links open at this band');
    await expect(pill(page, 20).locator('.pb-you')).toBeHidden();
    for (const band of BANDS.slice(1))
      await expect(pill(page, band)).not.toHaveAttribute('aria-current', 'true');
  });

  test('every spec row links to its page at 20 to 29 with no faction, and no row carries a figure', async ({
    page,
  }) => {
    const catalog = readSpecCatalog();
    const links = page.locator('[data-spec-link]');
    await expect(links).toHaveCount(catalog.length);
    for (const row of catalog) {
      await expect(specLink(page, row.class_slug, row.spec_slug)).toHaveAttribute(
        'href',
        `/bis/${row.class_slug}/${row.spec_slug}#band-alliance-20`,
      );
    }
    await expect(page.getByTestId('bis-landing')).not.toContainText(/\bDPS\b|\bHPS\b|effective health/i);
    await expect(page.getByTestId('bis-landing')).not.toContainText(/example character/i);
  });

  test('a pill selects its band and rewrites every spec link, keeping the choice in the URL', async ({
    page,
  }) => {
    await pill(page, 40).click();
    await expect(pill(page, 40)).toHaveAttribute('aria-current', 'true');
    await expect(pill(page, 20)).not.toHaveAttribute('aria-current', 'true');
    await expect(page).toHaveURL(/\/bis\?band=40$/);
    for (const row of readSpecCatalog()) {
      await expect(specLink(page, row.class_slug, row.spec_slug)).toHaveAttribute(
        'href',
        `/bis/${row.class_slug}/${row.spec_slug}#band-alliance-40`,
      );
    }
    await specLink(page, 'warrior', 'fury').click();
    await expect(page).toHaveURL(/\/bis\/warrior\/fury#band-alliance-40$/);
  });

  test('the choice survives a reload (the URL) and a fresh visit (storage)', async ({ page }) => {
    await pill(page, 50).click();
    await page.reload();
    await expect(pill(page, 50)).toHaveAttribute('aria-current', 'true');
    await page.goto('/bis');
    await expect(pill(page, 50)).toHaveAttribute('aria-current', 'true');
    await expect(specLink(page, 'mage', 'frost')).toHaveAttribute('href', '/bis/mage/frost#band-alliance-50');
  });

  test('every pill carries a caption, and the arrow keys move between the pills', async ({ page }) => {
    for (const band of BANDS)
      await expect(page.getByTestId(`bis-landing-band-caption-${band}`)).not.toBeEmpty();
    await expect(page.getByTestId('bis-landing-band-caption-20')).toHaveText(/^Top sources: /);
    await pill(page, 20).focus();
    await page.keyboard.press('ArrowRight');
    await expect(pill(page, 30)).toBeFocused();
    await page.keyboard.press('End');
    await expect(pill(page, 60)).toBeFocused();
  });

  test('hovering a spec row moves nothing', async ({ page }) => {
    const row = specLink(page, 'rogue', 'combat');
    const before = await row.boundingBox();
    const cardBefore = await page.getByTestId('bis-class-rogue').boundingBox();
    await row.hover();
    expect(await row.boundingBox()).toEqual(before);
    expect(await page.getByTestId('bis-class-rogue').boundingBox()).toEqual(cardBefore);
    await expect(row).toHaveCSS('background-color', 'rgb(16, 22, 36)');
  });

  test('a tank and a healer row carry their role word and no DPS row does', async ({ page }) => {
    await expect(specLink(page, 'priest', 'holy')).toContainText('Healer');
    await expect(specLink(page, 'warrior', 'protection')).toContainText('Tank');
    await expect(specLink(page, 'warrior', 'fury')).not.toContainText(/Tank|Healer/);
  });

  test('every class shows one circular crest and every spec row a spec icon', async ({ page }) => {
    for (const slug of ['warrior', 'druid']) {
      const card = page.getByTestId(`bis-class-${slug}`);
      await expect(card.locator('img.crest')).toHaveCount(1);
      await expect(card.locator('img.crest')).toHaveAttribute('src', `/icons/hd/crests/${slug}.webp`);
    }
    const iconSrcs = await page
      .locator('.spr-icon')
      .evaluateAll((nodes) => nodes.map((node) => node.getAttribute('src')));
    expect(iconSrcs).toHaveLength(readSpecCatalog().length);
    for (const src of iconSrcs) expect(src).toMatch(/\/data\/[^/]+\/icons\/[a-z0-9_]+\.webp$/);
  });

  test('a Druid card has four rows', async ({ page }) => {
    await expect(page.getByTestId('bis-class-druid').locator('[data-spec-link]')).toHaveCount(4);
  });
});

test.describe('signed in, desktop', () => {
  test.use({ viewport: DESKTOP });

  test('the lead shows the character, its band, its count and the one button to its own page', async ({
    page,
  }) => {
    await openSignedIn(page, [FURY, HUNTER]);
    await expect(page.getByTestId('bis-lead-descriptor')).toContainText('Fury Warrior');
    await expect(page.getByTestId('bis-lead-descriptor')).toContainText('Level 47');
    await expect(page.getByTestId('bis-lead-descriptor')).toContainText('Alliance');
    await expect(page.getByTestId('bis-lead-sync')).toContainText('Addon · ');
    await expect(page.getByTestId('bis-lead-count')).toHaveText(
      `${WORN_COUNT} of ${sourcedPicks.length} slots`,
    );
    await expect(page.getByTestId('bis-lead-cta')).toHaveText('Fury gear for 40 to 49 →');
    await expect(page.getByTestId('bis-lead-cta')).toHaveAttribute(
      'href',
      '/bis/warrior/fury?faction=alliance#band-alliance-40',
    );
    await expect(page.getByTestId('bis-lead-link')).toHaveAttribute(
      'href',
      '/bis/warrior/fury?faction=alliance#band-alliance-40',
    );
    if (unsourcedSlots.length > 0) await expect(lead(page)).toContainText('no pick at this band');
    else await expect(lead(page)).not.toContainText('no pick at this band');
    await expect(page.getByTestId('bis-grid-heading')).toHaveText('Another class or spec', {
      useInnerText: true,
    });
  });

  test('the character’s band is selected and marked, and its spec row is Yours', async ({ page }) => {
    await openSignedIn(page, [FURY, HUNTER]);
    await expect(pill(page, 40)).toHaveAttribute('aria-current', 'true');
    await expect(pill(page, 40).locator('.pb-you')).toBeVisible();
    await expect(pill(page, 40).locator('.pb-you')).toHaveText('you');
    await expect(specLink(page, 'warrior', 'fury')).toHaveClass(/is-you/);
    await expect(specLink(page, 'warrior', 'fury')).toContainText('Yours');
    await expect(page.locator('[data-yours]:visible')).toHaveCount(1);
    await expect(specLink(page, 'warrior', 'fury')).toHaveAttribute(
      'href',
      '/bis/warrior/fury?faction=alliance#band-alliance-40',
    );
  });

  test('the strip retargets the spec rows and never the lead', async ({ page }) => {
    await openSignedIn(page, [FURY, HUNTER]);
    await pill(page, 50).click();
    await expect(specLink(page, 'mage', 'frost')).toHaveAttribute(
      'href',
      '/bis/mage/frost?faction=alliance#band-alliance-50',
    );
    await expect(page.getByTestId('bis-lead-cta')).toHaveAttribute(
      'href',
      '/bis/warrior/fury?faction=alliance#band-alliance-40',
    );
    await expect(page.getByTestId('bis-lead-cta')).toHaveText('Fury gear for 40 to 49 →');
    await expect(page.getByTestId('bis-lead-count')).toHaveText(
      `${WORN_COUNT} of ${sourcedPicks.length} slots`,
    );
    await expect(pill(page, 40).locator('.pb-you')).toBeVisible();
    await expect(pill(page, 40)).not.toHaveAttribute('aria-current', 'true');
  });

  test('the gold button opens the character’s own spec at its band', async ({ page }) => {
    await openSignedIn(page, [FURY, HUNTER]);
    await page.getByTestId('bis-lead-cta').click();
    await expect(page).toHaveURL(/\/bis\/warrior\/fury\?faction=alliance#band-alliance-40$/);
  });

  test('a header switch moves the lead, the you pill and Yours in place, with no reload', async ({
    page,
  }) => {
    await openSignedIn(page, [FURY, HUNTER], [WARRIOR_SPEC, 'hunter-marksmanship']);
    await choose(page, HUNTER);
    await expect(page.getByTestId('bis-lead-name')).toHaveText('Quickshot');
    await expect(page.getByTestId('bis-lead-descriptor')).toContainText('Marksmanship Hunter');
    await expect(page.getByTestId('bis-lead-descriptor')).toContainText('Horde');
    await expect(page.getByTestId('bis-lead-cta')).toHaveText('Marksmanship gear for 30 to 39 →');
    await expect(page.getByTestId('bis-lead-cta')).toHaveAttribute(
      'href',
      '/bis/hunter/marksmanship?faction=horde#band-horde-30',
    );
    await expect(pill(page, 30)).toHaveAttribute('aria-current', 'true');
    await expect(pill(page, 30).locator('.pb-you')).toBeVisible();
    await expect(pill(page, 40).locator('.pb-you')).toBeHidden();
    await expect(specLink(page, 'hunter', 'marksmanship')).toHaveClass(/is-you/);
    await expect(specLink(page, 'warrior', 'fury')).not.toHaveClass(/is-you/);
    await expect(page.locator('[data-yours]:visible')).toHaveCount(1);
    await expect(specLink(page, 'mage', 'frost')).toHaveAttribute(
      'href',
      '/bis/mage/frost?faction=horde#band-horde-30',
    );
    expect(await page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept)).toBe(true);
  });

  test('a stale sync says so in the words the addon needs, and relabels the count', async ({ page }) => {
    await openSignedIn(page, [{ ...FURY, buildDaysAgo: 19 }]);
    await expect(page.getByTestId('bis-lead-sync')).toContainText('Addon · 19 days ago');
    await expect(page.getByTestId('bis-lead-sync')).toContainText(
      'Log out or /reload in the game; the companion then syncs it',
    );
    await expect(lead(page)).toContainText('Gear you wore at the last sync');
  });

  test('with no addon sync the count is one sentence and a link, never a zero', async ({ page }) => {
    await openSignedIn(page, [{ ...FURY, gear: undefined, buildSource: 'blizzard' }]);
    await expect(lead(page)).toContainText('Add the addon to see how many of these slots you already wear.');
    await expect(lead(page).getByRole('link', { name: 'Get the addon →' })).toHaveAttribute('href', '/setup');
    await expect(page.getByTestId('bis-lead-count')).toBeHidden();
    await expect(page.getByTestId('bis-lead-sync')).toBeHidden();
    await expect(page.getByTestId('bis-lead-cta')).toHaveText('Fury gear for 40 to 49 →');
  });

  test('a character under 20 gets one sentence and the first band', async ({ page }) => {
    await openSignedIn(page, [{ ...FURY, level: 12 }]);
    await expect(lead(page)).toContainText('Best in slot starts at 20.');
    await expect(page.getByTestId('bis-lead-count')).toBeHidden();
    await expect(page.getByTestId('bis-lead-cta')).toHaveText('Fury gear for 20 to 29 →');
    await expect(pill(page, 20)).toHaveAttribute('aria-current', 'true');
    await expect(page.locator('[data-band-pill][data-you]')).toHaveCount(0);
  });

  test('a level 60 character reads the plain top band', async ({ page }) => {
    await openSignedIn(page, [{ ...FURY, level: 60, gear: undefined }]);
    await expect(page.getByTestId('bis-lead-cta')).toHaveText('Fury gear for 60 →');
    await expect(pill(page, 60)).toHaveAttribute('aria-current', 'true');
  });

  test('the lead reserves its box while the character resolves', async ({ page }) => {
    await signInWith(page, [FURY]);
    await storePointer(page, armoryPointer(FURY));
    await stubPublishedBis(page, [WARRIOR_SPEC]);
    let release: () => void = () => undefined;
    const held = new Promise<void>((resolve) => {
      release = resolve;
    });
    await page.route('**/v1/me', async (route) => {
      await held;
      await route.fallback();
    });
    await page.goto('/bis');
    await expect(page.getByTestId('bis-lead-skeleton')).toBeVisible();
    const before = await page.getByTestId('bis-class-warrior').boundingBox();
    release();
    await expect(page.getByTestId('bis-lead-name')).toHaveText(FURY.name);
    const after = await page.getByTestId('bis-class-warrior').boundingBox();
    expect(Math.abs((after?.y ?? 0) - (before?.y ?? 0))).toBeLessThanOrEqual(2);
  });
});

test.describe('phone', () => {
  test.use({ viewport: PHONE });

  test('pills are 44 px tall, each label is on one line and nothing scrolls sideways', async ({ page }) => {
    await signedOut(page);
    await page.goto('/bis');
    for (const band of BANDS) {
      const box = await pill(page, band).boundingBox();
      expect(box?.height ?? 0).toBeGreaterThanOrEqual(MIN_TOUCH_TARGET_PX);
      const labelHeight = await pill(page, band)
        .locator('.pb-label')
        .evaluate((node) => node.getBoundingClientRect().height);
      const lineHeight = await pill(page, band)
        .locator('.pb-label')
        .evaluate((node) => parseFloat(getComputedStyle(node).fontSize) * 1.6);
      expect(labelHeight).toBeLessThanOrEqual(lineHeight);
    }
    const { scrollWidth, clientWidth } = await page.evaluate(() => ({
      scrollWidth: document.documentElement.scrollWidth,
      clientWidth: document.documentElement.clientWidth,
    }));
    expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
  });

  test('the character’s band is a gold dot, not a word, and the strip shows the selected caption only', async ({
    page,
  }) => {
    await openSignedIn(page, [FURY]);
    await expect(pill(page, 40).locator('.pb-dot')).toBeVisible();
    await expect(pill(page, 40).locator('.pb-you')).toBeHidden();
    const visibleCaptions = page.locator('[data-testid^="bis-landing-band-caption-"]:visible');
    await expect(visibleCaptions).toHaveCount(1);
    await expect(page.getByTestId('bis-lead-cta')).toBeVisible();
    const cta = await page.getByTestId('bis-lead-cta').boundingBox();
    expect(cta?.height ?? 0).toBeGreaterThanOrEqual(48);
  });

  test('spec rows are at least 44 px tall', async ({ page }) => {
    await signedOut(page);
    await page.goto('/bis');
    const box = await specLink(page, 'warrior', 'arms').boundingBox();
    expect(box?.height ?? 0).toBeGreaterThanOrEqual(MIN_TOUCH_TARGET_PX);
  });
});
