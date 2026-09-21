// web/src/lib/sim/item-search.test.ts
import { describe, expect, it } from 'vitest';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import lootJson from '../../fixtures/planner/loot.json';
import {
  defaultItemQuery,
  isDevFixtureItem,
  matchCount,
  noResultsReason,
  searchItems,
  SEARCH_LIMIT,
  type SearchContext,
} from './item-search';
import { sourcesByItem, type LootFile } from './loot';
import type { Item } from '../planner/types';

const items = itemsJson.items as unknown as Item[];
const ctx: SearchContext = {
  level: 60,
  sourcesByItem: sourcesByItem(lootJson as unknown as LootFile),
};

function item(overrides: Partial<Item> & Pick<Item, 'id' | 'name' | 'slot'>): Item {
  return {
    quality: 4,
    required_level: 0,
    item_level: 60,
    armor: 0,
    stats: {},
    set_id: null,
    unique: false,
    icon: `fixture_${overrides.id}`,
    ...overrides,
  };
}

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

describe('isDevFixtureItem', () => {
  it('flags a name that opens with its own item level -- no real item name does (dps D26)', () => {
    expect(isDevFixtureItem(item({ id: 1, name: '90 Epic Warrior Axe', slot: 'main_hand' }))).toBe(true);
    expect(isDevFixtureItem(item({ id: 2, name: '63 Green Rogue Dagger', slot: 'main_hand' }))).toBe(true);
  });

  it('never flags a real item, including one that legitimately reads ilvl 90 (Atiesh)', () => {
    expect(
      isDevFixtureItem(item({ id: 22589, name: 'Atiesh, Greatstaff of the Guardian', slot: 'main_hand' })),
    ).toBe(false);
    expect(isDevFixtureItem(item({ id: 12784, name: 'Arcanite Reaper', slot: 'main_hand' }))).toBe(false);
  });
});

describe('searchItems excludes dev fixtures (dps D26)', () => {
  const fixtures = [
    item({ id: 901, name: '90 Epic Frost Staff', slot: 'main_hand', item_level: 90 }),
    item({ id: 902, name: '90 Epic Rogue Dagger', slot: 'main_hand', item_level: 90 }),
  ];
  const real = item({ id: 903, name: 'Thunderfury, Blessed Blade of the Windseeker', slot: 'main_hand' });
  const withFixtures = [...fixtures, real];

  it('drops dev fixtures from the default (usable-only) search', () => {
    const found = searchItems(withFixtures, defaultItemQuery(), ctx);
    expect(found.map((row) => row.id)).toEqual([real.id]);
  });

  it('still drops them with usable-only off -- a fixture is never obtainable, at any level', () => {
    const found = searchItems(withFixtures, { ...defaultItemQuery(), usableOnly: false }, ctx);
    expect(found.map((row) => row.id)).toEqual([real.id]);
  });

  it('keeps them out of matchCount too, so the truncation note never counts them', () => {
    expect(matchCount(withFixtures, defaultItemQuery(), ctx)).toBe(1);
  });
});

describe('noResultsReason (dps D27)', () => {
  const twoHander = item({
    id: 950,
    name: 'Ashkandi, Greatsword of the Brotherhood',
    slot: 'main_hand',
    two_hand: true,
  });
  const oneHander = item({ id: 951, name: 'Brutality Blade', slot: 'main_hand' });
  const pool = [twoHander, oneHander];

  it('says nothing extra when the query is empty or unfiltered by slot', () => {
    expect(noResultsReason(pool, defaultItemQuery(), ctx)).toEqual({
      existsElsewhere: false,
      twoHanded: false,
    });
    expect(noResultsReason(pool, { ...defaultItemQuery(), text: 'Ashkandi' }, ctx)).toEqual({
      existsElsewhere: false,
      twoHanded: false,
    });
  });

  it('flags a two-handed weapon searched under off_hand as findable elsewhere, and why', () => {
    const query = { ...defaultItemQuery(), text: 'Ashkandi', slot: 'off_hand' };
    expect(noResultsReason(pool, query, ctx)).toEqual({ existsElsewhere: true, twoHanded: true });
  });

  it('flags a one-hander searched under the wrong slot as findable elsewhere, not two-handed', () => {
    const query = { ...defaultItemQuery(), text: 'Brutality', slot: 'off_hand' };
    expect(noResultsReason(pool, query, ctx)).toEqual({ existsElsewhere: true, twoHanded: false });
  });

  it('reports nothing extra when the name matches nowhere at all, in any slot', () => {
    const query = { ...defaultItemQuery(), text: 'Maladath', slot: 'off_hand' };
    expect(noResultsReason(pool, query, ctx)).toEqual({ existsElsewhere: false, twoHanded: false });
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
