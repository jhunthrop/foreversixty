import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';

// Same fixture as sim-run.spec.ts, sim-settings.spec.ts and sim-sources.spec.ts: an addon
// export needs no API stub, so this reads the site's own active build id rather than
// hardcoding one.
const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

async function loadFuryAndRun(page: Page): Promise<void> {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await page.getByTestId('sim-run-button').click();
  // The fake engine (engine-fake.ts) always resolves with the fixture's own summary, so
  // this is the moment `store.result` lands and SimResults' own lazy chunk starts loading.
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
  await expect(page.getByTestId('sim-results')).toBeVisible();
}

// Every figure asserted below comes from src/fixtures/sim/result.json's own summary --
// engine-fake.ts's simRun and simCombine both answer with `summary: fixture.summary`
// verbatim, whatever DPS numbers the fake's seeded normal distribution invents for `dps`.
// public/data/<build>/simnames/warrior.json does not exist under FOREVER_DATA=fixture (the
// fixture planner tree has no spellconst directory, so scripts/sync-data.mjs's own
// writeSimNames writes nothing), so `store.actionNames` stays null. D48 (dps-minmaxer
// review round 2): that used to mean every ability, aura and cast name on this page was
// the engine's own raw action key -- "spell:20662" on screen. resolveActionName's own
// humanised fallback (action-names.ts) means a null table now reads "Spell 20662" instead
// -- English, never a raw key, even before the build's own table has loaded.

test('the sentence names the two biggest damage sources, their share and the biggest buff', async ({
  page,
}) => {
  await loadFuryAndRun(page);

  // spell:25286 (64,767) and spell:23894 (33,210) of the actor's 186,849 total is 52%;
  // spell:9910 is the first 100%-uptime buff in the fixture's own array order (the tiebreak
  // summarySentence uses -- see sentence.ts -- keeps a stable sort's original order, unlike
  // AuraTable's own name-tiebroken sort below).
  const sentence = page.getByTestId('sim-sentence');
  await expect(sentence).toHaveText(
    'Spell 25286 and Spell 23894 are 52% of your damage; Spell 9910 is up 100% of the fight.',
  );
  await expect(sentence).not.toContainText(/\bspell:/);
});

test('the Damage tab is selected first and the actor table reads a humanised label, never the raw key', async ({
  page,
}) => {
  await loadFuryAndRun(page);

  const damageTab = page.getByTestId('sim-tab-damage');
  await expect(damageTab).toHaveAttribute('aria-selected', 'true');

  await page.getByTestId('actor-sim-player').click();
  const row = page.getByTestId('row-abilities');
  await expect(row).toContainText('Spell 25286');
  await expect(row).not.toContainText(/\bspell:/);
});

test('the Buffs tab shows the first aura row at 100% uptime', async ({ page }) => {
  await loadFuryAndRun(page);

  await page.getByTestId('sim-tab-buffs').click();
  const first = page.getByTestId('aura-table').locator('li').first();
  await expect(first).toContainText('Spell 15366');
  await expect(first).not.toContainText(/\bspell:/);
  await expect(page.getByTestId('aura-uptime').first()).toHaveText('100.0%');
});

test('the Debuffs tab is empty: the fixture summary has no debuff-type aura', async ({ page }) => {
  await loadFuryAndRun(page);

  await page.getByTestId('sim-tab-debuffs').click();
  await expect(page.getByTestId('table-empty')).toHaveText('No debuffs in this window.');
});

test('the Casts tab reads a cast row by its humanised label and count, never the raw key', async ({
  page,
}) => {
  await loadFuryAndRun(page);

  await page.getByTestId('sim-tab-casts').click();
  const row = page.getByTestId('cast-sim-player-23894');
  await expect(row).toContainText('Spell 23894');
  await expect(row).not.toContainText(/\bspell:/);
  await expect(row).toContainText('27');
});

