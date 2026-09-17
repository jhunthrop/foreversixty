// web/src/lib/report/window.ts
// Rescoping a fight to a time window, in the browser, from the summary alone.
//
// What the summary can answer exactly, it answers exactly:
//   an actor's total       the sum of its one-second series over the window
//   auras                  segments intersected with the window
//   casts                  the millisecond sequence filtered
//   deaths                 timestamps filtered
//   resources              the series sliced
// What it cannot: the per-ability and per-target split, which the engine stores only as
// whole-fight totals. Those are scaled by the window's share of the actor's total and
// marked `approximate`, which the tables render with a leading `~`. The exact answer is a
// query away -- the Queries view reads the fight's events.parquet -- and that is the
// deliberate trade: every table rescopes in under a frame with no network, and the one
// question that needs precision says so.
import { splitUnitName } from '../characters';
import { formatDuration } from './format';
import type {
  Actor,
  AuraTrack,
  CastRow,
  Death,
  ResourceTrack,
  RosterRow,
  Summary,
  ThreatPair,
  ThreatRow,
} from './types';
import type { ReportState } from './url';

/** The engine's summary.Options.Bucket default. */
export const BUCKET_MS = 1000;

export interface TimeWindow {
  startMs: number;
  endMs: number;
}

export function fullWindow(durationMs: number): TimeWindow {
  return { startMs: 0, endMs: Math.max(durationMs, 0) };
}

export function windowMs(window: TimeWindow): number {
  return Math.max(window.endMs - window.startMs, 0);
}

/**
 * Clamped to the fight and snapped outward to whole seconds. The series are one-second
 * buckets, so a window that starts mid-bucket either counts or drops that bucket
 * whole; snapping the bounds to the buckets makes the tables, the chart and the Queries
 * view (which reads the same bounds) agree to the number.
 */
export function clampWindow(window: TimeWindow, durationMs: number): TimeWindow {
  const startMs = Math.max(0, Math.min(Math.floor(window.startMs / BUCKET_MS) * BUCKET_MS, durationMs));
  const endMs = Math.max(startMs, Math.min(Math.ceil(window.endMs / BUCKET_MS) * BUCKET_MS, durationMs));
  return { startMs, endMs };
}

export function isFullWindow(window: TimeWindow, durationMs: number): boolean {
  return window.startMs <= 0 && window.endMs >= durationMs;
}

/** The URL's start and end, clamped; a window that covers the fight is the whole fight. */
export function windowOf(state: Pick<ReportState, 'start' | 'end'>, durationMs: number): TimeWindow {
  if (state.start === null || state.end === null) return fullWindow(durationMs);
  const clamped = clampWindow({ startMs: state.start, endMs: state.end }, durationMs);
  return windowMs(clamped) === 0 ? fullWindow(durationMs) : clamped;
}

function bucketRange(window: TimeWindow): { from: number; to: number } {
  return { from: Math.floor(window.startMs / BUCKET_MS), to: Math.ceil(window.endMs / BUCKET_MS) };
}

export function sliceSeries(series: number[], window: TimeWindow): number[] {
  const { from, to } = bucketRange(window);
  return series.slice(from, to);
}

export function sumSeries(series: number[], window: TimeWindow): number {
  return sliceSeries(series, window).reduce((total, value) => total + value, 0);
}

/** One line for the chart: the element-wise sum of every actor's series. */
export function combinedSeries(actors: Actor[]): number[] {
  const length = actors.reduce((longest, actor) => Math.max(longest, actor.series.length), 0);
  const combined = new Array<number>(length).fill(0);
  for (const actor of actors) {
    actor.series.forEach((value, index) => {
      combined[index] += value;
    });
  }
  return combined;
}

/** An Actor rescoped to a window. `approximate` is true when the split was scaled. */
export interface ScopedActor extends Actor {
  approximate: boolean;
}

function scale(value: number, ratio: number): number {
  return Math.round(value * ratio);
}

