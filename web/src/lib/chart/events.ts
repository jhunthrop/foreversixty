// web/src/lib/chart/events.ts
// The moments a chart marks (controls and charts spec §B5.6, §B8): their lines of text, the ±2 s
// event list a tooltip reads, and the marker lane's collision and collapse rules.
import { splitUnitName } from '../characters';
import { formatAmount } from '../report/format';
import type { Death, Taunt, TargetChange } from '../report/types';
import type { ChartEvent, ChartPhase } from './types';

/** A tooltip lists every event this close to the hovered second. */
export const EVENT_RADIUS_MS = 2000;
/** A marker this close to the one before it steps right by the same distance. */
export const LANE_STEP_PX = 22;
/** Three or more markers of a kind inside this span collapse into one counted marker. */
export const LANE_CLUSTER_PX = 44;
/** A collapsed death cluster's tooltip lists this many deaths, then counts the rest. */
export const CLUSTER_ROWS = 6;

const short = (name: string): string => splitUnitName(name).name;

/** A death, with its killing blow (spell, source, amount, overkill) when the summary has one. */
export function deathEvent(death: Death, href: string): ChartEvent {
  const who = `${short(death.name)} died`;
  const blow = death.killing_blow;
  const base = { key: `death-${death.guid}-${death.at_ms}`, kind: 'death' as const, atMs: death.at_ms, href };
  if (blow === undefined) return { ...base, text: who, shortText: who, clusterText: short(death.name) };
  const overkill = (blow.overkill ?? 0) > 0 ? ` (${formatAmount(blow.overkill ?? 0)} overkill)` : '';
  return {
    ...base,
    text: `${who} · ${blow.spell_name} from ${short(blow.source_name)} · ${formatAmount(blow.amount)}${overkill}`,
    shortText: `${who} · ${blow.spell_name} · ${formatAmount(blow.amount)}`,
    clusterText: `${short(death.name)} · ${blow.spell_name}`,
  };
}

export function tauntEvent(taunt: Taunt, iconUrl: string | null = null): ChartEvent {
  return {
    key: `taunt-${taunt.source_guid}-${taunt.at_ms}`,
    kind: 'taunt',
    atMs: taunt.at_ms,
    text: `Taunt · ${short(taunt.source_name)} on ${short(taunt.target_name)}`,
    iconUrl,
  };
}

export function targetChangeEvent(change: TargetChange): ChartEvent {
  return {
    key: `target-${change.enemy_guid}-${change.at_ms}`,
    kind: 'target-change',
    atMs: change.at_ms,
    text: `${short(change.enemy_name)} turned to ${short(change.to_name)}`,
  };
}

export function castEvent(
  atMs: number,
  spell: string,
  caster: string,
  spellId: number,
  iconUrl: string | null = null,
): ChartEvent {
  return {
    key: `cast-${spellId}-${caster}-${atMs}`,
    kind: 'cast',
    atMs,
    text: `${spell} · ${short(caster)}`,
    iconUrl,
  };
}

export function phaseEvent(phase: ChartPhase): ChartEvent {
  return {
    key: `phase-${phase.name}-${phase.startMs}`,
    kind: 'phase',
    atMs: phase.startMs,
    text: `Phase · ${phase.name}`,
  };
}

/** Within one second, the order a raider reads a wipe in: the death, who the boss turned to, the taunt. */
const KIND_ORDER: Readonly<Record<ChartEvent['kind'], number>> = {
  death: 0,
  'target-change': 1,
  taunt: 2,
  cast: 3,
  phase: 4,
};

/** Every event within ±2 s of `ms`, by time (to the second) and then by kind (spec §B5.6, fix round 2). */
export function eventsNear(
  events: readonly ChartEvent[],
  ms: number,
  radius = EVENT_RADIUS_MS,
): ChartEvent[] {
  return events
    .filter((event) => Math.abs(event.atMs - ms) <= radius)
    .sort(
      (a, b) =>
        Math.floor(a.atMs / 1000) - Math.floor(b.atMs / 1000) ||
        KIND_ORDER[a.kind] - KIND_ORDER[b.kind] ||
        a.atMs - b.atMs,
    );
}

/** One marker in the lane: a single event, or a counted cluster of them. */
export interface LaneItem {
  key: string;
  kind: 'single' | 'cluster';
  /** `death` draws the death circle; `gold` every other marker. */
  tone: 'death' | 'gold';
  events: ChartEvent[];
  /** Where the marker is drawn, after the collision step. */
  x: number;
  /** Where its time is: a stepped marker draws a leader back to here. */
  timeX: number;
  leader: boolean;
  startMs: number;
  endMs: number;
}

/** Runs of three or more events inside 44 px become one cluster; shorter runs stay single. */
function collapse(
  events: readonly ChartEvent[],
  xOf: (ms: number) => number,
  tone: LaneItem['tone'],
): Omit<LaneItem, 'x' | 'leader'>[] {
  const sorted = [...events].sort((a, b) => a.atMs - b.atMs);
  const out: Omit<LaneItem, 'x' | 'leader'>[] = [];
  let index = 0;
  while (index < sorted.length) {
    const startX = xOf(sorted[index].atMs);
    let end = index;
    while (end + 1 < sorted.length && xOf(sorted[end + 1].atMs) - startX <= LANE_CLUSTER_PX) end += 1;
    const run = sorted.slice(index, end + 1);
    if (run.length >= 3) {
      out.push({
        key: `cluster-${run[0].key}`,
        kind: 'cluster',
        tone,
        events: run,
        timeX: startX,
        startMs: run[0].atMs,
        endMs: run[run.length - 1].atMs,
      });
      index = end + 1;
    } else {
      const event = sorted[index];
      out.push({
        key: event.key,
        kind: 'single',
        tone,
        events: [event],
        timeX: startX,
        startMs: event.atMs,
        endMs: event.atMs,
      });
      index += 1;
    }
  }
  return out;
}

/** The lane's markers, collapsed and stepped apart (phases have no lane marker). */
export function laneItems(events: readonly ChartEvent[], xOf: (ms: number) => number): LaneItem[] {
  const marked = events.filter((event) => event.kind !== 'phase');
  const items = [
    ...collapse(
      marked.filter((event) => event.kind === 'death'),
      xOf,
      'death',
    ),
    ...collapse(
      marked.filter((event) => event.kind !== 'death'),
      xOf,
      'gold',
    ),
  ].sort((a, b) => a.timeX - b.timeX);
  let previous = Number.NEGATIVE_INFINITY;
  return items.map((item) => {
    const x = Math.max(item.timeX, previous + LANE_STEP_PX);
    previous = x;
    return { ...item, x, leader: x !== item.timeX };
  });
}

/** A death cluster's tooltip rows: six, then a count of the rest. */
export function clusterRows(events: readonly ChartEvent[]): { shown: ChartEvent[]; more: number } {
  return { shown: events.slice(0, CLUSTER_ROWS), more: Math.max(0, events.length - CLUSTER_ROWS) };
}
