// Talents start at level 10: a character below it has no point to spend, so the header's tiles
// say "0 points left" and "0/0 spent" instead of the 51 a finished build holds. A level-21
// character has earned 12 of them.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';

import { meAddonFixture } from '../../src/fixtures/me-addon';

const ok = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

function me(level: number) {
  const own = meAddonFixture.characters[0];
  return {
    ...meAddonFixture,
    main_character_key: 'us/pvp/obnoxious-yell',
    characters: [
      {
        ...own,
        key: 'us/pvp/obnoxious-yell',
        region: 'us',
        ruleset: 'pvp',
        name: 'Obnoxious Yell',
        class: 'Warrior',
        spec: undefined,
        level,
        build: { ...own.build, talents: { trees: ['0', '0', '0'], points: [0, 0, 0] } },
      },
    ],
  };
}

async function openAs(page: Page, level: number) {
  await page.route('**/v1/me', (route) => route.fulfill(ok(me(level))));
  await page.route('**/v1/devices', (route) => route.fulfill(ok([])));
  await page.goto('/planner');
}

test.describe('the planner counts only the points a character has earned', () => {
  test.skip(({ isMobile }) => isMobile, 'The facts rail is asserted at desktop width');
  test.use({ viewport: { width: 1280, height: 900 } });

  test('a level 9 character has no points left', async ({ page }) => {
    await openAs(page, 9);
    await expect(page.getByTestId('planner-level')).toHaveText('9');
    await expect(page.getByTestId('planner-remaining')).toHaveText('0');
    await expect(page.getByTestId('planner-spent')).toHaveText('0/0');
  });

  test('a level 21 character has twelve', async ({ page }) => {
    await openAs(page, 21);
    await expect(page.getByTestId('planner-level')).toHaveText('21');
    await expect(page.getByTestId('planner-remaining')).toHaveText('12');
    await expect(page.getByTestId('planner-spent')).toHaveText('0/12');
  });
});
