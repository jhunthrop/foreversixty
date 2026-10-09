// web/src/lib/tiers/tier-rules.ts
// The tier list's two stated rules, in a module with no imports so the browser script and the
// copy can read them without pulling in the ranking code.

/** Measurement lines in the DPS list, in percent behind the top spec. They claim nothing about a spec. */
export const RULER_PERCENTS: readonly number[] = [10, 20, 30];

/** Two specs within this many percent of each other tie: the site's own gear adoption margin. */
export const TIE_MARGIN_PERCENT = 1;
