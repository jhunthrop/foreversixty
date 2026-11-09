// web/tests/e2e/guides-sections.spec.ts
// W.18, in the flat system (design/specs/2026-11-08-guides.md §4.3): the spec guide reads like a
// class guide. This walks the Frost Mage page section by section at desktop (1280) and phone
// (390) widths: the header facts, the Overview blocks, the builds and the key talents, the
// rotation lists, the stat lane beside the plate's ranked stats, the one band of Best in slot
// picks, the enchant and consumable rows, the racials and the leveling lane, and checks the page
// never scrolls sideways. Then the contents: the rail from 1760, the On this page select below.
//
// The content comes from data/builds and the BiS ranker, not the planner fixture, so it is the
// same under either FOREVER_DATA; only the check that icon files actually load needs the real
// icon tree and so runs under FOREVER_DATA=real.
import { readFileSync } from 'node:fs';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { choose } from './support/select';
import { guidesCopy } from '../../src/lib/guides/copy';
import { raceNameFor } from '../../src/lib/guides/race-table';
import { slugify, SPEC_SECTIONS } from '../../src/lib/guides/sections';
import { ACTIVE_BUILD } from './support/active-build';
import { bisBand } from './support/bis-file';

const DESKTOP = { width: 1280, height: 900 };
const WIDE = { width: 1920, height: 1000 };
const PHONE = { width: 390, height: 844 };
const PAGE = '/guides/mage/frost';
const BIS_SPEC = 'mage-frost';
/** The race the ranker played the Horde 30 band as: every nightly re-rank may move it. */
const HORDE_30_RACE = raceNameFor(ACTIVE_BUILD, bisBand(BIS_SPEC, 'horde', 30, 'bare').race ?? '');
/** Every gear slot the BiS lane lists, one ItemRow each (an empty slot keeps its row). */
const BIS_SLOTS = 17;

/** The main-hand enchant the ranker put on the level 60 raid pick, from the build's enchants.json:
 *  the guide's lane must name the same one (a re-rank onto a two-hander moves it). */
function mainHandEnchant(): { name: string; spellPower: number } {
  const slot = bisBand(BIS_SPEC, 'alliance', 60, 'raid').slots?.find((row) => row.slot === 'main_hand');
  const records = JSON.parse(
    readFileSync(new URL(`../../../data/builds/${ACTIVE_BUILD}/enchants.json`, import.meta.url), 'utf8'),
  ) as { id: number; name: string; stats: Record<string, number> }[];
  const record = records.find((row) => row.id === slot?.enchant_id);
  if (record?.stats.spell_power === undefined)
    throw new Error('the Frost 60 main hand carries no spell power enchant');
  return { name: record.name, spellPower: record.stats.spell_power };
}

const SECTION_IDS = SPEC_SECTIONS.map(slugify);

const laneLabel = (page: Page, testid: string) => page.getByTestId(testid).locator('.ui-lane-head h3');
const bisPanel = (page: Page) => page.getByTestId('guide-bis').locator('[data-bis-current]');

async function noSidewaysScroll(page: Page): Promise<void> {
  const { scrollWidth, clientWidth } = await page.evaluate(() => ({
    scrollWidth: document.documentElement.scrollWidth,
    clientWidth: document.documentElement.clientWidth,
  }));
  expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
}

async function imagesLoaded(page: Page, selector: string): Promise<void> {
  if (process.env.FOREVER_DATA !== 'real') return;
  const broken = await page.evaluate(
    (css) =>
      [...document.querySelectorAll<HTMLImageElement>(css)]
        .filter((img) => img.complete && img.naturalWidth === 0)
        .map((img) => img.src),
    selector,
  );
  expect(broken).toEqual([]);
}

/** The names of a StatTable's rows whose weight is a figure (not "Not on any item" and the like). */
async function rankedStatNames(page: Page, rows: string): Promise<string[]> {
  return page
    .locator(rows)
    .evaluateAll((els) =>
      els
        .filter((row) => row.querySelector<HTMLElement>('.ui-st-f')?.dataset.muted !== 'true')
        .map((row) => row.querySelector('.ui-st-name')?.textContent?.trim() ?? ''),
    );
}

