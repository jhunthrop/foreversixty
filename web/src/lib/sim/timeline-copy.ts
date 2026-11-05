// web/src/lib/sim/timeline-copy.ts
// The saved sim timeline's player-facing words, verbatim from the spec (2026-11-04-sim-results
// section 6, Timeline, and 4.3's fallbacks). Figures are arguments; no component spells a sentence.

/**
 * The first day a saved run can carry `sample_auras` (spec 4.3's "{launch date of the field}"):
 * the day the engine pin that records them lands. One constant, so the integrator moves one line
 * if the pin slips.
 */
export const SAMPLE_AURAS_SINCE = '5 Nov 2026';

const OUTCOME_LABEL = {
  hit: 'Hit',
  crit: 'Crit',
  glance: 'Glance',
  miss: 'Miss',
  dodge: 'Dodge',
  parry: 'Parry',
  block: 'Block',
  tick: 'Hit',
  tick_crit: 'Crit',
} as const;

export const timelineCopy = {
  label: 'ONE FIGHT',
  title: 'Timeline',
  meta: (length: string): string => `The median run · ${length} long`,
  noSample: 'This run has no recorded fight to draw.',

  wholeFight: 'Whole fight',
  closeUp: 'Close up',
  zoomLabel: 'Zoom',
  swingsAndTicks: 'Swings and ticks',

  legend: {
    damage: 'Damage',
    crit: 'Crit',
    missed: 'Missed',
    cooldown: 'Cooldown',
    onYou: 'On you',
    onTarget: 'On the target',
  },

  groups: {
    casts: 'Casts',
    swings: 'Swings and ticks',
    player: 'On you',
    target: 'On the target',
  },
  hands: { 1: 'Main hand', 2: 'Off hand', 3: 'Extra attacks' } as Readonly<Record<number, string>>,
  ticksLane: (name: string): string => `${name} ticks`,
  overpowerWindow: 'Overpower window',
  offHandHit: (name: string): string => `${name} (off hand)`,

  axisTime: 'Time',
  axisPull: 'Pull',
  executeWash: 'EXECUTE PHASE',
  resourcePeak: (peak: number, cap: number): string => `peak ${peak} of ${cap}`,

  aurasFallback: `Buff and debuff bars need a run from ${SAMPLE_AURAS_SINCE} or later. The casts and rage here are from this one.`,
  openInSim: 'Open in the sim',

  readoutResource: (label: string, value: string): string => `${label} ${value}`,
  readoutOnYou: (names: readonly string[]): string =>
    names.length === 0 ? 'On you: nothing extra' : `On you: ${names.join(', ')}`,
  secondsAgo: (seconds: string): string => `${seconds} s ago`,
  closeReadout: 'Close the readout',

  rank: (rank: number): string => `Rank ${rank}`,
  tipWhen: (time: string, resource: string, value: string): string => `${time} · ${resource} after ${value}`,
  tipNext: (name: string, seconds: string): string => `Next ${name} ${seconds} s later`,
  tipLast: (name: string): string => `Last ${name} of the fight`,
  tipOutcome: (outcome: keyof typeof OUTCOME_LABEL, amount: string | null): string =>
    amount === null ? OUTCOME_LABEL[outcome] : `${OUTCOME_LABEL[outcome]} · ${amount} damage`,

  castsAll: (shown: number, total: number): string => `Casts: ${shown} of ${total}`,
  castsInView: (shown: number, total: number): string => `Casts in view: ${shown} of ${total}`,
  castsTail: (periodic: string | null): string =>
    periodic === null
      ? 'Swings are on the lanes above.'
      : `Swings and ${periodic} ticks are on the lanes above.`,
  showAll: (count: number): string => `Show all ${count}`,
  showFewer: 'Show fewer',
  queued: (name: string): string => `${name} queued`,

  plotLabel: (length: string, casts: number): string =>
    `Fight timeline, ${length} long, ${casts} casts. Arrow keys move between casts.`,
  plotRole: 'timeline',
  brushLabel: 'Close up window. Arrow keys pan it, Shift with an arrow key resizes it.',
} as const;
