// web/src/lib/rankings/rank-row-view.ts
// A ranking row as the shared RankRow draws it (logs spec 2026-11-07 section 7): the rank in its
// percentile colour (the WCL scale, floored: rankColor), the character's name and link in its
// class colour, the guild under it with the talent split (the spec on a board of several
// classes), the bar of its value over the board's top, the figures (the value, which opens the
// kill in its report, the percentile, Executed, Length) and the date. The row carries no link strip:
// a compact list beside the rail has no room for one, and the value is the way into the report.
import { characterHref, parseCharacterKey, splitUnitName } from '../characters';
import { formatAmount } from '../report/format';
import { executionHref, executionLabel, executionTitle } from '../sim/execution';
import type { RankFigure, RankGap, RankLink, RankNote } from '../ui/rank-list';
import type { RankingRow } from './api';
import { rankingsCopy } from './copy';
import { moderationLine, rankColor, rankPercentile, rankingDate } from './row-view';

const SECONDS_PER_MINUTE = 60;
const MS_PER_SECOND = 1000;

export interface RankRowView {
  rank: string;
  /** The percentile colour on the rank; left out where no legend explains it (the landing). */
  rankColor?: string;
  name: string;
  href: string;
  classSlug: string;
  sub: string;
  fraction: number;
  figures: RankFigure[];
  gap: RankGap;
  links: RankLink[];
  note?: RankNote;
}

/** "2:41": a kill's length in whole seconds. */
export function killLength(durationMs: number): string {
  const seconds = Math.max(0, Math.round(durationMs / MS_PER_SECOND));
  return `${Math.floor(seconds / SECONDS_PER_MINUTE)}:${String(seconds % SECONDS_PER_MINUTE).padStart(2, '0')}`;
}

/** The board a row is read on: its size, its top value and the value's unit. `multiClass`
 *  puts the spec and class under the name, where the crest alone would not say the spec. */
export interface RankBoardView {
  total: number;
  top: number;
  unit: string;
  multiClass?: boolean;
}

/** A rank's percentile on its board, floored to a whole number as the colour reads it: 100, 99, 95. */
export function percentileLabel(rank: number, total: number): string {
  return String(Math.max(0, Math.floor(rankPercentile(rank, total))));
}

export function rankRowView(row: RankingRow, board: RankBoardView): RankRowView {
  const name = splitUnitName(row.player.name).name;
  const key = parseCharacterKey(row.player.key);
  const moderation = moderationLine(row.state);
  return {
    rank: String(row.rank),
    rankColor: rankColor(row.rank, board.total),
    name,
    href: key === null ? reportHref(row) : characterHref(key.region, key.ruleset, row.player.name),
    classSlug: row.player.class.toLowerCase(),
    sub: [
      row.guild?.name,
      // A board of several classes names the spec; one class's board shows the talent split.
      board.multiClass === true ? `${row.player.spec} ${row.player.class}` : row.talent_split,
    ]
      .filter((part): part is string => part !== undefined && part !== '')
      .join(' · '),
    fraction: board.top > 0 ? row.value / board.top : 0,
    figures: [
      {
        value: formatAmount(Math.round(row.value)),
        label: board.unit,
        href: reportHref(row),
        title: rankingsCopy.openReport,
        testid: 'ranking-value',
      },
      {
        value: percentileLabel(row.rank, board.total),
        label: rankingsCopy.percentileLabel,
        title: rankingsCopy.percentileLabel,
        bare: true,
        color: rankColor(row.rank, board.total),
        testid: 'ranking-percentile',
      },
      {
        value: executionLabel(row.execution_score ?? null),
        label: 'Executed',
        title: executionTitle(row.execution_score ?? null),
        // A scored kill opens compare mode on that fight; an unscored one says why on hover.
        ...(row.execution_score === null || row.execution_score === undefined
          ? {}
          : { href: executionHref(row.report_id, row.fight_index) }),
        testid: 'ranking-execution',
      },
      { value: killLength(row.duration_ms), label: 'Length' },
    ],
    gap: { text: rankingDate(row.fought_at), label: 'Date' },
    links: [],
    ...(moderation === null ? {} : { note: { text: moderation } }),
  };
}

/** The kill in its report (the report alone when the row names no fight). */
export function reportHref(kill: { report_id: string; fight_index?: number }): string {
  return kill.fight_index === undefined
    ? `/reports/${kill.report_id}`
    : `/reports/${kill.report_id}?fight=${kill.fight_index}`;
}
