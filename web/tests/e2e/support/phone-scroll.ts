// web/tests/e2e/support/phone-scroll.ts
// Chromium's mobile emulation widens the layout viewport to fit content that overflows it
// (the "shrink-to-fit" behaviour Base.astro's `width=device-width, initial-scale=1` meta
// tag triggers, carrying no `shrink-to-fit=no`) -- window.innerWidth and
// document.documentElement.scrollWidth grow together when that happens, so a check that
// compares scrollWidth against the *live* window.innerWidth reads ~0 overflow even while
// the page genuinely overflows the device's own width. report-phone.spec.ts's own header
// comment documents catching this first ("innerWidth read 373 on this 360 viewport while
// Compare's table was overflowing... moves its own goal posts"); sim-phone.spec.ts hit the
// same false negative on SimResults' tab strip during Task 19's own fix round. Both specs
// need the same one comparison -- scrollWidth against the *fixed* device width the spec is
// auditing, never the live viewport -- so it lives here once instead of twice.
import { expect, type Page } from '@playwright/test';

/**
 * Asserts the page renders no wider than `width`, the phone width the calling spec fixed --
 * either via `test.use({ viewport })` (report-phone.spec.ts's own literal `PHONE_WIDTH`) or
 * by reading `page.viewportSize()` off the mobile project (sim-phone.spec.ts, which sets no
 * viewport of its own). Never `window.innerWidth`: see this file's header note.
 */
export async function assertNoHorizontalScroll(
  page: Page,
  width: number,
  message = 'page scrolls horizontally',
): Promise<void> {
  const scrollWidth = await page.evaluate(() => document.documentElement.scrollWidth);
  expect(scrollWidth, message).toBeLessThanOrEqual(width);
}