/** Waits until the ControlBar's script has wired the bar (after first paint). */
async function barBooted(page: Page): Promise<void> {
  await page.locator('[data-testid="guide-bar"][data-booted="true"]').waitFor({ state: 'attached' });
}

for (const [label, viewport] of [
  ['1280', DESKTOP],
  ['390', PHONE],
] as const) {
  test.describe(`Frost Mage guide at ${label}`, () => {
    test.use({ viewport });

    test('has the nine sections in order, and the header facts', async ({ page }) => {
      await page.goto(PAGE);
      const ids = await page.locator('main section[id]').evaluateAll((els) => els.map((el) => el.id));
      expect(ids.filter((id) => SECTION_IDS.includes(id))).toEqual(SECTION_IDS);

      await expect(page.getByTestId('guide-status-note')).toHaveCount(0);
      await expect(page.getByTestId('content-footnote')).toContainText('Numbers from client build');
      await expect(page.locator('[data-region="head-meta"]')).toContainText(
        /DPS · Updated \d{1,2} [A-Z][a-z]{2}/,
      );
      await expect(page.locator('.ui-page-header-lead')).toContainText(
        'Frost casts Frostbolt behind Ice Barrier',
      );
      await noSidewaysScroll(page);
    });

    test('lays out the Overview as strengths, weaknesses and the raid role', async ({ page }) => {
      await page.goto(PAGE);
      await expect(page.getByTestId('guide-strengths').locator('li')).toHaveCount(4);
      await expect(page.getByTestId('guide-weaknesses').locator('li')).toHaveCount(3);
      await expect(page.getByTestId('guide-overview-raid')).toContainText('Role:');
      await expect(page.getByTestId('guide-overview-raid')).toContainText('Brings:');
    });

    test('shows the one build, the key talents with icons, and a raid-build diff where one differs', async ({
      page,
    }) => {
      // Arcane's raid build differs from its level 60 build; Frost plays one build for both.
      await page.goto('/guides/mage/arcane');
      await expect(page.getByTestId('guide-build-leveling')).toBeVisible();
      await expect(page.getByTestId('guide-build-raid')).toBeVisible();
      await expect(page.getByTestId('guide-raid-diff')).toContainText(
        /^The raid build moves \d+ points from /,
      );
      await expect(page.getByTestId('guide-build-sim-raid')).toBeVisible();
      await page.goto(PAGE);
      await expect(page.getByTestId('guide-build-both')).toBeVisible();
      await expect(page.getByTestId('guide-load-build')).toBeVisible();
      await expect(laneLabel(page, 'guide-key-talents')).toHaveText('Key talents');
      const first = page.getByTestId('guide-key-talents').locator('.ui-row').first();
      await expect(first).toContainText('Ice Barrier');
      await expect(first.locator('img')).toHaveCount(1);
      await imagesLoaded(page, '[data-testid="guide-key-talents"] img');
    });

    test('numbers a priority with spell icons, lists real cooldowns, and opens on the pre-pull Ice Barrier', async ({
      page,
    }) => {
      await page.goto(PAGE);
      await expect(laneLabel(page, 'guide-rotation-priority')).toHaveText('Priority');
      const rows = page.getByTestId('guide-rotation-priority-row');
      // The order is the nightly APL's (Ice Barrier upkeep leads it today), so the walk pins
      // that Ice Lance is on the list with its icon, never which line comes first.
      const iceLance = rows.filter({ hasText: 'Ice Lance' });
      await expect(iceLance).toHaveCount(1);
      await expect(iceLance.locator('img')).toHaveCount(1);
      await expect(rows.first().locator('img')).toHaveCount(1);
      await expect(rows.first()).toContainText(/^1/);
      await expect(page.getByTestId('guide-rotation-opener-row').first()).toContainText('Ice Barrier');
      await expect(laneLabel(page, 'guide-rotation-cooldowns')).toHaveText('Cooldowns');
      const cooldowns = page.getByTestId('guide-rotation-cooldowns');
      for (const name of ['Evocation', 'Cold Snap', 'Ice Block']) await expect(cooldowns).toContainText(name);
      await imagesLoaded(page, '[data-testid^="guide-rotation-"] img');
    });

    // The stat pill row retires (G1): the plate's Stat priority column is the ranked summary now,
    // and it ranks the same stats, in the same order, as the lane's weights table.
    test('gives the stat lane the hit headline and the weights table, ranked as the plate', async ({
      page,
    }) => {
      await page.goto(PAGE);
      await expect(page.getByTestId('guide-glance-stats')).toBeVisible();
      await expect(page.getByTestId('guide-stat-table')).toBeVisible();
      // The cap, then each source the band's figure counts on the same line; the unit test in
      // src/pages/guides/_spec.test.ts pins the exact line.
      await expect(page.getByTestId('guide-stat-table-hit-cap')).toHaveText(
        /^Hit cap: \d+(\.\d)?%(?: · (?:\d+(\.\d)?% hit from BiS gear and buffs, )?(?:\d+(\.\d)?% from [^,]+, )*\d+(\.\d)?% still to find)?$/,
      );
      await expect(page.getByTestId('guide-stat-table')).toContainText('DPS per point');
      const plate = await rankedStatNames(page, '[data-testid="guide-glance-stats"] .ui-st-r');
      const table = await rankedStatNames(page, '[data-testid^="guide-stat-table-row-"]');
      expect(plate.length).toBeGreaterThan(0);
      expect(plate).toEqual(table);
    });

    // The level and faction pills retire (G5): the guide shows one band, the level 60 picks for
    // the guide's first race's faction, and links the full list; every other band waits in a
    // template for a signed-in character of this class.
    test('opens best in slot on the level 60 Alliance picks, one row a slot, with the full list linked', async ({
      page,
    }) => {
      await page.goto(PAGE);
      await expect(bisPanel(page)).toHaveAttribute('data-panel', '60-alliance');
      await expect(bisPanel(page).locator('.ui-lane-head h3')).toHaveText('Level 60');
      await expect(bisPanel(page).locator('.ui-item-row[data-testid^="guide-bis-"]')).toHaveCount(BIS_SLOTS);
      await expect(bisPanel(page).getByTestId('guide-bis-head')).toContainText('Head');
      await expect(bisPanel(page).getByTestId('guide-bis-head').locator('img')).toHaveCount(1);
      await expect(page.getByTestId('guide-bis-full-list')).toHaveAttribute('href', '/bis/mage/frost/60');
      await expect(page.getByTestId('guide-bis-full-list')).toHaveText(`${guidesCopy.fullListLink} ›`);
      const horde30 = await page
        .locator('template[data-bis-panel="30-horde"]')
        .evaluate((template) => (template as HTMLTemplateElement).content.textContent ?? '');
      expect(horde30).toContain(guidesCopy.bisCaption(HORDE_30_RACE, 'Horde'));
      await imagesLoaded(page, '[data-testid="guide-bis"] img');
      await noSidewaysScroll(page);
    });

    // A pasted export (an addon or code pointer, signed out) still picks its own band: bis-lane.ts
    // reads its spent points (ownBand).
    test('opens best in slot on a pasted character’s own level range when signed out', async ({ page }) => {
      await page.addInitScript(() => {
        // 24 talent points spent: level 33, the 30 to 39 range.
        localStorage.setItem(
          'fs.currentCharacter',
          JSON.stringify({
            source: 'code',
            ref: 'FS1:1.60.1.70532:mage:gnome:55/55/4:',
            label: 'Test',
            classSlug: 'mage',
            savedAt: new Date().toISOString(),
          }),
        );
      });
      await page.goto(PAGE);
      await expect(bisPanel(page)).toHaveAttribute('data-panel', '30-alliance');
      await expect(bisPanel(page).locator('.ui-lane-head h3')).toHaveText(guidesCopy.levelRangeLabel(30));
    });

    test('lists enchants and consumables from the data, only for the slots it can verify', async ({
      page,
    }) => {
      await page.goto(PAGE);
      await expect(laneLabel(page, 'guide-enchants')).toHaveText('Enchants');
      const enchant = mainHandEnchant();
      const mainHand = page.getByTestId('guide-enchant-main_hand');
      await expect(mainHand).toContainText(enchant.name);
      await expect(mainHand).toContainText(`+${enchant.spellPower} spell power`);
      await expect(laneLabel(page, 'guide-consumables')).toHaveText('Consumables');
      await expect(page.getByTestId('guide-consumables')).toContainText('Flask of Supreme Power');
      await expect(page.getByTestId('guide-consumables')).toContainText('Elixir of Frost Power');
    });

    // The race pill row retires (G1): the Racials lane is the one list, the recommended races
    // tagged in it.
    test('has a racials lane with the recommended races marked', async ({ page }) => {
      await page.goto(PAGE);
      const lane = page.getByTestId('guide-race-table');
      await expect(lane.locator('[data-testid^="guide-race-"]')).toHaveCount(6);
      const gnome = page.getByTestId('guide-race-gnome');
      await expect(gnome).toContainText('Recommended');
      await expect(gnome).toContainText('Eureka!');
      await expect(gnome).toContainText('Expansive Mind');
      // The prose names both picks, so no recommended-race line is said again under the lane.
      await expect(lane.locator(':scope > p')).toHaveCount(0);
      await expect(page.getByTestId('guide-rotation-cooldowns')).toContainText('Eureka! (Gnome)');
      await expect(page.getByTestId('guide-race-troll')).toContainText('Berserking');
      await expect(page.getByTestId('guide-race-human')).not.toContainText('Recommended');
    });

    test('renders the leveling lane at level 60 with its planner link, before the island loads', async ({
      page,
    }) => {
      await page.goto(PAGE);
      const tree = page.getByTestId('guide-leveling-tree');
      await expect(tree.getByTestId('guide-leveling-caption')).toContainText('51 points spent, all placed');
      await expect(tree.getByTestId('guide-leveling-caption')).toContainText(/Level 60: .+, rank \d/);
      await expect(tree.getByTestId('guide-leveling-load')).toHaveAttribute(
        'href',
        /^\/planner\?spec=mage-frost&talents=[\d-]+$/,
      );
      await expect(tree.getByTestId('guide-leveling-slider-range')).toHaveAttribute(
        'aria-valuetext',
        'Level 60',
      );
      await expect(tree.getByTestId('guide-leveling-trees')).toBeVisible();
      await noSidewaysScroll(page);
    });

    // The lane now sits beside the prose from 1280 and before it below (G3, G4), so the old
    // "tree between Levels 1 to 19 and Levels 20 to 60" order retires; the prose keeps its own.
    test('#leveling opens on the Leveling section, Levels 1 to 19 before Levels 20 to 60', async ({
      page,
    }) => {
      await page.goto(`${PAGE}#leveling`);
      await expect(page.locator('#leveling h2')).toBeInViewport();
      const top = async (selector: string): Promise<number> =>
        (await page.locator(selector).boundingBox())?.y ?? Number.NaN;
      expect(await top('#levels-1-to-19')).toBeLessThan(await top('#levels-20-to-60'));
    });
  });
}

