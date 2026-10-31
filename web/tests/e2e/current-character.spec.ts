// web/tests/e2e/current-character.spec.ts
// End-to-end coverage for the one-product spec's current-character hand-offs (Tasks 1-11):
// a character loaded on one tool is offered back on the next, every paste box points at
// the addon, a Droptimizer upgrade opens straight into the planner, the simulator is
// honest about the level it models, and sharing a build asks before it writes.
import { expect, test } from './support/test';
import { addonCopy } from '../../src/lib/addon/copy';
import { ACTIVE_BUILD } from './support/active-build';
import { shareBuild } from './support/planner';
import { nativeSelect } from './support/select';

// The fixture Fury Warrior every other /sim* spec in this suite already pastes (grepped
// across tests/e2e for `FS1:`) -- a known-good FS1 v2 code for the fixture data build,
// reused rather than inventing a new one that may not decode against it.
const FURY = `FS1:${ACTIVE_BUILD}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;
// The same build from an addon that wrote the character's name: the only kind of export that can
// become the current character (a code that names nobody never does).
const NAMED_FURY = `${FURY}|who=Simfury:Realm`;

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ data, error: null }),
});

test.describe('current character', () => {
  test('a named export pasted on /sim, then opening /planner bare, restores it', async ({ page }) => {
    await page.goto('/sim');
    await page.getByTestId('sim-cold-paste-input').fill(NAMED_FURY);
    await page.getByTestId('sim-cold-paste-run').click();
    await expect(page.getByTestId('sim-character')).toBeVisible();

    // Planner rebuild spec §3: the proof of a restore on this page is the planner itself --
    // the pointer bridge (`decidePlannerLoad` still restores a bare /planner from
    // the stored 'code'/'addon' pointer), so the proof here is the restored build itself:
    // the fixture's Fury Warrior string reconstructs to 24 points in Fury with nothing
    // dropped (sim-saved.spec.ts's own fixture result carries the identical final string,
    // "0-5530515-").
    await page.goto('/planner');
    await expect(nativeSelect(page, 'Class')).toHaveValue('warrior');
    await expect(nativeSelect(page, 'Race')).toHaveValue('orc');
    await expect(page.getByTestId('tree-points-164')).toHaveText('24');
  });

  test('a code that names nobody is looked at and never becomes the current character', async ({ page }) => {
    await page.goto(`/sim?code=${encodeURIComponent(FURY)}`);
    await expect(page.getByTestId('sim-character')).toBeVisible();
    await expect(page.getByTestId('paperdoll-name')).toHaveCount(0);
    expect(await page.evaluate(() => window.localStorage.getItem('fs.currentCharacter'))).toBeNull();
    await page.goto('/sim');
    await expect(page.getByTestId('sim-character')).toHaveCount(0);
  });

  test('Forget in the header clears a restored pointer for the next bare load', async ({ page, context }) => {
    await page.goto('/sim');
    await page.getByTestId('sim-cold-paste-input').fill(NAMED_FURY);
    await page.getByTestId('sim-cold-paste-run').click();
    // The island boots after the shell's first paint (IslandScript.astro), so the pointer
    // is written a beat after `load`: wait for the character to show before leaving.
    await expect(page.getByTestId('sim-character')).toBeVisible();

    // Chosen in this visit: the selector says nothing about a restore.
    await page.getByTestId('character-selector').click();
    await expect(page.getByTestId('selector-panel')).toBeVisible();
    await expect(page.getByTestId('selector-restored')).toHaveCount(0);

    // A later visit (a new tab: a new session, the same stored pointer) restores it.
    const later = await context.newPage();
    await later.goto('/sim');
    await expect(later.getByTestId('sim-character')).toBeVisible();
    await later.getByTestId('character-selector').click();
    await expect(later.getByTestId('selector-restored')).toContainText('Restored your last character');
    await later.getByTestId('selector-forget').click();

    await later.goto('/sim');
    await expect(later.getByTestId('sim-character')).toHaveCount(0);
    expect(await later.evaluate(() => window.localStorage.getItem('fs.currentCharacter'))).toBeNull();
  });

  test('every paste box links to /setup', async ({ page }) => {
    await page.goto('/sim');
    await expect(page.getByTestId('sim-cold-paste-setup')).toHaveAttribute('href', '/setup');

    await page.goto('/planner');
    await expect(page.getByTestId('import-get-addon')).toHaveAttribute('href', '/setup');
  });

  test('the merged scope caveat is on /sim', async ({ page }) => {
    await page.goto('/sim');
    await expect(page.getByTestId('sim-scope-note')).toHaveText(
      'Damage specs are simulated; healing and tanking specs are not simulated yet.',
    );
  });

  test("Droptimizer's every-upgrade rows offer Plan it, opening the planner with the item's code", async ({
    page,
  }) => {
    // One upgrade (a ring drop off Ragnaros), enough to exercise the flat "every upgrade"
    // list's own Plan it link without the larger multi-combo fixture
    // sim-drops.spec.ts's own STUB_RESULT builds for a different purpose.
    const SIM_ID = 'zzzzzzzzzzz9';
    const DROP_RESULT = {
      engine_version: 'test-engine',
      request: {
        engine_version: 'test-engine',
        spec: 'warrior-fury',
        iterations: 3000,
        random_seed: 0,
        character: {
          name: 'Fury',
          race: 'orc',
          class: 'warrior',
          level: 60,
          talents: '0-5530515-',
          gear: [{ slot: 'main_hand', item_id: 12784 }],
          buffs: [],
          consumes: [],
        },
      },
      lane: 'server',
      dps: { mean: 1080, stddev: 50, error: 5, min: 980, max: 1180 },
      iterations_run: 3000,
      duration_ms: 500,
      summary: {},
      equipped: { mean: 1000, stddev: 50, error: 5, min: 900, max: 1100 },
      stages: [{ iterations: 3000, combos: 1 }],
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
      ],
    };

    await page.route('**/v1/me', (route) =>
      route.fulfill(envelope({ user: { premium: true }, characters: [], guilds: [] })),
    );
    await page.route('**/v1/sims/run', (route) => route.fulfill(envelope({ sim_id: SIM_ID })));
    await page.route(`**/v1/sims/${SIM_ID}/progress`, (route) =>
      route.fulfill(envelope({ state: 'done', iterations_done: 3000 })),
    );
    await page.route(`**/v1/sims/${SIM_ID}`, (route) => route.fulfill(envelope(DROP_RESULT)));

    await page.goto('/sim/drops');
    await page.getByTestId('sim-addon-input').fill(FURY);
    await page.getByTestId('sim-addon-load').click();
    await expect(page.getByTestId('sim-source-picker')).toBeVisible();

    await page.getByTestId('sim-upcoming').check();
    await page.getByTestId('sim-source-raid:molten-core:11502').check();
    await expect(page.getByTestId('sim-server-run')).toBeVisible();
    await page.getByTestId('sim-server-run').click();

    const planIt = page.getByTestId('sim-drops-plan-it-finger1:19325');
    await expect(planIt).toBeVisible({ timeout: 10_000 });
    await planIt.click();
    await expect(page).toHaveURL(/\/planner\?code=/);
  });

  test('Share shows a confirm step before writing, and Cancel leaves nothing saved', async ({ page }) => {
    let saveCalls = 0;
    await page.route('**/v1/builds', async (route) => {
      saveCalls += 1;
      await route.fulfill({
        status: 201,
        contentType: 'application/json',
        body: JSON.stringify({
          ok: true,
          data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
          error: null,
          request_id: 'req-1',
        }),
      });
    });

    await page.goto('/planner');
    await page.getByTestId('talent-1001').click();

    await page.getByTestId('share-open').click();
    await expect(page.getByTestId('share-confirm')).toBeVisible();
    await expect(page.getByTestId('share-link')).toHaveCount(0);

    await page.getByTestId('share-confirm-cancel').click();
    await expect(page.getByTestId('share-confirm')).toHaveCount(0);
    await expect(page.getByTestId('share-link')).toHaveCount(0);
    expect(saveCalls).toBe(0);
  });

  test('the confirm’s no-write alternatives copy without saving', async ({ page, context, browserName }) => {
    test.skip(browserName !== 'chromium', 'clipboard permissions are Chromium-only here');
    await context.grantPermissions(['clipboard-read', 'clipboard-write']);

    let saveCalls = 0;
    await page.route('**/v1/builds', async (route) => {
      saveCalls += 1;
      await route.fulfill({
        status: 201,
        contentType: 'application/json',
        body: JSON.stringify({
          ok: true,
          data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
          error: null,
          request_id: 'req-1',
        }),
      });
    });

    await page.goto('/planner');
    await page.getByTestId('talent-1001').click();
    await page.getByTestId('share-open').click();
    await expect(page.getByTestId('share-confirm')).toBeVisible();

    await page.getByTestId('share-confirm-copy-code').click();
    await expect(page.getByTestId('share-confirm-copy-code')).toHaveText(addonCopy.copiedForGame);

    await page.getByTestId('share-confirm-copy-unsaved').click();
    await expect(page.getByTestId('share-confirm-copy-unsaved')).toHaveText('Copied');

    await expect(page.getByTestId('share-link')).toHaveCount(0);
    expect(saveCalls).toBe(0);
  });

  test('sharing a build still writes it once "Share anyway" is chosen', async ({ page }) => {
    await page.route('**/v1/builds', (route) =>
      route.fulfill({
        status: 201,
        contentType: 'application/json',
        body: JSON.stringify({
          ok: true,
          data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
          error: null,
          request_id: 'req-1',
        }),
      }),
    );

    await page.goto('/planner');
    await page.getByTestId('talent-1001').click();
    await shareBuild(page);
    await expect(page.getByTestId('share-link')).toHaveText('https://foreversixty.gg/b/k7x2qm4a');
  });
});
