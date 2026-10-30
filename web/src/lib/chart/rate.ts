// web/src/lib/chart/rate.ts
// The damage, healing and damage-taken chart's model (controls and charts spec D5, §B5.3): every
// player's average over the fight's bucket (buckets.ts: every second on a short pull, 5 s on a long fight), the top 8 of N plus the always-drawn players, the raid median and a class
// median, and the head's raid figure, which is the sum of the figures the legend shows.
import { classColorVar } from '../report/format';
import { BUCKET_MS, type TimeWindow } from '../report/window';
import { bucketMeans, chartBucket, type ChartBucket } from './buckets';
import type { ScopeNoun } from '../report/group';
import type { RateMetric } from './copy';
import { chartCopy } from './copy';
import {
  type DeadSpan,
  type DrawPreset,
  median,
  medianSeries,
  roundedSum,
  selectDrawn,
  stopAtDeath,
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
  /** The whole fight: it picks the bucket the lines are averaged over. */
  durationMs: number;
  /** The Source player and you: always drawn. */
  pinned: ReadonlySet<string>;
  youGuid?: string;
  /** The Players control's choice; null keeps the preset. */
  chosen: ReadonlySet<string> | null;
  /** The default draw; `split` (top 4 and bottom 4) unless the Players control picked another. */
  preset?: DrawPreset;
  /** The Compare group: the raid median, and the median of one class (a class name) or none. */
  compare: { raid: boolean; className: string | null };
  deadSpans?: ReadonlyMap<string, readonly DeadSpan[]>;
  /** One player as the Source: their line is the gold area (today's look). */
  single?: string;
  /** What the players are (lib/report/group.ts): it names the median line. A raid when absent. */
  noun?: ScopeNoun;
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
  /** The width every line's points are averaged over. */
  bucket: ChartBucket;
  /** One more line (a charted ability) on the same points: averaged over the bucket, with its step. */
  pointsOf: (values: readonly number[]) => Pick<ChartSeries, 'values' | 'stepMs'>;
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

/** A series as the chart draws it: averaged over the bucket, across the fight's seconds. */
type Averaged = (values: readonly number[]) => number[];

function medianLine(
  id: string,
  label: string,
  colour: string,
  members: readonly RatePlayer[],
  window: TimeWindow,
  averaged: Averaged,
): ChartSeries {
  return {
    id,
    label,
    colour,
    kind: 'median',
    dash: MEDIAN_DASH,
    values: medianSeries(members.map((player) => averaged(player.series))),
    figure: median(members.map((player) => windowRate(player.series, window))),
  };
}

export function buildRateModel(input: RateInput): RateModel {
  const { metric, players, window, pinned, chosen, compare } = input;
  const bucket = chartBucket(input.durationMs);
  const stepMs = bucket.seconds * BUCKET_MS;
  const spanS = Math.ceil(Math.max(0, input.durationMs) / BUCKET_MS);
  const averaged: Averaged = (values) => bucketMeans(values, bucket.seconds, spanS);
  const medianName = chartCopy.median(input.noun ?? 'raid');
  const figures = new Map(players.map((player) => [player.guid, windowRate(player.series, window)]));
  const inFight = new Set(players.map((player) => player.guid));
  const pinnedHere = new Set([...pinned].filter((guid) => inFight.has(guid)));
  const drawnIds = selectDrawn({
    candidates: players.map((player) => ({ id: player.guid, figure: figures.get(player.guid) ?? 0 })),
    pinned: pinnedHere,
    chosen,
    preset: input.preset ?? 'split',
    // The bottom is the metric's own role: a healer's damage is not a low performer on Damage done.
    bottomPool: new Set(
      players.filter((player) => inMedianGroup(metric, player)).map((player) => player.guid),
    ),
  });
  const drawn = new Set(drawnIds);
  const byGuid = new Map(players.map((player) => [player.guid, player]));

  const group = players.filter((player) => inMedianGroup(metric, player));
  const classMembers = compare.className === null ? [] : classMembersOf(players, compare.className);
  const medians: MedianInfo[] = [
    {
      id: RAID_MEDIAN_ID,
      label: medianName,
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
      ? [medianLine(RAID_MEDIAN_ID, medianName, 'var(--color-text)', group, window, averaged)]
      : []),
    ...(compare.className !== null && classMembers.length > 0
      ? [
          medianLine(
            medians[1].id,
            medians[1].label,
            classColorVar(compare.className),
            classMembers,
            window,
            averaged,
          ),
        ]
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
      values: stopAtDeath(averaged(player.series), spans, stepMs),
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
    series: [...medianSeriesDrawn, ...playerSeries].map((line) => ({ ...line, stepMs })),
    entries,
    raidFigure: roundedSum([...figures.values()]),
    total: players.length,
    isDefault: chosen === null,
    medians,
    bucket,
    pointsOf: (values) => ({ values: averaged(values), stepMs }),
  };
}
