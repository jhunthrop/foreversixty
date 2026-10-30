// web/src/lib/rankings/copy.ts
// Every visible string the /rankings encounter picker uses, in the site's own honest-copy
// voice: never demands a choice the visitor cannot make.
export const encounterPickerCopy = {
  heading: 'Choose an encounter',
  failed: 'Encounters did not load.',
  noneYet: 'No encounters have been ranked yet.',
  tryGuildsProgress: 'The Guilds board ranks guild progress across every encounter and needs none.',
  openGuildsBoard: 'Open the Guilds board',
  openProgressBoard: 'Open the Progress board',
  noEncountersOption: 'No encounters yet',
} as const;

/** The home page's Rankings panel (spec 2026-09-23 §2 item 3): the top five guilds by
 *  progression, from `GET /v1/rankings/guilds?kind=progress` (needs no encounter). */
export const homeTopGuildsCopy = {
  /** The placeholder's screen-reader name, so the panel never says "Loading" in page text. */
  placeholderLabel: 'Top guilds',
  failed: 'Rankings did not load.',
  empty: 'No guilds ranked yet. Claim yours from the Guild menu.',
  /** The ghost button under the empty sentence (P.1). */
  emptyAction: 'Open the guild page',
  emptyHref: '/guild',
  /** "1 boss", "12 bosses": the count beside each guild. */
  bosses: (count: number): string => `${count} ${count === 1 ? 'boss' : 'bosses'}`,
} as const;

/** Spec 2026-09-25 §6: every ranking board's own empty state, character or guild. */
export const rankingsEmptyCopy = {
  message: 'No ranked fights yet for this filter. Rankings fill in as reports are uploaded.',
  action: 'Upload a log',
  href: '/logs',
} as const;

/**
 * The page's own words (design/specs/2026-10-11-rankings.md section 6). The shared character
 * header owns its identity and signed-out line; this is the band, the board and the header's
 * answer.
 */
export const rankingsCopy = {
  eyebrow: 'Characters',
  h1: 'Rankings',
  jobSignedIn: 'Where your character stands among characters like it.',
  jobSignedOut: 'How characters rank on each encounter, kill by kill.',
  tierQuestion: 'Comparing specs, not characters?',
  tierLink: 'Tier list',
  tierHref: '/tiers',
  updated: (date: string) => `Updated ${date}`,

  answerLabel: 'Your rank',
  action: 'See your parses →',
  answerFailed: 'Your rank did not load.',
  /** "of 212 Fury Warriors" once the API ranks each character's best kill; "kills" until then. */
  answerUnit: (total: number, cohort: string, deduped: boolean) =>
    deduped ? `of ${total} ${cohort}s` : `of ${total} ${cohort} kills`,
  /** `rulesetScope` reads "Normal ruleset" or "all rulesets". */
  answerNote: (encounter: string, rulesetScope: string) => `On ${encounter}, ${rulesetScope}`,
  answerUnranked: (encounter: string) =>
    `Not ranked yet on ${encounter}. Upload a log of a kill to get on the board.`,

  everyClassAndSpec: 'Every class and spec',
  anotherBoard: 'Another class or spec',
  showingNote: (shown: string, own: string) => `Showing ${shown}. Your rank is on the ${own} board.`,
  backTo: (own: string) => `Back to ${own}`,

  legendLead: 'Rank colour is the percentile:',
  legendUnderLast: 'under 25',

  ownBest: 'Your best kill',
  removed: 'Removed by a moderator',
  atRisk: 'Under review by a moderator',
  filtersButton: 'Filters',
  boardTabs: 'Board',
  pickerHeading: 'Choose an encounter',

  encounterLabel: 'Encounter',
  allRulesets: 'All rulesets',
  allRegions: 'All regions',
  allPhases: 'All phases',
  everyClass: 'Every class',
  everySpec: 'Every spec',
  bothFactions: 'Both',
  todayOnly: 'Today only',
  sizeMan: (size: number) => `${size}-man`,
} as const;
