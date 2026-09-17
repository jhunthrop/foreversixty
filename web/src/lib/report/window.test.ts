import { describe, expect, it } from 'vitest';
import fixtureSummary from '../../fixtures/report/fights/3/summary.json';
import type { Actor, AuraTrack, CastRow, Summary, ThreatPair } from './types';
import {
  BUCKET_MS,
  aroundWindow,
  combinedSeries,
  deathWindow,
  fullWindow,
  isFullWindow,
  scopeActor,
  scopeAuraTrack,
  scopeCastRow,
  scopeResource,
  scopeSummary,
  scopeThreatPairs,
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

  // One player, one enemy, four seconds: the smallest fight that has a threat split to scale.
  const base: Summary = {
    engine_version: 't',
    fight_index: 1,
    duration_ms: 4000,
    damage_done: [
      {
        guid: 'P1',
        name: 'Tank',
        total: 800,
        effective: 800,
        active_ms: 4000,
        abilities: [],
        targets: [],
        series: [200, 200, 200, 200],
      },
    ],
    damage_taken: [],
    healing: [],
    healing_taken: [],
    deaths: [],
    auras: [],
    casts: [],
    interrupts: [],
    dispels: [],
    resources: [],
    threat: [{ guid: 'P1', name: 'Tank', threat: 1000, model_version: 'base-1', complete: false }],
    combatants: [],
    roster: [],
  };

  it('scales a player’s per-target threat by the same ratio as their total', () => {
    const halfWindow = { startMs: 0, endMs: 2000 };
    const withPairs: Summary = {
      ...base,
      threat_by_target: [
        { guid: 'P1', name: 'Tank', target_guid: 'E1', target_name: 'Boss', threat: 600 },
        { guid: 'P1', name: 'Tank', target_guid: 'E2', target_name: 'Add', threat: 400 },
      ],
      taunts: [
        {
          at_ms: 500,
          source_guid: 'P1',
          source_name: 'Tank',
          target_guid: 'E1',
          target_name: 'Boss',
          spell_id: 355,
          spell_name: 'Taunt',
        },
        {
          at_ms: 3000,
          source_guid: 'P1',
          source_name: 'Tank',
          target_guid: 'E1',
          target_name: 'Boss',
          spell_id: 355,
          spell_name: 'Taunt',
        },
      ],
    };
    const scoped = scopeSummary(withPairs, halfWindow);
    const ratio = scoped.threat[0].threat / 1000;
    expect(scoped.threat_by_target?.[0].threat).toBeCloseTo(600 * ratio);
    expect(scoped.threat_by_target?.[1].threat).toBeCloseTo(400 * ratio);
    // The taunt inside the window survives; the one after it is cut, like a death would be.
    expect(scoped.taunts?.map((taunt) => taunt.at_ms)).toEqual([500]);
  });

  it('leaves a player with no threat at zero rather than dividing by it', () => {
    // A pair whose player holds no threat over the whole fight has no ratio to scale by:
    // the guard hands back zero instead of a NaN that would render as "~NaN".
    const orphan: Summary = {
      ...base,
      threat: [{ guid: 'P1', name: 'Tank', threat: 0, model_version: 'base-1', complete: false }],
      threat_by_target: [{ guid: 'P1', name: 'Tank', target_guid: 'E1', target_name: 'Boss', threat: 600 }],
    };
    const scoped = scopeSummary(orphan, { startMs: 0, endMs: 2000 });
    expect(scoped.threat_by_target?.[0].threat).toBe(0);
  });

  it('leaves the per-target split undefined when the summary never had the key', () => {
    // The same sentence taunts get below: a report parsed before the engine split threat
    // per target has no key, which the table reads differently from an empty split.
    const whole = fullWindow(base.duration_ms);
    const half = { startMs: 0, endMs: 2000 };
    expect(scopeSummary({ ...base, threat_by_target: undefined }, half).threat_by_target).toBeUndefined();
    expect(scopeSummary({ ...base, threat_by_target: [] }, half).threat_by_target).toEqual([]);
    expect(scopeSummary({ ...base, threat_by_target: undefined }, whole).threat_by_target).toBeUndefined();
  });

  it('leaves taunts undefined when the summary never had the key', () => {
    // A summary the engine wrote before it kept taunts: "none in this window" and "this
    // report has none to keep" are different sentences, so the absence has to survive.
    const whole = fullWindow(summary.duration_ms);
    expect(scopeSummary({ ...summary, taunts: undefined }, whole).taunts).toBeUndefined();
    expect(scopeSummary({ ...summary, taunts: [] }, whole).taunts).toEqual([]);
  });
});

