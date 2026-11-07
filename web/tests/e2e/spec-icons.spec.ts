// web/tests/e2e/spec-icons.spec.ts
// Every page that draws a spec icon draws the spec-icon map's (`src/lib/sim/specs.ts`, generated
// from the active client's TalentTab and form-spell icons): the /guides spec rows and the /bis
// class cards, a class landing's spec rows, and a guide's and a BiS page's ControlBar spec options. The data
// builders behind the apps' files are held by src/lib/bis/spec-icon.test.ts.
import type { Page } from '@playwright/test';
import { SPEC_BY_KEY, SPECS } from '../../src/lib/sim/specs';
import { expect, test } from './support/test';

function iconOfSrc(src: string | null): string | undefined {
  return src?.match(/\/icons\/([a-z0-9_]+)\.webp$/)?.[1];
}

/** A ControlBar's spec options each draw their map icon (`control-spec-<slug>`). */
async function expectBarSpecs(
  page: Page,
  classSlug: string,
  slugOf: (spec: (typeof SPECS)[number]) => string,
) {
  const specs = SPECS.filter((s) => s.class_slug === classSlug);
  await expect(page.locator('[data-testid^="control-spec-"]')).toHaveCount(specs.length);
  for (const spec of specs) {
    const icon = page.getByTestId(`control-spec-${slugOf(spec)}`).locator('img');
    expect(iconOfSrc(await icon.getAttribute('src')), spec.spec).toBe(spec.icon);
  }
}

test.describe('spec icons come from the spec-icon map', () => {
  test('the /guides spec rows', async ({ page }) => {
    await page.goto('/guides');
    const rows = page.locator('a[data-testid^="guide-spec-row-"]');
    await expect(rows).toHaveCount(SPECS.length);
    for (const row of await rows.all()) {
      const [, , classSlug, specSlug] = (await row.getAttribute('href'))?.split('/') ?? [];
      const spec = SPEC_BY_KEY[`${classSlug}-${specSlug}`];
      expect(iconOfSrc(await row.locator('img').getAttribute('src')), `${classSlug}-${specSlug}`).toBe(
        spec.icon,
      );
    }
  });

  test('the /bis class and spec plate', async ({ page }) => {
    await page.goto('/bis');
    const rows = page.locator('a[data-testid^="bis-spec-row-"]');
    await expect(rows).toHaveCount(SPECS.length);
    for (const row of await rows.all()) {
      const key = ((await row.getAttribute('data-testid')) ?? '')
        .slice('bis-spec-row-'.length)
        .replace('/', '-');
      expect(iconOfSrc(await row.locator('img').getAttribute('src')), key).toBe(SPEC_BY_KEY[key].icon);
    }
  });

  test('a class landing gives each spec row its map icon', async ({ page }) => {
    await page.goto('/guides/druid');
    for (const spec of SPECS.filter((s) => s.class_slug === 'druid')) {
      const icon = page.getByTestId(`guide-spec-row-druid/${spec.spec_slug}`).locator('img');
      expect(iconOfSrc(await icon.getAttribute('src')), spec.spec).toBe(spec.icon);
    }
  });

  // The guide's header spec tabs and its meta-line icon retire (spec §4.3): the specs are the
  // ControlBar's segmented options, each with its icon, as on a BiS page.
  test('a guide control bar: its spec options', async ({ page }) => {
    await page.goto('/guides/hunter/marksmanship');
    await expectBarSpecs(page, 'hunter', (spec) => spec.spec_slug);
  });

  test('a BiS control bar: its spec options', async ({ page }) => {
    await page.goto('/bis/warrior/fury');
    await expectBarSpecs(page, 'warrior', (spec) => spec.spec);
  });
});
