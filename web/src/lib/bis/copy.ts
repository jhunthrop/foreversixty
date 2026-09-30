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
  /** "Crafted: Blacksmithing" when the source's own name already IS the profession --
   *  every real crafted source today (`loot.json`'s `name` "Blacksmithing" and `profession`
   *  "blacksmithing") is the same word, differently cased, so this must never print the
   *  "Blacksmithing (blacksmithing)" a bare `${name} (${profession})` used to (bis-web-
   *  polish, 2026-09-30). "Crafted: <Profession> · <name>" only when they genuinely differ,
   *  for a future crafted source named by something other than its own profession (a
   *  specific recipe, say). Shared by the panel's source cell (`source-cell.ts`) and the
   *  item tooltip (`items/tooltip.ts`) so both say the same thing for the same source. */
  craftedSourceLabel: (name: string, profession?: string): string =>
    profession === undefined || profession.toLowerCase() === name.toLowerCase()
      ? `Crafted: ${name}`
      : `Crafted: ${capitalise(profession)} · ${name}`,
  vendorSourceLabel: (npc: string): string => `Vendor: ${npc}`,
  repSourceLabel: (factionName: string, standing?: string): string =>
    standing === undefined ? factionName : `${factionName} (${standing})`,
  placeSourceLabel: (place: string): string => place,
  /** "PvP rank 11 · Knight-Lieutenant · Alliance" -- never the bare bucket name "Rank 11"
   *  (third wow-player sweep defect, 2026-09-29): a player reads a source by what they'd
   *  actually see at the Quartermaster, the rank NUMBER and the reward's own TITLE and
   *  faction, not a pipeline id. Falls back to naming just the rank and faction when the
   *  rank falls outside the known ladder (`pvpRankTitle` returning `undefined` -- should
   *  not happen on real data, but never worse than an incomplete-but-true line). */
  pvpSourceLabel: (rank: number, faction: 'alliance' | 'horde'): string => {
    const title = pvpRankTitle(faction, rank);
    const factionLabel = faction === 'alliance' ? 'Alliance' : 'Horde';
    return title === undefined
      ? `PvP rank ${rank} · ${factionLabel}`
      : `PvP rank ${rank} · ${title} · ${factionLabel}`;
  },
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
  /** The alternative row's own gap-from-the-pick line: anything under 0.05 DPS reads as
   *  "same DPS" rather than "+0.0 DPS behind" (fix round 1, wow-player review -- the
   *  ranker's real dps_delta is a float that is essentially never an exact 0, so a strict
   *  `=== 0` check missed every practical tie; 0.05 is below the 0.1 the line itself
   *  rounds to, so nothing that would still show as a non-zero number reads as a tie). */
  alternativeGapLabel: (dpsDelta: number): string =>
    Math.abs(dpsDelta) < 0.05
      ? 'same DPS'
      : `${dpsDelta > 0 ? '+' : '−'}${Math.abs(dpsDelta).toFixed(1)} DPS`,
  /** An alternative's item level, and its required level only when that is above the
   *  band it's shown at (fix round 1, wow-player review): a requirement at or under the
   *  band's own level is already implied by the row being shown here at all, so naming it
   *  every time would be noise, not information. */
  alternativeMetaLabel: (itemLevel: number, requiredLevel: number, band: number): string =>
    requiredLevel > band ? `ilvl ${itemLevel} · needs ${requiredLevel}` : `ilvl ${itemLevel}`,
  weightsReferenceDpsLine: (label: string, dpsPerPoint: number): string =>
    `1 ${label} = ${dpsPerPoint.toFixed(2)} DPS`,
  weightsRowDpsLine: (label: string, weight: number, refAbbrev: string, dpsPerPoint: number): string =>
    `${label} ${weight.toFixed(2)} ${refAbbrev} · ${dpsPerPoint.toFixed(2)} DPS per point`,
  /** A rating-family row's native `title` hover (spec addendum 2, §C(2)): the client's own
   *  rating-per-percent conversion, so "Crit rating" doesn't leave a player guessing what
   *  "rating" means -- `14 Crit rating = 1% Crit`. */
  weightsRatingFactorLine: (label: string, ratingFactor: number): string =>
    `${ratingFactor} ${label} rating = 1% ${label}`,
  weightsNoEffect: 'No effect',
  /** The reference stat's own row in the weight rail's plain value column -- never a
   *  tautological "1 Attack power = 1.00 Attack power" or a bare "1.00" (the rail's own
   *  first line, `weightsReferenceDpsLine`, already states the real "1 <reference> = <n>
   *  DPS" conversion; this row exists only to anchor the bar chart, and reads as what it
   *  is). */
  weightsReferenceRowValue: 'Reference',
  /** The whole weight rail's line when the band carries `weights_reason` (spec addendum
   *  §A): a band whose reference measurement was not positive beyond its own error never
   *  gets a bar list at all (`panel-view.ts`'s `bandInfosFor` picks this over both of the
   *  rail's ordinary branches) -- one honest sentence instead of a list of numbers this
   *  band's own sim could not stand behind. Every `weights_reason` value maps to this same
   *  line today (only one exists in real data); a future, more specific code gets its own
   *  line only when a person writes one, never a fabricated one in the meantime. */
  weightsUnmeasuredLine:
    "Weights couldn't be measured for this gear set. The picks below are still real sim results.",
  /** The empty off-hand row when the main hand is a two-hander -- never
   *  `noKnownSourceForSlot`, which would read as a data gap rather than the game rule it
   *  actually is (wow-player fix 4). */
  twoHanderEquippedLabel: 'Two-hander equipped',

  // --- evidence, flags and empty reasons (fourth wow-player sweep, day 3) -----------------
  /** The row's own evidence line, one muted line under the source line, in player words
   *  rather than the ranker's own `swap_note` sentence -- `panel-view.ts`'s `parseSwapNote`
   *  supplies the item name and the two DPS numbers, this row's own pick first. */
  evidenceLine: (itemName: string, pickDps: number, altDps: number): string =>
    `Sim-checked against ${itemName}: ${pickDps.toFixed(1)} vs ${altDps.toFixed(1)} DPS`,
  /** The row's own evidence line when the slot carries a real `dps_delta` (bis-ranker-
   *  integrity-5): the one number that stays true regardless of which snapshot `swap_note`'s
   *  own two absolute numbers were measured at (`panel-view.ts`'s own doc) -- preferred over
   *  `evidenceLine` above whenever `dps_delta` is present, so the row never shows two
   *  absolute numbers that can silently disagree with the band's own header `set_dps`. */
  evidenceLineDelta: (itemName: string, delta: number): string =>
    `Sim-checked against ${itemName}: +${delta.toFixed(1)} DPS`,
  /** The verified glyph's own title when a pick carries `sim_dps` but no `swap_note` --
   *  still a real sim result (a trinket/proc/weapon-pair tournament winner), just not one
   *  phrased as a swap against a named runner-up. */
  simDpsVerifiedTitle: (dps: number): string =>
    `Confirmed by a full sim: ${dps.toFixed(1)} DPS with this item`,
  effectUnmodelledTag: 'effect not simulated',
  effectUnmodelledTitle:
    "This item's proc or use effect is not modelled yet; it was ranked on its stats alone",
  lowValueTag: 'best available',
  lowValueTitle: 'No sourced weapon at this level adds DPS; this is the best by item level',
  /** An item sourced from classic-db's 1.12 tables rather than the client's own shipped
   *  tables (spec addendum §B, `Item.client_unconfirmed`) -- a data-confidence flag on the
   *  item itself, never conflated with `sourceLines` (where you GET the item) or coloured
   *  like an achievement (`.new-pill`'s gold): this marks a limitation of the data, so it
   *  borrows the same quiet `.flag-tag` family `effectUnmodelledTag`/`lowValueTag` already
   *  use. Never any form of the word "confirmed" (wow-player fix, day 3): this row already
   *  carries `VerifiedGlyph`'s own "Confirmed by a Top Gear simulation pass" title and its
   *  own "Unverified" empty state, so a second, unrelated use of "confirmed" on the same row
   *  reads as a contradiction of the verified glyph rather than a comment on the item's data
   *  provenance -- the data field name (`client_unconfirmed`) and every testid keep their
   *  names unchanged; only the two player-facing strings change. */
  clientUnconfirmedTag: '1.12 stats',
  clientUnconfirmedTitle: "Stats from the original 1.12 tables; not yet seen in Forever's client",
  /** `empty_reason` copy, one line per value the ranker publishes -- `no_sourced_item` and
   *  any value this page does not recognise both fall back to `noKnownSourceForSlot`'s own
   *  text (spec's "unknown -> the last"), never a fabricated reason. */
  emptyReasonNoDpsValue: 'Nothing sourced at this level helps your DPS',
  emptyReasonEffectNotModelled: "Relic effects aren't simulated yet",
} as const;

