// web/tests/captures/selected/capture.spec.ts
// Review captures for the one selected state (design/reviews/2026-11-09-selected): every page the
// owner named, signed in as a Beast Mastery Hunter where the page marks the viewer's own row, at
// 1440 and 390. Not a test of behaviour. CAPTURE_VARIANT=before writes before/, else after/.
// Run from web/ against a served copy of dist (see playwright.capture.config.ts).
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test, type Page } from '@playwright/test';
import { openBis } from '../../e2e/support/bis-page';
import { openBisLanding } from '../../e2e/support/bis-landing-page';
import { openGuidesState } from '../../e2e/support/guides-page';
import { openPlanner } from '../../e2e/support/planner-page';
import { armoryPointer, signInWith, storePointer, type FixtureCharacter } from '../../e2e/support/selector';
import { openSimState } from '../../e2e/support/sim-parity';

const VARIANT = process.env.CAPTURE_VARIANT === 'before' ? 'before' : 'after';
const OUT = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../../design/reviews/2026-11-09-selected',
  VARIANT,
);
const WIDTHS = [
  { width: 1440, height: 900, scale: 1 },
  { width: 390, height: 844, scale: 2 },
] as const;

const FELSNAP: FixtureCharacter = {
  name: 'Felsnap',
  class: 'Hunter',
  spec: 'Beast Mastery',
  level: 60,
  realm: 'Living Flame',
  faction: 'alliance',
  buildDaysAgo: 0.5,
};

async function signInHunter(page: Page): Promise<void> {
  await signInWith(page, [FELSNAP, { ...FELSNAP, name: 'Quickshot', spec: 'Marksmanship', level: 33 }]);
  await storePointer(page, armoryPointer(FELSNAP));
}

const phone = (page: Page) => (page.viewportSize()?.width ?? 1440) < 768;
const chestSlot = (page: Page) => (phone(page) ? 'paperdoll-list-slot-chest' : 'paperdoll-frame-slot-chest');
const ZOOM_MARGIN_PX = 120;

interface Shot {
  id: string;
  open: (page: Page) => Promise<void>;
  /** Viewport only (an open overlay), not the full page. */
  viewport?: boolean;
  /** Also writes `<id>-zoom-<width>.png`, the area around this test id. */
  zoom?: (page: Page) => string;
}

const SHOTS: readonly Shot[] = [
  { id: 'bis-landing', open: (page) => openBisLanding(page, '/bis', '34') },
  {
    id: 'bis-spec-slot',
    open: async (page) => {
      await signInHunter(page);
      await openBis(page, '/bis/hunter/beast-mastery/30-39');
      await page.getByTestId(chestSlot(page)).first().click();
      await page.mouse.move(0, 0);
    },
    zoom: chestSlot,
  },
  {
    id: 'tiers-own',
    open: async (page) => {
      await signInHunter(page);
      await page.goto('/tiers');
      await expect(page.getByTestId('tier-list').first()).toBeVisible();
      await page.waitForLoadState('networkidle');
    },
  },
  { id: 'sim-buffs', open: (page) => openSimState(page, 'landing-in') },
  { id: 'planner', open: (page) => openPlanner(page, '/planner', 'felsnap') },
  { id: 'guides-spec-picker', open: (page) => openGuidesState(page, '/guides/hunter', 'hunter34') },
  {
    id: 'selector-panel',
    viewport: true,
    open: async (page) => {
      await openBisLanding(page, '/bis', '34');
      await page.getByTestId('character-selector').first().click();
      await expect(page.getByTestId('selector-panel')).toBeVisible();
      await page.waitForTimeout(400);
    },
  },
];

for (const shot of SHOTS) {
  for (const size of WIDTHS) {
    test(`${shot.id} at ${size.width}`, async ({ browser }) => {
      const context = await browser.newContext({
        viewport: { width: size.width, height: size.height },
        deviceScaleFactor: size.scale,
        isMobile: size.width < 768,
        hasTouch: size.width < 768,
        reducedMotion: 'reduce',
      });
      const page = await context.newPage();
      await shot.open(page);
      await page.evaluate(() => document.fonts.ready);
      await page.screenshot({
        path: path.join(OUT, `${shot.id}-${size.width}.png`),
        fullPage: shot.viewport !== true,
      });
      if (shot.zoom !== undefined) {
        const box = await page.getByTestId(shot.zoom(page)).first().boundingBox();
        if (box === null) throw new Error(`${shot.id}: nothing to zoom on`);
        await page.screenshot({
          path: path.join(OUT, `${shot.id}-zoom-${size.width}.png`),
          clip: {
            x: Math.max(0, box.x - ZOOM_MARGIN_PX),
            y: Math.max(0, box.y - ZOOM_MARGIN_PX),
            width: box.width + ZOOM_MARGIN_PX * 2,
            height: box.height + ZOOM_MARGIN_PX * 2,
          },
        });
      }
      await context.close();
    });
  }
}
