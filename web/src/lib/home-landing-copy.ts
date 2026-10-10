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

export const homeHeroCopy = {
  eyebrow: 'World of Warcraft: Forever',
  // Tenet 14 / design/DESIGN-SYSTEM.md principle 1's one fixed sentence -- not a pitch, the
  // one thing this page is allowed to say about itself. Verbatim, capital P only, full stop
  // included (home rebuild spec §1); this is also the hero H1's exact glyph set that
  // scripts/hero-font.py subsets into hero-font.css, so a change here needs that script
  // re-run (`python3 scripts/hero-font.py` from web/).
  headline: 'Play your class better.',
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
  betaEyebrow: 'Beta is live',
  levelCapLabel: 'level cap 30',
  examplePill: 'Example',
  bandsLine: 'Leveling best in slot, bands 20 to 60',
} as const;

export interface HomeProductPanelCopy {
  label: string;
  sentence: string;
  linkLabel: string;
  href: string;
}

/** Spec section 2 item 3, one entry per product panel, in display order. Review round 1 fix
 *  item 2 adds Guides last: the fifth door this site offers, previously reachable only from
 *  the header nav, now presented the same way as the other four. */
export const homeProductPanels: readonly HomeProductPanelCopy[] = [
  {
    label: 'Planner',
    sentence: 'Every talent tree for 1.60, every race and class, shared by link.',
    linkLabel: 'Open the planner',
    href: '/planner',
  },
  {
    label: 'Simulator',
    sentence: 'Your DPS, upgrades from any loot table, stat weights.',
    linkLabel: 'Open the simulator',
    href: '/sim',
  },
  {
    label: 'Logs',
    sentence: 'Upload a night or log live with the companion; every fight, every parse.',
    linkLabel: 'Open the logs',
    href: '/logs',
  },
  {
    label: 'Rankings',
    sentence: 'Guild progression and character parses, per boss.',
    linkLabel: 'Open the rankings',
    href: '/rankings',
  },
  {
    label: 'Guides',
    sentence: '28 spec guides, talent builds and rotations for every class.',
    linkLabel: 'Open guides',
    href: '/guides',
  },
] as const;

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
    description: 'Sign in, install the addon, and pair the companion -- one page.',
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
