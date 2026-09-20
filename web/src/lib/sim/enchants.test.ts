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
  type EnchantStatKey,
  type SuffixRow,
} from './enchants';
import type { Item } from '../planner/types';

/**
 * Every `EnchantStatKey` member, as a compile-time-checked object: adding a member to the
 * union without adding it here fails `tsc`, so this stays exhaustive alongside the type.
 */
const ALL_ENCHANT_STAT_KEYS: Record<EnchantStatKey, true> = {
  strength: true,
  agility: true,
  stamina: true,
  intellect: true,
  spirit: true,
  armor: true,
  crit: true,
  hit: true,
  spell_power: true,
  healing: true,
  attack_power: true,
  defense: true,
  dodge: true,
  parry: true,
  block: true,
  mp5: true,
  spell_penetration: true,
  fire_res: true,
  frost_res: true,
  nature_res: true,
  shadow_res: true,
  arcane_res: true,
  arcane_power: true,
  block_value: true,
  bonus_armor: true,
  fire_power: true,
  frost_power: true,
  health: true,
  holy_power: true,
  mana: true,
  melee_haste: true,
  nature_power: true,
  ranged_attack_power: true,
  shadow_power: true,
  spell_damage: true,
};

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

  it('gates a class-restricted enchant on classSlug (none of the fixture rows are restricted, so this uses a synthetic row)', () => {
    const restricted: EnchantRow = { ...enchants[0], id: 9999, classes: ['warrior'], slots: ['head'] };
    expect(enchantsForSlot([restricted], 'head', 'mage')).toEqual([]);
    expect(enchantsForSlot([restricted], 'head', 'warrior')).toEqual([restricted]);
    // The default classSlug ('') applies no class restriction at all -- it is "unknown
    // class", not "no class."
    expect(enchantsForSlot([restricted], 'head')).toEqual([restricted]);
  });
});

describe('suffixesForItem', () => {
  it('offers only the suffixes the item rolls', () => {
    expect(suffixesForItem(suffixes, charm).map((row) => row.name)).toEqual(['of the Bear', 'of the Tiger']);
    expect(suffixesForItem(suffixes, head)).toEqual([]);
  });

  it('is empty for an item with suffixes: [] explicitly, same as one with no suffixes field at all', () => {
    const explicitlyEmpty: Item = { ...head, suffixes: [] };
    expect(head.suffixes).toBeUndefined();
    expect(suffixesForItem(suffixes, explicitlyEmpty)).toEqual([]);
  });
});

describe('EnchantStatKey', () => {
  it('is a closed union that accepts every stat key the fixture files carry', () => {
    const keys = new Set<string>();
    for (const row of [...enchants, ...suffixes]) {
      for (const key of Object.keys(row.stats)) keys.add(key);
    }
    expect(keys.size).toBeGreaterThan(0);
    for (const key of keys) expect(key in ALL_ENCHANT_STAT_KEYS).toBe(true);
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
