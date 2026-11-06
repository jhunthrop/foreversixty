// web/src/lib/sim/timeline-lanes.ts
// The timeline's lane rules and row heights (spec 2026-11-04-sim-results 3.7), small and pure, so
// the page's main chunk can size the timeline's loading reserve from the very rules the lazy card
// draws with: which event lands on which lane, which aura earns a bar, and how tall each row is.
// lib/sim/timeline.ts builds the drawn model on top of these; nothing here names or formats.
import { clockLabel } from '../chart/time-axis';
import { parseActionKey } from './action-names';
import { timelineCopy } from './timeline-copy';
import type { SampleAura, SampleCast, SimResult } from './types';

/** Spec 8: the most events a timeline draws. */
export const MAX_EVENTS = 5_000;
/** Without `sample_duration_ms` the axis ends this long after the last event (spec 4.3). */
const AXIS_TAIL_MS = 1_000;
const SWING_LABEL = 'attack';
export const OVERPOWER_KEY = 'spell:11585';
/** Lane E: Thrash is not a buff and left the results; an older result may still carry it. */
export const NEVER_DRAWN: ReadonlySet<string> = new Set(['spell:21919']);
/** An aura up this much of the run, from the pull, is the raid's and not the fight's: the Buffs card lists it. */
const ALWAYS_UP_SHARE = 0.99;

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

export const ROW_HEIGHT = {
  axis: 30,
  group: 30,
  lane: 24,
  /** A phone prints each lane's name on its own row above the lane, never over the plot (B1). */
  laneName: 14,
  window: 16,
  resource: 72,
  /** Spec 4.3's one line in place of the bar groups. */
  notice: 64,
  /** The same line and its link, wrapped in a phone's column. */
  noticePhone: 112,
  /** The cast count line under the plot (the list itself opens on "List every cast"). */
  castFoot: 30,
} as const;

/** `spell:11567/1+r3` to `spell:11567`. A key outside the grammar is its own base. */
export function baseActionKey(key: string): string {
  const parsed = parseActionKey(key);
  if (parsed === null || parsed.kind === 'unknown' || parsed.kind === 'other') return key;
  return `${parsed.kind}:${parsed.label}`;
}

export function schoolToken(mask: number | null | undefined): string | null {
  if (mask === null || mask === undefined) return null;
  const bit = Object.keys(SCHOOL_TOKENS)
    .map(Number)
    .find((candidate) => (mask & candidate) !== 0);
  return bit === undefined ? null : SCHOOL_TOKENS[bit]!;
}

export interface AbilityFacts {
  /** Base key to the school token of the row that dealt damage. */
  schools: Map<string, string>;
  /** Base keys whose row records periodic ticks. */
  periodic: Set<string>;
}

