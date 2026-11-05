// web/src/lib/sim/timeline.ts
// The saved sim timeline as data (spec 2026-11-04-sim-results 3.7): one pass over the median
// run's events into lanes, the aura segments, the resource line, the zoom window's maths, the
// crosshair's lookups and the keyboard model. Pure, with no DOM and no Svelte, so a test can pin
// every rule and the main chunk can size the card's skeleton from the same lane rules.
import { parseActionKey, resolveActionName, type ActionNames } from './action-names';
import { simCopy } from './copy';
import { timelineCopy } from './timeline-copy';
import type { SampleCast, SampleOutcome, SimResult } from './types';

/** Spec 8: the most events a timeline draws. */
export const MAX_EVENTS = 5_000;
/** Close up's first window. */
export const CLOSE_UP_MS = 30_000;
/** The brush never resizes below this. */
export const MIN_WINDOW_MS = 10_000;
/** One arrow key on the brush. */
export const BRUSH_STEP_MS = 5_000;
/** How many casts the readout lists. */
export const READOUT_CASTS = 5;
/** An aura up longer than this share of the run is left out of the readout's "On you" line: it
 *  says nothing about this moment (Flurry, spec 3.7). */
export const READOUT_AURA_MAX_SHARE = 0.5;
/** Without `sample_duration_ms` the axis ends this long after the last event (spec 4.3). */
const AXIS_TAIL_MS = 1_000;

/** Every rank of Execute: the phase wash starts at the first of these in the sample. */
const EXECUTE_SPELL_IDS: ReadonlySet<number> = new Set([5308, 20658, 20660, 20661, 20662]);
const OVERPOWER_KEY = 'spell:11585';
const SWING_LABEL = 'attack';
const PERIODIC_ORDER = 99;

/** The bit the engine's school mask sets, to the design system's school token. */
const SCHOOL_TOKENS: Readonly<Record<number, string>> = {
  1: '--color-school-physical',
  2: '--color-school-holy',
  4: '--color-school-fire',
  8: '--color-school-nature',
  16: '--color-school-frost',
  32: '--color-school-shadow',
  64: '--color-school-arcane',
};

/** The resources a timeline can draw, in the order it prefers them. */
const LINE_RESOURCES = ['rage', 'energy', 'focus', 'mana'] as const;
const RESOURCE_CAP: Readonly<Record<string, number>> = { rage: 100, energy: 100, focus: 100 };

export type LaneGroup = 'casts' | 'swings' | 'player' | 'target';

export interface TimelineEvent {
  atMs: number;
  /** The event's own action key, tags included. */
  action: string;
  resource: number | null;
  outcome?: SampleOutcome;
  amount?: number;
  /** A queue press (`spell:11567/1`): hollow, in Close up only, and never a landing. */
  queued: boolean;
}

export interface TickLane {
  kind: 'ticks';
  id: string;
  group: 'casts' | 'swings';
  /** The key its name and icon resolve from. */
  action: string;
  name: string;
  /** A school token, or null for a cast with no damage row (a cooldown). */
  schoolToken: string | null;
  /** A lane of periodic ticks: drawn plain at every outcome. */
  periodic: boolean;
  /** Landings, in time order. */
  events: TimelineEvent[];
  /** Queue presses, in time order. */
  queued: TimelineEvent[];
}

export interface BarLane {
  kind: 'bars';
  id: string;
  group: 'player' | 'target' | 'window';
  action: string;
  name: string;
  segments: readonly (readonly [number, number])[];
  stacks: readonly (readonly [number, number])[];
  /** Share of the run this aura was up. */
  share: number;
}

export type TimelineLane = TickLane | BarLane;

export interface ResourceLine {
  key: string;
  label: string;
  /** [at_ms, value after that event]. */
  points: readonly (readonly [number, number])[];
  peak: number;
  cap: number;
}

export interface Landing {
  event: TimelineEvent;
  lane: TickLane;
}