/** `dwarf` -> `Dwarf`: the pipeline's own race strings are not reliably capitalised (owner
 *  screenshot finding, 2026-09-29), and this is the one place every `/bis` race line reads
 *  one, so every caller gets the fix for free rather than re-capitalising it themselves. */
function capitalise(word: string): string {
  return word.length === 0 ? word : word[0]!.toUpperCase() + word.slice(1);
}

// --- pvp rank titles (third wow-player sweep defect, 2026-09-29) --------------------------
// Vanilla's own Alliance/Horde PvP rank ladders, ranks 1-14 in order -- mirrors
// data/pipeline/loot/pvp_faction.py's ALLIANCE_TITLES/HORDE_TITLES, the primary source for
// both. Blizzard's own client `RequiredPVPRank` column (loot.json's `LootSource.rank`) is
// these ranks + 4, which `pvpRankTitle` undoes.
const ALLIANCE_PVP_TITLES = [
  'Private',
  'Corporal',
  'Sergeant',
  'Master Sergeant',
  'Sergeant Major',
  'Knight',
  'Knight-Lieutenant',
  'Knight-Captain',
  'Knight-Champion',
  'Lieutenant Commander',
  'Commander',
  'Marshal',
  'Field Marshal',
  'Grand Marshal',
] as const;

const HORDE_PVP_TITLES = [
  'Scout',
  'Grunt',
  'Sergeant',
  'Senior Sergeant',
  'First Sergeant',
  'Stone Guard',
  'Blood Guard',
  'Legionnaire',
  'Centurion',
  'Champion',
  'Lieutenant General',
  'General',
  'Warlord',
  'High Warlord',
] as const;

/** loot.json's own pvp source `rank` (Blizzard's client RequiredPVPRank, 5-18) -> the
 *  in-game rank title for `faction` ("Knight-Lieutenant" for alliance rank 11). `undefined`
 *  for a rank outside the ladder -- should not happen; loot.json only ever writes 5-18, but
 *  a caller sees a title-less line rather than an out-of-bounds crash if it ever did. */
export function pvpRankTitle(faction: 'alliance' | 'horde', rank: number): string | undefined {
  const titles = faction === 'alliance' ? ALLIANCE_PVP_TITLES : HORDE_PVP_TITLES;
  return titles[rank - 5];
}
