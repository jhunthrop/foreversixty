// web/src/lib/sim/input-cards-copy.ts
// The words of a sim result's lower sections (spec 2026-11-04-sim-results section 6, verbatim;
// sim spec 2026-11-06 section 6 where it changed them): Gear, Talents, Resource, Buffs and
// Rotation. Apart from result-copy.ts because these sections are their own chunks and the boot
// chunk should not carry them.

export const gearCardCopy = {
  title: 'Gear',
  meta: (worn: number, total: number, enchants: number) =>
    `${worn} of ${total} slots · ${enchants === 0 ? 'no enchants' : `${enchants} ${enchants === 1 ? 'enchant' : 'enchants'}`}`,
  itemLevel: (level: string) => `item level ${level}`,
  statsHead: (level: number) => `In the fight, against a level ${level} boss.`,
  stats: {
    attackPower: 'Attack power',
    hit: 'Hit',
    crit: 'Crit on the boss',
    skill: 'Weapon skill',
    haste: 'Haste from gear',
    spellPower: 'Spell power',
    spellHit: 'Spell hit',
    spellCrit: 'Spell crit',
  },
  missedInRun: (parts: string) => `Missed in this run: ${parts}.`,
  missedSpecials: (share: string) => `${share} of specials`,
  missedWhite: (share: string) => `${share} of white swings`,
  noSet: 'No set bonus is active: no two pieces share a set.',
  setLine: (name: string, worn: number, total: number) => `Set: ${name} ${worn} of ${total}`,
  bonus: (pieces: number, text: string) => `(${pieces}) ${text}`,
  bisLink: 'Upgrades for this character on Best in slot',
  /** The landing's plated doll: its Section action. */
  bisAction: 'Upgrades on Best in slot',
  noGear: 'This run had no gear on file.',
} as const;

export const talentsCardCopy = {
  title: 'Talents',
  footDesktop: 'Read only. Hover or focus a talent for its tooltip.',
  footPhone: 'Read only. Tap a talent for its tooltip.',
  planner: 'Open in planner',
  doesNotFit:
    'These talents were saved on an older talent layout. The point split above is as saved; the trees have changed since, so they are not drawn.',
  failed: "Talents didn't load.",
} as const;

export const resourceCardCopy = {
  meta: 'Average fight',
  gained: 'Gained',
  spent: 'Spent',
  lostAtCap: 'Lost at the cap',
  endedWith: 'Ended with',
  gainedFrom: 'Gained from',
  spentOn: 'Spent on',
  /** The engine's rage refund: 80% of a special's cost back when it misses, is dodged or parried. */
  refunded: 'Refunded on a miss, dodge or parry',
  lost: (amount: string) => `−${amount} lost`,
  capValue: (amount: string, percent: string) => `${amount} · ${percent}%`,
  needsNewerRun: (resource: string) => `Where the ${resource} came from and went needs a newer run.`,
} as const;

export const aurasCardCopy = {
  title: 'Buffs and debuffs',
  meta: 'Uptime, then times applied',
  onYou: 'On you',
  onTarget: 'On the target',
  comeAndGo: (count: number) => (count === 1 ? '1 that comes and goes' : `${count} that come and go`),
  upAllFight: 'Up all fight',
  dimmed: 'Dimmed: no effect on your damage.',
  stacks: (count: number) => `× ${count}`,
  applied: (count: number) => `${count}×`,
  targetNeedsNewerRun: 'Uptime on the target needs a newer run.',
  mainHand: '(main hand)',
  offHand: '(off hand)',
  improved: ', improved',
} as const;

export const rotationCardCopy = {
  title: 'Rotation',
  /** The section's meta: before a run, then with this run's casts. */
  meta: (spec: string, level: number) => `${spec}, at ${level}`,
  metaRan: (spec: string, level: number) => `${spec}, at ${level} · casts per fight`,
  today: " · today's rotation",
  beforePull: 'Before the pull',
  thenPriority: 'Then, from the top at every decision',
  cast: 'Cast',
  notCast: 'Not cast in this run',
  casts: (count: number) => (count === 1 ? '1 cast' : `${count} casts`),
  otherCooldowns: 'Other cooldowns as they come up',
  thisRun: (names: string) => `This run: ${names}.`,
  beforeRun: 'Trinkets, racials and potions, as each comes off cooldown.',
  rank: (rank: number) => `Rank ${rank}`,
  prepullAt: (seconds: number) => `-${seconds} s`,
  foot: (spec: string) => `The ${spec} rotation the engine follows.`,
  footRan: (spec: string) =>
    `The ${spec} rotation the engine follows. A line with no cast in this run is dimmed.`,
  guide: (spec: string) => `${spec} guide`,
  noRotation: (spec: string) => `The ${spec} rotation is not on file for this build.`,
} as const;
