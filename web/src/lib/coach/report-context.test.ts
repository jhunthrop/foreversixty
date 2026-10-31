import { describe, expect, it } from 'vitest';
import type { FightEntry } from '../report/types';
import { reviewableFight } from './report-context';

const fight = (overrides: Partial<FightEntry>): FightEntry => ({
  index: 3,
  kind: 'encounter',
  name: 'Onyxia',
  kill: true,
  in_progress: false,
  start: '',
  end: '',
  duration_ms: 232_000,
  players: [],
  deaths: 0,
  ...overrides,
});

describe('reviewableFight', () => {
  it('a long boss pull is reviewable', () => {
    const f = fight({});
    expect(reviewableFight({ reportId: 'abc123', fight: f, firstFight: 1 })).toBe(f);
  });
  it('trash, short pulls and the whole night are not', () => {
    expect(
      reviewableFight({ reportId: 'abc123', fight: fight({ kind: 'trash' }), firstFight: 1 }),
    ).toBeNull();
    expect(
      reviewableFight({ reportId: 'abc123', fight: fight({ duration_ms: 29_999 }), firstFight: 1 }),
    ).toBeNull();
    expect(reviewableFight({ reportId: 'abc123', fight: null, firstFight: 1 })).toBeNull();
  });
});