/**
 * Active time, measured rather than capped: one bucket per second the actor actually put
 * a number out. The engine's own `active_ms` is millisecond-accurate but whole-fight, and
 * `min(active_ms, windowMs)` is a ceiling on it, not a measurement -- a short window over
 * an idle stretch would come back fully active, which is the one figure here that would
 * look right while being wrong. The series is what the summary knows about *when*, so the
 * window's answer comes from the series, at its one-second resolution. The whole-fight
 * path never reaches this: it returns the engine's exact figure untouched.
 *
 * Clamped to the window because the two are on different grids. The range inputs step by
 * BUCKET_MS, but a canvas drag does not: it yields arbitrary millisecond bounds, and
 * sliceSeries rounds those outwards (floor the start, ceil the end), so an off-grid
 * 1000ms window can touch two buckets. Counting both would claim 2000ms of activity
 * inside a 1000ms window, which is impossible on its face. The clamp is a bound on a real
 * measurement, not the old ceiling standing in for one.
 */
function activeMs(series: number[], window: TimeWindow): number {
  const measured = sliceSeries(series, window).filter((value) => value !== 0).length * BUCKET_MS;
  return Math.min(measured, windowMs(window));
}

export function scopeActor(actor: Actor, window: TimeWindow): ScopedActor {
  const total = sumSeries(actor.series, window);
  const whole = actor.series.reduce((sum, value) => sum + value, 0);
  const full = window.startMs <= 0 && window.endMs >= actor.series.length * BUCKET_MS;
  if (full) return { ...actor, approximate: false };

  const ratio = whole === 0 ? 0 : total / whole;
  return {
    ...actor,
    approximate: true,
    total: scale(actor.total, ratio),
    effective: total,
    overheal: actor.overheal === undefined ? undefined : scale(actor.overheal, ratio),
    absorbed: actor.absorbed === undefined ? undefined : scale(actor.absorbed, ratio),
    active_ms: activeMs(actor.series, window),
    series: sliceSeries(actor.series, window),
    abilities: actor.abilities.map((ability) => ({
      ...ability,
      total: scale(ability.total, ratio),
      effective: scale(ability.effective, ratio),
      hits: scale(ability.hits, ratio),
      crits: scale(ability.crits, ratio),
      ticks: scale(ability.ticks, ratio),
      overheal: ability.overheal === undefined ? undefined : scale(ability.overheal, ratio),
      absorbed: ability.absorbed === undefined ? undefined : scale(ability.absorbed, ratio),
    })),
    targets: actor.targets.map((target) => ({
      ...target,
      total: scale(target.total, ratio),
      overheal: target.overheal === undefined ? undefined : scale(target.overheal, ratio),
    })),
  };
}

export function scopeAuraTrack(track: AuraTrack, window: TimeWindow): AuraTrack {
  const segments = track.segments
    .map((segment) => ({
      ...segment,
      start_ms: Math.max(segment.start_ms, window.startMs),
      end_ms: Math.min(segment.end_ms, window.endMs),
    }))
    .filter((segment) => segment.end_ms > segment.start_ms);
  return {
    ...track,
    segments,
    applications: track.segments.filter(
      (segment) => segment.start_ms >= window.startMs && segment.start_ms < window.endMs,
    ).length,
    uptime_ms: segments.reduce((total, segment) => total + (segment.end_ms - segment.start_ms), 0),
  };
}

/** Null when the caster did nothing inside the window, so the row disappears. */
export function scopeCastRow(row: CastRow, window: TimeWindow): CastRow | null {
  const sequence = row.sequence.filter((at) => at >= window.startMs && at < window.endMs);
  if (sequence.length === 0) return null;
  const ratio = row.sequence.length === 0 ? 0 : sequence.length / row.sequence.length;
  return {
    ...row,
    sequence,
    succeeded: Math.round(row.succeeded * ratio),
    started: Math.round(row.started * ratio),
    failed: Math.round(row.failed * ratio),
  };
}

/**
 * The series is exact under a window -- this only slices it -- so the time at the cap is
 * recomputed from the window's own buckets and stays a measured figure. `wasted`,
 * `gained`, `spent` and `zero_ms` are not sliceable from the summary and come through
 * untouched, which is the whole fight's figure and is marked as one where it is printed.
 */
export function scopeResource(track: ResourceTrack, window: TimeWindow): ResourceTrack {
  const series = sliceSeries(track.series, window);
  if (track.max === undefined) return { ...track, series };
  return { ...track, series, at_max_ms: atMaxMs(series, track.max) };
}

