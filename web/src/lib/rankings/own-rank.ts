// web/src/lib/rankings/own-rank.ts
// What the shared character header says on /rankings (design/specs/2026-10-11-rankings.md
// sections 1 and 6): "Your rank: 14th of 212 Fury Warriors, on Onyxia, Normal ruleset", read
// from the board's own cohort so the number is the one the list shows. The answer is always
// about the character's own class and spec, whichever board the page is looking at.
import type { MeCharacter } from '../account/api';
import { characterHref, rulesetLabel } from '../characters';
import { ordinal } from '../report/format';
import type { AnswerSpec, HeaderAction, HeaderAnswer } from '../character-header/types';
import { fetchRankings, type RankingsPage, type RankingsQuery } from './api';
import { rankingsCopy } from './copy';
import { rankColor } from './row-view';
import { cohortName } from './scope';
import { type RankingsDefaults, type RankingsState } from './url';

/** What the answer is read on: the encounter and the board's own filters. */
export interface OwnRankScope {
  encounterSlug: string;
  encounterName: string;
  state: Pick<RankingsState, 'metric' | 'phase' | 'region' | 'ruleset' | 'faction' | 'since'>;
  /** The character's own class and spec: the cohort the rank is among. */
  own: Pick<RankingsDefaults, 'class' | 'spec'>;
}

/** The request that places the character: its cohort, one best kill per character. */
export function ownRankQuery(scope: OwnRankScope, characterKey: string): RankingsQuery {
  const { state, own } = scope;
  return {
    encounter: scope.encounterSlug,
    metric: state.metric as RankingsQuery['metric'],
    class: own.class,
    spec: own.spec,
    phase: state.phase,
    region: state.region,
    ruleset: state.ruleset,
    faction: state.faction,
    since: state.since,
    distinct: 'character',
    character: characterKey,
  };
}

function rulesetScope(ruleset: string): string {
  return ruleset === '' ? 'all rulesets' : `${rulesetLabel(ruleset)} ruleset`;
}

/**
 * The header's answer from the board's response. An API that does not dedupe cannot place a
 * character among characters, so it is the failed state rather than a rank that would not match
 * the list.
 */
export function ownRankAnswer(page: RankingsPage, scope: OwnRankScope): HeaderAnswer {
  if (page.deduped !== true) throw new Error('the rankings API does not place characters yet');
  const placed = page.character_rank;
  if (placed === undefined || placed === null) {
    return { kind: 'sentence', text: rankingsCopy.answerUnranked(scope.encounterName) };
  }
  return {
    kind: 'figures',
    label: rankingsCopy.answerLabel,
    figures: [
      {
        value: ordinal(placed.rank),
        unit: rankingsCopy.answerUnit(page.total, cohortName(scope.own.class, scope.own.spec), true),
        color: rankColor(placed.rank, page.total),
      },
    ],
    note: { text: rankingsCopy.answerNote(scope.encounterName, rulesetScope(scope.state.ruleset)) },
  };
}

/** Holds the header's answer in its loading state until the board knows what it is aimed at. */
export const PENDING_ANSWER: AnswerSpec = {
  strings: { failed: rankingsCopy.answerFailed },
  load: () => new Promise<HeaderAnswer>(() => undefined),
};

/** The header's answer spec for a scope; `read` is injectable so a test needs no network. */
export function ownRankSpec(
  scope: OwnRankScope,
  read: (query: RankingsQuery) => Promise<RankingsPage> = fetchRankings,
): AnswerSpec {
  return {
    strings: { failed: rankingsCopy.answerFailed },
    load: async (character: MeCharacter) =>
      scope.own.class === '' || scope.encounterSlug === ''
        ? { kind: 'none' }
        : ownRankAnswer(await read(ownRankQuery(scope, character.key)), scope),
  };
}

/** The one action: the character's own page, where its parses are. */
export function ownParsesAction(character: MeCharacter): HeaderAction {
  return {
    label: rankingsCopy.action,
    href: characterHref(character.region, character.ruleset, character.name),
  };
}
