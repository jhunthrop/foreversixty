// web/src/lib/rankings/own-rank.test.ts
import { describe, expect, it, vi } from 'vitest';
import type { MeCharacter } from '../account/api';
import type { RankingRow, RankingsPage } from './api';
import { ownParsesAction, ownRankAnswer, ownRankQuery, ownRankSpec, type OwnRankScope } from './own-rank';
import { defaultRankingsState } from './url';

const SCOPE: OwnRankScope = {
  encounterSlug: 'onyxia',
  encounterName: 'Onyxia',
  state: { ...defaultRankingsState(), ruleset: 'normal', phase: 'raids-1' },
  own: { class: 'warrior', spec: 'Fury' },
};
const CHARACTER: MeCharacter = {
  key: 'us/normal/obnoxious-yell',
  region: 'us',
  ruleset: 'normal',
  name: 'Obnoxious Yell',
  class: 'Warrior',
  spec: 'Fury',
};
const PAGE: RankingsPage = {
  rows: [],
  total: 212,
  page: 1,
  per_page: 100,
  updated_at: '2026-12-13T00:00:00Z',
  deduped: true,
  character_rank: { rank: 14, row: { player: { key: CHARACTER.key } } as RankingRow },
};

describe('ownRankQuery', () => {
  it('asks for the character in its own cohort, one best kill per character', () => {
    expect(ownRankQuery(SCOPE, CHARACTER.key)).toMatchObject({
      encounter: 'onyxia',
      class: 'warrior',
      spec: 'Fury',
      ruleset: 'normal',
      phase: 'raids-1',
      distinct: 'character',
      character: 'us/normal/obnoxious-yell',
    });
  });
});

describe('ownRankAnswer', () => {
  it('reads the rank as the figure, in the board colour, with the cohort and encounter', () => {
    const answer = ownRankAnswer(PAGE, SCOPE);
    expect(answer).toEqual({
      kind: 'figures',
      label: 'Your rank',
      figures: [{ value: '14th', unit: 'of 212 Fury Warriors', color: 'var(--color-parse-purple)' }],
      note: { text: 'On Onyxia, Normal ruleset' },
    });
  });

  it('says every ruleset when the board is not narrowed to one', () => {
    const answer = ownRankAnswer(PAGE, { ...SCOPE, state: { ...SCOPE.state, ruleset: '' } });
    expect(answer).toMatchObject({ note: { text: 'On Onyxia, all rulesets' } });
  });

  it('says so when the character has no ranked kill in the cohort', () => {
    expect(ownRankAnswer({ ...PAGE, character_rank: null }, SCOPE)).toEqual({
      kind: 'sentence',
      text: 'Not ranked yet on Onyxia. Upload a log of a kill to get on the board.',
    });
  });

  it('fails rather than show a per-kill rank as a place among characters', () => {
    expect(() => ownRankAnswer({ ...PAGE, deduped: undefined }, SCOPE)).toThrow();
  });
});

describe('ownRankSpec', () => {
  it('loads through the injected read, with the character key', async () => {
    const read = vi.fn().mockResolvedValue(PAGE);
    const answer = await ownRankSpec(SCOPE, read).load(CHARACTER);
    expect(read).toHaveBeenCalledWith(expect.objectContaining({ character: CHARACTER.key }));
    expect(answer.kind).toBe('figures');
  });

  it('has nothing to say for a character with no class', async () => {
    const read = vi.fn();
    const answer = await ownRankSpec({ ...SCOPE, own: { class: '', spec: '' } }, read).load(CHARACTER);
    expect(answer).toEqual({ kind: 'none' });
    expect(read).not.toHaveBeenCalled();
  });

  it('words its failure in the page voice', () => {
    expect(ownRankSpec(SCOPE).strings.failed).toBe('Your rank did not load.');
  });
});

describe('ownParsesAction', () => {
  it('opens the character page', () => {
    expect(ownParsesAction(CHARACTER)).toEqual({
      label: 'See your parses →',
      href: '/character/us/normal/obnoxious-yell',
    });
  });
});
