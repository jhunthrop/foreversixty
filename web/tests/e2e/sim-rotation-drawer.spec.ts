// web/tests/e2e/sim-rotation-drawer.spec.ts
// Task 6 (newcomer BLOCKER, tank MAJOR): "what it does" beside ROTATION used to link to
// /sim/specs#<spec>, which explains parse fidelity rather than the rotation, and threw away
// the loaded character and the finished run on the way -- browser Back did not restore
// either. This proves the fix from the browser, on the reviewers' own repro (newcomer
// review, load a Frost Mage, run a sim, click "what it does"): the trigger opens an
// in-page drawer, the character and the result stay on screen, and the URL never changes.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

// The fixture data build (FOREVER_DATA=fixture) ships only a warrior talent tree
// (src/fixtures/planner/talents/warrior.json) -- sim-specs.spec.ts's own ROGUE_TALENTS
// established the pattern this follows: stub /data/<build>/talents/mage.json with the
// smallest legal three-tree file (a mage always has three) that decodes to mage-frost, one
// point in a tier-0 talent of the third tree, no prereq.
const MAGE_TALENTS = {
  build: activeBuild.build,
  class_id: 8,
  class_slug: 'mage',
  trees: [
    { id: 81, name: 'Arcane', position: 0, background: 'fixture_arcane', talents: [] },
    { id: 82, name: 'Fire', position: 1, background: 'fixture_fire', talents: [] },
    {
      id: 83,
      name: 'Frost',
      position: 2,
      background: 'fixture_frost',
      talents: [
        {
          id: 831,
          name: 'Frost Warding',
          icon: 'fixture_frost_warding',
          max_rank: 2,
          tier: 0,
          column: 0,
          prereq_talent_id: null,
          prereq_rank: null,
          spell_id: 11189,
          ranks: Array.from({ length: 2 }, (_, i) => ({
            spell_id: 11189 + i,
            description: `Rank ${i + 1}.`,
          })),
        },
      ],
    },
  ],
};
const FROST_FS1 = `FS1:${activeBuild.build}:mage:gnome:0/0/1:`;

async function loadFrostMage(page: Page): Promise<void> {
  await page.route('**/data/*/talents/mage.json', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(MAGE_TALENTS) }),
  );
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FROST_FS1);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

test('"what it does" opens the rotation in a drawer, without losing the character or the run', async ({
  page,
}) => {
  await loadFrostMage(page);

  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 5000 });
  const finishedFigure = await page.getByTestId('sim-dps').textContent();
  const url = page.url();

  const card = page.getByTestId('sim-rotation-card');
  await card.getByTestId('sim-rotation-card-link').click();

  const panel = card.getByTestId('sim-rotation-drawer-panel');
  await expect(panel).toBeVisible();
  await expect(panel.getByTestId('sim-rotation-drawer-steps')).toContainText(
    'Frostbolt is the whole rotation.',
  );

  // The point of the fix: no navigation happened at all, so nothing was there to lose.
  expect(page.url()).toBe(url);
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-dps')).toHaveText(finishedFigure ?? '');

  // The fidelity detail is still one click away, in a new tab rather than carrying this one
  // away.
  const fidelityLink = panel.getByTestId('sim-rotation-drawer-fidelity-link');
  await expect(fidelityLink).toHaveAttribute('href', '/sim/specs#mage-frost');
  await expect(fidelityLink).toHaveAttribute('target', '_blank');
});

// Escape closes the drawer and returns focus to the trigger -- Disclosure.svelte's own
// contract (Ruling 3), proven here rather than assumed for this caller.
test('Escape closes the rotation drawer and returns focus to the trigger', async ({ page }) => {
  await loadFrostMage(page);
  // RotationCard only renders once a run has finished (SimView.svelte: `store.result !==
  // null`), so the drawer trigger does not exist until this.
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 5000 });

  const card = page.getByTestId('sim-rotation-card');
  const trigger = card.getByTestId('sim-rotation-card-link');
  await trigger.click();
  await expect(card.getByTestId('sim-rotation-drawer-panel')).toBeVisible();

  await page.keyboard.press('Escape');
  await expect(card.getByTestId('sim-rotation-drawer-panel')).toHaveCount(0);
  await expect(trigger).toBeFocused();
});

// task-6-brief.md: "/sim/specs' own per-spec card gains the same rotation prose, so the
// page a visitor lands on from the strip answers 'what does this rotation do' too."
test('/sim/specs carries the same rotation prose on the mage-frost card', async ({ page }) => {
  const specFixture = JSON.parse(
    readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'fixtures', 'sim', 'specs.json'), 'utf8'),
  ) as unknown[];
  await page.route('**/v1/specs**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { specs: specFixture }, error: null, request_id: 'req-test' }),
    }),
  );

  await page.goto('/sim/specs');
  const card = page.getByTestId('spec-mage-frost');
  await expect(card).toBeVisible();
  await card.getByTestId('spec-rotation-trigger').click();
  await expect(card.getByTestId('spec-rotation-steps')).toContainText('Frostbolt is the whole rotation.');
  // Same trigger text as RotationCard's own -- one voice for "what does this rotation do".
  await expect(card.getByTestId('spec-rotation-trigger')).toHaveText(simCopy.rotationLink);
});