export interface TimelineModel {
  startMs: number;
  endMs: number;
  casts: TickLane[];
  swings: TickLane[];
  onYou: BarLane[];
  onTarget: BarLane[];
  /** The Overpower window lane: only when the rotation casts Overpower. */
  overpower: BarLane | null;
  resource: ResourceLine | null;
  executeAtMs: number | null;
  /** False on an older result: the bars give way to spec 4.3's one line. */
  hasAuras: boolean;
  hasOutcomes: boolean;
  /** Every press in fight order: the cast list, and what the plot's label counts. */
  presses: TimelineEvent[];
  /** Every landing on a Casts lane in fight order: what the readout lists. */
  landings: readonly Landing[];
  /** The periodic lane's own name for the cast list's foot, when the fight has one. */
  periodicName: string | null;
}

/** `spell:11567/1+r3` to `spell:11567`. A key outside the grammar is its own base. */
export function baseActionKey(key: string): string {
  const parsed = parseActionKey(key);
  if (parsed === null || parsed.kind === 'unknown' || parsed.kind === 'other') return key;
  return `${parsed.kind}:${parsed.label}`;
}

function schoolToken(mask: number | null | undefined): string | null {
  if (mask === null || mask === undefined) return null;
  const bit = Object.keys(SCHOOL_TOKENS)
    .map(Number)
    .find((candidate) => (mask & candidate) !== 0);
  return bit === undefined ? null : SCHOOL_TOKENS[bit]!;
}

interface AbilityFacts {
  /** Base key to the school token of the row that dealt damage. */
  schools: Map<string, string>;
  /** Base keys whose row records periodic ticks. */
  periodic: Set<string>;
}

function abilityFacts(result: SimResult): AbilityFacts {
  const schools = new Map<string, string>();
  const periodic = new Set<string>();
  for (const actor of result.summary.damage_done ?? []) {
    for (const ability of actor.abilities) {
      const base = baseActionKey(ability.name);
      const token = ability.total > 0 ? schoolToken(ability.school) : null;
      if (token !== null) schools.set(base, token);
      if ((ability.ticks ?? 0) > 0) periodic.add(base);
    }
  }
  return { schools, periodic };
}

function castKeys(result: SimResult): Set<string> {
  return new Set(
    result.summary.casts.filter((row) => row.succeeded > 0).map((row) => baseActionKey(row.spell_name)),
  );
}

function resourceOf(cast: SampleCast, key: string | null): number | null {
  return key === null ? null : (cast.resources?.[key] ?? null);
}

/** The pool this class plays: rage, energy or focus before mana (a cat's energy, not its mana). */
function lineResource(sample: readonly SampleCast[]): string | null {
  const seen = new Set<string>();
  for (const cast of sample) {
    for (const key of Object.keys(cast.resources ?? {})) seen.add(key);
  }
  return LINE_RESOURCES.find((key) => seen.has(key)) ?? null;
}

function buildResource(sample: readonly SampleCast[], key: string | null): ResourceLine | null {
  if (key === null) return null;
  const points: [number, number][] = [];
  let peak = 0;
  for (const cast of sample) {
    const value = resourceOf(cast, key);
    if (value === null) continue;
    points.push([cast.at_ms, value]);
    peak = Math.max(peak, value);
  }
  if (points.length === 0) return null;
  const cap = Math.max(RESOURCE_CAP[key] ?? peak, peak);
  return { key, label: simCopy.resourceLabel[key] ?? key, points, peak: Math.round(peak), cap };
}

interface Placement {
  group: 'casts' | 'swings';
  laneId: string;
  action: string;
  periodic: boolean;
}

/** Where one event draws (spec 3.7, lane membership). Null: an event no lane shows. */
function placeEvent(cast: SampleCast, facts: AbilityFacts, casts: ReadonlySet<string>): Placement | null {
  const parsed = parseActionKey(cast.action);
  if (parsed?.kind === 'other') {
    if (parsed.label !== SWING_LABEL) return null;
    return { group: 'swings', laneId: cast.action, action: cast.action, periodic: false };
  }
  const base = baseActionKey(cast.action);
  const isTick = cast.outcome === 'tick' || cast.outcome === 'tick_crit';
  if (isTick || (facts.periodic.has(base) && !casts.has(base))) {
    return { group: 'swings', laneId: `ticks:${base}`, action: base, periodic: true };
  }
  return { group: 'casts', laneId: base, action: base, periodic: false };
}

