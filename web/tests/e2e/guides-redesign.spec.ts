// web/tests/e2e/guides-redesign.spec.ts
// Acceptance sweep for the guides in the flat system (design/specs/2026-11-08-guides.md §4): the
// /guides class picker (nine crests, 28 spec rows), the class landing's one Plate of spec rows
// (name and role, no date and no blurb), the spec guide's At a glance plate, and the Leveling lane's level
// slider and planner link -- at desktop and phone, with the site's own "nothing scrolls
// sideways" sweep every page-level spec runs (layout.spec.ts's own convention).
import type { Page } from '@playwright/test';
import activeBuild from '../../src/data/active-build.json' with { type: 'json' };
import { buildRacesForClass, raceIconsFor } from '../../src/lib/guides/build-races';
import { SPEC_BY_KEY } from '../../src/lib/sim/specs';
import { hydratedLevelingTree, levelingTree } from './support/guides-page';
import { expect, test } from './support/test';

const ACTIVE_BUILD = activeBuild.build;
const DESKTOP = { width: 1440, height: 900 };
const PHONE = { width: 390, height: 844 };

function iconOfSrc(src: string | null): string | undefined {
  return src?.match(/\/icons\/([a-z0-9_]+)\.webp$/)?.[1];
}

const specRows = (page: Page) => page.locator('[data-testid^="guide-spec-row-"]');

async function noSidewaysScroll(page: Page): Promise<void> {
  const { scrollWidth, clientWidth } = await page.evaluate(() => ({
    scrollWidth: document.documentElement.scrollWidth,
    clientWidth: document.documentElement.clientWidth,
  }));
  expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
}

