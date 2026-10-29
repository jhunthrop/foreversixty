// Every user-visible string the guides lane's own components (GuideBuildTree,
// BuildActionButtons, RacePillRow, StatPriorityPills) render, reference voice throughout
// (design/DESIGN-SYSTEM.md: state the thing and stop).
export const guidesCopy = {
  loadThisBuild: 'Load this build',
  simThisBuild: 'Sim this build',
  treeLoadFailed: 'This build did not load',
  recommendedRaceLabel: 'Recommended',
  // --- rebuild (2026-10-04) --------------------------------------------------------------
  /** `ClassHeader`'s own eyebrow row, all three guide surfaces (spec §6's copy table). */
  eyebrow: 'Guides',
  /** `/guides` index's own signed-in callout (spec §4.A/§6): `{specName}` and `{className}`
   *  come from the visitor's own character, e.g. "Your guide: Fury Warrior →". */
  signedInCallout: (specName: string, className: string): string => `Your guide: ${specName} ${className} →`,
  indexHonestyLine: "The beta stops at level 30. Level 60 picks come from this site's simulator.",
  roleLabel: (role: 'dps' | 'healer' | 'tank'): string =>
    role === 'dps' ? 'DPS' : role === 'healer' ? 'Healer' : 'Tank',
  readTheGuide: 'Read the guide',
  readTheGuideAriaLabel: (specName: string, className: string): string =>
    `Read the ${specName} ${className} guide`,
  yourSpecPill: 'Your spec',
  setDpsLine: (dps: number): string => `${dps.toFixed(1)} DPS`,
  setDpsLineSuffix: 'at 60, this set',
  buildRailLabel: 'Build',
  rotationRailLabel: (bandLabel: string): string => `Rotation · ${bandLabel}`,
  fullPriorityLink: 'Full priority ↓',
  statPriorityRailLabel: 'Stat priority',
  levelingPointsLabel: (points: number): string => `${points} point${points === 1 ? '' : 's'}`,
  loadInPlanner: 'Load in planner',
  loadBuildRailAriaLabel: 'Load this build in the planner (summary)',
  simBuildRailAriaLabel: 'Sim this build (summary)',
  // --- raid build (2026-10-07) -------------------------------------------------------------
  levelingBuild: 'Level 60 build',
  raidBuild: 'Raid build',
  /** The one heading a spec shows when its raid build is its leveling build. */
  levelingAndRaidBuild: 'Level 60 and raid build',
  loadLevelingBuildAriaLabel: 'Load the level 60 build in the planner',
  simLevelingBuildAriaLabel: 'Sim the level 60 build',
  loadRaidBuildAriaLabel: 'Load the raid build in the planner',
  simRaidBuildAriaLabel: 'Sim the raid build',
  // --- guides like Wowhead's (W.18, 2026-10-22) ---------------------------------------------
  /** The one status note for the whole page; the sections are then written as a guide. */
  statusNote:
    "The beta stops at level 30. Everything above it on this page, the talents, rotation, stat weights and best in slot, comes from this site's simulator running on the client's own data, not from play. Class mechanics carry over from 1.12 where the demo showed no change.",
  statusNoteLabel: 'About this guide',
  /** The muted tag beside `Updated` in the header: the guide is simulated, the beta is not over. */
  betaTag: 'Beta',
  headerMeta: (role: 'dps' | 'healer' | 'tank', updated: string): string =>
    `${role === 'dps' ? 'DPS' : role === 'healer' ? 'Healer' : 'Tank'} · Updated ${updated}`,
  strengths: 'Strengths',
  weaknesses: 'Weaknesses',
  inARaid: 'In a raid',
  keyTalents: 'Key talents',
  opener: 'Opener',
  priority: 'Priority',
  cooldowns: 'Cooldowns',
  bestInSlot: 'Best in slot',
  levelRange: 'Level range',
  bisCaption: (race: string, faction: string): string =>
    `Picks for a ${race} (${faction}), ranked by this site's simulator with every item equipped.`,
  fullListLink: 'Full list and alternatives',
  preRaidNote:
    'These are the best picks before the first raid tier. Raid loot is not itemized in the client yet, and this list updates when it is.',
  notSimChecked: 'Not sim-checked',
  slotColumn: 'Slot',
  itemColumn: 'Item',
  sourceColumn: 'Source',
  enchants: 'Enchants',
  consumables: 'Consumables',
  enchantsNote:
    'Only enchants the simulator can score are listed. Crit, hit and haste enchants are not scored yet.',
  consumablesNote: 'What the simulator gives this spec in its raid-ready run.',
  enchantColumn: 'Enchant',
  givesColumn: 'Gives',
  consumableColumn: 'Consumable',
  raceColumn: 'Race',
  factionColumn: 'Faction',
  racialsColumn: 'Racials',

  recommendedTag: 'Recommended',
  /** The Stat priority headline: "Hit cap: 16% · 3% from Elemental Precision, 13% from gear". */
  hitToCap: (
    percent: string,
    split?: { talentPercent: string; talent: string; gearPercent: string },
  ): string =>
    split === undefined
      ? `Hit cap: ${percent}%`
      : `Hit cap: ${percent}% · ${split.talentPercent}% from ${split.talent}, ${split.gearPercent}% from gear`,
  hitToCapWhite: (percent: string, white: string): string =>
    `Hit cap: ${percent}% (special attacks), ${white}% (white swings)`,
  /** The Stat priority note under the table: "Expertise cap: 6.5% (dodge)", parry added for a tank. */
  expertiseCap: (dodge: string, parry: string | undefined): string =>
    `Expertise cap: ${dodge}% (dodge)${parry === undefined ? '' : `, ${parry}% (parry)`}`,
  notSignificantNote: (labels: string, plural: boolean): string =>
    `${labels} shows as "Not significant": the simulator cannot tell ${plural ? 'them' : 'it'} from zero.`,
  weightColumn: (anchor: string): string => `Weight (${anchor} = 1.00)`,
  dpsPerPointColumn: 'DPS per point',
  unitLine: (stat: string, dps: string): string => `1 point of ${stat} = ${dps} DPS at level 60`,
  notSignificant: 'Not significant',
  levelingNewTitle: 'More points here than the range before',
  levelingLegend: 'Outlined: points added in this level range.',
  levelRangeLabel: (band: number): string => (band === 60 ? 'Level 60' : `Level ${band} to ${band + 9}`),
  levelRangePoints: (points: number, band: number): string =>
    band === 60 ? `${points} points` : `${points} to ${points + 9} points`,
  tocBarLabel: 'On this page',
  tocBarHint: 'Jump to a section',
  alliancePicksHeading: 'Alliance',
  hordePicksHeading: 'Horde',
  factionGroup: 'Faction',
  withoutFlask: 'Without a flask',
  replacesBothElixirs: 'replaces both elixirs',
  alsoUsed: 'Food, oil and potions',
  recommendedRacesLine: (name: string, faction: string, article: string, racials: string): string =>
    `${name} is ${article} ${faction} pick: ${racials}.`,
} as const;
