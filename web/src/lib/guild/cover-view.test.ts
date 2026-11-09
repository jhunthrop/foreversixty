// web/src/lib/guild/cover-view.test.ts
// The cover grid's view: the gaps first, each cell's tone, the count and the icon url.
import { describe, expect, it } from 'vitest';
import type { GuildRaidCoverRow } from './api';
import { coverCount, coverIconUrl, coverSigned, coverThinnestFirst, coverTone } from './cover-view';

function row(over: Partial<GuildRaidCoverRow>): GuildRaidCoverRow {
  return {
    key: 'fortitude',
    name: 'Power Word: Fortitude',
    kind: 'buff',
    spell_id: 10938,
    icon: 'spell_holy_wordfortitude',
    need: 1,
    per: 'raid',
    providers_in: 0,
    providers_late: 0,
    providers_bench: 0,
    likely: 0,
    state: 'missing',
    ...over,
  };
}

describe('coverThinnestFirst', () => {
  it('lists nobody first, then too few, then covered, the fewest signed first within each', () => {
    const rows = [
      row({ key: 'kings', providers_in: 3 }),
      row({ key: 'motw', providers_in: 1 }),
      row({ key: 'lotp' }),
      row({ key: 'trueshot', providers_in: 1, likely: 1 }),
      row({ key: 'sunder', kind: 'debuff' }),
    ];
    expect(coverThinnestFirst(rows).map((r) => r.key)).toEqual(['lotp', 'sunder', 'trueshot', 'motw', 'kings']);
  });

  it('reads nothing as an empty list and never reorders its input', () => {
    expect(coverThinnestFirst(null)).toEqual([]);
    expect(coverThinnestFirst(undefined)).toEqual([]);
    const rows = [row({ key: 'a', providers_in: 1 }), row({ key: 'b' })];
    coverThinnestFirst(rows);
    expect(rows.map((r) => r.key)).toEqual(['a', 'b']);
  });
});

describe('coverTone', () => {
  it('is ok when In and Late meet the need', () => {
    expect(coverTone(row({ providers_in: 1, providers_late: 1, need: 2, state: 'ok' }))).toBe('ok');
  });

  it('warns when the need is met only on a talent the sign-up cannot show', () => {
    expect(coverTone(row({ providers_in: 1, likely: 1, state: 'ok' }))).toBe('warn');
    expect(coverTone(row({ providers_in: 2, likely: 1, state: 'ok' }))).toBe('ok');
  });

  it('warns when some are signed but too few', () => {
    expect(coverTone(row({ providers_in: 1, need: 3, state: 'short' }))).toBe('warn');
  });

  it('misses at 0, whatever the bench holds', () => {
    expect(coverTone(row({ providers_bench: 2 }))).toBe('miss');
  });
});

describe('coverCount and coverSigned', () => {
  it('counts In and Late, never Bench', () => {
    const r = row({ providers_in: 2, providers_late: 1, providers_bench: 4, need: 3 });
    expect(coverSigned(r)).toBe(3);
    expect(coverCount(r)).toBe('3 of 3');
  });
});

describe('coverIconUrl', () => {
  it('is the build icon file SpellIcon draws', () => {
    expect(coverIconUrl('1.60.1.70790', row({}))).toBe(
      '/data/1.60.1.70790/icons/spell_holy_wordfortitude.webp',
    );
  });

  it('is null for the placeholder or no stem', () => {
    expect(coverIconUrl('1.60.1.70790', row({ icon: 'inv_misc_questionmark' }))).toBeNull();
    expect(coverIconUrl('1.60.1.70790', row({ icon: '' }))).toBeNull();
  });
});