export function abilityFacts(result: SimResult): AbilityFacts {
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

export function castKeys(result: SimResult): Set<string> {
  return new Set(
    result.summary.casts.filter((row) => row.succeeded > 0).map((row) => baseActionKey(row.spell_name)),
  );
}

/** The pool this class plays: rage, energy or focus before mana (a cat's energy, not its mana). */
export function lineResource(sample: readonly SampleCast[]): string | null {
  const seen = new Set<string>();
  for (const cast of sample) {
    for (const key of Object.keys(cast.resources ?? {})) seen.add(key);
  }
  return LINE_RESOURCES.find((key) => seen.has(key)) ?? null;
}

export interface Placement {
  group: 'casts' | 'swings';
  laneId: string;
  action: string;
  periodic: boolean;
}

/** Where one event draws (spec 3.7, lane membership). Null: an event no lane shows. */
export function placeEvent(
  cast: SampleCast,
  facts: AbilityFacts,
  casts: ReadonlySet<string>,
  hasOutcomes: boolean,
): Placement | null {
  const parsed = parseActionKey(cast.action);
  if (parsed?.kind === 'other') {
    if (parsed.label !== SWING_LABEL) return null;
    return { group: 'swings', laneId: cast.action, action: cast.action, periodic: false };
  }
  const base = baseActionKey(cast.action);
  if (NEVER_DRAWN.has(base)) return null;
  const isTick = cast.outcome === 'tick' || cast.outcome === 'tick_crit';
  const tickLane = { group: 'swings', laneId: `ticks:${base}`, action: base, periodic: true } as const;
  if (isTick) return tickLane;
  if (facts.periodic.has(base) && !casts.has(base)) {
    // A run with outcomes records each tick as its own row, so a bare row here is an application:
    // neither a tick nor a press. An older run cannot tell them apart and draws them as they are.
    return hasOutcomes ? null : tickLane;
  }
  // Damage with no cast row is a hit no one pressed (the off-hand Whirlwind, a proc): it lands on
  // Swings and ticks, so a lane on Casts never reads as a second press.
  if (casts.size > 0 && !casts.has(base) && facts.schools.has(base)) {
    return { group: 'swings', laneId: base, action: base, periodic: false };
  }
  return { group: 'casts', laneId: base, action: base, periodic: false };
}

/** A tagged key's event is a queue press when the same lane also records landings. */
export function isQueuePress(cast: SampleCast, base: string, seen: ReadonlySet<string>): boolean {
  return cast.action !== base && (parseActionKey(cast.action)?.tag ?? 0) !== 0 && seen.has(base);
}

/** The sample a timeline draws: capped at MAX_EVENTS. */
export function drawnSample(result: SimResult): SampleCast[] {
  return (result.sample ?? []).slice(0, MAX_EVENTS);
}

/** Where the axis ends: the sample's own length, or just past its last event on an older result. */
export function axisEndMs(result: SimResult, sample: readonly SampleCast[]): number {
  return result.sample_duration_ms ?? (sample.at(-1)?.at_ms ?? 0) + AXIS_TAIL_MS;
}

/** Share of the run (to `endMs`) an aura's segments cover. */
export function auraShare(aura: Pick<SampleAura, 'segments'>, endMs: number): number {
  const up = aura.segments.reduce((total, [start, end]) => total + Math.max(0, end - start), 0);
  return endMs <= 0 ? 0 : up / endMs;
}

/**
 * Up from the pull to the end, on you or on the target: the raid's buffs and the raid's debuffs.
 * The Buffs card lists them; nine identical full-length bars say nothing about this fight.
 */
export function upAllFight(aura: Pick<SampleAura, 'segments'>, endMs: number): boolean {
  const first = aura.segments[0];
  return first !== undefined && first[0] <= 0 && auraShare(aura, endMs) >= ALWAYS_UP_SHARE;
}

/** A bar worth a lane: not a removed effect, not a queue's own window, not up all fight. */
export function isFightAura(aura: SampleAura, seen: ReadonlySet<string>, endMs: number): boolean {
  const base = baseActionKey(aura.action);
  if (NEVER_DRAWN.has(base)) return false;
  if (aura.action !== base && seen.has(base) && seen.has(aura.action)) return false;
  return !upAllFight(aura, endMs);
}

/** How many of each row the timeline draws with Swings and ticks off: its ready height's inputs. */
export interface TimelineShape {
  castLanes: number;
  /** The Overpower window lane (16 px), when the rotation casts Overpower. */
  windowLane: boolean;
  playerLanes: number;
  targetLanes: number;
  hasAuras: boolean;
  hasResource: boolean;
  presses: number;
}

/** The timeline's shape in one light pass, with no names: null when the run recorded no fight. */
export function timelineShape(result: SimResult): TimelineShape | null {
  const sample = drawnSample(result);
  if (sample.length === 0) return null;
  const facts = abilityFacts(result);
  const casted = castKeys(result);
  const seen = new Set(sample.map((cast) => cast.action));
  const hasOutcomes = sample.some((cast) => cast.outcome !== undefined);
  const landings = new Map<string, number>();
  const queued = new Map<string, number>();
  for (const cast of sample) {
    const place = placeEvent(cast, facts, casted, hasOutcomes);
    if (place?.group !== 'casts') continue;
    const counts = isQueuePress(cast, place.action, seen) ? queued : landings;
    counts.set(place.laneId, (counts.get(place.laneId) ?? 0) + 1);
    if (counts === queued && !landings.has(place.laneId)) landings.set(place.laneId, 0);
  }
  const presses = [...landings].reduce((sum, [lane, count]) => sum + (queued.get(lane) ?? count), 0);

  const endMs = axisEndMs(result, sample);
  const auras = result.sample_auras ?? [];
  const windowLane =
    casted.has(OVERPOWER_KEY) &&
    auras.some((aura) => aura.on !== 'target' && baseActionKey(aura.action) === OVERPOWER_KEY);
  const fight = auras.filter(
    (aura) => baseActionKey(aura.action) !== OVERPOWER_KEY && isFightAura(aura, seen, endMs),
  );
  return {
    castLanes: landings.size,
    windowLane,
    playerLanes: fight.filter((aura) => aura.on !== 'target').length,
    targetLanes: fight.filter((aura) => aura.on === 'target').length,
    hasAuras: result.sample_auras !== undefined,
    hasResource: lineResource(sample) !== null,
    presses,
  };
}

/** One lane's height: a phone adds its name row above it. */
export function laneHeight(window: boolean, phone: boolean): number {
  return (window ? ROW_HEIGHT.window : ROW_HEIGHT.lane) + (phone ? ROW_HEIGHT.laneName : 0);
}

/** The plot's height, axis to resource line, with Swings and ticks off. */
export function plotHeightOf(shape: TimelineShape, phone: boolean): number {
  /** A group's heading and its lanes; a group with no lanes does not draw. */
  const group = (lanes: number, extra = 0): number =>
    lanes === 0 && extra === 0 ? 0 : ROW_HEIGHT.group + lanes * laneHeight(false, phone) + extra;
  const casts = group(shape.castLanes, shape.windowLane ? laneHeight(true, phone) : 0);
  const auras = shape.hasAuras
    ? group(shape.playerLanes) + group(shape.targetLanes)
    : phone
      ? ROW_HEIGHT.noticePhone
      : ROW_HEIGHT.notice;
  return ROW_HEIGHT.axis + casts + auras + (shape.hasResource ? ROW_HEIGHT.resource : 0);
}

/** The card head's meta, "The median run · 2:53 long": the loading head prints it too, so nothing pops in. */
export function timelineHeadMeta(result: SimResult): string | undefined {
  const sample = drawnSample(result);
  return sample.length === 0 ? undefined : timelineCopy.meta(clockLabel(axisEndMs(result, sample)));
}
