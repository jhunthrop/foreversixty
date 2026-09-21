// web/tests/e2e/sim-talents.spec.ts
// Design 3.4: /sim/talents is /sim/gear's own island, TopGear.svelte, with `store.tool`
// set to 'talents' instead of 'gear'. There is no separate component to test here -- this
// file proves the mode difference: gear locked, no gear-shaped sections, the talent list is
// the only thing that can be picked.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { bulkCopy } from '../../src/lib/sim/copy';
import { poolQualityCopy } from '../../src/lib/sim/pool-quality-copy';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;
// A second, genuinely different point allocation for the same class, pasted through ADD A
// BUILD's own ImportBox -- no gear after the last colon is valid FS1 (import.test.ts's own
// `'FS1:1.60.1.69893:paladin:human:20/0/0:'`), and the inline planner locks gear to the
// loaded character regardless (Task 14's own rule), so none is needed here.
const ALT_BUILD = `FS1:${activeBuild.build}:warrior:orc:5530515/0/0:`;

async function loadTalents(page: Page): Promise<void> {
  await page.goto('/sim/talents');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

test('there is no slot grid, no item search and no named sets — only the talent list', async ({ page }) => {
  await loadTalents(page);
  await expect(page.getByTestId('sim-talent-candidates')).toBeVisible();
  await expect(page.getByTestId('sim-slot-grid')).toHaveCount(0);
  await expect(page.getByTestId('sim-item-search')).toHaveCount(0);
  await expect(page.getByTestId('sim-named-sets')).toHaveCount(0);
});

test('ticking the character’s own build makes the run button live and ranks it', async ({ page }) => {
  await loadTalents(page);
  await page.getByTestId('sim-loadout-current').check();
  await expect(page.getByTestId('sim-combo-count')).toHaveText(bulkCopy.combinations(1));
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 20_000 });
  await expect(page.getByTestId('sim-combo-row').first()).toContainText(bulkCopy.talentsOwn);
});

test('the page has its own heading and its own canonical', async ({ page }) => {
  await page.goto('/sim/talents');
  await expect(page.getByRole('heading', { level: 1 })).toHaveText(bulkCopy.talentsTitle);
  await expect(page.locator('link[rel="canonical"]')).toHaveAttribute('href', /\/sim\/talents$/);
});

// dps D31/E2: signed out (every e2e run is), the saved-builds panel used to read "Your
// saved builds could not be read; the rest of the page still works." -- worded like a bug
// report for the expected, permanent state of a visitor with no account -- and a build
// pasted through ADD A BUILD had no checkbox anywhere once added, so the ranked table could
// never carry two builds side by side without signing in. This proves both: the message
// reads as guidance, not an error, and a pasted build can be ticked and run beside the
// character's own with nothing saved and nobody signed in.
//
// `/v1/builds?mine=1` is routed to a real 401 explicitly (the same envelope
// `curl https://api.foreversixty.gg/v1/builds?mine=1` returns signed out) rather than left
// to actually reach the production API: this dev server's origin has no CORS grant there,
// so the browser reports a network failure, not the 401 itself, and `TalentCandidates.
// svelte`'s own signed-out branch (`error.status === 401`) would never see the status that
// causes it on production.
test('signed out, the copy is guidance and a pasted build can be ranked beside the current one', async ({
  page,
}) => {
  await page.route('**/v1/builds?mine=1*', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: false,
        data: null,
        error: { code: 'unauthorized', message: 'sign in on the site first' },
        request_id: 'r',
      }),
    }),
  );
  await loadTalents(page);

  await expect(page.getByTestId('sim-loadouts-unavailable')).toHaveText(
    poolQualityCopy.talentsSavedSignedOut,
  );
  await expect(page.getByTestId('sim-loadouts-unavailable')).not.toHaveText(bulkCopy.talentsSavedUnavailable);

  await page.getByTestId('sim-loadout-add').click();
  await expect(page.getByTestId('sim-inline-planner')).toBeVisible();
  await page.getByTestId('import-code').fill(ALT_BUILD);
  await page.getByTestId('import-submit').click();
  await expect(page.getByTestId('import-error')).toHaveCount(0);
  await page.getByTestId('sim-loadout-accept').click();

  // The pasted build now has its own visible, ticked checkbox -- not silently folded into
  // `picked` with nothing on screen to show for it.
  const pastedBuild = page.getByTestId('sim-loadout-Build 1');
  await expect(pastedBuild).toBeVisible();
  await expect(pastedBuild).toBeChecked();

  await page.getByTestId('sim-loadout-current').check();
  // The original bug (E2): RUN still reported "1 valid combination" no matter what was
  // pasted, because the pasted build was never offered a checkbox to tick in the first
  // place. Both are ticked now, so both are submitted.
  await expect(page.getByTestId('sim-combo-count')).toHaveText(bulkCopy.combinations(2));

  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 20_000 });
  await expect(page.getByTestId('sim-combo-row').first()).toContainText('Build 1');
});
