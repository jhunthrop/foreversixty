// web/src/lib/chart/copy.ts
// Every string the report charts show (controls and charts spec §B8), verbatim. Components read
// these and never spell their own.
import type { ScopeNoun } from '../report/group';
import { clockLabel } from './time-axis';

export type RateMetric = 'damage-done' | 'healing' | 'damage-taken';

/** The head's figure, named for the set it sums. */
const SET_NAME: Record<ScopeNoun, string> = { raid: 'Raid', group: 'Group', everyone: 'Total' };
const MEDIAN_NAME: Record<ScopeNoun, string> = {
  raid: 'Raid median',
  group: 'Group median',
  everyone: 'Median player',
};
const UNIT: Record<RateMetric, string> = { 'damage-done': 'DPS', healing: 'HPS', 'damage-taken': 'DTPS' };
const NOUN: Record<RateMetric, string> = {
  'damage-done': 'damage',
  healing: 'healing',
  'damage-taken': 'damage taken',
};

export const chartCopy = {
  title: { 'damage-done': 'Damage done', healing: 'Healing', 'damage-taken': 'Damage taken' } as Record<
    RateMetric,
    string
  >,
  unit: (metric: RateMetric): string => UNIT[metric],
  /** The span every figure on a chart is over: the head, the legend and the Players list agree. */
  span: (startMs: number, endMs: number, whole: boolean): string =>
    whole ? 'the whole fight' : `${clockLabel(startMs)} to ${clockLabel(endMs)}`,
  /**
   * Fix round 2: the head says what its figure is, the sum of every player's figure over the span,
   * named for the set it sums: a raid, a group, or (everyone in range) a total.
   */
  rateMeta: (
    metric: RateMetric,
    figure: string,
    span: string,
    target?: string,
    noun: ScopeNoun = 'raid',
  ): string => `${SET_NAME[noun]} ${figure} ${UNIT[metric]} over ${span}${target ? ` · on ${target}` : ''}`,
  playerMeta: (metric: RateMetric, name: string, figure: string, span: string, target?: string): string =>
    `${name} · ${figure} ${UNIT[metric]} over ${span}${target ? ` · on ${target}` : ''}`,
  threatMeta: (startMs: number, endMs: number, whole: boolean): string =>
    `Cumulative · ${whole ? 'Whole fight' : `${clockLabel(startMs)} to ${clockLabel(endMs)}`}`,
  emptyMeta: (noun: string, phase: string | undefined, startMs: number, endMs: number): string =>
    `${phase ?? `${clockLabel(startMs)} to ${clockLabel(endMs)}`} · no ${noun}`,
  rateNoun: (metric: RateMetric): string => NOUN[metric],
  /** Fix round 2: the title names what the Enemy field picked, at every width. */
  threatTitle: (enemy: string | null): string =>
    enemy === null ? 'Estimated threat, every enemy' : `Estimated threat on ${enemy}`,
  threatSubject: (enemy: string | null): string => (enemy === null ? 'every enemy' : enemy),
  threatCaveat:
    'An estimate: 1 threat per point of damage and 0.5 per point of healing, split across the enemies in combat. No stance, form, talent or taunt modifiers yet.',
  threatCaveatPhone: 'An estimate: 1 per damage, 0.5 per heal.',
  threatCaveatLink: 'What it leaves out',
  threatTableTitle: 'Estimated threat',
  threatTableMeta: (enemy: string, atMs: number, players: number): string =>
    `On ${enemy} · standing at ${clockLabel(atMs)} · ${players} players`,
  threatColumn: 'Estimated threat',
  standingColumn: 'Standing',
  shareColumn: 'Share of estimated threat',
  tooltipThreatColumn: 'Estimated threat',
  tooltipShareColumn: 'Share',
  players: 'Players',
  window: 'Window',
  enemy: 'Enemy',
  /** The Players field: what is drawn, never a limit above the roster (`All 3`, not `Top 8 of 3`). */
  /** `seen`: the set is everyone in range, so the count says so (`of 90 players seen`). */
  playersField: (
    drawn: number,
    total: number,
    choice: 'split' | 'top' | 'bottom' | 'custom',
    seen = false,
  ): string => {
    const of = `${total}${seen ? ' players seen' : ''}`;
    if (choice === 'custom') return `${drawn} of ${of}`;
    if (total <= 8) return `All ${of}`;
    if (choice === 'top') return `Top 8 of ${of}`;
    if (choice === 'bottom') return `Bottom 8 of ${of}`;
    return `Top 4, bottom 4 of ${of}`;
  },
  presets: { split: 'Top 4 + bottom 4', top: 'Top 8', bottom: 'Bottom 8', custom: 'Custom' },
  groups: { compare: 'Compare', tank: 'Tanks', healer: 'Healers', dps: 'Damage', other: 'Others' },
  /** The median line, named for the set it is the median of. */
  median: (noun: ScopeNoun): string => MEDIAN_NAME[noun],
  classMedian: (className: string): string => `${className} median`,
  raidMedianDescription: (metric: RateMetric, count: number): string =>
    metric === 'healing'
      ? `Median HPS of the ${count} healers`
      : metric === 'damage-taken'
        ? `Median DTPS of the ${count} players`
        : `Median DPS of the ${count} damage dealers`,
  classMedianDescription: (count: number, plural: string): string =>
    `Median DPS of the ${count} damage-dealing ${plural}`,
  drawnCount: (n: number): string => `${n} drawn`,
  filterPlayers: 'Filter players',
  alwaysDrawn: 'Always drawn',
  youList: 'You',
  youChip: 'YOU',
  more: (n: number): string => `+${n} more`,
  morePhone: (n: number): string => `+${n}`,
  only: 'Only',
  hide: (name: string): string => `Hide ${name}`,
  show: (name: string): string => `Show ${name}`,
  showOnly: (name: string): string => `Show only ${name}`,
  sheetMeta: (drawn: number, total: number): string => `${drawn} of ${total} drawn`,
  sheetDrawn: 'Drawn',
  sheetNotDrawn: 'Not drawn',
  done: 'Done',
  wholeFight: 'Whole fight',
  logScale: 'Log scale',
  dragLabel: (startMs: number, endMs: number): string =>
    `${clockLabel(startMs)} to ${clockLabel(endMs)} · ${Math.round((endMs - startMs) / 1000)} s`,
  /** The tooltip head's right side: the phase, then what the rows are (the threat chart names its enemy). */
  tooltipRight: (phase: string | undefined, rows: string): string =>
    `${phase === undefined ? '' : `${phase} phase · `}${rows}`,
  /** What a rate chart's points are: the average over their bucket (lib/chart/buckets.ts). */
  tooltipRate: (metric: RateMetric, bucketS: number): string =>
    `${UNIT[metric]}, ${bucketS === 1 ? 'each second' : `${bucketS} s average`}`,
  tooltipThreat: (subject: string): string => `estimated threat, ${subject}`,
  tooltipMore: (n: number): string => `+${n} more`,
  diedAt: (ms: number): string => `Died ${clockLabel(ms)}`,
  deathCluster: (n: number): string => `${n} deaths`,
  clusterMore: (n: number): string => `+${n} more · open on the Deaths tab`,
  markerCluster: (n: number): string => `${n} events`,
  deathAria: (line: string): string => `${line}. Open on the Deaths tab.`,
  plotAria: (title: string, startMs: number, endMs: number): string =>
    `${title} chart, ${clockLabel(startMs)} to ${clockLabel(endMs)}. Arrow keys move through time.`,
  live: (ms: number, rows: readonly { name: string; value: string }[]): string =>
    `${clockLabel(ms)}. ${rows.map((row) => `${row.name} ${row.value}`).join(', ')}.`,
  loading: 'Loading the chart.',
  error: 'The chart didn’t load.',
  emptyRate: (metric: RateMetric): string => `No ${NOUN[metric]} in this window.`,
  emptyThreat: (enemy: string): string => `No threat on ${enemy} in this window.`,
  touchHint: 'Touch and hold the chart to read a moment.',
} as const;
