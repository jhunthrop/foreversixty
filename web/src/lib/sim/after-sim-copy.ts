// web/src/lib/sim/after-sim-copy.ts
// The after-sim results card's first line (spec 2026-09-25 §6): the one best next action
// from data a run already produced. A new module, not another key on copy.ts (1,288 lines,
// already over this codebase's 800-line ceiling).
import type { LastUpgrade } from './last-upgrade';

export const afterSimCopy = {
  upgrade: (itemName: string, sourceName: string, gain: string): string =>
    `Upgrade: ${itemName} from ${sourceName}, ${gain} DPS.`,
  noUpgrade: 'Run Top Gear to find your next upgrade.',
  noUpgradeLink: 'Run Top Gear',
  noUpgradeHref: '/sim/gear',
} as const;

/** The results card's first line: the last Droptimizer upgrade on record, or the Top Gear
 *  pointer when there is none. Pure -- takes the already-read value, fetches nothing. */
export function afterSimSentence(topUpgrade: LastUpgrade | null): string {
  return topUpgrade === null
    ? afterSimCopy.noUpgrade
    : afterSimCopy.upgrade(topUpgrade.itemName, topUpgrade.sourceName, topUpgrade.gain);
}
