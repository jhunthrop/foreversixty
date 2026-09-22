// web/tests/e2e/character-render.spec.ts
// Character.svelte's header render/avatar (brief 2026-09-22 §B7): the render at lg when
// GET /v1/characters/... carries render_url, the 44px avatar otherwise when it carries
// avatar_url, and neither when the character has no Battle.net media at all. Reuses the
// one prerendered character page (support/character-fixture.ts's CHARACTER_PATH) that
// every character-page e2e spec shares, routing its own fetch to a variant body per case.
import { expect, test } from '@playwright/test';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL } from './support/character-fixture';

function withCharacter(overrides: Record<string, unknown>) {
  return {
    ...CHARACTER,
    data: { ...CHARACTER.data, character: { ...CHARACTER.data.character, ...overrides } },
  };
}

test('shows the render on the right at lg when render_url is present', async ({ page }, testInfo) => {
  test.skip(testInfo.project.name !== 'desktop', 'lg-only render');
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify(
        withCharacter({
          avatar_url: 'https://render.worldofwarcraft.com/avatar.jpg',
          render_url: 'https://render.worldofwarcraft.com/main-raw.png',
        }),
      ),
    }),
  );
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-render')).toHaveAttribute(
    'src',
    'https://render.worldofwarcraft.com/main-raw.png',
  );
  await expect(page.getByTestId('character-avatar')).toHaveCount(0);
});

test('shows the 44px avatar when there is no render_url', async ({ page }) => {
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify(withCharacter({ avatar_url: 'https://render.worldofwarcraft.com/avatar.jpg' })),
    }),
  );
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-avatar')).toHaveAttribute(
    'src',
    'https://render.worldofwarcraft.com/avatar.jpg',
  );
  await expect(page.getByTestId('character-render')).toHaveCount(0);
});

test('shows neither when the character has no Battle.net media', async ({ page }) => {
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(CHARACTER) }),
  );
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-render')).toHaveCount(0);
  await expect(page.getByTestId('character-avatar')).toHaveCount(0);
});
