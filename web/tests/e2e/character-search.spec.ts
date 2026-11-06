// web/tests/e2e/character-search.spec.ts
// Task L8.2 ruling 7: the character search box (on /characters since the home redesign, design/specs/2026-11-06-home.md) routes Name (optional -Realm) to
// /character/<region>/<ruleset>/<name>. The fixture build prerenders one character page,
// us/hardcore/elyra-duskvale, so the search lands on a real profile; its API read is the
// shared character fixture.
import { expect, test } from './support/test';
import { CHARACTER_PATH, stubCharacter } from './support/character-fixture';
import { nativeSelect } from './support/select';

test('typing a name and submitting lands on the character’s profile', async ({ page }) => {
  await stubCharacter(page);
  await page.goto('/characters');

  const search = page.getByTestId('home-character-search');
  await search.getByRole('combobox', { name: 'Character name' }).fill('Elyra Duskvale-Hardcore');
  await search.getByRole('combobox', { name: 'Character name' }).press('Enter');

  await expect(page).toHaveURL(new RegExp(`${CHARACTER_PATH}$`));
  await expect(page.getByTestId('character-head').getByTestId('page-title')).toHaveText('Elyra Duskvale');
});

test('"Look up another character" lands on the lookup page with the name filled in', async ({ page }) => {
  await page.goto('/characters?character=Elyra%20Duskvale#lookup');
  await expect(
    page.getByTestId('home-character-search').getByRole('combobox', { name: 'Character name' }),
  ).toHaveValue('Elyra Duskvale');
});

test('the lookup link’s region and ruleset fill the selects, and the name box has the cursor', async ({
  page,
}) => {
  await page.goto('/characters?character=Elyra%20Duskvale&region=eu&ruleset=hardcore#lookup');
  const search = page.getByTestId('home-character-search');
  await expect(search.getByRole('combobox', { name: 'Character name' })).toHaveValue('Elyra Duskvale');
  await expect(nativeSelect(search, 'Region')).toHaveValue('eu');
  await expect(nativeSelect(search, 'Ruleset')).toHaveValue('hardcore');
  await expect(search.getByRole('combobox', { name: 'Character name' })).toBeFocused();
});
