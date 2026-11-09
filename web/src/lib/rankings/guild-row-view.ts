// web/src/lib/rankings/guild-row-view.ts
// A guild board's row as the shared RankRow draws it: the guild's name and link, its region and
// ruleset under it, the bar against the board's best, the one figure the kind ranks by (bosses
// killed, the kill time, the deaths), the date and the Report link. The best is the most bosses
// on Progress and the least time or deaths on Speed and Execution.
import { guildHref, rulesetLabel } from '../characters';
import type { RankFigure, RankGap, RankLink } from '../ui/rank-list';
import type { GuildRankingRow } from './api';
import { homeTopGuildsCopy, rankingsCopy } from './copy';
import { killLength } from './rank-row-view';
import { rankingDate } from './row-view';
import type { GuildKind } from './url';

export interface GuildRowView {
  rank: string;
  name: string;
  href: string;
  sub: string;
  fraction: number;
  figures: RankFigure[];
  gap: RankGap;
  links: RankLink[];
}

const FIGURE_LABEL: Record<GuildKind, string> = {
  progress: 'Bosses',
  speed: 'Kill time',
  execution: 'Deaths',
};

export function guildFigureLabel(kind: GuildKind): string {
  return FIGURE_LABEL[kind];
}

function figureText(kind: GuildKind, value: number): string {
  if (kind === 'progress') return homeTopGuildsCopy.bosses(Math.round(value));
  if (kind === 'speed') return killLength(value);
  return rankingsCopy.deaths(Math.round(value));
}

/** The bar: the row against the board's best, 1 for the best. Fewer is better on speed and deaths
 *  (a death count can be 0, so both sides count one more). */
function fractionOf(kind: GuildKind, value: number, best: number): number {
  if (kind === 'progress') return best > 0 ? value / best : 0;
  if (kind === 'speed') return value > 0 ? best / value : 0;
  return (best + 1) / (value + 1);
}

export function guildRowView(row: GuildRankingRow, board: { kind: GuildKind; best: number }): GuildRowView {
  const { guild } = row;
  return {
    rank: String(row.rank),
    name: guild.name,
    href: guildHref(guild.region, guild.ruleset, guild.name),
    sub: `${guild.region.toUpperCase()} · ${rulesetLabel(guild.ruleset)}`,
    fraction: fractionOf(board.kind, row.value, board.best),
    figures: [
      {
        value: figureText(board.kind, row.value),
        label: FIGURE_LABEL[board.kind],
        testid: 'guild-ranking-value',
      },
    ],
    gap: { text: rankingDate(row.fought_at), label: 'Date' },
    links:
      row.report_id === undefined || row.report_id === ''
        ? []
        : [
            {
              label: 'Report',
              href: reportHref(row.report_id, row.fight_index),
              primary: true,
              testid: 'guild-ranking-report',
            },
          ],
  };
}

function reportHref(reportId: string, fightIndex: number | undefined): string {
  return fightIndex === undefined ? `/reports/${reportId}` : `/reports/${reportId}?fight=${fightIndex}`;
}
