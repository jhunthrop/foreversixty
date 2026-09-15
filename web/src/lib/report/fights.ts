// web/src/lib/report/fights.ts
// One rule, kept out of the component so it can be tested on its own: which fight the
// island shows when the url names one the report does not have.
//
// parseReportState is deliberately total -- a pasted link renders something rather than
// throwing -- and it validates ?fight= as a run of digits, which is all it can do: the
// query string is read before report.json arrives, so the real list of fights is not
// known there. `?fight=0` therefore parses even though fight_index is 1-based and no
// report has a fight 0, and so does `?fight=999` and a number that falls in a gap.
import type { FightEntry } from './types';

/** `wanted` when the report has that fight, otherwise `fallback` (the first fight). */
export function resolveFightIndex(fights: readonly FightEntry[], wanted: number, fallback: number): number {
  return fights.some((fight) => fight.index === wanted) ? wanted : fallback;
}
