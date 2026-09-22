// web/tests/e2e/sim-saved.spec.ts
// Task 17: /sim/<sim_id> (read-only, from the prerendered fixture and from a stubbed shell
// for an id nothing prerendered) and the save flow on /sim itself.
//
// Read as JSON rather than imported as an ES module: Playwright's own Node runtime needs an
// import attribute this repo's other e2e specs do not carry for a plain `.json` import, the
// same reason sim-sources.spec.ts and sim-specs.spec.ts read their own fixtures this way.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';
import { plannerHrefForSpec } from '../../src/lib/sim/character';
import { simCopy } from '../../src/lib/sim/copy';
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

test.describe('a saved sim from the prerendered fixture', () => {
  test('renders read-only, with the results and no gear grid claim it cannot make', async ({ page }) => {
    await page.goto('/sim/simfixtureab');

    // Read from the fixture rather than a literal (fixture values drift as the engine
    // golden is refreshed; the plan's own "1,131" is stale against the checked-in file).
    const figure = Math.round(fixtureResult.dps.mean).toLocaleString('en-US');
    await expect(page.getByTestId('sim-dps')).toHaveText(figure);
    await expect(page.getByTestId('sim-saved-encounter')).toHaveText('raid-buffed, 3:00, single target');

    await expect(page.getByTestId('sim-results')).toBeVisible();
    await expect(page.getByTestId('sim-tab-damage')).toHaveAttribute('aria-selected', 'true');

    // Fix round 1: the header no longer states iterations/processing time/lane a second
    // time (they used to duplicate the details card below it) -- the card is now the one
    // place these live, so this is where the saved page's own figures are pinned down.
    await expect(page.getByTestId('sim-details-card')).toBeVisible();
    await expect(page.getByTestId('sim-details-iterations')).toHaveText(
      fixtureResult.iterations_run.toLocaleString('en-US'),
    );
    await expect(page.getByTestId('sim-details-processing')).toHaveText(
      `${(fixtureResult.duration_ms / 1000).toFixed(1)} s`,
    );
    await expect(page.getByTestId('sim-details-lane')).toHaveText(
      fixtureResult.lane === 'server' ? simCopy.detailsLaneServer : simCopy.detailsLaneBrowser,
    );
    await expect(page.getByTestId('sim-details-engine')).toHaveAttribute('href', '/sim/specs');

    // The fixture's stored request carries real gear entries, so `gearKnown` is true and
    // the strip renders the grid rather than `simCopy.savedNoGear` -- the inverse of a
    // request with no gear at all, exercised below against a stubbed id.
    await expect(page.getByTestId('sim-slot-head')).toBeVisible();
    await expect(page.getByTestId('sim-no-gear')).toHaveCount(0);

    // Controller addendum (Task 17's review, MEDIUM): the strip's "Change source" button
    // was rendering live and focusable on this read-only page but wired to a no-op. There
    // is nothing to change a source into here, so CharacterStrip's `readonly` prop hides it
    // rather than leaving a button that does nothing for a player to find.
    await expect(page.getByTestId('sim-change-source')).toHaveCount(0);
  });

  // Defect fix: SavedSim.svelte used to hand CharacterStrip a `character` with
  // `point_order: []`, which fell through to the strip's own `plannerHrefFor` -- built for
  // the live /sim page, where `point_order` is the truth -- and that function can only
  // encode zeroed talents from an empty order (`toCharacterSpec`'s own
  // `talentsString(index, [])`). A saved sim's own stored request already carries the true,
  // final talents string (result.json's fixture: `-5530515-`, a real Fury Warrior build,
  // not `----`), so the fix reaches it directly through `plannerHrefForSpec` instead of
  // reconstructing anything through a point order the page never had.
  test('"Open in planner" carries the saved sim’s own talents, not zeroed', async ({ page }) => {
    await page.goto('/sim/simfixtureab');
    await expect(page.getByTestId('sim-character')).toBeVisible();

    const expectedHref = plannerHrefForSpec(fixtureResult.request.character, activeBuild.build);
    // The fixture's own talents string is non-empty, so a correct link's code cannot be the
    // all-zero shape defect B produced (":0/0/0:").
    expect(fixtureResult.request.character.talents.replace(/-/g, '')).not.toBe('');
    expect(expectedHref).not.toContain(':0/0/0:');
    await expect(page.getByTestId('sim-open-planner')).toHaveAttribute('href', expectedHref);
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
          spell: { '1680': 'Whirlwind', '25289': 'Battle Shout' },
          item: {},
        }),
      }),
    );

    await page.goto('/sim/simfixtureab');

    // Same second ability and buff sim-results.spec.ts's live-run sentence test names, at
    // the same 78%/100% figures -- both read off src/fixtures/sim/result.json's summary.
    // The top ability is the fixture's own auto-attack row, which reads as prose
    // ("main-hand white hits") regardless of the name table.
    const sentence = page.getByTestId('sim-sentence');
    await expect(sentence).toHaveText(
      'main-hand white hits and Whirlwind are 78% of your damage; Battle Shout is up 100% of the fight.',
    );
    await expect(sentence).not.toContainText(/\bspell:/);
  });

  // H4 (final whole-branch review): the fixture's saved sim is addon-sourced with an empty
  // ref (src/fixtures/sim/result.json's request.source is {kind:'addon', ref:''}), exactly
  // the shape that used to land on a dead `/sim?source=addon` with no character and no
  // message. The fix adopts the saved request's own request.character through /sim's
  // existing ?code= bootstrap instead, so a character is on screen after the click.
  test('Run this yourself opens /sim with the saved request’s own character, not a dead ref', async ({
    page,
  }) => {
    await page.goto('/sim/simfixtureab');
    await page.getByTestId('sim-run-yourself').click();
    await expect(page).toHaveURL(/\/sim\?code=/);
    await expect(page.getByTestId('sim-character')).toBeVisible();
    // The stored request's own character, not a guess: the same class and gear the saved
    // sim ran with (src/fixtures/sim/result.json's request.character).
    await expect(page.getByTestId('sim-slot-head')).toBeVisible();
  });
});

