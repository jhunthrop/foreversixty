import { expect, test } from './support/test';
import { signedOut } from './support/selector';

// design/DESIGN-SYSTEM.md: "Second screen first ... 44px minimum hit targets",
// "Page gutter 48px desktop, 18px phone". 360x800 is the narrowest phone we design for.
// Crawlers read the served HTML, not the hydrated page: the character header's reserved box must
// not leak text ("Placeholder") or a second h1 into it.
test.describe('served HTML of the pages that mount the character header', () => {
  for (const path of ['/', '/bis', '/rankings', '/logs', '/planner', '/bis/warrior/fury']) {
    test(`${path} has no reserved-box text and exactly one h1`, async ({ request }) => {
      const html = await (await request.get(path)).text();
      // Case-sensitive: the word survives lowercase in CSS ::placeholder and input attributes.
      expect(html).not.toContain('Placeholder');
      expect(html.match(/<h1[\s>]/g) ?? []).toHaveLength(1);
    });
  }
});

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
});

test('a content page carries the same character selector as every other page', async ({ page }) => {
  // Owner 2026-10-04: /guides used to render a static "Sign in" while every tool page showed
  // the signed-in chip. The layout now mounts the character selector everywhere; signed out,
  // it resolves to the Sign in slot, whose panel carries the sign-in route.
  await signedOut(page);
  await page.goto('/guides');
  await expect(page.locator('astro-island[component-url*="CharacterSelector"]')).toHaveCount(1);
  await page.getByRole('banner').getByRole('button', { name: 'Sign in or choose a character' }).click();
  await expect(page.getByTestId('selector-sign-in-bnet')).toHaveAttribute('href', /battlenet/);
});

// Owner rule (wide-screen): one centred site column for the whole site. Base.astro wraps every
// page in `.page-column`, so at ~2000px every page's first content sits in the same centred
// column with equal left and right margins. The BiS page was the one page that ran edge to edge.
test.describe('one centred column at 2000px', () => {
  const VIEWPORT_WIDTH = 2000;
  const EQUAL_MARGIN_TOLERANCE_PX = 2;
  const MIN_SIDE_MARGIN_PX = 100;
  const PAGES = [
    { path: '/bis/hunter/marksmanship', content: '[data-testid="bis-plate"]' },
    { path: '/bis', content: 'main' },
    // The tier list is a split page (list and spec panel side by side): its column is the page body.
    { path: '/tiers', content: '[data-testid="tier-body"]' },
    { path: '/tiers/tank', content: '[data-testid="tier-body"]' },
    { path: '/tiers/healer', content: '[data-testid="tier-body"]' },
    { path: '/guides', content: '[data-testid="guides-class-picker"]' },
    { path: '/changelog', content: 'main' },
  ];

  test.use({ viewport: { width: VIEWPORT_WIDTH, height: 1000 } });

  // The widest check (review focus 2): the tier list's split column stays centred at 2560 too.
  for (const path of ['/tiers', '/tiers/tank', '/tiers/healer']) {
    test(`${path} sits in a centred column at 2560`, async ({ page }) => {
      const width = 2560;
      await page.setViewportSize({ width, height: 1000 });
      await page.goto(path);
      const box = (await page.getByTestId('tier-body').boundingBox())!;
      const right = width - (box.x + box.width);
      expect(Math.abs(box.x - right)).toBeLessThanOrEqual(EQUAL_MARGIN_TOLERANCE_PX);
      expect(box.x).toBeGreaterThan(MIN_SIDE_MARGIN_PX);
    });
  }

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

  test('the BiS control bar is full-bleed while its content stays in the column', async ({ page }) => {
    await page.goto('/bis/hunter/marksmanship');
    const band = await page.getByTestId('bis-controls').boundingBox();
    expect(band!.x).toBeLessThanOrEqual(1);
    expect(band!.width).toBeGreaterThanOrEqual(VIEWPORT_WIDTH - 1);
    const inner = await page.getByTestId('bis-controls').locator('.page-column').first().boundingBox();
    const innerRight = VIEWPORT_WIDTH - (inner!.x + inner!.width);
    expect(Math.abs(inner!.x - innerRight)).toBeLessThanOrEqual(EQUAL_MARGIN_TOLERANCE_PX);
  });
});
