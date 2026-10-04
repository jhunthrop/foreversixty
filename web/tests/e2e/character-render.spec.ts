// web/tests/e2e/character-render.spec.ts
// Character.svelte's header: the circular class crest is always there (one crest language,
// design/DESIGN-SYSTEM.md tenet 7) and the Battle.net body render sits on the right at lg
// only when GET /v1/characters/... carries render_url. Reuses the one prerendered character
// page (support/character-fixture.ts's CHARACTER_PATH) every character-page e2e spec shares.
import { expect, test } from '@playwright/test';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL } from './support/character-fixture';

function withCharacter(overrides: Record<string, unknown>) {
  return {
    ...CHARACTER,
    data: { ...CHARACTER.data, character: { ...CHARACTER.data.character, ...overrides } },
  };
}

test('shows the render on the right at lg beside the crest when render_url is present', async ({
  page,
}, testInfo) => {
  test.skip(testInfo.project.name !== 'desktop', 'lg-only render');
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify(withCharacter({ render_url: 'https://render.worldofwarcraft.com/main-raw.png' })),
    }),
  );
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-render')).toHaveAttribute(
    'src',
    'https://render.worldofwarcraft.com/main-raw.png',
  );
  await expect(page.getByTestId('character-avatar')).toHaveAttribute('src', /\/icons\/hd\/crests\/.+\.webp$/);
});

test('shows the crest and no render when the character has no Battle.net media', async ({ page }) => {
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(CHARACTER) }),
  );
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-render')).toHaveCount(0);
  await expect(page.getByTestId('character-avatar')).toHaveAttribute('src', /\/icons\/hd\/crests\/.+\.webp$/);
});