describe('presets', () => {
  it('sets the window to the twenty seconds before a death', () => {
    expect(deathWindow(10_100, 40_000)).toEqual({ startMs: 0, endMs: 11_000 });
    expect(deathWindow(30_000, 40_000)).toEqual({ startMs: 10_000, endMs: 30_000 });
  });

  it('sets the window to the span around a taunt, snapped out and clamped', () => {
    expect(aroundWindow(8_500, 60_000)).toEqual({ startMs: 3_000, endMs: 14_000 });
    expect(aroundWindow(1_000, 60_000)).toEqual({ startMs: 0, endMs: 6_000 });
    expect(aroundWindow(59_000, 60_000)).toEqual({ startMs: 54_000, endMs: 60_000 });
  });

  it('offers the whole fight, both thirty-second ends, one per phase, and one per death', () => {
    const presets = windowPresets(summary);
    expect(presets.map((preset) => preset.label)).toEqual([
      'Whole fight',
      'First 30s',
      'Last 30s',
      'Phase 1 · 0.0s to 14.0s',
      'Phase 2 · 14.0s to 1:00',
      '20s before Thalgrit died · 10.1s',
    ]);
    expect(presets[0].window).toBeNull();
    expect(presets[5].window).toEqual({ startMs: 0, endMs: 11_000 });
  });

  it('offers each phase as a preset, in order, after the fixed ones', () => {
    const phased: Summary = {
      ...summary,
      phases: [
        { name: 'Phase 1', start_ms: 0, end_ms: 14_000 },
        { name: 'Phase 2', start_ms: 14_000, end_ms: 60_000 },
      ],
    };
    const presets = windowPresets(phased);
    const phases = presets.filter((preset) => preset.label.startsWith('Phase'));
    expect(phases.map((preset) => preset.label)).toEqual([
      'Phase 1 · 0.0s to 14.0s',
      'Phase 2 · 14.0s to 1:00',
    ]);
    expect(phases[1].window).toEqual({ startMs: 14_000, endMs: 60_000 });
  });

  it('clamps a phase against the whole fight, not against a window already set', () => {
    const scopedToAWindow: Summary = {
      ...summary,
      duration_ms: 5000,
      phases: [{ name: 'Phase 2', start_ms: 14_000, end_ms: 60_000 }],
    };
    const [phase] = windowPresets(scopedToAWindow, 60_000).filter((preset) =>
      preset.label.startsWith('Phase'),
    );
    expect(phase.window).toEqual({ startMs: 14_000, endMs: 60_000 });
  });

  it('offers no phase presets for a fight that has none', () => {
    const unphased: Summary = { ...summary, phases: undefined };
    expect(windowPresets(unphased).some((preset) => preset.label.startsWith('Phase'))).toBe(false);
  });

  it('scales an ability’s overheal with its total, so the share holds inside a window', () => {
    const actor: Actor = {
      guid: 'H',
      name: 'Healer',
      total: 1000,
      effective: 600,
      overheal: 400,
      active_ms: 4000,
      abilities: [
        {
          spell_id: 1,
          name: 'Heal',
          total: 1000,
          effective: 600,
          overheal: 400,
          hits: 4,
          crits: 0,
          ticks: 0,
          min: 0,
          max: 0,
        },
      ],
      targets: [],
      series: [150, 150, 150, 150],
    };
    const scoped = scopeActor(actor, { startMs: 0, endMs: 2000 });
    expect(scoped.abilities[0].overheal).toBe(200);
    expect(scoped.abilities[0].total).toBe(500);
  });
});

