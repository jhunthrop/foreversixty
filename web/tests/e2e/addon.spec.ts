import { expect, test } from '@playwright/test';
import { createServer } from 'vite';
import { addonCopy } from '../../src/lib/addon/copy';
import type * as Fsb1Module from '../../src/lib/addon/fsb1';
import { ACTIVE_BUILD } from './support/active-build';
import { openImportBox } from './support/planner';

// Not a plain `import { decodeFSB1 } from '../../src/lib/addon/fsb1'`. fsb1.ts used to
// import PINNED_STATS from sim/stats.ts, which imports a generated `.json` file at module
// scope (as do several other modules in this codebase -- planner/reference.ts,
// sim/buffs.ts, sim/phase.ts). Every browser build and every Vitest run goes through Vite,
// which inlines that JSON import; Playwright's plain Node ESM loader does not, and Node 22
// refuses a raw `.json` import without a `with { type: 'json' }` attribute the source
// doesn't carry. Dropping the stat-vocabulary sort took that particular chain away -- at
// the time of writing fsb1.ts reaches no JSON at runtime -- but the Vite route is kept
// deliberately rather than reverted to a plain import: it loads the module exactly as the
// shipped build and Vitest do, so this spec cannot start disagreeing with them the next
// time anything in the import graph grows a JSON or an alias.
//
// Also not `page.evaluate(() => import('/src/...'))`: the browser suite runs against
// `astro preview`, a production build with no `/src/...` path served, so that 404s.
let vite: Awaited<ReturnType<typeof createServer>>;
let decodeFSB1: typeof Fsb1Module.decodeFSB1;

test.beforeAll(async () => {
  vite = await createServer({ server: { middlewareMode: true }, appType: 'custom', logLevel: 'silent' });
  const mod = (await vite.ssrLoadModule('/src/lib/addon/fsb1.ts')) as typeof Fsb1Module;
  decodeFSB1 = mod.decodeFSB1;
});

test.afterAll(async () => {
  await vite.close();
});

test.describe('the addon flows', () => {
  test('the share panel copies a code the FSB1 decoder reads', async ({ page, context }) => {
    await context.grantPermissions(['clipboard-read', 'clipboard-write']);
    // The e2e fixture (FOREVER_DATA=fixture) ships a two-tree warrior with talent ids
    // 1001-1004 and no paladin at all, so this exercises the planner's default class.
    await page.goto('/planner');
    // Rebuild spec §11 (owner ruling): the gear panel left the planner entirely -- this
    // code now carries talents alone, which `decodeFSB1`'s own `gear.length` of 0 proves.
    await page.getByTestId('talent-1001').click();

    await page.getByTestId('copy-addon-code').click();
    await expect(page.getByTestId('copy-addon-code')).toHaveText(addonCopy.copiedAddonCode);

    const code = await page.evaluate(() => navigator.clipboard.readText());
    expect(code.startsWith('FSB1:')).toBe(true);
    // Decoded by the same module the site ships, so the test proves the button and the
    // decoder agree rather than re-parsing the grammar here.
    const decoded = decodeFSB1(code);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.order.length).toBe(1);
    expect(decoded.build.gear.length).toBe(0);
  });

  test('pasting an export string shows the tree and the approximated-order note', async ({ page }) => {
    // The fixture (FOREVER_DATA=fixture) ships a two-tree warrior, talent ids 1001-1007
    // (Arms) and 2001-2007 (Fury), and no paladin at all -- so this is built from the
    // fixture's own Arms tree rather than the paladin string a generic brief would use.
    // Tree field "3502" is base-36 per-talent ranks in tab order: 1001 (Improved Heroic
    // Strike, tier 0) gets 3, 1002 (Deflection, tier 0) gets 5 -- eight points, enough to
    // open tier 1 -- and 1004 (Tactical Mastery, tier 1, needs Deflection rank 2) gets 2,
    // for 3 + 5 + 2 = 10 legally reachable points with nothing dropped. The Fury and third
    // fields are both "0" -- decodeFS1 still requires three slash-separated tree fields
    // even though this class has two trees; orderFromRanks ignores the field it has no
    // tree for.
    await page.goto('/planner');
    await openImportBox(page);
    await page.getByTestId('import-code').fill(`FS1:${ACTIVE_BUILD}:warrior:human:3502/0/0:head=12640`);
    await page.getByTestId('import-submit').click();

    await expect(page.getByTestId('planner-spent')).toHaveText('10/51');
    await expect(page.getByTestId('import-note').first()).toHaveText(addonCopy.importOrderApproximated);
    await expect(page.getByTestId('import-error')).toHaveCount(0);
  });

  test('a code from another format is refused by name', async ({ page }) => {
    await page.goto('/planner');
    await openImportBox(page);
    await page.getByTestId('import-code').fill('FS2:nope');
    await page.getByTestId('import-submit').click();
    // The literal, not `addonCopy.wrongPrefix('FS2', 'FS1')`: this refusal comes out of
    // `planner/fs1.ts`, which owns its own strings and has no copy file, and `addonCopy`'s
    // template is FSB1's. The two read identically today by coincidence, so asserting the
    // FSB1 constant here would keep passing if fs1's wording changed. `import.test.ts`
    // asserts the same message as a literal for the same reason.
    await expect(page.getByTestId('import-error')).toHaveText('That code is FS2; this site reads FS1.');
  });

  test('an export for another class switches the planner to that class rather than reconstructing', async ({
    page,
  }) => {
    // The fixture has no paladin data at all, which is exactly the point: decodeFS1 only
    // parses the string, so this well-formed paladin export decodes fine, and the class
    // check has to catch it before anything tries to reconstruct an order against the
    // warrior tree the planner actually has loaded. The standalone planner then switches
    // to paladin itself (the import completes once that class's talents load; here they
    // never do, since the fixture has none) instead of asking the player to switch.
    await page.goto('/planner');
    await openImportBox(page);
    await page.getByTestId('import-code').fill(`FS1:${ACTIVE_BUILD}:paladin:human:0/0/0:`);
    await page.getByTestId('import-submit').click();
    // The box (and its "switching" line) leaves with the warrior trees the moment the class
    // changes; what persists is the class itself, and here the fixture's own missing-data
    // state for it.
    await expect(page.getByLabel('Class')).toHaveValue('paladin');
    await expect(page.getByTestId('planner-load-error')).toBeVisible();
  });

  // Rebuild spec §11 (owner ruling): GearPanel/ItemPicker (and their score sort, and the
  // stat-weights disclosure) left the planner page entirely -- the simulator keeps its own
  // gear step, and the BiS page keeps the weights rail. No replacement test lives here:
  // neither surface is this spec's own scope.
});
