import { expect, test } from '@playwright/test';
import { signedOut } from './support/selector';

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
    // On a phone the Discord link waits behind the header's Menu button.
    await page.getByTestId('menu-button').click();
    const discord = await page.getByRole('link', { name: 'Discord' }).boundingBox();
    expect(discord?.height ?? 0).toBeGreaterThanOrEqual(44);
  });
});

test('a content page carries the same character selector as every other page', async ({ page }) => {
  // Owner 2026-10-04: /guides used to render a static "Sign in" while every tool page showed
  // the signed-in chip. The layout now mounts the character selector everywhere; signed out,
  // it resolves to the Sign in slot, whose panel carries the sign-in route.
  await signedOut(page);
  await page.goto('/guides');
  await expect(page.locator('astro-island[component-url*="CharacterSelector"]')).toHaveCount(1);
  await page.getByRole('banner').getByRole('button', { name: 'Sign in or choose a character' }).click();
  await expect(page.getByTestId('selector-sign-in')).toHaveAttribute('href', /\/login/);
});

// Owner rule (wide-screen): one centred site column for the whole site. Base.astro wraps every
// page in `.page-column`, so at ~2000px every page's first content sits in the same centred
// column with equal left and right margins. The BiS page was the one page that ran edge to edge.
test.describe('one centred column at 2000px', () => {
  const VIEWPORT_WIDTH = 2000;
  const EQUAL_MARGIN_TOLERANCE_PX = 2;
  const MIN_SIDE_MARGIN_PX = 100;
  const PAGES = [
    { path: '/bis/hunter/marksmanship', content: '[data-testid="bis-faction-panel-alliance"]' },
    { path: '/bis', content: 'main' },
    { path: '/tiers/tank', content: '[data-testid="tier-page"] .tier-body' },
    { path: '/guides', content: '.guides-index-body' },
    { path: '/changelog', content: 'main' },
  ];

  test.use({ viewport: { width: VIEWPORT_WIDTH, height: 1000 } });

  for (const { path, content } of PAGES) {
    test(`${path} content sits in a centred column`, async ({ page }) => {
      await page.goto(path);
      const box = await page.locator(content).first().boundingBox();
      expect(box).not.toBeNull();
      const left = box!.x;
      const right = VIEWPORT_WIDTH - (box!.x + box!.width);
      expect(Math.abs(left - right)).toBeLessThanOrEqual(EQUAL_MARGIN_TOLERANCE_PX);
      expect(left).toBeGreaterThan(MIN_SIDE_MARGIN_PX);
    });
  }

  test('the BiS header band is full-bleed while its content stays in the column', async ({ page }) => {
    await page.goto('/bis/hunter/marksmanship');
    const band = await page.locator('.art-panel').first().boundingBox();
    expect(band!.x).toBeLessThanOrEqual(1);
    expect(band!.width).toBeGreaterThanOrEqual(VIEWPORT_WIDTH - 1);
    const inner = await page.locator('.art-panel .page-column').first().boundingBox();
    const innerRight = VIEWPORT_WIDTH - (inner!.x + inner!.width);
    expect(Math.abs(inner!.x - innerRight)).toBeLessThanOrEqual(EQUAL_MARGIN_TOLERANCE_PX);
  });
});