// Task 13: the rotation card. The prerendered fixture's own client-side fetch of /v1/specs
// has to be stubbed, the same way sim-specs.spec.ts stubs it, or the fidelity note has
// nothing to render off of -- the card would show with no note and the second assertion
// below would be checking an element that never appears.
//
// Task 6: "what it does" no longer links straight to /sim/specs#<spec> -- it opens the
// rotation drawer in place, and the fidelity link moves inside that drawer. Clicking it
// open is this test's own proof the trigger is a Disclosure now, not a navigating anchor.
// A saved sim is somebody else's result opened cold: no character is "loaded" here, and the
// chip's empty line ("No character loaded...") sat directly above a full result. The casts
// tab also carried the log viewer's note about other players' rows, which a one-player
// simulated fight does not have.
test('a saved sim opened cold says nothing about a missing character or other players', async ({ page }) => {
  await page.goto('/sim/simfixtureab');
  await expect(page.getByTestId('sim-dps')).toBeVisible();

  await expect(page.getByTestId('sim-chip-slot')).toBeAttached();
  await expect(page.getByTestId('current-character-chip')).toHaveCount(0);

  await page.getByTestId('sim-tab-casts').click();
  const note = page.getByTestId('cast-time-note');
  await expect(note).toBeVisible();
  await expect(note).not.toContainText('other players');
});

test('a saved sim names the rotation it used, opens its drawer, and carries its fidelity', async ({
  page,
}) => {
  await page.route('**/v1/specs', (route) =>
    route.fulfill(
      envelope({
        specs: [
          {
            spec: 'warrior-fury',
            state: 'in_progress',
            median_gap: 0.08,
            parses: 12,
            worst_actions: [],
            engine_version: activeBuild.build,
            updated_at: '2026-01-01',
          },
        ],
      }),
    ),
  );

  await page.goto('/sim/simfixtureab');
  const card = page.getByTestId('sim-rotation-card');
  await expect(card).toBeVisible();
  const url = page.url();

  // The trigger opens the drawer in place -- no navigation.
  await card.getByTestId('sim-rotation-card-link').click();
  const panel = card.getByTestId('sim-rotation-card-drawer-panel');
  await expect(panel).toBeVisible();
  expect(page.url()).toBe(url);

  // The fidelity link moved inside the drawer, and opens a new tab rather than this one.
  const fidelityLink = panel.getByTestId('sim-rotation-drawer-fidelity-link');
  await expect(fidelityLink).toHaveAttribute('href', '/sim/specs#warrior-fury');
  await expect(fidelityLink).toHaveAttribute('target', '_blank');

  // The fixture's warrior-fury row is not validated, so the note is there; a validated
  // spec renders the card without one.
  await expect(card.getByTestId('sim-rotation-card-note')).toBeVisible();
});