function laneName(place: Placement, names: ActionNames | null): string {
  const parsed = parseActionKey(place.action);
  if (parsed?.kind === 'other' && parsed.label === SWING_LABEL) {
    const hand = timelineCopy.hands[parsed.tag];
    if (hand !== undefined) return hand;
  }
  const name = resolveActionName(place.action, names);
  return place.periodic ? timelineCopy.ticksLane(name) : name;
}

function isExecute(key: string): boolean {
  const parsed = parseActionKey(key);
  return parsed?.kind === 'spell' && EXECUTE_SPELL_IDS.has(parsed.id);
}

/** Moves a tagged key's events to queue presses when the same lane also records landings. */
function isQueuePress(cast: SampleCast, base: string, seen: ReadonlySet<string>): boolean {
  return cast.action !== base && (parseActionKey(cast.action)?.tag ?? 0) !== 0 && seen.has(base);
}

function buildBars(result: SimResult, names: ActionNames | null, spanMs: number): BarLane[] {
  return (result.sample_auras ?? []).map((aura) => {
    const up = aura.segments.reduce((total, [start, end]) => total + Math.max(0, end - start), 0);
    const group = aura.on === 'target' ? ('target' as const) : ('player' as const);
    return {
      kind: 'bars' as const,
      id: `${aura.on}:${aura.action}`,
      group,
      action: aura.action,
      name: resolveActionName(aura.action, names),
      segments: aura.segments,
      stacks: aura.stacks ?? [],
      share: spanMs <= 0 ? 0 : up / spanMs,
    };
  });
}

/** Main hand, off hand, extra attacks, then the periodic lanes in first use (spec 3.7). */
function swingOrder(lane: TickLane): number {
  return lane.periodic ? PERIODIC_ORDER : (parseActionKey(lane.action)?.tag ?? PERIODIC_ORDER);
}

const firstStart = (lane: BarLane): number => lane.segments[0]?.[0] ?? Number.POSITIVE_INFINITY;
const byFirstApplication = (a: BarLane, b: BarLane): number => firstStart(a) - firstStart(b);

/** The whole timeline, in one pass over the sample. Null when the run recorded no fight. */
export function buildTimeline(result: SimResult, names: ActionNames | null): TimelineModel | null {
  const sample = (result.sample ?? []).slice(0, MAX_EVENTS);
  const first = sample[0];
  const last = sample.at(-1);
  if (first === undefined || last === undefined) return null;

  const facts = abilityFacts(result);
  const casted = castKeys(result);
  const resourceKey = lineResource(sample);
  const seen = new Set(sample.map((cast) => cast.action));
  const lanes = new Map<string, TickLane>();
  const presses: TimelineEvent[] = [];
  let executeAtMs: number | null = null;
  let hasOutcomes = false;

  for (const cast of sample) {
    if (cast.outcome !== undefined) hasOutcomes = true;
    const place = placeEvent(cast, facts, casted);
    if (place === null) continue;
    let lane = lanes.get(place.laneId);
    if (lane === undefined) {
      lane = {
        kind: 'ticks',
        id: place.laneId,
        group: place.group,
        action: place.action,
        name: laneName(place, names),
        schoolToken: facts.schools.get(baseActionKey(place.action)) ?? null,
        periodic: place.periodic,
        events: [],
        queued: [],
      };
      lanes.set(place.laneId, lane);
    }
    const queued = place.group === 'casts' && isQueuePress(cast, place.action, seen);
    const event: TimelineEvent = {
      atMs: cast.at_ms,
      action: cast.action,
      resource: resourceOf(cast, resourceKey),
      queued,
      ...(cast.outcome === undefined ? {} : { outcome: cast.outcome }),
      ...(cast.amount === undefined ? {} : { amount: cast.amount }),
    };
    (queued ? lane.queued : lane.events).push(event);
    if (executeAtMs === null && isExecute(cast.action)) executeAtMs = cast.at_ms;
  }

  const allLanes = [...lanes.values()];
  const casts = allLanes.filter((lane) => lane.group === 'casts');
  for (const lane of casts) {
    // The list counts presses (spec 3.7): a queued spell's presses, not its landings as well.
    presses.push(...(lane.queued.length > 0 ? lane.queued : lane.events));
  }
  presses.sort((a, b) => a.atMs - b.atMs);

  const endMs = result.sample_duration_ms ?? last.at_ms + AXIS_TAIL_MS;
  const bars = buildBars(result, names, endMs - Math.min(first.at_ms, 0));
  const overpower = casted.has(OVERPOWER_KEY)
    ? bars.find((bar) => bar.group === 'player' && baseActionKey(bar.action) === OVERPOWER_KEY)
    : undefined;
  const others = bars.filter((bar) => baseActionKey(bar.action) !== OVERPOWER_KEY);
  const swings = allLanes
    .filter((lane) => lane.group === 'swings')
    .sort((a, b) => swingOrder(a) - swingOrder(b));
  const periodic = swings.find((lane) => lane.periodic);

  return {
    startMs: first.at_ms,
    endMs,
    casts,
    swings,
    onYou: others.filter((bar) => bar.group === 'player').sort(byFirstApplication),
    onTarget: others.filter((bar) => bar.group === 'target').sort(byFirstApplication),
    overpower:
      overpower === undefined
        ? null
        : { ...overpower, group: 'window', id: 'window:overpower', name: timelineCopy.overpowerWindow },
    resource: buildResource(sample, resourceKey),
    executeAtMs,
    hasAuras: result.sample_auras !== undefined,
    hasOutcomes,
    presses,
    landings: casts
      .flatMap((lane) => lane.events.map((event) => ({ event, lane })))
      .sort((a, b) => a.event.atMs - b.event.atMs),
    periodicName: periodic === undefined ? null : resolveActionName(periodic.action, names),
  };
}

