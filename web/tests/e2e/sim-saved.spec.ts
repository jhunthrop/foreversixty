// web/tests/e2e/sim-saved.spec.ts
// Task 17: /sim/<sim_id> (read-only, from the prerendered fixture and from a stubbed shell
// for an id nothing prerendered) and the save flow on /sim itself.
//
// Read as JSON rather than imported as an ES module: Playwright's own Node runtime needs an
// import attribute this repo's other e2e specs do not carry for a plain `.json` import, the
// same reason sim-sources.spec.ts and sim-specs.spec.ts read their own fixtures this way.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { curatedRotationNotes } from './support/rotation-notes';
import {
  expectSimGearSlot,
  gotoSavedSim,
  gotoSim,
  openSimSection,
  simCharacter,
  simGearDoll,
} from './support/sim-page';
import { expect, test } from './support/test';
import type { CharacterSpec } from '../../src/lib/sim/types';

const ROOT = path.join(import.meta.dirname, '..', '..');
const activeBuild = JSON.parse(readFileSync(path.join(ROOT, 'src', 'data', 'active-build.json'), 'utf8')) as {
  build: string;
};
const fixtureResult = JSON.parse(
  readFileSync(path.join(ROOT, 'src', 'fixtures', 'sim', 'result.json'), 'utf8'),
) as {
  dps: { mean: number };
  engine_version: string;
  iterations_run: number;
  duration_ms: number;
  lane: 'browser' | 'server';
  request: { spec: string; character: CharacterSpec };
};

// Task 20: a saved Top Gear/talents/drops result and a saved weights result, at /sim/<id>.
// Read as plain JSON (not an ES import) for the same reason `fixtureResult` above is --
// Playwright's own Node loader needs an import attribute this repo's other e2e specs never
// carry for a bare `.json` import.
const bulkFixture: Record<string, unknown> = JSON.parse(
  readFileSync(path.join(ROOT, 'src', 'fixtures', 'sim', 'bulk-result.json'), 'utf8'),
) as Record<string, unknown>;
const weightsFixture: Record<string, unknown> = JSON.parse(
  readFileSync(path.join(ROOT, 'src', 'fixtures', 'sim', 'weights-result.json'), 'utf8'),
) as Record<string, unknown>;

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'req-test' }),
  };
}

function failure(message: string, status: number) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: false, data: null, error: { message }, request_id: 'req-test' }),
  };
}

// The minimal shell `sim.astro` (and the Worker's live re-serve of the same static asset
// for any other id) emits: a mount div carrying no `data-sim-id`, so the island falls back
// to reading the id off `window.location.pathname` -- `sim-island.ts`'s own `simIdFor`.
// `astro preview` has no Worker in front of it (that rewrite is Cloudflare-only), so a
// non-prerendered `/sim/<id>` 404s unless the navigation itself is stubbed, the same reason
// shared-build.spec.ts stubs `/b/k7x2qm4a` rather than navigating to it directly.
const SHELL_HTML = `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><title>Saved sim</title>
<link rel="stylesheet" href="/sim-island.css"></head>
<body><main id="main">
<div id="sim" data-sim-mount></div>
<script type="module" src="/sim-island.js"></script>
</main></body></html>`;

