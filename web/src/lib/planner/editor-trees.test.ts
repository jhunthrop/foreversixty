// web/src/lib/planner/editor-trees.test.ts
import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import type { BandDiffView } from './band-compare';
import { editorTrees, type EditorTreesInput } from './editor-trees';
import { indexTalents } from './rules';
import type { TalentFile } from './types';

const index = indexTalents(fixtureTalents as TalentFile);
const base: EditorTreesInput = {
  index,
  order: [],
  classSlug: 'warrior',
  characterLevel: null,
  readOnly: false,
  bandDiff: null,
  asset: (path) => `/d/${path}`,
};
const cell = (input: Partial<EditorTreesInput>, id: number) =>
  editorTrees({ ...base, ...input })
    .flatMap((tree) => tree.cells)
    .find((c) => c.id === id)!;

describe('editorTrees', () => {
  it('draws the game colours: can spend, partial, maxed, locked (P2)', () => {
    expect(cell({}, 1001).state).toBe('avail');
    expect(cell({ order: [1001] }, 1001).state).toBe('part');
    expect(cell({ order: [1001, 1001, 1001] }, 1001).state).toBe('max');
    // A maxed talent's tooltip says nothing past its rank, as the client's.
    expect(cell({ order: [1001, 1001, 1001] }, 1001).req).toBe('');
    const locked = cell({}, 1003);
    expect(locked.state).toBe('lock');
    expect(locked.req).not.toBe('');
  });

  it('counts each tree and serves its art', () => {
    const trees = editorTrees({ ...base, order: [1001, 2001, 2001] });
    expect(trees.map((tree) => tree.points)).toEqual([1, 2]);
    expect(trees[0].cells.length).toBeGreaterThan(0);
  });

  it('greys the unspent cells at a character cap and says why (P3), marking points past it (P4)', () => {
    const atCap = cell({ order: [1001], characterLevel: 10 }, 1002);
    expect(atCap.state).toBe('lock');
    expect(atCap.req).toBe('No points left at level 10.');
    const planned = cell({ order: [1001, 1001], characterLevel: 10 }, 1001);
    expect(planned.planned).toBe(true);
    expect(planned.plannedLabel).toBe('planned past level 10');
    expect(cell({ order: [1001], characterLevel: 10 }, 1001).planned).toBeUndefined();
  });

  it('dots a cell that differs from the band once a point is spent, with its band line', () => {
    const bandDiff = {
      bandLabel: '60',
      diffTalentIds: new Set([1002]),
      bandRankById: new Map([[1002, 5]]),
    } as unknown as BandDiffView;
    expect(cell({ bandDiff }, 1002).differs).toBe(false);
    const dotted = cell({ bandDiff, order: [1001] }, 1002);
    expect(dotted.differs).toBe(true);
    expect(dotted.band).toBe('Level 60 build: 5 of 5');
  });
});
