// web/src/lib/guild/loot-view.ts
// The loot gain's one wording (guild loot spec 2026-11-05 §3.2), shared by the loot council
// (loot-council-view.ts): the API ranks candidates and names the awardee; this only formats.
import type { GuildLootCandidate } from './api';

export const LOOT_LOAD_ERROR = 'Loot did not load.';
/** A tier-1 gain cell when the candidate wears nothing there: the item-level gap would only be the
 *  item's own level (guild round 3). */
export const FILLS_EMPTY_SLOT = 'Fills an empty slot';
export const ON_BIS_LIST = 'On their BiS list';

/** A no-break space: "+16 item levels" never breaks inside. */
const NBSP = '\u00a0';

/**
 * A candidate's gain. Tier 0 reads the sim's gain in the BiS file's own unit (`+12.4 DPS`, a
 * healer's `HPS`, a tank's `TPS`), or `On their BiS list` when no figure was measured. Tier 1 reads
 * the item-level gap (`+14 item levels`), or `Fills an empty slot` when the candidate wears nothing
 * there (the gap would only be the item's own level). `gain_dps` is null on every tier-1 row, so
 * nothing may call a number method on it without checking.
 */
export function candidateGainLabel(candidate: GuildLootCandidate): string {
  if (candidate.tier === 0) {
    if (candidate.gain_dps === null || candidate.gain_unit === null) return ON_BIS_LIST;
    return `+${candidate.gain_dps.toFixed(1)}${NBSP}${candidate.gain_unit.toUpperCase()}`;
  }
  if (candidate.fills_empty_slot === true) return FILLS_EMPTY_SLOT;
  const delta = candidate.ilvl_delta ?? 0;
  return `+${delta}${NBSP}item${NBSP}level${delta === 1 ? '' : 's'}`;
}
