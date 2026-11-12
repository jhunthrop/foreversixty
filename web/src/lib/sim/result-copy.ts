// web/src/lib/sim/result-copy.ts
// Every player-facing word of the saved run result page that lane A draws (spec
// design/specs/2026-11-04-sim-results.md section 6, verbatim; sim spec 2026-11-06 section 6 where
// it changed them): the band, Worth knowing, Damage by ability, the DPS spread and the page's errors. Its own
// module rather than more keys on copy.ts, which is already past the 800-line ceiling. Settings
// has its own (settings-copy.ts), because that card is its own chunk.
import { uiCopy } from '../ui/copy';

export const bandCopy = {
  eyebrow: 'SAVED SIM',
  /** A finished run on /sim: where it ran, which the saved band leaves to its Settings card. */
  liveEyebrow: (lane: 'browser' | 'server') =>
    lane === 'server' ? 'RUN ON OUR SERVERS' : 'RUN IN YOUR BROWSER',
  level: (level: number) => `Level ${level}`,
  talents: (split: string) => `${split} talents`,
  dps: 'DPS',
  error: 'error',
  openInSim: 'Open in the sim',
  copyLink: 'Copy link',
  facts: {
    fight: 'Fight',
    targets: 'Targets',
    buffs: 'Buffs',
    consumables: 'Consumables',
    range: 'Range',
    runs: 'Runs',
    saved: 'Saved',
  },
  buffsRaid: 'Raid',
  buffsCustom: 'Custom',
  none: 'None',
  range: (low: string, high: string) => `${low} to ${high}`,
  fight: (length: string, variation: number) => `${length} ± ${variation}%`,
} as const;

/** The word for the player's power, the resource section's title. */
export const powerCopy = { rage: 'Rage', energy: 'Energy', mana: 'Mana' } as const;

export const warningsCopy = {
  stale: 'This run used an older engine. A run today may give a different number.',
  capstone: (spell: string) =>
    `${spell} was never pressed in this run. Check the talents in Open in the sim.`,
  missingTalent: (spell: string, spec: string) =>
    `${spell} is in the ${spec} rotation, but these talents do not take it.`,
  atCap: (amount: string, resource: string, percent: string) =>
    `${amount} ${resource} (${percent}%) was lost at the cap.`,
  atCapRage: 'Queue Heroic Strike sooner so swings do not land at full rage.',
  atCapEnergy: 'Spend before the bar fills.',
  wideError: (margin: string) =>
    `This run's error is ± ${margin} DPS. A difference smaller than that between two runs is noise.`,
  openInSim: 'Open in the sim',
  openInPlanner: 'Open in planner',
} as const;

export const damageCopy = {
  title: 'Damage by ability',
  meta: (dps: string, length: string) => `${dps} DPS over ${length} on average`,
  headers: {
    ability: 'Ability',
    shareOfDamage: 'Share of damage',
    dps: 'DPS',
    share: 'Share',
    casts: 'Casts',
    avgHit: 'Avg hit',
    crit: 'Crit',
    biggestHit: 'Biggest hit',
    uptime: 'Uptime',
  },
  critTip:
    'Crits out of the hits that landed, glancing blows included. Misses, dodges and parries are left out, so a row can read above or below your crit chance on the boss.',
  rank: (rank: number) => `Rank ${rank}`,
  outcome: { MISS: 'Miss', DODGE: 'Dodge', PARRY: 'Parry', GLANCING: 'Glance', BLOCK: 'Block' },
  ticks: (count: string, crit: string) => [`${count} ticks`, `tick crit ${crit}%`] as const,
  /** A DoT's uptime on its sub line (the Uptime column retired, sim spec 2026-11-06 §4.7). */
  upFor: (percent: string) => `up ${percent}% of the fight`,
  /** A spell that hits several enemies: its landed attempts a cast, and the fight's targets. */
  hitsACast: (hits: string, targets: number) =>
    `${hits} hits a cast, ${targets} ${targets === 1 ? 'target' : 'targets'}`,
  becameHeroicStrike: (count: string) => `${count} swings became Heroic Strike`,
  fromSources: 'From',
  phoneCasts: (count: string, swings: boolean) => `${count} ${swings ? 'swings' : 'casts'}`,
  phoneAvg: (avg: string) => `${avg} avg hit`,
  phoneCrit: (crit: string) => `crit ${crit}%`,
  pet: (name: string) => `Pet · ${name}`,
  total: (sources: number) => `Total, ${sources} sources`,
  totalDamage: (damage: string) => `${damage} damage a fight`,
  dash: '–',
  /** The swing rows' own names (spec section 6). */
  hands: { main: 'Main hand', off: 'Off hand', extra: 'Extra attacks' },
  offHandSuffix: '(off hand)',
} as const;

export const spreadCopy = {
  meta: (runs: string) => `${runs} runs`,
  title: 'DPS spread',
  axis: 'DPS per run',
  bar: (low: string, high: string, runs: string) => `${low} to ${high} DPS: ${runs} runs`,
  image: (runs: string, low: string, high: string, mean: string) =>
    `DPS of each of ${runs} runs, from ${low} to ${high}, most near ${mean}`,
  rangeCaption:
    'This run kept the range, not the shape. The shaded part is one spread either side of the mean.',
  rangeMean: (mean: string) => `mean ${mean}`,
  tiles: {
    mean: 'Mean',
    error: 'Error',
    stddev: 'Std dev',
    lowest: 'Lowest',
    highest: 'Highest',
    runs: 'Runs',
  },
  /** The published set the run is read against (lib/sim/bis-benchmark.ts): "Level 60 raid BiS". */
  benchmark: (band: number, preset: 'raid' | 'bare'): string => `Level ${band} ${preset} BiS`,
  /** "9.9% under": the run against the set, beside the set's DPS. */
  benchmarkGap: (percent: string, under: boolean): string => `${percent}% ${under ? 'under' : 'over'}`,
} as const;

export const pageErrorCopy = {
  label: 'SAVED SIM',
  title: 'This sim did not load',
  body: 'The server did not answer. Your link is fine; try again in a moment.',
  retry: uiCopy.retry,
  partFailed: "This part of the page didn't load.",
} as const;