test.describe('the plate', () => {
  test('phone: Rotation and Stat priority stack under the Build column with no disclosure', async ({
    page,
  }) => {
    await page.setViewportSize(PHONE);
    await page.goto(PAGE);
    await expect(page.getByTestId('guide-glance').locator('details, [aria-expanded]')).toHaveCount(0);
    await expect(page.getByTestId('guide-glance-rotation')).toBeVisible();
    await expect(page.getByTestId('guide-glance-stats')).toBeVisible();
  });

  test('desktop: the plate shows all three columns', async ({ page }) => {
    await page.setViewportSize(DESKTOP);
    await page.goto(PAGE);
    await expect(page.getByTestId('guide-glance-build')).toBeVisible();
    await expect(page.getByTestId('guide-glance-rotation')).toBeVisible();
    await expect(page.getByTestId('guide-glance-stats')).toBeVisible();
  });
});

test.describe('table of contents', () => {
  test('from 1760: the contents rail jumps to a section, and the select is gone', async ({ page }) => {
    await page.setViewportSize(WIDE);
    await page.goto(PAGE);
    const rail = page.getByTestId('guide-toc');
    await expect(rail).toBeVisible();
    await expect(page.getByTestId('guide-toc-select')).toBeHidden();
    await rail.getByRole('link', { name: 'Best in slot' }).click();
    await expect(page).toHaveURL(/#best-in-slot$/);
    await expect(page.locator('#best-in-slot')).toBeInViewport();
  });

  test('from 1760: the rail marks the section you are in', async ({ page }) => {
    await page.setViewportSize(WIDE);
    await page.goto(PAGE);
    const rail = page.getByTestId('guide-toc');
    const scrollTo = (id: string) =>
      page.evaluate(
        (target) =>
          window.scrollTo(
            0,
            (document.getElementById(target)?.getBoundingClientRect().top ?? 0) + window.scrollY - 80,
          ),
        id,
      );
    await scrollTo('stat-priority');
    await expect(rail.locator('[aria-current="location"]')).toHaveText('Stat priority');
    await scrollTo('races');
    await expect(rail.locator('[aria-current="location"]')).toHaveText('Races');
    await expect(rail.locator('[aria-current="location"]')).toHaveCount(1);
  });

  test('below 1760: the rail is hidden and the bar’s On this page select jumps to a section', async ({
    page,
  }) => {
    await page.setViewportSize(DESKTOP);
    await page.goto(PAGE);
    await expect(page.getByTestId('guide-toc')).toBeHidden();
    await expect(page.getByTestId('guide-toc-select')).toBeVisible();
    await barBooted(page);
    await choose(page.locator('#guide-toc'), 'best-in-slot');
    await expect(page).toHaveURL(/#best-in-slot$/);
    await expect(page.locator('#best-in-slot')).toBeInViewport();
  });

  test('phone: the contents open from the bar and jump to a section', async ({ page }) => {
    await page.setViewportSize(PHONE);
    await page.goto(PAGE);
    await expect(page.getByTestId('guide-toc')).toBeHidden();
    await barBooted(page);
    const toggle = page.getByTestId('guide-bar').locator('[data-control-bar-toggle]');
    await expect(toggle).toContainText(`Frost Mage · ${guidesCopy.tocBarLabel}`);
    await toggle.click();
    await expect(toggle).toHaveAttribute('aria-expanded', 'true');
    // Open, the bar lists the sections as rows (the mock's open phone bar), not the select.
    await expect(page.getByTestId('guide-toc-select')).toBeHidden();
    await expect(page.getByTestId('guide-toc-phone')).toBeVisible();
    await page.getByTestId('guide-toc-phone-races').click();
    await expect(page).toHaveURL(/#races$/);
    await expect(page.locator('#races')).toBeInViewport();
  });

  // Round 19 (integrator, wow-player): the jump list stayed open after a tap, so its rows
  // covered the section the player jumped to. A jump closes it ("it jumps and closes", guides
  // spec §4.3), and the heading lands below the sticky bar.
  test('phone: a jump closes the contents and the heading lands below the bar', async ({ page }) => {
    await page.setViewportSize(PHONE);
    await page.goto(PAGE);
    await barBooted(page);
    const bar = page.getByTestId('guide-bar');
    const toggle = bar.locator('[data-control-bar-toggle]');
    await toggle.click();
    await page.getByTestId('guide-toc-phone-rotation-and-priority').click();
    await expect(page).toHaveURL(/#rotation-and-priority$/);
    await expect(toggle).toHaveAttribute('aria-expanded', 'false');
    // Focus follows the jump, not the row the closed bar now hides.
    await expect(page.locator('#rotation-and-priority')).toBeFocused();
    const barBox = (await bar.boundingBox())!;
    const headingTop = (await page.locator('#rotation-and-priority').boundingBox())!.y;
    expect(headingTop).toBeGreaterThanOrEqual(barBox.y + barBox.height);
  });
});
