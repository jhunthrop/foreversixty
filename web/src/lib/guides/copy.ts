// Every user-visible string the guide pages (/guides, the class landings, the spec guides) and
// their components render, reference voice throughout (design/DESIGN-SYSTEM.md: state the
// thing and stop).

/** The unit a role's per-point and set figures are in: a tank's per-point figures are
 *  tank-score points (mitigation, risk and threat combined), never damage. */
const ROLE_RATE_UNIT = { dps: 'DPS', healer: 'HPS', tank: 'tank score' } as const;
/** The same units as a column header leads with them. */
const ROLE_RATE_LABEL = { dps: 'DPS', healer: 'HPS', tank: 'Tank score' } as const;

/** What a stat that adds nothing would have raised, by role. */
const ROLE_OUTPUT = { dps: 'damage', healer: 'healing', tank: 'survival or threat' } as const;

type GuideRole = keyof typeof ROLE_RATE_UNIT;

export const guidesCopy = {
  loadThisBuild: 'Load this build',
  simThisBuild: 'Sim this build',
  // --- rebuild (2026-10-04) --------------------------------------------------------------
  /** `/guides` index's own signed-in callout (spec §4.A/§6): `{specName}` and `{className}`
   *  come from the visitor's own character, e.g. "Your guide: Fury Warrior ›". */
  signedInCallout: (specName: string, className: string): string => `Your guide: ${specName} ${className} ›`,
  readTheGuide: 'Read the guide',
  yourSpecPill: 'Your spec',
  /** The class landing's spec plate's accessible name. */
  classLandingSpecsHeading: 'Specs',
  buildRailLabel: 'Build',
  /** The rail's Rotation card title; its band label is the head's meta (site alignment §5.4). */
  rotationRailTitle: 'Rotation',
  fullPriorityLink: 'Full priority ↓',
  statPriorityRailLabel: 'Stat priority',
  loadInPlanner: 'Load in planner',
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
  /** Said once above the lists when a row carries that tag. It never says "level 60 build":
   *  a raider is level 60 too and read the old wording as his own build (round 22 mage
   *  wow-player SHOULD). */
  raidBuildLegend:
    'Raid build only: needs a talent only the raid build takes. Not on the raid build? Skip that row.',
  /** "Picks for a Human (Alliance).", "Picks for an Undead (Horde)." */
  bisCaption: (race: string, faction: string): string =>
    `Picks for ${/^[AEIOU]/i.test(race) ? 'an' : 'a'} ${race} (${faction}).`,
  fullListLink: 'Full list and alternatives',
  /** Before the BiS lane's band links, each the band's own BiS page (lib/bis/band-path.ts). */
  bisByLevelLabel: 'Best in slot by level:',
  /** The BiS lane's link to the spec's role tier list. */
  tierListLink: 'Where it ranks at 60',
  preRaidNote:
    'These are the best picks before the first raid tier. Raid loot is not itemized in the client yet, and this list updates when it is.',
  notSimChecked: 'Ranked by stats only',
  enchants: 'Enchants',
  consumables: 'Consumables',
  /** What the enchant table is: each slot's best flat-stat enchant by the band's weights
   *  (gear-extras.ts); crit, hit and haste enchants are not scored. wow-player SHIP 2026-11-04. */
  enchantsNote:
    'The enchant for each slot that adds the most of your weighted stats. Crit, hit and haste enchants are not included.',
  consumablesNote: 'What to bring to a raid.',
  racialsColumn: 'Racials',

  recommendedTag: 'Recommended',
  /** The Stat priority headline's cap: "Hit cap: 16%". Caps arrive formatted with their own
   *  "%" (or as `capReached`); the breakdown joins through `capWithDetail`. */
  hitToCap: (cap: string): string => `Hit cap: ${cap}`,
  /** Several caps, each with what it is for (`capLabelled`, joined): "Hit cap: 9% (special
   *  attacks), 28% (white swings)", "Hit cap: reached (weapon attacks), 16% (spells)". */
  hitToCapEach: (caps: string): string => `Hit cap: ${caps}`,
  /** A cap line and its breakdown as one string: "Hit cap: 16% · 1% hit from BiS gear and
   *  buffs, 6% from Precision, 9% still to find". The page sets the two apart. */
  capWithDetail: (cap: string, detail: string | undefined): string =>
    `${cap}${detail === undefined ? '' : ` · ${detail}`}`,
  /** One of those caps: "16% (spells)". */
  capLabelled: (cap: string, label: string): string => `${cap} (${label})`,
  capLabelSpecials: 'special attacks',
  capLabelWhite: 'white swings',
  capLabelWeapon: 'weapon attacks',
  capLabelSpells: 'spells',
  /** The Stat priority note under the table: "Expertise cap: 6.5% (dodge)", parry added for a
   *  tank, "reached" for a cap the character is already past. */
  expertiseCap: (dodge: string, parry: string | undefined): string =>
    `Expertise cap: ${dodge} (dodge)${parry === undefined ? '' : `, ${parry} (parry)`}`,
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
  /** "DPS per point", "HPS per point", "Tank score per point"; "DPS per point or rating" when
   *  the table prices a rating stat (hit, crit), whose point is a rating point. */
  perPointColumn: (role: GuideRole, rated: boolean): string =>
    `${ROLE_RATE_LABEL[role]} per point${rated ? ' or rating' : ''}`,
  /** "1 point of spell power = ...", or "1 hit rating = ..." for a rating stat. */
  unitLine: (stat: string, figure: string, role: GuideRole, rated = false): string =>
    rated
      ? `1 ${stat} rating = ${figure} ${ROLE_RATE_UNIT[role]} at level 60.`
      : `1 point of ${stat} = ${figure} ${ROLE_RATE_UNIT[role]} at level 60`,
  /** Under a table with rating stats (hit, crit): "Hit and Critical strike are per rating point: ...". */
  ratingNote: (labels: string, plural: boolean, conversions: string): string =>
    `${labels} ${plural ? 'are' : 'is'} per rating point: ${conversions}.`,
  /** One rating stat's conversion: "10 hit rating is 1% hit". */
  ratingConversion: (stat: string, factor: number): string => `${factor} ${stat} rating is 1% ${stat}`,
  /** A haste or expertise weight, which is per 1% rather than per point: "2.5 per 1% haste". */
  perPercentValue: (figure: string, stat: 'haste' | 'expertise'): string => `${figure} per 1% ${stat}`,
  notSignificant: 'Not significant',
  noGain: 'No gain',
  /** Haste no item at level 60 carries: worth something, but not a stat to gear for. */
  notOnItems: 'Not on any item',
  /** The note under the table for that row: "Spell haste: 8.1 per 1%, but no item at level 60
   *  has it." (the BiS rail's own haste caption, `bisCopy.weightsHasteCaption`, says the same). */
  notOnItemsNote: (label: string, figure: string | undefined, role: GuideRole = 'dps'): string =>
    figure === undefined
      ? `${label}: no item at level 60 has it.`
      : `${label}: ${figure} ${ROLE_RATE_UNIT[role]} per 1%, but no item at level 60 has it.`,
  levelRangeLabel: (band: number): string => (band === 60 ? 'Level 60' : `Level ${band} to ${band + 9}`),
  // --- the leveling tree (design/specs/2026-11-04-leveling-tree.md) ---------------------------
  talentsByLevel: 'Talents by level',
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
  alliancePicksHeading: 'Alliance',
  hordePicksHeading: 'Horde',
  alsoUsed: 'Food, oil and potions',
  // --- the flat system (design/specs/2026-11-08-guides.md, rulings 0A) -----------------------
  /** The signed-in answer's label, all three guide pages. */
  answerLabel: 'Your character',
  /** The answer's link on the guide of the character's own spec below 60 (to #leveling). */
  talentsAtLevel: (level: number): string => `Talents at level ${level}`,
  /** Beside it: the ladder's next point (Q5), "Next point at 35: Expose Prey". */
  nextPointAt: (level: number, name: string): string => `Next point at ${level}: ${name}`,
  noCharacter: 'No character on this account yet.',
  /** The guide's ControlBar: its accessible name and its one group's label. */
  barLabel: 'Guide choices',
  specGroupLabel: 'Spec',
  /** The At a glance plate's accessible name. */
  glance: 'At a glance',
  /** A healer's At a glance rotation column (MUST 2). */
  healingPriorityRailTitle: 'Healing priority',
  /** The At a glance filler line when the filler is below the first four (MUST 1). */
  glanceFiller: (spell: string): string => `Filler: ${spell}`,
  /** A row cast below its top rank on purpose: "Rank 7 (cheap rank)". */
  rankLabel: (rank: number, cheap = false): string => `Rank ${rank}${cheap ? ' (cheap rank)' : ''}`,
  /** The Talents and builds lane's label over the build rows. */
  builds: 'Builds',
  sources: 'Sources',
  statColumn: 'Stat',
  indexTitle: 'Guides',
  indexDescription:
    'Class and spec guides for World of Warcraft: Forever: talents, rotation, stat weights and best in slot for every spec.',
  recommendedRacesLine: (name: string, faction: string, article: string, racials: string): string =>
    `${name} is ${article} ${faction} pick: ${racials}.`,
} as const;
