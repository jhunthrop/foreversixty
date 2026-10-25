import type { Page } from '@playwright/test';

declare global {
  interface Window {
    __cls: number;
  }
}

interface LayoutShiftEntry extends PerformanceEntry {
  value: number;
  hadRecentInput: boolean;
}

/** Starts summing layout shifts before the page's own scripts run. Call before `page.goto`. */
export async function trackLayoutShift(page: Page): Promise<void> {
  await page.addInitScript(() => {
    window.__cls = 0;
    new PerformanceObserver((list) => {
      for (const entry of list.getEntries() as LayoutShiftEntry[]) {
        if (!entry.hadRecentInput) window.__cls += entry.value;
      }
    }).observe({ type: 'layout-shift', buffered: true });
  });
}

/** The cumulative layout shift so far, after letting pending shifts be reported. */
export async function readLayoutShift(page: Page): Promise<number> {
  await page.waitForTimeout(500);
  return page.evaluate(() => window.__cls);
}
