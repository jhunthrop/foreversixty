// web/src/lib/bis/copy.ts
// Every string the /bis pages show, in one module (the site's own pattern -- see
// lib/planner/copy.ts and lib/sim/copy.ts). `craftedSourceLabel`/`pvpSourceLabel` (and
// `pvpRankTitle`, re-exported below for existing importers) live in
// `lib/items/source-labels.ts`, not here, and are only re-attached to this object -- see
// that file's own header for why (the item tooltip needs those two functions but must
// never pay for the rest of this much larger, /bis-page-only object).
import { craftedSourceLabel, pvpRankTitle, pvpSourceLabel } from '../items/source-labels';
import type { RateUnit } from './types';

export { pvpRankTitle };

/** `1 buff`, `2 buffs`. */
function countOf(count: number, noun: string): string {
  return `${count} ${noun}${count === 1 ? '' : 's'}`;
}

/** What the unit measures, in a sentence: damage for DPS, healing for HPS. */
function outcomeNoun(unit: RateUnit): string {
  return unit === 'HPS' ? 'healing' : 'damage';
}

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
  // A non-breaking space between the label word and its number (here and in every other
  // "Word N" metadata token below) -- a source/meta line wraps at its own " · " separators
  // on a narrow phone (bis rebuild spec §6), never mid-token; a plain space let the browser
  // orphan the number on its own line ("Quest: The Defias Brotherhood · Level" / "22",
  // wow-player review round 2, 390px capture). Never a `white-space: nowrap` on the line's
  // own outer element instead: several of these lines (`alternativeMetaLabel` below) join
  // two such tokens with " · ", which must still be free to wrap between them.
  questLevelLabel: (level: number): string => `Level\u00a0${level}`,
  /** The row's small item-level figure, labelled so a bare number never has to be guessed at. */
  itemLevelShort: (level: number): string => `ilvl\u00a0${level}`,
  dungeonSourceLabel: (instance: string, boss?: string): string =>
    boss === undefined ? instance : `${instance} · ${boss}`,
  /** "Crafted: Blacksmithing" when the source's own name already IS the profession --
   *  every real crafted source today (`loot.json`'s `name` "Blacksmithing" and `profession`
   *  "blacksmithing") is the same word, differently cased, so this must never print the
   *  "Blacksmithing (blacksmithing)" a bare `${name} (${profession})` used to (bis-web-
   *  polish, 2026-09-30). "Crafted: <Profession> · <name>" only when they genuinely differ,
   *  for a future crafted source named by something other than its own profession (a
   *  specific recipe, say). Shared by the panel's source cell (`source-cell.ts`) and the
   *  item tooltip (`items/tooltip.ts`) so both say the same thing for the same source --
   *  its own implementation lives in `items/source-labels.ts`, this file's own import. */
  craftedSourceLabel,
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
  pvpSourceLabel,
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
  indexSpecDps60: (dps: number, unit: RateUnit = 'DPS'): string => `Level 60: ${dps.toFixed(1)} ${unit}`,

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
  setDpsDelta: (delta: number, unit: RateUnit = 'DPS'): string =>
    `${delta >= 0 ? '+' : ''}${delta.toFixed(1)} ${unit} since the last band`,
  /** Shown under `setDpsLabel`'s big number when `BisBand.set_dps_partial` is true (spec
   *  addendum 3, §E): `count` always equals the number of rows this same band actually
   *  renders with `notSimCheckedTag` (`panel-view.ts`'s `BandInfo.setDpsPartialCount`,
   *  computed off the same `rows` array the paperdoll renders), so the number named here
   *  always matches a row a player can go find and read for themselves. */
  setDpsPartialNote: (count: number, unit: RateUnit = 'DPS'): string =>
    `Set ${unit} leaves out ${count} pick${count === 1 ? '' : 's'} marked not sim-checked`,
  raceTalentsLine: (race: string, points: number): string =>
    `${capitalise(race)} · ${points} talent point${points === 1 ? '' : 's'} spent`,

  // --- character panel redesign (bis-character-panel, 2026-09-29) -------------------------
  alternativesLabel: 'Also:',
  /** The alternative row's own gap-from-the-pick line: anything under 0.05 DPS reads as
   *  "same DPS" rather than "+0.0 DPS behind" (fix round 1, wow-player review -- the
   *  ranker's real dps_delta is a float that is essentially never an exact 0, so a strict
   *  `=== 0` check missed every practical tie; 0.05 is below the 0.1 the line itself
   *  rounds to, so nothing that would still show as a non-zero number reads as a tie). */
  alternativeGapLabel: (dpsDelta: number, unit: RateUnit = 'DPS'): string =>
    Math.abs(dpsDelta) < 0.05
      ? `same ${unit}`
      : `${dpsDelta > 0 ? '+' : '−'}${Math.abs(dpsDelta).toFixed(1)} ${unit}`,
  /** An alternative's item level, and its required level only when that is above the
   *  band it's shown at (fix round 1, wow-player review): a requirement at or under the
   *  band's own level is already implied by the row being shown here at all, so naming it
   *  every time would be noise, not information. */
  alternativeMetaLabel: (itemLevel: number, requiredLevel: number, band: number): string =>
    requiredLevel > band ? `ilvl\u00a0${itemLevel} · needs\u00a0${requiredLevel}` : `ilvl\u00a0${itemLevel}`,
  /** A rating-family row's native `title` hover (spec addendum 2, §C(2)): the client's own
   *  rating-per-percent conversion, so "Crit rating" doesn't leave a player guessing what
   *  "rating" means -- `14 Crit rating = 1% Crit`. */
  weightsRatingFactorLine: (label: string, ratingFactor: number): string =>
    `${ratingFactor} ${label} rating = 1% ${label}`,
  /** The whole weight rail's line when the band carries `weights_reason` (spec addendum
   *  §A): a band whose reference measurement was not positive beyond its own error never
   *  gets a bar list at all (`panel-view.ts`'s `bandInfosFor` picks this over both of the
   *  rail's ordinary branches) -- one honest sentence instead of a list of numbers this
   *  band's own sim could not stand behind. Every `weights_reason` value maps to this same
   *  line today (only one exists in real data); a future, more specific code gets its own
   *  line only when a person writes one, never a fabricated one in the meantime. */
  weightsUnmeasuredLine:
    "Weights couldn't be measured for this gear set. The picks below are still real sim results.",
  /** The scale rail's own first line (this lane's brief, bis-weights-simc, owner: "we need
   *  to make the stat weights align with simcraft stat weights output"): `topLabel` is
   *  `scale_reference_stat`'s own display label -- the single highest-weighted PER-POINT
   *  stat every other row's `scale_factor` is normalized against (never a haste stat). */
  weightsScaleNote: (topLabel: string, unit: RateUnit = 'DPS'): string =>
    `Per point of stat, normalized to ${topLabel} = 1.00, with ${unit} per point and the sim error.`,
  /** Haste's own one-line caption (owner correction, 2026-09-30, after player review: haste
   *  is not a per-point stat and never appears as a table row) -- `scaleFactor` is the same
   *  number `band.haste_scale_factor` (or its client-side fallback) publishes. `noItemAtBand`
   *  is `!band.haste_on_items` (`sim/cmd/leveling-bis/report.go`'s own `bandHasHasteCandidate`
   *  -- a plain inventory check, "does any candidate this band considered carry a haste
   *  stat"), never a haste row's own `insignificant` flag: that answers a different,
   *  statistical question (was the sweep's own sample clean enough to trust the number) and
   *  conflating the two produced this line's own second-draft bug -- a significant-but-
   *  noisy haste weight on a band that DOES have haste gear rendered the redundant
   *  "Haste: 1.58 per 1%, per 1%" (owner fix, 2026-09-30, found on screenshot review). The
   *  suffix now only ever appears when no candidate item actually carries the stat. */
  weightsHasteCaption: (scaleFactor: number, noItemAtBand: boolean): string =>
    `Haste: ${scaleFactor.toFixed(2)} per 1%${
      noItemAtBand ? ', not in the table because no item at this band has it' : ''
    }`,
  /** A per-point row's own `scale_factor`/`dps_per_point`/`scale_error` the sweep's own 25%-
   *  of-value error bar could not clear (`weight.insignificant`) -- greyed on the page
   *  (`weight-bar-row-faint`), but the row itself is never dropped (tenet 4). */
  weightsNotSignificant: 'Not significant',
  /** The crit/hit rating-per-percent conversion, stated once as a standing caption rather
   *  than only in each row's own hover (`weightsRatingFactorLine`) -- this ruleset's own
   *  `gametables/combatratings.txt` conversion is fixed at level 60 regardless of the band a
   *  page is showing (`sim/cmd/leveling-bis/data.go`'s own `loadRatingFactors` doc), so one
   *  note for the whole rail, not per band. */
  weightsRatingConversionNote: 'At level 60, 1% crit = 14 rating and 1% hit = 10 rating.',
  /** What every BiS sim is run against, under the preset caption. Static copy, so it may
   *  name the target level and fight length; no data-rendered figure appears here. */
  simTargetNote:
    'Simulated against a level-63 target of no creature type over 180 seconds; slaying bonuses such as attack power against Undead count for nothing here.',
  /** The hit-to-cap line under the weights (`hit-cap.ts`'s `hitCapLine`); `white` is
   *  `undefined` for a spec that does not dual wield. Percent strings arrive formatted. */
  hitToCapLine: (specials: string, white: string | undefined): string =>
    `Hit to cap: ${specials}% for specials${white === undefined ? '' : `, ${white}% for white swings`}`,
  /** The caster's line: `school` is the distance for the spells a talent adds hit to, and
   *  `undefined` when no talent does. Percent strings arrive formatted. */
  spellHitToCapLine: (spell: string, school: string | undefined): string =>
    `Spell hit to cap: ${spell}%${school === undefined ? '' : ` (school ${school}%)`}`,
  /** The hover on the caster's line. */
  spellHitToCapTitle:
    'Spell hit is worth its full weight until the cap and nothing past it: spells miss 17% of the time against a boss three levels up, and 1% always remains. School talents such as Elemental Precision add hit to their own spells only. The weights above are per rating point: 10 hit rating is 1%.',
  /** The hover on that line: hit is worth its full weight until the cap and nothing past it,
   *  and the rail's weights are per rating point. */
  hitToCapTitle:
    'Hit is worth its full weight until the cap and nothing past it. The weights above are per rating point: 10 hit rating and 14 crit rating are 1%.',
  /** The addon cross-sell under the rail (owner correction, 2026-09-30: "no Pawn anywhere...
   *  under the rail, replace the string box with one secondary button"). The addon feature
   *  itself is a separate lane; `weightsAddonButton`'s own link is `/addon` until it lands. */
  weightsAddonBlurb: 'The Forever Sixty addon rates every item in your bags and tooltips with these weights.',
  weightsAddonButton: 'Use these weights in the addon',
  /** The empty off-hand row when the main hand is a two-hander -- never
   *  `noKnownSourceForSlot`, which would read as a data gap rather than the game rule it
   *  actually is (wow-player fix 4). */
  twoHanderEquippedLabel: 'Two-hander equipped',

  // --- evidence, flags and empty reasons (fourth wow-player sweep, day 3) -----------------
  /** The row's own evidence line, one muted line under the source line, in player words
   *  rather than the ranker's own `swap_note` sentence -- `panel-view.ts`'s `parseSwapNote`
   *  supplies the item name and the two DPS numbers, this row's own pick first. */
  evidenceLine: (itemName: string, pickDps: number, altDps: number, unit: RateUnit = 'DPS'): string =>
    `Sim-checked against ${itemName}: ${pickDps.toFixed(1)} vs ${altDps.toFixed(1)} ${unit}`,
  /** The row's own evidence line when the slot carries a real `dps_delta` (bis-ranker-
   *  integrity-5): the one number that stays true regardless of which snapshot `swap_note`'s
   *  own two absolute numbers were measured at (`panel-view.ts`'s own doc) -- preferred over
   *  `evidenceLine` above whenever `dps_delta` is present, so the row never shows two
   *  absolute numbers that can silently disagree with the band's own header `set_dps`. */
  evidenceLineDelta: (itemName: string, delta: number, unit: RateUnit = 'DPS'): string =>
    `Sim-checked against ${itemName}: +${delta.toFixed(1)} ${unit}`,
  /** The verified glyph's own title when a pick carries `sim_dps` but no `swap_note` --
   *  still a real sim result (a trinket/proc/weapon-pair tournament winner), just not one
   *  phrased as a swap against a named runner-up. */
  simDpsVerifiedTitle: (dps: number, unit: RateUnit = 'DPS'): string =>
    `Confirmed by a full sim: ${dps.toFixed(1)} ${unit} with this item`,
  /** `sim_status === 'not_in_sim'` (spec addendum 3, §D): this build's sim database does
   *  not carry the item at all, so the pick is ranked by stat weights only, with no full
   *  sim run behind it -- the umbrella statement, shown before `effectUnmodelledTag` when
   *  both fire on the same row (that row's own proc/use effect being unmodelled is the
   *  narrower, more specific fact). Same quiet `.flag-tag` family: no gold, no red. */
  notSimCheckedTag: 'not sim-checked',
  notSimCheckedTitle:
    "This build's sim database doesn't carry this item; the pick is ranked by stat weights only, with no full sim run behind it",
  effectUnmodelledTag: 'effect not simulated',
  effectUnmodelledTitle:
    "This item's proc or use effect is not modelled yet; it was ranked on its stats alone",
  lowValueTag: 'best available',
  lowValueTitle: (unit: RateUnit = 'DPS'): string =>
    `No sourced weapon at this level adds ${unit}; this is the best by item level`,
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
  emptyReasonNoDpsValue: (unit: RateUnit = 'DPS'): string =>
    `Nothing sourced at this level helps your ${unit}`,
  emptyReasonEffectNotModelled: "Relic effects aren't simulated yet",

  // --- "The list" rebuild (bis rebuild spec, 2026-09-30) -----------------------------------
  headerEyebrow: 'Best in slot · leveling',
  /** `{band}` -> `"20 to 29"` through `"50 to 59"`, and plain `"60"` for the top band --
   *  design system's own explicit rule, never `"60 to 60"` (spec §4.A). Every band in this
   *  contract is a ten-level bracket (`bands` are `[20, 30, 40, 50, 60]`, one entry per
   *  lower bound) except the last, which has no bracket to show. */
  bandRangeLabel: (band: number): string => (band === 60 ? '60' : `${band} to ${band + 9}`),
  /** The header's one summary sentence (spec §4.A) -- `bandHigh` is `undefined` only for
   *  the plain-60 band, which drops the whole "from X to Y" clause rather than say "at 60
   *  to 60" or invent a range the contract does not have. */
  headerSummary: (bandLow: number, bandHigh: number | undefined, unit: RateUnit = 'DPS'): string =>
    bandHigh === undefined
      ? `The gear that raises your ${outcomeNoun(unit)} most at ${bandLow}, ranked by the simulator with every item equipped, and where each piece comes from.`
      : `The gear that raises your ${outcomeNoun(unit)} most from ${bandLow} to ${bandHigh}, ranked by the simulator with every item equipped, and where each piece comes from.`,
  thisSetLabel: 'This set',
  /** "DPS on the training dummy, level 29 Troll, 11 talent points" -- `characterLevel` is
   *  the band's own training-dummy level (never a signed-in character's real level, spec
   *  §4.C.1's own explicit warning: the two are unrelated facts that happen to coincide in
   *  the mock's own worked example). */
  thisSetFigureLine: (characterLevel: number, race: string, talentPoints: number): string =>
    `DPS on the training dummy, level ${characterLevel} ${capitalise(race)}, ${talentPoints} talent point${talentPoints === 1 ? '' : 's'}`,
  simmedLine: (generatedAtDate: string): string =>
    `Simmed ${generatedAtDate}, every pick measured with the item equipped`,
  openInSimulator: 'Open in simulator',
  talentsInPlanner: 'Talents in planner',
  statWeightsAtBand: (bandLabel: string): string => `Stat weights at ${bandLabel}`,
  playItHeading: (bandLabel: string): string => `Play it · ${bandLabel}`,
  /** The rotation line's own icon placeholder (spec §4.C.3) -- this build carries no spell-
   *  icon-by-id data for a trained (non-talent) ability at all (verified: `spells.json`,
   *  `spellranks.json` and `spellconst/<class>.json` all checked), so every line falls back
   *  to a plain bordered square. Its own title, never `noKnownSourceForSlot`'s wording
   *  (that line describes an EMPTY GEAR SLOT with no pick at all -- a different fact from
   *  "this ability's icon isn't resolved yet"; ux-designer review round 1). */
  iconNotAvailableYet: 'Icon not available yet',
  openGuideLink: (specName: string): string => `Open the ${specName} guide`,
  presetControlLabel: 'Conditions',
  presetOptionRaid: 'Raid-ready',
  presetOptionBare: 'Bare',
  presetRaidCaption: (label: string, buffs: number, debuffs: number, consumes: number): string =>
    `${label}: ${countOf(buffs, 'buff')}, ${countOf(debuffs, 'debuff')} and ${countOf(consumes, 'consumable')}.`,
  presetBareCaption: 'Bare character: no raid buffs, debuffs or consumables.',
  presetDetailsSummary: 'What is included',
  presetBuffsHeading: 'Buffs',
  presetDebuffsHeading: 'Debuffs',
  presetConsumesHeading: 'Consumables',
  theListHeading: 'The list',
  hoverOrTapCaption: 'Hover or tap an item for its stats',
  slotHeaderLabel: 'Slot',
  pickHeaderLabel: 'Best in slot · where it comes from',
  runnersUpHeaderLabel: (unit: RateUnit = 'DPS'): string => `Runners-up · ${unit} vs the pick`,
  youHeaderLabel: 'You',
  whereToGetItHeading: 'Where to get it',
  newAtThisBandHeading: 'New at this band',
  /** `source_kind === 'world_drop'`'s own group label for "Where to get it" (spec §4.E) --
   *  never `worldDropSourceLabel`'s `"World drop (BoE)"` phrasing, which is pipeline jargon
   *  for the same row's own individual source line; this is the one place on the page
   *  naming the actual action a player takes (the auction house), for a whole GROUP of
   *  items rather than one row. */
  worldDropGroupLabel: 'World drop (auction house)',
  groupCountSuffix: (count: number): string => `×${count}`,
  firstBandEveryPickNew: 'This is the first band — every pick here is new.',
  /** Sentence 1 of "New at this band" (spec §4.E) for every band past the first -- `parts`
   *  is the already-built, ordered list of instance names and `"the <faction> rewards"`
   *  clauses (`joinWithAnd` below joins them into one clause). */
  reachingBandOpens: (bandLow: number, joinedParts: string, n: number, total: number): string =>
    `Reaching ${bandLow} opens ${joinedParts}: ${n} of the ${total} picks come from there.`,
  /** Sentence 2, the next band's own preview -- `joinedParts` is `''` when the next band
   *  opens nothing a dungeon/rep clause would name (every new pick there is a vendor/quest/
   *  world/crafted source), in which case the sentence starts straight from the slot clause
   *  (or is empty entirely when neither clause has anything to say). */
  nextBandLine: (bandLabel: string, joinedParts: string, helpfulSlotLabel: string | undefined): string => {
    const head = `Next band (${bandLabel})`;
    const slotClause = helpfulSlotLabel === undefined ? '' : `the first ${helpfulSlotLabel} that helps`;
    if (joinedParts === '' && slotClause === '') return `${head}.`;
    if (joinedParts === '') return `${head}: ${slotClause}.`;
    if (slotClause === '') return `${head}: ${joinedParts}.`;
    return `${head}: ${joinedParts}, and ${slotClause}.`;
  },
  seeBandLink: (bandLabel: string): string => `See ${bandLabel}`,
  /** The off-hand's own empty row when the main hand is a two-hander, NAMING the weapon
   *  (spec §4.D) -- replaces the older, generic `twoHanderEquippedLabel` for this rebuild's
   *  list row (kept above, unchanged, for any other surface that still reads it). */
  twoHanderEquippedNamed: (mainHandName: string): string =>
    `${mainHandName} is a two-hander; the off hand is taken.`,
  /** `no_sourced_item`, this slot's own first empty occurrence in the file (spec §4.D) --
   *  `nextRealBand` is `undefined` when no later band in this same file ever sources the
   *  slot either, which drops the second sentence entirely rather than name a band that
   *  also turns out empty. Named as the plain band NUMBER ("comes at 30"), not the band's
   *  own range label -- verified against the approved mock's own exact wording ("No
   *  trinket you can get at 20 to 29 raises your damage. The first that does comes at
   *  30."): the band this sentence is ABOUT gets the range, the band a bare "comes at"
   *  points to does not. */
  noSourcedItemFirst: (
    slotLower: string,
    bandLabel: string,
    nextRealBand: number | undefined,
    unit: RateUnit = 'DPS',
  ): string =>
    nextRealBand === undefined
      ? `No ${slotLower} you can get at ${bandLabel} raises your ${outcomeNoun(unit)}.`
      : `No ${slotLower} you can get at ${bandLabel} raises your ${outcomeNoun(unit)}. The first that does comes at ${nextRealBand}.`,
  /** `no_sourced_item`, a LATER empty band for a slot this file's already named once (spec
   *  §4.D) -- `undefined` the same way `noSourcedItemFirst`'s own clause is: no later band
   *  ever sources it either. Same plain-band-number rule as above (mock: "Nothing here
   *  either until 30."). */
  noSourcedItemLater: (nextRealBand: number | undefined): string =>
    nextRealBand === undefined
      ? 'Nothing here helps at this level either.'
      : `Nothing here either until ${nextRealBand}.`,

  // --- Character card (bis rebuild spec §4.B) ----------------------------------------------
  yourCharacterLabel: 'Your character',
  switchLabel: 'Switch',
  syncedRelative: (relative: string): string => `synced ${relative}`,
  notSyncedYetLabel: 'Not synced yet',
  characterCardIdentityLine: (level: number, race: string): string => `${level} ${capitalise(race)}`,
  upgradesLabel: 'Upgrades',
  togetherLabel: 'Together',
  togetherDpsValue: (delta: number): string => `+${delta.toFixed(1)} DPS`,
  equippedLabel: 'Equipped',
  equippedValue: (n: number, total: number): string => `${n} / ${total}`,
  sendListToAddon: 'Send this list to the addon',
  installTheAddon: 'Install the addon',
  installAddonToCompare: 'Install the addon to compare your gear against this list.',
  gearSyncNotHereYet: "Gear sync isn't here yet — install the addon to compare your gear against this list.",
  characterCardLoadError: 'Your character did not load.',
  /** The list row's own "You" column (spec §4.D/§7 `GearRow`'s `you` slot) -- unused today
   *  (the column itself is omitted entirely while the worn-gear gap is open, §9), kept here
   *  so the component that renders it needs no second copy module once a future lane
   *  supplies real data. */
  equippedTickLabel: 'equipped',
  youGainOverWorn: (dpsDelta: number, wornItemName: string, unit: RateUnit = 'DPS'): string =>
    `+${dpsDelta.toFixed(1)} ${unit} over your ${wornItemName}`,

  // --- healer band panel (healer BiS lane) ---------------------------------------------------
  // A healing sim ranks gear under a STATED incoming-damage profile and never says one healer
  // beats another: every healer figure below is worded with the profile's name beside it.
  healerFigureUnit: 'HPS',
  healerFigureCaption: (profileLabel: string): string => `Effective healing per second under ${profileLabel}`,
  /** Fallback when a healer band names no profile at all (a file the ranker published without
   *  `heal_profile`/`profile`): the figure still says it depends on one. */
  healerProfileUnstated: 'an unstated incoming-damage profile',
  healerOverhealLabel: 'Overheal',
  healerOverhealValue: (share: number): string => `${Math.round(share * 100)}%`,
  healerOverhealLine: (share: number): string => `${Math.round(share * 100)}% overheal`,
  healerManaLabel: 'Mana lasts',
  healerManaWholeFightValue: 'Whole fight',
  healerManaWholeFightLine: (duration: string | undefined): string =>
    duration === undefined ? 'Mana lasts the whole fight' : `Mana lasts the whole ${duration} fight`,
  healerManaOutValue: (clock: string): string => clock,
  healerManaOutLine: (clock: string): string => `Out of mana at ${clock}`,
  healerHpmLabel: 'Healing per mana',
  healerHpmValue: (hpm: number): string => hpm.toFixed(1),
  healerHpmLine: (hpm: number): string => `${hpm.toFixed(1)} healing per mana`,
  healerProfileSummary: 'What this profile assumes',
  healerProfileTankHeading: 'Tank',
  healerProfileTankLine: (health: number, hit: number, swingSeconds: number): string =>
    `${health.toLocaleString('en-US')} health, a ${hit.toLocaleString('en-US')} hit every ${swingSeconds} seconds`,
  healerProfileMembersHeading: 'Raid members',
  healerProfileMembersLine: (health: number): string => `${health.toLocaleString('en-US')} health each`,
  healerProfilePulseHeading: 'Raid pulses',
  healerProfilePulseLine: (damage: number, intervalSeconds: number, members: number): string =>
    `${damage.toLocaleString('en-US')} damage to ${members} members every ${intervalSeconds} seconds`,
  healerProfileFightHeading: 'Fight',
  healerProfileFightLine: (duration: string): string => `${duration} long`,
} as const;

