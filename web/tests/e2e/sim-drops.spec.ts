// web/tests/e2e/sim-drops.spec.ts
// The Droptimizer against the checked-in fake engine (PUBLIC_SIM_ENGINE defaults to
// 'fake'), so no Go toolchain and no wasm artifact is needed -- the identical setup
// tests/e2e/sim-gear.spec.ts uses.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { bulkCopy } from '../../src/lib/sim/copy';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

// Read as plain JSON, not imported from `../../src/lib/sim/phase` -- that module's own
// `import builtIn from '../../data/phases.json'` has no `with { type: 'json' }` attribute,
// which Vite accepts (and every browser-side caller goes through Vite) but Playwright's own
// Node-based test loader refuses outright, failing this whole spec file to load.
const phasesData = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'phases.json'), 'utf8'),
) as { name: string; start: string }[];
const raidsOneStart = phasesData.find((phase) => phase.name === 'raids-1')?.start ?? '';
const raidsOneDateLabel = new Intl.DateTimeFormat('en-GB', {
  day: 'numeric',
  month: 'long',
  year: 'numeric',
  timeZone: 'UTC',
}).format(new Date(raidsOneStart));

// The fixture warrior, wearing the Arcanite Reaper (also Ragnaros's own drop and
// Blacksmithing's crafted item -- contract 10.1 A6's provenance-merge path needs an item
// that already has a home before a pin or a boss tick gives it a second one).
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;
const FURY_BLACKSMITH = `${FURY}|professions=blacksmithing`;

/**
 * A bare `/sim/drops` now restores the stored current-character pointer (current-character
 * spec section 1: every /sim* page, including the tools island, bootstraps from it) --
 * calling this a second time in one test, after an earlier load already wrote that pointer,
 * lands on the loaded `CharacterStrip` rather than the paste box `SourceSwitcher` shows.
 * "Change source" is that strip's own way back to the paste box (`sim-change-source`), so
 * this reaches for it whenever the restore beat the fresh paste this call wants instead of
 * asserting the paste box is unconditionally the first thing on the page.
 */
async function loadDrops(page: Page, code = FURY): Promise<void> {
  await page.goto('/sim/drops');
  const addonInput = page.getByTestId('sim-addon-input');
  const changeSource = page.getByTestId('sim-change-source');
  await expect(addonInput.or(changeSource)).toBeVisible();
  if (await changeSource.isVisible()) await changeSource.click();
  await addonInput.fill(code);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-source-picker')).toBeVisible();
}

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ data, error: null }),
});

const STUB_SIM_ID = 'zzzzzzzzzzz2';

/**
 * A hand-built drops result with genuine, distinct deltas -- the checked-in fake engine
 * (browser lane) ties every combination with the equipped baseline exactly (same
 * `random_seed` and `iterations` for every request in a stage, `engine-fake.ts`'s own
 * header), so proving "Every upgrade" excludes a tie/downgrade and "Best here" names the
 * true best needs a result the fake engine did not compute. Dispatched through the premium
 * lane (`page.route`, real HTTP calls the engine never touches) so the production code
 * under test -- `DropResults.svelte`'s own `isUpgrade` filter, restored to `> 0` in fix
 * round 1 -- is exercised unmodified, and nothing about the fake engine's determinism
 * changes.
 *
 * Three combos, one boss (Ragnaros): +80 (the true best), +30 (a lesser but real upgrade,
 * proving "Best here" picks the leader rather than merely the only or first member), and an
 * exact tie at the already-equipped Arcanite Reaper (not an upgrade).
 */
// `request.character` (required on the real `SimRequest` shape, `types.ts`) is what
// `combos.ts`'s `planItHref` reads to build each row's "Plan it" link -- omitted here
// before that link existed, a hand-built server-run stub with no `character` crashed
// `DropResults.svelte` at render (`gearForCombo` reading `.gear` off `undefined`). Mirrors
// FURY's own gear (head=12640, main_hand=12784) plus finger1=19325, the item one of this
// stub's own combos names as an upgrade over it.
const STUB_CHARACTER = {
  name: 'Fury',
  race: 'orc',
  class: 'warrior',
  level: 60,
  talents: '0-5530515-',
  gear: [
    { slot: 'head', item_id: 12640 },
    { slot: 'main_hand', item_id: 12784 },
    { slot: 'finger1', item_id: 19325 },
  ],
  buffs: [],
  consumes: [],
};

