// web/tests/captures/launch-pricing/plans-capture-kit.ts
// What the launch-pricing capture specs share: the review folder, the clock instants, the file
// naming, the /premium opener, and the overlaid diff of two same-size captures. Not a test.
import { readFileSync } from 'node:fs';
import type { Locator, Page } from '@playwright/test';
import { VIEWPORT_HEIGHT, cropOf, settle, sizeTo, targetIn, two } from '../support/capture-kit';

export const target = targetIn('2026-10-27-launch-pricing');
export const LAUNCH_DAY = '2026-10-27T18:00:00Z';
export const DECEMBER = '2026-12-08T18:00:00Z';

export const file = (row: number, state: string, width: number, kind = ''): string =>
  target(`launch-pricing-${two(row)}-${state}${kind}-${width}.png`);

/** A full-page capture and, with a subject, a crop of it. */
export async function shoot(
  page: Page,
  row: number,
  state: string,
  width: number,
  subject?: Locator,
): Promise<void> {
  await page.screenshot({ path: file(row, state, width), fullPage: true });
  if (subject !== undefined) await cropOf(page, subject, file(row, state, width, '-crop'));
}

/** Clears storage before each load: the site caches /v1/me and the Premium source hint there. */
export async function clearStorageOnLoad(page: Page): Promise<void> {
  await page.addInitScript(() => {
    try {
      localStorage.clear();
      sessionStorage.clear();
    } catch {
      // Storage blocked: nothing to clear.
    }
  });
}

/**
 * Opens /premium at `width` on the clock `time`. Storage is cleared before every load: the site
 * caches /v1/me and the Premium source hint there, so a stale state would leak into the next
 * (the Premium card swaps its source line from that hint).
 */
export async function openPlans(page: Page, width: number, time: string): Promise<void> {
  await sizeTo(page, width);
  await clearStorageOnLoad(page);
  await page.clock.install({ time });
  await page.goto('/premium');
  await settle(page);
}

const dataUri = (png: string): string => `data:image/png;base64,${readFileSync(png).toString('base64')}`;

/** Two captures of the same width stacked, the top one in `difference` blend: unchanged pixels go black. */
export async function overlayDiff(
  page: Page,
  baselinePng: string,
  changedPng: string,
  out: string,
  width: number,
): Promise<void> {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await page.setContent(
    `<body style="margin:0;background:#000">
      <div style="position:relative;width:${width}px">
        <img src="${dataUri(baselinePng)}" style="display:block;width:${width}px">
        <img src="${dataUri(changedPng)}" style="position:absolute;left:0;top:0;width:${width}px;mix-blend-mode:difference">
      </div>
    </body>`,
  );
  await page.waitForFunction(() => [...document.images].every((image) => image.complete));
  await page.screenshot({ path: out, fullPage: true });
}