test.describe('desktop', () => {
  test.use({ viewport: DESKTOP });

  test('/guides renders nine class crests and 28 spec links', async ({ page }) => {
    await page.goto('/guides');
    const picker = page.getByTestId('guides-class-picker');
    await expect(picker.locator('[data-testid^="class-crest-"]')).toHaveCount(9);
    await expect(picker.locator('a[data-testid^="guides-class-link-"]')).toHaveCount(9);
    await expect(picker.locator('a[data-testid^="guide-spec-row-"][href^="/guides/"]')).toHaveCount(28);
  });

  test('/guides/warrior leads with three spec rows that are each the link to their guide, with no figure', async ({
    page,
  }) => {
    await page.goto('/guides/warrior');
    await expect(page.locator('h1')).toHaveText('Warrior');
    await expect(page.getByRole('tablist', { name: 'Spec' })).toHaveCount(0);
    const rows = page.getByTestId('guide-class-spec-cards').locator('a[data-testid^="guide-spec-row-"]');
    await expect(rows).toHaveCount(3);
    const fury = rows.nth(1);
    await expect(fury).toHaveAttribute('href', '/guides/warrior/fury');
    // The row's own name says it (the spec retires the card's "Read the ... guide" aria label):
    // the spec name and its line, all one link and never a nested anchor. The row is the link,
    // so no "Read the guide ›" trail repeats it (h-audit round 8).
    await expect(fury).toContainText('Fury');
    await expect(fury).not.toContainText('Read the guide');
    await expect(fury.locator('a')).toHaveCount(0);
    await expect(page.getByTestId('guide-class-spec-cards')).not.toContainText(/DPS at 60|per second/);
  });

  test('every spec row names its role, with no date, no blurb and no class intro', async ({ page }) => {
    await page.goto('/guides/warrior');
    await expect(page.getByTestId('guide-spec-row-warrior/protection')).toContainText('Tank');
    await expect(page.getByTestId('guide-class-spec-cards')).not.toContainText('Updated');
    await expect(page.getByTestId('guide-class-spec-cards')).not.toContainText(/ is the /);
    await expect(page.getByTestId('guide-class-identity')).toHaveCount(0);
  });

  test('the class facts follow the spec rows under their own headings, with no sourcing caveat', async ({
    page,
  }) => {
    await page.goto('/guides/warrior');
    await expect(page.getByTestId('class-landing-what-changed-from-1-12').locator('h2')).toHaveText(
      'What changed from 1.12',
    );
    await expect(page.getByTestId('class-landing-races').locator('h2')).toHaveText('Races');
    // A floor, not a pin: the guide rounds add changes to the list as they confirm them.
    const changes = page.getByTestId('class-landing-what-changed-from-1-12').locator('li');
    await expect(changes.first()).toBeVisible();
    expect(await changes.count()).toBeGreaterThanOrEqual(8);
    await expect(page.getByTestId('confidence-footer-note')).toHaveCount(0);
  });

  test('the Races panel lists every race the class can be, each with its client portrait', async ({
    page,
  }) => {
    await page.goto('/guides/warrior');
    const races = buildRacesForClass(ACTIVE_BUILD, 'warrior');
    const list = page.getByTestId('guide-class-races');
    await expect(list.locator('li')).toHaveCount(races.length);
    const portraits = raceIconsFor(ACTIVE_BUILD);
    for (const race of races) {
      const row = list.getByTestId(`guide-class-race-${race.slug}`);
      await expect(row).toContainText(race.name);
      await expect(row.locator(`[data-race-icon="${race.name}"]`)).toHaveCount(
        portraits.has(race.slug) ? 1 : 0,
      );
    }
  });

  test('druid gives Feral and Feral Bear their own form icons from the spec-icon map', async ({ page }) => {
    await page.goto('/guides/druid');
    const iconOf = async (spec: string) =>
      iconOfSrc(await page.getByTestId(`guide-spec-row-druid/${spec}`).locator('img').getAttribute('src'));
    const cat = await iconOf('feral');
    const bear = await iconOf('feral-bear');
    expect(cat).toBe(SPEC_BY_KEY['druid-feral'].icon);
    expect(bear).toBe(SPEC_BY_KEY['druid-feral-bear'].icon);
    expect(cat).not.toBe(bear);
  });

  // The plate is the build alone (h-audit 2026-11-09): the Rotation and Stat priority sections
  // below own their lists, so the plate never repeats them.
  test('/guides/warrior/fury At a glance is the build with Load and Sim, no rotation or stats', async ({
    page,
  }) => {
    await page.goto('/guides/warrior/fury');
    await expect(page.getByTestId('guide-glance-build')).toBeVisible();
    await expect(page.getByTestId('guide-load-build')).toBeVisible();
    await expect(page.getByTestId('guide-sim-build')).toBeVisible();
    await expect(page.getByTestId('guide-glance-rotation')).toHaveCount(0);
    await expect(page.getByTestId('guide-glance-stats')).toHaveCount(0);
  });

  // The rail's summary Load and the full tree's Load were two buttons that needed distinct
  // names; the spec retires the rail (G1, "one Load per page now"): the plate's Load this build
  // is the page's one primary, and each build row in Talents and builds carries a plain link.
  test('the page has one Load this build button, the plate’s primary', async ({ page }) => {
    await page.goto('/guides/warrior/fury');
    await expect(page.getByTestId('guide-load-build')).toHaveCount(1);
    await expect(page.getByTestId('guide-load-build')).toHaveAccessibleName('Load this build');
    await expect(page.getByTestId('guide-glance-build').getByTestId('guide-load-build')).toHaveCount(1);
  });

  test('the leveling tree walks 10 to 60 from the keyboard and the buttons', async ({ page }) => {
    await page.goto('/guides/warrior/fury');
    await hydratedLevelingTree(page);
    const tree = levelingTree(page);
    const slider = tree.getByTestId('guide-leveling-slider-range');
    await expect(slider).toBeEnabled();
    await expect(tree.getByTestId('guide-leveling-trees')).toBeVisible();
    const caption = tree.getByTestId('guide-leveling-caption');
    await expect(caption).not.toContainText('points spent');
    await expect(tree.getByTestId('guide-leveling-slider-next')).toBeDisabled();
    // Home is 10, PageUp ten levels on, End is 60.
    await slider.focus();
    await page.keyboard.press('Home');
    await expect(caption).toContainText(/Level 10: .+, rank 1/);
    await expect(caption).toContainText('1 point spent, 50 to go');
    await expect(tree.getByTestId('guide-leveling-slider-previous')).toBeDisabled();
    await page.keyboard.press('PageUp');
    await expect(caption).toContainText('Level 20:');
    await tree.getByTestId('guide-leveling-slider-next').click();
    await expect(caption).toContainText('12 points spent, 39 to go');
    await expect(tree.getByTestId('guide-leveling-load')).toHaveAccessibleName(
      'Load the level 21 build in the planner',
    );
    await slider.focus();
    await page.keyboard.press('End');
    await expect(caption).not.toContainText('points spent');
  });

  test('the leveling tree’s Load in planner link carries the chosen level’s talent string', async ({
    page,
  }) => {
    await page.goto('/guides/warrior/fury');
    await hydratedLevelingTree(page);
    const tree = levelingTree(page);
    const at60 = await tree.getByTestId('guide-leveling-load').getAttribute('href');
    expect(at60).toMatch(/^\/planner\?spec=warrior-fury&talents=/);
    await tree.getByTestId('guide-leveling-slider-previous').click();
    await expect(tree.getByTestId('guide-leveling-load')).not.toHaveAttribute('href', at60 ?? '');
    const at59 = await tree.getByTestId('guide-leveling-load').getAttribute('href');
    expect(at59).toMatch(/^\/planner\?spec=warrior-fury&talents=/);
  });
});

