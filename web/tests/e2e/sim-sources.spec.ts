import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';

// Read as JSON rather than imported as an ES module: Playwright's own Node runtime (not
// Vite, which every in-app import goes through) needs an import attribute this repo's
// other e2e specs do not carry for a plain `.json` import, so this reads the file the way
// src/test-support/sim-api.ts's own Node-side fixture reads already do.
const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

// The addon path needs no API: FS1 decoding and the talent/item lookups it drives are all
// `/data/<build>/…` reads the preview server serves for real under FOREVER_DATA=fixture, so
// nothing here is stubbed. Unlike the unit tests (which stub fetch and never look at the
// build segment of the URL), this carries the site's own active build id rather than the
// fixture files' own recorded "1.15.9.69722": `characterFromFs1` stamps the resulting
// character's `tree_version` from this field verbatim (character.ts), and `adopt()`
// (store.svelte.ts) then fetches items and talents by that value -- so against a real
// server the code has to name the build id that server actually publishes under, or the
// gear grid 404s into "Empty" slots.
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

// sim-one-paste lane: ColdPasteHero.svelte is /sim's own one paste surface whenever no
// character is loaded -- SourceSwitcher's addon card (and the `sim-sources` section it
// sits in) still renders underneath, for "From a build"/"From a logged fight", but
// `showAddonCard` drops its own addon card so there is only ever one paste box on screen.
// "Change source" reopens the switcher with a character already loaded, which is no longer
// the hero's own cold-start moment (`showColdPasteHero` requires `store.character ===
// null`), so the full four-source switcher -- addon card included -- is what that reopen
// still shows; this test only needs it visible again, not the card itself.
test('pasting a valid addon export renders the strip, and Change source returns to the switcher', async ({
  page,
}) => {
  await page.goto('/sim');
  await expect(page.getByTestId('sim-cold-paste')).toBeVisible();
  await expect(page.getByTestId('sim-sources')).toBeVisible();

  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await page.getByTestId('sim-cold-paste-run').click();

  const strip = page.getByTestId('sim-character');
  await expect(strip).toBeVisible();
  await expect(page.getByTestId('sim-character-name')).toBeVisible();
  await expect(page.getByTestId('sim-source-pill')).toHaveText('Addon export, just now');
  await expect(page.getByTestId('sim-character-descriptor')).toContainText('Fury Warrior');
  await expect(page.getByTestId('sim-slot-head')).toContainText('Lionheart Helm');

  await page.getByTestId('sim-change-source').click();
  await expect(page.getByTestId('sim-sources')).toBeVisible();
  await expect(strip).toBeHidden();
});

// The hero validates locally and synchronously (cold-paste.ts's `validateColdPaste`, the
// same decode AddonPasteBox.svelte and the old addon card both already proved), so a
// nonsense paste never reaches `store.loadAddon` or `store.message` at all now -- it shows
// the decoder's own refusal as the hero's own error, with Run disabled, rather than the
// switcher's `sim-source-message`. Unlike the old addon card (which handed `store.loadAddon`
// the raw textarea value directly), the hero first classifies the paste as an addon export
// or a build link (`parseBuildInput`, since the one box takes both) -- a bare wrong-prefix
// string is caught at that classification step, with cold-paste.ts's own generic message
// (already proved by cold-paste.test.ts's own "rejects unrecognisable text" case and
// sim-cold-paste.spec.ts's own e2e coverage); `decodeFS1`'s own FS2-specific sentence is
// only reachable through the hero wrapped in a `?code=` link, the same path
// cold-paste.test.ts's own "wrapped in a planner link's ?code=" case proves for a valid
// code -- this is that same path for an invalid one, the same refusal sources.test.ts pins
// for fromAddonExport.
test('a nonsense paste shows the decoder’s own sentence and Run stays disabled', async ({ page }) => {
  await page.goto('/sim');

  const run = page.getByTestId('sim-cold-paste-run');
  await page
    .getByTestId('sim-cold-paste-input')
    .fill('https://foreversixty.gg/planner?code=FS2:1:warrior:orc:0/0/0:');

  await expect(run).toBeDisabled();
  await expect(page.getByTestId('sim-cold-paste-error')).toHaveText('That code is FS2; this site reads FS1.');
  await expect(page.getByTestId('sim-character')).toBeHidden();
});

// Planner.svelte's "Sim this build" link, before the build is saved: `/sim?code=<FS1 code>`.
// The store adopts it at init, ahead of the component's first render.
test('opening /sim?code=<fs1> adopts the character on load, as a manual source', async ({ page }) => {
  await page.goto(`/sim?code=${encodeURIComponent(FURY)}`);

  const strip = page.getByTestId('sim-character');
  await expect(strip).toBeVisible();
  await expect(page.getByTestId('sim-character-name')).toBeVisible();
  // Not "Addon export, …": a code and an addon export share a decoder but not a source, and
  // the pill must not claim the addon read gear it never saw.
  await expect(page.getByTestId('sim-source-pill')).toHaveText('Entered by hand');
  await expect(page.getByTestId('sim-character-descriptor')).toContainText('Fury Warrior');
  await expect(page.getByTestId('sim-slot-head')).toContainText('Lionheart Helm');
  await expect(page.getByTestId('sim-sources')).toBeHidden();
});

// "From a build" says it takes a planner link. An unsaved planner link carries the build in
// ?code=, and the box used to read its last path segment ("planner") as a saved build id.
test('an unsaved planner link pasted into From a build loads its character', async ({ page }) => {
  await page.goto('/sim');
  await page
    .getByTestId('sim-build-input')
    .fill(`https://foreversixty.gg/planner?code=${encodeURIComponent(FURY)}`);
  await page.getByTestId('sim-build-load').click();

  await expect(page.getByTestId('sim-character-descriptor')).toContainText('Fury Warrior');
  await expect(page.getByTestId('sim-slot-head')).toContainText('Lionheart Helm');
});
