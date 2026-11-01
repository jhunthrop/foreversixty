// web/tests/e2e/sim-follows-selector.spec.ts
// The owner's 2026-10-31 report on /sim: the header selector moved to another character and
// the page kept the old one, and the Run block's "Run sim" did nothing. Both characters were
// stored addon exports from an older client build (1.60.1.70450) whose tree the current build
// lays out differently; read against the current tree the second export was "not a legal
// build", so neither the follow nor the pick could load it. Here the older build's warrior
// tree swaps the Fury tab's first and last talents, so the export's five points are Booming
// Voice on its own build and an unreachable Flurry on the current one.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { armoryPointer, keyOf, signInWith, storePointer, type FixtureCharacter } from './support/selector';

const ACTIVE_BUILD = (
  JSON.parse(
    readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
  ) as { build: string }
).build;
const OLD_BUILD = '1.60.1.70450';
const FIXTURE_WARRIOR_TREE = path.join(
  import.meta.dirname,
  '..',
  '..',
  'src',
  'fixtures',
  'planner',
  'talents',
  'warrior.json',
);

const OBNOXIOUS: FixtureCharacter = {
  name: 'Obnoxious Yell',
  class: 'Warrior',
  spec: 'Fury',
  level: 21,
  buildDaysAgo: 9,
};
const IRONHIDE: FixtureCharacter = {
  name: 'Ironhide',
  class: 'Warrior',
  spec: 'Fury',
  level: 24,
  buildDaysAgo: 9,
};

const CURRENT_EXPORT = `FS1:${ACTIVE_BUILD}:warrior:orc:0/5530515/0:head=12640,main_hand=11726|level=21`;
const OLD_EXPORT = `FS1:${OLD_BUILD}:warrior:orc:0/0000005/0:head=12640|level=24`;

function envelope(data: unknown) {
  return {
    status: 200,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
  };
}

function oldWarriorTree(): string {
  const file = JSON.parse(readFileSync(FIXTURE_WARRIOR_TREE, 'utf8')) as { trees: { talents: unknown[] }[] };
  const swap = (talents: unknown[]) => [talents[6], ...talents.slice(1, 6), talents[0]];
  const trees = file.trees.map((tree, index) =>
    index === 1 ? { ...tree, talents: swap(tree.talents) } : tree,
  );
  return JSON.stringify({ ...file, build: OLD_BUILD, trees });
}

async function stubAccount(page: Page, characters: FixtureCharacter[]): Promise<void> {
  await signInWith(page, characters);
  for (const [character, code] of [
    [OBNOXIOUS, CURRENT_EXPORT],
    [IRONHIDE, OLD_EXPORT],
  ] as const) {
    await page.route(`**/v1/characters/${keyOf(character)}/sim-input`, (route) =>
      route.fulfill(
        envelope({
          spec: 'warrior-fury',
          gear: code,
          talents: '',
          buffs: [],
          captured_at: new Date(Date.now() - 9 * 86_400_000).toISOString(),
          source: 'addon',
        }),
      ),
    );
  }
  await page.route('**/v1/sims?mine=1*', (route) =>
    route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  // The older build's files: its own warrior tree, and the current build's copy of the rest.
  await page.route(`**/data/${OLD_BUILD}/**`, async (route) => {
    const url = route.request().url();
    if (url.endsWith('/talents/warrior.json')) {
      await route.fulfill({ status: 200, contentType: 'application/json', body: oldWarriorTree() });
      return;
    }
    await route.fulfill({
      response: await route.fetch({ url: url.replace(`/data/${OLD_BUILD}/`, `/data/${ACTIVE_BUILD}/`) }),
    });
  });
}

const selectorRow = (page: Page, character: FixtureCharacter) =>
  page
    .getByTestId('selector-panel')
    .getByTestId(`selector-row-${keyOf(character)}`)
    .getByRole('button')
    .first();

test.describe('/sim and the header selector', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  // The older build's files are proxied; one still in flight when a test ends is not a failure.
  test.afterEach(async ({ page }) => {
    await page.unrouteAll({ behavior: 'ignoreErrors' });
  });

  test('choosing another character in the header loads it on the page', async ({ page }) => {
    await stubAccount(page, [OBNOXIOUS, IRONHIDE]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/sim');
    await expect(page.getByTestId('paperdoll-name')).toHaveText('Obnoxious Yell');

    await page.getByTestId('character-selector').click();
    await selectorRow(page, IRONHIDE).click();
    await expect(page.getByTestId('selector-name')).toHaveText('Ironhide');

    await expect(page.getByTestId('paperdoll-name')).toHaveText('Ironhide');
    await expect(page.getByTestId('sim-character')).toContainText('Level 24');
  });

  test("the Run block's Run sim loads the current character and runs it", async ({ page }) => {
    // The Run block stands while no character is loaded: no pointer, and this session's
    // default already settled (as after Forget), so the block offers the account's main.
    await stubAccount(page, [IRONHIDE, OBNOXIOUS]);
    await page.addInitScript(() => window.sessionStorage.setItem('fs.pointerHealed', '1'));
    await page.goto('/sim');
    const block = page.getByTestId('sim-run-block');
    await expect(block).toContainText('Ironhide');

    await block.getByTestId('sim-run-block-action').click();

    await expect(page.getByTestId('paperdoll-name')).toHaveText('Ironhide');
    await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
  });
});