// Fix round 1, Finding 3: the card's defining behaviour -- a validated spec renders no
// note at all, rather than a green "all is well" line -- had no coverage. Dropping the
// `!== 'validated'` clause in needsFidelityNote would break nothing without this.
test('a saved sim for a validated spec carries the rotation card with no fidelity note', async ({ page }) => {
  await page.route('**/v1/specs', (route) =>
    route.fulfill(
      envelope({
        specs: [
          {
            spec: 'warrior-fury',
            state: 'validated',
            median_gap: 0.02,
            parses: 50,
            worst_actions: [],
            engine_version: activeBuild.build,
            updated_at: '2026-01-01',
          },
        ],
      }),
    ),
  );

  await page.goto('/sim/simfixtureab');
  const card = page.getByTestId('sim-rotation-card');
  await expect(card).toBeVisible();
  await card.getByTestId('sim-rotation-card-link').click();
  await expect(
    card.getByTestId('sim-rotation-card-drawer-panel').getByTestId('sim-rotation-drawer-fidelity-link'),
  ).toHaveAttribute('href', '/sim/specs#warrior-fury');
  await expect(card.getByTestId('sim-rotation-card-note')).toHaveCount(0);
});

test('a saved sim whose stored request has no gear shows the line, not the grid', async ({ page }) => {
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

  await expect(page.getByTestId('sim-no-gear')).toBeVisible();
  await expect(page.getByTestId('sim-slot-head')).toHaveCount(0);
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

  await expect(page.getByTestId('sim-saved-title')).toHaveText(named);
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

  test('a saved weights result renders its table and its Pawn string', async ({ page }) => {
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
    await expect(page.getByTestId('sim-pawn')).toContainText('( Pawn: v1:');
    // A weights result has no gear story (design 7): the strip's grid does not claim one.
    await expect(page.getByTestId('sim-slot-head')).toHaveCount(0);
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
  await expect(page.getByTestId('sim-saved-header')).toHaveCount(0);

  release?.();

  await expect(page.getByTestId('sim-saved-header')).toBeVisible();
  await expect(page.getByTestId('sim-saved-skeleton')).toHaveCount(0);
});

// Task 7: the saved-sim fetch is a named, retriggerable function (loadSavedSim) rather than
// a dead-end anonymous block, so a failed GET /v1/sims/{id} offers a real retry through
// LoadError -- the same request re-fired in place, not a page reload.
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

test('a stale engine version shows the pill and the sentence, and is never re-run automatically', async ({
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

  await expect(page.getByTestId('sim-stale-pill')).toHaveText('Engine 6a1c2d9');
  await expect(page.getByTestId('sim-stale')).toBeVisible();

  // No auto re-run: the remedy is the "Run this yourself" link, never a second fetch fired
  // on its own.
  await page.waitForTimeout(500);
  expect(calls).toBe(1);
});

test.describe('saving a browser run from /sim', () => {
  async function runFury(page: import('@playwright/test').Page): Promise<void> {
    await page.goto('/sim');
    await page.getByTestId('sim-addon-input').fill(FURY);
    await page.getByTestId('sim-addon-load').click();
    await expect(page.getByTestId('sim-character')).toBeVisible();
    await page.getByTestId('sim-run-button').click();
    await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 5000 });
  }

  test('opens pre-filled, saves, and shows the link without navigating', async ({ page }) => {
    await runFury(page);

    await expect(page.getByTestId('sim-save-open')).toBeEnabled();
    await page.getByTestId('sim-save-open').click();
    await expect(page.getByTestId('sim-save-title')).toHaveValue('Raid-buffed, 3:00, single target');

    await page.route('**/v1/sims', (route) => route.fulfill(envelope({ sim_id: 'simnew234567' }, 201)));
    await page.getByTestId('sim-save-confirm').click();

    await expect(page.getByTestId('sim-save-link')).toBeVisible();
    await expect(page.getByTestId('sim-save-link')).toHaveValue(/\/sim\/simnew234567$/);
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
