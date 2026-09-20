// web/tests/e2e/sim-weights.spec.ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

async function loadWeights(page: Page): Promise<void> {
  await page.goto('/sim/weights');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

test('the page opens with the warning, above everything else', async ({ page }) => {
  await page.goto('/sim/weights');
  const warning = page.getByTestId('sim-weights-warning');
  await expect(warning).toBeVisible();
  await expect(warning).toContainText('not a straight line');
  await expect(warning.getByRole('link', { name: /Top Gear/i })).toHaveAttribute('href', '/sim/gear');
});

test('the stat picker defaults to the spec’s reference stat, first and ticked', async ({ page }) => {
  await page.route('**/v1/specs', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        request_id: 'r',
        error: null,
        data: {
          specs: [
            {
              spec: 'warrior-fury',
              state: 'validated',
              median_gap: 0.03,
              parses: 50,
              worst_actions: [],
              engine_version: 'x',
              updated_at: null,
              // Contract 10.1 A7: proto.Stat enum names in snake case.
              reference_stat: 'attack_power',
            },
          ],
        },
      }),
    }),
  );
  await loadWeights(page);
  await expect(page.getByTestId('sim-weight-pick-attack_power')).toBeChecked();
  await expect(page.getByTestId('sim-weights-reference')).toHaveText(/Attack power/);
});

test('a run renders a weight per stat with an error bar and a Pawn string', async ({ page, context }) => {
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  await loadWeights(page);
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-weights')).toBeVisible({ timeout: 25_000 });
  await expect(page.getByTestId('sim-weight-attack_power')).toContainText('1.00');
  // Contract 10.8: haste is split, hit and crit are not.
  await expect(page.getByTestId('sim-weight-melee_haste')).toBeVisible();
  await expect(page.getByTestId('sim-weight-crit')).toBeVisible();
  await expect(page.getByTestId('sim-weight-hit')).toBeVisible();
  for (const wrong of ['haste', 'melee_crit', 'melee_hit']) {
    await expect(page.getByTestId(`sim-weight-${wrong}`)).toHaveCount(0);
  }
  await expect(page.getByTestId('sim-pawn')).toContainText('( Pawn: v1:');
  await page.getByTestId('sim-pawn-copy').click();
  await expect(page.getByTestId('sim-pawn-copy')).toHaveText('Copied');
});

test('a weights run never shows a DPS figure on the progress line', async ({ page }) => {
  await loadWeights(page);
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-weights')).toBeVisible({ timeout: 25_000 });
  // Engine-lane rule 3: a weights run's progress ticks carry dps: 0 and must never render as
  // a DPS figure. This page has no bulk stage count to show for one wasm call either, so the
  // progress line should simply never have appeared during the run.
  await expect(page.getByTestId('sim-stage-progress')).toHaveCount(0);
});

test('there is no slot grid, no source picker and no named sets', async ({ page }) => {
  await loadWeights(page);
  await expect(page.getByTestId('sim-slot-grid')).toHaveCount(0);
  await expect(page.getByTestId('sim-source-picker')).toHaveCount(0);
  await expect(page.getByTestId('sim-named-sets')).toHaveCount(0);
});
