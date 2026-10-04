// web/tests/e2e/guides-redesign.spec.ts
// Acceptance sweep for the guides rebuild (design/specs/2026-10-04-guides.md, §11): the
// nine-crest index, the class landing's spec cards (DPS only for a ranked spec), the spec
// guide's 3-up action rail, and the Leveling band strip's own planner links -- at desktop
// and phone, with the site's own "nothing scrolls sideways" sweep every page-level spec
// runs (layout.spec.ts's own convention).
import { expect, test } from '@playwright/test';

const DESKTOP = { width: 1440, height: 900 };
const PHONE = { width: 390, height: 844 };

test.describe('desktop', () => {
  test.use({ viewport: DESKTOP });

  test('/guides renders nine class crests and 27 spec links', async ({ page }) => {
    await page.goto('/guides');
    const crests = page.locator('[data-testid^="class-crest-"]');
    await expect(crests).toHaveCount(9);
    const specLinks = page.locator('[data-testid="guides-class-picker"] a[href^="/guides/"]');
    await expect(specLinks).toHaveCount(27);
  });

  test('/guides/warrior shows three spec cards, DPS only on the ranked specs', async ({ page }) => {
    await page.goto('/guides/warrior');
    const cards = page.locator('[data-testid="guide-class-spec-cards"] > div');
    await expect(cards).toHaveCount(3);
    const dpsLines = page.locator('[data-testid="guide-spec-card-dps"]');
    // Arms and Fury both carry a ranked band-60 BiS file; Protection does not.
    await expect(dpsLines).toHaveCount(2);
  });

  test('/guides/warrior/fury rail shows four rotation lines with icons, Load/Sim buttons and stat priority lines', async ({
    page,
  }) => {
    await page.goto('/guides/warrior/fury');
    const rotationLines = page.locator('[data-testid^="guide-rail-rotation-line-"]');
    await expect(rotationLines).toHaveCount(4);
    for (const line of await rotationLines.all()) {
      await expect(line.locator('img')).toHaveCount(1);
    }
    await expect(page.getByTestId('guide-rail-load-build')).toBeVisible();
    await expect(page.getByTestId('guide-rail-sim-build')).toBeVisible();
    const statRows = page.locator(
      '[data-testid="guide-rail-stat-priority"] >> text=/Attack power|Strength|Agility|Critical strike|Hit|Melee haste/',
    );
    expect(await statRows.count()).toBeGreaterThanOrEqual(6);
  });

  test('the rail and the full-tree build buttons carry distinct accessible names', async ({ page }) => {
    await page.goto('/guides/warrior/fury');
    await expect(page.getByTestId('guide-rail-load-build')).toHaveAccessibleName(
      'Load this build in the planner (summary)',
    );
    await expect(page.getByTestId('guide-load-build')).toHaveAccessibleName('Load this build');
  });

  test('the Leveling band strip has five rows, each Load in planner link carrying that band’s talent string', async ({
    page,
  }) => {
    await page.goto('/guides/warrior/fury');
    const strip = page.getByTestId('guide-leveling-strip');
    await expect(strip).toBeVisible();
    const rows = strip.locator('[data-testid^="guide-leveling-row-"]:not([data-testid$="-thumbnail"])');
    await expect(rows).toHaveCount(5);
    for (const row of await rows.all()) {
      const link = row.getByText('Load in planner');
      const href = await link.getAttribute('href');
      expect(href).toMatch(/^\/planner\?spec=warrior-fury&talents=/);
    }
  });
});

test.describe('phone', () => {
  test.use({ viewport: PHONE });

  test('/guides/warrior/fury rail stacks to one column with every line intact', async ({ page }) => {
    await page.goto('/guides/warrior/fury');
    const rotationLines = page.locator('[data-testid^="guide-rail-rotation-line-"]');
    await expect(rotationLines).toHaveCount(4);
    const scrollWidth = await page.evaluate(() => document.documentElement.scrollWidth);
    const clientWidth = await page.evaluate(() => document.documentElement.clientWidth);
    expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
  });

  test('the Leveling band strip keeps all five rows on phone (thumbnails dropped)', async ({ page }) => {
    await page.goto('/guides/warrior/fury');
    const rows = page
      .getByTestId('guide-leveling-strip')
      .locator('[data-testid^="guide-leveling-row-"]:not([data-testid$="-thumbnail"])');
    await expect(rows).toHaveCount(5);
    await expect(page.getByTestId('guide-leveling-row-thumbnail').first()).toBeHidden();
  });
});

test.describe('no horizontal overflow', () => {
  for (const [label, viewport] of [
    ['390', PHONE],
    ['1440', DESKTOP],
  ] as const) {
    test.describe(label, () => {
      test.use({ viewport });

      for (const path of ['/guides', '/guides/warrior', '/guides/warrior/fury']) {
        test(`${path} fits the viewport without scrolling sideways`, async ({ page }) => {
          await page.goto(path);
          const { scrollWidth, clientWidth } = await page.evaluate(() => ({
            scrollWidth: document.documentElement.scrollWidth,
            clientWidth: document.documentElement.clientWidth,
          }));
          expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
        });
      }
    });
  }
});
