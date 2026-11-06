// web/src/lib/rankings/rank-row-view.ts
// A ranking row as the shared RankRow draws it (logs spec 2026-11-07 section 7): the rank in its
// percentile colour (the WCL scale, floored: rankColor), the character's name and link in its
// class colour, `{guild} · {spec} {class} · {size}-man` under it, the bar of its value over the
// board's top, the figures (the value, Item level, Executed, Length), the date and the links
// (Report gold, Build). An item level the row does not carry yet reads a dash (G12).
import { characterHref, parseCharacterKey, splitUnitName } from '../characters';
import { formatAmount } from '../report/format';
import { executionHref, executionLabel, executionTitle } from '../sim/execution';
import type { RankFigure, RankGap, RankLink, RankNote } from '../ui/rank-list';
import type { RankingRow } from './api';
import { rankingsCopy } from './copy';
import { moderationLine, rankColor, rankingDate } from './row-view';

const SECONDS_PER_MINUTE = 60;
const MS_PER_SECOND = 1000;
const UNKNOWN = '–';

export interface RankRowView {
  rank: string;
  rankColor: string;
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

export function rankRowView(
  row: RankingRow & { item_level?: number | null },
  board: { total: number; top: number; unit: string; now: Date },
): RankRowView {
  const name = splitUnitName(row.player.name).name;
  const key = parseCharacterKey(row.player.key);
  const moderation = moderationLine(row.state);
  return {
    rank: String(row.rank),
    rankColor: rankColor(row.rank, board.total),
    name,
    href:
      key === null
        ? `/reports/${row.report_id}?fight=${row.fight_index}`
        : characterHref(key.region, key.ruleset, row.player.name),
    classSlug: row.player.class.toLowerCase(),
    sub: [row.guild?.name, `${row.player.spec} ${row.player.class}`, rankingsCopy.sizeMan(row.size)]
      .filter((part): part is string => part !== undefined && part !== '')
      .join(' · '),
    fraction: board.top > 0 ? row.value / board.top : 0,
    figures: [
      {
        value: formatAmount(Math.round(row.value)),
        unit: board.unit,
        label: board.unit,
        testid: 'ranking-value',
      },
      {
        value: row.item_level === undefined || row.item_level === null ? UNKNOWN : String(row.item_level),
        unit: '',
        label: 'Item level',
      },
      {
        value: executionLabel(row.execution_score),
        unit: '',
        label: 'Executed',
        title: executionTitle(row.execution_score),
        // A scored kill opens compare mode on that fight; an unscored one says why on hover.
        ...(row.execution_score === null ? {} : { href: executionHref(row.report_id, row.fight_index) }),
        testid: 'ranking-execution',
      },
      { value: killLength(row.duration_ms), unit: '', label: 'Length' },
    ],
    gap: { text: rankingDate(row.fought_at, board.now), label: 'Date' },
    links: [
      {
        label: 'Report',
        href: `/reports/${row.report_id}?fight=${row.fight_index}`,
        primary: true,
        testid: 'ranking-report',
      },
      ...(row.build_id === undefined
        ? []
        : [{ label: 'Build', href: `/b/${row.build_id}`, testid: 'ranking-build' }]),
    ],
    ...(moderation === null ? {} : { note: { text: moderation } }),
  };
}
