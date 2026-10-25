// web/tests/e2e/support/layout-shift.ts
// Chromium's layout-shift observer for the specs that gate cumulative shift: install it before the
// page loads, then drain what it has seen. Member-caused shifts and load shifts use the same two
// calls; the observer ignores shifts that follow the member's own input.
import type { Page } from '@playwright/test';

interface LayoutShiftEntry extends PerformanceEntry {
  value: number;
  hadRecentInput: boolean;
}

declare global {
  interface Window {
    __shifts: number[];
  }
}

const SETTLE_AFTER_FRAMES_MS = 300;

/** From the next navigation on, collects every layout shift (buffered, so none is missed). */
export async function recordShifts(page: Page): Promise<void> {
  await page.addInitScript(() => {
    window.__shifts = [];
    new PerformanceObserver((list) => {
      for (const entry of list.getEntries() as LayoutShiftEntry[]) {
        if (!entry.hadRecentInput) window.__shifts.push(entry.value);
      }
    }).observe({ type: 'layout-shift', buffered: true });
  });
}

/** The shifts reported since the last call, once the page has had two frames and a beat to settle. */
export async function drainShifts(page: Page): Promise<number[]> {
  return page.evaluate(
    (settleMs) =>
      new Promise<number[]>((resolve) => {
        requestAnimationFrame(() =>
          requestAnimationFrame(() => {
            setTimeout(() => resolve(window.__shifts.splice(0)), settleMs);
          }),
        );
      }),
    SETTLE_AFTER_FRAMES_MS,
  );
}

export const sumShifts = (values: number[]): number => values.reduce((total, value) => total + value, 0);
