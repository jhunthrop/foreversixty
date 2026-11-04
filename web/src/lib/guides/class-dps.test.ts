// web/src/lib/guides/class-dps.test.ts
import { describe, expect, it } from 'vitest';
import { bandEntry, loadBisFile } from '../bis/load';
import { band60Weights, factionForFirstRace } from './class-dps';

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
    expect(result!.haste).toEqual({
      scaleFactor: expected.haste_scale_factor ?? null,
      onItems: expected.haste_on_items ?? true,
    });
    expect(result!.talents).toBe(expected.talents);
    expect(result!.generatedAt).toBe(file.generated_at);
    expect(result!.hitToCap).toEqual(expected.hit_to_cap ?? null);
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
