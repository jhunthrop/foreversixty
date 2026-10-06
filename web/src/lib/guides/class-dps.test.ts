// web/src/lib/guides/class-dps.test.ts
import { describe, expect, it } from 'vitest';
import { bandEntry, loadBisFile } from '../bis/load';
import { classLandingSetDps, factionForFirstRace } from './class-dps';

const BUILD = '1.60.1.70009';

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

describe('classLandingSetDps', () => {
  it('reads Fury’s band-60 set DPS at the faction its first recommended race resolves to', () => {
    // Read straight from the committed file rather than pinned: the nightly republishes it
    // (data contracts are floors, never exact values).
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const expected = bandEntry(file, 60, 'alliance')?.set_dps;
    const dps = classLandingSetDps(BUILD, 'warrior-fury', 'warrior', ['human', 'troll']);
    expect(dps).toBeDefined();
    expect(dps).toBeCloseTo(expected!, 5);
    expect(dps).toBeGreaterThan(0);
  });

  it('reads the Horde row when the first recommended race is a Horde one', () => {
    const allianceDps = classLandingSetDps(BUILD, 'warrior-fury', 'warrior', ['human', 'troll']);
    const hordeDps = classLandingSetDps(BUILD, 'warrior-fury', 'warrior', ['troll', 'human']);
    expect(hordeDps).not.toBe(allianceDps);
  });

  it('is undefined for a spec with no ranked BiS file (Protection)', () => {
    expect(classLandingSetDps(BUILD, 'warrior-protection', 'warrior', ['dwarf'])).toBeUndefined();
  });
});
