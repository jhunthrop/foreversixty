// web/tests/e2e/character-render.spec.ts
// The character page head's crest (one crest language, design/DESIGN-SYSTEM.md tenet 7): the
// circular class crest is always drawn, and the Battle.net body render is cut (character spec
// 2026-11-06 C3: one paperdoll, no second figure of the character) even when the API sends one.
import { expect, test } from './support/test';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL } from './support/character-fixture';

test('the head draws the class crest and never the Battle.net render, even with render_url', async ({
  page,
}) => {
  const renderUrl = 'https://render.worldofwarcraft.com/main-raw.png';
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ...CHARACTER,
        data: { ...CHARACTER.data, character: { ...CHARACTER.data.character, render_url: renderUrl } },
      }),
    }),
  );
  await page.goto(CHARACTER_PATH);

  const head = page.getByTestId('character-head');
  await expect(head.locator('img[src*="/icons/hd/crests/"]')).toHaveAttribute(
    'src',
    /\/icons\/hd\/crests\/priest\.webp$/,
  );
  await expect(page.locator(`img[src="${renderUrl}"]`)).toHaveCount(0);
});
