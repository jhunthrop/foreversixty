// web/src/lib/planner/point-order.test.ts
import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { aheadOf, pointOrderRows } from './point-order';
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
    expect(view.capText).toBe('Felsnap is level 11.');
  });

  it('gives each point the talent tooltip: its name, the rank it reached and that rank text', () => {
    const view = pointOrderRows(index, [1001, 1001], null, null, icon);
    const talent = index.byId.get(1001)!;
    expect(view.rows[0].slots[1].tip).toEqual({
      name: talent.name,
      rankLine: `Rank 2 of ${talent.max_rank}`,
      now: talent.ranks[1].description,
    });
  });

  it("fills a character's sockets past its build with the guide's order, dimmed, under its line", () => {
    const ahead = aheadOf(index, [1001, 1001, 1002], [1001, 1002, 1001, 1001, 1002]);
    expect(ahead).toEqual([1001, 1002]);
    const view = pointOrderRows(index, [1001, 1001, 1002], 11, 'Felsnap', icon, {
      specName: 'Arms',
      points: ahead,
    });
    expect(view.rows[0].slots.map((slot) => slot.kind).slice(0, 6)).toEqual([
      'spent',
      'spent',
      'planned',
      'ahead',
      'ahead',
      'empty',
    ]);
    expect(view.rows[0].slots[3].rank).toBe(3);
    expect(view.capText).toBe('Felsnap is 11. Then Arms guide order.');
    // An empty build previews the guide's order, dimmed.
    expect(
      pointOrderRows(index, [], null, null, icon, { specName: 'Arms', points: [1001] }).rows[0].slots[0].kind,
    ).toBe('ahead');
    // A bare build has no guide plan: its next socket stays.
    expect(
      pointOrderRows(index, [1001], null, null, icon, { specName: 'Arms', points: [1001] }).rows[0].slots[1]
        .kind,
    ).toBe('next');
  });

  it('skips a guide point that is not legal on top of the build', () => {
    expect(aheadOf(index, [], [9999, 1001])).toEqual([1001]);
  });

  it('shows no level line before the build reaches the character level', () => {
    const view = pointOrderRows(index, [1001], 34, 'Felsnap', icon);
    expect(view.capAfter).toBe(2);
    expect(view.capText).toBe('');
  });
});
