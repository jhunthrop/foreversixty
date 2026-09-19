// web/src/lib/rankings/url.test.ts
import { describe, expect, it } from 'vitest';
import {
  GUILD_KINDS,
  RANKING_METRICS,
  defaultRankingsState,
  parseRankingsState,
  rankingsSearch,
} from './url';

describe('the rankings URL state', () => {
  it('offers the three character metrics and the three guild kinds spec section 6 names', () => {
    expect(RANKING_METRICS.map((metric) => metric.id)).toEqual(['dps', 'hps', 'damage_taken', 'execution']);
    expect(RANKING_METRICS.map((metric) => metric.label)).toEqual([
      'Damage',
      'Healing',
      'Damage taken',
      'Execution',
    ]);
    expect(GUILD_KINDS.map((kind) => kind.id)).toEqual(['progress', 'speed', 'execution']);
  });

  it('defaults to character damage, all time, page one', () => {
    expect(defaultRankingsState()).toEqual({
      board: 'character',
      metric: 'dps',
      kind: 'progress',
      spec: '',
      class: '',
      phase: '',
      region: '',
      ruleset: '',
      faction: '',
      since: '',
      page: 1,
    });
  });

  it('parses every filter and rejects a value that is not one of ours', () => {
    const state = parseRankingsState(
      '?board=guild&kind=speed&metric=hps&spec=Fury&class=Warrior&phase=raids-1&region=eu&ruleset=hardcore&faction=horde&since=today&page=3',
    );
    expect(state.board).toBe('guild');
    expect(state.kind).toBe('speed');
    expect(state.ruleset).toBe('hardcore');
    expect(state.region).toBe('eu');
    expect(state.page).toBe(3);

    const bad = parseRankingsState('?board=nonsense&ruleset=nightslayer&region=uk&page=0');
    expect(bad.board).toBe('character');
    expect(bad.ruleset).toBe('');
    expect(bad.region).toBe('');
    expect(bad.page).toBe(1);
  });

  it('serialises only what is set, and round-trips', () => {
    expect(rankingsSearch(defaultRankingsState())).toBe('');
    const state = { ...defaultRankingsState(), ruleset: 'pvp', metric: 'hps', page: 2 };
    expect(rankingsSearch(state)).toBe('?metric=hps&ruleset=pvp&page=2');
    expect(parseRankingsState(rankingsSearch(state))).toEqual(state);
  });
});

describe('the execution metric', () => {
  it('is offered as a sort order', () => {
    expect(RANKING_METRICS.map((metric) => metric.id)).toContain('execution');
  });

  it('survives a round trip through the query string', () => {
    expect(
      parseRankingsState(rankingsSearch({ ...defaultRankingsState(), metric: 'execution' })).metric,
    ).toBe('execution');
  });

  it('is still refused when the query names something else', () => {
    expect(parseRankingsState('?metric=execution-but-not').metric).toBe('dps');
  });
});
