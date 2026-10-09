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
// A finished (51-point) build is the band a player reaches at level 60: the live figure
// waits for one (live-gate.ts), so the numeric-figure tests open each spec on it.
const FINISHED_BAND = 60;

const CASES = [
  { classSlug: 'priest', spec: 'priest-holy', label: 'HPS', band: /^\d+\.\d HPS$/ },
  { classSlug: 'warrior', spec: 'warrior-protection', label: 'Tank score', band: /effective health/i },
  { classSlug: 'mage', spec: 'mage-fire', label: 'DPS', band: /^\d+\.\d DPS$/ },
] as const;

for (const { classSlug, spec, label, band } of CASES) {
  test(`${spec} reads ${label} in the score strip and the band card`, async ({ page }) => {
    const { talents } = bisBand(spec, 'alliance', LEVEL_BAND, 'bare') as { talents?: string };
    expect(talents).toBeTruthy();
    await page.goto(`/planner?class=${classSlug}&spec=${spec}&talents=${talents}`);

    const facts = page.getByTestId('planner-facts');
    await expect(facts).toContainText(label);
    if (label !== 'DPS') await expect(facts).not.toContainText('DPS');

    await expect(page.getByTestId('band-compare-set-dps')).toHaveText(band, { timeout: 10_000 });
  });
}

// The live strip headlines the role's own figure, computed by the sim/score code the nightly
// ranker uses (here the fake engine stands in for the wasm: it only has to hand back the
// role's block, and the strip has to print that and not the damage number). A phone is asked
// before it simulates (live-gate.ts), so only the desktop project waits for the figure.
const LIVE_CASES = [
  { classSlug: 'priest', spec: 'priest-holy', figure: /^\d{1,3}\.\d$/ },
  { classSlug: 'warrior', spec: 'warrior-protection', figure: /^\d{1,3}(,\d{3})*$/ },
] as const;

for (const { classSlug, spec, figure } of LIVE_CASES) {
  test(`${spec} shows a numeric live figure with its error`, async ({ page, isMobile }) => {
    test.skip(isMobile, 'a phone is asked before it simulates');
    const { talents } = bisBand(spec, 'alliance', FINISHED_BAND, 'bare') as { talents?: string };
    expect(talents).toBeTruthy();
    await page.goto(`/planner?class=${classSlug}&spec=${spec}&talents=${talents}`);

    await expect(page.getByTestId('planner-dps')).toHaveText(figure, { timeout: 15_000 });
    await expect(page.getByTestId('planner-dps-error')).toHaveText(/^± \d/);
    await expect(page.getByTestId('planner-dps-error')).not.toHaveText(/no live/i);
  });
}