test('the Resources tab is empty: the fixture summary carries no non-zero resource series', async ({
  page,
}) => {
  await loadFuryAndRun(page);

  await page.getByTestId('sim-tab-resources').click();
  await expect(page.getByTestId('table-empty')).toHaveText('No resource changes in this window.');
});

test('the Timeline tab draws one lane, the player rostered from the sample iteration', async ({ page }) => {
  await loadFuryAndRun(page);

  await page.getByTestId('sim-tab-timeline').click();
  // The design's intent is a sample cast timeline (TimelinesView draws CastRow.sequence
  // as ticks, exactly as it draws a logged fight's), but src/fixtures/sim/result.json's
  // own cast rows and aura tracks both carry empty `sequence`/`segments` arrays -- only
  // their aggregate counts (`succeeded`, `uptime_ms`) are populated. So this fixture's own
  // lane draws no ticks or bars; what it does prove is the one-lane-per-roster-row wiring,
  // which is what this asserts.
  await expect(page.getByTestId('lane-sim-player')).toBeVisible();
  await expect(page.getByTestId('lane-sim-player')).toContainText('Sim');
});

test('the Distribution tab shows the mean and five rows, and stands in for the missing chart', async ({
  page,
}) => {
  await loadFuryAndRun(page);

  await page.getByTestId('sim-tab-distribution').click();
  const distribution = page.getByTestId('sim-distribution');
  await expect(distribution).toBeVisible();
  await expect(distribution).toContainText('Mean DPS');
  await expect(distribution).toContainText('Standard deviation');
  await expect(distribution).toContainText('Lowest iteration');
  await expect(distribution).toContainText('Highest iteration');
  await expect(distribution).toContainText('Iterations');
  await expect(distribution.locator('dl > div')).toHaveCount(5);
});

test('no brushable time chart exists anywhere on the page', async ({ page }) => {
  await loadFuryAndRun(page);

  for (const tab of [
    'damage',
    'buffs',
    'debuffs',
    'casts',
    'resources',
    'timeline',
    'sample',
    'distribution',
  ]) {
    await page.getByTestId(`sim-tab-${tab}`).click();
    await expect(page.getByTestId('time-chart')).toHaveCount(0);
  }
});

test('the buffs tab counts applications and the sample tab shows one iteration', async ({ page }) => {
  await loadFuryAndRun(page);

  await page.getByTestId('sim-tab-buffs').click();
  await expect(page.getByTestId('aura-applied').first()).toHaveText(/^\d+$/);

  await page.getByTestId('sim-tab-sample').click();
  const log = page.getByTestId('sim-sample-log');
  await expect(log).toBeVisible();
  // The fixture's first two casts are before the pull and are labelled as such.
  await expect(log.getByText(simCopy.samplePrePull).first()).toBeVisible();
  // public/data/<build>/simnames/warrior.json does not exist under FOREVER_DATA=fixture
  // (see the note above loadFuryAndRun), so the sample log resolves nothing and renders
  // resolveActionName's humanised fallback -- the same thing the cast table renders in
  // this environment. The resolved-name path is covered instead, with names injected, by
  // src/lib/sim/sample-log.test.ts.
  await expect(log).toContainText('Spell 25286');
  await expect(log).not.toContainText(/\bspell:/);
  // Resources after each cast: the fixture's only pool is rage.
  await expect(log.locator('th', { hasText: simCopy.resourceLabel.rage })).toHaveCount(1);
});

test.describe('phone', () => {
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('every results tab clears 44px', async ({ page }) => {
    await loadFuryAndRun(page);

    for (const tab of [
      'damage',
      'buffs',
      'debuffs',
      'casts',
      'resources',
      'timeline',
      'sample',
      'distribution',
    ]) {
      const box = await page.getByTestId(`sim-tab-${tab}`).boundingBox();
      expect(box?.height ?? 0, tab).toBeGreaterThanOrEqual(44);
    }
  });
});
