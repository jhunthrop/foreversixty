// web/src/lib/rating/roster-join.test.ts
import { describe, expect, it } from 'vitest';
import { classForPlayer, guidForPlayer } from './roster-join';
import type { RatingCardPlayer } from './types';

const ROSTER = [{ guid: 'Player-4184-000000A1', name: 'Simfury', class: 'Warrior' }];

function player(overrides: Partial<RatingCardPlayer> = {}): RatingCardPlayer {
  return {
    player_key: 'us/normal/simfury',
    player_name: 'Simfury',
    class: 'Warrior',
    spec: 'Fury',
    role: 'dps',
    overall: 70,
    overall_uncapped: 70,
    overall_capped: false,
    basis: 'percentile',
    components: [],
    ...overrides,
  };
}

describe('guidForPlayer', () => {
  it('joins by display name, stripping the log’s trailing -<segment> on both sides', () => {
    expect(guidForPlayer(player({ player_name: 'Simfury-A1B2' }), ROSTER)).toBe('Player-4184-000000A1');
  });

  it('a name with no roster match returns an empty string, not a guess', () => {
    expect(guidForPlayer(player({ player_name: 'Nobody' }), ROSTER)).toBe('');
  });
});

describe('classForPlayer', () => {
  it('prefers the roster’s own class, falling back to the rating row’s', () => {
    expect(classForPlayer('Player-4184-000000A1', 'Priest', ROSTER)).toBe('Warrior');
    expect(classForPlayer('', 'Priest', ROSTER)).toBe('Priest');
  });
});
