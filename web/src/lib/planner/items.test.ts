// web/src/lib/planner/items.test.ts
import { describe, expect, it } from 'vitest';
import fixtureItems from '../../fixtures/planner/items/warrior.json';
import { itemsForSlot, rarityBorderColorFor, rarityClassFor, searchItems, wornItemLabel } from './items';
import type { ItemFile } from './types';

const items = (fixtureItems as ItemFile).items;

describe('rarityClassFor', () => {
  it('maps the client quality values onto the rarity tokens', () => {
    expect(rarityClassFor(0)).toBe('text-rarity-poor');
    expect(rarityClassFor(1)).toBe('text-rarity-common');
    expect(rarityClassFor(2)).toBe('text-rarity-uncommon');
    expect(rarityClassFor(3)).toBe('text-rarity-rare-text');
    expect(rarityClassFor(4)).toBe('text-rarity-epic-text');
    expect(rarityClassFor(5)).toBe('text-rarity-legendary');
  });

  it('falls back to common for a quality it does not know', () => {
    expect(rarityClassFor(9)).toBe('text-rarity-common');
  });
});

describe('rarityBorderColorFor', () => {
  it('maps the client quality values onto the raw (unlightened) rarity tokens', () => {
    expect(rarityBorderColorFor(0)).toBe('var(--color-rarity-poor)');
    expect(rarityBorderColorFor(2)).toBe('var(--color-rarity-uncommon)');
    expect(rarityBorderColorFor(3)).toBe('var(--color-rarity-rare)');
    expect(rarityBorderColorFor(4)).toBe('var(--color-rarity-epic)');
    expect(rarityBorderColorFor(5)).toBe('var(--color-rarity-legendary)');
  });

  it('falls back to common for a quality it does not know', () => {
    expect(rarityBorderColorFor(9)).toBe('var(--color-rarity-common)');
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

describe('wornItemLabel', () => {
  const copy = {
    itemNotSimmed: (name: string) => `${name} · not simmed`,
    unknownItem: (id: number) => `Unknown item ${id} · not in our data yet`,
  };

  it('names a planner item, says Empty for no item, and names an outside item as not simmed', () => {
    expect(wornItemLabel('Band of Earthen Might', 21182, {}, copy)).toBe('Band of Earthen Might');
    expect(wornItemLabel(undefined, undefined, {}, copy)).toBe('Empty');
    expect(wornItemLabel(undefined, 264908, { '264908': 'Ancient Heirloom' }, copy)).toBe(
      'Ancient Heirloom · not simmed',
    );
  });

  it("is honest that an id nothing knows is our gap, not the player's", () => {
    expect(wornItemLabel(undefined, 999999, {}, copy)).toBe('Unknown item 999999 · not in our data yet');
  });
});
