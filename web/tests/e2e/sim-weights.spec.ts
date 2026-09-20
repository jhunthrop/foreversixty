// web/tests/e2e/sim-weights.spec.ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { bulkCopy } from '../../src/lib/sim/copy';

const ROOT = path.join(import.meta.dirname, '..', '..');
const activeBuild = JSON.parse(readFileSync(path.join(ROOT, 'src', 'data', 'active-build.json'), 'utf8')) as {
  build: string;
};

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

// Task 20's fixture, read as plain JSON rather than imported as an ES module -- the same
// reason sim-saved.spec.ts reads its own `.json` fixtures this way (Playwright's own Node
// loader needs an import attribute this repo's other e2e specs never carry for a bare
// import).
const weightsFixture: Record<string, unknown> = JSON.parse(
  readFileSync(path.join(ROOT, 'src', 'fixtures', 'sim', 'weights-result.json'), 'utf8'),
) as Record<string, unknown>;

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'req-test' }),
  };
}

// sim/[id].astro only prerenders the one fixture id; `astro preview` carries no Worker to
// rewrite any other `/sim/<id>` onto that shell (Cloudflare-only), so a saved page reached
// by a fresh id has to have its shell and its `GET /v1/sims/<id>` stubbed -- the identical
// pattern sim-saved.spec.ts's own SHELL_HTML uses.
const SHELL_HTML = `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><title>Saved sim</title>
<link rel="stylesheet" href="/sim-island.css"></head>
<body><main id="main">
<div id="sim" data-sim-mount></div>
<script type="module" src="/sim-island.js"></script>
</main></body></html>`;

async function loadWeights(page: Page): Promise<void> {
  await page.goto('/sim/weights');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

test('the page opens with the warning, above everything else', async ({ page }) => {
  await page.goto('/sim/weights');
  const warning = page.getByTestId('sim-weights-warning');
  await expect(warning).toBeVisible();
  await expect(warning).toContainText(bulkCopy.weightsWarning);
  await expect(warning.getByRole('link', { name: bulkCopy.weightsWarningLink })).toHaveAttribute(
    'href',
    '/sim/gear',
  );
});

test('the stat picker defaults to the spec’s reference stat, first and ticked', async ({ page }) => {
  await page.route('**/v1/specs', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        request_id: 'r',
        error: null,
        data: {
          specs: [
            {
              spec: 'warrior-fury',
              state: 'validated',
              median_gap: 0.03,
              parses: 50,
              worst_actions: [],
              engine_version: 'x',
              updated_at: null,
              // Contract 10.1 A7: proto.Stat enum names in snake case.
              reference_stat: 'attack_power',
            },
          ],
        },
      }),
    }),
  );
  await loadWeights(page);
  await expect(page.getByTestId('sim-weight-pick-attack_power')).toBeChecked();
  await expect(page.getByTestId('sim-weights-reference')).toHaveText(/Attack power/);
});

// D45 (MAJOR, dps-minmaxer review round 1): the picker used to offer the full retail stat
// list -- Expertise, spell haste, armor penetration, MP5, feral attack power -- to every
// spec, including ones (all of them, in 1.60) none of those apply to. Once the spec's own
// `weight_stats` names the stats the engine actually weighs, the picker offers only those,
// by construction, and says so.
test('the picker offers only the spec’s own weight_stats, and says so, when the engine sent one', async ({
  page,
}) => {
  await page.route('**/v1/specs', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        request_id: 'r',
        error: null,
        data: {
          specs: [
            {
              spec: 'warrior-fury',
              state: 'validated',
              median_gap: 0.03,
              parses: 50,
              worst_actions: [],
              engine_version: 'x',
              updated_at: null,
              reference_stat: 'attack_power',
              weight_stats: ['attack_power', 'strength', 'crit'],
            },
          ],
        },
      }),
    }),
  );
  await loadWeights(page);
  await expect(page.getByTestId('sim-weights-stats-note')).toBeVisible();
  await expect(page.getByTestId('sim-weight-pick-attack_power')).toBeVisible();
  await expect(page.getByTestId('sim-weight-pick-strength')).toBeVisible();
  await expect(page.getByTestId('sim-weight-pick-crit')).toBeVisible();
  for (const excluded of [
    'expertise',
    'spell_haste',
    'armor_penetration',
    'mp5',
    'feral_attack_power',
    'hit',
  ]) {
    await expect(page.getByTestId(`sim-weight-pick-${excluded}`)).toHaveCount(0);
  }
});

