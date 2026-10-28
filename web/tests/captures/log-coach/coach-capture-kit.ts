// web/tests/captures/log-coach/coach-capture-kit.ts
// What the log-coach capture spec shares: the review folder, the widths, file naming, storage
// cleared before each load, and the same-scale board against a mock render when one exists.
import { existsSync, mkdirSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { MOCK_DIR, settle, sideBySide, sizeTo, targetIn } from '../support/capture-kit';

export const target = targetIn('2026-10-28-log-coach');
/** 390 phone, 1024 narrow desktop, 1440 the mock width, 1920 and 2000 the owner's wide screens. */
export const WIDTHS = [390, 1024, 1440, 1920, 2000] as const;
export const PHONE_AND_DESKTOP = [390, 1440] as const;
export const file = (row: string, state: string, width: number): string =>
  target(`log-coach-${row}-${state}-${width}.png`);

/** Clears storage before every load: the site caches /v1/me there, so a stale state would leak. */
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

export async function openAt(page: Page, url: string, width: number): Promise<void> {
  await sizeTo(page, width);
  await clearStorageOnLoad(page);
  await page.goto(url);
  await settle(page);
}

/** Leaves no control holding focus or hover, so a capture shows the state at rest. */
export async function atRest(page: Page): Promise<void> {
  await page.evaluate(() => (document.activeElement as HTMLElement | null)?.blur());
  await page.mouse.move(0, 0);
}

/** A full-page capture; with a mock board of the same subject, also the same-scale side-by-side. */
export async function shoot(
  page: Page,
  row: string,
  state: string,
  width: number,
  mock?: string,
): Promise<void> {
  mkdirSync(path.dirname(target('x')), { recursive: true });
  const built = file(row, state, width);
  await page.screenshot({ path: built, fullPage: true });
  const mockPng = mock === undefined ? null : path.join(MOCK_DIR, mock);
  if (mockPng !== null && existsSync(mockPng)) {
    const board = await page.context().newPage();
    await sideBySide(board, mockPng, built, file(row, `${state}-board`, width), width);
    await board.close();
  }
}

export function writeLog(name: string, data: unknown): void {
  mkdirSync(path.dirname(target('x')), { recursive: true });
  writeFileSync(target(name), `${JSON.stringify(data, null, 2)}\n`);
}
