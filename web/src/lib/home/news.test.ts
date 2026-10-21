import { describe, expect, it } from 'vitest';
import { isStaleSiteNews, latestNews, newsAgeDays } from './news';

const entry = (id: string, kind: string, date: string) => ({ id, data: { kind, date: new Date(date) } });
const NOW = new Date('2026-10-20T12:00:00Z');

describe('the landing LATEST block (W.9)', () => {
  it('counts whole days', () => {
    expect(newsAgeDays(new Date('2026-10-16'), NOW)).toBe(4);
    expect(newsAgeDays(new Date('2026-10-25'), NOW)).toBe(0);
  });

  it('treats a This site item as stale at a month old, and only that kind', () => {
    expect(isStaleSiteNews('site', new Date('2026-09-20'), NOW)).toBe(true);
    expect(isStaleSiteNews('site', new Date('2026-09-21'), NOW)).toBe(false);
    expect(isStaleSiteNews('blizzard', new Date('2026-09-14'), NOW)).toBe(false);
  });

  it('hides the This site row when nothing is newer than a month, and keeps the rest newest first', () => {
    const rows = latestNews(
      [
        entry('a', 'site', '2026-09-20'),
        entry('b', 'blizzard', '2026-10-16'),
        entry('c', 'datamined', '2026-09-14'),
        entry('d', 'blizzard', '2026-09-13'),
        entry('e', 'community', '2026-09-12'),
      ],
      NOW,
    );
    expect(rows.map((row) => row.id)).toEqual(['b', 'c', 'd', 'e']);
    expect(rows.some((row) => row.data.kind === 'site')).toBe(false);
  });

  it('shows a This site item from inside the month, and caps the list at four', () => {
    const rows = latestNews(
      [
        entry('new', 'site', '2026-10-19'),
        entry('b', 'blizzard', '2026-10-16'),
        entry('c', 'blizzard', '2026-10-15'),
        entry('d', 'blizzard', '2026-10-14'),
        entry('e', 'blizzard', '2026-10-13'),
      ],
      NOW,
    );
    expect(rows.map((row) => row.id)).toEqual(['new', 'b', 'c', 'd']);
  });
});