const STUB_RESULT = {
  engine_version: 'test-engine',
  request: {
    engine_version: 'test-engine',
    spec: 'warrior-fury',
    iterations: 3000,
    random_seed: 0,
    character: STUB_CHARACTER,
  },
  lane: 'server',
  dps: { mean: 1000, stddev: 50, error: 5, min: 900, max: 1100 },
  iterations_run: 3000,
  duration_ms: 500,
  summary: {},
  equipped: { mean: 1000, stddev: 50, error: 5, min: 900, max: 1100 },
  stages: [{ iterations: 3000, combos: 3 }],
  combos: [
    {
      substitutions: [
        {
          kind: 'item',
          slot: 'finger1',
          item_id: 19325,
          name: 'Band of Accuria',
          origin: 'drop:raid:molten-core:11502',
          source_name: 'Ragnaros',
        },
      ],
      dps: { mean: 1080, stddev: 50, error: 5, min: 980, max: 1180 },
      delta: { mean: 80, stddev: 0, error: 8, min: 0, max: 0 },
      group: 0,
    },
    {
      substitutions: [
        {
          kind: 'item',
          slot: 'head',
          item_id: 16963,
          name: 'Helm of Wrath',
          origin: 'drop:raid:molten-core:11502',
          source_name: 'Ragnaros',
        },
      ],
      dps: { mean: 1030, stddev: 50, error: 5, min: 930, max: 1130 },
      delta: { mean: 30, stddev: 0, error: 6, min: 0, max: 0 },
      group: 1,
    },
    {
      substitutions: [
        {
          kind: 'item',
          slot: 'main_hand',
          item_id: 12784,
          name: 'Arcanite Reaper',
          origin: 'drop:raid:molten-core:11502',
          source_name: 'Ragnaros',
        },
      ],
      dps: { mean: 1000, stddev: 50, error: 5, min: 900, max: 1100 },
      delta: { mean: 0, stddev: 0, error: 5, min: 0, max: 0 },
      group: 2,
    },
  ],
};

/**
 * Signs the visitor in as premium and answers the premium lane's own three routes --
 * `POST /v1/sims/run`, `GET /v1/sims/<id>/progress`, `GET /v1/sims/<id>` -- so
 * `runOnServer()` lands on `STUB_RESULT` without a real API anywhere. `runServerJob` always
 * waits one full `serverPollMs` (2,000ms, unconfigured here) before its first poll, so
 * callers give the run bar a generous timeout rather than the browser lane's own.
 */
async function stubPremiumRun(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill(envelope({ user: { premium: true }, characters: [] })),
  );
  await page.route('**/v1/sims/run', (route) => route.fulfill(envelope({ sim_id: STUB_SIM_ID })));
  await page.route(`**/v1/sims/${STUB_SIM_ID}/progress`, (route) =>
    route.fulfill(envelope({ state: 'done', iterations_done: 3000 })),
  );
  await page.route(`**/v1/sims/${STUB_SIM_ID}`, (route) => route.fulfill(envelope(STUB_RESULT)));
}

test('the picker groups every source kind and leaves quests off', async ({ page }) => {
  await loadDrops(page);
  // Molten Core (raids-1) and Azuregos (opens: "later") are both gated ahead of today, so
  // "show unreleased content" is what makes every kind's group actually render -- without
  // it, Raids and World bosses would have nothing visible under them.
  await page.getByTestId('sim-upcoming').check();
  for (const label of [
    bulkCopy.sourcesRaids,
    bulkCopy.sourcesDungeons,
    bulkCopy.sourcesWorld,
    bulkCopy.sourcesCrafted,
    bulkCopy.sourcesRep,
    bulkCopy.sourcesPvp,
  ]) {
    await expect(page.getByTestId('sim-source-picker')).toContainText(label);
  }
  await expect(page.getByTestId('sim-kind-quest')).not.toBeChecked();
  await expect(page.getByTestId('sim-source-quest')).toHaveCount(0);
  await page.getByTestId('sim-kind-quest').check();
  await expect(page.getByTestId('sim-source-quest')).toBeVisible();
});

test('an unreleased source is hidden until show-upcoming, and then says it has no date', async ({ page }) => {
  await loadDrops(page);
  // The fixture's world boss carries `opens: "later"` (contract 10.4), which never opens
  // whatever the clock says -- so this is a date-independent assertion.
  await expect(page.getByTestId('sim-source-world:azuregos')).toHaveCount(0);
  await page.getByTestId('sim-upcoming').check();
  await expect(page.getByTestId('sim-source-world:azuregos')).toBeVisible();
  await expect(page.getByTestId('sim-source-picker')).toContainText(bulkCopy.opensLater);
});

