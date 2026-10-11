// web/src/lib/score/copy.ts
// Every visible string the Forever Score surfaces use: the character page's score block and
// "In game" section, and the /ratings explainer. Reference voice (design/DESIGN-SYSTEM.md):
// what the number is and what moves it, never a pitch.
//
// The score line, the five component labels, the tier names, "as of" and the no-logs line
// are the addon's own (addon/ForeverSixty/Locale.lua), read from the shared fixture
// docs/contracts/fixtures/score-labels.json that addon/tests/score_labels_spec.lua pins.
import labels from '../../../../docs/contracts/fixtures/score-labels.json';
import type { ScoreComponentKey } from './score';

/** What moves each component, in plain words (api/openapi.yaml's `ForeverScore` descriptions,
 *  rewritten for a player: "how your parses rank", never "percentile"). */
const MOVES: Readonly<Record<string, string>> = {
  gear: 'How close the gear you wear is to your spec’s best-in-slot set at level 60. Every upgrade for your spec moves it up.',
  experience:
    '25 for each different dungeon cleared and 40 for each different raid boss killed, across your main and the alts the site has linked to it.',
  parses:
    'How your parses rank: your best parse on each boss against the same spec, averaged over the bosses you have killed.',
  attendance: 'How many of your guild’s raid nights in the last 90 days you were in the log for.',
  discipline:
    'How many of your logged fights you went into with a flask, food, and enchants on chest, feet, back and main hand.',
};

export const scoreCopy = {
  heading: 'Forever Score',
  /** "Forever Score 1830", the addon's score line. */
  line: (score: number): string => labels.line.replace('%d', String(score)),
  /** "as of 4 Nov". */
  asOf: (date: string): string => labels.asOf.replace('%s', date),
  /** Where the number would be, for a character with no logs on the site. */
  noLogs: labels.noLogs,
  /** The five parts, in the order the addon's tooltip lists them. */
  components: labels.components.map((component) => ({
    key: component.key as ScoreComponentKey,
    label: component.label,
    moves: MOVES[component.key] ?? '',
  })),
  /** "412 of 500" beside a component's bar. */
  ofMax: (points: number, max: number): string => `${points} of ${max}`,
  explainLink: 'What the score counts',
  explainHref: '/ratings#forever-score',

  inGameHeading: 'In game',
  inGameScoreLead: 'The addon shows the same score on this character’s tooltip in game:',
  inGameNoScore: 'The addon shows no score on this character’s tooltip until the site has a log of it.',
  inGamePlannerLead: 'To put this character into the planner: in game, /fs, Copy my character, then',
  inGamePasteLink: 'paste it here',
  inGamePasteHref: '/setup#paste',
  /** "Last seen in game 4 Nov 14:03": when the site last saw the owner's addon export. */
  lastSeen: (when: string): string => `Last seen in game ${when}`,

  explainerHeading: 'Forever Score',
  explainerWhat:
    'The Forever Score is one number for a character, from 0 to 2100, made of five parts that add up to it. It comes only from the logs the site has: fights in public reports, and the gear those logs show. The addon shows the same number on the character’s tooltip in game, dated with the day of the newest fight it counts. A character with no logs here has no score, not a score of zero.',
  explainerExample: 'How a score reads, here and in the addon’s tooltip:',
  explainerPartsHeading: 'What moves each part',
  /** "up to 500" beside a component's label. */
  upTo: (max: number): string => `up to ${max}`,
  explainerTiersHeading: 'Tiers',
  tierColumn: 'Tier',
  rangeColumn: 'Score',
  tierBelow: (next: number): string => `below ${next}`,
  tierBetween: (from: number, to: number): string => `${from} to ${to}`,
  tierFrom: (from: number): string => `${from} and up`,
  tiersNote:
    'The tiers take the colours of item quality in game, as the addon colours the score. Their lines are set from the scores players have and may move.',
} as const;
