// web/src/lib/bis/expertise-cap.ts
// The "Expertise to cap" line under a band's stat weights, beside the hit line (`hit-cap.ts`).
// One formatter shared by the BiS page's weight rail (`panel-view.ts`) and the guides' stat
// table, so the two never word it twice.
import { bisCopy } from './copy';
import type { BisExpertiseToCap, BisStatWeight } from './types';
import { formatCapPercent } from './hit-cap';

export interface ExpertiseCapLine {
  text: string;
  title: string;
}

/**
 * Whether anything the band's character wears or talents into gives Expertise (`baseline`, from
 * `sim/cmd/leveling-bis/expertisecap.go`: gear and talents). No Forever item carries the stat,
 * so for most specs it is 0, and a cap line for a stat the player cannot raise reads as a
 * joke; Combat's Weapon Expertise and a bear's 8% are what make it worth a line.
 */
export function expertiseIsGained(expertiseToCap: BisExpertiseToCap | null | undefined): boolean {
  return expertiseToCap !== null && expertiseToCap !== undefined && expertiseToCap.baseline > 0;
}

/** `undefined` when the band publishes no `expertise_to_cap`, or its character gains no
 *  Expertise at all (`expertiseIsGained`) -- the caller then renders nothing. */
export function expertiseCapLine(
  expertiseToCap: BisExpertiseToCap | null | undefined,
): ExpertiseCapLine | undefined {
  if (expertiseToCap === null || expertiseToCap === undefined) return undefined;
  if (!expertiseIsGained(expertiseToCap)) return undefined;
  const parry = expertiseToCap.parry === undefined ? undefined : formatCapPercent(expertiseToCap.parry);
  return {
    text: bisCopy.expertiseToCapLine(formatCapPercent(expertiseToCap.dodge), parry),
    title: bisCopy.expertiseToCapTitle,
  };
}

/** A band's weights without Expertise when its character gains none (`expertiseIsGained`):
 *  the pills, the stat table and the weights panel then agree with the cap line, which is
 *  left out on the same rule. */
export function weightsForGainedStats(
  weights: readonly BisStatWeight[],
  expertiseToCap: BisExpertiseToCap | null | undefined,
): BisStatWeight[] {
  return expertiseIsGained(expertiseToCap) ? [...weights] : weights.filter((w) => w.stat !== 'expertise');
}
