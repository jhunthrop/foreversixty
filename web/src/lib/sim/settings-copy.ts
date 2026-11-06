// web/src/lib/sim/settings-copy.ts
// The This run section's words (sim spec 2026-11-06 section 6, on spec 2026-11-04-sim-results
// section 6). Apart from result-copy.ts because the section is its own chunk and the boot chunk
// should not carry them.

export const settingsCopy = {
  title: 'This run',
  meta: 'Inputs and engine',
  keys: {
    fight: 'Fight length',
    target: 'Target',
    execute: 'Execute phase',
    race: 'Race',
    runs: 'Runs',
    apm: 'Actions per minute',
    engine: 'Engine',
  },
  fight: (length: string, variation: number) => `${length}, ± ${variation}%`,
  range: (low: string, high: string) => `${low} to ${high}`,
  target: (level: number) => `Level ${level} boss`,
  armor: (armor: string) => `${armor} armor`,
  afterDebuffs: (after: string) => `${after} after debuffs`,
  execute: (percent: number) => `Last ${percent}% of the fight`,
  executeLength: (length: string, fight: string) => `${length} of a ${fight} fight`,
  groups: { player: 'On you', target: 'On the target', consumables: 'Consumables' },
  group: (label: string, count: number) => `${label} · ${count}`,
  dimmed: 'Dimmed: no effect on your damage.',
  copyForGame: 'Copy for the game',
  replacedBy: (winner: string) => `replaced by ${winner}`,
  mainHand: '(main hand)',
  offHand: '(off hand)',
  improved: ', improved',
  copyJson: 'Copy JSON',
  downloadJson: 'Download JSON',
  jsonNote: 'The JSON is this whole result, request included.',
  fileName: (simId: string) => `forever-sixty-sim-${simId}.json`,
} as const;
