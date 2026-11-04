// Every user-visible string the guides lane's own components (GuideBuildTree,
// BuildActionButtons, RacePillRow, StatPriorityPills) render, reference voice throughout
// (design/DESIGN-SYSTEM.md: state the thing and stop).

/** The unit a role's per-point and set figures are in: a tank's per-point figures are
 *  tank-score points (mitigation, risk and threat combined), never damage. */
const ROLE_RATE_UNIT = { dps: 'DPS', healer: 'HPS', tank: 'tank score' } as const;

/** What a stat that adds nothing would have raised, by role. */
const ROLE_OUTPUT = { dps: 'damage', healer: 'healing', tank: 'survival or threat' } as const;

type GuideRole = keyof typeof ROLE_RATE_UNIT;

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
  /** The /guides index line, under the title: what every guide gives, for launch day on. */
  indexLine: 'Talents, rotation and gear for every spec, from your first talent point to the raid at 60.',
  roleLabel: (role: 'dps' | 'healer' | 'tank'): string =>
    role === 'dps' ? 'DPS' : role === 'healer' ? 'Healer' : 'Tank',
  readTheGuide: 'Read the guide',
  readTheGuideAriaLabel: (specName: string, className: string): string =>
    `Read the ${specName} ${className} guide`,
  yourSpecPill: 'Your spec',
  /** The class landing card's accessible name once it carries the "Your spec" pill. */
  yourSpecAriaLabel: (cardLabel: string): string => `${cardLabel}, your spec`,
  /** The class landing's visually hidden heading over the spec cards. */
  classLandingSpecsHeading: 'Specs',
  buildRailLabel: 'Build',
  /** The rail's Rotation card title; its band label is the head's meta (site alignment §5.4). */
  rotationRailTitle: 'Rotation',
  fullPriorityLink: 'Full priority ↓',
  statPriorityRailLabel: 'Stat priority',
  levelingPointsLabel: (points: number): string => `${points} point${points === 1 ? '' : 's'}`,
  loadInPlanner: 'Load in planner',
  loadBuildRailAriaLabel: 'Load this build in the planner (summary)',
  simBuildRailAriaLabel: 'Sim this build (summary)',
  loadRaidBuildRailAriaLabel: 'Load the raid build in the planner (summary)',
  simRaidBuildRailAriaLabel: 'Sim the raid build (summary)',
  /** Under "Rotation and priority" when the guide has its own raid build: the cards are worked
   *  out for it (its talents, cooldowns and cast times). */
  rotationForRaidBuild: 'For the raid build.',
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
  /** The one data-source line, in the page's footer beside the Updated stamp, never in the reading flow. */
  dataSourceLine: (build: string): string =>
    `Numbers from client build ${build}, checked against launch logs as they come in.`,
  headerMeta: (role: 'dps' | 'healer' | 'tank', updated: string): string =>
    `${role === 'dps' ? 'DPS' : role === 'healer' ? 'Healer' : 'Tank'} · Updated ${updated}`,
  strengths: 'Strengths',
  weaknesses: 'Weaknesses',
  inARaid: 'In a raid',
  keyTalents: 'Key talents',
  /** A key talent the two builds take at different ranks: "Level 60 3/3" and "Raid 2/3", each
   *  kept on one line. */
  keyTalentBuildRank: (build: 'leveling' | 'raid', rank: number, max: number): string =>
    `${build === 'leveling' ? 'Level 60' : 'Raid'} ${rank}/${max}`,
  opener: 'Opener',
  priority: 'Priority',
  cooldowns: 'Cooldowns',
  /** A rotation row for a talent only the raid build takes, or a row that needs one. */
  raidBuildTag: 'Raid build only',
  /** Said once above the lists when a row carries that tag. */
  raidBuildLegend:
    'Raid build only: needs a talent the level 60 build does not take. On the level 60 build, skip that row.',
  bestInSlot: 'Best in slot',
  levelRange: 'Level range',
  bisCaption: (race: string, faction: string): string => `Picks for a ${race} (${faction}).`,
  fullListLink: 'Full list and alternatives',
  preRaidNote:
    'These are the best picks before the first raid tier. Raid loot is not itemized in the client yet, and this list updates when it is.',
  notSimChecked: 'Ranked by stats only',
  slotColumn: 'Slot',
  itemColumn: 'Item',
  sourceColumn: 'Source',
  enchants: 'Enchants',
  consumables: 'Consumables',
  /** What the enchant table is: each slot's best flat-stat enchant by the band's weights
   *  (gear-extras.ts); crit, hit and haste enchants are not scored. wow-player SHIP 2026-11-04. */
  enchantsNote:
    'The enchant for each slot that adds the most of your weighted stats. Crit, hit and haste enchants are not included.',
  consumablesNote: 'What to bring to a raid.',
  enchantColumn: 'Enchant',
  givesColumn: 'Gives',
  consumableColumn: 'Consumable',
  raceColumn: 'Race',
  factionColumn: 'Faction',
  racialsColumn: 'Racials',

  recommendedTag: 'Recommended',
  /** The Stat priority headline: "Hit cap: 16% · 1% hit from BiS gear and buffs, 6% from
   *  Precision, 9% still to find". Caps
   *  arrive formatted with their own "%" (or as `capReached`); `split` arrives joined. */
  hitToCap: (cap: string, split?: string): string =>
    `Hit cap: ${cap}${split === undefined ? '' : ` · ${split}`}`,
  /** Several caps, each with what it is for (`capLabelled`, joined): "Hit cap: 9% (special
   *  attacks), 28% (white swings)", "Hit cap: reached (weapon attacks), 16% (spells)". */
  hitToCapEach: (caps: string, split?: string): string =>
    `Hit cap: ${caps}${split === undefined ? '' : ` · ${split}`}`,
  /** One of those caps: "16% (spells)". */
  capLabelled: (cap: string, label: string): string => `${cap} (${label})`,
  capLabelSpecials: 'special attacks',
  capLabelWhite: 'white swings',
  capLabelWeapon: 'weapon attacks',
  capLabelSpells: 'spells',
  /** The Stat priority note under the table: "Expertise cap: 6.5% (dodge)", parry added for a
   *  tank, "reached" for a cap the character is already past. */
  expertiseCap: (dodge: string, parry: string | undefined, split?: string): string =>
    `Expertise cap: ${dodge} (dodge)${parry === undefined ? '' : `, ${parry} (parry)`}${split === undefined ? '' : ` · ${split}`}`,
  capReached: 'reached',
  /** What the band's best-in-slot set and buffs already give toward a cap (`baseline` less its
   *  talents, which `capFrom` names one by one): "10.2% hit from BiS gear and buffs". Not "built
   *  in": two specs with the same talents differ here because their sets differ. `sources` is
   *  what the ranker counts besides talents (`sim/cmd/leveling-bis/hitcap.go`, `expertisecap.go`). */
  capOnSet: (percent: string, stat: string, sources: string): string =>
    `${percent}% ${stat} from BiS ${sources}`,
  /** The hit that set is still short of the cap. */
  capStillToFind: (percent: string): string => `${percent}% still to find`,
  /** Past the weapon cap, what a melee spec's spells are still short of. */
  capStillToFindForSpells: (percent: string): string => `${percent}% still to find for spells`,
  capFrom: (percent: string, source: string): string => `${percent}% from ${source}`,
  /** A school talent's share, which only that school's spells get: "5% from Shadow Focus on
   *  Shadow spells". */
  capFromOnSchools: (percent: string, source: string, schools: string): string =>
    `${percent}% from ${source} on ${schools} spells`,
  notSignificantNote: (labels: string, plural: boolean): string =>
    `${labels} ${plural ? 'show' : 'shows'} as "Not significant": no measurable value at level 60.`,
  /** A stat whose weight is below zero beyond its error: no number is printed for it. */
  noGainNote: (labels: string, plural: boolean, role: GuideRole): string =>
    `${labels} ${plural ? 'show' : 'shows'} as "No gain": more of it does not raise your ${ROLE_OUTPUT[role]} at level 60.`,
  weightColumn: (anchor: string): string => `Weight (${anchor} = 1.00)`,
  /** "DPS per point", "HPS per point", "Tank score per point". */
  perPointColumn: {
    dps: 'DPS per point',
    healer: 'HPS per point',
    tank: 'Tank score per point',
  } as const satisfies Record<GuideRole, string>,
  unitLine: (stat: string, figure: string, role: GuideRole): string =>
    `1 point of ${stat} = ${figure} ${ROLE_RATE_UNIT[role]} at level 60`,
  /** A haste or expertise weight, which is per 1% rather than per point: "2.5 per 1% haste". */
  perPercentValue: (figure: string, stat: 'haste' | 'expertise'): string => `${figure} per 1% ${stat}`,
  notSignificant: 'Not significant',
  noGain: 'No gain',
  /** Haste no item at level 60 carries: worth something, but not a stat to gear for. */
  notOnItems: 'Not on any item',
  /** The note under the table for that row: "Spell haste: 8.1 per 1%, but no item at level 60
   *  has it." (the BiS rail's own haste caption, `bisCopy.weightsHasteCaption`, says the same). */
  notOnItemsNote: (label: string, figure: string | undefined): string =>
    figure === undefined
      ? `${label}: no item at level 60 has it.`
      : `${label}: ${figure} per 1%, but no item at level 60 has it.`,
  levelRangeLabel: (band: number): string => (band === 60 ? 'Level 60' : `Level ${band} to ${band + 9}`),
  // --- the leveling tree (design/specs/2026-11-04-leveling-tree.md) ---------------------------
  talentsByLevel: 'Talents by level',
  levelingTreeIntro:
    'One talent point a level, from 10 to 60. Move the level to see the build at that point.',
  /** "Level 23: Improved Gouge, rank 2"; a level the ladder places no point at reads its level alone. */
  levelingPoint: (level: number, name: string | null, rank: number): string =>
    name === null ? `Level ${level}` : `Level ${level}: ${name}, rank ${rank}`,
  /** "14 points spent, 37 to go", "1 point spent, 50 to go", "51 points spent, all placed". */
  levelingPointsSpent: (spent: number, toGo: number): string =>
    `${spent} ${spent === 1 ? 'point' : 'points'} spent, ${toGo === 0 ? 'all placed' : `${toGo} to go`}`,
  previousLevel: 'Previous level',
  nextLevel: 'Next level',
  levelSliderLabel: 'Level',
  levelValueText: (level: number): string => `Level ${level}`,
  loadLevelBuildAriaLabel: (level: number): string => `Load the level ${level} build in the planner`,
  tocBarLabel: 'On this page',
  tocBarHint: 'Jump to a section',
  alliancePicksHeading: 'Alliance',
  hordePicksHeading: 'Horde',
  factionGroup: 'Faction',
  alsoUsed: 'Food, oil and potions',
  recommendedRacesLine: (name: string, faction: string, article: string, racials: string): string =>
    `${name} is ${article} ${faction} pick: ${racials}.`,
} as const;
