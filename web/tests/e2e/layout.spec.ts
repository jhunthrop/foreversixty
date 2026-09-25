import { expect, test } from '@playwright/test';

// design/DESIGN-SYSTEM.md: "Second screen first ... 44px minimum hit targets",
// "Page gutter 48px desktop, 18px phone". 360x800 is the narrowest phone we design for.
test.describe('phone layout', () => {
  test.use({ viewport: { width: 360, height: 800 } });

  for (const path of ['/', '/planner', '/guides', '/changelog']) {
    test(`${path} fits the viewport without scrolling sideways`, async ({ page }) => {
      await page.goto(path);
      const { scrollWidth, clientWidth } = await page.evaluate(() => ({
        scrollWidth: document.documentElement.scrollWidth,
        clientWidth: document.documentElement.clientWidth,
      }));
      expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
    });
  }

  test('the Discord link clears 44px', async ({ page }) => {
    await page.goto('/');
    const discord = await page.getByRole('link', { name: 'Discord' }).boundingBox();
    expect(discord?.height ?? 0).toBeGreaterThanOrEqual(44);
  });
});

test('Sign in is in the header of a content page that passes no session prop, as plain HTML', async ({
  page,
}) => {
  // /guides ships no client JavaScript, so its Sign in is a link in the markup rather than
  // a hydrated island; nav.spec.ts pins that no script is requested for it.
  await page.goto('/guides');
  await expect(page.getByRole('banner').getByRole('link', { name: 'Sign in' })).toHaveAttribute(
    'href',
    '/login',
  );
});
