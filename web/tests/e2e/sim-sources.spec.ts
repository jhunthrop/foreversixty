import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from './support/test';
import { expectSimGearSlot, gotoSim, simCharacter } from './support/sim-page';

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

// The landing's paste Plate (SimLanding.svelte, spec 2026-11-06-sim §4.2) is /sim's one paste
// surface whenever no character is loaded; "Other ways in" sits under it for a build or a logged
// fight. The head then carries the character: its name as the title (ruling S1), the meta line
// and where the gear came from. "Change source" is retired: the header selector changes the
// character (one way to switch).
test('pasting a valid addon export loads the character into the head and the gear', async ({ page }) => {
  await gotoSim(page);
  await expect(page.getByTestId('sim-cold-paste')).toBeVisible();
  await expect(page.getByTestId('sim-sources')).toBeVisible();

  await page.getByTestId('sim-cold-paste-input').fill(`${FURY}|who=Simfury:Realm`);
  await page.getByTestId('sim-cold-paste-run').click();

  await expect(simCharacter(page)).toBeVisible();
  await expect(page.getByTestId('page-title')).toHaveText('Simfury');
  await expect(page.getByTestId('sim-head-source')).toHaveText(/^Gear from the addon/);
  await expect(simCharacter(page)).toContainText('Fury Warrior');
  await expect(page.getByTestId('sim-cold-paste')).toHaveCount(0);
  // The paste runs at once, so the gear is the report's open Gear section (ruling S3).
  await expect(page.getByTestId('sim-live-report')).toBeVisible({ timeout: 15_000 });
  await page.getByTestId('sim-section-gear').scrollIntoViewIfNeeded();
  await expectSimGearSlot(page, 'head', 'Lionheart Helm');
});

// The Plate validates locally and synchronously (cold-paste.ts's `validateColdPaste`), so a
// nonsense paste never reaches `store.loadAddon` at all -- it shows the decoder's own refusal
// under the field, with Run disabled. The box classifies the paste as an addon export or a
// build link first (`parseBuildInput`); `decodeFS1`'s own FS2-specific sentence is reached
// through a `?code=` link, the path cold-paste.test.ts proves for a valid code.
test('a nonsense paste shows the decoder’s own sentence and Run stays disabled', async ({ page }) => {
  await gotoSim(page);

  const run = page.getByTestId('sim-cold-paste-run');
  await page
    .getByTestId('sim-cold-paste-input')
    .fill('https://foreversixty.gg/planner?code=FS2:1:warrior:orc:0/0/0:');

  await expect(run).toBeDisabled();
  await expect(page.locator('#sim-cold-paste-input-error')).toHaveText(
    'That code is FS2; this site reads FS1.',
  );
  await expect(simCharacter(page)).toHaveCount(0);
});

// Planner.svelte's "Sim this build" link, before the build is saved: `/sim?code=<FS1 code>`.
// The store adopts it at init, ahead of the component's first render.
test('opening /sim?code=<fs1> adopts the character on load, as a manual source', async ({ page }) => {
  await gotoSim(page, `/sim?code=${encodeURIComponent(FURY)}`);

  await expect(simCharacter(page)).toBeVisible();
  // A code names nobody, so the title is the spec (characterTitle), never a made-up name.
  await expect(page.getByTestId('page-title')).toHaveText('Fury Warrior');
  // Not "the addon": a code and an addon export share a decoder but not a source, and the
  // line must not claim the addon read gear it never saw.
  await expect(page.getByTestId('sim-head-source')).toHaveText(/^Gear from a pasted export/);
  await expect(page.getByTestId('sim-head-source')).not.toContainText('the addon');
  await expect(simCharacter(page)).toContainText('Fury Warrior');
  // Not run yet: the doll is the page's one Plate.
  await expectSimGearSlot(page, 'head', 'Lionheart Helm');
  await expect(page.getByTestId('sim-sources')).toHaveCount(0);
});

// "From a build" says it takes a planner link. An unsaved planner link carries the build in
// ?code=, and the box used to read its last path segment ("planner") as a saved build id.
test('an unsaved planner link pasted into From a build loads its character', async ({ page }) => {
  await gotoSim(page);
  await page
    .getByTestId('sim-build-input')
    .fill(`https://foreversixty.gg/planner?code=${encodeURIComponent(FURY)}`);
  await page.getByTestId('sim-build-load').click();

  await expect(simCharacter(page)).toContainText('Fury Warrior');
  await expectSimGearSlot(page, 'head', 'Lionheart Helm');
});
