import { expect, test } from '@playwright/test';

// The 44px hit-target floor is a mobile guarantee only -- Footer.astro's
// `min-h-11 md:min-h-0` (plus NEEDS_HIT_TARGET_PADDING's px-2) drops the padding above the
// `md` breakpoint, same as layout.spec.ts's phone-layout describe and nav.spec.ts's phone
// nav describe. Scoped to the mobile project so it measures the same thing those do.
test.use({ viewport: { width: 360, height: 800 } });
test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

test('the footer links to Premium, Terms, Privacy and Refunds, each clearing the hit-target floor', async ({
  page,
}) => {
  await page.goto('/about');
  const nav = page.getByRole('navigation', { name: 'Footer' });
  for (const [label, href] of [
    ['Premium', '/premium'],
    ['Terms', '/terms'],
    ['Privacy', '/privacy'],
    ['Refunds', '/refunds'],
  ] as const) {
    const link = nav.getByRole('link', { name: label });
    await expect(link).toHaveAttribute('href', href);
    const box = await link.boundingBox();
    expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
  }
});
