// web/src/lib/sim/candidates.test.ts
import { describe, expect, it } from 'vitest';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import {
  addRow,
  buildBulkSpec,
  candidateKey,
  copyAndModify,
  envelopeSlotOf,
  removeRow,
  rowFor,
  toCandidates,
  toggleRow,
  uiSlotsOf,
  validateBulk,
  type CandidateRow,
} from './candidates';
import { KEEP_CURRENT_ENCHANT, NO_ENCHANT } from './enchants';
import type { Item } from '../planner/types';

const items = itemsJson.items as unknown as Item[];
const helm = items.find((item) => item.id === 16963)!;
const ring = items.find((item) => item.id === 19325)!;
const charm = items.find((item) => item.id === 13968)!;

describe('envelopeSlotOf', () => {
  it('is "" for anything that fits more than one slot, and the slot otherwise', () => {
    expect(envelopeSlotOf(helm)).toBe('head');
    expect(envelopeSlotOf(ring)).toBe('');
    expect(envelopeSlotOf(charm)).toBe('');
  });
});

describe('uiSlotsOf', () => {
  it('gives the grid one row for rings and trinkets, not two', () => {
    expect(uiSlotsOf(helm)).toEqual(['head']);
    expect(uiSlotsOf(ring)).toEqual(['finger1']);
    expect(uiSlotsOf(charm)).toEqual(['trinket1']);
  });
});

describe('candidateKey', () => {
  it('identifies a row by slot, item, enchant and suffix, so a copy is a different row', () => {
    const base = rowFor(helm, 'head', 'bag');
    const enchanted = { ...base, enchant: 2543 };
    expect(candidateKey(base)).toBe('head:16963:0:0');
    expect(candidateKey(enchanted)).toBe('head:16963:2543:0');
  });
});

describe('the row list', () => {
  const start: CandidateRow[] = [rowFor(helm, 'head', 'bag'), rowFor(ring, 'finger1', 'equipped')];

  it('toggles one row and leaves the list otherwise identical', () => {
    const next = toggleRow(start, candidateKey(start[0]));
    expect(next[0].checked).toBe(!start[0].checked);
    expect(next[1]).toEqual(start[1]);
    expect(start[0].checked).toBe(false);
  });

  it('adds a row once, and re-adding it ticks the one already there', () => {
    const added = addRow(start, rowFor(charm, 'trinket1', 'search'));
    expect(added).toHaveLength(3);
    const again = addRow(added, { ...rowFor(charm, 'trinket1', 'search'), checked: true });
    expect(again).toHaveLength(3);
    expect(again[2].checked).toBe(true);
  });

  it('copies and modifies into a new, ticked row beside the original', () => {
    const next = copyAndModify(start, candidateKey(start[0]), { enchant: 2543 });
    expect(next).toHaveLength(3);
    expect(next[1].enchant).toBe(2543);
    expect(next[1].checked).toBe(true);
    expect(next[1].origin).toBe('bag');
  });

  it('removes by key', () => {
    expect(removeRow(start, candidateKey(start[0]))).toHaveLength(1);
  });
});

describe('toCandidates', () => {
  it('sends only ticked rows, with the envelope slot and no sentinel enchant', () => {
    const rows = [
      { ...rowFor(helm, 'head', 'bag'), checked: true },
      { ...rowFor(ring, 'finger1', 'equipped'), checked: true, enchant: KEEP_CURRENT_ENCHANT },
      rowFor(charm, 'trinket1', 'search'),
    ];
    expect(toCandidates(rows, [])).toEqual([
      { slot: 'head', item_id: 16963, origin: 'bag' },
      { slot: '', item_id: 19325, origin: 'equipped' },
    ]);
  });

  it('carries the source name a drop was picked from (contract 10.1 A6)', () => {
    const rows = [{ ...rowFor(helm, 'head', 'drop:raid:mc:11502', 'Ragnaros'), checked: true }];
    expect(toCandidates(rows, [])).toEqual([
      { slot: 'head', item_id: 16963, origin: 'drop:raid:mc:11502', source_name: 'Ragnaros' },
    ]);
  });

  it('never sends a candidate on a locked slot', () => {
    const rows = [{ ...rowFor(helm, 'head', 'bag'), checked: true }];
    expect(toCandidates(rows, ['head'])).toEqual([]);
  });

  it('sends an explicit "no enchant" as an omitted enchant, not as -1', () => {
    const rows = [{ ...rowFor(helm, 'head', 'bag'), checked: true, enchant: NO_ENCHANT }];
    expect(toCandidates(rows, [])[0].enchant).toBeUndefined();
  });
});

describe('buildBulkSpec and validateBulk', () => {
  it('builds a gear spec with the cap and precision it was given', () => {
    const spec = buildBulkSpec({
      mode: 'gear',
      rows: [{ ...rowFor(helm, 'head', 'bag'), checked: true }],
      locked: ['main_hand'],
      loadouts: [],
      sets: [],
      precision: 'fast',
      cap: 400,
    });
    expect(spec.mode).toBe('gear');
    expect(spec.cap).toBe(400);
    expect(spec.locked).toEqual(['main_hand']);
    expect(validateBulk(spec)).toBeNull();
  });

  it('refuses a talents spec that carries candidates, and one with no loadout', () => {
    const withCandidates = buildBulkSpec({
      mode: 'talents',
      rows: [{ ...rowFor(helm, 'head', 'bag'), checked: true }],
      locked: [],
      loadouts: [{ name: 'A', talents: '0-1-' }],
      sets: [],
      precision: 'normal',
      cap: 400,
    });
    expect(withCandidates.candidates).toEqual([]);
    const empty = buildBulkSpec({
      mode: 'talents',
      rows: [],
      locked: [],
      loadouts: [],
      sets: [],
      precision: 'normal',
      cap: 400,
    });
    expect(validateBulk(empty)).not.toBeNull();
  });

  it('refuses a drops spec whose candidates are not all from a drop', () => {
    const spec = buildBulkSpec({
      mode: 'drops',
      rows: [{ ...rowFor(helm, 'head', 'bag'), checked: true }],
      locked: [],
      loadouts: [],
      sets: [],
      precision: 'normal',
      cap: 400,
    });
    expect(validateBulk(spec)).not.toBeNull();
  });

  it('refuses a gear spec with nothing ticked at all', () => {
    const spec = buildBulkSpec({
      mode: 'gear',
      rows: [],
      locked: [],
      loadouts: [],
      sets: [],
      precision: 'fast',
      cap: 400,
    });
    expect(validateBulk(spec)).not.toBeNull();
  });
});
