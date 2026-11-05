// web/src/lib/sim/settings-copy.ts
// The Settings card's words (spec 2026-11-04-sim-results section 6, verbatim). Apart from
// result-copy.ts because the card is its own chunk and the boot chunk should not carry them.

export const settingsCopy = {
  label: "THIS RUN'S INPUT",
  title: 'Settings',
  keys: {
    fight: 'Fight length',
    target: 'Target',
    execute: 'Execute phase',
    race: 'Race',
    runs: 'Runs',
    apm: 'Actions per minute',
    engine: 'Engine',
  },
  fight: (length: string, variation: number, low: string, high: string) =>
    `${length}, ± ${variation}% (${low} to ${high})`,
  target: (level: number) => `Level ${level} boss`,
  armor: (armor: string) => `${armor} armor`,
  afterDebuffs: (after: string) => `${after} after debuffs`,
  execute: (percent: number) => `Last ${percent}% of the fight, standing in for the boss under 20% health`,
  olderEngine: 'Older engine',
  groups: { player: 'On you', target: 'On the target', consumables: 'Consumables' },
  noEffect: 'No effect on your damage',
  replacedBy: (winner: string) => `replaced by ${winner}`,
  mainHand: '(main hand)',
  offHand: '(off hand)',
  improved: ', improved',
  empty: 'None in this run.',
  showAll: (count: number) => `Show all ${count}`,
  copyJson: 'Copy JSON',
  downloadJson: 'Download JSON',
  jsonNote: 'The JSON is this whole result, request included.',
  fileName: (simId: string) => `forever-sixty-sim-${simId}.json`,
} as const;