// The saved page (spec 2026-11-06-sim §4.5, rulings S1 and S5): the head carries the title, the
// run's facts in its meta line, "Saved {date}", and the answer; then the report body. No tool
// tabs, no run line, no section nav (ruling S4).
test.describe('a saved sim from the prerendered fixture', () => {
  test('renders read-only: the head, the answer and every section on one page, no tabs', async ({ page }) => {
    await gotoSavedSim(page, '/sim/simfixtureab');

    // Read from the fixture rather than a literal (fixture values drift as the engine
    // golden is refreshed).
    const figure = fixtureResult.dps.mean.toLocaleString('en-US', {
      minimumFractionDigits: 1,
      maximumFractionDigits: 1,
    });
    await expect(page.getByTestId('sim-dps')).toHaveText(
      new RegExp(`^\\s*${figure.replace('.', '\\.')}\\s*DPS\\s*$`),
    );
    await expect(page.getByTestId('sim-head-meta')).toContainText('3:00');
    // A saved result names no lane and no label: its date is the meta's, its runs the report's.
    await expect(page.getByTestId('sim-answer-label')).toHaveText('');

    await expect(page.getByTestId('sim-damage-card')).toBeVisible();
    await expect(page.getByTestId('sim-section-nav')).toHaveCount(0);
    await expect(page.getByTestId('sim-tab-damage')).toHaveCount(0);
    await expect(page.getByTestId('sim-nav-tabs')).toHaveCount(0);
    await expect(page.getByTestId('sim-settings')).toHaveCount(0);
    await expect(page.getByTestId('sim-change-source')).toHaveCount(0);

    // An older result: the range strip stands in for the bins, and says so.
    await expect(page.getByTestId('sim-spread-range')).toBeVisible();
    await expect(page.getByTestId('sim-spread-plot')).toHaveCount(0);

    // The card below the fold mounts as it nears the viewport.
    await openSimSection(page, 'settings');
    await expect(page.getByTestId('sim-settings-runs')).toContainText(
      fixtureResult.iterations_run.toLocaleString('en-US'),
    );
    await expect(page.getByTestId('sim-saved-copy-for-game')).toBeVisible();
    await expect(page.getByTestId('sim-copy-json')).toBeVisible();
  });

  test('the full fixture draws the real bins, the saved date and a sortable table', async ({
    page,
    isMobile,
  }) => {
    await gotoSavedSim(page, '/sim/simfullfixt1');
    await expect(page.getByTestId('sim-spread-plot')).toBeVisible();
    await expect(page.getByTestId('sim-spread-bar').first()).toBeVisible();
    await expect(page.getByTestId('sim-spread-caption')).toHaveCount(0);
    await expect(page.getByTestId('sim-head-source')).toHaveText(/^Saved 4 Nov( 2026)?$/);

    // Below 1024 the table is three columns with no heads (spec §4.7, phone), so the sort
    // heads are a desktop control: a press on Crit sorts it, a second press flips it.
    if (!isMobile) {
      const header = page.locator('[role="columnheader"]', { has: page.getByTestId('stat-sort-crit') });
      await page.getByTestId('stat-sort-crit').click();
      await expect(header).toHaveAttribute('aria-sort', 'descending');
      await page.getByTestId('stat-sort-crit').click();
      await expect(header).toHaveAttribute('aria-sort', 'ascending');
    }
  });

  // D48 (dps-minmaxer review round 2, BLOCKER): a saved run loaded cold used to hardcode
  // `actionNames={null}` (SavedSim.svelte), so the build's own name table was never
  // fetched and the headline sentence read "spell:20662", live and saved disagreeing --
  // the defect's own name. This pins that a saved page fetches and uses the same
  // `loadActionNames` a live run does, exactly as `sim-results.spec.ts`'s own sentence
  // test exercises for a live run.
  test('a saved sim resolves ability and buff names from the build’s own table, not the raw action key', async ({
    page,
  }) => {
    await page.route(`**/data/${activeBuild.build}/simnames/warrior.json`, (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          spell: { '1680': 'Whirlwind', '25289': 'Battle Shout', '2687': 'Bloodrage' },
          item: {},
        }),
      }),
    );

    await gotoSavedSim(page, '/sim/simfixtureab');

    await expect(page.getByTestId('ability-row-spell:1680').getByTestId('ability-name')).toHaveText(
      'Whirlwind',
    );
    await expect(page.getByTestId('sim-damage-card')).not.toContainText(/\bspell:/);
  });

  // H4 (final whole-branch review): the fixture's saved sim is addon-sourced with an empty
  // ref (src/fixtures/sim/result.json's request.source is {kind:'addon', ref:''}), exactly
  // the shape that used to land on a dead `/sim?source=addon` with no character and no
  // message. The fix adopts the saved request's own request.character through /sim's
  // existing ?code= bootstrap instead, so a character is on screen after the click.
  test('Open in the sim opens /sim with the saved request’s own character, not a dead ref', async ({
    page,
  }) => {
    await gotoSavedSim(page, '/sim/simfixtureab');
    await page.getByTestId('sim-run-yourself').click();
    await expect(page).toHaveURL(/\/sim\?code=/);
    await expect(simCharacter(page)).toBeVisible();
    // The stored request's own character, not a guess: the same class and gear the saved
    // sim ran with (src/fixtures/sim/result.json's request.character).
    await expectSimGearSlot(page, 'head');
  });
});

// A saved sim is somebody else's result opened cold: no character is "loaded" here, and the
// chip's empty line ("No character loaded...") sat directly above a full result. The casts
// tab also carried the log viewer's note about other players' rows, which a one-player
// simulated fight does not have. Neither the old chip nor the retired spine bar is on the page.
test('a saved sim opened cold says nothing about a missing character or other players', async ({ page }) => {
  await gotoSavedSim(page, '/sim/simfixtureab');
  await expect(page.getByTestId('sim-dps')).toBeVisible();

  await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
  await expect(page.getByTestId('current-character-chip')).toHaveCount(0);

  await expect(page.getByTestId('sim-saved')).toBeVisible();
  await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
  await expect(page.getByTestId('current-character-chip')).toHaveCount(0);
});

