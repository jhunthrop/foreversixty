import { describe, expect, it } from 'vitest';
import { averageWornItemLevel, characterMeta, dayMonth, raceWords } from './head-meta';

describe('the simulator head meta', () => {
  it('reads the race in words', () => {
    expect(raceWords('high-order-skyborne')).toBe('High Order Skyborne');
  });

  it('averages worn item levels, or says nothing when one is unknown', () => {
    const levels = new Map([
      [1, 60],
      [2, 59],
    ]);
    expect(averageWornItemLevel([1, 2], (id) => levels.get(id))).toBe('59.5');
    expect(averageWornItemLevel([1, 3], (id) => levels.get(id))).toBeNull();
  });

  it('names the gear source and its day', () => {
    const [line1, line2] = characterMeta(
      {
        race: 'troll',
        spec: 'warrior-fury',
        level: 60,
        talents: '05-32-2',
        source: { kind: 'addon', ref: '', captured_at: '2026-11-04T10:00:00Z' },
      },
      '59.9',
    );
    expect(line1).toMatch(/^Troll .*Warrior · Level 60 · 5\/5\/2 talents · item level 59\.9$/);
    expect(line2).toBe('Gear from the addon, 4 Nov');
    expect(dayMonth('')).toBeNull();
  });
});