test('a dated, not-yet-open source carries the date it opens', async ({ page }) => {
  await loadDrops(page);
  // Molten Core opens on "raids-1" (phases.json: 2026-12-09), a real date rather than the
  // "later" sentinel -- a genuinely different case from the world boss above.
  await expect(page.getByTestId('sim-source-raid:molten-core')).toHaveCount(0);
  await page.getByTestId('sim-upcoming').check();
  await expect(page.getByTestId('sim-source-raid:molten-core')).toBeVisible();
  // GET /v1/phases fails (the API is not running in e2e) so the page falls back to its own
  // build-time phases.json -- the same file read above -- and this asserts the exact date
  // that fallback produces, not a re-typed one.
  await expect(page.getByTestId('sim-source-picker')).toContainText(raidsOneDateLabel);
});

test('crafted sources split by profession once the character records one', async ({ page }) => {
  await loadDrops(page);
  await expect(page.getByTestId('sim-source-picker')).toContainText(bulkCopy.sourcesProfessionsUnknown);
  await loadDrops(page, FURY_BLACKSMITH);
  await expect(page.getByTestId('sim-source-picker')).toContainText(bulkCopy.sourcesMyProfessions);
});

test('picking a boss counts its drops, and running ranks them by source', async ({ page }) => {
  await loadDrops(page);
  await page.getByTestId('sim-upcoming').check();
  await page.getByTestId('sim-source-raid:molten-core:11502').check();
  await expect(page.getByTestId('sim-combo-count')).toHaveText(/\d+ valid combinations?/);
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-drops-by-boss')).toBeVisible({ timeout: 25_000 });
  // The boss is named from Substitution.SourceName, which the candidate carried in
  // (contract 10.1 A6) -- the page does not re-join the origin id to loot.json.
  await expect(page.getByTestId('sim-drops-by-boss')).toContainText('Ragnaros');
  await expect(page.getByTestId('sim-drops-flat')).toBeVisible();
});

test('the flat list only lists genuine upgrades, and the boss card names the true best', async ({ page }) => {
  await stubPremiumRun(page);
  await loadDrops(page);
  await page.getByTestId('sim-upcoming').check();
  await page.getByTestId('sim-source-raid:molten-core:11502').check();
  await expect(page.getByTestId('sim-server-run')).toBeVisible();
  await page.getByTestId('sim-server-run').click();
  await expect(page.getByTestId('sim-drops-flat')).toBeVisible({ timeout: 10_000 });

  // Three drops came back: +80, +30 and an exact tie (the drop that is already equipped).
  // Only the two genuine upgrades get a pin button -- a tie is not "an upgrade" (design
  // 6.3; the production predicate is `> 0`, restored in fix round 1, Finding 1).
  //
  // The id is `<slot>:<item>`, not the item alone: a candidate fitting more than one slot
  // carries `Candidate.Slot === ""` and can be tried in either of its slots, which is two
  // rows with one item id (final whole-branch review, Minor 7).
  await expect(page.getByTestId('sim-drops-pin-finger1:19325')).toBeVisible();
  await expect(page.getByTestId('sim-drops-pin-head:16963')).toBeVisible();
  await expect(page.getByTestId('sim-drops-pin-main_hand:12784')).toHaveCount(0);

  await expect(page.getByTestId('sim-drops-by-boss')).toContainText(bulkCopy.dropsUpgrades(2, 3));
  // "Best here" is the +80 drop specifically, not merely the first or only member of the
  // group (fix round 1, Minor: two real upgrades under one boss, not one).
  await expect(page.getByTestId('sim-drops-best')).toContainText('+80');
});

/**
 * Task 7 (spec 2026-09-25 §6): the after-sim sentence reads back a small localStorage
 * record Droptimizer.svelte writes on a genuine top-row upgrade, with no fetch of its own.
 * This is the one end-to-end proof that the whole chain -- the write here, the read on a
 * later, wholly separate /sim page load -- actually reaches the results card, not just the
 * two halves in isolation (last-upgrade.test.ts's round-trip, SimResults.test.ts's SSR
 * render).
 */
test('a real Droptimizer upgrade names the next action on a later plain sim', async ({ page }) => {
  await stubPremiumRun(page);
  await loadDrops(page);
  await page.getByTestId('sim-upcoming').check();
  await page.getByTestId('sim-source-raid:molten-core:11502').check();
  await expect(page.getByTestId('sim-server-run')).toBeVisible();
  await page.getByTestId('sim-server-run').click();
  await expect(page.getByTestId('sim-drops-flat')).toBeVisible({ timeout: 10_000 });

  // Wait on the write itself, not the DOM's own render timing, for proof the effect ran.
  await expect
    .poll(() => page.evaluate(() => window.localStorage.getItem('fs.lastDroptimizerUpgrade')))
    .toContain('Band of Accuria');

  // A wholly separate page load, per current-character.ts's own "nothing fetched anew"
  // rule: no ?code= carries the upgrade across, only the localStorage record -- the
  // current-character pointer `loadDrops` above already wrote restores the same FURY
  // character here on its own, the same restore-over-paste-box behaviour `loadDrops`
  // itself accounts for.
  await page.goto('/sim');
  const addonInput = page.getByTestId('sim-addon-input');
  const characterStrip = page.getByTestId('sim-character');
  await expect(addonInput.or(characterStrip)).toBeVisible();
  if (await addonInput.isVisible()) {
    await addonInput.fill(FURY);
    await page.getByTestId('sim-addon-load').click();
    await expect(characterStrip).toBeVisible();
  }
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-results')).toBeVisible({ timeout: 10_000 });

  await expect(page.getByTestId('sim-next-action')).toContainText('Upgrade: ');
});

