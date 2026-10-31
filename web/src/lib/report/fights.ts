// web/src/lib/report/fights.ts
// One rule, kept out of the component so it can be tested on its own: which fight the
// island shows when the url names one the report does not have.
//
// parseReportState is deliberately total -- a pasted link renders something rather than
// throwing -- and it validates ?fight= as a run of digits, which is all it can do: the
// query string is read before report.json arrives, so the real list of fights is not
// known there. `?fight=0` therefore parses even though fight_index is 1-based and no
// report has a fight 0, and so does `?fight=999` and a number that falls in a gap.
import { type KillCounts, sumKillCounts } from './kill-counts';
import { formatDuration } from './format';
import type { FightEntry } from './types';
import { ALL_FIGHTS, trashScope, trashScopeKey } from './url';

/**
 * `wanted` when the report has that fight, otherwise `fallback` (the default fight).
 * ALL_FIGHTS (the whole night) is a fight the report has whenever it has a boss pull, and a
 * trash scope whenever the report has that zone's trash group.
 */
export function resolveFightIndex(fights: readonly FightEntry[], wanted: number, fallback: number): number {
  if (wanted === ALL_FIGHTS)
    return fights.some((fight) => fight.kind === 'encounter') ? ALL_FIGHTS : fallback;
  if (trashScopeKey(wanted) !== null) return combinedFights(fights, wanted) === null ? fallback : wanted;
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

export interface TrashGroup {
  type: 'trash-group';
  /** The first pull's index: the group's key, its url (`trash-<key>`) and its test id. */
  key: number;
  zone: string;
  fights: FightEntry[];
  /** The zone's summed kill counts; null when any pull predates them (lib/report/kill-counts). */
  kills: KillCounts | null;
  duration_ms: number;
}

export type FightRow = { type: 'fight'; fight: FightEntry } | TrashGroup;

function trashGroup(zone: string, fights: FightEntry[]): TrashGroup {
  return {
    type: 'trash-group',
    key: fights[0].index,
    zone,
    fights,
    kills: sumKillCounts(fights),
    duration_ms: fights.reduce((sum, entry) => sum + entry.duration_ms, 0),
  };
}

/**
 * The fight list's rows, zone by zone in the order the night reached them: each boss pull its
 * own row, then that zone's trash folded into one entry, the way Warcraft Logs collapses it, so
 * forty-two pulls of "Trash" read as one line per zone. Trash the engine gave no zone keeps its
 * own rows, in the night's order.
 */
export function fightRows(fights: readonly FightEntry[]): FightRow[] {
  const byZone = new Map<string, FightEntry[]>();
  for (const fight of fights) {
    const zone = fight.zone ?? '';
    byZone.set(zone, [...(byZone.get(zone) ?? []), fight]);
  }
  return [...byZone].flatMap(([zone, zoneFights]): FightRow[] => {
    const own = (fight: FightEntry): FightRow => ({ type: 'fight', fight });
    if (zone === '') return zoneFights.map(own);
    const trash = zoneFights.filter((fight) => fight.kind !== 'encounter');
    const bosses = zoneFights.filter((fight) => fight.kind === 'encounter').map(own);
    return trash.length === 0 ? bosses : [...bosses, trashGroup(zone, trash)];
  });
}

/** Every zone's trash entry, as fightRows draws it. */
export function trashGroups(fights: readonly FightEntry[]): TrashGroup[] {
  return fightRows(fights).filter((row): row is TrashGroup => row.type === 'trash-group');
}

/** The selection that opens a trash entry: the combined view, or a lone pull's own view. */
export function trashGroupSelection(group: TrashGroup): number {
  return group.fights.length === 1 ? group.fights[0].index : trashScope(group.key);
}

/**
 * The pulls a combined view folds together: every boss pull for ALL_FIGHTS, one zone's trash for
 * its trash scope; null for a single fight (or a scope the report does not have).
 */
export function combinedFights(fights: readonly FightEntry[], selected: number): FightEntry[] | null {
  if (selected === ALL_FIGHTS) return fights.filter((fight) => fight.kind === 'encounter');
  const key = trashScopeKey(selected);
  if (key === null) return null;
  return trashGroups(fights).find((group) => group.key === key)?.fights ?? null;
}

/** "6 boss pulls", "1 boss pull", "No boss pulls": the fight list's count, which never counts trash. */
export function bossPullsLabel(count: number): string {
  if (count === 0) return 'No boss pulls';
  return `${count} ${count === 1 ? 'boss pull' : 'boss pulls'}`;
}

/** What a combined view is called: "All boss pulls", or "Trash · Zephras Isle". */
export function combinedLabel(fights: readonly FightEntry[], selected: number): string {
  if (selected === ALL_FIGHTS) return 'All boss pulls';
  const zone = combinedFights(fights, selected)?.[0]?.zone ?? '';
  return zone === '' ? 'Trash' : `Trash · ${zone}`;
}
