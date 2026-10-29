// web/src/lib/home-landing-copy.ts
// The home page's own words (spec 2026-09-23, "the landing page is the product, not the
// wiki"; spec 2026-09-25 §3.5 drops the Reference band; spec 2026-09-28's own guild card in
// the hero's right column was, in turn, replaced by a Switch character panel (home rebuild
// spec 2026-09-30 §3.B.4, itself removed by the signed-in panel spec 2026-10-10: the header
// selector is the one way to change character), HomeGuildCard.svelte and lib/guild/home-card.ts removed as
// dead code; home rebuild spec §1 replaces the headline with tenet 14's own fixed sentence):
// the sky-band hero, the four product panels, and the "Get set up" card. `home-panel-copy.ts`
// stays the account-aware strip's own copy (the signed-out sentence there is now also the
// hero's sentence, so it stays the one source rather than a second copy of the same words
// here).

const HEADLINE_LEAD = 'Play your class';
const HEADLINE_ACCENT = 'better.';

export const homeHeroCopy = {
  eyebrow: 'World of Warcraft: Forever',
  // Tenet 14 / design/DESIGN-SYSTEM.md principle 1's one fixed sentence -- not a pitch, the
  // one thing this page is allowed to say about itself. Verbatim, capital P only, full stop
  // included (home rebuild spec §1); this is also the hero H1's exact glyph set that
  // scripts/hero-font.py subsets into hero-font.css, so a change here needs that script
  // re-run (`python3 scripts/hero-font.py` from web/).
  headline: `${HEADLINE_LEAD} ${HEADLINE_ACCENT}`,
  // The mock's two-tone split (polish spec 4.4): the lead in --strong, the accent word in gold.
  headlineLead: HEADLINE_LEAD,
  headlineAccent: HEADLINE_ACCENT,
  sentence:
    'Pick your class for its best-in-slot list, talents and rotation, or sign in and your own characters arrive with their gear, talents and guild.',
  addonAltLabel: 'or paste an addon export',
  addonAltHref: '/setup#paste',
} as const;

/** Home rebuild spec §3.A.4/§3.B.5: the two nine-class rows below the hero -- "Best in slot
 *  by class" signed out, "Another class" signed in, same aside link either way. */
export const homeClassPickerCopy = {
  bestInSlotByClassHeading: 'Best in slot by class',
  anotherClassHeading: 'Another class',
  allSpecsLink: 'All 28 specs',
  allSpecsHref: '/bis',
  tierListLink: 'Tier list',
  tierListHref: '/tiers',
} as const;

/** Home rebuild spec §3.A.2: the signed-out hero's right-column example panel -- a
 *  leveling BiS preview, captioned unambiguously as a stand-in class, never a claim about
 *  the visitor's own character. */
export const homeHeroStatePanelCopy = {
  /** While the beta runs (W.9: the card follows the clock, lib/home/hero-card.ts). */
  betaEyebrow: 'Beta is live',
  levelCapLabel: 'level cap 30',
  betaBandsLine: 'Leveling best in slot, levels 10 to 30',
  /** Between the beta and launch: "Launch Nov 4 · 15:00 PST", and the days left. */
  launchEyebrow: (key: string, note: string | undefined): string =>
    note === undefined || note === '' ? `Launch ${key}` : `Launch ${key} · ${note}`,
  /** The three countdown tiles' micro labels. */
  countdownLabels: { days: 'Days', hours: 'Hours', minutes: 'Minutes' },
  /** Once launch day is over. */
  launchedEyebrow: (key: string): string => `Launched ${key}`,
  launchedTag: 'level cap 60',
  /** After launch the tiles count to the first raids. */
  firstRaidsTag: (key: string): string => `First raids ${key}`,
  /** The card's link row: the real Hunter list, never a claim about the visitor's own character. */
  linkTitle: 'Marksmanship Hunter best in slot',
  linkHref: '/bis/hunter/marksmanship',
  bandsLine: 'Leveling best in slot, levels 20, 30, 40, 50, 60',
} as const;

/** Spec section 2.4, "Around the site": the two compact tables, each five rows "as today"
 *  (RecentReports compact and HomeTopGuilds), with a "See all" link apiece. */
export const homeAroundTheSiteCopy = {
  reportsLabel: 'Recent reports',
  guildsLabel: 'Top guilds',
  seeAll: 'See all',
} as const;

export interface HomeCompanionRowCard {
  title: string;
  description: string;
  href: string;
}

/** Spec 2026-09-25 §3.5: the addon and companion row becomes one card. Signed-out only --
 *  see `homeGetSetUpCopy` for the state-aware, signed-in line (review round 1 fix item 3). */
export const homeCompanionRow: readonly HomeCompanionRowCard[] = [
  {
    title: 'Get set up',
    description: 'Sign in, install the addon, and run the companion — one page.',
    href: '/setup',
  },
] as const;

/** Review round 1 fix item 3: a signed-in visitor never sees the three-step pitch above --
 *  `lib/home/get-set-up.ts` builds one slim line from these words, naming only the step(s)
 *  `me` proves are not done yet. The companion has no signal here (no API tells this page
 *  whether one is paired), so its own item never carries a done mark either way. */
export const homeGetSetUpCopy = {
  /** Both known steps (signed in, addon linked) done: the one line names them and points at
   *  the one step left. Home rebuild spec §3.B.6: takes the addon's own sync time (the same
   *  relative figure the hero's sync line shows, `build.captured_at`) so the two never
   *  disagree about when the addon last spoke -- a function, not a fixed string, since the
   *  hero and this card read the same fact from two different characters' data in the
   *  general case. */
  bothDone: (syncedAgo: string): string =>
    `Signed in and addon linked · synced ${syncedAgo} → Set up the companion`,
  addonRemaining: 'Install the addon',
  companionRemaining: 'Set up the companion',
} as const;

/** Task L8.2 ruling 7: the character search box. Typing `Name`, or `Name-Realm` as the game
 *  writes it, opens that character's profile (lib/character-search.ts). */
export const homeCharacterSearchCopy = {
  heading: 'Look up a character',
  /** The one muted line beside the heading. */
  aside: 'See any character’s Forever Score.',
  nameLabel: 'Character name',
  namePlaceholder: 'Character name',
  hint: 'Type the name, then pick the region and ruleset.',
  regionLabel: 'Region',
  rulesetLabel: 'Ruleset',
  submit: 'Look up',
  empty: 'Type a character’s name.',
  invalid: 'No character can have that name.',
  /** The keyboard hint inside the box (design system: Search), and what it says aloud. */
  slashKey: '/',
  slashHint: 'Press / to search',
  /** The typeahead list's accessible name (both lookup boxes). */
  typeaheadLabel: 'Characters',
  /** The one muted row when no character matches what was typed. */
  typeaheadEmpty: (typed: string): string => `No character named "${typed}" yet`,
} as const;
