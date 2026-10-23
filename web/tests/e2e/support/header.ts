// web/tests/e2e/support/header.ts
import { expect, type Page } from '@playwright/test';

/**
 * Wait until the header's panels, lookup and "/" key are wired.
 *
 * The header loads that wiring after first paint (Header.astro's loader, so the module chain
 * is not ahead of the page's largest contentful paint) and stamps `data-header-wired` on the
 * banner when it is done, the way the tooltip host stamps `data-tooltip-host`. A spec that
 * drives a door's panel by touch, the "/" key or the compact search toggle waits for it, or
 * its first input lands on markup nothing listens to yet.
 */
export async function headerWired(page: Page): Promise<void> {
  await expect(page.locator('[data-site-header][data-header-wired]')).toHaveCount(1);
}