// --- Time ---

/** `2:25.0`, and `-0:02.5` before the pull. */
export function preciseClock(ms: number): string {
  const tenths = Math.round(Math.abs(ms) / 100);
  const minutes = Math.floor(tenths / 600);
  const seconds = (tenths % 600) / 10;
  return `${ms < 0 ? '-' : ''}${minutes}:${seconds.toFixed(1).padStart(4, '0')}`;
}

// --- The zoom window ---

export interface ZoomWindow {
  startMs: number;
  endMs: number;
}

function clampWindow(window: ZoomWindow, model: Pick<TimelineModel, 'startMs' | 'endMs'>): ZoomWindow {
  const length = Math.min(window.endMs - window.startMs, model.endMs - model.startMs);
  const startMs = Math.max(model.startMs, Math.min(model.endMs - length, window.startMs));
  return { startMs, endMs: startMs + length };
}

/** Close up: 30 s centred on the crosshair, or opening at the pull when there is none. */
export function closeUpWindow(
  model: Pick<TimelineModel, 'startMs' | 'endMs'>,
  crosshairMs: number | null,
): ZoomWindow {
  const startMs = crosshairMs === null ? Math.max(model.startMs, 0) : crosshairMs - CLOSE_UP_MS / 2;
  return clampWindow({ startMs, endMs: startMs + CLOSE_UP_MS }, model);
}

/** The window moved by `deltaMs`, its length kept, inside the fight. */
export function panWindow(
  window: ZoomWindow,
  deltaMs: number,
  model: Pick<TimelineModel, 'startMs' | 'endMs'>,
): ZoomWindow {
  return clampWindow({ startMs: window.startMs + deltaMs, endMs: window.endMs + deltaMs }, model);
}

/** One edge moved by `deltaMs`: never under 10 s (or the whole fight, if that is shorter), never out of the fight. */
export function resizeWindow(
  window: ZoomWindow,
  edge: 'start' | 'end',
  deltaMs: number,
  model: Pick<TimelineModel, 'startMs' | 'endMs'>,
): ZoomWindow {
  const least = Math.min(MIN_WINDOW_MS, model.endMs - model.startMs);
  if (edge === 'start') {
    const startMs = Math.max(model.startMs, Math.min(window.endMs - least, window.startMs + deltaMs));
    return { startMs, endMs: window.endMs };
  }
  const endMs = Math.min(model.endMs, Math.max(window.startMs + least, window.endMs + deltaMs));
  return { startMs: window.startMs, endMs };
}

