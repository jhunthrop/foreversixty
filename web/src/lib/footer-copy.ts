// web/src/lib/footer-copy.ts
// The footer's copy and link set, one source for Footer.astro and its tests. Every string is a
// row of design/specs/2026-10-28-polish.md section 2.6; an edit here edits that row too.
import communityLinks from '../data/links.json';

export interface FooterLink {
  readonly label: string;
  readonly href: string;
}

export interface FooterColumn {
  readonly heading: string;
  readonly links: readonly FooterLink[];
}

export const FOOTER_PROMISE =
  'Play your class better: gear, talents, sims and logs for World of Warcraft: Forever.';

export const FOOTER_DISCORD_BUTTON: FooterLink = {
  label: 'Join the Discord',
  href: communityLinks.discordInvite,
};

export const FOOTER_CONTRIBUTE_BUTTON: FooterLink = {
  label: 'Contribute',
  href: `${communityLinks.githubRepo}/issues`,
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
    ],
  },
  {
    heading: 'Logs',
    links: [
      { label: 'Upload a log', href: '/logs' },
      { label: 'Your reports', href: '/reports' },
      { label: 'Rankings', href: '/rankings' },
      { label: 'Tier list', href: '/tiers' },
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
    heading: 'Forever Sixty',
    links: [
      { label: 'About', href: '/about' },
      { label: 'Premium', href: '/premium' },
      { label: 'Changelog', href: '/changelog' },
      { label: 'Sources', href: '/sources' },
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
  return `Forever Sixty is operated by Commish LLC. © ${year} Commish LLC.`;
}
