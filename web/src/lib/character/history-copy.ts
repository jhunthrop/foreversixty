// web/src/lib/character/history-copy.ts
// Every string of experience spec 2026-10-24-character-history section 6 for the History
// section. Functions take the data and return the string; no number or date is typed into
// prose. The shared retry word is `uiCopy.retry`; tier names are `tierLabel`'s.
import { scoreCopy } from '../score/copy';
import type { HistoryVisibility } from './history-api';

const DETAIL_SEPARATOR = ' · ';

function count(value: number, one: string, many: string): string {
  return `${value} ${value === 1 ? one : many}`;
}

export type ListKind = 'fights' | 'kills';

export const historyCopy = {
  heading: 'History',
  caption: 'Includes your private and guild reports. Only you see this section.',

  scoreHeading: 'Score over time',
  /** The control row: the collapse rule, and with rings (private reports) the key to them. */
  scoreNote: (hasRings: boolean): string =>
    `One point a day: the day's last rated fight.${hasRings ? ' Rings are private reports.' : ''}`,
  parseHeading: 'Parse by boss',
  /** The boss picker's accessible name. */
  pickerLabel: 'Boss',

  /** The pill's accessible name: "Golemagg, 5 kills". */
  pillName: (boss: string, kills: number): string => `${boss}, ${count(kills, 'kill', 'kills')}`,
  /** A boss killed on two difficulties: "Golemagg · Hard". */
  pillBoss: (boss: string, difficultyWord: string | null): string =>
    difficultyWord === null ? boss : `${boss}${DETAIL_SEPARATOR}${difficultyWord}`,

  /** Under the score: "11 Dec 2026 · 2 fights". */
  scoreDetail: (date: string, fights: number): string =>
    `${date}${DETAIL_SEPARATOR}${count(fights, 'fight', 'fights')}`,
  /** Under the parse: "11 Dec 2026 · kill 5 of 5". */
  parseDetail: (date: string, position: number, total: number): string =>
    `${date}${DETAIL_SEPARATOR}kill ${position} of ${total}`,
  /** The score readout as one line: "1830 Epic · 11 Dec 2026 · 2 fights". */
  scoreReadoutText: (score: number, tierWord: string, detail: string): string =>
    `${[String(score), tierWord].filter((part) => part !== '').join(' ')}${DETAIL_SEPARATOR}${detail}`,
  parseWord: 'parse',

  /** "Up 1090 since 12 Oct": the card's own sentence, for both charts. */
  netChange: scoreCopy.historyChange,

  /** Under the chart, for a gap too narrow for its label: "Gap 20 Jul to 9 Aug: No rated fights". */
  gapNote: (from: string, to: string, reason: string): string => `Gap ${from} to ${to}: ${reason}`,

  step: {
    score: { earlier: 'Earlier point', later: 'Later point' },
    parse: { earlier: 'Earlier kill', later: 'Later kill' },
  },
  scorePlotLabel: 'Score chart. Use the arrow keys to move between points.',
  parsePlotLabel: (boss: string): string =>
    `Parse chart for ${boss}. Use the arrow keys to move between kills.`,

  srScoreCaption: 'Forever Score by date, oldest first',
  srParseCaption: 'Parse by kill, oldest first',
  srColumns: { date: 'Date', score: 'Score', parse: 'Parse' },
  srGapRow: (from: string, to: string, reason: string): string => `Gap from ${from} to ${to}: ${reason}`,

  scoreListLabel: 'Fights behind the score',
  parseListLabel: (boss: string): string => `Kills of ${boss}`,

  /** The word after a boss name in a list row; nothing for a public report. */
  visibilityWord: (visibility: HistoryVisibility): string | null =>
    visibility === 'private' ? 'Private' : visibility === 'guild' ? 'Guild' : null,

  showAll: (total: number, kind: ListKind): string =>
    `Show all ${kind === 'fights' ? count(total, 'fight', 'fights') : count(total, 'kill', 'kills')}`,
  showFewer: 'Show fewer',

  loadingSr: 'Reading your history…',
  loadFailed: 'Your history did not load.',
  empty: 'No history yet. Each rated fight you upload adds a point.',
  scoreOnePoint: (date: string): string => `First score recorded ${date}. A line needs a second point.`,
  parseOneKill: 'One kill so far. A line needs two.',
  parseNoKills: 'No kills on record yet. Parse charts start with your first kill.',

  /** One sentence under the score card; `freeOwnerLink` is its closing link to /premium. */
  freeOwnerSentence: 'History charts draw your score and parses over time.',
  freeOwnerLink: 'See plans',
} as const;