/** The window pulled over `atMs` when it sits outside: a keyboard move never leaves the view. */
export function windowIncluding(
  window: ZoomWindow,
  atMs: number,
  model: Pick<TimelineModel, 'startMs' | 'endMs'>,
): ZoomWindow {
  if (atMs >= window.startMs && atMs <= window.endMs) return window;
  const length = window.endMs - window.startMs;
  return clampWindow({ startMs: atMs - length / 2, endMs: atMs + length / 2 }, model);
}

// --- Merged ticks ---

export interface DrawnTick {
  x: number;
  event: TimelineEvent;
  /** The events this tick stands for: more than one when they sat under 1 px apart. */
  count: number;
  /** The event's index in its lane's list. */
  index: number;
}

/** The px between two ticks below which they merge into one (spec 8). */
export const MERGE_UNDER_PX = 1;

/** A lane's events inside the window, those under 1 px apart merged into the first. */
export function mergeTicks(
  events: readonly TimelineEvent[],
  window: ZoomWindow,
  pxPerMs: number,
): DrawnTick[] {
  const drawn: DrawnTick[] = [];
  let lastX = Number.NEGATIVE_INFINITY;
  events.forEach((event, index) => {
    if (event.atMs < window.startMs || event.atMs > window.endMs) return;
    const x = (event.atMs - window.startMs) * pxPerMs;
    const previous = drawn.at(-1);
    if (previous !== undefined && x - lastX < MERGE_UNDER_PX) {
      drawn[drawn.length - 1] = { ...previous, count: previous.count + 1 };
    } else {
      drawn.push({ x, event, count: 1, index });
      lastX = x;
    }
  });
  return drawn;
}

/** Segments cut to the window, with whatever falls outside dropped. */
export function clipSegments(
  segments: readonly (readonly [number, number])[],
  window: ZoomWindow,
): [number, number][] {
  const clipped: [number, number][] = [];
  for (const [start, end] of segments) {
    const from = Math.max(start, window.startMs);
    const to = Math.min(end, window.endMs);
    if (to > from) clipped.push([from, to]);
  }
  return clipped;
}

/** The stack count in force at `atMs`; 0 when a stacking aura has no change point before it. */
export function stacksAt(stacks: readonly (readonly [number, number])[], atMs: number): number {
  let count = 0;
  for (const [at, value] of stacks) {
    if (at > atMs) break;
    count = value;
  }
  return count;
}

// --- The crosshair's lookups ---

/** The last index whose time is at or before `atMs`, or -1. Binary, so a hover stays cheap. */
function lastAtOrBefore<T>(rows: readonly T[], atMs: number, timeOf: (row: T) => number): number {
  let low = 0;
  let high = rows.length - 1;
  let found = -1;
  while (low <= high) {
    const middle = (low + high) >> 1;
    if (timeOf(rows[middle]!) <= atMs) {
      found = middle;
      low = middle + 1;
    } else {
      high = middle - 1;
    }
  }
  return found;
}

/** The resource held at `atMs`: the value after the last event at or before it. */
export function resourceAt(line: ResourceLine | null, atMs: number): number | null {
  if (line === null) return null;
  const index = lastAtOrBefore(line.points, atMs, (point) => point[0]);
  return index < 0 ? null : line.points[index]![1];
}

/** The drawn tick within `tolerancePx` of `x`, nearest first; null when the pointer is on bare lane. */
export function tickNear(ticks: readonly DrawnTick[], x: number, tolerancePx: number): DrawnTick | null {
  let best: DrawnTick | null = null;
  let bestDistance = tolerancePx;
  for (const tick of ticks) {
    const distance = Math.abs(tick.x - x);
    if (distance <= bestDistance) {
      best = tick;
      bestDistance = distance;
    }
    if (tick.x > x + tolerancePx) break;
  }
  return best;
}

export interface ReadoutCast {
  action: string;
  name: string;
  agoMs: number;
}

export interface Readout {
  atMs: number;
  time: string;
  resource: { label: string; value: string } | null;
  onYou: string[];
  casts: ReadoutCast[];
}

