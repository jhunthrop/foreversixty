import { grouped } from '../format/number';
// web/src/lib/rankings/copy.ts
// Every visible string the /rankings encounter picker uses, in the site's own honest-copy
// voice: never demands a choice the visitor cannot make.
export const encounterPickerCopy = {
  heading: 'Encounters',
  noneYet: 'No boss kills are ranked yet.',
  /** What fills it: the line under the empty sentence. */
  fill: 'A boss gets its board when the first log of its kill is uploaded.',
  /** "Onyxia's Lair · 40-player"; either half left out when no log named it. */
  where: (zone: string | null, size: number | null): string =>
    [zone, size === null ? null : `${size}-player`]
      .filter((part) => part !== null && part !== '')
      .join(' · '),
  /** "212 ranked": the encounter's ranked kills. */
  ranked: (kills: number): string => `${grouped(kills)} ranked`,
  /** One line in place of an encounter's top parses, or of the spec leaders. */
  topFailed: 'Top parses did not load.',
  topNone: 'No ranked kill on this board yet.',
  leadersHeading: 'Leaders by spec',
  /** The leaders' two lists, by what each spec is ranked on. */
  damageHead: 'Damage',
  healingHead: 'Healing',
  /** Over the encounters past the busiest few. */
  more: 'More encounters',
} as const;

/** The home page's Rankings panel (spec 2026-09-23 §2 item 3): the top five guilds by
 *  progression, from `GET /v1/rankings/guilds?kind=progress` (needs no encounter). */
export const homeTopGuildsCopy = {
  /** The placeholder's screen-reader name, so the panel never says "Loading" in page text. */
  placeholderLabel: 'Top guilds',
  failed: 'Rankings did not load.',
  empty: 'No guild has a boss kill logged yet.',
  /** The ghost button under the empty sentence (P.1). */
  emptyAction: 'Open the guild page',
  emptyHref: '/guild',
  /** "1 boss", "12 bosses": the count beside each guild. */
  bosses: (count: number): string => `${count} ${count === 1 ? 'boss' : 'bosses'}`,
  /** The guild landing's figure: "1 boss down". */
  bossesDown: (count: number): string => `${count} ${count === 1 ? 'boss' : 'bosses'} down`,
  /** The guild landing's row line after where the guild lives: its newest kill. */
  lastKill: (day: string): string => `last kill ${day}`,
} as const;

/** Spec 2026-09-25 §6: every ranking board's own empty state, character or guild: what is
 *  missing, then what fills it, then the one action (and, on a narrowed board, the wider one). */
export const rankingsEmptyCopy = {
  /** "No Fury Warrior kills on Onyxia yet." */
  board: (cohort: string, encounter: string): string => `No ${cohort} kills on ${encounter} yet.`,
  boardFill: 'Upload a log with the kill and it ranks here.',
  guilds: (encounter: string): string => `No guild has a ranked kill on ${encounter} yet.`,
  guildsAnywhere: 'No guild has a ranked boss kill yet.',
  guildsFill: 'A guild ranks once a member uploads a log of the kill.',
  action: 'Upload a log',
  href: '/logs',
  widen: 'Every class and spec',
  /** Under an empty board: the bosses that do have ranked kills. */
  elsewhere: 'Bosses with ranked kills',
} as const;

/**
 * The page's own words (design/specs/2026-10-11-rankings.md section 6). The shared character
 * header owns its identity and signed-out line; this is the band, the board and the header's
 * answer.
 */
export const rankingsCopy = {
  eyebrow: 'Characters',
  h1: 'Rankings',
  /** The page's lead: what the boards hold and how a player gets on them. */
  lead: 'Logged boss kills on Forever, ranked by DPS and HPS for each spec, class, ruleset and guild. The companion uploads your combat log as you play, so your kills join the board.',
  tierLink: 'Tier list',
  tierHref: '/tiers',
  updated: (date: string) => `Updated ${date}`,

  answerLabel: 'Your rank',
  action: 'See your parses →',
  /** The board's answer link (the arrow is the link's own chevron). */
  parsesLink: 'See your parses',
  answerFailed: 'Your rank did not load.',
  /** "of 212 Fury Warriors" once the API ranks each character's best kill; "kills" until then. */
  answerUnit: (total: number, cohort: string, deduped: boolean) =>
    deduped ? `of ${total} ${cohort}s` : `of ${total} ${cohort} kills`,
  /** `rulesetScope` reads "Normal ruleset" or "all rulesets". */
  answerNote: (encounter: string, rulesetScope: string) => `On ${encounter}, ${rulesetScope}`,
  answerUnranked: (encounter: string) =>
    `Not ranked yet on ${encounter}. Upload a log of a kill to get on the board.`,

  everyClassAndSpec: 'Every class and spec',
  /** The board panel's action on a narrowed board: back to every class and spec. */
  widen: 'Every class and spec',
  showingNote: (shown: string, own: string) => `Showing ${shown}. Your rank is on the ${own} board.`,
  backTo: (own: string) => `Back to ${own}`,

  legendLead: 'Percentile',
  legendUnderLast: 'under 25',

  ownBest: 'Your best kill',
  removed: 'Removed by a moderator',
  atRisk: 'Under review by a moderator',
  filtersButton: 'Filters',
  clearFilters: 'Clear filters',
  boardTabs: 'Board',
  pickerHeading: 'Choose an encounter',

  encounterLabel: 'Encounter',
  /** The filter panel's other field labels (spec 2026-10-30-controls-charts §A10 row 14). */
  classLabel: 'Class',
  specLabel: 'Spec',
  rulesetLabel: 'Ruleset',
  regionLabel: 'Region',
  phaseLabel: 'Phase',
  factionLabel: 'Faction',
  metricLabel: 'Metric',
  boardLabel: 'Board',
  allRulesets: 'All rulesets',
  allRegions: 'All regions',
  allPhases: 'All phases',
  everyClass: 'Every class',
  /** The Guilds Progress board's encounter choice: every boss at once. */
  everyEncounter: 'Every encounter',
  everySpec: 'Every spec',
  bothFactions: 'Both',
  todayOnly: 'Today only',
  percentileLabel: 'Percentile',
  /** The value's hover line: it opens the kill in its report. */
  openReport: 'Open the kill in its report',
  /** A board row's link to the kill's log. */
  reportLink: 'Report',
  /** The guild board's panel titles, by kind. */
  guildTitle: { progress: 'Most bosses killed', speed: 'Fastest kills', execution: 'Fewest deaths' },
  /** An execution guild row's figure (a progress row's is homeTopGuildsCopy.bosses). */
  deaths: (count: number): string => `${count} ${count === 1 ? 'death' : 'deaths'}`,
} as const;
