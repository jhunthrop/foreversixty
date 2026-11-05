// web/src/lib/sim/input-cards-copy.ts
// The words of the saved run's lower cards (spec 2026-11-04-sim-results section 6, verbatim):
// Gear, Talents, Resource, Buffs and Rotation. Apart from result-copy.ts because these cards are
// their own chunks and the boot chunk should not carry them.

export const gearCardCopy = {
  label: "THIS RUN'S INPUT",
  title: 'Gear',
  meta: (worn: number, total: number, enchants: number) =>
    `${worn} of ${total} slots · ${enchants === 0 ? 'no enchants' : `${enchants} ${enchants === 1 ? 'enchant' : 'enchants'}`}`,
  slotLine: (slot: string, itemLevel: number) => `${slot} · ${itemLevel}`,
  statsHead: (level: number) => `Character stats in the fight, against a level ${level} boss`,
  stats: {
    attackPower: 'Attack power',
    hit: 'Hit',
    crit: 'Crit on the boss',
    skill: 'Weapon skill',
    haste: 'Haste',
    spellPower: 'Spell power',
    spellHit: 'Spell hit',
    spellCrit: 'Spell crit',
  },
  missedInRun: (parts: string) => `Missed in this run: ${parts}`,
  missedSpecials: (share: string) => `${share} of specials`,
  missedWhite: (share: string) => `${share} of white swings`,
  noSet: 'No set bonus is active: no two pieces share a set.',
  setLine: (name: string, worn: number, total: number) => `Set: ${name} ${worn} of ${total}`,
  bonus: (pieces: number, text: string) => `(${pieces}) ${text}`,
  bisLink: 'Upgrades for this character on Best in slot',
  noGear: 'This run had no gear on file.',
} as const;

export const talentsCardCopy = {
  label: "THIS RUN'S INPUT",
  title: 'Talents',
  footDesktop: 'Read only. Hover or focus a talent for its tooltip.',
  footPhone: 'Read only. Tap a talent for its tooltip.',
  planner: 'Open in planner',
  doesNotFit:
    'These talents were saved on an older talent layout. The point split above is as saved; the trees have changed since, so they are not drawn.',
  failed: "Talents didn't load.",
} as const;

export const resourceCardCopy = {
  label: 'AVERAGE FIGHT',
  gained: 'Gained',
  spent: 'Spent',
  lostAtCap: 'Lost at the cap',
  endedWith: 'Ended with',
  gainedFrom: 'Gained from',
  spentOn: 'Spent on',
  lost: (amount: string) => `−${amount} lost`,
  capValue: (amount: string, percent: string) => `${amount} · ${percent}%`,
  needsNewerRun: (resource: string) => `Where the ${resource} came from and went needs a newer run.`,
} as const;

export const aurasCardCopy = {
  label: 'AVERAGE FIGHT',
  title: 'Buffs and debuffs',
  meta: 'Uptime, then times applied',
  onYou: 'On you',
  onTarget: 'On the target',
  comeAndGo: (count: number) => (count === 1 ? '1 that comes and goes' : `${count} that come and go`),
  upAllFight: 'Up all fight',
  noEffect: 'No effect on your damage',
  stacks: (count: number) => `× ${count}`,
  applied: (count: number) => `${count}×`,
  targetNeedsNewerRun: 'Uptime on the target needs a newer run.',
  mainHand: '(main hand)',
  offHand: '(off hand)',
  improved: ', improved',
} as const;

export const rotationCardCopy = {
  label: (spec: string, level: number) => `${spec.toUpperCase()}, AT ${level}`,
  title: 'Rotation',
  meta: (lines: number) => `${lines} ${lines === 1 ? 'line' : 'lines'} · casts are per fight`,
  today: " · today's rotation",
  beforePull: 'Before the pull',
  thenPriority: 'Then, from the top at every decision',
  cast: 'Cast',
  notCast: 'Not cast in this run',
  casts: (count: number) => (count === 1 ? '1 cast' : `${count} casts`),
  otherCooldowns: 'Other cooldowns as they come up',
  thisRun: (names: string) => `This run: ${names}.`,
  rank: (rank: number) => `Rank ${rank}`,
  prepullAt: (seconds: number) => `-${seconds} s`,
  foot: (spec: string) =>
    `The ${spec} rotation the engine follows. A line with no cast in this run is dimmed.`,
  guide: (spec: string) => `Read the ${spec} guide`,
  noRotation: (spec: string) => `The ${spec} rotation is not on file for this build.`,
} as const;
