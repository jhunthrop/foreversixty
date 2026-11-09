// web/tests/e2e/bis-landing.spec.ts
// /bis in the flat system (design/specs/2026-11-07-bis-landing.md): the head, the level range bar
// (the character's own range by default, owner 2026-11-07), the summary line and the class plate,
// signed out and signed in with the header selector stubbed. Layout is the parity spec's job
// (parity-bis-landing.spec.ts); this spec pins the behaviour.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
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
import { ctaCopy } from '../../src/lib/addon/cta-copy';

const DESKTOP = { width: 1440, height: 900 };
const PHONE = { width: 390, height: 844 };
const BANDS = [20, 30, 40, 50, 60];
const MIN_TOUCH_TARGET_PX = 44;

const WARRIOR_SPEC = 'warrior-fury';
const warriorBand = bisBand(WARRIOR_SPEC, 'alliance', 40, 'bare');
const sourcedPicks = (warriorBand.slots ?? []).filter((slot) => typeof slot.item_id === 'number');
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

const specRow = (page: Page, classSlug: string, specSlug: string) =>
  page.getByTestId(`bis-spec-row-${classSlug}/${specSlug}`);
const band = (page: Page, value: number) => page.getByTestId(`control-band-${value}`);
const answer = (page: Page) => page.getByTestId('bis-answer');
const action = (page: Page) => page.getByTestId('bis-answer-button');

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
  await expect(answer(page)).toHaveAttribute('data-state', 'ready');
}

async function choose(page: Page, character: FixtureCharacter): Promise<void> {
  await page.evaluate(() => {
    (window as unknown as { __kept: boolean }).__kept = true;
  });
  await page.getByTestId('character-selector').click();
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

  test('one Marcellus title, the levels and one Updated stamp, the tier list link, no answer', async ({
    page,
  }) => {
    const title = page.getByTestId('page-title');
    await expect(title).toHaveText('Best in slot');
    await expect(title).toHaveCSS('font-family', /Marcellus/);
    await expect(title).toHaveCSS('text-transform', 'none');
    await expect(page.locator('h1')).toHaveCount(1);
    await expect(page.getByTestId('bis-landing-updated')).toHaveCount(1);
    await expect(page.getByTestId('bis-index-tier-link')).toHaveAttribute('href', '/tiers');
    await expect(answer(page)).toBeHidden();
  });

  test('the bar defaults to 20 to 29 and every spec row opens at 20 to 29 with no faction', async ({
    page,
  }) => {
    await expect(band(page, 20)).toHaveAttribute('aria-checked', 'true');
    const rows = page.locator('a[data-testid^="bis-spec-row-"]');
    await expect(rows).toHaveCount(readSpecCatalog().length);
    for (const row of await rows.all()) {
      await expect(row).toHaveAttribute('href', /^\/bis\/[a-z]+\/[a-z-]+$/);
    }
  });

  test('a range rewrites every spec link and the summary, keeps the choice in the URL and in storage', async ({
    page,
  }) => {
    await band(page, 40).click();
    await expect(band(page, 40)).toHaveAttribute('aria-checked', 'true');
    await expect(page).toHaveURL(/\?band=40$/);
    await expect(specRow(page, 'warrior', 'fury')).toHaveAttribute('href', '/bis/warrior/fury/40-49');
    await expect(page.locator('[data-band-text]:visible')).toHaveCount(1);
    await expect(page.locator('[data-band-text="40"]')).toBeVisible();
    await page.goto('/bis');
    await expect(band(page, 40)).toHaveAttribute('aria-checked', 'true');
  });

  test('every range has its summary; ?band= opens on it', async ({ page }) => {
    for (const value of BANDS) {
      await page.goto(`/bis?band=${value}`);
      await expect(band(page, value)).toHaveAttribute('aria-checked', 'true');
      await expect(page.locator(`[data-band-text="${value}"]`)).toContainText('Top sources:');
    }
  });

  test('tanks and healers carry their role tag; every class has a crest and every spec an icon', async ({
    page,
  }) => {
    await expect(specRow(page, 'warrior', 'protection')).toContainText('Tank');
    await expect(specRow(page, 'priest', 'holy')).toContainText('Healer');
    await expect(specRow(page, 'warrior', 'fury')).not.toContainText(/Tank|Healer/);
    await expect(
      page.locator('[data-testid^="bis-class-"] img.crest, [data-testid^="bis-class-"] img[src*="crests"]'),
    ).toHaveCount(9);
    await expect(page.locator('a[data-testid^="bis-spec-row-"] img')).toHaveCount(readSpecCatalog().length);
  });
});

