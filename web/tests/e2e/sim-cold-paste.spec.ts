// web/tests/e2e/sim-cold-paste.spec.ts
// Persona review 2026-10-06 (retail-raider), §5/§8 item 1: "Let the sim take my real
// gear/talents in one step... no sign-in detour -- the one surface where Raidbots still
// wins". The engine runs as wasm in this browser (lib/sim/engine.ts, /_sim/<version>/), so
// a visitor signed out should be able to paste an export and get a number with no queue --
// this is the end-to-end proof of the landing's paste Plate (SimLanding.svelte, spec
// 2026-11-06-sim §4.2), through the real /sim UI rather than through cold-paste.test.ts, which
// cannot exercise `$state` or a real run against the fake engine.
import { expect, test } from './support/test';
import { simPageCopy } from '../../src/lib/sim/page-copy';
import { gotoSim, SIM_FURY as FURY, simCharacter } from './support/sim-page';
// Only an export that names its character becomes the current one.
const NAMED_FURY = `${FURY}|who=Simfury:Realm`;

test('a cold paste decodes, runs with no further click, and writes the current-character pointer', async ({
  page,
}) => {
  await gotoSim(page);

  // No session stub anywhere in this test: the preview server behind it has no API, so
  // fetchMeOnce fails and the page is signed out by default -- the same ground truth every
  // other signed-out /sim spec (sim-run.spec.ts, sim-sources.spec.ts) already relies on.
  const run = page.getByTestId('sim-cold-paste-run');
  await expect(run).toBeDisabled();

  await page.getByTestId('sim-cold-paste-input').fill(NAMED_FURY);
  await expect(run).toBeEnabled();
  await expect(page.getByText(simPageCopy.paste.invalid)).toHaveCount(0);

  await run.click();

  // The character's head and the run line both land from the one click -- no separate
  // "Run sim" press afterward (runPastedInput in SimView.svelte).
  await expect(simCharacter(page)).toBeVisible();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 5000 });
  await expect(page.getByTestId('sim-answer')).toHaveAttribute('data-state', 'result');
  await expect(page.getByTestId('sim-dps')).toHaveText(/\d/);

  // The pointer current-character.spec.ts's own restore tests prove a bare /planner or
  // /sim load reads back -- this proves this box writes the same one, not a parallel copy.
  const pointer = await page.evaluate(() => localStorage.getItem('fs.currentCharacter'));
  expect(pointer).not.toBeNull();
  const parsed = JSON.parse(pointer ?? '{}') as { source: string; ref: string };
  expect(parsed.source).toBe('addon');
  expect(parsed.ref).toBe(NAMED_FURY);
});

test('a bad paste shows the validation message and Run stays disabled', async ({ page }) => {
  await gotoSim(page);

  const run = page.getByTestId('sim-cold-paste-run');
  await page.getByTestId('sim-cold-paste-input').fill('not a real export');

  await expect(run).toBeDisabled();
  // The Field's own error line under the box (spec §5, "Paste plate").
  const error = page.locator('#sim-cold-paste-input-error');
  await expect(error).toHaveText('That does not look like an addon export or a build code.');
  await expect(page.getByTestId('sim-cold-paste-input')).toHaveAttribute('aria-invalid', 'true');

  // Clearing it back to empty returns to the silent, not-yet-an-error state (cold-paste.ts).
  await page.getByTestId('sim-cold-paste-input').fill('');
  await expect(error).toHaveCount(0);
  await expect(run).toBeDisabled();
});

// Signed out, the answer's save Button reads "Sign in to save" (spec §4.5, §5.1): saving still
// needs an account, and the Button is the way to one rather than a dead control.
test('Save keeps the sign-in requirement after a cold-paste run', async ({ page }) => {
  await gotoSim(page);

  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 5000 });

  const save = page.getByTestId('sim-save-open');
  await expect(save).toHaveText(simPageCopy.answer.signInToSave);
  await expect(page.getByTestId('sim-save-title')).toHaveCount(0);
});

// Build item 3: phone -- the paste box is the first thing after the title, every control
// is a real 44px hit target, and nothing overflows at 390 (sim.astro's own documented
// measurement width). A literal override rather than the shared mobile project's Pixel 7
// width (412px, phone-scroll.ts's own note): 390 is the narrower, stricter check, and the
// desktop project never runs this block at all.
test.describe('phone, 390', () => {
  test.use({ viewport: { width: 390, height: 844 } });
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('the paste box renders ahead of the Battle.net sign-in, with no horizontal overflow', async ({
    page,
  }) => {
    await gotoSim(page);

    const pasteTop = await page
      .getByTestId('sim-cold-paste')
      .evaluate((el) => el.getBoundingClientRect().top);
    const accountTop = await page.getByTestId('sim-sign-in').evaluate((el) => el.getBoundingClientRect().top);
    expect(pasteTop).toBeLessThan(accountTop);

    const scrollWidth = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(scrollWidth, 'page scrolls horizontally at 390').toBeLessThanOrEqual(390);
  });

  test('the paste field and Run are real 44px hit targets', async ({ page }) => {
    await gotoSim(page);
    const input = await page.getByTestId('sim-cold-paste-input').boundingBox();
    const run = await page.getByTestId('sim-cold-paste-run').boundingBox();
    expect(input?.height ?? 0).toBeGreaterThanOrEqual(44);
    expect(run?.height ?? 0).toBeGreaterThanOrEqual(44);
  });
});
