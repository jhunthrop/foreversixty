// web/src/lib/tiers/tier-rules.ts
// The tier list's stated rules, in a module with no imports so the browser script and the
// copy can read them without pulling in the ranking code.

export type TierLetter = 'S' | 'A' | 'B' | 'C' | 'D';

export interface TierBand {
  letter: TierLetter;
  /** Percent behind the top spec at which this tier starts, inclusive. The next tier's `from` ends it. */
  from: number;
}

/** The DPS tiers, best first: S within 5% of the top, A 5 to 10%, B 10 to 20%, C 20 to 30%, D
 *  beyond. A gap exactly on a line belongs to the lower tier (5.0 is A, 30.0 is D). The colours
 *  are the item quality colours and live in the page's CSS, keyed by letter. */
export const TIER_BANDS: readonly TierBand[] = [
  { letter: 'S', from: 0 },
  { letter: 'A', from: 5 },
  { letter: 'B', from: 10 },
  { letter: 'C', from: 20 },
  { letter: 'D', from: 30 },
];

/** Two specs within this many percent of each other tie: the site's own gear adoption margin.
 *  A tie may straddle a tier line; the mark then says the letters differ by less than the sim can tell. */
export const TIE_MARGIN_PERCENT = 1;