test.describe('signed in, desktop', () => {
  test.use({ viewport: DESKTOP });

  test('the answer names the character, counts its range and has one sentence-case button', async ({
    page,
  }) => {
    await openSignedIn(page, [FURY]);
    await expect(answer(page)).toContainText(FURY.name);
    await expect(page.getByTestId('bis-answer-figure')).toContainText(`${WORN_COUNT}`);
    await expect(page.getByTestId('bis-answer-figure')).toContainText(`of ${sourcedPicks.length} slots`);
    await expect(action(page)).toHaveText('Fury gear for 40 to 49');
    await expect(action(page)).toHaveCSS('text-transform', 'none');
    await expect(action(page)).toHaveAttribute('href', '/bis/warrior/fury/40-49');
  });

  test('the range follows the character, even over a stored pick, and its spec row is Yours', async ({
    page,
  }) => {
    await page.addInitScript(() => window.localStorage.setItem('fs.bisBand', '60'));
    await openSignedIn(page, [FURY]);
    await expect(band(page, 40)).toHaveAttribute('aria-checked', 'true');
    await expect(specRow(page, 'warrior', 'fury')).toHaveAttribute('aria-current', 'true');
    await expect(page.locator('[data-your-spec]:visible')).toHaveCount(1);
    await expect(specRow(page, 'mage', 'frost')).toHaveAttribute('href', '/bis/mage/frost/40-49');
  });

  test('another range keeps the answer and offers the way back to the own range', async ({ page }) => {
    await openSignedIn(page, [FURY]);
    await band(page, 60).click();
    await expect(action(page)).toHaveText('Fury gear for 40 to 49');
    const own = page.locator('[data-your-range]');
    await expect(own).toHaveText('Your range: 40 to 49 ›');
    await own.click();
    await expect(band(page, 40)).toHaveAttribute('aria-checked', 'true');
    await expect(own).toBeHidden();
  });

  test('a header switch moves the answer, the range and Yours in place, with no reload', async ({ page }) => {
    await openSignedIn(page, [FURY, HUNTER], [WARRIOR_SPEC, 'hunter-marksmanship']);
    await choose(page, HUNTER);
    await expect(answer(page)).toContainText('Quickshot');
    await expect(action(page)).toHaveText('Marksmanship gear for 30 to 39');
    await expect(action(page)).toHaveAttribute('href', '/bis/hunter/marksmanship/30-39?faction=horde');
    await expect(band(page, 30)).toHaveAttribute('aria-checked', 'true');
    await expect(specRow(page, 'hunter', 'marksmanship')).toHaveAttribute('aria-current', 'true');
    await expect(specRow(page, 'warrior', 'fury')).not.toHaveAttribute('aria-current', 'true');
    expect(await page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept)).toBe(true);
  });

  test('a stale sync relabels the count', async ({ page }) => {
    await openSignedIn(page, [{ ...FURY, buildDaysAgo: 19 }]);
    await expect(page.getByTestId('bis-answer-figure')).toContainText('Pieces you wore at the last sync');
  });

  test('with no gear on file the count is one sentence and a link, never a zero', async ({ page }) => {
    await openSignedIn(page, [{ ...FURY, gear: undefined, buildSource: 'blizzard' }]);
    await expect(answer(page)).toContainText(
      'Add the addon to see how many of these slots you already wear.',
    );
    await expect(
      answer(page).getByRole('link', { name: `${ctaCopy.getAddonAndCompanion} ›` }),
    ).toHaveAttribute('href', '/setup');
    await expect(page.getByTestId('bis-answer-figure')).toBeHidden();
    await expect(action(page)).toHaveText('Fury gear for 40 to 49');
  });

  test('a character under 20 gets one sentence and the first range', async ({ page }) => {
    await openSignedIn(page, [{ ...FURY, level: 12 }]);
    await expect(answer(page)).toContainText('Best in slot starts at 20.');
    await expect(action(page)).toHaveText('Fury gear for 20 to 29');
    await expect(band(page, 20)).toHaveAttribute('aria-checked', 'true');
  });

  test('a level 60 character reads the plain top range', async ({ page }) => {
    await openSignedIn(page, [{ ...FURY, level: 60, gear: undefined }]);
    await expect(action(page)).toHaveText('Fury gear for 60');
    await expect(band(page, 60)).toHaveAttribute('aria-checked', 'true');
  });
});

test.describe('phone', () => {
  test.use({ viewport: PHONE });

  test('nothing scrolls sideways and the spec rows are at least 44 px tall', async ({ page }) => {
    await signedOut(page);
    await page.goto('/bis');
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(PHONE.width);
    for (const row of (await page.locator('a[data-testid^="bis-spec-row-"]').all()).slice(0, 6)) {
      expect((await row.boundingBox())!.height).toBeGreaterThanOrEqual(MIN_TOUCH_TARGET_PX);
    }
  });
});