/** What a readout reads at `atMs` (spec 3.7). */
export function readoutAt(model: TimelineModel, atMs: number): Readout {
  const value = resourceAt(model.resource, atMs);
  const upTo = lastAtOrBefore(model.landings, atMs, (row) => row.event.atMs);
  return {
    atMs,
    time: preciseClock(atMs),
    resource:
      model.resource === null || value === null
        ? null
        : { label: model.resource.label, value: String(Math.round(value)) },
    onYou: model.onYou
      .filter(
        (bar) =>
          bar.share < READOUT_AURA_MAX_SHARE &&
          bar.segments.some(([start, end]) => start <= atMs && atMs < end),
      )
      .map((bar) => bar.name),
    casts: model.landings
      .slice(Math.max(0, upTo - READOUT_CASTS + 1), upTo + 1)
      .reverse()
      .map(({ event, lane }) => ({ action: lane.action, name: lane.name, agoMs: atMs - event.atMs })),
  };
}

export interface TickFollow {
  name: string;
  /** Ms to the next event on this lane; null on its last. */
  nextInMs: number | null;
}

/** The tooltip's third line: the next landing on the same lane. */
export function tickFollow(lane: TickLane, index: number): TickFollow {
  const current = lane.events[index];
  const next = lane.events[index + 1];
  return {
    name: lane.name,
    nextInMs: current === undefined || next === undefined ? null : next.atMs - current.atMs,
  };
}

/** One decimal, for "0.5 s ago" and "1.5 s later". */
export function seconds(ms: number): string {
  return (ms / 1000).toFixed(1);
}

// --- The keyboard model ---

export interface Caret {
  /** The index into the visible tick lanes. */
  lane: number;
  /** The index into that lane's landings. */
  index: number;
}

export type CaretKey = 'ArrowLeft' | 'ArrowRight' | 'ArrowUp' | 'ArrowDown' | 'Home' | 'End';

/** The lanes a keyboard moves over: Casts, then Swings and ticks when they are on. */
export function caretLanes(model: TimelineModel, swings: boolean): TickLane[] {
  return swings ? [...model.casts, ...model.swings] : model.casts;
}

/** Every landing across `lanes`, in time order; ties keep the lane order. */
function ordered(lanes: readonly TickLane[]): (Caret & { atMs: number })[] {
  return lanes
    .flatMap((lane, laneIndex) =>
      lane.events.map((event, index) => ({ lane: laneIndex, index, atMs: event.atMs })),
    )
    .sort((a, b) => a.atMs - b.atMs || a.lane - b.lane);
}

/** The event in `lane` nearest in time to `atMs`. */
function nearestIn(lane: TickLane, atMs: number): number {
  const before = lastAtOrBefore(lane.events, atMs, (event) => event.atMs);
  const after = before + 1;
  if (before < 0) return after;
  if (after >= lane.events.length) return before;
  const toBefore = atMs - lane.events[before]!.atMs;
  const toAfter = lane.events[after]!.atMs - atMs;
  return toAfter < toBefore ? after : before;
}

/** The first lane with landings past `from`, stepping by `step`; null at the edge. */
function neighbour(lanes: readonly TickLane[], from: number, step: 1 | -1): number | null {
  for (let lane = from + step; lane >= 0 && lane < lanes.length; lane += step) {
    if (lanes[lane]!.events.length > 0) return lane;
  }
  return null;
}

/** Where a key moves the caret (spec 3.7); null when it has nowhere to go. */
export function moveCaret(lanes: readonly TickLane[], caret: Caret | null, key: CaretKey): Caret | null {
  const all = ordered(lanes);
  const first = all[0];
  const last = all.at(-1);
  if (first === undefined || last === undefined) return null;
  const home = { lane: first.lane, index: first.index };
  const end = { lane: last.lane, index: last.index };
  if (key === 'Home') return home;
  if (key === 'End') return end;
  if (caret === null) return home;

  if (key === 'ArrowUp' || key === 'ArrowDown') {
    const lane = neighbour(lanes, caret.lane, key === 'ArrowDown' ? 1 : -1);
    if (lane === null) return null;
    const atMs = lanes[caret.lane]?.events[caret.index]?.atMs ?? 0;
    return { lane, index: nearestIn(lanes[lane]!, atMs) };
  }
  const at = all.findIndex((row) => row.lane === caret.lane && row.index === caret.index);
  const next = all[at + (key === 'ArrowRight' ? 1 : -1)];
  return next === undefined ? null : { lane: next.lane, index: next.index };
}

