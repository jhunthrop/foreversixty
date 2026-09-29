// web/src/lib/bis/copy.ts
// Every string the /bis pages show, in one module (the site's own pattern -- see
// lib/planner/copy.ts and lib/sim/copy.ts).
export const bisCopy = {
  navLabel: 'Leveling BiS',
  indexTitle: 'Leveling BiS',
  indexDescription:
    'The best gear you can wear at every level band, from the simulator: one list per class and spec, built once and never per character.',
  indexIntro:
    'A band is "the best you can wear at that level" from a named source -- a quest, a vendor, a dungeon, a drop you can farm. Pick a class to see its specs.',
  noDataYet: 'No leveling BiS list yet for this spec.',
  noDataYetBody:
    'The simulator has not ranked this spec’s gear across the leveling bands yet. Check back once the nightly ranking run covers it.',
  backToIndex: 'Back to Leveling BiS',
  classGuideLink: 'See leveling BiS gear for this class',
  factionAlliance: 'Alliance',
  factionHorde: 'Horde',
  levelLabel: (band: number): string => `Level ${band}`,
  bandPillGroupLabel: 'Level band',
  bandHeading: (band: number): string => `Best gear at level ${band}`,
  newAtBand: (band: number): string => `New at ${band}`,
  newAtBandEmpty: 'Nothing changed from the previous band.',
  newAtBandFirst: 'The first band -- everything here is new.',
  noKnownSourceForSlot: 'No sourced item at this level yet',
  slotHeading: 'Slot',
  itemHeading: 'Item',
  levelHeading: 'Item level',
  sourceHeading: 'Source',
  setDpsLabel: 'Set DPS',
  statWeightsLabel: 'Stat weights',
  talentsLabel: 'Talents',
  raceLabel: 'Race',
  talentPointsLabel: (points: number): string => `${points} point${points === 1 ? '' : 's'} spent`,
  verified: 'Verified',
  unverified: 'Unverified',
  verifiedTitle: 'Confirmed by a Top Gear simulation pass at this band.',
  unverifiedTitle: 'Ranked by stat weights only; not yet settled by a Top Gear pass.',
  swapNoteSummary: 'Swap note',
  noSourceCount: (count: number): string =>
    count === 0
      ? ''
      : `${count} item${count === 1 ? '' : 's'} at this band had no known source and ${count === 1 ? 'is' : 'are'} left out.`,
  questFactionBadge: (faction: 'alliance' | 'horde'): string =>
    faction === 'alliance' ? 'Alliance quest' : 'Horde quest',
  generatedFrom: (engineVersion: string): string => `Simulated against engine ${engineVersion}`,
  metaDescription: (specName: string, className: string): string =>
    `The best gear for a leveling ${className} ${specName} at every level band, ranked by the simulator and verified with Top Gear.`,

  // --- source cell (step 1) ---------------------------------------------------------------
  questSourceLabel: (questName: string): string => `Quest: ${questName}`,
  questLevelLabel: (level: number): string => `Level ${level}`,
  dungeonSourceLabel: (instance: string, boss?: string): string =>
    boss === undefined ? instance : `${instance} · ${boss}`,
  craftedSourceLabel: (profession: string): string => `Crafted: ${profession}`,
  vendorSourceLabel: (npc: string): string => `Vendor: ${npc}`,
  repSourceLabel: (factionName: string, standing?: string): string =>
    standing === undefined ? factionName : `${factionName} (${standing})`,
  placeSourceLabel: (place: string): string => place,
  runnerUpBeatBy: (dps: number): string => `+${dps.toFixed(1)} DPS`,
  runnerUpTitle: (name: string, higherDps: number, lowerDps: number): string =>
    `${name} measured higher at this band: ${higherDps.toFixed(1)} vs ${lowerDps.toFixed(1)} set DPS.`,

  // --- band navigation (step 2) ------------------------------------------------------------
  bandStripGroupLabel: 'Jump to level',
  bandNewCount: (count: number): string => `${count} new`,
  changedSinceHeading: (previousBand: number): string => `What changed since level ${previousBand}`,
  changedSinceEmpty: 'Nothing changed from the previous band.',
  changedSinceFirst: 'The first band -- everything here is new.',
  wasLabel: 'Was',
  nowLabel: 'Now',
  newRowMarker: 'New',
  runnerUpLabel: 'Runner-up',

  // --- header (step 3) ----------------------------------------------------------------------
  indexSpecDps60: (dps: number): string => `Level 60: ${dps.toFixed(1)} DPS`,
} as const;
