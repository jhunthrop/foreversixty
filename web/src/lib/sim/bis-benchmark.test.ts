import { describe, expect, it } from 'vitest';
import fixture from '../../fixtures/sim/result-full.json';
import { benchmarkFor, type BisBenchmarks } from './bis-benchmark';
import { buildBisBenchmarks, specBenchmarks } from './bis-benchmark-load';
import type { BisFile, SpecCatalogEntry } from '../bis/types';
import type { SimResult } from './types';

const RESULT = fixture as unknown as SimResult;
const RUN = { duration_sec: 180, variation: 0.2, targets: 1, execute_ratio: 0.25, target_level_offset: 3 };
const TABLE: BisBenchmarks = {
  'warrior-fury': {
    run: RUN,
    bands: [
      [60, 'horde', 'raid', 1004.02, 1.06],
      [60, 'horde', 'bare', 362.84, 0.4],
      [60, 'alliance', 'raid', 1004.21, 0.96],
    ],
  },
};

function withRequest(patch: (result: SimResult) => SimResult): SimResult {
  return patch(structuredClone(RESULT));
}

describe('benchmarkFor', () => {
  it('reads the fixture run (Troll, 60, raid, the ranker fight) against the Horde raid set', () => {
    const match = benchmarkFor(RESULT, TABLE);
    expect(match).not.toBeNull();
    expect(match?.faction).toBe('horde');
    expect(match?.preset).toBe('raid');
    expect(match?.dps).toBe(1004.02);
    expect(match?.gap).toBeCloseTo(RESULT.dps.mean - 1004.02, 6);
    expect(match?.margin).toBeCloseTo(1.96 * Math.hypot(RESULT.dps.error, 1.06), 6);
    expect(match?.withinNoise).toBe(false);
  });

  it('says nothing for a fight the set was not ranked in', () => {
    const shorter = withRequest((r) => ({
      ...r,
      request: { ...r.request, encounter: { ...r.request.encounter, duration_sec: 120 } },
    }));
    expect(benchmarkFor(shorter, TABLE)).toBeNull();
    const cleave = withRequest((r) => ({
      ...r,
      request: { ...r.request, encounter: { ...r.request.encounter, targets: 3 } },
    }));
    expect(benchmarkFor(cleave, TABLE)).toBeNull();
  });

  it('says nothing for a level with no band, an unknown spec or a raid run with no buffs', () => {
    const level59 = withRequest((r) => ({
      ...r,
      request: { ...r.request, character: { ...r.request.character, level: 59 } },
    }));
    expect(benchmarkFor(level59, TABLE)).toBeNull();
    expect(benchmarkFor(RESULT, {})).toBeNull();
    const unbuffed = withRequest((r) => ({
      ...r,
      request: { ...r.request, character: { ...r.request.character, buffs: [] } },
    }));
    expect(benchmarkFor(unbuffed, TABLE)).toBeNull();
  });

  it('calls a gap inside both runs’ band within the noise', () => {
    const close = withRequest((r) => ({ ...r, dps: { ...r.dps, mean: 1004.5 } }));
    expect(benchmarkFor(close, TABLE)?.withinNoise).toBe(true);
  });
});

describe('specBenchmarks', () => {
  const band = {
    band: 60,
    faction: 'horde',
    preset: 'raid',
    spec: 'warrior-fury',
    set_dps: 1004.015220120736,
    set_dps_error: 1.0583860418252813,
  };
  const file = {
    spec: 'warrior-fury',
    bands: [band],
    run: { ...RUN, iterations: {}, seed: 7 },
  } as unknown as BisFile;

  it('keeps each damage band to a hundredth with the file’s run block', () => {
    expect(specBenchmarks(file)).toEqual({ run: RUN, bands: [[60, 'horde', 'raid', 1004.02, 1.06]] });
  });

  it('leaves out a file with no run block, and a tank or partial band', () => {
    expect(specBenchmarks({ ...file, run: undefined })).toBeNull();
    const tank = { ...file, bands: [{ ...band, role: 'tank' }] } as unknown as BisFile;
    expect(specBenchmarks(tank)).toBeNull();
    const partial = { ...file, bands: [{ ...band, set_dps_partial: true }] } as unknown as BisFile;
    expect(specBenchmarks(partial)).toBeNull();
  });

  it('builds the table for damage specs only', () => {
    const catalog = [
      { spec: 'warrior-fury', role: 'dps' },
      { spec: 'warrior-protection', role: 'tank' },
    ] as SpecCatalogEntry[];
    expect(Object.keys(buildBisBenchmarks(catalog, () => file))).toEqual(['warrior-fury']);
  });
});
