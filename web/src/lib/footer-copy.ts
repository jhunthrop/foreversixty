// web/src/lib/footer-copy.ts
// The footer's copy and link set, one source for Footer.astro and its tests: round 5's compact
// footer (design/specs/2026-11-09-directions.md, coordinator ruling 2026-11-09): the seal and name,
// Get set up, Premium, the Discord and Report a bug, the legal links, the not-affiliated sentence and
// the operator line. FOOTER_COLUMNS is no longer drawn; it is the site's page list the 404 page
// suggests from (lib/not-found/suggest.ts).
import communityLinks from '../data/links.json';

export interface FooterLink {
  readonly label: string;
  readonly href: string;
}

export interface FooterColumn {
  readonly heading: string;
  readonly links: readonly FooterLink[];
}

/** The footer's site links, beside the seal: Get set up left the bar for here and the account menu. */
export const FOOTER_LINKS: readonly FooterLink[] = [
  { label: 'Get set up', href: '/setup' },
  { label: 'Premium', href: '/premium' },
];

export const FOOTER_DISCORD_BUTTON: FooterLink = {
  label: 'Join the Discord',
  href: communityLinks.discordInvite,
};

// Bugs are reported in the Discord's #bug-reports forum; the invite lands a newcomer there.
export const FOOTER_REPORT_BUG_BUTTON: FooterLink = {
  label: 'Report a bug',
  href: communityLinks.discordInvite,
};

export const FOOTER_COLUMNS: readonly FooterColumn[] = [
  {
    heading: 'Play',
    links: [
      { label: 'Get set up', href: '/setup' },
      { label: 'Guides', href: '/guides' },
      { label: 'Best in slot', href: '/bis' },
      { label: 'Talent planner', href: '/planner' },
      { label: 'Simulator', href: '/sim' },
      { label: 'Tier list', href: '/tiers' },
    ],
  },
  {
    heading: 'Logs',
    links: [
      { label: 'Upload a log', href: '/logs' },
      { label: 'Your reports', href: '/reports' },
      { label: 'Rankings', href: '/rankings' },
    ],
  },
  {
    heading: 'Guilds',
    links: [
      { label: 'Find a guild', href: '/guild' },
      { label: 'Claim your guild', href: '/guild#claim' },
      { label: 'Guild plans', href: '/premium#guild-leaders' },
    ],
  },
  {
    heading: 'About',
    links: [
      { label: 'Premium', href: '/premium' },
      { label: 'About', href: '/about' },
      { label: 'Changelog', href: '/changelog' },
      { label: 'Data sources', href: '/sources' },
    ],
  },
];

export const FOOTER_LEGAL_LINKS: readonly FooterLink[] = [
  { label: 'Terms', href: '/terms' },
  { label: 'Privacy', href: '/privacy' },
  { label: 'Refunds', href: '/refunds' },
];

export const FOOTER_TRADEMARK =
  'Not affiliated with or endorsed by Blizzard Entertainment. World of Warcraft and Warcraft are trademarks of Blizzard Entertainment, Inc.';

export function footerOperatorLine(year: number): string {
  return `Forever Sixty is a registered trade name of Commish LLC. © ${year} Commish LLC.`;
}
