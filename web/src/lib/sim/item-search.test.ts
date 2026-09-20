// web/src/lib/sim/item-search.test.ts
import { describe, expect, it } from 'vitest';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import lootJson from '../../fixtures/planner/loot.json';
import { defaultItemQuery, matchCount, searchItems, SEARCH_LIMIT, type SearchContext } from './item-search';
import { sourcesByItem, type LootFile } from './loot';
import type { Item } from '../planner/types';

const items = itemsJson.items as unknown as Item[];
const ctx: SearchContext = {
  level: 60,
  sourcesByItem: sourcesByItem(lootJson as unknown as LootFile),
};

describe('searchItems', () => {
  it('returns everything for an empty query, sorted by item level then name', () => {
    const found = searchItems(items, defaultItemQuery(), ctx);
    expect(found).toHaveLength(items.length);
    expect(found[0].item_level).toBeGreaterThanOrEqual(found[found.length - 1].item_level);
  });

  it('matches on name, case-insensitively, anywhere in the name', () => {
    const found = searchItems(items, { ...defaultItemQuery(), text: 'wrath' }, ctx);
    expect(found.map((item) => item.id).sort()).toEqual([16963, 16966]);
  });

  it('filters on minimum item level', () => {
    const found = searchItems(items, { ...defaultItemQuery(), minItemLevel: 76 }, ctx);
    expect(found.map((item) => item.id).sort()).toEqual([16963, 16966, 19325]);
  });

  it('filters on slot, following the finger and trinket aliases', () => {
    expect(searchItems(items, { ...defaultItemQuery(), slot: 'finger1' }, ctx).map((i) => i.id)).toEqual([
      19325,
    ]);
    expect(
      searchItems(items, { ...defaultItemQuery(), slot: 'head' }, ctx)
        .map((i) => i.id)
        .sort(),
    ).toEqual([12640, 16963]);
  });

  it('filters on source', () => {
    const found = searchItems(items, { ...defaultItemQuery(), sourceId: 'world:azuregos' }, ctx);
    expect(found.map((item) => item.id)).toEqual([19325]);
  });

  it('drops what the character cannot equip when usableOnly is on, and keeps it when off', () => {
    const lowLevel: SearchContext = { ...ctx, level: 59 };
    const on = searchItems(items, defaultItemQuery(), lowLevel);
    expect(on.map((item) => item.id)).toEqual([13968]);
    const off = searchItems(items, { ...defaultItemQuery(), usableOnly: false }, lowLevel);
    expect(off).toHaveLength(items.length);
  });

  it('drops an item the engine does not know about when a known set is given, and keeps it when there is none', () => {
    // 21550 (Idol of the White Stag) is the fixture item deliberately left out of
    // simitems.json -- see web/src/fixtures/planner/simitems.json's own header comment.
    const known = new Set(items.filter((item) => item.id !== 21550).map((item) => item.id));
    const filtered = searchItems(items, defaultItemQuery(), { ...ctx, known });
    expect(filtered.map((item) => item.id)).not.toContain(21550);
    expect(filtered).toHaveLength(items.length - 1);

    // `known` absent (every caller before this field existed) and `known: null` (a build
    // that ships no simitems.json) both mean "nothing to filter" -- 21550 is back.
    expect(searchItems(items, defaultItemQuery(), ctx).map((item) => item.id)).toContain(21550);
    expect(searchItems(items, defaultItemQuery(), { ...ctx, known: null }).map((item) => item.id)).toContain(
      21550,
    );
  });

  it('never returns more than the limit', () => {
    const many = Array.from({ length: SEARCH_LIMIT + 20 }, (_, i) => ({ ...items[0], id: 1000 + i }));
    expect(searchItems(many, defaultItemQuery(), ctx)).toHaveLength(SEARCH_LIMIT);
  });
});

describe('matchCount', () => {
  it('agrees with searchItems under the limit, and counts past it where searchItems truncates', () => {
    expect(matchCount(items, defaultItemQuery(), ctx)).toBe(items.length);
    expect(matchCount(items, { ...defaultItemQuery(), text: 'wrath' }, ctx)).toBe(2);
    const many = Array.from({ length: SEARCH_LIMIT + 20 }, (_, i) => ({ ...items[0], id: 1000 + i }));
    expect(matchCount(many, defaultItemQuery(), ctx)).toBe(SEARCH_LIMIT + 20);
  });
});

describe('defaultItemQuery', () => {
  it('opens with usable-only on, the way the design specifies', () => {
    expect(defaultItemQuery()).toEqual({
      text: '',
      minItemLevel: 0,
      slot: '',
      sourceId: '',
      usableOnly: true,
    });
  });
});
