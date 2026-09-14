import { describe, expect, it } from 'vitest';
import fixtureSummary from '../../fixtures/report/fights/3/summary.json';
import type { Actor, AuraTrack, CastRow, Summary } from './types';
import {
  BUCKET_MS,
  combinedSeries,
  deathWindow,
  fullWindow,
  isFullWindow,
  scopeActor,
  scopeAuraTrack,
  scopeCastRow,
  scopeSummary,
  sliceSeries,
  sumSeries,
  windowMs,
  windowOf,
  windowPresets,
} from './window';

const summary = fixtureSummary as Summary;

const actor: Actor = {
  guid: 'g1',
  name: 'Tester',
  total: 100,
  effective: 100,
  active_ms: 4000,
  abilities: [
    { spell_id: 1, name: 'A', total: 60, effective: 60, hits: 6, crits: 0, ticks: 0, min: 10, max: 10 },
    { spell_id: 2, name: 'B', total: 40, effective: 40, hits: 4, crits: 0, ticks: 0, min: 10, max: 10 },
  ],
  targets: [{ guid: 't1', name: 'Boss', total: 100 }],
  series: [10, 20, 30, 40],
};

describe('the window itself', () => {
  it('buckets at one second, which is what the engine writes', () => {
    expect(BUCKET_MS).toBe(1000);
  });

  it('defaults to the whole fight and recognises it', () => {
    expect(fullWindow(40000)).toEqual({ startMs: 0, endMs: 40000 });
    expect(isFullWindow({ startMs: 0, endMs: 40000 }, 40000)).toBe(true);
    expect(isFullWindow({ startMs: 1000, endMs: 40000 }, 40000)).toBe(false);
    expect(windowMs({ startMs: 1000, endMs: 4000 })).toBe(3000);
  });

  it('reads the window out of the URL state, clamped to the fight', () => {
    const base = { fight: 3, mode: 'analyze', view: 'tables', tab: 'summary', source: 'friendlies' } as const;
    expect(windowOf({ ...base, start: null, end: null }, 40000)).toEqual({ startMs: 0, endMs: 40000 });
    expect(windowOf({ ...base, start: 2000, end: 9000 }, 40000)).toEqual({ startMs: 2000, endMs: 9000 });
    expect(windowOf({ ...base, start: -5, end: 99999 }, 40000)).toEqual({ startMs: 0, endMs: 40000 });
  });
});

describe('series slicing', () => {
  it('takes the buckets the window covers', () => {
    expect(sliceSeries([10, 20, 30, 40], { startMs: 1000, endMs: 3000 })).toEqual([20, 30]);
    expect(sumSeries([10, 20, 30, 40], { startMs: 1000, endMs: 3000 })).toBe(50);
  });

  it('is total over the whole fight, and empty past the end', () => {
    expect(sumSeries([10, 20, 30, 40], { startMs: 0, endMs: 4000 })).toBe(100);
    expect(sliceSeries([10, 20], { startMs: 9000, endMs: 10_000 })).toEqual([]);
  });

  it('sums several actors into one line for the chart', () => {
    expect(
      combinedSeries([
        { ...actor, series: [1, 2] },
        { ...actor, series: [10, 20, 30] },
      ]),
    ).toEqual([11, 22, 30]);
  });
});

