// web/src/lib/guild/roster-sort.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildRosterRow } from './api';
import {
  DEFAULT_ROSTER_FILTERS,
  altOfMainName,
  applyRosterFilters,
  isBelowRatingFloor,
  orderRosterForTab,
  pinOwnRowFirst,
  ratingFloor,
  sortRoster,
} from './roster-sort';

function row(overrides: Partial<GuildRosterRow> = {}): GuildRosterRow {
  return {
    character_key: 'us/pvp/kraggor',
    region: 'us',
    ruleset: 'pvp',
    name: 'Kraggor',
    class: 'warrior',
    spec: 'Protection',
    role: 'tank',
    rank: 'leader',
    verified: true,
    logged_recently: true,
    consent: 'gear_bags',
    item_level: 68,
    may_remove: false,
    ...overrides,
  };
}

describe('ratingFloor / isBelowRatingFloor', () => {
  it('is the 25th percentile of the roster’s own rated population', () => {
    const rows = [10, 20, 30, 40, 50, 60, 70, 80].map((overall, i) =>
      row({
        character_key: `c${i}`,
        rating: { overall, output: 0, survival: 0, mechanics: 0, utility: 0, preparation: 0, activity: 0 },
      }),
    );
    expect(ratingFloor(rows)).toBe(30);
  });

  it('is null when nobody has a rating yet', () => {
    expect(ratingFloor([row({ rating: null }), row({ rating: undefined })])).toBeNull();
  });

  it('never counts a missing rating as "below"', () => {
    expect(isBelowRatingFloor(row({ rating: undefined }), 50)).toBe(false);
    expect(isBelowRatingFloor(row({ rating: null }), 50)).toBe(false);
  });
});

describe('applyRosterFilters', () => {
  const rows = [
    row({ character_key: 'tank', role: 'tank', class: 'warrior', verified: true }),
    row({ character_key: 'healer', role: 'healer', class: 'priest', verified: true }),
    row({ character_key: 'unverified-dps', role: 'dps', class: 'mage', verified: false }),
  ];

  it('combines filters as AND', () => {
    const result = applyRosterFilters(
      rows,
      { ...DEFAULT_ROSTER_FILTERS, role: 'healer', classFilter: 'priest' },
      null,
    );
    expect(result.map((r) => r.character_key)).toEqual(['healer']);
  });

  it('filters verified-only', () => {
    const result = applyRosterFilters(rows, { ...DEFAULT_ROSTER_FILTERS, verifiedOnly: true }, null);
    expect(result.map((r) => r.character_key)).toEqual(['tank', 'healer']);
  });

  it('filters below-rating-floor', () => {
    const withRatings = [
      row({
        character_key: 'low',
        rating: {
          overall: 10,
          output: 0,
          survival: 0,
          mechanics: 0,
          utility: 0,
          preparation: 0,
          activity: 0,
        },
      }),
      row({
        character_key: 'high',
        rating: {
          overall: 90,
          output: 0,
          survival: 0,
          mechanics: 0,
          utility: 0,
          preparation: 0,
          activity: 0,
        },
      }),
    ];
    const result = applyRosterFilters(withRatings, { ...DEFAULT_ROSTER_FILTERS, belowFloorOnly: true }, 50);
    expect(result.map((r) => r.character_key)).toEqual(['low']);
  });
});

describe('sortRoster', () => {
  it('sorts by item level descending', () => {
    const rows = [row({ character_key: 'a', item_level: 55 }), row({ character_key: 'b', item_level: 70 })];
    expect(sortRoster(rows, 'ilvl').map((r) => r.character_key)).toEqual(['b', 'a']);
  });

  it('sorts by rating descending', () => {
    const rows = [
      row({
        character_key: 'a',
        rating: {
          overall: 40,
          output: 0,
          survival: 0,
          mechanics: 0,
          utility: 0,
          preparation: 0,
          activity: 0,
        },
      }),
      row({
        character_key: 'b',
        rating: {
          overall: 90,
          output: 0,
          survival: 0,
          mechanics: 0,
          utility: 0,
          preparation: 0,
          activity: 0,
        },
      }),
    ];
    expect(sortRoster(rows, 'rating').map((r) => r.character_key)).toEqual(['b', 'a']);
  });

  it('sorts by attendance ratio descending', () => {
    const rows = [
      row({ character_key: 'a', attendance: { present: 2, nights: 8 } }),
      row({ character_key: 'b', attendance: { present: 7, nights: 8 } }),
    ];
    expect(sortRoster(rows, 'attendance').map((r) => r.character_key)).toEqual(['b', 'a']);
  });

  it('sorts rank leader > officer > member', () => {
    const rows = [
      row({ character_key: 'member', rank: 'member' }),
      row({ character_key: 'leader', rank: 'leader' }),
      row({ character_key: 'officer', rank: 'officer' }),
    ];
    expect(sortRoster(rows, 'rank').map((r) => r.character_key)).toEqual(['leader', 'officer', 'member']);
  });

  it('never mutates the input array', () => {
    const rows = [row({ character_key: 'a', item_level: 55 }), row({ character_key: 'b', item_level: 70 })];
    const original = [...rows];
    sortRoster(rows, 'ilvl');
    expect(rows).toEqual(original);
  });
});

describe('pinOwnRowFirst', () => {
  it('pins the viewer’s own row first, stable otherwise', () => {
    const rows = [row({ character_key: 'a' }), row({ character_key: 'mine' }), row({ character_key: 'b' })];
    expect(pinOwnRowFirst(rows, 'mine').map((r) => r.character_key)).toEqual(['mine', 'a', 'b']);
  });

  it('is a no-op for a null viewer key', () => {
    const rows = [row({ character_key: 'a' }), row({ character_key: 'b' })];
    expect(pinOwnRowFirst(rows, null)).toEqual(rows);
  });
});

describe('orderRosterForTab', () => {
  it('always leads with unverified rows, then filters/sorts/pins the rest', () => {
    const rows = [
      row({ character_key: 'verified-low', verified: true, item_level: 55 }),
      row({ character_key: 'unverified', verified: false }),
      row({ character_key: 'mine', verified: true, item_level: 60 }),
      row({ character_key: 'verified-high', verified: true, item_level: 70 }),
    ];
    const ordered = orderRosterForTab(rows, {
      filters: DEFAULT_ROSTER_FILTERS,
      floor: null,
      sortKey: 'ilvl',
      myCharacterKey: 'mine',
    });
    expect(ordered.map((r) => r.character_key)).toEqual([
      'unverified',
      'mine',
      'verified-high',
      'verified-low',
    ]);
  });
});

describe('altOfMainName', () => {
  it('tags the lower-item-level row of a shared account as "alt of" the higher one', () => {
    const rows = [
      row({ character_key: 'main', name: 'Kraggor', account_key: 'u:1', item_level: 68 }),
      row({ character_key: 'alt', name: 'Grimtotem', account_key: 'u:1', item_level: 58 }),
    ];
    const map = altOfMainName(rows);
    expect(map.get('alt')).toBe('Kraggor');
    expect(map.has('main')).toBe(false);
  });

  it('tags nobody for a solo account', () => {
    const rows = [row({ character_key: 'solo', account_key: 'u:2' })];
    expect(altOfMainName(rows).size).toBe(0);
  });

  it('ignores rows with no account_key', () => {
    const rows = [
      row({ character_key: 'a', account_key: undefined }),
      row({ character_key: 'b', account_key: undefined }),
    ];
    expect(altOfMainName(rows).size).toBe(0);
  });
});
