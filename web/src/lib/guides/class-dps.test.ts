// web/src/lib/guides/class-dps.test.ts
import { describe, expect, it } from 'vitest';
import { bandEntry, loadBisFile } from '../bis/load';
import { isTankMetrics } from '../bis/tank-view';
import { band60Weights, classLandingSetMetric, factionForFirstRace } from './class-dps';

const BUILD = '1.60.1.70374';

describe('factionForFirstRace', () => {
  it('resolves the faction of the guide’s own first recommended race', () => {
    expect(factionForFirstRace('warrior', ['human', 'troll'])).toBe('alliance');
    expect(factionForFirstRace('warrior', ['troll', 'human'])).toBe('horde');
  });

  it('is undefined when the guide names no recommended race', () => {
    expect(factionForFirstRace('warrior', [])).toBeUndefined();
  });

  it('is undefined for a race the class cannot legally play', () => {
    expect(factionForFirstRace('warrior', ['not-a-real-race'])).toBeUndefined();
  });
});

describe('classLandingSetMetric', () => {
  it('reads Fury’s band-60 set DPS at the faction its first recommended race resolves to', () => {
    // Read straight from the committed file rather than pinned: the nightly republishes it
    // (data contracts are floors, never exact values).
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const expected = bandEntry(file, 60, 'alliance')?.set_dps;
    const metric = classLandingSetMetric(BUILD, 'warrior-fury', 'warrior', ['human', 'troll']);
    expect(metric?.role).toBe('dps');
    expect(metric?.value).toBeCloseTo(expected!, 5);
    expect(metric!.value).toBeGreaterThan(0);
  });

  it('reads the Horde row when the first recommended race is a Horde one', () => {
    const alliance = classLandingSetMetric(BUILD, 'warrior-fury', 'warrior', ['human', 'troll']);
    const horde = classLandingSetMetric(BUILD, 'warrior-fury', 'warrior', ['troll', 'human']);
    expect(horde?.value).not.toBe(alliance?.value);
  });

  it('gives a healer its healing per second, never a DPS figure', () => {
    const file = loadBisFile('paladin-holy', BUILD);
    if (file === null) throw new Error('paladin-holy BiS file missing');
    const band = bandEntry(file, 60, 'alliance');
    const metric = classLandingSetMetric(BUILD, 'paladin-holy', 'paladin', ['human']);
    expect(metric?.role).toBe('healer');
    const hps =
      band?.metrics !== null && band?.metrics !== undefined && 'hps' in band.metrics
        ? band.metrics.hps
        : band?.set_dps;
    expect(metric?.value).toBeCloseTo(hps!, 5);
  });

  it('gives a tank its damage taken per second', () => {
    const file = loadBisFile('warrior-protection', BUILD);
    if (file === null) throw new Error('warrior-protection BiS file missing');
    const band = bandEntry(file, 60, 'alliance');
    if (band === undefined || !isTankMetrics(band.metrics)) throw new Error('tank metrics missing');
    const metric = classLandingSetMetric(BUILD, 'warrior-protection', 'warrior', ['human']);
    expect(metric).toEqual({ role: 'tank', value: band.metrics.dtps });
  });

  it('is undefined for a spec with no ranked BiS file', () => {
    expect(classLandingSetMetric(BUILD, 'warrior-nonexistent', 'warrior', ['dwarf'])).toBeUndefined();
  });
});

describe('band60Weights', () => {
  it('reads Fury’s band-60 weights and file generated_at at the faction its first race resolves to', () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const expected = bandEntry(file, 60, 'alliance');
    if (expected === undefined) throw new Error('band 60 alliance missing');

    const result = band60Weights(BUILD, 'warrior-fury', 'warrior', ['human', 'troll']);
    expect(result).toBeDefined();
    expect(result!.weights).toEqual(expected.weights);
    expect(result!.role).toBe('dps');
    expect(result!.hasteScaleFactor).toBe(expected.haste_scale_factor ?? null);
    expect(result!.generatedAt).toBe(file.generated_at);
    expect(result!.hitToCap).toEqual(expected.hit_to_cap ?? null);
    expect(result!.expertiseToCap).toEqual(expected.expertise_to_cap ?? null);
  });

  it('is undefined for a spec with no ranked BiS file', () => {
    expect(band60Weights(BUILD, 'warrior-nonexistent', 'warrior', ['dwarf'])).toBeUndefined();
  });

  it('carries the band role, which decides the per-point unit', () => {
    expect(band60Weights(BUILD, 'warrior-protection', 'warrior', ['human'])?.role).toBe('tank');
    expect(band60Weights(BUILD, 'priest-holy', 'priest', ['human'])?.role).toBe('healer');
  });

  it('is undefined when the guide names no recommended race', () => {
    expect(band60Weights(BUILD, 'warrior-fury', 'warrior', [])).toBeUndefined();
  });
});
