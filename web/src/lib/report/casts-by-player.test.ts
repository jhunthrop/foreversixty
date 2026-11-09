import { describe, expect, it } from 'vitest';
import { castsByPlayer } from './casts-by-player';
import type { CastRow } from './types';

function cast(guid: string, spell_id: number, succeeded: number, failed = 0, owner_guid?: string): CastRow {
  return {
    guid,
    name: guid,
    owner_guid,
    spell_id,
    spell_name: `Spell ${spell_id}`,
    started: 0,
    succeeded,
    failed,
    cast_time_ms: 0,
    sequence: [],
  };
}

describe('castsByPlayer', () => {
  it('counts each player’s own casts, not their pets’ or totems’, per minute over the window, busiest first', () => {
    const rows = [
      cast('A', 1, 6, 2),
      cast('A', 2, 3),
      cast('Pet-1', 3, 3, 0, 'A'),
      cast('B', 1, 20, 0),
      cast('Creature-9', 4, 50),
    ];
    expect(castsByPlayer(rows, new Set(['A', 'B']), 30_000)).toEqual([
      { guid: 'B', casts: 20, failed: 0, spells: 1, perMinute: 40 },
      { guid: 'A', casts: 9, failed: 2, spells: 2, perMinute: 18 },
    ]);
  });

  it('counts a spell that only ever failed as no spell used, and leaves out a player who cast nothing', () => {
    const rows = [cast('A', 1, 0, 4), cast('B', 1, 0, 0)];
    expect(castsByPlayer(rows, new Set(['A', 'B']), 60_000)).toEqual([
      { guid: 'A', casts: 0, failed: 4, spells: 0, perMinute: 0 },
    ]);
  });
});
