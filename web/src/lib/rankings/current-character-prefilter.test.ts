// web/src/lib/rankings/current-character-prefilter.test.ts
import { describe, expect, it } from 'vitest';
import { applyCurrentCharacterPrefilter, pinCurrentCharacterRow } from './current-character-prefilter';
import { parseRankingsState } from './url';
import type { CurrentCharacter } from '../current-character';
import type { RankingRow } from './api';

const ARMORY: CurrentCharacter = {
  source: 'armory',
  ref: 'us/normal/simfury',
  label: 'Simfury · Fury Warrior',
  classSlug: 'warrior',
  savedAt: '2026-09-25T00:00:00.000Z',
};

describe('applyCurrentCharacterPrefilter', () => {
  it('pre-fills class and ruleset from an armory pointer on a bare URL', () => {
    const state = applyCurrentCharacterPrefilter(parseRankingsState(''), '', ARMORY);
    expect(state.class).toBe('warrior');
    expect(state.ruleset).toBe('normal');
  });

  it('leaves an explicit ?class= or ?ruleset= alone', () => {
    const state = applyCurrentCharacterPrefilter(parseRankingsState('?class=mage'), '?class=mage', ARMORY);
    expect(state.class).toBe('mage');
  });

  it('is a no-op with no pointer', () => {
    const state = applyCurrentCharacterPrefilter(parseRankingsState(''), '', null);
    expect(state.class).toBe('');
    expect(state.ruleset).toBe('');
  });

  it('fills class only for a non-armory pointer (no known ruleset)', () => {
    const code: CurrentCharacter = { ...ARMORY, source: 'code', ref: 'FS1:1:warrior:orc:0/0/0:' };
    const state = applyCurrentCharacterPrefilter(parseRankingsState(''), '', code);
    expect(state.class).toBe('warrior');
    expect(state.ruleset).toBe('');
  });
});

const ROW = (rank: number, key: string): RankingRow => ({
  rank,
  player: { key, name: key.split('/')[2], class: 'warrior', spec: 'fury' },
  value: 100,
  size: 20,
  fought_at: '2026-09-25T00:00:00Z',
  duration_ms: 60_000,
  talent_split: '31/20/0',
  trinkets: [],
  buff_count: 0,
  report_id: 'r1',
  fight_index: 0,
  state: 'ok',
  execution_score: null,
});

describe('pinCurrentCharacterRow', () => {
  it('moves the current character row to the top when present in the page', () => {
    const rows = [ROW(1, 'us/normal/other'), ROW(2, 'us/normal/simfury'), ROW(3, 'us/normal/third')];
    const pinned = pinCurrentCharacterRow(rows, ARMORY);
    expect(pinned[0].player.key).toBe('us/normal/simfury');
    expect(pinned).toHaveLength(3);
  });

  it('leaves the order alone when the current character is not on this page', () => {
    const rows = [ROW(1, 'us/normal/other'), ROW(2, 'us/normal/third')];
    expect(pinCurrentCharacterRow(rows, ARMORY)).toEqual(rows);
  });

  it('leaves the order alone with no pointer', () => {
    const rows = [ROW(1, 'us/normal/other')];
    expect(pinCurrentCharacterRow(rows, null)).toEqual(rows);
  });
});
