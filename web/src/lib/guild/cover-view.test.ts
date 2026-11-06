// web/src/lib/guild/cover-view.test.ts
// The cover grid's view: the two sections in order, each cell's tone, the count and the icon url.
import { describe, expect, it } from 'vitest';
import type { GuildRaidCoverRow } from './api';
import { coverCount, coverIconUrl, coverSections, coverSigned, coverTone } from './cover-view';

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

describe('coverSections', () => {
  it('puts buffs before debuffs and keeps the API order inside each', () => {
    const rows = [
      row({ key: 'sunder_armor', kind: 'debuff' }),
      row({ key: 'windfury_totem' }),
      row({ key: 'faerie_fire', kind: 'debuff' }),
      row({ key: 'fortitude' }),
    ];
    const sections = coverSections(rows);
    expect(sections.map((s) => s.kind)).toEqual(['buff', 'debuff']);
    expect(sections[0].rows.map((r) => r.key)).toEqual(['windfury_totem', 'fortitude']);
    expect(sections[1].rows.map((r) => r.key)).toEqual(['sunder_armor', 'faerie_fire']);
  });

  it('has nothing for a dungeon or an older fixture', () => {
    expect(coverSections(null)).toEqual([]);
    expect(coverSections(undefined)).toEqual([]);
    expect(coverSections([row({ kind: 'debuff' })]).map((s) => s.kind)).toEqual(['debuff']);
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
