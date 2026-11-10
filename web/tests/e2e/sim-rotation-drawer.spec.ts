// web/tests/e2e/sim-rotation-drawer.spec.ts
// Task 6 (newcomer BLOCKER, tank MAJOR): "what it does" beside ROTATION used to link to
// /sim/specs#<spec>, which explains parse fidelity rather than the rotation, and threw away
// the loaded character and the finished run on the way -- browser Back did not restore
// either.
//
// Final whole-branch review, C1: both reviewers' own repros -- newcomer: load a Frost Mage,
// click "what it does"; tank: load an Arms Warrior, click "what it does" -- load a
// character and click, without ever running a sim first. RotationCard, whose Disclosure
// this suite originally tested, does not render until a run finishes (SimView.svelte:
// `store.result !== null`), so a test that ran a sim before clicking never actually
// exercised the reviewers' repro -- it exercised the settings bar's still-broken plain
// `<a>`'s already-safe successor instead, once main had one. The tests below run the
// reviewers' own repro on the settings bar's own trigger (`sim-rotation-link`), the one
// drawer on /sim now that a finished run shows the report's always-open Rotation card.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { curatedRotationNotes } from './support/rotation-notes';
import { gotoSim, simCharacter } from './support/sim-page';

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
  // A code link loads the character without running it: the rail's Rotation section before a
  // run, the moment the newcomer's repro was about.
  await gotoSim(page, `/sim?code=${encodeURIComponent(FROST_FS1)}`);
  await expect(simCharacter(page)).toBeVisible();
}

// The tank review's own repro: an Arms Warrior. No talent stub needed -- the fixture data
// build ships the warrior tree directly (src/fixtures/planner/talents/warrior.json),
// exactly the one sim-run.spec.ts's own FURY_FS1 already loads with no route at all.
// `3/0/0` spends 3 points on Improved Heroic Strike (id 1001, tier 0, column 0, no
// prereq, max_rank 3) -- the fixture's own first Arms talent, so `specForSplit` resolves
// warrior-arms the same way FURY_FS1's `0/5530515/0` resolves warrior-fury. The trailing
// `/0` third segment is never read: `orderFromRanks` (fs1.ts) walks only `index.trees`,
// and this fixture's warrior.json carries two trees (Arms, Fury), not three -- the same
// reason FURY_FS1's own all-zero third segment needs no Protection tree to exist either.
const ARMS_FS1 = `FS1:${activeBuild.build}:warrior:orc:3/0/0:head=12640,main_hand=11726`;

async function loadArmsWarrior(page: Page): Promise<void> {
  await gotoSim(page, `/sim?code=${encodeURIComponent(ARMS_FS1)}`);
  await expect(simCharacter(page)).toBeVisible();
}

// The newcomer's own repro, verbatim at the time: /sim -> load a Frost Mage -> "what it does"
// used to navigate away. The rotation is the rail's open Rotation section now (spec
// 2026-11-06-sim §4.3), each line with its note in place: there is no drawer to open, no
// navigation, and nothing to lose. The fidelity detail is the answer's caption (§4.6).
test('the rotation reads in place in the rail, without losing the character (newcomer repro)', async ({
  page,
}) => {
  await loadFrostMage(page);

  const url = page.url();
  const section = page.getByTestId('sim-rotation-saved-card');
  await section.scrollIntoViewIfNeeded();
  await expect(section).toBeVisible();
  await expect(section).toContainText('Frost guide');
  const notes = curatedRotationNotes('mage-frost');
  expect(notes.length).toBeGreaterThan(0);
  for (const note of notes) await expect(section).toContainText(note);

  // Nothing navigated, so nothing was there to lose.
  expect(page.url()).toBe(url);
  await expect(simCharacter(page)).toBeVisible();
  await expect(section.getByRole('link', { name: /Frost guide/ })).toHaveAttribute(
    'href',
    '/guides/mage/frost',
  );
});

// The tank's own repro (tank review, MAJOR), verbatim at the time: /sim -> load an Arms
// Warrior -> "what it does" used to land on /sim/specs#warrior-arms with the character gone.
test('the tank’s repro: an Arms Warrior’s rotation reads in place, with no navigation needed', async ({
  page,
}) => {
  await loadArmsWarrior(page);

  const url = page.url();
  const section = page.getByTestId('sim-rotation-saved-card');
  await section.scrollIntoViewIfNeeded();
  await expect(section).toContainText(
    'Whenever rage is under 80: 10 rage at once and 10 more over 10 seconds, back every minute',
  );

  expect(page.url()).toBe(url);
  await expect(simCharacter(page)).toBeVisible();
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
  // Final whole-branch review, I2: suffixed with the spec so the 20 damage-spec cards on
  // this page never share one testid (or one unscoped accessible name) between them.
  await card.getByTestId('spec-rotation-trigger-mage-frost').click();
  const steps = card.getByTestId('spec-rotation-steps-mage-frost');
  for (const note of curatedRotationNotes('mage-frost')) await expect(steps).toContainText(note);
  // h-audit 2026-11-09: the spec's own name is the trigger, not a "What it does" on every row.
  await expect(card.getByTestId('spec-rotation-trigger-mage-frost')).toHaveText('Frost Mage');
});
