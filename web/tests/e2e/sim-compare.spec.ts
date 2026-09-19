// web/tests/e2e/sim-compare.spec.ts
// Compare mode: the sim beside the fight it was built from. No second parse and no new
// API -- the report fixture's fight files are already served, statically, at
// /logs-data/reports/fixture2abcd/ (scripts/sync-report-fixture.mjs), the same files
// tests/e2e/report-*.spec.ts read. The one thing compare mode asks for that a report page
// does not is GET /v1/reports/{id} itself (report pages get it inlined at build time,
// src/pages/reports/[id].astro's own `data-report`), so that is the one endpoint this
// suite stubs.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';

const meta = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'fixtures', 'report', 'meta.json'), 'utf8'),
) as Record<string, unknown>;

// Fight 3 (Warden Kelthas) is the report's own default fight (the first encounter) and
// the only one in the fixture whose COMBATANT_INFO row carries gear and talents -- fights
// 1, 2 and 4 record no `combatants` at all, and `fromLoggedFight` refuses a fight without
// one (simCopy.fightNoCombatant). It is therefore the only fight this fixture can open
// compare mode on.
const REF = 'fixture2abcd:3';

async function stubReportMeta(page: Page): Promise<void> {
  await page.route('**/v1/reports/fixture2abcd', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: meta, error: null, request_id: 'r' }),
    }),
  );
}

async function openCompare(page: Page): Promise<void> {
  await stubReportMeta(page);
  await page.goto(`/sim?source=fight&ref=${encodeURIComponent(REF)}&mode=compare`);
  // The fake engine (engine-fake.ts) ticks a few times before it resolves, so the headline
  // lands a moment after navigation rather than on first paint.
  await expect(page.getByTestId('compare-headline')).toBeVisible({ timeout: 10_000 });
}

test('the headline states both DPS figures and the fraction between them', async ({ page }) => {
  await openCompare(page);
  await expect(page.getByTestId('compare-headline')).toHaveText(
    /This fight did [\d,]+ DPS; the sim expects [\d,]+\. That is \d+% of what this gear can do\./,
  );
});

test('the lines explain at least one cast gap in words', async ({ page }) => {
  await openCompare(page);
  const lines = page.getByTestId('compare-lines').locator('li');
  await expect(lines.first()).toBeVisible();
  const texts = await lines.allTextContents();
  expect(texts.some((line) => /cast \d+ times, the sim expects \d+\./.test(line))).toBe(true);
});

test('both tables and the footnote render', async ({ page }) => {
  await openCompare(page);
  await expect(page.getByTestId('compare-abilities')).toBeVisible();
  await expect(page.getByTestId('compare-auras')).toBeVisible();
  await expect(page.getByTestId('compare-footnote')).toHaveText(simCopy.compareFootnote);
});

test('the sentence and the ordinary results are gone in compare mode', async ({ page }) => {
  await openCompare(page);
  await expect(page.getByTestId('sim-sentence')).toHaveCount(0);
  await expect(page.getByTestId('sim-results')).toHaveCount(0);
});

test('the report links to compare mode on its own current fight', async ({ page }) => {
  await stubReportMeta(page);
  await page.goto('/reports/fixture2abcd');
  const link = page.getByTestId('report-sim-fight');
  await expect(link).toBeVisible();
  await link.click();
  await expect(page).toHaveURL(/\/sim\?source=fight&ref=fixture2abcd(?:%3A|:)3&mode=compare/);
});

test.describe('phone', () => {
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('the page does not scroll sideways and every ability row clears 44px', async ({ page }) => {
    await openCompare(page);

    const { scrollWidth, clientWidth } = await page.evaluate(() => ({
      scrollWidth: document.documentElement.scrollWidth,
      clientWidth: document.documentElement.clientWidth,
    }));
    expect(scrollWidth).toBeLessThanOrEqual(clientWidth);

    const rows = await page.locator('[data-testid^="compare-ability-"]').all();
    expect(rows.length).toBeGreaterThan(0);
    for (const row of rows) {
      const box = await row.boundingBox();
      expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
    }
  });
});
