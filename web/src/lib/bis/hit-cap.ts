// web/src/lib/bis/hit-cap.ts
// The "Hit to cap" line under a band's stat weights. One formatter shared by the BiS page's
// weight rail (`panel-view.ts`) and the guides' stat table, so the two never word it twice.
import { bisCopy } from './copy';
import type { BisHitToCap } from './types';

export interface HitCapLine {
  text: string;
  title: string;
}

/** Whole percents print bare ("6"), fractions to one place ("6.5"). */
function formatPercent(value: number): string {
  return String(Number(value.toFixed(1)));
}

/** `undefined` when the band publishes no `hit_to_cap` -- the caller then renders nothing. */
export function hitCapLine(hitToCap: BisHitToCap | null | undefined): HitCapLine | undefined {
  if (hitToCap === null || hitToCap === undefined) return undefined;
  const white = hitToCap.white === undefined ? undefined : formatPercent(hitToCap.white);
  return {
    text: bisCopy.hitToCapLine(formatPercent(hitToCap.specials), white),
    title: bisCopy.hitToCapTitle,
  };
}