test('the picker falls back to the full list, with no curated-list claim, when weight_stats is absent', async ({
  page,
}) => {
  await page.route('**/v1/specs', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        request_id: 'r',
        error: null,
        data: {
          specs: [
            {
              spec: 'warrior-fury',
              state: 'validated',
              median_gap: 0.03,
              parses: 50,
              worst_actions: [],
              engine_version: 'x',
              updated_at: null,
              reference_stat: 'attack_power',
            },
          ],
        },
      }),
    }),
  );
  await loadWeights(page);
  await expect(page.getByTestId('sim-weights-stats-note')).toHaveCount(0);
  await expect(page.getByTestId('sim-weight-pick-expertise')).toBeVisible();
});

test('a run renders a weight per stat with an error bar and a Pawn string', async ({ page, context }) => {
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  await loadWeights(page);
  // Deselecting an optional stat drops it from the run; the reference stat cannot be
  // deselected at all -- its checkbox is disabled, so WeightsSpec.Reference (required)
  // can never go empty through this control.
  //
  // `sim-character` becomes visible as soon as `adopt()` sets `character` -- before the
  // store's own async `loadDataFor` finishes and seeds `stats` -- so the picker can still
  // be showing every box unchecked at that instant. Waiting for the default seed to land
  // (stamina ticked) avoids unchecking a box that was never checked yet.
  await expect(page.getByTestId('sim-weight-pick-stamina')).toBeChecked();
  await page.getByTestId('sim-weight-pick-stamina').uncheck();
  await expect(page.getByTestId('sim-weight-pick-stamina')).not.toBeChecked();
  await expect(page.getByTestId('sim-weight-pick-attack_power')).toBeDisabled();
  await page.getByTestId('sim-weight-pick-attack_power').click({ force: true });
  await expect(page.getByTestId('sim-weight-pick-attack_power')).toBeChecked();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-weights')).toBeVisible({ timeout: 25_000 });
  await expect(page.getByTestId('sim-weight-stamina')).toHaveCount(0);
  await expect(page.getByTestId('sim-weight-attack_power')).toContainText('1.00');
  // Contract 10.8: haste is split, hit and crit are not.
  await expect(page.getByTestId('sim-weight-melee_haste')).toBeVisible();
  await expect(page.getByTestId('sim-weight-crit')).toBeVisible();
  await expect(page.getByTestId('sim-weight-hit')).toBeVisible();
  for (const wrong of ['haste', 'melee_crit', 'melee_hit']) {
    await expect(page.getByTestId(`sim-weight-${wrong}`)).toHaveCount(0);
  }
  await expect(page.getByTestId('sim-pawn')).toContainText('( Pawn: v1:');
  await page.getByTestId('sim-pawn-copy').click();
  await expect(page.getByTestId('sim-pawn-copy')).toHaveText(bulkCopy.weightsCopied);
});

test('a weights run never shows a DPS figure on the progress line', async ({ page }) => {
  await loadWeights(page);
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-weights')).toBeVisible({ timeout: 25_000 });
  // Engine-lane rule 3: a weights run's progress ticks carry dps: 0 and must never render as
  // a DPS figure. This page has no bulk stage count to show for one wasm call either, so the
  // progress line should simply never have appeared during the run.
  await expect(page.getByTestId('sim-stage-progress')).toHaveCount(0);
});

