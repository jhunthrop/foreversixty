// web/src/lib/report/compare.ts
// Compare mode's arithmetic, kept out of the component so it can be tested without a DOM.
//
// Two pulls, or two players inside one pull, ability by ability. There is no engine call
// and no query here: every figure is already in the summaries the page has loaded, and the
// window has already been applied to both sides by window.ts's scopeSummary before these
// are asked. The one thing the summary cannot split is threat, which the engine keeps per
// player and never per ability, so a threat comparison expands to nothing and says so.
import { abilityKey, type Ability, type Summary } from './types';
import type { TimeWindow } from './window';

export type CompareMetric =
  'damage_done' | 'dps' | 'healing_done' | 'hps' | 'damage_taken' | 'dtps' | 'threat' | 'tps';

/** The metric picker's order, which is the order the options are listed in. */
export const COMPARE_METRICS: readonly CompareMetric[] = [
  'dps',
  'damage_done',
  'hps',
  'healing_done',
  'dtps',
  'threat',
  'tps',
  'damage_taken',
];

/** The caption's words for each metric: a key like dtps is not a sentence. */
export const METRIC_LABELS: Record<CompareMetric, string> = {
  damage_done: 'damage done',
  dps: 'DPS',
  healing_done: 'healing done',
  hps: 'HPS',
  damage_taken: 'damage taken',
  dtps: 'damage taken per second',
  threat: 'threat',
  tps: 'threat per second',
};

export const PER_SECOND_METRICS: ReadonlySet<CompareMetric> = new Set<CompareMetric>([
  'dps',
  'hps',
  'dtps',
  'tps',
]);

/** Which actor table a metric's ability split comes from; null for threat, which has none. */
export function metricTable(metric: CompareMetric): 'damage_done' | 'healing' | 'damage_taken' | null {
  switch (metric) {
    case 'damage_done':
    case 'dps':
      return 'damage_done';
    case 'healing_done':
    case 'hps':
      return 'healing';
    case 'damage_taken':
    case 'dtps':
      return 'damage_taken';
    default:
      return null;
  }
}

/** One ability on both sides of a comparison. */
export interface AbilityDiff {
  /** types.ts's abilityKey: the spell and the pet it came via. */
  key: string;
  name: string;
  via?: string;
  /** Null when that side never used the ability, which the table prints as a dash. */
  a: number | null;
  b: number | null;
  /** (a ?? 0) - (b ?? 0): positive means this side did more. */
  delta: number;
}

/** One actor's abilities out of the table a metric reads, keyed the way the tables key them. */
function abilityAmounts(summary: Summary | null, guid: string, metric: CompareMetric): Map<string, Ability> {
  // A plain Map: built once and returned to diffOf, never read reactively by key.
  const out = new Map<string, Ability>();
  const table = metricTable(metric);
  if (summary === null || table === null) return out;
  const actor = summary[table].find((row) => row.guid === guid);
  for (const ability of actor?.abilities ?? []) out.set(abilityKey(ability), ability);
  return out;
}

/**
 * The union of two ability tables as one signed list, largest difference first, so the
 * thing that changed most is the first line a reader sees. Ties break on the name, so two
 * runs over the same pull produce the same order.
 */
function diffOf(a: Map<string, Ability>, b: Map<string, Ability>): AbilityDiff[] {
  const out: AbilityDiff[] = [];
  for (const key of new Set([...a.keys(), ...b.keys()])) {
    const left = a.get(key);
    const right = b.get(key);
    const named = left ?? right;
    if (named === undefined) continue;
    const x = left === undefined ? null : left.effective;
    const y = right === undefined ? null : right.effective;
    out.push({ key, name: named.name, via: named.via, a: x, b: y, delta: (x ?? 0) - (y ?? 0) });
  }
  return out.sort((p, q) => Math.abs(q.delta) - Math.abs(p.delta) || p.name.localeCompare(q.name));
}

/** One player's abilities in two fights, largest difference first. */
export function abilityDiff(
  left: Summary | null,
  right: Summary | null,
  guid: string,
  metric: CompareMetric,
): AbilityDiff[] {
  return diffOf(abilityAmounts(left, guid, metric), abilityAmounts(right, guid, metric));
}

/** Two players inside one fight, ability by ability, largest difference first. */
export function playerAbilityDiff(
  summary: Summary | null,
  leftGuid: string,
  rightGuid: string,
  metric: CompareMetric,
): AbilityDiff[] {
  return diffOf(abilityAmounts(summary, leftGuid, metric), abilityAmounts(summary, rightGuid, metric));
}

/**
 * The phase names both pulls reached, in the first pull's order. Aligning by phase only
 * makes sense for a phase both sides got to: a pull that wiped in Phase 2 has no Phase 3
 * to compare, and offering one would compare a real span against nothing.
 */
export function sharedPhases(left: Summary | null, right: Summary | null): string[] {
  const theirs = new Set((right?.phases ?? []).map((phase) => phase.name));
  return (left?.phases ?? []).map((phase) => phase.name).filter((name) => theirs.has(name));
}

/** One pull's own span for a named phase, or null when it has none of that name. */
export function phaseWindow(summary: Summary | null, name: string): TimeWindow | null {
  const found = (summary?.phases ?? []).find((phase) => phase.name === name);
  return found === undefined ? null : { startMs: found.start_ms, endMs: found.end_ms };
}
