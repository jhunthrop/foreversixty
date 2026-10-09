// web/src/lib/bis/expertise-cap.ts
// The "Expertise to cap" line under a band's stat weights, beside the hit line (`hit-cap.ts`).
// One formatter shared by the BiS page's weight rail (`panel-view.ts`) and the guides' stat
// table, so the two never word it twice.
import { bisCopy } from './copy';
import type { BisExpertiseToCap } from './types';
import { formatCapPercent } from './hit-cap';

export interface ExpertiseCapLine {
  text: string;
  title: string;
}

/** `undefined` when the band publishes no `expertise_to_cap` -- the caller then renders nothing. */
export function expertiseCapLine(
  expertiseToCap: BisExpertiseToCap | null | undefined,
): ExpertiseCapLine | undefined {
  if (expertiseToCap === null || expertiseToCap === undefined) return undefined;
  const parry = expertiseToCap.parry === undefined ? undefined : formatCapPercent(expertiseToCap.parry);
  return {
    text: bisCopy.expertiseToCapLine(formatCapPercent(expertiseToCap.dodge), parry),
    title: bisCopy.expertiseToCapTitle,
  };
}
