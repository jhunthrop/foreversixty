import { expect, test } from './support/test';

import { meAddonFixture } from '../../src/fixtures/me-addon';
import { choose } from './support/select';

const ok = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

function me(race: string) {
  const own = meAddonFixture.characters[0];
  return {
    ...meAddonFixture,
    characters: [
      {
        ...own,
        key: 'us/hardcore/obnoxious-yell',
        name: 'Obnoxious Yell',
        class: 'Warrior',
        race,
        level: 21,
      },
    ],
  };
}

test.describe('The planner starts on the character’s race', () => {
  test.skip(({ isMobile }) => isMobile, 'The header is asserted at desktop width');

  test('a Gnome warrior opens as Gnome, and a manual pick wins', async ({ page }) => {
    await page.route('**/v1/me', (route) => route.fulfill(ok(me('Gnome'))));
    await page.route('**/v1/devices', (route) => route.fulfill(ok([])));
    await page.setViewportSize({ width: 1280, height: 900 });
    await page.goto('/planner');

    await expect(page.getByRole('heading', { level: 1 })).toHaveText('Warrior talents');
    await expect(page.getByTestId('page-head-meta')).toContainText('Gnome · Level 21');
    await expect(page.getByTestId('planner-race-select')).toHaveValue('gnome');

    await choose(page.getByTestId('planner-race-select'), 'orc');
    await expect(page.getByTestId('page-head-meta')).toContainText('Orc · Level 21');
  });
});
