import { expect, test } from '@playwright/test';
import { createServer } from 'vite';
import { addonCopy } from '../../src/lib/addon/copy';
import type * as Fsb1Module from '../../src/lib/addon/fsb1';

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
});
