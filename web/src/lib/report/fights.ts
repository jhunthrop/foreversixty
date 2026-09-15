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
import { ALL_FIGHTS } from './url';

/**
 * `wanted` when the report has that fight, otherwise `fallback` (the default fight).
 * ALL_FIGHTS (the whole night) is a fight the report has whenever it has a boss pull.
 */
export function resolveFightIndex(fights: readonly FightEntry[], wanted: number, fallback: number): number {
  if (wanted === ALL_FIGHTS)
    return fights.some((fight) => fight.kind === 'encounter') ? ALL_FIGHTS : fallback;
  return fights.some((fight) => fight.index === wanted) ? wanted : fallback;
}

export interface PullNumber {
  /** 1-based, counted per boss in the order the pulls happened. */
  pull: number;
  /** How many pulls of that boss the report has. */
  of: number;
}

/**
 * Which pull of its boss each encounter fight is. Two rows both reading "General Kaal"
 * tell a raid leader nothing about which was the wipe they want; "pull 1 of 2" does.
 * Trash fights are not numbered.
 */
export function pullNumbers(fights: readonly FightEntry[]): Map<number, PullNumber> {
  const counts = new Map<string, number>();
  const pulls = new Map<number, number>();
  for (const fight of fights) {
    if (fight.kind !== 'encounter') continue;
    const pull = (counts.get(fight.name) ?? 0) + 1;
    counts.set(fight.name, pull);
    pulls.set(fight.index, pull);
  }
  const out = new Map<number, PullNumber>();
  for (const fight of fights) {
    const pull = pulls.get(fight.index);
    if (pull !== undefined) out.set(fight.index, { pull, of: counts.get(fight.name) ?? pull });
  }
  return out;
}

/**
 * The fight a report opens on when the url names none: the first boss pull. A night's
 * first fight is nearly always a trash pull, often a few seconds long with one player in
 * it, and landing there reads as a broken report rather than as a choice.
 */
export function defaultFightIndex(fights: readonly FightEntry[]): number {
  const encounter = fights.find((fight) => fight.kind === 'encounter');
  return encounter?.index ?? fights[0]?.index ?? 1;
}
