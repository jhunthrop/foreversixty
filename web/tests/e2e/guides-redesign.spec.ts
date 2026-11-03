// web/tests/e2e/guides-redesign.spec.ts
// Acceptance sweep for the guides rebuild (design/specs/2026-10-04-guides.md, §11): the
// nine-crest index, the class landing's spec cards (spec 2026-11-02), the spec
// guide's 3-up action rail, and the Leveling band strip's own planner links -- at desktop
// and phone, with the site's own "nothing scrolls sideways" sweep every page-level spec
// runs (layout.spec.ts's own convention).
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { cardLineFor } from '../../src/lib/guides/class-landing';
import { expect, test } from './support/test';

const DESKTOP = { width: 1440, height: 900 };
const PHONE = { width: 390, height: 844 };

/** The card line the class landing shows: the spec guide's own Overview summary, read through
 *  the same helper the page uses, so the test and the page share one source of truth. */
function cardLine(classSlug: string, specSlug: string): string {
  const path = fileURLToPath(
    new URL(`../../src/content/guides/${classSlug}/${specSlug}.md`, import.meta.url),
  );
  const body = readFileSync(path, 'utf8').split(/^---$/m).slice(2).join('---');
  return cardLineFor(body, path);
}

test.describe('desktop', () => {
  test.use({ viewport: DESKTOP });

  test('/guides renders nine class crests and 28 spec links', async ({ page }) => {
    await page.goto('/guides');
    const crests = page.locator('[data-testid^="class-crest-"]');
    await expect(crests).toHaveCount(9);
    const specLinks = page.locator('[data-testid="guides-class-picker"] a[href^="/guides/"]');
    await expect(specLinks).toHaveCount(28);
  });

  test('/guides/warrior leads with three spec cards that are each the link to their guide, with no figure', async ({
    page,
  }) => {
    await page.goto('/guides/warrior');
    await expect(page.locator('h1')).toHaveText('Warrior');
    await expect(page.getByRole('tablist', { name: 'Spec' })).toHaveCount(0);
    const cards = page.locator('[data-testid="guide-class-spec-cards"] > a');
    await expect(cards).toHaveCount(3);
    await expect(cards.nth(1)).toHaveAttribute('href', '/guides/warrior/fury');
    await expect(cards.nth(1)).toHaveAccessibleName('Read the Fury Warrior guide');
    await expect(page.locator('[data-testid="guide-class-spec-cards"]')).not.toContainText(
      /DPS at 60|per second/,
    );
  });

  test('every spec card says how the spec plays, from its guide’s Overview', async ({ page }) => {
    await page.goto('/guides/warrior');
    for (const specSlug of ['arms', 'fury', 'protection']) {
      await expect(
        page.getByTestId(`guide-spec-card-${specSlug}`).getByTestId('guide-spec-card-line'),
      ).toHaveText(cardLine('warrior', specSlug));
    }
  });

  test('the class facts follow the cards under their own headings, with no sourcing caveat', async ({
    page,
  }) => {
    await page.goto('/guides/warrior');
    await expect(page.locator('.class-landing-body h2')).toHaveText(['What changed from 1.12', 'Races']);
    await expect(page.getByTestId('class-landing-what-changed-from-1-12').locator('li')).toHaveCount(8);
    await expect(page.getByTestId('confidence-footer-note')).toHaveCount(0);
  });

  test('druid gives Feral and Feral Bear their own icons', async ({ page }) => {
    await page.goto('/guides/druid');
    // data-icon names the spec's icon whether its art is in this build or not (the fixture
    // carries a subset, so a card can draw the neutral square instead of the image).
    const iconOf = (spec: string) =>
      page.getByTestId(`guide-spec-card-${spec}`).locator('.spec-card-icon').getAttribute('data-icon');
    const cat = await iconOf('feral');
    const bear = await iconOf('feral-bear');
    expect(cat).toBe('ability_druid_catformattack');
    expect(cat).not.toBe(bear);
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

  test('/guides/warrior/protection shows the Build, Rotation and Stat priority rail cards and runs 3-up', async ({
    page,
  }) => {
    // The rail renders a card only when it has real data (a spec with nothing to show has it
    // omitted and the rail narrows to 2-up or 1-up). Protection now publishes a ranked
    // tank BiS file and a rotation, so it carries all three.
    await page.goto('/guides/warrior/protection');
    await expect(page.getByTestId('guide-rail-build')).toBeVisible();
    await expect(page.getByTestId('guide-rail-rotation')).toBeVisible();
    await expect(page.getByTestId('guide-rail-stat-priority')).toBeVisible();
    const rail = page.getByTestId('guide-action-rail');
    await expect(rail).toHaveCSS('grid-template-columns', /^[\d.]+px [\d.]+px [\d.]+px$/);
  });

  test('the Leveling band strip has five rows, each Load in planner link carrying that band’s talent string', async ({
    page,
  }) => {
    await page.goto('/guides/warrior/fury');
    const strip = page.getByTestId('guide-leveling-strip');
    await expect(strip).toBeVisible();
    const rows = strip.locator('[data-testid^="guide-leveling-row-"]:not([data-testid$="-talents"])');
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

  test('the Leveling band strip keeps all five rows on phone, each listing its talents', async ({ page }) => {
    await page.goto('/guides/warrior/fury');
    const rows = page
      .getByTestId('guide-leveling-strip')
      .locator('[data-testid^="guide-leveling-row-"]:not([data-testid$="-talents"])');
    await expect(rows).toHaveCount(5);
    await expect(page.getByTestId('guide-leveling-row-talents').first()).toBeVisible();
    await expect(page.getByTestId('guide-leveling-row-talents').first().locator('li').first()).toBeVisible();
  });
});

const CLASS_SLUGS = ['druid', 'hunter', 'mage', 'paladin', 'priest', 'rogue', 'shaman', 'warlock', 'warrior'];
/** The least room between the spec cards and the class notes below them. */
const MIN_GAP_BELOW_CARDS = 16;

test.describe('class landings', () => {
  for (const [label, viewport] of [
    ['390', PHONE],
    ['1440', DESKTOP],
  ] as const) {
    test.describe(label, () => {
      test.use({ viewport });

      for (const classSlug of CLASS_SLUGS) {
        test(`/guides/${classSlug} keeps every spec tab on screen and the notes clear of the cards`, async ({
          page,
        }) => {
          await page.goto(`/guides/${classSlug}`);
          const clientWidth = await page.evaluate(() => document.documentElement.clientWidth);
          for (const tab of await page.locator('[data-testid^="bis-spec-tab-"]').all()) {
            const box = await tab.boundingBox();
            expect(box).not.toBeNull();
            expect(box!.x).toBeGreaterThanOrEqual(0);
            expect(box!.x + box!.width).toBeLessThanOrEqual(clientWidth);
          }
          const gap = await page.getByTestId('guide-class-spec-cards').evaluate((cards) => {
            const next = cards.nextElementSibling;
            return next === null
              ? null
              : next.getBoundingClientRect().top - cards.getBoundingClientRect().bottom;
          });
          if (gap !== null) expect(gap).toBeGreaterThanOrEqual(MIN_GAP_BELOW_CARDS);
        });
      }
    });
  }

  test.describe('1440 row', () => {
    test.use({ viewport: DESKTOP });

    test('/guides/druid sets its four spec cards in one row', async ({ page }) => {
      await page.goto('/guides/druid');
      const tops = await page
        .locator('[data-testid="guide-class-spec-cards"] > a')
        .evaluateAll((cards) => cards.map((card) => Math.round(card.getBoundingClientRect().top)));
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