/** The caret on the landing nearest `atMs`: Close up and the crosshair keep their place. */
export function caretAt(lanes: readonly TickLane[], atMs: number): Caret | null {
  const all = ordered(lanes);
  if (all.length === 0) return null;
  const nearest = all.reduce((best, row) =>
    Math.abs(row.atMs - atMs) < Math.abs(best.atMs - atMs) ? row : best,
  );
  return { lane: nearest.lane, index: nearest.index };
}

// --- The cast list ---

/** The presses inside the window, in fight order. */
export function pressesInView(model: TimelineModel, window: ZoomWindow): TimelineEvent[] {
  return model.presses.filter((press) => press.atMs >= window.startMs && press.atMs <= window.endMs);
}

// --- Rows, and the card's ready height ---

export const ROW_HEIGHT = {
  toolbar: 56,
  axis: 30,
  group: 30,
  lane: 24,
  window: 16,
  resource: 72,
  /** Spec 4.3's one line in place of the bar groups. */
  notice: 64,
  /** The same line and its link, wrapped in a phone's column. */
  noticePhone: 112,
  brush: 44,
  castFoot: 52,
  castRow: 40,
} as const;

/** How many rows the collapsed cast list shows. */
export const COLLAPSED_CAST_ROWS = 5;

export type PlotRow =
  | { kind: 'axis'; y: number; height: number }
  | { kind: 'group'; group: LaneGroup; label: string; y: number; height: number }
  | { kind: 'lane'; lane: TimelineLane; y: number; height: number }
  | { kind: 'notice'; y: number; height: number }
  | { kind: 'resource'; line: ResourceLine; y: number; height: number };

/** The plot's rows, top to bottom (spec 3.7): a group with no lanes does not draw. */
export function layoutRows(
  model: TimelineModel,
  swings: boolean,
  phone = false,
): { rows: PlotRow[]; height: number } {
  const rows: PlotRow[] = [];
  let y = 0;
  const advance = (height: number): { y: number; height: number } => {
    const place = { y, height };
    y += height;
    return place;
  };
  const group = (id: LaneGroup, lanes: readonly TimelineLane[]): void => {
    if (lanes.length === 0) return;
    rows.push({ kind: 'group', group: id, label: timelineCopy.groups[id], ...advance(ROW_HEIGHT.group) });
    for (const lane of lanes) {
      const height = lane.kind === 'bars' && lane.group === 'window' ? ROW_HEIGHT.window : ROW_HEIGHT.lane;
      rows.push({ kind: 'lane', lane, ...advance(height) });
    }
  };

  rows.push({ kind: 'axis', ...advance(ROW_HEIGHT.axis) });
  group('casts', model.overpower === null ? model.casts : [...model.casts, model.overpower]);
  if (swings) group('swings', model.swings);
  if (model.hasAuras) {
    group('player', model.onYou);
    group('target', model.onTarget);
  } else {
    rows.push({ kind: 'notice', ...advance(phone ? ROW_HEIGHT.noticePhone : ROW_HEIGHT.notice) });
  }
  if (model.resource !== null) {
    rows.push({ kind: 'resource', line: model.resource, ...advance(ROW_HEIGHT.resource) });
  }
  return { rows, height: y };
}

/**
 * The card's body at its ready height, CardHead excluded (spec 3.7's loading formula), computed
 * from the same lane rules the plot uses. Null when the run has no fight: the card is one line.
 */
export function timelineReadyHeight(result: SimResult): number | null {
  const model = buildTimeline(result, null);
  if (model === null) return null;
  const castRows = Math.min(COLLAPSED_CAST_ROWS, model.presses.length);
  return (
    ROW_HEIGHT.toolbar +
    layoutRows(model, false).height +
    ROW_HEIGHT.brush +
    ROW_HEIGHT.castFoot +
    ROW_HEIGHT.castRow * castRows
  );
}
