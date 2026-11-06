// web/src/lib/sim/result-sections.ts
// The resource section's title (sim spec 2026-11-06 §4.5): the player's power, from the class.
import { primaryPowerTrack } from './primary-power';
import { powerCopy } from './result-copy';
import type { SimResult } from './types';

/** The word for the player's power, from the class (primary-power.ts); null when the run has no track for it. */
export function resourceLabel(result: SimResult): string | null {
  const found = primaryPowerTrack(result);
  return found === null ? null : powerCopy[found.power];
}
