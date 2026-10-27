// web/tests/captures/support/capture-kit.ts
// The helpers the acceptance-capture specs share: where PNGs go, the settle pause, a padded crop
// of one element, a width-parametrised test, and the same-scale side-by-side board. Not a test.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Locator, type Page } from '@playwright/test';

const HERE = path.dirname(fileURLToPath(import.meta.url));
export const VIEWPORT_HEIGHT = 1000;
export const MOCK_DIR = path.resolve(HERE, '../../../../design/mocks/renders');
const REVIEWS_DIR = path.resolve(HERE, '../../../../design/reviews');
const PAD = 24;
const SETTLE_MS = 400;
const DEFAULT_CROP_MAX_PX = 1200;

/** `targetIn('2026-10-27-launch-pricing')` -> a function that resolves a file name into that review folder. */
export function targetIn(folder: string): (file: string) => string {
  return (file) => path.join(REVIEWS_DIR, folder, file);
}

export const settle = (page: Page): Promise<void> => page.waitForTimeout(SETTLE_MS);
export const fresh = (page: Page): Promise<Page> => page.context().newPage();
export const two = (n: number): string => String(n).padStart(2, '0');

export function atWidths(
  title: string,
  widths: readonly number[],
  body: (page: Page, width: number) => Promise<void>,
): void {
  for (const width of widths) {
    test(`${title} @${width}`, ({ page }) => body(page, width));
  }
}

export async function sizeTo(page: Page, width: number, height = VIEWPORT_HEIGHT): Promise<void> {
  await page.setViewportSize({ width, height });
}

/** A padded crop of `element` from a full-page capture, to `file`. */
export async function cropOf(
  page: Page,
  element: Locator,
  file: string,
  maxHeight = DEFAULT_CROP_MAX_PX,
): Promise<void> {
  const rect = await element.evaluate((node) => {
    const box = node.getBoundingClientRect();
    return {
      x: box.x + window.scrollX,
      y: box.y + window.scrollY,
      width: box.width,
      height: box.height,
      pageWidth: document.documentElement.clientWidth,
    };
  });
  const x = Math.max(0, rect.x - PAD);
  const y = Math.max(0, rect.y - PAD);
  await page.screenshot({
    path: file,
    fullPage: true,
    clip: {
      x,
      y,
      width: Math.min(rect.width + PAD * 2, rect.pageWidth - x),
      height: Math.min(rect.height + PAD * 2, maxHeight),
    },
  });
}

/** Mock on the left, built on the right, both drawn at `width`. */
export async function sideBySide(
  page: Page,
  mockPng: string,
  builtPng: string,
  file: string,
  width: number,
): Promise<void> {
  const data = (p: string): string => `data:image/png;base64,${readFileSync(p).toString('base64')}`;
  await page.setViewportSize({ width: width * 2 + 48, height: VIEWPORT_HEIGHT });
  await page.setContent(
    `<body style="margin:0;background:#000;color:#ccc;font:14px sans-serif;display:flex;gap:48px;align-items:flex-start">
      <figure style="margin:0;width:${width}px"><figcaption>Mock</figcaption><img src="${data(mockPng)}" style="width:${width}px"></figure>
      <figure style="margin:0;width:${width}px"><figcaption>Built (fixture)</figcaption><img src="${data(builtPng)}" style="width:${width}px"></figure>
    </body>`,
  );
  await page.screenshot({ path: file, fullPage: true });
}
