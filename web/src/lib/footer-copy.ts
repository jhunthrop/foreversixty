// web/src/lib/footer-copy.ts
// The footer's copy and link set, one source for Footer.astro and its tests: round 5's one-row
// footer (the mock, design/mocks/directions/png/g-home-1440.png; coordinator ruling 2026-11-09):
// the seal and name, the seven links and the fan-made line. FOOTER_LEGAL_LINKS are the legal pages'
// own tabs (layouts/Content.astro). FOOTER_COLUMNS is not drawn; it is the site's page list the 404
// page suggests from (lib/not-found/suggest.ts).
import communityLinks from '../data/links.json';
import { fillShared } from './copy/shared';
import { sharedFooterFanLine, sharedFooterOperator } from './copy/shared/footer.generated';

export interface FooterLink {
  readonly label: string;
  readonly href: string;
}

export interface FooterColumn {
  readonly heading: string;
  readonly links: readonly FooterLink[];
}

// Bugs are reported in the Discord's #bug-reports forum; the invite lands a newcomer there.
/** The footer's one row of links, beside the seal (the round 5 mock, design/mocks/directions/png/
 *  g-home-1440.png): Get set up left the bar for here and the account menu. */
export const FOOTER_LINKS: readonly FooterLink[] = [
  { label: 'Get set up', href: '/setup' },
  { label: 'Addon', href: '/setup#addon' },
  { label: 'Discord', href: communityLinks.discordInvite },
  { label: 'Premium', href: '/premium' },
  { label: 'Report a bug', href: communityLinks.discordInvite },
  { label: 'Privacy', href: '/privacy' },
  { label: 'Terms', href: '/terms' },
];

/** The footer's not-official line, at the right of the row (owner rule: never mistakable for
 *  Blizzard's). The full trademark sentence and the operator line are on /about and /terms
 *  (lib/legal-notice.ts). */
export const FOOTER_FAN_LINE = sharedFooterFanLine;

/** The operator under the fan line (owner 2026-11-10): the legal entity and its trade name on every
 *  page. The year is the build's, so it never goes stale. Both lines are shared keys
 *  (apps/shared/copy/shared.json shared.footer.*), so every email's footer reads the same words. */
export const footerOperatorLine = (year: number): string => fillShared(sharedFooterOperator, { year });

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
    ],
  },
];

export const FOOTER_LEGAL_LINKS: readonly FooterLink[] = [
  { label: 'Terms', href: '/terms' },
  { label: 'Privacy', href: '/privacy' },
  { label: 'Refunds', href: '/refunds' },
];
