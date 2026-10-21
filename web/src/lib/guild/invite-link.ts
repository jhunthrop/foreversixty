// web/src/lib/guild/invite-link.ts
// The invite link an officer hands out: the API answers a path ("/guild/invite/fsg_..."),
// which is not something to paste into Discord. This makes it the full URL on the site the
// officer is on (the live origin in production, the dev origin when walking the local stack),
// and falls back to the production origin when there is no window (server render, tests).

/** The production origin; astro.config.mjs's `site` is the same value. */
export const INVITE_FALLBACK_ORIGIN = 'https://foreversixty.gg';

export function inviteUrl(path: string, origin?: string): string {
  const base = origin ?? (typeof window !== 'undefined' ? window.location.origin : INVITE_FALLBACK_ORIGIN);
  return `${base.replace(/\/+$/, '')}${path.startsWith('/') ? path : `/${path}`}`;
}

const SHORT_MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

/** "20 Oct": a day, in UTC as lib/dates.ts reads its days (the claim's and the link's own stamps). */
export function formatMadeOn(iso: string): string {
  const d = new Date(iso);
  if (Number.isNaN(d.getTime())) return '';
  return `${d.getUTCDate()} ${SHORT_MONTHS[d.getUTCMonth()]}`;
}