/** The spec list's answer for warrior-fury, in one fidelity state. */
function fidelityRoute(state: 'validated' | 'in_progress') {
  return envelope({
    specs: [
      {
        spec: 'warrior-fury',
        state,
        median_gap: state === 'validated' ? 0.02 : 0.08,
        parses: state === 'validated' ? 50 : 12,
        worst_actions: [],
        engine_version: activeBuild.build,
        updated_at: '2026-01-01',
      },
    ],
  });
}

// Spec 2026-11-04-sim-results 3.11: the saved page's Rotation section is an always-open list (no
// drawer). It names the rotation the run used with the run's spec and level, prints each line's
// note in place and links the spec's guide. The spec's fidelity note, when the sim has not
// validated the spec (tenet 8), is the answer's caption now (spec 2026-11-06-sim §4.6). The
// prerendered fixture's own client-side fetch of /v1/specs is stubbed, as sim-specs.spec.ts
// stubs it, so the note has a row to read.
test('a saved sim names the rotation it used, prints its notes in place, and carries its fidelity', async ({
  page,
}) => {
  await page.route('**/v1/specs', (route) => route.fulfill(fidelityRoute('in_progress')));
  await gotoSavedSim(page, '/sim/simfixtureab');
  await openSimSection(page, 'rotation');
  const card = page.getByTestId('sim-rotation-saved-card');
  await expect(card).toBeVisible();
  await expect(card).toContainText('Fury guide');
  await expect(card).toContainText('Then, from the top at every decision');

  const notes = curatedRotationNotes('warrior-fury');
  await expect(card.getByTestId('rotation-line').first()).toBeVisible();
  // A note that only says "on cooldown" is dropped (h-audit, lib/sim/rotation-card.ts obviousNote):
  // Bloodthirst's line carries no note; a note that says when to press is printed in place.
  const bloodthirst = notes.find((note) => note.startsWith('On cooldown: the button'));
  expect(bloodthirst).toBeDefined();
  await expect(card).not.toContainText(bloodthirst!);
  const execute = notes.find((note) => note.startsWith('In the Execute phase, or in the last 15 seconds'));
  expect(execute).toBeDefined();
  await expect(card).toContainText(execute!);

  const url = page.url();
  const guide = card.getByRole('link', { name: /Fury guide/ });
  await expect(guide).toHaveAttribute('href', '/guides/warrior/fury');
  expect(page.url()).toBe(url);

  // The fixture's warrior-fury row is not validated, so the note is there, pill linking the spec list.
  const note = page.getByTestId('sim-answer').getByTestId('spec-fidelity-note');
  await expect(note).toBeVisible();
  await expect(note.getByRole('link')).toHaveAttribute('href', '/sim/specs#warrior-fury');
});

// A validated spec renders no note at all, rather than a green "all is well" line: dropping the
// `!== 'validated'` clause in needsFidelityNote would break nothing without this.
test('a saved sim for a validated spec carries the rotation section with no fidelity note', async ({
  page,
}) => {
  await page.route('**/v1/specs', (route) => route.fulfill(fidelityRoute('validated')));
  await gotoSavedSim(page, '/sim/simfixtureab');
  await openSimSection(page, 'rotation');
  const card = page.getByTestId('sim-rotation-saved-card');
  await expect(card).toBeVisible();
  await expect(card.getByTestId('rotation-line').first()).toBeVisible();
  await expect(page.getByTestId('spec-fidelity-note')).toHaveCount(0);
});

