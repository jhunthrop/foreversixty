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
  /** The row's small item-level figure, labelled so a bare number never has to be guessed at. */
  itemLevelShort: (level: number): string => `ilvl ${level}`,
  dungeonSourceLabel: (instance: string, boss?: string): string =>
    boss === undefined ? instance : `${instance} · ${boss}`,
  craftedSourceLabel: (profession: string): string => `Crafted: ${profession}`,
  vendorSourceLabel: (npc: string): string => `Vendor: ${npc}`,
  repSourceLabel: (factionName: string, standing?: string): string =>
    standing === undefined ? factionName : `${factionName} (${standing})`,
  placeSourceLabel: (place: string): string => place,
  /** "40% from Lord Serpentis", "6% from Deadmines trash" -- classic-db's own drop
   *  chance (src-classicdb lane, 2026-09-29), shown in front of the place a source cell
   *  would otherwise just name plainly. */
  /** Rounded the way the client and Wowhead show it: whole percent, "<1%" below one, and
   *  never "0%" -- a chance of 0 means unknown and the caller omits the label. */
  dropChanceLabel: (chance: number, from: string): string =>
    `${chance < 1 ? '<1' : Math.round(chance)}% from ${from}`,
  worldDropSourceLabel: (levelMin?: number, levelMax?: number): string =>
    levelMin === undefined || levelMax === undefined
      ? 'World drop (BoE)'
      : `World drop (BoE) · levels ${levelMin}-${levelMax}`,

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

  // --- list-first redesign (bis-ux, 2026-09-29 -- owner: "still looks like shit") ----------
  levelScaleGroupLabel: 'Jump to level',
  newAtBandCount: (count: number): string => `${count} new`,
  newPillLabel: 'New',
  newPillTitle: (replaced: string | undefined): string =>
    replaced === undefined ? 'New this band -- nothing was equipped here before.' : `Replaces ${replaced}.`,
  verifiedGlyphTitle: 'Confirmed by a Top Gear simulation pass at this band.',
  unverifiedGlyphTitle: 'Ranked by stat weights only; not yet settled by a Top Gear pass.',
  upgradesSinceHeading: (count: number, previousBand: number): string =>
    count === 0
      ? `Nothing changed since level ${previousBand}`
      : `${count} upgrade${count === 1 ? '' : 's'} since level ${previousBand}`,
  newSlotLabel: 'New slot',
  setDpsDelta: (delta: number): string =>
    `${delta >= 0 ? '+' : ''}${delta.toFixed(1)} DPS since the last band`,
  raceTalentsLine: (race: string, points: number): string =>
    `${capitalise(race)} · ${points} talent point${points === 1 ? '' : 's'} spent`,

  // --- character panel redesign (bis-character-panel, 2026-09-29) -------------------------
  alternativesLabel: 'Also:',
  /** The alternative row's own gap-from-the-pick line: an exact tie (the ranker's own
   *  `dps_delta` 0) reads as "same DPS" rather than "+0.0 DPS behind" -- a signed zero is
   *  never a real distinction a player should have to parse. */
  alternativeGapLabel: (dpsDelta: number): string =>
    dpsDelta === 0 ? 'same DPS' : `${dpsDelta > 0 ? '+' : '−'}${Math.abs(dpsDelta).toFixed(1)} DPS`,
  weightsReferenceDpsLine: (label: string, dpsPerPoint: number): string =>
    `1 ${label} = ${dpsPerPoint.toFixed(2)} DPS`,
  weightsRowDpsLine: (label: string, weight: number, refAbbrev: string, dpsPerPoint: number): string =>
    `${label} ${weight.toFixed(2)} ${refAbbrev} · ${dpsPerPoint.toFixed(2)} DPS per point`,
  weightsNoEffect: 'No effect',
  /** The empty off-hand row when the main hand is a two-hander -- never
   *  `noKnownSourceForSlot`, which would read as a data gap rather than the game rule it
   *  actually is (wow-player fix 4). */
  twoHanderEquippedLabel: 'Two-hander equipped',
} as const;

/** `dwarf` -> `Dwarf`: the pipeline's own race strings are not reliably capitalised (owner
 *  screenshot finding, 2026-09-29), and this is the one place every `/bis` race line reads
 *  one, so every caller gets the fix for free rather than re-capitalising it themselves. */
function capitalise(word: string): string {
  return word.length === 0 ? word : word[0]!.toUpperCase() + word.slice(1);
}
