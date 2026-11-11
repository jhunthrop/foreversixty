// web/src/lib/golinks/funnel.ts
// The /admin/links table's arithmetic: the short link itself, each step's share of the step
// before it, and the totals row.
import { SITE_URL } from '../seo';
import type { GoLinkFunnel } from './api';
import { GO_PREFIX } from './config';

import { EMPTY_FIGURE, grouped } from '../format/number';
export type FunnelFigures = Pick<GoLinkFunnel, 'clicks' | 'bot_clicks' | 'sign_ins' | 'linked' | 'premium'>;

/** The URL the owner pastes into a post. */
export function shortLink(slug: string): string {
  return `${SITE_URL}${GO_PREFIX}${slug}`;
}

/** A step as a whole percent of the one before it; an en dash when there is nothing before. */
export function shareOf(part: number, whole: number): string {
  if (whole <= 0) return EMPTY_FIGURE;
  return `${Math.round((part / whole) * 100)}%`;
}

/** A figure as the table shows it: 6,132. */
export function countLabel(value: number): string {
  return grouped(value);
}

export function funnelTotals(rows: readonly GoLinkFunnel[]): FunnelFigures {
  return rows.reduce<FunnelFigures>(
    (sum, r) => ({
      clicks: sum.clicks + r.clicks,
      bot_clicks: sum.bot_clicks + r.bot_clicks,
      sign_ins: sum.sign_ins + r.sign_ins,
      linked: sum.linked + r.linked,
      premium: sum.premium + r.premium,
    }),
    { clicks: 0, bot_clicks: 0, sign_ins: 0, linked: 0, premium: 0 },
  );
}

/** Busiest first; a tie keeps the slugs in order. A new array: the input is left as it was. */
export function sortedByClicks(rows: readonly GoLinkFunnel[]): GoLinkFunnel[] {
  return [...rows].sort((a, b) => b.clicks - a.clicks || a.slug.localeCompare(b.slug));
}