test('a saved sim whose stored request has no gear says so and draws no doll', async ({ page }) => {
  const id = 'simnogearabc';
  const noGearResult = {
    ...fixtureResult,
    sim_id: id,
    request: { ...fixtureResult.request, character: { ...fixtureResult.request.character, gear: [] } },
  };
  await page.route(`**/sim/${id}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
  );
  await page.route(`**/v1/sims/${id}`, (route) => route.fulfill(envelope(noGearResult)));

  await page.goto(`/sim/${id}`);

  await expect(page.getByTestId('sim-saved')).toBeVisible();
  await page.getByTestId('sim-section-gear').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('sim-gear-empty')).toBeVisible();
  await expect(simGearDoll(page)).toHaveCount(0);
});

// Defect fix: the name typed into "Name this sim" (SaveSimForm.svelte) used to be discarded
// well before the page -- GET /v1/sims/{id} dropped it entirely, so `result.title` was
// always undefined here and the heading fell back to the composed spec line unconditionally.
// This pins the page's own half of the fix: given a `title` on the fetched result (what a
// fixed GET now sends), the heading shows it, not the fallback. The field report's own
// string, exercised end to end.
test('a saved sim heads with the name a member gave it, not the composed spec line', async ({ page }) => {
  const id = 'simtitledabc';
  const named = 'Thoradin - Fury Warrior, raid-buffed BWL night';
  const titledResult = { ...fixtureResult, sim_id: id, title: named };
  await page.route(`**/sim/${id}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
  );
  await page.route(`**/v1/sims/${id}`, (route) => route.fulfill(envelope(titledResult)));

  await page.goto(`/sim/${id}`);

  await expect(page.getByTestId('page-title')).toHaveText(named);
});

test.describe('a saved sim renders the results view its own kind calls for', () => {
  test('a saved Top Gear result renders its ranked table, not a damage breakdown', async ({ page }) => {
    const id = 'simbulk23456';
    await page.route(`**/sim/${id}`, (route) =>
      route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
    );
    await page.route(`**/v1/sims/${id}`, (route) => route.fulfill(envelope({ ...bulkFixture, sim_id: id })));

    await page.goto(`/sim/${id}`);

    await expect(page.getByTestId('sim-combos')).toBeVisible();
    await expect(page.getByTestId('sim-combo-row').first()).toBeVisible();
    // The plain-run damage-breakdown view (SimResults.svelte's own `data-testid="sim-results"`)
    // never mounts for a bulk kind -- the kind switch picks exactly one results view.
    await expect(page.getByTestId('sim-results')).toHaveCount(0);
  });

  test('a saved weights result renders its scale-factor table', async ({ page }) => {
    // The brief's own draft id ("simweight3456") is 13 characters -- one over
    // `SIM_ID_PATTERN`'s `[a-z2-7]{12}` (contract: a sim_id is always exactly 12) -- and
    // silently fails to match, which routes the page to the plain, no-character /sim
    // landing state instead of a saved sim at all. This id is the same twelve characters,
    // trimmed to fit the real pattern.
    const id = 'simweight345';
    await page.route(`**/sim/${id}`, (route) =>
      route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
    );
    await page.route(`**/v1/sims/${id}`, (route) =>
      route.fulfill(envelope({ ...weightsFixture, sim_id: id })),
    );

    await page.goto(`/sim/${id}`);

    await expect(page.getByTestId('sim-weights')).toBeVisible();
    await expect(page.getByTestId('sim-weights-table')).toContainText('1.00');
    // A weights result has no gear story (design 7): neither doll form claims one.
    await expect(page.locator('[data-testid^="paperdoll-"][data-testid$="-slot-head"]')).toHaveCount(0);
  });
});

// Round 3 (Lighthouse): the skeleton reserves the saved result's shape while a real,
// non-prerendered id's GET /v1/sims/{id} is in flight -- sim/[id].astro's own shell (the
// one Lighthouse measures) renders the identical string statically and never shows this
// branch for the prerendered fixture (its result is inlined), so this is the only place
// the loading skeleton is exercisable end to end. Mirrors report-phone.spec.ts's own
// reserved-skeleton check for reports/[id].astro's report-skeleton.
test('a saved sim shows the loading skeleton until the fetch resolves, then the real result', async ({
  page,
}) => {
  const id = 'simloadingab';
  await page.route(`**/sim/${id}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
  );
  let release: (() => void) | undefined;
  const held = new Promise<void>((resolve) => {
    release = resolve;
  });
  await page.route(`**/v1/sims/${id}`, async (route) => {
    await held;
    await route.fulfill(envelope({ ...fixtureResult, sim_id: id }));
  });

  await page.goto(`/sim/${id}`);

  await expect(page.getByTestId('sim-saved-skeleton')).toBeVisible();
  await expect(page.getByTestId('sim-saved')).toHaveCount(0);

  release?.();

  await expect(page.getByTestId('sim-saved')).toBeVisible();
  await expect(page.getByTestId('sim-saved-skeleton')).toHaveCount(0);
});

// The saved-sim fetch is a named, retriggerable function (createSavedSimState().load)
// rather than a dead-end anonymous block, so a failed GET /v1/sims/{id} offers a real
// retry through LoadError -- the same request re-fired in place, not a page reload.
test('a failed saved-sim fetch offers a retry that re-fires the same request', async ({ page }) => {
  const id = 'simfailonceb';
  await page.route(`**/sim/${id}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
  );
  let attempts = 0;
  await page.route(`**/v1/sims/${id}`, async (route) => {
    attempts += 1;
    if (attempts === 1) {
      await route.fulfill(failure('boom', 500));
    } else {
      await route.fulfill(envelope({ ...fixtureResult, sim_id: id }));
    }
  });

  await page.goto(`/sim/${id}`);

  await expect(page.getByTestId('sim-saved-error')).toBeVisible();
  await page.getByTestId('sim-saved-error-retry').click();
  await expect(page.getByTestId('sim-saved-error')).not.toBeVisible();
});

