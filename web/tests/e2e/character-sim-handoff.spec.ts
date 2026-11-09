// web/tests/e2e/character-sim-handoff.spec.ts
// The owner's 2026-11-09 report: on another player's character page, the sim link opened the
// simulator on the reader's own header character. The link now carries the character on screen
// (`?source=armory&ref=<key>&keep=1`): the sim loads that character's gear, talents, race, level
// and spec from its sim-input, and the reader's header character stays theirs.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from './support/test';
import { routeCharacter } from './support/character-page';
import { armoryPointer, OBNOXIOUS, signInWith, storePointer } from './support/selector';

const ACTIVE_BUILD = (
  JSON.parse(
    readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
  ) as { build: string }
).build;
const FELSNAP_KEY = 'us/normal/felsnap';
const FELSNAP_EXPORT = `FS1:${ACTIVE_BUILD}:hunter:orc:0/0/0:|level=34`;

function envelope(data: unknown) {
  return {
    status: 200,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
  };
}

test.describe("Sim from another player's character page", () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  test('opens the character on screen in the sim and keeps the header character', async ({ page }) => {
    await routeCharacter(page, {
      id: 'no-score',
      mock: '',
      path: '/character/us/normal/felsnap',
      character: 'character-34.json',
    });
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    const simInputs: string[] = [];
    await page.route('**/v1/characters/*/*/*/sim-input', (route) => {
      simInputs.push(new URL(route.request().url()).pathname);
      return route.fulfill(
        envelope({
          spec: 'hunter-beast-mastery',
          gear: FELSNAP_EXPORT,
          talents: '',
          buffs: [],
          captured_at: '2026-11-05T18:00:00Z',
          source: 'addon',
        }),
      );
    });
    await page.route('**/v1/sims?mine=1*', (route) =>
      route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 100 })),
    );

    await page.goto('/character/us/normal/felsnap');
    const sim = page.getByTestId('character-sim');
    await expect(sim).toHaveAttribute(
      'href',
      `/sim?source=armory&ref=${encodeURIComponent(FELSNAP_KEY)}&keep=1`,
    );
    await sim.click();

    await expect(page).toHaveURL(/\/sim\?source=armory/);
    await expect(page.getByTestId('page-title')).toHaveText('Felsnap');
    await expect(page.getByTestId('sim-head-meta')).toContainText('Level 34');
    expect(simInputs).toContain(`/v1/characters/${FELSNAP_KEY}/sim-input`);
    const pointer = await page.evaluate(() => window.localStorage.getItem('fs.currentCharacter'));
    expect(JSON.parse(pointer ?? '{}')).toMatchObject({ source: 'armory', ref: 'us/normal/obnoxious-yell' });
  });
});