describe('threat inside a window', () => {
  const pair = (guid: string, threat: number, series: number[]): ThreatPair => ({
    guid,
    name: guid,
    target_guid: 'Creature-1',
    target_name: 'Boss',
    threat,
    series,
  });

  it('measures standing at the window’s end and built inside it, from the pair’s own series', () => {
    const [scoped] = scopeThreatPairs([pair('P1', 100, [10, 20, 30, 40])], {
      startMs: 1000,
      endMs: 3000,
    });
    // Standing is everything up to the window's end: 10 + 20 + 30.
    expect(scoped.standing).toBe(60);
    // Built is the window's own buckets: 20 + 30.
    expect(scoped.built).toBe(50);
    expect(scoped.measured).toBe(true);
    // The table sorts and draws on standing, so `threat` is standing.
    expect(scoped.threat).toBe(60);
    expect(scoped.series).toEqual([20, 30]);
  });

  it('leaves a whole-fight window alone: standing at the end is the total', () => {
    const [scoped] = scopeThreatPairs([pair('P1', 100, [10, 20, 30, 40])], {
      startMs: 0,
      endMs: 4000,
    });
    expect(scoped.standing).toBe(100);
    expect(scoped.built).toBe(100);
    expect(scoped.threat).toBe(100);
  });

  it('says nothing was measured when the pair was written before the engine kept a series', () => {
    const old: ThreatPair = {
      guid: 'P1',
      name: 'P1',
      target_guid: 'Creature-1',
      target_name: 'Boss',
      threat: 100,
    };
    const [scoped] = scopeThreatPairs([old], { startMs: 1000, endMs: 3000 });
    expect(scoped.measured).toBeUndefined();
    expect(scoped.standing).toBeUndefined();
    expect(scoped.threat).toBe(100);
  });

  it('scopeSummary measures the pairs when they carry a series and scales them when they do not', () => {
    const base = summary;
    const withSeries: Summary = {
      ...base,
      threat: [{ guid: 'P1', name: 'P1', threat: 100, model_version: 'base-1', complete: false }],
      threat_by_target: [pair('P1', 100, [10, 20, 30, 40])],
      duration_ms: 4000,
      damage_done: [],
      healing: [],
    };
    const scoped = scopeSummary(withSeries, { startMs: 1000, endMs: 3000 });
    expect(scoped.threat_by_target?.[0].measured).toBe(true);
    expect(scoped.threat_by_target?.[0].standing).toBe(60);

    const withoutSeries: Summary = {
      ...withSeries,
      threat_by_target: [{ ...pair('P1', 100, []), series: undefined }],
    };
    const legacy = scopeSummary(withoutSeries, { startMs: 1000, endMs: 3000 });
    // No series, no damage and no healing in the window: the old ratio scales it to zero.
    expect(legacy.threat_by_target?.[0].measured).toBeUndefined();
    expect(legacy.threat_by_target?.[0].threat).toBe(0);
  });
});

describe('a resource track in a window', () => {
  const track = {
    guid: 'P1',
    name: 'P1',
    power_type: 1,
    series: [0, 100, 100, 40, 100],
    gained: 300,
    spent: 60,
    zero_ms: 1000,
    max: 100,
    at_max_ms: 3000,
    wasted: 25,
  };

  it('recomputes the time at the cap from the window’s own buckets', () => {
    const scoped = scopeResource(track, { startMs: 1000, endMs: 4000 });
    expect(scoped.series).toEqual([100, 100, 40]);
    expect(scoped.at_max_ms).toBe(2000);
  });

  it('leaves the cap and the waste alone: one is a property of the bar, the other whole-fight', () => {
    const scoped = scopeResource(track, { startMs: 1000, endMs: 4000 });
    expect(scoped.max).toBe(100);
    expect(scoped.wasted).toBe(25);
    expect(scoped.zero_ms).toBe(1000);
  });

  it('leaves a track written before the engine kept a cap untouched', () => {
    const { max: _max, at_max_ms: _atMax, wasted: _wasted, ...old } = track;
    const scoped = scopeResource(old, { startMs: 1000, endMs: 4000 });
    expect(scoped.max).toBeUndefined();
    expect(scoped.at_max_ms).toBeUndefined();
  });
});
