// web/src/lib/guides/build-diff.test.ts
import { describe, expect, it } from 'vitest';
import { buildDiffLine, diffBuilds } from './build-diff';
import type { DecodedTalentView, DecodedTreeView } from './build-tree';

function talent(id: number, name: string, rank: number): DecodedTalentView {
  return {
    id,
    name,
    icon: `icon_${id}`,
    max_rank: 5,
    tier: id,
    column: 0,
    prereq_talent_id: null,
    prereq_rank: null,
    ranks: [],
    spell_id: id,
    rank,
  };
}

function tree(...talents: DecodedTalentView[]): DecodedTreeView {
  return {
    name: 'Tree',
    position: 0,
    talents,
    points: talents.reduce((sum, t) => sum + t.rank, 0),
  };
}

describe('diffBuilds', () => {
  const leveling = [
    tree(talent(1, 'Improved Fireball', 2), talent(2, 'Impact', 3), talent(3, 'Incineration', 0)),
  ];
  const raid = [
    tree(talent(1, 'Improved Fireball', 0), talent(2, 'Impact', 0), talent(3, 'Incineration', 5)),
  ];

  it('lists what loses points, what gains them and how many moved', () => {
    expect(diffBuilds(leveling, raid)).toEqual({
      from: [
        { name: 'Improved Fireball', points: 2 },
        { name: 'Impact', points: 3 },
      ],
      to: [{ name: 'Incineration', points: 5 }],
      moved: 5,
    });
  });

  it('is undefined for the same build', () => {
    expect(diffBuilds(leveling, leveling)).toBeUndefined();
  });
});

describe('buildDiffLine', () => {
  it('names the move in one sentence', () => {
    const line = buildDiffLine({
      from: [
        { name: 'A', points: 2 },
        { name: 'B', points: 3 },
      ],
      to: [{ name: 'C', points: 5 }],
      moved: 5,
    });
    expect(line).toBe('The raid build moves 5 points from A (2) and B (3) to C (5).');
  });

  it('says "1 point" in the singular and caps a long list', () => {
    const line = buildDiffLine({
      from: [{ name: 'A', points: 1 }],
      to: ['C', 'D', 'E', 'F', 'G', 'H'].map((name) => ({ name, points: 1 })),
      moved: 1,
    });
    expect(line).toBe('The raid build moves 1 point from A (1) to C (1), D (1), E (1), F (1) and 2 more.');
  });

  it('handles a build that only adds or only removes points', () => {
    expect(buildDiffLine({ from: [], to: [{ name: 'C', points: 2 }], moved: 0 })).toBe(
      'The raid build puts 2 more points into C (2).',
    );
    expect(buildDiffLine({ from: [{ name: 'A', points: 2 }], to: [], moved: 2 })).toBe(
      'The raid build takes 2 points out of A (2).',
    );
  });

  it('is undefined when there is no difference', () => {
    expect(buildDiffLine(undefined)).toBeUndefined();
  });
});
