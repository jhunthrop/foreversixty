// web/src/lib/home-landing-copy.ts
// The home page's own words (spec 2026-09-23, "the landing page is the product, not the
// wiki"; spec 2026-09-25 §3.5 drops the Reference band): the sky-band hero, the four product
// panels, the Your guild panel, and the "Get set up" card. `home-panel-copy.ts` stays the
// account-aware strip's own copy (the signed-out sentence there is now also the hero's
// sentence, so it stays the one source rather than a second copy of the same words here).
import { battlenetStartUrl } from './account/api';

export const homeHeroCopy = {
  eyebrow: 'World of Warcraft: Forever',
  headline: 'Your character, planned, simmed, logged and ranked.',
  addonAltLabel: 'or paste an addon export',
  addonAltHref: '/setup#paste',
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
    sentence: '27 spec guides, talent builds and rotations for every class.',
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

export const homeGuildPanelCopy = {
  label: 'Your guild',
  /** Signed out, or signed in with no guild yet: the same claim invitation either way. */
  claimSentence:
    'Claim your guild with a Battle.net sign-in: a home page, roster, progression and officer tools.',
  /** Ruling 3 (plan): no generic guild directory exists, so the claim CTA is the sign-in
   *  itself, the same helper the hero button uses. */
  claimHref: battlenetStartUrl('/account?signed_in=1'),
  claimLinkLabel: 'Sign in with Battle.net',
  failed: 'Your guild did not load.',
  progressionOf: (killed: number, total: number): string => `${killed} of ${total} bosses down`,
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
   *  the one step left, verbatim. */
  bothDone: 'Signed in and addon linked · Set up the companion →',
  addonRemaining: 'Install the addon',
  companionRemaining: 'Set up the companion',
} as const;
