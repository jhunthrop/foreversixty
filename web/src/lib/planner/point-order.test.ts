// web/src/lib/planner/point-order.test.ts
import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { pointOrderRows } from './point-order';
import { indexTalents } from './rules';
import type { TalentFile } from './types';

const index = indexTalents(fixtureTalents as TalentFile);
const icon = (name: string) => `/i/${name}`;

describe('pointOrderRows', () => {
  it('draws six rows of ten levels, the last one level 60 alone', () => {
    const view = pointOrderRows(index, [], null, null, icon);
    expect(view.rows.map((row) => row.label)).toEqual([
      'Levels 10 to 19',
      'Levels 20 to 29',
      'Levels 30 to 39',
      'Levels 40 to 49',
      'Levels 50 to 59',
      'Level 60',
    ]);
    expect(view.rows.map((row) => row.slots.length)).toEqual([10, 10, 10, 10, 10, 1]);
    expect(view.rows[0].slots[0]).toEqual({ kind: 'next', name: 'Next point: level 10' });
    expect(view.capAfter).toBe(-1);
    expect(view.capText).toBe('');
  });

  it('names each spent point with the rank it reached, then the next socket', () => {
    const view = pointOrderRows(index, [1001, 1001], null, null, icon);
    const [first, second, next] = view.rows[0].slots;
    expect(first.kind).toBe('spent');
    expect(first.rank).toBe(1);
    expect(second.rank).toBe(2);
    expect(second.name).toMatch(/^Level 11: .+ 2$/);
    expect(first.iconSrc).toMatch(/^\/i\//);
    expect(next).toEqual({ kind: 'next', name: 'Next point: level 12' });
  });

  it('marks points past a character level as planned, with the level line after its row (P4)', () => {
    const view = pointOrderRows(index, [1001, 1001, 1002], 11, 'Felsnap', icon);
    expect(view.rows[0].slots.map((slot) => slot.kind).slice(0, 4)).toEqual([
      'spent',
      'spent',
      'planned',
      'empty',
    ]);
    expect(view.capAfter).toBe(0);
    expect(view.capText).toBe('Felsnap is level 11: points after level 11 are planned ahead.');
  });

  it('shows no level line before the build reaches the character level', () => {
    const view = pointOrderRows(index, [1001], 34, 'Felsnap', icon);
    expect(view.capAfter).toBe(2);
    expect(view.capText).toBe('');
  });
});