test('a stale engine version raises the note and names the engine, and is never re-run automatically', async ({
  page,
}) => {
  const id = 'simstaleabcd';
  const staleResult = { ...fixtureResult, sim_id: id, engine_version: '6a1c2d9' };
  let calls = 0;
  await page.route(`**/sim/${id}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
  );
  await page.route(`**/v1/sims/${id}`, (route) => {
    calls += 1;
    return route.fulfill(envelope(staleResult));
  });

  await page.goto(`/sim/${id}`);

  await expect(page.getByTestId('sim-note-stale')).toBeVisible();
  await openSimSection(page, 'settings');
  // This run's Engine fact names the old build (the stale tag beside it retired).
  await expect(page.getByTestId('sim-settings-engine')).toContainText('6a1c2d9');

  // No auto re-run: the remedy is the "Open in the sim" link, never a second fetch fired
  // on its own.
  await page.waitForTimeout(500);
  expect(calls).toBe(1);
});

// Persona review 2026-10-06 (retail-raider), §5/§8 item 1: SimSavePanel now requires
// sign-in before Save enables (ColdPasteHero.svelte's cold-paste run never needed one, but
// saving the result still does), so this describe block's own save flow needs a signed-in
// `/v1/me` -- a cold-paste run reaching a result stays covered, signed out, by
// sim-cold-paste.spec.ts.
const SIGNED_IN_ME = {
  user: { id: 1, battletag: 'Fixture#0001', email: null, role: 'user', anonymize: false, premium: false },
  characters: [],
  guilds: [],
};

test.describe('saving a browser run from /sim', () => {
  // sim-one-paste lane: ColdPasteHero.svelte's Run pastes and runs in one click, so the
  // separate `sim-run-button` click this helper used to make after Load is gone.
  async function runFury(page: import('@playwright/test').Page): Promise<void> {
    await page.route('**/v1/me', (route) => route.fulfill(envelope(SIGNED_IN_ME)));
    await gotoSim(page);
    await page.getByTestId('sim-cold-paste-input').fill(FURY);
    await page.getByTestId('sim-cold-paste-run').click();
    await expect(simCharacter(page)).toBeVisible();
    await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
  }

  test('opens pre-filled, saves, and shows the link without navigating', async ({ page }) => {
    await runFury(page);

    await expect(page.getByTestId('sim-save-open')).toBeEnabled();
    await page.getByTestId('sim-save-open').click();
    await expect(page.getByTestId('sim-save-title')).toHaveValue('Bare, 3:00, single target');

    await page.route('**/v1/sims', (route) => route.fulfill(envelope({ sim_id: 'simnew234567' }, 201)));
    await page.getByTestId('sim-save-confirm').click();

    // The Button becomes Copy link, the head's actions say Saved, and the link is the new sim's.
    await expect(page.getByTestId('sim-save-copy')).toHaveText('Copy link');
    await expect(page.getByTestId('sim-head')).toContainText('Saved');
    await expect(page.getByTestId('sim-save-link')).toHaveAttribute('href', /\/sim\/simnew234567$/);
    await expect(page).toHaveURL(/\/sim(\?.*)?$/);
    await expect(page.getByTestId('sim-view')).toBeVisible();
  });

  test('a failed save shows saveFailed and keeps the typed title', async ({ page }) => {
    await runFury(page);

    await page.getByTestId('sim-save-open').click();
    await page.getByTestId('sim-save-title').fill('My opener sim');

    await page.route('**/v1/sims', (route) => route.fulfill(failure('boom', 500)));
    await page.getByTestId('sim-save-confirm').click();

    await expect(page.getByTestId('sim-save-error')).toBeVisible();
    await expect(page.getByTestId('sim-save-title')).toHaveValue('My opener sim');
  });
});
