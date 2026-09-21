import { expect, test } from '@playwright/test';

// design/DESIGN-SYSTEM.md: "Second screen first ... 44px minimum hit targets",
// "Page gutter 48px desktop, 18px phone". 360x800 is the narrowest phone we design for.
test.describe('phone layout', () => {
  test.use({ viewport: { width: 360, height: 800 } });

  for (const path of ['/', '/planner', '/dungeons', '/dungeons/hall-of-thanes', '/changelog']) {
    test(`${path} fits the viewport without scrolling sideways`, async ({ page }) => {
      await page.goto(path);
      const { scrollWidth, clientWidth } = await page.evaluate(() => ({
        scrollWidth: document.documentElement.scrollWidth,
        clientWidth: document.documentElement.clientWidth,
      }));
      expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
    });
  }

  test('the Discord link and the search box clear 44px', async ({ page }) => {
    await page.goto('/');
    const discord = await page.getByRole('link', { name: 'Discord' }).boundingBox();
    expect(discord?.height ?? 0).toBeGreaterThanOrEqual(44);
    const searchbox = await page.getByRole('combobox', { name: 'Search the site' }).boundingBox();
    expect(searchbox?.height ?? 0).toBeGreaterThanOrEqual(44);
  });
});

test('Sign in is in the header with no session prop needed, on a page that never passed one', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
    }),
  );
  // /classes never passed `session` to Base.astro before this lane's change.
  await page.goto('/classes');
  await expect(page.getByTestId('session-nav').getByRole('link', { name: 'Sign in' })).toBeVisible();
});