/** `dwarf` -> `Dwarf`: the pipeline's own race strings are not reliably capitalised (owner
 *  screenshot finding, 2026-09-29), and this is the one place every `/bis` race line reads
 *  one, so every caller gets the fix for free rather than re-capitalising it themselves. */
function capitalise(word: string): string {
  return word.length === 0 ? word : word[0]!.toUpperCase() + word.slice(1);
}

/** `["A"]` -> `"A"`, `["A", "B"]` -> `"A and B"`, `["A", "B", "C"]` -> `"A, B and C"` --
 *  the one join "New at this band"'s two computed sentences use for an instance/rep list of
 *  any length, English's own list convention (no Oxford comma, one "and" before the last
 *  item). `[]` -> `""`, so a caller can always check for the empty string rather than a
 *  separate "were there any parts at all" branch. */
export function joinWithAnd(parts: readonly string[]): string {
  if (parts.length === 0) return '';
  if (parts.length === 1) return parts[0]!;
  return `${parts.slice(0, -1).join(', ')} and ${parts[parts.length - 1]}`;
}

/** The tank band's own words (stage 4 of the tanks job). A tank band's per-slot figures
 *  (`sim_dps`, `dps_delta`) and its per-point reference are tank-score points -- a composite
 *  of mitigation, risk and threat -- so they read "score", never "DPS". */
