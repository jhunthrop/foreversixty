// web/tests/e2e/guides.spec.ts
// The guide pages in the flat system (design/specs/2026-11-08-guides.md §4, §5): each page has
// one Plate and one h1, a spec guide has the nine sections in SPEC_SECTIONS order, nothing
// scrolls sideways at 390, the leveling slider moves its caption, and the signed-in answer, the
// marked spec row, the Best in slot band and the slider follow the header selector's character
// (B1, B2). Last, spec section 9 (lane C): Load this build lands in the planner with the points
// lit.
//
// Guide pages read their talents, ladder and BiS from data/builds directly, so everything but
// the planner landing is the same under either FOREVER_DATA. The planner itself decodes against
// the published talent file: the fixture's small warrior cannot hold a real-build FS1 code, so
// that one test runs only under FOREVER_DATA=real (`FOREVER_DATA=real E2E_PORT=4387 npx
// playwright test tests/e2e/guides.spec.ts`).
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { openGuidesState } from './support/guides-page';
import {
  armoryPointer,
  keyOf,
  OBNOXIOUS,
  signInWith,
  storePointer,
  type FixtureCharacter,
} from './support/selector';
import { guidesCopy } from '../../src/lib/guides/copy';
import { slugify, SPEC_SECTIONS } from '../../src/lib/guides/sections';

const PHONE = { width: 390, height: 844 };
const DESKTOP = { width: 1440, height: 900 };
const PAGES = ['/guides', '/guides/hunter', '/guides/hunter/survival'] as const;
/** The panels each page draws (directions mocks b2-guides, b2-guide-class, b2-guide-spec): a panel
 *  per class on /guides, a panel per spec on a class landing, the one At a glance on a spec guide. */
const PLATES: Record<(typeof PAGES)[number], number> = {
  '/guides': 9,
  '/guides/hunter': 3,
  '/guides/hunter/survival': 1,
};
const SURVIVAL = '/guides/hunter/survival';

/** A level 34 Survival hunter, the second character on the account (Felsnap's facts). */
const FELSNAP: FixtureCharacter = {
  name: 'Felsnap',
  class: 'Hunter',
  spec: 'Survival',
  race: 'Dwarf',
  level: 34,
  faction: 'alliance',
  buildDaysAgo: 0.01,
};

const answer = (page: Page) => page.getByTestId('guide-answer');
const bisPanel = (page: Page) => page.getByTestId('guide-bis').locator('[data-bis-current]');
const levelingTree = (page: Page) => page.getByTestId('guide-leveling-tree');

/** Scrolls the leveling lane into view and waits for its island (client:visible) to hydrate. */
async function hydratedLevelingTree(page: Page): Promise<void> {
  await levelingTree(page).scrollIntoViewIfNeeded();
  await expect(page.locator('astro-island:has([data-testid="guide-leveling-tree"])')).not.toHaveAttribute(
    'ssr',
    /.*/,
  );
}

/** Picks `character` in the header selector, as a visitor switches. */
async function chooseInSelector(page: Page, character: FixtureCharacter): Promise<void> {
  await page.getByTestId('character-selector').click();
  await page
    .getByTestId('selector-panel')
    .getByTestId(`selector-row-${keyOf(character)}`)
    .getByRole('button')
    .first()
    .click();
  await expect(page.getByTestId('selector-name')).toHaveText(character.name);
}

async function noSidewaysScroll(page: Page): Promise<void> {
  const { scrollWidth, clientWidth } = await page.evaluate(() => ({
    scrollWidth: document.documentElement.scrollWidth,
    clientWidth: document.documentElement.clientWidth,
  }));
  expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
}

test.describe('every guide page', () => {
  for (const path of PAGES) {
    test(`${path} has its panels and one h1, the Marcellus title`, async ({ page }) => {
      await page.goto(path);
      await expect(page.locator('[data-region="plate"]')).toHaveCount(PLATES[path]);
      await expect(page.locator('h1')).toHaveCount(1);
      await expect(page.locator('h1')).toHaveAttribute('data-testid', 'page-title');
    });

    test(`${path} fits 390 without scrolling sideways`, async ({ page }) => {
      await page.setViewportSize(PHONE);
      await page.goto(path);
      await noSidewaysScroll(page);
    });
  }

  test('a spec guide carries the nine section anchors in SPEC_SECTIONS order', async ({ page }) => {
    await page.goto(SURVIVAL);
    const anchors = await page
      .locator('main section[id][data-region^="section-"]')
      .evaluateAll((sections) => sections.map((section) => section.id));
    expect(anchors).toEqual(SPEC_SECTIONS.map(slugify));
  });
});

