// web/src/lib/chart/copy.ts
// Every string the report charts show (controls and charts spec §B8), verbatim. Components read
// these and never spell their own.
import { clockLabel } from './time-axis';

export type RateMetric = 'damage-done' | 'healing' | 'damage-taken';

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
  rateMeta: (metric: RateMetric, figure: string, target?: string): string =>
    `Raid ${figure} ${UNIT[metric]}${target ? ` · on ${target}` : ''} · 5 s average`,
  playerMeta: (metric: RateMetric, name: string, figure: string, target?: string): string =>
    `${name} · ${figure} ${UNIT[metric]}${target ? ` · on ${target}` : ''} · 5 s average`,
  threatMeta: (startMs: number, endMs: number, whole: boolean): string =>
    `Cumulative · ${whole ? 'Whole fight' : `${clockLabel(startMs)} to ${clockLabel(endMs)}`}`,
  emptyMeta: (noun: string, phase: string | undefined, startMs: number, endMs: number): string =>
    `${phase ?? `${clockLabel(startMs)} to ${clockLabel(endMs)}`} · no ${noun}`,
  rateNoun: (metric: RateMetric): string => NOUN[metric],
  threatTitle: (enemy: string): string => `Estimated threat on ${enemy}`,
  threatTitlePhone: 'Estimated threat',
  threatCaveat:
    'An estimate: 1 threat per point of damage and 0.5 per point of healing, split across the enemies in combat. No stance, form, talent or taunt modifiers yet.',
  threatCaveatPhone: 'Damage and healing only.',
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
  playersField: (drawn: number, total: number, isDefault: boolean): string =>
    isDefault ? `Top 8 of ${total}` : `${drawn} of ${total}`,
  groups: { compare: 'Compare', tank: 'Tanks', healer: 'Healers', dps: 'Damage', other: 'Others' },
  raidMedian: 'Raid median',
  classMedian: (className: string): string => `${className} median`,
  raidMedianDescription: (metric: RateMetric, count: number): string =>
    metric === 'healing'
      ? `Median HPS of the ${count} healers`
      : metric === 'damage-taken'
        ? `Median DTPS of the ${count} players`
        : `Median DPS of the ${count} damage dealers`,
  classMedianDescription: (count: number, plural: string): string =>
    `Median DPS of the ${count} damage-dealing ${plural}`,
  topN: 'Top 8',
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
  tooltipRight: (phase: string | undefined, threat: boolean): string =>
    `${phase === undefined ? '' : `${phase} phase · `}${threat ? 'estimated threat' : '5 s average'}`,
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
