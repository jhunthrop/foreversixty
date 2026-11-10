// web/src/lib/sim/result-band.test.ts
// Spec 2026-11-04-sim-results 3.3 and section 10, check 3.
import { describe, expect, it } from 'vitest';
import full from '../../fixtures/sim/result-full.json';
import older from '../../fixtures/sim/result.json';
import {
  bandDescriptor,
  bandError,
  bandFacts,
  bandHeading,
  buffsFact,
  savedDate,
  savedDay,
  SPACED_SPLIT,
  talentSplit,
} from './result-band';
import type { SimResult } from './types';

const result = full as unknown as SimResult;
const oldResult = older as unknown as SimResult;
const facts = (from: SimResult) =>
  Object.fromEntries(bandFacts(from, null).map((fact) => [fact.id, fact.value]));

describe('the band', () => {
  it('heads with the spec when the name is the ranker placeholder', () => {
    expect(result.request.character.name).toBe('breakdown');
    expect(bandHeading(result)).toBe('Fury Warrior');
  });

  it('prefers a title, then a real name', () => {
    expect(bandHeading({ ...result, title: 'BWL night' })).toBe('BWL night');
    const named = { ...result.request.character, name: 'Thoradin' };
    expect(bandHeading({ ...result, request: { ...result.request, character: named } })).toBe('Thoradin');
  });

  it('reads the one talent split', () => {
    expect(talentSplit('30323013002000000-051531025120100231-200000000000000000')).toBe('17/32/2');
    expect(talentSplit('30323013002000000-051531025120100231-200000000000000000', SPACED_SPLIT)).toBe(
      '17 / 32 / 2',
    );
    expect(bandDescriptor(result)).toEqual(['Troll Fury Warrior', 'Level 60', '17/32/2 talents']);
  });

  it('states the error as ± 0.5, 0.05%', () => {
    expect(bandError(result)).toEqual({ margin: '± 0.5', percent: '0.05%' });
  });

  it('prints the facts in order with their values', () => {
    expect(bandFacts(result, null).map((fact) => fact.label)).toEqual([
      'Fight',
      'Targets',
      'Buffs',
      'Consumables',
      'Range',
      'Runs',
      'Saved',
    ]);
    expect(facts(result)).toEqual({
      fight: '3:00 ± 20%',
      targets: '1',
      buffs: 'Raid',
      consumables: '8',
      range: '731 to 1,116',
      runs: '10,000',
      saved: '4 Nov 2026',
    });
  });

  it('an older result leaves Saved out rather than using the capture date', () => {
    expect(facts(oldResult).saved).toBeUndefined();
    expect(savedDate(undefined)).toBeNull();
    expect(savedDate('not a date')).toBeNull();
  });

  it('the head says the save day, with its year only when it is not this year', () => {
    const now = new Date('2026-11-09T12:00:00Z');
    expect(savedDay('2026-11-04T10:00:00Z', now)).toBe('4 Nov');
    expect(savedDay('2025-11-04T10:00:00Z', now)).toBe('4 Nov 2025');
    expect(savedDay('', now)).toBeNull();
  });

  it('Buffs reads Raid only when the list equals the preset', () => {
    const { buffs } = result.request.character;
    expect(buffsFact(result, [...buffs].reverse())).toBe('Raid');
    expect(buffsFact(result, buffs.slice(1))).toBe('Custom');
    const none = { ...result.request.character, buffs: [] };
    expect(buffsFact({ ...result, request: { ...result.request, character: none } }, null)).toBe('None');
  });
});
