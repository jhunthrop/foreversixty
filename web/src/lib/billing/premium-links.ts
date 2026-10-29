// web/src/lib/billing/premium-links.ts
// Where each contextual Premium link lands on /premium (spec 2026-10-29-premium-entry sections 7
// and 8): never the page top, always the card or the feature line the player just met.
import type { PremiumAnchor } from './plans-page-copy';

const PLANS_PAGE = '/premium';

const at = (anchor: PremiumAnchor | 'premium' | 'guild-leaders'): string => `${PLANS_PAGE}#${anchor}`;

export const PREMIUM_LINKS = {
  /** The Premium card itself: the account panel's Plan row and the account page's See plans. */
  card: at('premium'),
  coach: at('coach'),
  serverSims: at('server-sims'),
  history: at('history'),
  notifications: at('notifications'),
  guildLeaders: at('guild-leaders'),
} as const;
