// web/src/lib/character/history-layout.ts
// Small layout rules for the History section that are not chart geometry: how a gap's
// reason wraps inside its band, and the stored height the next visit reserves (spec 3.6, 3.11).
import { writeReserve } from '../layout-reserve';
import type { ChartSize } from './history-scales';

/** Read back by the pre-paint script in Base.astro. */
export const HISTORY_RESERVE_KEY = 'fs.history.reserve';

/** Two full lists stacked on a phone stand over 2000 px, so the section's sanity bound is higher. */
export const HISTORY_MAX_RESERVE_PX = 4000;

/** The size the model is built for before the plot has measured its box (the 1440 desktop chart). */
export const PLACEHOLDER_SIZE: ChartSize = { width: 588, height: 294, phone: false };

/** An 11 px Barlow glyph averages about this wide. */
const GLYPH_WIDTH_PX = 5.6;
const BAND_PADDING_PX = 8;
const MAX_LINES = 2;

/** Scrolls jump rather than glide for a visitor who asked for less motion. */
export function scrollBehavior(): ScrollBehavior {
  const reduced =
    typeof window !== 'undefined' &&
    typeof window.matchMedia === 'function' &&
    window.matchMedia('(prefers-reduced-motion: reduce)').matches;
  return reduced ? 'auto' : 'smooth';
}

/**
 * The height worth keeping for the next visit's reserve: the settled, collapsed section. The
 * skeleton is the reserve itself, an error is a short box, and an opened list is the owner's
 * choice, so none of them is stored (null).
 */
export function reserveToStore(status: string, listOpen: boolean, height: number): number | null {
  return status === 'ready' && !listOpen ? height : null;
}

/** The rows a committed selection marks (CharacterFightTable's `data-marked`). */
export const MARKED_ROW_SELECTOR = '[data-marked="true"]';

/**
 * Jumps the first marked row inside `host` to the middle of the window (its `scroll-mt-24`
 * margin holds). Instant on purpose: a smooth scroll can be cut short by a layout that moves
 * under it, and headless Chromium never finished one. The row is looked up now, not earlier, so
 * the list is measured as it stands; an unmounted host or a list with no mark does nothing.
 */
export function scrollRowIntoView(host: Element | null | undefined): void {
  if (host === null || host === undefined || !host.isConnected) return;
  host.querySelector(MARKED_ROW_SELECTOR)?.scrollIntoView({ block: 'center', behavior: 'auto' });
}

/** The two waits before a scroll can measure freshly rendered rows. */
export interface LayoutWait {
  /** Svelte's `tick`: the rows are in the DOM. */
  tick: () => Promise<void>;
  /** `requestAnimationFrame`: the rows have been laid out and have height. */
  frame: (run: () => void) => unknown;
}

/** Runs `run` once the DOM holds the new rows and they have been laid out. */
export async function afterLayout(run: () => void, wait: LayoutWait): Promise<void> {
  await wait.tick();
  wait.frame(run);
}

/**
 * Whether a change of the committed marks opens the list and scrolls: only a new set of marked
 * rows (a rebuilt chart hands over an equal one) whose first row sits beyond the collapsed eight.
 */
export function opensOnSelection(previous: string, signature: string, needsExpand: boolean): boolean {
  return signature !== previous && needsExpand;
}

export function writeHistoryReserve(pixels: number): void {
  writeReserve(HISTORY_RESERVE_KEY, pixels, HISTORY_MAX_RESERVE_PX);
}

/** A gap's reason as up to two lines that fit a band `width` px wide, split at the space nearest the middle. */
export function wrapReason(reason: string, width: number): string[] {
  const fits = reason.length * GLYPH_WIDTH_PX <= width - BAND_PADDING_PX;
  if (fits || !reason.includes(' ')) return [reason];
  const middle = reason.length / 2;
  const spaces = [...reason].flatMap((char, index) => (char === ' ' ? [index] : []));
  const splitAt = spaces.reduce((best, index) =>
    Math.abs(index - middle) < Math.abs(best - middle) ? index : best,
  );
  return [reason.slice(0, splitAt), reason.slice(splitAt + 1)].slice(0, MAX_LINES);
}
