// web/tests/e2e/sim-cold-paste.spec.ts
// Persona review 2026-10-06 (retail-raider), §5/§8 item 1: "Let the sim take my real
// gear/talents in one step... no sign-in detour -- the one surface where Raidbots still
// wins". The engine runs as wasm in this browser (lib/sim/engine.ts, /_sim/<version>/), so
// a visitor signed out should be able to paste an export and get a number with no queue --
// this is the end-to-end proof of ColdPasteHero.svelte, through the real /sim UI rather
// than through the component's own unit tests (ColdPasteHero.test.ts, cold-paste.test.ts),
// which cannot exercise `$state` or a real run against the fake engine.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';
import { landingCopy } from '../../src/lib/sim/landing-copy';
import { simCopy } from '../../src/lib/sim/copy';

// Same fixture as sim-run.spec.ts, sim-settings.spec.ts and sim-sources.spec.ts: an addon
// export needs no API stub, so this reads the site's own active build id rather than
// hardcoding one that may not match the fixture data this suite runs against.
const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

test('a cold paste decodes, runs with no further click, and writes the current-character pointer', async ({
  page,
}) => {
  await page.goto('/sim');

  // No session stub anywhere in this test: the preview server behind it has no API, so
  // fetchMeOnce fails and the page is signed out by default -- the same ground truth every
  // other signed-out /sim spec (sim-run.spec.ts, sim-sources.spec.ts) already relies on.
  const run = page.getByTestId('sim-cold-paste-run');
  await expect(run).toBeDisabled();

  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await expect(run).toBeEnabled();
  await expect(page.getByTestId('sim-cold-paste-error')).toHaveCount(0);

  await run.click();

  // The character strip and the run control both land from the one click -- no separate
  // "Run sim" press afterward (runPastedInput in SimView.svelte).
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 5000 });
  await expect(page.getByTestId('sim-dps')).not.toHaveText('—');

  // The pointer current-character.spec.ts's own restore tests prove a bare /planner or
  // /sim load reads back -- this proves this box writes the same one, not a parallel copy.
  const pointer = await page.evaluate(() => localStorage.getItem('fs.currentCharacter'));
  expect(pointer).not.toBeNull();
  const parsed = JSON.parse(pointer ?? '{}') as { source: string; ref: string };
  expect(parsed.source).toBe('addon');
  expect(parsed.ref).toBe(FURY);
});

test('a bad paste shows the validation message and Run stays disabled', async ({ page }) => {
  await page.goto('/sim');

  const run = page.getByTestId('sim-cold-paste-run');
  await page.getByTestId('sim-cold-paste-input').fill('not a real export');

  await expect(run).toBeDisabled();
  await expect(page.getByTestId('sim-cold-paste-error')).toHaveText(
    'That does not look like an addon export or a build code.',
  );

  // Clearing it back to empty returns to the silent, not-yet-an-error state (cold-paste.ts).
  await page.getByTestId('sim-cold-paste-input').fill('');
  await expect(page.getByTestId('sim-cold-paste-error')).toHaveCount(0);
  await expect(run).toBeDisabled();
});

test('Save keeps the sign-in requirement after a cold-paste run', async ({ page }) => {
  await page.goto('/sim');

  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 5000 });

  const save = page.getByTestId('sim-save-open');
  await expect(save).toHaveText(simCopy.signInToSave);
  await expect(save).toBeDisabled();
});

test('the reassurance line says the run stays local', async ({ page }) => {
  await page.goto('/sim');
  await expect(page.getByText(landingCopy.pasteHeroRunsLocally)).toBeVisible();
});

// Build item 3: phone -- the paste box is the first thing after the title, every control
// is a real 44px hit target, and nothing overflows at 390 (sim.astro's own documented
// measurement width). A literal override rather than the shared mobile project's Pixel 7
// width (412px, phone-scroll.ts's own note): 390 is the narrower, stricter check, and the
// desktop project never runs this block at all.
test.describe('phone, 390', () => {
  test.use({ viewport: { width: 390, height: 844 } });
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('the paste box renders ahead of the Battle.net hero, with no horizontal overflow', async ({
    page,
  }) => {
    await page.goto('/sim');

    const pasteTop = await page
      .getByTestId('sim-cold-paste')
      .evaluate((el) => el.getBoundingClientRect().top);
    const accountTop = await page
      .getByTestId('sim-account-card')
      .evaluate((el) => el.getBoundingClientRect().top);
    expect(pasteTop).toBeLessThan(accountTop);

    const scrollWidth = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(scrollWidth, 'page scrolls horizontally at 390').toBeLessThanOrEqual(390);
  });

  test('the textarea and Run are real 44px hit targets', async ({ page }) => {
    await page.goto('/sim');
    const input = await page.getByTestId('sim-cold-paste-input').boundingBox();
    const run = await page.getByTestId('sim-cold-paste-run').boundingBox();
    expect(input?.height ?? 0).toBeGreaterThanOrEqual(44);
    expect(run?.height ?? 0).toBeGreaterThanOrEqual(44);
  });
});
