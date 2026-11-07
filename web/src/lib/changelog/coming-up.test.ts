// web/src/lib/changelog/coming-up.test.ts
import { describe, expect, it } from 'vitest';
import { changelogSchema } from '../../content.config';
import { COMING_UP, COMING_UP_IDS, COMING_UP_SOURCE, comingUpHtml, comingUpRows } from './coming-up';

const at = (iso: string): Date => new Date(iso);
const ids = (rows: readonly { id: string }[]): string[] => rows.map((row) => row.id);

describe('Coming up', () => {
  it('draws the rows as the ruled table, escaped', () => {
    const html = comingUpHtml(COMING_UP);
    expect(html.startsWith('<table><thead><tr><th>Date</th><th>Content</th></tr></thead><tbody>')).toBe(true);
    expect(html.match(/<tr><td>/g)).toHaveLength(COMING_UP.length);
    expect(html).toContain("<td>Dec 9</td><td>Barrow Deeps, Hyjal Summit, and Onyxia's Lair open</td>");
    expect(comingUpHtml([{ id: 'x', when: '<b>', what: 'a & b' }])).toContain(
      '<td>&lt;b&gt;</td><td>a &amp; b</td>',
    );
  });

  it("names Blizzard's own recap as its source", () => {
    expect(COMING_UP_SOURCE.kind).toBe('blizzard');
    expect(COMING_UP_SOURCE.url).toMatch(/^https:\/\/news\.blizzard\.com\//);
  });

  it('keeps every row while its day is ahead', () => {
    expect(ids(comingUpRows(at('2026-11-07T12:00:00Z')))).toEqual(COMING_UP_IDS);
  });

  it('keeps a dated row through its own UTC day and drops it the day after', () => {
    expect(ids(comingUpRows(at('2026-12-09T23:59:59Z')))).toContain('raids-dec-9');
    const after = ids(comingUpRows(at('2026-12-10T00:00:00Z')));
    expect(after).not.toContain('raids-dec-9');
    expect(after).toEqual(['hardcore', 'spring-2027', 'summer-2027']);
  });

  it('keeps an undated season however late the clock', () => {
    expect(ids(comingUpRows(at('2030-01-01T00:00:00Z')))).toEqual(['hardcore', 'spring-2027', 'summer-2027']);
  });

  it('drops a row a changelog entry covers, dated or not', () => {
    const now = at('2026-11-07T00:00:00Z');
    expect(ids(comingUpRows(now, ['raids-dec-9', 'hardcore']))).toEqual(['spring-2027', 'summer-2027']);
  });

  it('works on any rows passed in', () => {
    const rows = [
      { id: 'a', on: '2026-01-02', when: 'Jan 2', what: 'a' },
      { id: 'b', when: 'Later', what: 'b' },
    ];
    expect(ids(comingUpRows(at('2026-01-03T00:00:00Z'), [], rows))).toEqual(['b']);
  });

  it('gives every row a unique id, and a dated row an ISO day', () => {
    expect(new Set(COMING_UP_IDS).size).toBe(COMING_UP.length);
    for (const row of COMING_UP) {
      if ('on' in row) expect(row.on).toMatch(/^\d{4}-\d{2}-\d{2}$/);
    }
  });

  it('lets a changelog entry cover only a known row', () => {
    const base = {
      title: 't',
      date: '2026-12-09',
      kind: 'blizzard',
      updated: '2026-12-09',
      sources: [{ label: 'l', url: 'https://example.com', kind: 'blizzard' }],
    };
    expect(changelogSchema.safeParse({ ...base, covers: 'raids-dec-9' }).success).toBe(true);
    expect(changelogSchema.safeParse(base).success).toBe(true);
    expect(changelogSchema.safeParse({ ...base, covers: 'nope' }).success).toBe(false);
  });
});