test.describe('phone', () => {
  test.use({ viewport: PHONE });

  test('/guides/warrior/fury At a glance fits the phone with no sideways scroll', async ({ page }) => {
    await page.goto('/guides/warrior/fury');
    await expect(page.getByTestId('guide-glance')).toBeVisible();
    await noSidewaysScroll(page);
  });

  test('the leveling tree keeps its control and link on phone, with no sideways scroll', async ({ page }) => {
    await page.goto('/guides/warrior/fury');
    await hydratedLevelingTree(page);
    const tree = levelingTree(page);
    await expect(tree.getByTestId('guide-leveling-slider-range')).toBeEnabled();
    await expect(tree.getByTestId('guide-leveling-slider-previous')).toBeVisible();
    await expect(tree.getByTestId('guide-leveling-slider-next')).toBeVisible();
    await expect(tree.getByTestId('guide-leveling-load')).toBeVisible();
    await expect(tree.getByTestId('guide-leveling-trees')).toBeVisible();
    await noSidewaysScroll(page);
  });
});

const CLASS_SLUGS = ['druid', 'hunter', 'mage', 'paladin', 'priest', 'rogue', 'shaman', 'warlock', 'warrior'];
/** The least room between the spec rows' plate and the class notes below them. */
const MIN_GAP_BELOW_CARDS = 16;

test.describe('class landings', () => {
  for (const [label, viewport] of [
    ['390', PHONE],
    ['1440', DESKTOP],
  ] as const) {
    test.describe(label, () => {
      test.use({ viewport });

      for (const classSlug of CLASS_SLUGS) {
        test(`/guides/${classSlug} keeps every spec row on screen and the notes clear of the rows`, async ({
          page,
        }) => {
          await page.goto(`/guides/${classSlug}`);
          const clientWidth = await page.evaluate(() => document.documentElement.clientWidth);
          const rows = await specRows(page).all();
          expect(rows.length).toBeGreaterThanOrEqual(3);
          for (const row of rows) {
            const box = await row.boundingBox();
            expect(box).not.toBeNull();
            expect(box!.x).toBeGreaterThanOrEqual(0);
            expect(box!.x + box!.width).toBeLessThanOrEqual(clientWidth);
          }
          const plate = await page.getByTestId('guide-class-spec-cards').boundingBox();
          const notes = await page.locator('[data-testid^="class-landing-"]').first().boundingBox();
          expect(notes!.y - (plate!.y + plate!.height)).toBeGreaterThanOrEqual(MIN_GAP_BELOW_CARDS);
        });
      }
    });
  }

  test.describe('1440 row', () => {
    test.use({ viewport: DESKTOP });

    test('/guides/druid sets its four spec rows in one row', async ({ page }) => {
      await page.goto('/guides/druid');
      const tops = await specRows(page).evaluateAll((rows) =>
        rows.map((row) => Math.round(row.getBoundingClientRect().top)),
      );
      expect(tops).toHaveLength(4);
      expect(new Set(tops).size).toBe(1);
    });
  });
});

test.describe('no horizontal overflow', () => {
  for (const [label, viewport] of [
    ['390', PHONE],
    ['1440', DESKTOP],
  ] as const) {
    test.describe(label, () => {
      test.use({ viewport });

      for (const path of ['/guides', '/guides/warrior', '/guides/druid', '/guides/warrior/fury']) {
        test(`${path} fits the viewport without scrolling sideways`, async ({ page }) => {
          await page.goto(path);
          await noSidewaysScroll(page);
        });
      }
    });
  }
});
