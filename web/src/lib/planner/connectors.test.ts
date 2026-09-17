// web/src/lib/planner/connectors.test.ts
import { describe, expect, it } from 'vitest';
import { cellCentre, connectorPath, connectorsFor, gridViewBox, PITCH_PX } from './connectors';
import type { Talent, TalentTree } from './types';

function talent(overrides: Partial<Talent> & Pick<Talent, 'id' | 'tier' | 'column'>): Talent {
  return {
    name: `Talent ${overrides.id}`,
    icon: 'icon',
    max_rank: 5,
    prereq_talent_id: null,
    prereq_rank: null,
    ranks: [],
    spell_id: 1,
    ...overrides,
  };
}

function tree(talents: Talent[]): TalentTree {
  return { id: 161, name: 'Arms', position: 0, background: 'warriorarms', talents };
}

describe('cellCentre', () => {
  it('puts row 0 column 0 half a cell in, and steps by the pitch', () => {
    expect(cellCentre({ tier: 0, column: 0 })).toEqual({ x: 22, y: 22 });
    expect(cellCentre({ tier: 1, column: 1 })).toEqual({ x: 22 + PITCH_PX, y: 22 + PITCH_PX });
    expect(cellCentre({ tier: 6, column: 3 })).toEqual({ x: 178, y: 334 });
  });
});

describe('connectorPath', () => {
  it('runs straight down the column, edge to edge', () => {
    // 67 of the beta build's 69 prerequisites are this shape.
    expect(connectorPath({ tier: 0, column: 1 }, { tier: 1, column: 1 })).toBe('M 74 44 V 52');
  });

  it('spans more than one row when the prerequisite is further up', () => {
    expect(connectorPath({ tier: 0, column: 0 }, { tier: 2, column: 0 })).toBe('M 22 44 V 104');
  });

  it('runs straight across the row to the right', () => {
    // Priest: Mind Flay (2,2) -> Improved Mind Flay (2,3).
    expect(connectorPath({ tier: 2, column: 2 }, { tier: 2, column: 3 })).toBe('M 148 126 H 156');
  });

  it('runs straight across the row to the left', () => {
    // Paladin: Holy Shock (4,1) -> Divine Precision (4,0).
    expect(connectorPath({ tier: 4, column: 1 }, { tier: 4, column: 0 })).toBe('M 52 230 H 44');
  });

  it('elbows down the prerequisite column and then across', () => {
    // No prerequisite in build 1.60.1.69893 needs this, but the shape is the
    // game's and a future build that adds one must not draw a wrong line.
    expect(connectorPath({ tier: 0, column: 0 }, { tier: 1, column: 1 })).toBe('M 22 44 V 74 H 52');
    expect(connectorPath({ tier: 0, column: 2 }, { tier: 1, column: 1 })).toBe('M 126 44 V 74 H 96');
  });
});

describe('gridViewBox', () => {
  it('is the grid with no trailing gap', () => {
    expect(gridViewBox(7, 4)).toBe('0 0 200 356');
    expect(gridViewBox(4, 2)).toBe('0 0 96 200');
  });
});

describe('connectorsFor', () => {
  const arms = tree([
    talent({ id: 1001, tier: 0, column: 1 }),
    talent({ id: 1002, tier: 1, column: 1, prereq_talent_id: 1001, prereq_rank: 5 }),
    talent({ id: 1003, tier: 2, column: 0 }),
  ]);

  it('draws one connector per prerequisite, unmet until the rank is reached', () => {
    expect(connectorsFor(arms, new Map())).toEqual([
      { id: '1001-1002', from: 1001, to: 1002, d: 'M 74 44 V 52', met: false },
    ]);
  });

  it('marks a connector met once the prerequisite holds the rank it needs', () => {
    expect(connectorsFor(arms, new Map([[1001, 4]]))[0].met).toBe(false);
    expect(connectorsFor(arms, new Map([[1001, 5]]))[0].met).toBe(true);
    expect(connectorsFor(arms, new Map([[1001, 5]]))[0].id).toBe('1001-1002');
  });

  it('ignores a prerequisite that is not in the tree', () => {
    const broken = tree([talent({ id: 1, tier: 1, column: 0, prereq_talent_id: 99, prereq_rank: 1 })]);
    expect(connectorsFor(broken, new Map())).toEqual([]);
  });
});
