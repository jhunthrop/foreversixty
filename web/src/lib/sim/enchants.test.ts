// web/src/lib/sim/enchants.test.ts
import { describe, expect, it } from 'vitest';
import enchantsJson from '../../fixtures/planner/enchants.json';
import suffixesJson from '../../fixtures/planner/suffixes.json';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import {
  ENCHANTS_PER_SLOT_CAP,
  KEEP_CURRENT_ENCHANT,
  NO_ENCHANT,
  enchantsForSlot,
  suffixesForItem,
  type EnchantRow,
  type SuffixRow,
} from './enchants';
import type { Item } from '../planner/types';

const enchants = enchantsJson as unknown as EnchantRow[];
const suffixes = suffixesJson as unknown as SuffixRow[];
const items = itemsJson.items as unknown as Item[];
const head = items.find((item) => item.id === 12640)!;
const charm = items.find((item) => item.id === 13968)!;

describe('enchantsForSlot', () => {
  it('offers only the enchants whose slot list names this slot', () => {
    expect(enchantsForSlot(enchants, 'head').map((row) => row.id)).toEqual([2543]);
    expect(enchantsForSlot(enchants, 'main_hand').map((row) => row.id)).toEqual([1900]);
    expect(enchantsForSlot(enchants, 'neck')).toEqual([]);
  });

  it('is empty rather than wrong when the build ships no enchant file', () => {
    expect(enchantsForSlot([], 'head')).toEqual([]);
  });
});

describe('suffixesForItem', () => {
  it('offers only the suffixes the item rolls', () => {
    expect(suffixesForItem(suffixes, charm).map((row) => row.name)).toEqual(['of the Bear', 'of the Tiger']);
    expect(suffixesForItem(suffixes, head)).toEqual([]);
  });
});

describe('the two sentinel enchant values', () => {
  it('are distinct and are not real effect ids', () => {
    expect(KEEP_CURRENT_ENCHANT).toBe(-1);
    expect(NO_ENCHANT).toBe(0);
    expect(ENCHANTS_PER_SLOT_CAP).toBe(4);
  });
});

describe('the enchant row shape (contract 10.4, as corrected by the real build file)', () => {
  it('is keyed by id, which is what GearSlot.enchant carries', () => {
    for (const row of enchants) expect(typeof row.id).toBe('number');
  });

  it("carries the real file's string-enum item_types and a numeric phase", () => {
    for (const row of enchants) {
      expect(Array.isArray(row.item_types)).toBe(true);
      for (const type of row.item_types) expect(typeof type).toBe('string');
      expect(typeof row.phase).toBe('number');
    }
  });
});
