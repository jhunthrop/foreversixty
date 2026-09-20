import { expect, test } from '@playwright/test';
import { createServer } from 'vite';
import { addonCopy } from '../../src/lib/addon/copy';
import type * as Fsb1Module from '../../src/lib/addon/fsb1';
import { ACTIVE_BUILD } from './support/active-build';
import { openGear } from './support/planner';

// Not a plain `import { decodeFSB1 } from '../../src/lib/addon/fsb1'`: fsb1.ts imports
// PINNED_STATS from sim/stats.ts, which imports a generated `.json` file at module scope
// (as do several other modules in this codebase -- planner/reference.ts, sim/buffs.ts,
// sim/phase.ts). Every browser build and every Vitest run goes through Vite, which inlines
// that JSON import; Playwright's plain Node ESM loader does not, and Node 22 refuses a raw
// `.json` import without a `with { type: 'json' }` attribute the source doesn't carry --
// changing a shared, multiply-imported source file to add one is outside this task. So
// fsb1.ts is loaded the same way Vitest already loads it: through Vite's own module graph,
// in SSR mode, which resolves the JSON import exactly as the shipped module does.
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
    // Spend a point and equip something, so the code carries both halves. Talent 1001
    // (Improved Heroic Strike) is spendable from an empty build.
    await page.getByTestId('talent-1001').click();
    // Below md the gear panel is the last tab rather than part of the column, so the slot
    // has to be brought on screen before it can be clicked. The talent click comes first
    // because selecting the Gear tab hides the tree grid it needs.
    await openGear(page);
    await page.getByTestId('slot-head').click();
    // The picker's header carries its own "Close" button ahead of the item rows, so the
    // item rows are matched by their own testid rather than "first button in the panel".
    await page.getByTestId('item-picker').locator('[data-testid^="item-"]').first().click();

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
    expect(decoded.build.gear.length).toBe(1);
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
    await page.getByTestId('import-code').fill(`FS1:${ACTIVE_BUILD}:warrior:human:3502/0/0:head=12640`);
    await page.getByTestId('import-submit').click();

    await expect(page.getByTestId('planner-spent')).toHaveText('10/51');
    await expect(page.getByTestId('import-note').first()).toHaveText(addonCopy.importOrderApproximated);
    await expect(page.getByTestId('import-error')).toHaveCount(0);
  });

  test('a code from another format is refused by name', async ({ page }) => {
    await page.goto('/planner');
    await page.getByTestId('import-code').fill('FS2:nope');
    await page.getByTestId('import-submit').click();
    // The literal, not `addonCopy.wrongPrefix('FS2', 'FS1')`: this refusal comes out of
    // `planner/fs1.ts`, which owns its own strings and has no copy file, and `addonCopy`'s
    // template is FSB1's. The two read identically today by coincidence, so asserting the
    // FSB1 constant here would keep passing if fs1's wording changed. `import.test.ts`
    // asserts the same message as a literal for the same reason.
    await expect(page.getByTestId('import-error')).toHaveText('That code is FS2; this site reads FS1.');
  });

  test('an export for another class is refused by name rather than reconstructed', async ({ page }) => {
    // The fixture has no paladin data at all, which is exactly the point: decodeFS1 only
    // parses the string, so this well-formed paladin export decodes fine, and the class
    // check has to refuse it before anything tries to reconstruct an order against the
    // warrior tree the planner actually has loaded.
    await page.goto('/planner');
    await page.getByTestId('import-code').fill(`FS1:${ACTIVE_BUILD}:paladin:human:0/0/0:`);
    await page.getByTestId('import-submit').click();
    await expect(page.getByTestId('import-error')).toHaveText(
      addonCopy.importWrongClass('paladin', 'warrior'),
    );
  });

  test('the item picker sorts by score', async ({ page }) => {
    // No ?class=paladin: the fixture (FOREVER_DATA=fixture) has no paladin data at all, so
    // this runs against the planner's default class, the fixture's own two-tree warrior.
    // The head slot carries the fixture's only two head items (Lionheart Helm, Helm of
    // Wrath), which is enough to exercise the sort.
    await page.goto('/planner');
    await openGear(page);
    await page.getByTestId('slot-head').click();
    const picker = page.getByTestId('item-picker');
    await picker.getByTestId('sort-by-score').check();

    const scores = await picker.getByTestId('item-score').allTextContents();
    expect(scores.length).toBeGreaterThan(1);
    const numbers = scores.map(Number);
    expect(numbers).toEqual([...numbers].sort((a, b) => b - a));
  });

  test('the gear panel shows the spec’s weights with their sources', async ({ page }) => {
    await page.goto('/planner');
    await openGear(page);
    const weights = page.getByTestId('gear-weights');
    // `toContainText` reads textContent, which a closed <details> still has -- so the proof
    // that the disclosure actually opened has to come first, and from its `open` state.
    await expect(weights).not.toHaveAttribute('open');
    await weights.click();
    await expect(weights).toHaveAttribute('open', '');
    await expect(weights).toContainText(addonCopy.weightsAreOpinions);
    // The weights list reads its stats through `statLabel`, the same table the totals above
    // it use -- "Attack power", never the raw contract id `attack_power`.
    await expect(weights).toContainText('Attack power');
    await expect(weights).not.toContainText('attack_power');
    await expect(weights.getByRole('link').first()).toHaveAttribute('href', /^https:/);
  });
});