describe('scoping a row to the window', () => {
  it('takes an actor’s total exactly from its series and scales the split', () => {
    const scoped = scopeActor(actor, { startMs: 1000, endMs: 3000 });
    expect(scoped.total).toBe(50);
    expect(scoped.effective).toBe(50);
    expect(scoped.series).toEqual([20, 30]);
    // 60/100 and 40/100 of 50.
    expect(scoped.abilities.map((a) => a.total)).toEqual([30, 20]);
    expect(scoped.abilities.map((a) => a.hits)).toEqual([3, 2]);
    expect(scoped.targets[0].total).toBe(50);
    expect(scoped.approximate).toBe(true);
  });

  it('measures active time from the series instead of capping the whole-fight figure', () => {
    // Three seconds of window, one of them with output. The old ceiling
    // min(active_ms, windowMs) called all three active, which is the whole fight's
    // activity read onto an idle stretch.
    const idle: Actor = { ...actor, active_ms: 4000, series: [10, 0, 0, 40] };
    expect(scopeActor(idle, { startMs: 0, endMs: 3000 }).active_ms).toBe(1000);
    expect(scopeActor(idle, { startMs: 1000, endMs: 3000 }).active_ms).toBe(0);
  });

  it('never reports more active time than the window is long', () => {
    // The range inputs step by BUCKET_MS, but a canvas drag does not: onUp produces
    // arbitrary millisecond bounds, gated only by a minimum width of one bucket. A legal
    // 1000ms drag landing off the bucket grid touches two buckets, and two buckets of a
    // fully active series is 2000ms of activity inside a 1000ms window.
    const busy: Actor = { ...actor, series: [10, 20, 30, 40] };
    const offGrid = { startMs: 999, endMs: 1999 };
    expect(scopeActor(busy, offGrid).active_ms).toBeLessThanOrEqual(windowMs(offGrid));
  });

  it('leaves a whole-fight actor exactly as the engine wrote it', () => {
    const scoped = scopeActor(actor, { startMs: 0, endMs: 4000 });
    expect(scoped.total).toBe(100);
    expect(scoped.abilities).toEqual(actor.abilities);
    expect(scoped.approximate).toBe(false);
  });

  it('intersects an aura’s segments exactly', () => {
    const track: AuraTrack = {
      target_guid: 't',
      target_name: 'T',
      spell_id: 1,
      name: 'Buff',
      type: 'BUFF',
      applications: 2,
      max_stacks: 1,
      uptime_ms: 9000,
      segments: [
        { start_ms: 0, end_ms: 4000, stacks: 1 },
        { start_ms: 6000, end_ms: 11_000, stacks: 1 },
      ],
      appliers: ['a'],
    };
    const scoped = scopeAuraTrack(track, { startMs: 3000, endMs: 8000 });
    expect(scoped.segments).toEqual([
      { start_ms: 3000, end_ms: 4000, stacks: 1 },
      { start_ms: 6000, end_ms: 8000, stacks: 1 },
    ]);
    expect(scoped.uptime_ms).toBe(3000);
    expect(scoped.applications).toBe(1);
  });

  it('filters a cast row’s sequence exactly, and drops a row with nothing left', () => {
    const row: CastRow = {
      guid: 'g',
      name: 'N',
      spell_id: 1,
      spell_name: 'S',
      started: 3,
      succeeded: 3,
      failed: 0,
      cast_time_ms: 1500,
      sequence: [1000, 5000, 9000],
    };
    const scoped = scopeCastRow(row, { startMs: 4000, endMs: 6000 });
    expect(scoped?.sequence).toEqual([5000]);
    expect(scoped?.succeeded).toBe(1);
    expect(scopeCastRow(row, { startMs: 20_000, endMs: 21_000 })).toBeNull();
  });

  it('scopes the whole fixture summary and keeps the death inside the window', () => {
    const scoped = scopeSummary(summary, { startMs: 8000, endMs: 12_000 });
    expect(scoped.deaths).toHaveLength(1);
    expect(scoped.deaths[0].name).toBe('Thalgrit-Nightslayer');
    expect(scopeSummary(summary, { startMs: 20_000, endMs: 30_000 }).deaths).toHaveLength(0);
    expect(scoped.duration_ms).toBe(4000);
  });

  it('rescopes the fixture’s damage so a window that excludes a caster excludes their row', () => {
    const early = scopeSummary(summary, { startMs: 0, endMs: 4000 });
    expect(early.damage_done.map((a) => a.name)).not.toContain('Morrowlyn-Nightslayer');
    const whole = scopeSummary(summary, { startMs: 0, endMs: 40_000 });
    expect(whole.damage_done.map((a) => a.name)).toContain('Morrowlyn-Nightslayer');
  });
});

describe('presets', () => {
  it('sets the window to the twenty seconds before a death', () => {
    expect(deathWindow(10_100, 40_000)).toEqual({ startMs: 0, endMs: 10_100 });
    expect(deathWindow(30_000, 40_000)).toEqual({ startMs: 10_000, endMs: 30_000 });
  });

  it('offers the whole fight, both thirty-second ends, and one preset per death', () => {
    const presets = windowPresets(summary);
    expect(presets.map((preset) => preset.label)).toEqual([
      'Whole fight',
      'First 30s',
      'Last 30s',
      'Before Thalgrit died',
    ]);
    expect(presets[0].window).toBeNull();
    expect(presets[3].window).toEqual({ startMs: 0, endMs: 10_100 });
  });
});