/** The buckets whose reading is at the cap, as milliseconds: the engine's own rule. */
function atMaxMs(series: number[], max: number): number {
  if (max <= 0) return 0;
  return series.filter((value) => value >= max).length * BUCKET_MS;
}

function scopeRoster(roster: RosterRow[], scoped: Summary, window: TimeWindow): RosterRow[] {
  const seconds = Math.max(windowMs(window) / 1000, 0.001);
  const find = (rows: Actor[], guid: string): number => rows.find((row) => row.guid === guid)?.effective ?? 0;
  return roster.map((row) => {
    const damage = find(scoped.damage_done, row.guid);
    const healing = find(scoped.healing, row.guid);
    const taken = find(scoped.damage_taken, row.guid);
    return {
      ...row,
      damage_done: damage,
      healing_done: healing,
      damage_taken: taken,
      deaths: scoped.deaths.filter((death) => death.guid === row.guid).length,
      dps: damage / seconds,
      hps: healing / seconds,
      dtps: taken / seconds,
    };
  });
}

function scopeThreat(threat: ThreatRow[], scoped: Summary, whole: Summary): ThreatRow[] {
  // Threat is accumulated from damage and healing, so it scales with the same ratio the
  // actor rows did rather than being recomputed from a model the browser does not have.
  return threat.map((row) => {
    const before =
      (whole.damage_done.find((a) => a.guid === row.guid)?.effective ?? 0) +
      (whole.healing.find((a) => a.guid === row.guid)?.effective ?? 0);
    const after =
      (scoped.damage_done.find((a) => a.guid === row.guid)?.effective ?? 0) +
      (scoped.healing.find((a) => a.guid === row.guid)?.effective ?? 0);
    const ratio = before === 0 ? 0 : after / before;
    return { ...row, threat: row.threat * ratio };
  });
}

/**
 * A player's per-target split scales with the same ratio their own threat total just
 * did -- the browser has no model to recompute the split from, only the whole-fight
 * numbers and the window's share of them.
 */
function scaleThreatPairs(pairs: ThreatPair[], whole: ThreatRow[], scoped: ThreatRow[]): ThreatPair[] {
  return pairs.map((pair) => {
    const before = whole.find((row) => row.guid === pair.guid)?.threat ?? 0;
    const after = scoped.find((row) => row.guid === pair.guid)?.threat ?? 0;
    const ratio = before === 0 ? 0 : after / before;
    return { ...pair, threat: pair.threat * ratio };
  });
}

/**
 * A pair's two honest window figures, both measured from its own per-second series:
 * `standing`, the cumulative threat at the window's end, which is the number that decides
 * who the enemy is looking at, and `built`, the threat made inside the window. `threat`
 * becomes standing, because that is what the table sorts by, draws a bar for and takes a
 * share of. A pair written before engine 0.4.0 has no series and comes back untouched, for
 * `scaleThreatPairs` to prorate the old way.
 */
export function scopeThreatPairs(pairs: ThreatPair[], window: TimeWindow): ThreatPair[] {
  return pairs.map((pair) => {
    if (pair.series === undefined) return pair;
    const standing = sumSeries(pair.series, { startMs: 0, endMs: window.endMs });
    const built = sumSeries(pair.series, window);
    return {
      ...pair,
      threat: standing,
      standing,
      built,
      measured: true,
      series: sliceSeries(pair.series, window),
    };
  });
}

/** A row with nothing in the window is dropped, not shown as a zero. */
function scopeActors(actors: Actor[], window: TimeWindow): ScopedActor[] {
  return actors
    .map((actor) => scopeActor(actor, window))
    .filter((actor) => actor.effective > 0 || actor.total > 0)
    .sort((a, b) => b.effective - a.effective || a.guid.localeCompare(b.guid));
}

/**
 * With no window set this returns the caller's own `summary` by reference, not a copy:
 * every consumer must treat the result as read-only.
 */
