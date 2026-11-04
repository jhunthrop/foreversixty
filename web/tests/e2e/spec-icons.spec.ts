// web/tests/e2e/spec-icons.spec.ts
// Every page that draws a spec icon draws the spec-icon map's (`src/lib/sim/specs.ts`, generated
// from the active client's TalentTab and form-spell icons): the /guides and /bis class cards, a
// class landing's spec cards, and a guide's and a BiS page's header tabs and meta line. The data
// builders behind the apps' files are held by src/lib/bis/spec-icon.test.ts.
import type { Page } from '@playwright/test';
import { SPEC_BY_KEY, SPECS } from '../../src/lib/sim/specs';
import { expect, test } from './support/test';

function iconOfSrc(src: string | null): string | undefined {
  return src?.match(/\/icons\/([a-z0-9_]+)\.webp$/)?.[1];
}

async function expectTabs(page: Page, classSlug: string): Promise<void> {
  const specs = SPECS.filter((s) => s.class_slug === classSlug);
  // A BiS page renders one header per faction view and hides the other: read the shown one.
  const tabs = page.locator('.class-header-spec-tabs:visible').first();
  await expect(tabs.locator('.class-header-spec-tab')).toHaveCount(specs.length);
  for (const spec of specs) {
    const icon = tabs.getByTestId(`bis-spec-tab-${spec.spec_slug}`).locator('[data-icon]');
    await expect(icon, spec.spec).toHaveAttribute('data-icon', spec.icon);
  }
}

test.describe('spec icons come from the spec-icon map', () => {
  test('the /guides class cards', async ({ page }) => {
    await page.goto('/guides');
    const rows = page.locator('[data-testid^="guides-class-card-"] a.spr');
    await expect(rows).toHaveCount(SPECS.length);
    for (const row of await rows.all()) {
      const [, , classSlug, specSlug] = (await row.getAttribute('href'))?.split('/') ?? [];
      const spec = SPEC_BY_KEY[`${classSlug}-${specSlug}`];
      expect(iconOfSrc(await row.locator('img').getAttribute('src')), `${classSlug}-${specSlug}`).toBe(
        spec.icon,
      );
    }
  });

  test('the /bis class cards', async ({ page }) => {
    await page.goto('/bis');
    const rows = page.locator('a.spr[data-spec-key]');
    await expect(rows).toHaveCount(SPECS.length);
    for (const row of await rows.all()) {
      const key = (await row.getAttribute('data-spec-key')) ?? '';
      expect(iconOfSrc(await row.locator('img').getAttribute('src')), key).toBe(SPEC_BY_KEY[key].icon);
    }
  });

  test('a class landing gives each spec card its map icon', async ({ page }) => {
    await page.goto('/guides/druid');
    for (const spec of SPECS.filter((s) => s.class_slug === 'druid')) {
      const card = page.getByTestId(`guide-spec-card-${spec.spec_slug}`).locator('.spec-card-icon');
      await expect(card, spec.spec).toHaveAttribute('data-icon', spec.icon);
    }
  });

  test('a guide header: its spec tabs and the meta line', async ({ page }) => {
    await page.goto('/guides/hunter/marksmanship');
    await expectTabs(page, 'hunter');
    const meta = page.getByTestId('guide-header-meta').locator('img');
    if ((await meta.count()) > 0) {
      await expect(meta).toHaveAttribute('data-icon', SPEC_BY_KEY['hunter-marksmanship'].icon);
    }
  });

  test('a BiS header: its spec tabs', async ({ page }) => {
    await page.goto('/bis/warrior/fury');
    await expectTabs(page, 'warrior');
  });
});
