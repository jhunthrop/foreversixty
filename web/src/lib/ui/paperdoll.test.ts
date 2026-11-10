import { describe, expect, it } from 'vitest';
import { defaultSelection, selectable, type DollCell } from './paperdoll';

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

describe('defaultSelection', () => {
  it('opens on the first upgrade, else the first filled slot', () => {
    expect(defaultSelection(CELLS)).toBe('waist');
    expect(defaultSelection(CELLS.map((c) => ({ ...c, badge: undefined })))).toBe('head');
    expect(defaultSelection([cell('head', null)])).toBeNull();
  });

  it('opens down the left column before the right', () => {
    const plain = CELLS.map((c) => ({ ...c, badge: undefined }));
    expect(defaultSelection(plain.filter((c) => c.slot !== 'head' && c.slot !== 'neck'))).toBe('main_hand');
  });
});

describe('selectable', () => {
  it('takes a slot with an item, never an empty one', () => {
    expect(selectable(cell('head', 1))).toBe(true);
    expect(selectable(cell('trinket1', null))).toBe(false);
  });
});
