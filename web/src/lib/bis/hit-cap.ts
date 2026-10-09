// web/src/lib/bis/hit-cap.ts
// The "Hit to cap" line under a band's stat weights. One formatter shared by the BiS page's
// weight rail (`panel-view.ts`) and the guides' stat table, so the two never word it twice.
import { bisCopy } from './copy';
import type { BisHitToCap, BisSpellHitToCap } from './types';

export interface HitCapLine {
  text: string;
  title: string;
}

/** Whole percents print bare ("6"), fractions to one place ("6.5"). */
export function formatCapPercent(value: number): string {
  return String(Number(value.toFixed(1)));
}

function spellHitCapLine(hitToCap: BisSpellHitToCap): HitCapLine {
  const school = hitToCap.school === undefined ? undefined : formatCapPercent(hitToCap.school.to_cap);
  return {
    text: bisCopy.spellHitToCapLine(formatCapPercent(hitToCap.spell), school),
    title: bisCopy.spellHitToCapTitle,
  };
}

/** `undefined` when the band publishes no `hit_to_cap` -- the caller then renders nothing. */
export function hitCapLine(hitToCap: BisHitToCap | null | undefined): HitCapLine | undefined {
  if (hitToCap === null || hitToCap === undefined) return undefined;
  if (hitToCap.kind === 'spell') return spellHitCapLine(hitToCap);
  const white = hitToCap.white === undefined ? undefined : formatCapPercent(hitToCap.white);
  return {
    text: bisCopy.hitToCapLine(formatCapPercent(hitToCap.specials), white),
    title: bisCopy.hitToCapTitle,
  };
}