export function scopeSummary(summary: Summary, window: TimeWindow): Summary {
  if (isFullWindow(window, summary.duration_ms)) return summary;

  const scoped: Summary = {
    ...summary,
    duration_ms: windowMs(window),
    damage_done: scopeActors(summary.damage_done, window),
    damage_taken: scopeActors(summary.damage_taken, window),
    healing: scopeActors(summary.healing, window),
    healing_taken: scopeActors(summary.healing_taken, window),
    deaths: summary.deaths.filter(
      (death: Death) => death.at_ms >= window.startMs && death.at_ms <= window.endMs,
    ),
    auras: summary.auras.map((track) => scopeAuraTrack(track, window)).filter((track) => track.uptime_ms > 0),
    casts: summary.casts
      .map((row) => scopeCastRow(row, window))
      .filter((row): row is CastRow => row !== null),
    resources: summary.resources.map((track) => scopeResource(track, window)),
  };
  scoped.threat = scopeThreat(summary.threat, scoped, summary);
  // Undefined survives scoping, the same way `taunts` does below: a summary written before
  // the engine kept the per-target split has no key at all, and the table reads that
  // differently from a split that came out empty. A split that carries per-second series
  // (engine 0.4.0) is measured rather than scaled, so the window's figures are the
  // window's own events; one that does not is prorated as before, and marked.
  scoped.threat_by_target =
    summary.threat_by_target === undefined
      ? undefined
      : summary.threat_by_target.every((pair) => pair.series !== undefined)
        ? scopeThreatPairs(summary.threat_by_target, window)
        : scaleThreatPairs(summary.threat_by_target, summary.threat, scoped.threat);
  // Undefined survives scoping: a summary written before the engine kept taunts has no
  // `taunts` key at all, and "none in this window" is a different sentence from "this
  // report never had them". An empty array from an engine that does keep them stays [].
  scoped.taunts = summary.taunts?.filter(
    (taunt) => taunt.at_ms >= window.startMs && taunt.at_ms <= window.endMs,
  );
  scoped.roster = scopeRoster(summary.roster, scoped, window);
  return scoped;
}

/** The twenty seconds before a death, which is what the deaths view opens. */
export function deathWindow(atMs: number, durationMs: number, leadMs = 20_000): TimeWindow {
  return clampWindow({ startMs: atMs - leadMs, endMs: atMs }, durationMs);
}

/**
 * The span around a moment, which is what a taunt's window link opens: a taunt is read
 * from what led to it and what followed, so it sits in the middle rather than at the end
 * the way a death does.
 */
export function aroundWindow(atMs: number, durationMs: number, spanMs = 10_000): TimeWindow {
  return clampWindow({ startMs: atMs - spanMs / 2, endMs: atMs + spanMs / 2 }, durationMs);
}

export interface WindowPreset {
  label: string;
  /** Null is the whole fight, which the URL carries as no start and no end. */
  window: TimeWindow | null;
}

/**
 * The window chips under the chart: the whole fight, the two fixed slices, each phase the
 * encounter's curated table named, and one per death.
 *
 * `wholeDurationMs` is the fight's own length, which is not `summary.duration_ms` when a
 * window is already set (scopeSummary narrows that to the window). The phases are the
 * fight's own spans and must be clamped against the fight, or a chip set inside one
 * window could not reach outside it.
 */
export function windowPresets(summary: Summary, wholeDurationMs = summary.duration_ms): WindowPreset[] {
  const duration = summary.duration_ms;
  const presets: WindowPreset[] = [
    { label: 'Whole fight', window: null },
    { label: 'First 30s', window: clampWindow({ startMs: 0, endMs: 30_000 }, duration) },
    { label: 'Last 30s', window: clampWindow({ startMs: duration - 30_000, endMs: duration }, duration) },
  ];
  for (const phase of summary.phases ?? []) {
    presets.push({
      label: `${phase.name} · ${formatDuration(phase.start_ms)} to ${formatDuration(phase.end_ms)}`,
      window: clampWindow({ startMs: phase.start_ms, endMs: phase.end_ms }, wholeDurationMs),
    });
  }
  for (const death of summary.deaths) {
    presets.push({
      // The time is part of the label: one player can die twice in a fight (a battle
      // rez), and two chips reading "Before Thalgrit died" leave the reader guessing.
      label: `20s before ${splitUnitName(death.name).name} died · ${formatDuration(death.at_ms)}`,
      window: deathWindow(death.at_ms, duration),
    });
  }
  return presets;
}
