// web/src/lib/planner/items.test.ts
import { describe, expect, it } from 'vitest';
import fixtureItems from '../../fixtures/planner/items/warrior.json';
import { itemsForSlot, rarityClassFor, searchItems } from './items';
import type { ItemFile } from './types';

const items = (fixtureItems as ItemFile).items;

describe('rarityClassFor', () => {
  it('maps the client quality values onto the rarity tokens', () => {
    expect(rarityClassFor(0)).toBe('text-rarity-poor');
    expect(rarityClassFor(1)).toBe('text-rarity-common');
    expect(rarityClassFor(2)).toBe('text-rarity-uncommon');
    expect(rarityClassFor(3)).toBe('text-rarity-rare');
    expect(rarityClassFor(4)).toBe('text-rarity-epic');
    expect(rarityClassFor(5)).toBe('text-rarity-legendary');
  });

  it('falls back to common for a quality it does not know', () => {
    expect(rarityClassFor(9)).toBe('text-rarity-common');
  });
});

describe('itemsForSlot', () => {
  it('keeps only the items that fit, sorted by required level then name', () => {
    expect(itemsForSlot(items, 'head').map((i) => i.name)).toEqual(['Helm of Wrath', 'Lionheart Helm']);
  });

  it('maps the finger alias onto both numbered slots', () => {
    expect(itemsForSlot(items, 'finger1').map((i) => i.id)).toEqual([19325]);
    expect(itemsForSlot(items, 'finger2').map((i) => i.id)).toEqual([19325]);
  });

  it('is empty for a slot nothing in the class list fits', () => {
    expect(itemsForSlot(items, 'legs')).toEqual([]);
  });
});

describe('searchItems', () => {
  it('matches on a case-insensitive substring of the name', () => {
    expect(
      searchItems(items, 'wrath')
        .map((i) => i.id)
        .sort(),
    ).toEqual([16963, 16966]);
    expect(searchItems(items, 'REAPER').map((i) => i.id)).toEqual([12784]);
  });

  it('returns everything for a blank query', () => {
    expect(searchItems(items, '   ')).toHaveLength(items.length);
  });
});
