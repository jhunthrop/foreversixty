// web/src/lib/sim/best-per-spec.ts
// 2026-09-26 layout pass, Finding 4: a compact "spec, best DPS, when" summary above a
// signed-in character's sim history, derived client-side from the rows the history panel
// already loaded -- no extra fetch, and no server change. A pure function of the same
// SimListRow[] SimHistory.svelte renders, so the two never disagree about which run was
// best for a given spec.
import type { SimListRow } from './types';

export interface BestPerSpecRow {
  spec: string;
  dps: number;
  createdAt: string;
}

/**
 * The highest-DPS row for each distinct spec, sorted best DPS first -- the order a raider
 * comparing specs wants, not creation order. Ties keep whichever row was seen first.
 */
export function bestPerSpec(rows: readonly SimListRow[]): BestPerSpecRow[] {
  const bestBySpec = new Map<string, SimListRow>();
  for (const row of rows) {
    const current = bestBySpec.get(row.spec);
    if (current === undefined || row.dps > current.dps) bestBySpec.set(row.spec, row);
  }
  return [...bestBySpec.values()]
    .sort((a, b) => b.dps - a.dps)
    .map((row) => ({ spec: row.spec, dps: row.dps, createdAt: row.created_at }));
}
