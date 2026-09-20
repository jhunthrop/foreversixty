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

async function loadDrops(page: Page, code = FURY): Promise<void> {
  await page.goto('/sim/drops');
  await page.getByTestId('sim-addon-input').fill(code);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-source-picker')).toBeVisible();
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

test('a drop pins into Top Gear, carrying its origin in the URL', async ({ page }) => {
  await loadDrops(page);
  await page.getByTestId('sim-upcoming').check();
  await page.getByTestId('sim-source-raid:molten-core:11502').check();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-drops-flat')).toBeVisible({ timeout: 25_000 });
  // Either of Ragnaros's two drops will do -- this test cares that pinning works, not
  // which item the flat list lists first.
  const pinButton = page.locator('[data-testid^="sim-drops-pin-"]').first();
  await expect(pinButton).toBeVisible();
  const itemId = (await pinButton.getAttribute('data-testid'))?.replace('sim-drops-pin-', '') ?? '';
  await pinButton.click();
  await expect(page).toHaveURL(new RegExp(`/sim/gear\\?.*pin=${itemId}.*pinOrigin=drop`));
  // No ?source=/?ref= rode along with the pin (the addon code was typed by hand, not
  // arrived at through a source-carrying link), so Top Gear opens on the switcher rather
  // than a character -- the pin itself is still held (ToolsView's own effect) for the
  // moment one loads.
  await expect(page.getByTestId('sim-tools-empty')).toBeVisible();
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
