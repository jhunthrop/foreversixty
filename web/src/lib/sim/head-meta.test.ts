import { describe, expect, it } from 'vitest';
import { characterMeta, raceWords, runDay } from './head-meta';

describe('the simulator head meta', () => {
  it('reads the race in words', () => {
    expect(raceWords('high-order-skyborne')).toBe('High Order Skyborne');
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
    expect(runDay('')).toBeNull();
  });

  it('says nothing of an addon export with no capture date (the default)', () => {
    const lines = characterMeta(
      {
        race: 'troll',
        spec: 'warrior-fury',
        level: 60,
        talents: '05-32-2',
        source: { kind: 'addon', ref: '', captured_at: '' },
      },
      null,
    );
    expect(lines).toHaveLength(1);
  });
});
