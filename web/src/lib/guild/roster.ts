// web/src/lib/guild/roster.ts
// Pure helpers for how the guild home's roster list is ordered and summarized (UX review
// defect 3, 2026-09-28): unverified rows were buried wherever the API happened to return
// them, with no signal to an officer that anyone was waiting on them.
import type { GuildRosterRow } from './api';

/**
 * Unverified rows first, verified rows after -- stable within each group, so two rows that
 * share a `verified` value keep the API's own relative order. `Array.prototype.sort` is
 * stability-guaranteed (ES2019+), so a plain comparator on `verified` alone is enough; no
 * secondary index-based tiebreak is needed. Never mutates `rows`.
 */
export function orderRoster(rows: readonly GuildRosterRow[]): GuildRosterRow[] {
  return [...rows].sort((a, b) => Number(a.verified) - Number(b.verified));
}

/** How many roster rows are still waiting on an officer's approval. */
export function unverifiedRosterCount(rows: readonly GuildRosterRow[]): number {
  return rows.filter((row) => !row.verified).length;
}
