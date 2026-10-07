// web/src/lib/bis/heal-view.test.ts
import { describe, expect, it } from 'vitest';
import { clockLabel, healerSetViewFor, rateUnitOf } from './heal-view';
import type { BisBand, BisHealMetrics, BisHealProfile } from './types';

const PROFILE: BisHealProfile = {
  id: 'onyxia-sized',
  label: 'Onyxia-sized tank hits and raid pulses',
  summary: "One healer's share of a Phase 1 boss fight.",
  notes: 'A stated assumption, not a measurement.',
  duration_sec: 300,
  damage_spread: 0.25,
  tank: { health: 9500, hit_damage: 1150, swing_seconds: 2, reason: 'Tank reason.' },
  members: { health: 5000, reason: 'Member reason.' },
  pulse: { damage: 450, interval_seconds: 4, members: 3, reason: 'Pulse reason.' },
  sources: [],
};

const METRICS: BisHealMetrics = {
  hps: 386.74,
  raw_hps: 489.5,
  overheal_pct: 0.214,
  mana_lasts_sec: 3600,
  hpm: 3.84,
};

function healerBand(overrides: Partial<BisBand> = {}): BisBand {
  return {
    spec: 'priest-holy',
    band: 60,
    faction: 'alliance',
    race: 'human',
    talents: '',
    talent_points: 51,
    weights: [],
    slots: [],
    set_dps: 386.74,
    no_source_count: 0,
    new_at_band: [],
    coverage: {},
    weights_run_seconds: 0,
    verify_run_seconds: 0,
    role: 'healer',
    profile: 'onyxia-sized',
    metrics: METRICS,
    ...overrides,
  };
}

describe('rateUnitOf', () => {
  it('reads HPS for a healer band and DPS for anything else', () => {
    expect(rateUnitOf(healerBand())).toBe('HPS');
    expect(rateUnitOf(healerBand({ role: undefined }))).toBe('DPS');
  });
});

describe('clockLabel', () => {
  it('formats seconds as m:ss', () => {
    expect(clockLabel(192)).toBe('3:12');
    expect(clockLabel(300)).toBe('5:00');
    expect(clockLabel(65.9)).toBe('1:05');
    expect(clockLabel(-4)).toBe('0:00');
  });
});

describe('healerSetViewFor', () => {
  it('is undefined for a band that is not a healer band', () => {
    expect(healerSetViewFor(healerBand({ role: undefined }), PROFILE)).toBeUndefined();
  });

  it('builds the headline figure, caption and metrics under the named profile', () => {
    const view = healerSetViewFor(healerBand(), PROFILE)!;
    expect(view.figure).toBe('386.7');
    expect(view.unit).toBe('HPS');
    expect(view.profileLabel).toBe('Onyxia-sized tank hits and raid pulses');
    expect(view.caption).toBe('Effective healing per second under Onyxia-sized tank hits and raid pulses');
    expect(view.metrics.map((m) => [m.key, m.value, m.line])).toEqual([
      ['overheal', '21%', '21% overheal'],
      ['mana', 'Whole fight', 'Mana lasts the whole 5:00 fight'],
      ['hpm', '3.8', '3.8 healing per mana'],
    ]);
  });

  it('says when mana runs out, by clock time, when it is gone before the fight ends', () => {
    const view = healerSetViewFor(healerBand({ metrics: { ...METRICS, mana_lasts_sec: 192 } }), PROFILE)!;
    const mana = view.metrics.find((m) => m.key === 'mana')!;
    expect(mana.value).toBe('3:12');
    expect(mana.line).toBe('Out of mana at 3:12');
  });

  it('counts mana lasting exactly the fight length as lasting the whole fight', () => {
    const view = healerSetViewFor(healerBand({ metrics: { ...METRICS, mana_lasts_sec: 300 } }), PROFILE)!;
    expect(view.metrics.find((m) => m.key === 'mana')!.line).toBe('Mana lasts the whole 5:00 fight');
  });

  it('shows different figures for the two presets', () => {
    const raid = healerSetViewFor(healerBand({ preset: 'raid' }), PROFILE)!;
    const bare = healerSetViewFor(
      healerBand({
        preset: 'bare',
        set_dps: 301.4,
        metrics: { hps: 301.4, raw_hps: 358.8, overheal_pct: 0.16, mana_lasts_sec: 192, hpm: 3.3 },
      }),
      PROFILE,
    )!;
    expect(raid.figure).not.toBe(bare.figure);
    expect(raid.metrics.map((m) => m.value)).not.toEqual(bare.metrics.map((m) => m.value));
  });

  it('carries the profile reasons for the disclosure', () => {
    const disclosure = healerSetViewFor(healerBand(), PROFILE)!.disclosure!;
    expect(disclosure.profileSummary).toBe(PROFILE.summary);
    expect(disclosure.notes).toBe(PROFILE.notes);
    expect(disclosure.rows.map((r) => r.heading)).toEqual(['Fight', 'Tank', 'Raid members', 'Raid pulses']);
    expect(disclosure.rows.find((r) => r.heading === 'Tank')).toEqual({
      heading: 'Tank',
      figures: '9,500 health, a 1,150 hit every 2 seconds',
      reason: 'Tank reason.',
    });
    expect(disclosure.rows.find((r) => r.heading === 'Fight')!.figures).toBe('5:00 long');
  });

  it('falls back to set_dps and no metric cells when the band publishes no metrics', () => {
    const view = healerSetViewFor(healerBand({ metrics: undefined, set_dps: 250.04 }), PROFILE)!;
    expect(view.figure).toBe('250.0');
    expect(view.metrics).toEqual([]);
  });

  it('still names a profile when the file has no heal_profile: the band id, else an unstated one', () => {
    const named = healerSetViewFor(healerBand(), undefined)!;
    expect(named.profileLabel).toBe('onyxia-sized');
    expect(named.disclosure).toBeUndefined();
    const unnamed = healerSetViewFor(healerBand({ profile: undefined }), undefined)!;
    expect(unnamed.caption).toBe('Effective healing per second under an unstated incoming-damage profile');
  });

  it('reads the capped mana figure as lasting the fight when there is no profile to give a fight length', () => {
    const view = healerSetViewFor(healerBand(), undefined)!;
    expect(view.metrics.find((m) => m.key === 'mana')!.line).toBe('Mana lasts the whole fight');
    const short = healerSetViewFor(healerBand({ metrics: { ...METRICS, mana_lasts_sec: 120 } }), undefined)!;
    expect(short.metrics.find((m) => m.key === 'mana')!.line).toBe('Out of mana at 2:00');
  });
});
