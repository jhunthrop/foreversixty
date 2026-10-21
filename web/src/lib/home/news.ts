// The landing page's LATEST block. A "This site" item is the site speaking for itself, and one
// older than a month reads as a site that has gone quiet right before it launches, so the block
// hides those and lets the dated news from Blizzard and the community carry it. A "This site" item
// inside the month stays. (W.9; UX review r5, finding 1.)
export const SITE_NEWS_MAX_AGE_DAYS = 30;

const DAY_MS = 24 * 60 * 60 * 1000;

/** Whole days between a news date and `now`, never negative. */
export function newsAgeDays(date: Date, now: Date = new Date()): number {
  return Math.max(0, Math.floor((now.getTime() - date.getTime()) / DAY_MS));
}

/** A "site" item is stale once it is a month old (30 days or more); other kinds are history and stay. */
export function isStaleSiteNews(kind: string, date: Date, now: Date = new Date()): boolean {
  return kind === 'site' && newsAgeDays(date, now) >= SITE_NEWS_MAX_AGE_DAYS;
}

interface NewsLike {
  data: { kind: string; date: Date };
}

/** The newest `limit` items, with the stale "This site" ones left out. */
export function latestNews<T extends NewsLike>(
  entries: readonly T[],
  now: Date = new Date(),
  limit = 4,
): T[] {
  return [...entries]
    .sort((a, b) => b.data.date.getTime() - a.data.date.getTime())
    .filter((entry) => !isStaleSiteNews(entry.data.kind, entry.data.date, now))
    .slice(0, limit);
}
