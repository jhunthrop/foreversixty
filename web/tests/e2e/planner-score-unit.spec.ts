// web/tests/e2e/planner-score-unit.spec.ts
// The planner names its score by the build's role: HPS for a healer, tank score for a tank,
// DPS for everyone else. The fixture planner data is a warrior only, so this runs against the
// real pipeline output (FOREVER_DATA=real), opening each spec on its own level-20 band build.
import { expect, test } from '@playwright/test';
import { bisBand } from './support/bis-file';

test.skip(
  process.env.FOREVER_DATA !== 'real',
  'needs the real class talent data; run it with FOREVER_DATA=real',
);

const LEVEL_BAND = 20;

const CASES = [
  { classSlug: 'priest', spec: 'priest-holy', label: 'HPS', band: /^\d+\.\d HPS$/ },
  { classSlug: 'warrior', spec: 'warrior-protection', label: 'Tank score', band: /effective health/i },
  { classSlug: 'mage', spec: 'mage-fire', label: 'DPS', band: /^\d+\.\d DPS$/ },
] as const;

for (const { classSlug, spec, label, band } of CASES) {
  test(`${spec} reads ${label} in the score strip and the band card`, async ({ page }) => {
    const { talents } = bisBand(spec, 'alliance', LEVEL_BAND) as { talents?: string };
    expect(talents).toBeTruthy();
    await page.goto(`/planner?class=${classSlug}&spec=${spec}&talents=${talents}`);

    const facts = page.getByTestId('planner-facts');
    await expect(facts).toContainText(label);
    if (label !== 'DPS') await expect(facts).not.toContainText('DPS');
    if (label !== 'DPS') await expect(page.getByTestId('planner-dps-error')).not.toHaveText(/reaches 51/);

    await expect(page.getByTestId('band-compare-set-dps')).toHaveText(band, { timeout: 10_000 });
  });
}