test('there is no slot grid, no source picker and no named sets', async ({ page }) => {
  await loadWeights(page);
  await expect(page.getByTestId('sim-slot-grid')).toHaveCount(0);
  await expect(page.getByTestId('sim-source-picker')).toHaveCount(0);
  await expect(page.getByTestId('sim-named-sets')).toHaveCount(0);
});

test('the run bar shows no combination count, but does show a working precision control', async ({
  page,
}) => {
  // Important 2, final whole-branch review: the weights tool sends no `bulk` block, so the
  // count never resolves (it would read "Counting combinations…" forever) -- that control
  // stays hidden. dps-minmaxer review round 1, D45 (BLOCKER): the page used to have no
  // working precision control at all, so an error bar bigger than its own weight could
  // never be tightened. `buildRequest` now reads `precision` for a weights run too (the
  // same `PRECISION_ITERATIONS` map /sim's own plain run uses), so the control belongs
  // here now, with the three fixed counts (no "until ±0.5%": a weights run is one wasm
  // call with no adaptive loop to stop, bulk-run.ts's own `runWeightsRun`).
  await loadWeights(page);
  await expect(page.getByTestId('sim-run-bulk-bar')).toBeVisible();
  await expect(page.getByTestId('sim-combo-count')).toHaveCount(0);
  const precision = page.getByTestId('sim-precision');
  await expect(precision).toBeVisible();
  await expect(precision.locator('option')).toHaveCount(3);
  // Playwright never reports a closed <select>'s own <option>s as visible; their text
  // content is what proves this select carries /sim's own "N iterations" labels rather
  // than the bulk tools' bare "Fast"/"Normal"/"High".
  await expect(precision.locator('option').first()).toHaveText(/iterations/);
});

// Contract 10.6: saved sims of every kind are public at /sim/<id>. Top Gear and talent
// compare already had this control (ComboResults.svelte, Task 16); this is the follow-up
// that gives a finished weights run the identical save form (SaveSimForm.svelte, the
// component the two now share) and proves the saved page renders the table Task 20 built
// for a `kind: weights` result.
test('a finished run can be saved, and the saved page renders its table', async ({ page }) => {
  await loadWeights(page);
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-weights')).toBeVisible({ timeout: 25_000 });

  await page.getByTestId('sim-save-open').click();
  await expect(page.getByTestId('sim-save-title')).not.toHaveValue('');

  // Exactly 12 characters, `[a-z2-7]` only (contract: a sim_id is always exactly 12,
  // SIM_ID_PATTERN) -- sim-saved.spec.ts's own comment on the same trap: one character too
  // many, or a digit outside 2-7, silently misses the pattern and routes the page to the
  // plain, no-character /sim landing state instead of a saved sim at all.
  const id = 'simweightsav';
  await page.route('**/v1/sims', (route) =>
    route.fulfill({
      status: 201,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { sim_id: id }, error: null, request_id: 'r' }),
    }),
  );
  await page.getByTestId('sim-save-confirm').click();

  await expect(page.getByTestId('sim-save-link')).toBeVisible();
  await expect(page.getByTestId('sim-save-link')).toHaveValue(new RegExp(`/sim/${id}$`));

  // Opening the saved link: the shell and the fetch it triggers both have to be stubbed --
  // `astro preview` carries no Worker to rewrite a fresh, non-prerendered id onto sim.html
  // the way Cloudflare does in production (sim-saved.spec.ts's own comment on the same
  // constraint).
  await page.route(`**/sim/${id}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
  );
  await page.route(`**/v1/sims/${id}`, (route) => route.fulfill(envelope({ ...weightsFixture, sim_id: id })));

  await page.goto(`/sim/${id}`);

  await expect(page.getByTestId('sim-weights')).toBeVisible();
  await expect(page.getByTestId('sim-pawn')).toContainText('( Pawn: v1:');
});
