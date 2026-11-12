// web/src/lib/golinks/config.ts
// The tracked short links (foreversixty.gg/go/<slug>) and the visit beacon: the constants the
// Worker's redirect, the base layout's arrival script and the owner's /admin/links page share.
// The slug rule is api/internal/golinks's (and migration 0066's check); change all three.
import { SITE_URL } from '../seo';

export const GO_PREFIX = '/go/';

/** A link's slug: lowercase letters, digits and dashes, 1 to 40, starting with a letter or digit. */
export const SLUG_PATTERN = '^[a-z0-9][a-z0-9-]{0,39}$';
const SLUG = new RegExp(SLUG_PATTERN);

export function isSlug(value: string): boolean {
  return SLUG.test(value);
}

/**
 * The query parameter the redirect hands the slug to the landing page in. Not `ref`: the
 * simulator already reads ?ref= as the character or fight it loads.
 */
export const LINK_PARAM = 'golink';

/**
 * Where the landing page keeps the slug until the first signed-in page sends it. localStorage
 * rather than sessionStorage so an emailed sign-in link, which opens in a new tab, still
 * carries it; ARRIVAL_TTL_MS bounds it, and it is removed once sent.
 */
export const ARRIVAL_KEY = 'fs.arrival';

/**
 * Where the inline capture puts the newest click. The first click stays in ARRIVAL_KEY until it
 * goes stale (later.ts promoteArrival), so an account keeps the first link that brought it.
 */
export const ARRIVAL_NEXT_KEY = 'fs.arrival.next';

/**
 * How long a kept click waits for the sign-in: a creator's link counts for 30 days before sign-up
 * (the API's golinks.ClickWindow and creators.LinkWindow). An older one is dropped unsent.
 */
export const ARRIVAL_TTL_MS = 30 * 24 * 60 * 60 * 1000;

export const ARRIVAL_PATH = '/v1/me/arrival';

/**
 * Cloudflare Web Analytics: cookieless, no personal data. The site token is public by
 * design (it only says which site a beacon is for). The beacon loads only on the production
 * host, so previews, the e2e suite and Lighthouse never fetch it.
 */
export const BEACON_SRC = 'https://static.cloudflareinsights.com/beacon.min.js';
export const BEACON_TOKEN = 'b6d2adf8cec64540805310e82ef35b3c';
export const BEACON_HOST = new URL(SITE_URL).hostname;

/**
 * Pages the beacon never loads on: /app/* carries the phone apps' one-time sign-in code in
 * its URL for a moment, and that page loads no third-party script at all.
 */
export const BEACON_SKIP_PREFIXES = ['/app/'] as const;