test('a drop pins into Top Gear, carrying its origin in the URL', async ({ page }) => {
  await stubPremiumRun(page);
  await loadDrops(page);
  await page.getByTestId('sim-upcoming').check();
  await page.getByTestId('sim-source-raid:molten-core:11502').check();
  await expect(page.getByTestId('sim-server-run')).toBeVisible();
  await page.getByTestId('sim-server-run').click();
  await expect(page.getByTestId('sim-drops-pin-finger1:19325')).toBeVisible({ timeout: 10_000 });
  await page.getByTestId('sim-drops-pin-finger1:19325').click();
  await expect(page).toHaveURL(/\/sim\/gear\?.*pin=19325.*pinOrigin=drop.*pinName=Ragnaros/);
  // No ?source=/?ref= rode along with the pin (the addon code was typed by hand, not
  // arrived at through a source-carrying link) -- but `loadDrops` above already pasted FURY,
  // which wrote the current-character pointer (current-character spec section 1: every
  // successful load stamps it). A bare /sim/gear now restores that pointer rather than
  // opening empty, so the pin lands on the restored character, not the switcher (updated
  // from this test's pre-pointer assumption that nothing but the URL could carry a
  // character across the navigation).
  await expect(page.getByTestId('sim-character')).toBeVisible();
  const row = page.getByTestId('sim-candidate-finger1-19325');
  await expect(row).toBeVisible();
  await expect(row.getByRole('checkbox')).toBeChecked();
  await expect(row).toContainText(bulkCopy.pinned);
});

test('a pin arriving with a source-carrying link lands ticked, upgrading an existing row rather than duplicating it', async ({
  page,
}) => {
  // Simulates a "sim this build" link into /sim/drops whose pin then carries ?source=&ref=
  // forward to /sim/gear, so the pin lands on a character that is already loading rather
  // than an empty switcher -- fix round 1, Finding 2's own race: `ToolsView`'s pin effect
  // used to gate on `character !== null` alone, which is already true before `items` is
  // populated, so `addSearchItem` could silently find nothing and no-op. Gating on
  // `phase === 'idle'` too closes that window; this proves it end to end rather than only
  // in the unit test.
  const params = new URLSearchParams({
    source: 'addon',
    ref: FURY,
    pin: '12784',
    pinOrigin: 'drop:raid:molten-core:11502',
    pinName: 'Ragnaros',
  });
  await page.goto(`/sim/gear?${params.toString()}`);
  await expect(page.getByTestId('sim-character')).toBeVisible();

  // main_hand=12784 is already an equipped row from FURY's own gear -- the pin must
  // upgrade it in place (drop origin, "Ragnaros" as its source, ticked), not add a second
  // row for the same item (candidates.ts's `addRow`, the provenance-merge rule).
  const row = page.getByTestId('sim-candidate-main_hand-12784');
  await expect(row).toHaveCount(1);
  await expect(row.getByRole('checkbox')).toBeChecked();
  await expect(row).toContainText(bulkCopy.pinned);
});

test('a pin for an item this character has no file entry for tells the player, not just the store', async ({
  page,
}) => {
  // Fix round 2: fix round 1's own fix only set `store.detail`, which `BulkRunBar` renders
  // solely inside `{#if message !== null}` -- a stale or cross-class pin id still failed
  // completely silently to the player. This asserts what actually reaches the screen.
  const params = new URLSearchParams({ source: 'addon', ref: FURY, pin: '999999' });
  await page.goto(`/sim/gear?${params.toString()}`);
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-message')).toBeVisible();
  await expect(page.getByTestId('sim-message')).toContainText(bulkCopy.itemNotAdded(999_999));
});

test('the page never shows a probability', async ({ page }) => {
  await loadDrops(page);
  await expect(page.getByTestId('sim-drops-note')).toHaveCount(0);
  await page.getByTestId('sim-upcoming').check();
  await page.getByTestId('sim-source-raid:molten-core:11502').check();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-drops-note')).toBeVisible({ timeout: 25_000 });
  await expect(page.getByTestId('sim-drops-note')).toHaveText(bulkCopy.dropsNoChance);
});
