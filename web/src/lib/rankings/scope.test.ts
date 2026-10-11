// web/src/lib/rankings/scope.test.ts
import { describe, expect, it } from 'vitest';
import { PHASES } from './phases';
import { cohortName, filterSummary, scopeHeading } from './scope';

const PHASE_NAMES = new Map(PHASES.map((phase) => [phase.id, phase.label]));
const BASE = { ruleset: 'normal', region: '', phase: 'raids-1', faction: '', since: '', metric: 'dps' };

describe('the board scope', () => {
  it('names the cohort and heads the board with its plural', () => {
    expect(cohortName('warrior', 'fury')).toBe('Fury Warrior');
    expect(cohortName('warrior', '')).toBe('Warrior');
    expect(cohortName('', '')).toBe('');
    expect(scopeHeading('warrior', 'Fury')).toBe('Fury Warriors');
    expect(scopeHeading('priest', '')).toBe('Priests');
    expect(scopeHeading('', '')).toBe('Every class and spec');
  });
});

describe('the phone summary line', () => {
  it('reads encounter, ruleset, region and phase', () => {
    expect(filterSummary({ encounterName: 'Onyxia', state: BASE, defaultMetric: 'dps' }, PHASE_NAMES)).toBe(
      'Onyxia · Normal · All regions · Phase 1',
    );
  });

  it('names an unset filter as all, and a chosen region in capitals', () => {
    const state = { ...BASE, ruleset: '', region: 'eu', phase: '' };
    expect(filterSummary({ encounterName: 'Onyxia', state, defaultMetric: 'dps' }, PHASE_NAMES)).toBe(
      'Onyxia · All rulesets · EU · All phases',
    );
  });

  it('adds only the narrowing choices that were made', () => {
    const state = { ...BASE, faction: 'horde', since: 'today', metric: 'hps' };
    expect(filterSummary({ encounterName: 'Onyxia', state, defaultMetric: 'dps' }, PHASE_NAMES)).toBe(
      'Onyxia · Normal · All regions · Phase 1 · Horde · Today only · Healing',
    );
    expect(
      filterSummary({ encounterName: 'Onyxia', state, defaultMetric: 'hps' }, PHASE_NAMES),
    ).not.toContain('Healing');
  });
});
