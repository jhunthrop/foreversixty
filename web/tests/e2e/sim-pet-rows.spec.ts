// web/tests/e2e/sim-pet-rows.spec.ts
// A warlock registers all four demons as pet units, and the engine used to report each as
// active for the whole fight. A saved Demonology result (a real engine run through
// sim/adapter, src/fixtures/sim/demonology-result.json) must list the player and the
// Succubus, the one demon out, and none of the three that never were.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from './support/test';
import { openSimSection } from './support/sim-page';

const FIXTURES = path.join(import.meta.dirname, '..', '..', 'src', 'fixtures', 'sim');
function fixture(name: string): Record<string, unknown> {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8')) as Record<string, unknown>;
}
// The pinned engine version lives in result.json (make engine-pin rewrites it there), so
// the saved result borrows it and never reads as a stale engine.
const engineVersion = fixture('result.json').engine_version;
const saved = fixture('demonology-result.json');
const demonology = { ...saved, engine_version: engineVersion };
const ID = String(saved.sim_id);
/** The three demons never out; whole words, so "Improved" never trips it. */
const ABSENT = /\b(Imp|Felhunter|Voidwalker)\b/;

// The same minimal shell sim-saved.spec.ts serves for an id nothing prerendered.
const SHELL_HTML = `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><title>Saved sim</title>
<link rel="stylesheet" href="/sim-island.css"></head>
<body><main id="main">
<div id="sim" data-sim-mount></div>
<script type="module" src="/sim-island.js"></script>
</main></body></html>`;

test('a Demonology result lists the Succubus and none of the demons never summoned', async ({ page }) => {
  await page.route(`**/sim/${ID}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
  );
  await page.route(`**/v1/sims/${ID}`, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: demonology, error: null, request_id: 'req-test' }),
    }),
  );

  await page.goto(`/sim/${ID}`);
  // The saved page's damage table (spec 2026-11-04-sim-results 3.6): the player's own rows, then
  // one "Pet · Succubus" group. withoutAbsentPets (lib/sim/pet-rows.ts) drops the three demons
  // never out before the table, the buffs and the timeline read the summary.
  const damage = page.getByTestId('sim-damage-card');
  await expect(damage).toBeVisible();
  const petHeadings = damage.getByTestId('ability-pet-heading');
  await expect(petHeadings).toHaveCount(1);
  await expect(petHeadings.first()).toContainText('Succubus');
  await expect(damage.getByTestId('ability-row-other:attack/1').first()).toBeVisible();
  await expect(damage).not.toContainText(ABSENT);

  await openSimSection(page, 'buffs');
  const buffs = page.getByTestId('sim-section-buffs');
  await expect(buffs).not.toContainText(ABSENT);
});
