// web/src/lib/home-landing-copy.ts
// The home page's own words (spec 2026-09-23, "the landing page is the product, not the
// wiki"): the sky-band hero, the four product panels, the reference band's sentence, the
// Your guild panel, and the addon/companion row. `home-panel-copy.ts` stays the account-
// aware strip's own copy (the signed-out sentence there is now also the hero's sentence, so
// it stays the one source rather than a second copy of the same words here).
import { battlenetStartUrl } from './account/api';

export const homeHeroCopy = {
  eyebrow: 'World of Warcraft: Forever',
  headline: 'Your character, planned, simmed, logged and ranked.',
  addonAltLabel: 'or paste an addon export',
  addonAltHref: '/addon#paste',
} as const;

export interface HomeProductPanelCopy {
  label: string;
  sentence: string;
  linkLabel: string;
  href: string;
}

/** Spec section 2 item 3, one entry per product panel, in display order. */
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
] as const;

/** Spec section 2.4, "Around the site": the two compact tables, each five rows "as today"
 *  (RecentReports compact and HomeTopGuilds), with a "See all" link apiece. */
export const homeAroundTheSiteCopy = {
  reportsLabel: 'Recent reports',
  guildsLabel: 'Top guilds',
  seeAll: 'See all',
} as const;

export const homeReferenceCopy = {
  sentence: 'Every fact dated and sourced.',
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

/** Spec section 2 item 6, two small cards, reference voice: what each does, one line. */
export const homeCompanionRow: readonly HomeCompanionRowCard[] = [
  {
    title: 'The addon',
    description:
      'The addon adds bags, bank and professions to your Battle.net build and shows the next talent point in game.',
    href: '/addon',
  },
  {
    title: 'The companion',
    description: 'The companion logs live from your desktop and sends your exports.',
    href: '/logs#companion',
  },
] as const;
