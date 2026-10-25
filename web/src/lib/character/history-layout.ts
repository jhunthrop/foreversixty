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
