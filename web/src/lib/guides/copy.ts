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
  indexHonestyLine:
    'The beta caps at level 30. Level 60 play is a projection until we can confirm it against a live raid.',
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
  levelingBuild: 'Leveling build',
  raidBuild: 'Raid build',
  /** The one heading a spec shows when its raid build is its leveling build. */
  levelingAndRaidBuild: 'Leveling and raid build',
  loadLevelingBuildAriaLabel: 'Load the leveling build in the planner',
  simLevelingBuildAriaLabel: 'Sim the leveling build',
  loadRaidBuildAriaLabel: 'Load the raid build in the planner',
  simRaidBuildAriaLabel: 'Sim the raid build',
} as const;