test.describe('the leveling slider', () => {
  test('moves the leveling caption, one level a step and Home to 10', async ({ page }) => {
    await page.goto(SURVIVAL);
    await hydratedLevelingTree(page);
    const caption = levelingTree(page).getByTestId('guide-leveling-caption');
    await expect(caption).toContainText('Level 60:');
    await expect(caption).not.toContainText('points spent');
    await levelingTree(page).getByTestId('guide-leveling-slider-previous').click();
    await expect(caption).toContainText('Level 59:');
    await expect(caption).toContainText('50 points spent, 1 to go');
    await levelingTree(page).getByTestId('guide-leveling-slider-range').focus();
    await page.keyboard.press('Home');
    await expect(caption).toContainText(/Level 10: .+, rank 1/);
    await expect(caption).toContainText('1 point spent, 50 to go');
  });
});

test.describe('signed in as a level 34 Survival hunter on the Survival guide', () => {
  test.beforeEach(async ({ page }) => {
    await page.setViewportSize(DESKTOP);
    await openGuidesState(page, SURVIVAL, 'survival34');
  });

  test('the answer offers the talents at 34 and names the next point', async ({ page }) => {
    await expect(answer(page)).toHaveAttribute('data-state', 'ready');
    await expect(answer(page).getByTestId('guide-answer-detail')).toHaveText('34 Dwarf Survival');
    const talents = answer(page).getByTestId('guide-answer-link-0');
    await expect(talents).toHaveText(`${guidesCopy.talentsAtLevel(34)} ›`);
    await expect(talents).toHaveAttribute('href', '#leveling');
    await expect(answer(page).getByTestId('guide-answer-link-1')).toHaveAttribute(
      'href',
      '/bis/hunter/survival?band=30',
    );
    await expect(answer(page).getByTestId('guide-answer-next')).toBeVisible();
    await expect(answer(page).getByTestId('guide-answer-next')).toHaveText(/^Next point at 3[5-9]: .+$/);
  });

  test('the leveling slider opens at 34', async ({ page }) => {
    await hydratedLevelingTree(page);
    await expect(levelingTree(page).getByTestId('guide-leveling-slider-range')).toHaveValue('34');
    await expect(levelingTree(page).getByTestId('guide-leveling-slider-range')).toHaveAttribute(
      'aria-valuetext',
      'Level 34',
    );
    await expect(levelingTree(page).getByTestId('guide-leveling-caption')).toContainText('Level 34:');
    await expect(levelingTree(page).getByTestId('guide-leveling-caption')).toContainText(
      '25 points spent, 26 to go',
    );
  });

  test('the Best in slot lane shows band 30 for its faction (a Battle.net character, B2)', async ({
    page,
  }) => {
    await expect(bisPanel(page)).toHaveAttribute('data-panel', '30-alliance');
    await expect(bisPanel(page).locator('[data-band]')).toHaveAttribute('data-band', '30');
  });
});

test.describe('the Best in slot lane follows the header selector', () => {
  test('a switch from the level 34 to a level 60 Survival hunter moves the lane from band 30 back to 60', async ({
    page,
  }) => {
    const capped: FixtureCharacter = { ...FELSNAP, name: 'Thandrel', race: 'Night Elf', level: 60 };
    await signInWith(page, [FELSNAP, capped]);
    await storePointer(page, armoryPointer(FELSNAP));
    await page.goto(SURVIVAL);
    await expect(bisPanel(page)).toHaveAttribute('data-panel', '30-alliance');
    await chooseInSelector(page, capped);
    await expect(answer(page).getByTestId('guide-answer-detail')).toHaveText(/^60 /);
    await expect(bisPanel(page)).toHaveAttribute('data-panel', '60-alliance');
    await expect(bisPanel(page).locator('[data-band]')).toHaveAttribute('data-band', '60');
  });
});

