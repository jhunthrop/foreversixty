// web/src/lib/chart/rate.ts
// The damage, healing and damage-taken chart's model (controls and charts spec D5, §B5.3): every
// player's 5 s average, the top 8 of N plus the always-drawn players, the raid median and a class
// median, and the head's raid figure, which is the sum of the figures the legend shows.
import { classColorVar } from '../report/format';
import type { TimeWindow } from '../report/window';
import type { RateMetric } from './copy';
import { chartCopy } from './copy';
import {
  type DeadSpan,
  median,
  medianSeries,
  roundedSum,
  selectDrawn,
  stopAtDeath,
  trailingAverage,
  windowRate,
} from './drawn';
import type { ChartSeries, PlayerEntry } from './types';

export const RAID_MEDIAN_ID = 'median:raid';
export const MEDIAN_DASH = [5, 4] as const;

export interface RatePlayer {
  guid: string;
  name: string;
  className?: string;
  role: PlayerEntry['role'];
  /** One value per second of the whole fight. */
  series: readonly number[];
  /** A colour for a unit with no class (an enemy): the palette's. */
  colour?: string;
}

export interface RateInput {
  metric: RateMetric;
  players: readonly RatePlayer[];
  window: TimeWindow;
  /** The Source player and you: always drawn. */
  pinned: ReadonlySet<string>;
  youGuid?: string;
  /** The Players control's choice; null keeps the top 8. */
  chosen: ReadonlySet<string> | null;
  /** The Compare group: the raid median, and the median of one class (a class name) or none. */
  compare: { raid: boolean; className: string | null };
  deadSpans?: ReadonlyMap<string, readonly DeadSpan[]>;
  /** One player as the Source: their line is the gold area (today's look). */
  single?: string;
}

export interface MedianInfo {
  id: string;
  label: string;
  description: string;
  count: number;
  on: boolean;
}

export interface RateModel {
  series: ChartSeries[];
  entries: PlayerEntry[];
  /** The head's figure: every player's figure rounded, then added. */
  raidFigure: number;
  total: number;
  isDefault: boolean;
  medians: MedianInfo[];
}

/** The players a median is taken over: damage dealers, healers, or (damage taken) everyone. */
function inMedianGroup(metric: RateMetric, player: RatePlayer): boolean {
  if (metric === 'healing') return player.role === 'healer';
  if (metric === 'damage-taken') return true;
  return player.role === 'dps';
}

function classPlural(className: string): string {
  return /s$/i.test(className) ? `${className.toLowerCase()}es` : `${className.toLowerCase()}s`;
}

/** The damage dealers of one class: a class median's members. */
function classMembersOf(players: readonly RatePlayer[], className: string): RatePlayer[] {
  return players.filter((player) => player.role === 'dps' && player.className === className);
}

/** A class median's Compare row, on or off: its name, its description and how many it covers. */
export function classMedianInfo(players: readonly RatePlayer[], className: string, on: boolean): MedianInfo {
  const count = classMembersOf(players, className).length;
  return {
    id: `median:${className.toLowerCase()}`,
    label: chartCopy.classMedian(className),
    description: chartCopy.classMedianDescription(count, classPlural(className)),
    count,
    on,
  };
}

function medianLine(
  id: string,
  label: string,
  colour: string,
  members: readonly RatePlayer[],
  window: TimeWindow,
): ChartSeries {
  return {
    id,
    label,
    colour,
    kind: 'median',
    dash: MEDIAN_DASH,
    values: medianSeries(members.map((player) => trailingAverage(player.series))),
    figure: median(members.map((player) => windowRate(player.series, window))),
  };
}

export function buildRateModel(input: RateInput): RateModel {
  const { metric, players, window, pinned, chosen, compare } = input;
  const figures = new Map(players.map((player) => [player.guid, windowRate(player.series, window)]));
  const inFight = new Set(players.map((player) => player.guid));
  const pinnedHere = new Set([...pinned].filter((guid) => inFight.has(guid)));
  const drawnIds = selectDrawn({
    candidates: players.map((player) => ({ id: player.guid, figure: figures.get(player.guid) ?? 0 })),
    pinned: pinnedHere,
    chosen,
  });
  const drawn = new Set(drawnIds);
  const byGuid = new Map(players.map((player) => [player.guid, player]));

  const group = players.filter((player) => inMedianGroup(metric, player));
  const classMembers = compare.className === null ? [] : classMembersOf(players, compare.className);
  const medians: MedianInfo[] = [
    {
      id: RAID_MEDIAN_ID,
      label: chartCopy.raidMedian,
      description: chartCopy.raidMedianDescription(metric, group.length),
      count: group.length,
      on: compare.raid,
    },
    ...(compare.className === null
      ? []
      : [
          {
            id: `median:${compare.className.toLowerCase()}`,
            label: chartCopy.classMedian(compare.className),
            description: chartCopy.classMedianDescription(
              classMembers.length,
              classPlural(compare.className),
            ),
            count: classMembers.length,
            on: true,
          },
        ]),
  ];

  const medianSeriesDrawn: ChartSeries[] = [
    ...(compare.raid && group.length > 0
      ? [medianLine(RAID_MEDIAN_ID, chartCopy.raidMedian, 'var(--color-text)', group, window)]
      : []),
    ...(compare.className !== null && classMembers.length > 0
      ? [medianLine(medians[1].id, medians[1].label, classColorVar(compare.className), classMembers, window)]
      : []),
  ];

  const playerSeries: ChartSeries[] = drawnIds.map((guid) => {
    const player = byGuid.get(guid)!;
    const spans = input.deadSpans?.get(guid) ?? [];
    const single = input.single === guid;
    return {
      id: guid,
      label: player.name,
      colour: single ? 'var(--color-gold)' : (player.colour ?? classColorVar(player.className)),
      kind: single ? 'single' : 'player',
      values: stopAtDeath(trailingAverage(player.series), spans),
      figure: figures.get(guid) ?? 0,
      pinned: pinnedHere.has(guid),
      you: guid === input.youGuid,
      diedAtMs: spans[0]?.startMs,
    };
  });

  const entries: PlayerEntry[] = [...players]
    .sort((a, b) => (figures.get(b.guid) ?? 0) - (figures.get(a.guid) ?? 0))
    .map((player) => ({
      id: player.guid,
      label: player.name,
      colour: player.colour ?? classColorVar(player.className),
      className: player.className,
      role: player.role,
      figure: figures.get(player.guid) ?? 0,
      drawn: drawn.has(player.guid),
      pinned: pinnedHere.has(player.guid),
      you: player.guid === input.youGuid,
    }));

  return {
    series: [...medianSeriesDrawn, ...playerSeries],
    entries,
    raidFigure: roundedSum([...figures.values()]),
    total: players.length,
    isDefault: chosen === null,
    medians,
  };
}
