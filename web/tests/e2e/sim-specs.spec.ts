import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';
// Not `dpsSpecs` from spec-label.ts: that module also imports classRows from
// planner/reference.ts, which does a raw `import combosJson from
// '../../data/generated/combos.json'` -- fine for every in-app import, which goes through
// Vite's own JSON handling, but Playwright's own Node runtime loads a spec file's imports
// itself and Node needs an import attribute for a bare `.json` import that nothing here
// carries. Reading SPECS directly and filtering it the way dpsSpecs() does avoids the whole
// chain while still reading the canonical list rather than a literal.
import { SPECS } from '../../src/lib/sim/specs';

const dpsSpecCount = SPECS.filter((spec) => spec.role === 'dps').length;

// Read as JSON rather than imported as an ES module: Playwright's own Node runtime needs an
// import attribute this repo's other e2e specs do not carry for a plain `.json` import, so
// this reads the files the way tests/e2e/sim-sources.spec.ts and src/test-support/sim-api.ts
// already do.
const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const specFixture = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'fixtures', 'sim', 'specs.json'), 'utf8'),
) as unknown[];

function envelope(data: unknown) {
  return { ok: true, data, error: null, request_id: 'req-test' };
}

function failure(message: string) {
  return { ok: false, data: null, error: { message }, request_id: 'req-test' };
}

async function stubSpecs(page: Page): Promise<void> {
  await page.route('**/v1/specs**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify(envelope({ specs: specFixture })),
    }),
  );
}

test.describe('/sim/specs', () => {
  test('renders one card per dps spec, from the canonical list, whatever the API said', async ({ page }) => {
    await stubSpecs(page);
    await page.goto('/sim/specs');

    const cards = page.locator('[data-testid^="spec-"][id]');
    await expect(cards).toHaveCount(dpsSpecCount);

    // Tanks and healers are not on this page at all -- the design's launch scope is Quick
    // Sim for damage specs, so a card for one of them would promise something deferred.
    await expect(page.getByTestId('spec-warrior-protection')).toHaveCount(0);
    await expect(page.getByTestId('spec-druid-restoration')).toHaveCount(0);

    const fury = page.getByTestId('spec-warrior-fury');
    await expect(fury.getByTestId('spec-state')).toHaveText('Validated');
    await expect(fury.getByTestId('spec-figure')).toHaveText('Median gap 3.1% over 50 parses');

    const frost = page.getByTestId('spec-mage-frost');
    await expect(frost.getByTestId('spec-state')).toHaveText('In progress');

    // druid-feral is a dps spec the fixture's own /v1/specs answer says nothing about: the
    // grid still renders it, as "Not yet" with no figure, rather than leaving it out.
    const feral = page.getByTestId('spec-druid-feral');
    await expect(feral.getByTestId('spec-state')).toHaveText('Not yet');
    await expect(feral.getByTestId('spec-figure')).toHaveText(simCopy.specNoParses);

    // C1: rogue-combat is the fixture's own default (unmeasured) card -- null updated_at,
    // '' engine_version -- exactly the shape the real API sends for every spec the nightly
    // job has not measured. Before the fix, SpecCard's unguarded `.slice(0, 10)` threw here
    // and took the whole grid down with it; the assertion above (every card rendered) would
    // never have been reached. This one names the card and asserts it renders no footer at
    // all, rather than a footer built from a null date.
    const rogue = page.getByTestId('spec-rogue-combat');
    await expect(rogue.getByTestId('spec-state')).toHaveText('Not yet');
    await expect(rogue.getByTestId('spec-footer')).toHaveCount(0);
  });

  test('a hash link scrolls straight to that spec’s card', async ({ page }) => {
    await stubSpecs(page);
    await page.goto('/sim/specs#warrior-fury');

    await expect(page.getByTestId('spec-warrior-fury')).toBeInViewport();
  });

  test('a failed read shows one alert and a way to try again', async ({ page }) => {
    await page.route('**/v1/specs**', (route) =>
      route.fulfill({ status: 500, contentType: 'application/json', body: JSON.stringify(failure('boom')) }),
    );
    await page.goto('/sim/specs');

    const alert = page.getByTestId('specs-error');
    await expect(alert).toBeVisible();
    await expect(alert).toContainText(simCopy.specsFailed);
    await expect(page.getByTestId('specs-retry')).toHaveText(simCopy.tryAgain);
  });
});

test.describe('the in-page fidelity note', () => {
  // Combat Rogue: the fixture's data build only ships a warrior talent tree
  // (src/fixtures/planner/talents/warrior.json is the one file FOREVER_DATA=fixture
  // publishes), so this stubs `/data/<build>/talents/rogue.json` with the smallest legal
  // two-tree file that decodes to rogue-combat -- one point, no prereq, in a tier-0 talent
  // of the second tree. rogue-combat is the fixture /v1/specs answer's own unsupported row.
  const ROGUE_TALENTS = {
    build: activeBuild.build,
    class_id: 4,
    class_slug: 'rogue',
    trees: [
      { id: 401, name: 'Assassination', position: 0, background: 'fixture_assassination', talents: [] },
      {
        id: 402,
        name: 'Combat',
        position: 1,
        background: 'fixture_combat',
        talents: [
          {
            id: 4021,
            name: 'Improved Sinister Strike',
            icon: 'fixture_improved_sinister_strike',
            max_rank: 5,
            tier: 0,
            column: 0,
            prereq_talent_id: null,
            prereq_rank: null,
            spell_id: 40211,
            ranks: Array.from({ length: 5 }, (_, i) => ({
              spell_id: 40211 + i,
              description: `Rank ${i + 1}.`,
            })),
          },
        ],
      },
    ],
  };
  const ROGUE_FS1 = `FS1:${activeBuild.build}:rogue:orc:0/1/0:`;

  // Fury Warrior: the fixture's data build ships this class's talent tree directly
  // (src/fixtures/planner/talents/warrior.json), so loading it needs no route stub -- the
  // same character sim-run.spec.ts and sim-settings.spec.ts load. It is the fixture
  // /v1/specs answer's own validated row.
  const FURY_FS1 = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

  test('an unmeasured spec still runs, with a label linking to /sim/specs instead of a blocking card', async ({
    page,
  }) => {
    await stubSpecs(page);
    await page.route('**/data/*/talents/rogue.json', (route) =>
      route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ROGUE_TALENTS) }),
    );

    await page.goto('/sim');
    await page.getByTestId('sim-addon-input').fill(ROGUE_FS1);
    await page.getByTestId('sim-addon-load').click();
    await expect(page.getByTestId('sim-character')).toBeVisible();

    await expect(page.getByTestId('sim-run-button')).toHaveText('Run sim');
    const note = page.getByTestId('spec-fidelity-note');
    await expect(note).toContainText(simCopy.specNotYet);
    await expect(note).toContainText(simCopy.specNotYetNote);
  });

  // Validated is the one state with nothing to say -- the settings bar above already named
  // the spec being simulated, so a note reading "Validated ..." beside it would be noise.
  test('a validated spec runs with no fidelity note at all', async ({ page }) => {
    await stubSpecs(page);

    await page.goto('/sim');
    await page.getByTestId('sim-addon-input').fill(FURY_FS1);
    await page.getByTestId('sim-addon-load').click();
    await expect(page.getByTestId('sim-character')).toBeVisible();

    await expect(page.getByTestId('sim-run-button')).toHaveText('Run sim');
    await expect(page.getByTestId('spec-fidelity-note')).toHaveCount(0);
  });
});