export const tankCopy = {
  /** The neutral unit word for every tank-score figure. */
  scoreWord: 'score',
  runnersUpHeaderLabel: 'Runners-up · score vs the pick',
  effectiveHealthLabel: 'Effective health',
  dtpsLabel: 'Damage taken per second',
  chanceOfDeathLabel: 'Chance of death',
  tpsLabel: 'Threat per second',
  effectiveHealthTitle: 'Hit points against the boss profile.',
  dtpsTitle: 'Damage taken per second at the boss profile. Lower is better.',
  chanceOfDeathTitle: 'Chance of dying during the simulated boss fight. Lower is better.',
  tpsTitle: 'Threat per second while tanking the boss profile.',
  tmiTitle:
    'Theck-Meloree Index: a risk index for the damage spikes a tank takes. Lower means a smoother, safer fight.',
  /** The tiny non-zero chance of death, so it never rounds to a misleading 0.0%. */
  chanceOfDeathBelowResolution: '<0.1%',
  secondaryLine: (tmi: string, ownDps: string): string => `TMI ${tmi} · own damage ${ownDps} DPS`,
  /** "Against the boss profile, level 60 Troll, 31 talent points". */
  figureLine: (characterLevel: number, race: string, talentPoints: number): string =>
    `Against the boss profile, level ${characterLevel} ${capitalise(race)}, ${talentPoints} talent point${talentPoints === 1 ? '' : 's'}`,
  /** The runner-up row's gap-from-the-pick line, in score points. */
  alternativeGapLabel: (scoreDelta: number): string =>
    Math.abs(scoreDelta) < 0.05
      ? 'same score'
      : `${scoreDelta > 0 ? '+' : '−'}${Math.abs(scoreDelta).toFixed(1)} score`,
  evidenceLineDelta: (itemName: string, delta: number): string =>
    `Sim-checked against ${itemName}: +${delta.toFixed(1)} score`,
  simScoreVerifiedTitle: (score: number): string =>
    `Confirmed by a full sim: ${score.toFixed(1)} score with this item`,
  weightsScaleNote: (topLabel: string): string =>
    `Per point of stat, normalized to ${topLabel} = 1.00, with score per point and the sim error.`,
  indexSpecSummary: (effectiveHealth: string): string => `Level 60: ${effectiveHealth} effective health`,
} as const;