test.describe('signed in as another class', () => {
  test('a level 60 Fire mage on the Survival guide gets its own guide link, and the lanes stay at 60', async ({
    page,
  }) => {
    await openGuidesState(page, SURVIVAL, 'fire60');
    await expect(answer(page).getByTestId('guide-answer-link-0')).toHaveText(
      guidesCopy.signedInCallout('Fire', 'Mage'),
    );
    await expect(answer(page).getByTestId('guide-answer-link-0')).toHaveAttribute(
      'href',
      '/guides/mage/fire',
    );
    await expect(answer(page).getByTestId('guide-answer-next')).toBeHidden();
    await expect(bisPanel(page)).toHaveAttribute('data-panel', /^60-/);
    await hydratedLevelingTree(page);
    await expect(levelingTree(page).getByTestId('guide-leveling-slider-range')).toHaveValue('60');
  });
});

// Bug B1: /guides and the class landing read `me.characters[0]` (here the Fury warrior); they
// follow the header selector's pointer (the Survival hunter) instead, and move when it moves.
test.describe('the answer and the marked spec row follow the header selector (B1)', () => {
  test.beforeEach(async ({ page }) => {
    await signInWith(page, [OBNOXIOUS, FELSNAP]);
    await storePointer(page, armoryPointer(FELSNAP));
  });

  for (const path of ['/guides', '/guides/hunter']) {
    test(`${path} answers for the selected character, not the account's first`, async ({ page }) => {
      await page.goto(path);
      await expect(answer(page)).toHaveAttribute('data-state', 'ready');
      const link = answer(page).getByTestId('guide-answer-link-0');
      await expect(link).toHaveText(guidesCopy.signedInCallout('Survival', 'Hunter'));
      await expect(link).toHaveAttribute('href', SURVIVAL);
      const mine = page.getByTestId('guide-spec-row-hunter/survival');
      await expect(mine).toHaveAttribute('aria-current', 'true');
      await expect(mine.locator('[data-your-spec]')).toBeVisible();
      await expect(page.locator('[data-testid^="guide-spec-row-"][aria-current="true"]')).toHaveCount(1);
      await expect(page.getByTestId('guide-spec-row-warrior/fury')).toHaveCount(path === '/guides' ? 1 : 0);
    });
  }

  test('/guides moves the answer and the mark when the header selector switches, without a reload', async ({
    page,
  }) => {
    await page.goto('/guides');
    await expect(page.getByTestId('guide-spec-row-hunter/survival')).toHaveAttribute('aria-current', 'true');
    await page.evaluate(() => {
      (window as unknown as { __kept: boolean }).__kept = true;
    });
    await chooseInSelector(page, OBNOXIOUS);
    await expect(answer(page).getByTestId('guide-answer-link-0')).toHaveText(
      guidesCopy.signedInCallout('Fury', 'Warrior'),
    );
    await expect(page.getByTestId('guide-spec-row-warrior/fury')).toHaveAttribute('aria-current', 'true');
    await expect(page.getByTestId('guide-spec-row-hunter/survival')).not.toHaveAttribute(
      'aria-current',
      'true',
    );
    expect(await page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept === true)).toBe(true);
  });
});

test.describe('Load this build lands in the planner (FOREVER_DATA=real)', () => {
  test.skip(
    process.env.FOREVER_DATA !== 'real',
    'the fixture talent tree cannot legally hold a real-build FS1 code. Run with FOREVER_DATA=real, e.g. FOREVER_DATA=real E2E_PORT=4387 npx playwright test tests/e2e/guides.spec.ts',
  );

  test('Load this build on the Fury Warrior guide lands in the planner with points already spent', async ({
    page,
  }) => {
    await page.goto('/guides/warrior/fury');
    const loadLink = page.getByTestId('guide-load-build');
    await expect(loadLink).toBeVisible();
    expect(await loadLink.getAttribute('href')).toMatch(/^\/planner\?code=/);

    await loadLink.click();
    await expect(page).toHaveURL(/\/planner\?code=/);

    // Bloodthirst is Fury's own named capstone (tier 6, single rank, tab-order's last talent
    // in the tree) -- fury.md's `build:` FS1 code spends its Fury segment's final, untrimmed
    // digit on it, so it comes back lit at rank 1 once the planner decodes the loaded code
    // against the real talent data this same build id ships.
    const bloodthirstCell = page.locator('[data-testid^="talent-"][aria-label^="Bloodthirst"]');
    await expect(bloodthirstCell).toHaveAttribute('data-rank', '1');
  });
});
