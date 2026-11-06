import { describe, expect, it } from 'vitest';
import { defaultSelection, dollColumns, dollKeyTarget, readingOrder, type DollCell } from './paperdoll';

const item = (id: number) => ({ id, name: `Item ${id}`, quality: 3, icon: 'x' });
const cell = (slot: string, id: number | null, extra: Partial<DollCell> = {}): DollCell => ({
  slot,
  label: slot,
  item: id === null ? null : item(id),
  ...extra,
});

const CELLS: DollCell[] = [
  cell('trinket1', null),
  cell('hands', 7),
  cell('head', 1),
  cell('neck', 2),
  cell('main_hand', 3),
  cell('off_hand', null),
  cell('ranged', 4),
  cell('waist', 8, { badge: 'upgrade' }),
];

describe('dollColumns', () => {
  it('puts head to ranged down the left and hands to trinkets down the right, in panel order', () => {
    const { left, right } = dollColumns(CELLS);
    expect(left.map((c) => c.slot)).toEqual(['head', 'neck', 'main_hand', 'off_hand', 'ranged']);
    expect(right.map((c) => c.slot)).toEqual(['hands', 'waist', 'trinket1']);
    expect(readingOrder(CELLS).at(-1)?.slot).toBe('trinket1');
  });
});

describe('dollKeyTarget', () => {
  it('moves down the left then the right, skipping empty slots', () => {
    expect(dollKeyTarget(CELLS, 'main_hand', 'ArrowDown')).toBe('ranged');
    expect(dollKeyTarget(CELLS, 'ranged', 'ArrowDown')).toBe('hands');
    expect(dollKeyTarget(CELLS, 'head', 'ArrowUp')).toBe('head');
    expect(dollKeyTarget(CELLS, 'neck', 'Home')).toBe('head');
    expect(dollKeyTarget(CELLS, 'neck', 'End')).toBe('waist');
  });

  it('crosses to the same row of the other column', () => {
    expect(dollKeyTarget(CELLS, 'neck', 'ArrowRight')).toBe('waist');
    expect(dollKeyTarget(CELLS, 'hands', 'ArrowLeft')).toBe('head');
  });

  it('ignores other keys', () => {
    expect(dollKeyTarget(CELLS, 'head', 'a')).toBeNull();
  });
});

describe('defaultSelection', () => {
  it('opens on the first upgrade, else the first filled slot', () => {
    expect(defaultSelection(CELLS)).toBe('waist');
    expect(defaultSelection(CELLS.map((c) => ({ ...c, badge: undefined })))).toBe('head');
    expect(defaultSelection([cell('head', null)])).toBeNull();
  });
});
