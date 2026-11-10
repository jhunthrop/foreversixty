// web/tests/e2e/talent-tree-captures.spec.ts
// The talent-tree lane's review captures (design/reviews/2026-11-10-talent-tree): every page that
// draws a talent tree at 1440, 390 and 2000, and its interactions (hover, keyboard focus, a tap on
// a touch screen; add and remove in the planner). Not a gate: it runs only with TREE_CAPTURE_DIR
// set, e.g. `TREE_CAPTURE_DIR=../design/reviews/2026-11-10-talent-tree/after npx playwright test
// talent-tree-captures --project=desktop --workers=1`.
import path from 'node:path';
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { ACTIVE_BUILD } from './support/active-build';
import { openSimState, openToolState, serveMockData, TOOL_CHARACTER } from './support/sim-parity';

const OUT = process.env.TREE_CAPTURE_DIR;
const WIDTHS = [1440, 390, 2000] as const;
const HEIGHT = 1000;

interface TreePage {
  name: string;
  open: (page: Page) => Promise<void>;
  /** The region that holds the trees. */
  region: string;
  editable?: boolean;
}

const PAGES: readonly TreePage[] = [
  {
    name: 'planner',
    open: async (page) => {
      await serveMockData(page);
      await page.goto(`/planner?code=${encodeURIComponent(TOOL_CHARACTER(ACTIVE_BUILD))}`);
      await expect(page.getByTestId('planner-trees')).toBeVisible({ timeout: 20_000 });
    },
    region: 'planner-trees',
    editable: true,
  },
  {
    name: 'guide-glance',
    open: async (page) => {
      await serveMockData(page);
      await page.goto('/guides/warrior/fury');
      await expect(page.getByTestId('guide-glance-trees')).toBeVisible();
    },
    region: 'guide-glance-trees',
  },
  {
    name: 'guide-leveling',
    open: async (page) => {
      await serveMockData(page);
      await page.goto('/guides/warrior/fury');
      await page.getByTestId('guide-leveling-tree').scrollIntoViewIfNeeded();
      await expect(page.getByTestId('guide-leveling-tree')).toBeVisible();
    },
    region: 'guide-leveling-tree',
  },
  {
    name: 'sim-landing',
    open: (page) => openSimState(page, 'landing-in'),
    region: 'sim-talents-card',
  },
  {
    name: 'sim-result',
    open: (page) => openSimState(page, 'result'),
    region: 'sim-talents-card',
  },
  {
    name: 'sim-talents',
    open: (page) => openToolState(page, 'talents'),
    region: 'sim-talent-candidates',
  },
];

/** A learned talent in the region, whichever tree implementation drew it. */
function learnedCell(region: Locator): Locator {
  return region
    .locator(
      '[data-rank]:not([data-rank="0"]), .ui-tree-cell[data-state="full"], .ui-tree-cell[data-state="part"]',
    )
    .first();
}

function file(name: string): string {
  return path.join(OUT ?? '', `${name}.png`);
}

test.describe('talent tree captures', () => {
  test.skip(!OUT, 'review captures only: set TREE_CAPTURE_DIR');
  test.skip(({ isMobile }) => isMobile, 'widths are set per capture');

  for (const tree of PAGES) {
    for (const width of WIDTHS) {
      test(`${tree.name} at ${width}`, async ({ page }) => {
        test.slow();
        await page.setViewportSize({ width, height: HEIGHT });
        await tree.open(page);
        const region = page.getByTestId(tree.region);
        await region.scrollIntoViewIfNeeded();
        // The page's islands (the tooltip host among them) hydrate when idle.
        await page.waitForTimeout(2000);
        await region.screenshot({ path: file(`${tree.name}-${width}`) });
        if (width !== 1440) return;
        const cell = learnedCell(region);
        if ((await cell.count()) === 0) return;
        await page.mouse.move(4, 4);
        await cell.hover();
        await page.waitForTimeout(1200);
        await page.screenshot({ path: file(`${tree.name}-hover`) });
        await page.mouse.move(0, 0);
        await cell.focus();
        await page.keyboard.press('ArrowRight');
        await page.waitForTimeout(500);
        await page.screenshot({ path: file(`${tree.name}-focus`) });
        if (tree.editable) {
          await page.keyboard.press('Escape');
          await cell.click();
          await page.waitForTimeout(300);
          await region.screenshot({ path: file(`${tree.name}-add`) });
          await cell.click({ button: 'right' });
          await page.waitForTimeout(300);
          await region.screenshot({ path: file(`${tree.name}-remove`) });
        }
      });
    }

    test(`${tree.name} tap on a touch screen`, async ({ browser }) => {
      test.slow();
      const context = await browser.newContext({
        viewport: { width: 390, height: 844 },
        hasTouch: true,
        isMobile: true,
      });
      const page = await context.newPage();
      await tree.open(page);
      const region = page.getByTestId(tree.region);
      await region.scrollIntoViewIfNeeded();
      await page.waitForTimeout(2000);
      const cell = learnedCell(region);
      if ((await cell.count()) > 0 && (await cell.isVisible())) {
        await cell.tap();
        await page.waitForTimeout(600);
      }
      await page.screenshot({ path: file(`${tree.name}-tap`) });
      await context.close();
    });
  }
});
