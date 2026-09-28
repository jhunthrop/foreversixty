// web/src/lib/guild/roster.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildRosterRow } from './api';
import { orderRoster, unverifiedRosterCount } from './roster';

function row(character_key: string, verified: boolean): GuildRosterRow {
  return {
    character_key,
    region: 'us',
    ruleset: 'hardcore',
    name: character_key,
    rank: 'member',
    verified,
    logged_recently: false,
    consent: 'roster',
    may_remove: false,
  };
}

describe('orderRoster', () => {
  it('puts unverified rows first, stable within each group', () => {
    const rows = [row('a', true), row('b', false), row('c', true), row('d', false)];
    expect(orderRoster(rows).map((r) => r.character_key)).toEqual(['b', 'd', 'a', 'c']);
  });

  it('leaves an all-verified roster in its original order', () => {
    const rows = [row('a', true), row('b', true)];
    expect(orderRoster(rows).map((r) => r.character_key)).toEqual(['a', 'b']);
  });

  it('leaves an all-unverified roster in its original order', () => {
    const rows = [row('a', false), row('b', false)];
    expect(orderRoster(rows).map((r) => r.character_key)).toEqual(['a', 'b']);
  });

  it('does not mutate the input array', () => {
    const rows = [row('a', true), row('b', false)];
    const copy = [...rows];
    orderRoster(rows);
    expect(rows).toEqual(copy);
  });

  it('handles an empty roster', () => {
    expect(orderRoster([])).toEqual([]);
  });
});

describe('unverifiedRosterCount', () => {
  it('counts unverified rows', () => {
    const rows = [row('a', true), row('b', false), row('c', false)];
    expect(unverifiedRosterCount(rows)).toBe(2);
  });

  it('is zero when every row is verified', () => {
    expect(unverifiedRosterCount([row('a', true)])).toBe(0);
  });

  it('is zero for an empty roster', () => {
    expect(unverifiedRosterCount([])).toBe(0);
  });
});
