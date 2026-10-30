// web/src/lib/report/fights.ts
// One rule, kept out of the component so it can be tested on its own: which fight the
// island shows when the url names one the report does not have.
//
// parseReportState is deliberately total -- a pasted link renders something rather than
// throwing -- and it validates ?fight= as a run of digits, which is all it can do: the
// query string is read before report.json arrives, so the real list of fights is not
// known there. `?fight=0` therefore parses even though fight_index is 1-based and no
// report has a fight 0, and so does `?fight=999` and a number that falls in a gap.
import { formatDuration } from './format';
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

/** The longest of `fights`, the first of equals; undefined for none. */
function longestFight(fights: readonly FightEntry[]): FightEntry | undefined {
  let best: FightEntry | undefined;
  for (const fight of fights) if (best === undefined || fight.duration_ms > best.duration_ms) best = fight;
  return best;
}

/**
 * The fight the bundled sample opens on, the one a visitor should see first: its first boss
 * kill, else its longest boss attempt, else its longest pull. A real report keeps
 * defaultFightIndex's rule. Today's sample is one night of trash with no boss, so it opens on
 * its longest pull; a sample with a boss in it opens on the boss.
 */
export function sampleFightIndex(fights: readonly FightEntry[]): number {
  const bosses = fights.filter((fight) => fight.kind === 'encounter');
  const chosen = bosses.find((fight) => fight.kill) ?? longestFight(bosses) ?? longestFight(fights);
  return chosen?.index ?? defaultFightIndex(fights);
}

/** What a fight's row is called: a boss keeps its name; trash is "Trash · 3:26", its length, since every trash pull is named "Trash". */
export function fightTitle(fight: Pick<FightEntry, 'kind' | 'name' | 'duration_ms'>): string {
  return fight.kind === 'encounter' ? fight.name : `Trash · ${formatDuration(fight.duration_ms)}`;
}

export type FightRow =
  | { type: 'fight'; fight: FightEntry }
  | {
      type: 'trash-group';
      /** The first pull's index: the group's key and its test id. */
      key: number;
      zone: string;
      fights: FightEntry[];
      kills: number;
      duration_ms: number;
      deaths: number;
    };

/**
 * The strip's rows: consecutive trash pulls in the same zone fold into one row, so forty-two
 * pulls of "Trash" read as a few zones. Trash the engine gave no zone, a lone pull and every
 * boss stay their own row.
 */
export function fightRows(fights: readonly FightEntry[]): FightRow[] {
  const rows: FightRow[] = [];
  let run: FightEntry[] = [];
  const flush = (): void => {
    if (run.length === 1) rows.push({ type: 'fight', fight: run[0] });
    else if (run.length > 1) {
      rows.push({
        type: 'trash-group',
        key: run[0].index,
        zone: run[0].zone ?? '',
        fights: run,
        kills: run.reduce((sum, entry) => sum + entry.npc_kills, 0),
        duration_ms: run.reduce((sum, entry) => sum + entry.duration_ms, 0),
        deaths: run.reduce((sum, entry) => sum + entry.deaths, 0),
      });
    }
    run = [];
  };
  for (const fight of fights) {
    const zone = fight.zone ?? '';
    if (fight.kind !== 'encounter' && zone !== '') {
      if (run.length > 0 && (run[0].zone ?? '') !== zone) flush();
      run.push(fight);
      continue;
    }
    flush();
    rows.push({ type: 'fight', fight });
  }
  flush();
  return rows;
}
