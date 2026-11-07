// web/src/lib/changelog/entries.test.ts
import { describe, expect, it } from 'vitest';
import type { Source } from '../sources';
import {
  changelogMonths,
  LAUNCH_DAY,
  showCounts,
  showFrom,
  showOf,
  siteAction,
  sortedEntries,
  type ChangelogInput,
} from './entries';

const BLIZZARD: Source = { label: 'Blizzard, a post', url: 'https://news.blizzard.com/a', kind: 'blizzard' };
const WOWHEAD: Source = { label: 'Wowhead, a story', url: 'https://www.wowhead.com/a', kind: 'community' };
const GUILD: Source = { label: 'This site, Guild', url: 'https://foreversixty.gg/guild', kind: 'site' };

function input(
  id: string,
  date: string,
  kind: Source['kind'],
  sources: Source[],
  html = '',
  note?: string,
): ChangelogInput {
  return { id, data: { title: id, date: new Date(date), kind, sources, ...(note ? { note } : {}) }, html };
}

const INPUTS = [
  input('a-sept', '2026-09-12', 'blizzard', [BLIZZARD, WOWHEAD, BLIZZARD], '<p>a</p>', 'One outlet.'),
  input('b-nov', '2026-11-05', 'site', [GUILD], '<p>b</p>'),
  input('c-nov', '2026-11-05', 'site', [GUILD], '<p>c</p>'),
  input('d-launch', LAUNCH_DAY, 'blizzard', [BLIZZARD], '<p>d</p>'),
  input('e-sept', '2026-09-12', 'community', [WOWHEAD], '', 'Only a note.'),
];

describe('the changelog model', () => {
  it('sorts newest first and keeps a day in collection order', () => {
    expect(sortedEntries(INPUTS).map((entry) => entry.id)).toEqual([
      'b-nov',
      'c-nov',
      'd-launch',
      'a-sept',
      'e-sept',
    ]);
  });

  it('folds every entry before launch day into a title row (ruling CL1)', () => {
    const folded = Object.fromEntries(sortedEntries(INPUTS).map((entry) => [entry.id, entry.folded]));
    expect(folded).toEqual({
      'b-nov': false,
      'c-nov': false,
      'd-launch': false,
      'a-sept': true,
      'e-sept': true,
    });
  });

  it("lists each source once, tagged only when its kind differs from the entry's", () => {
    const entry = sortedEntries(INPUTS).find((e) => e.id === 'a-sept')!;
    expect(entry.sources).toEqual([
      { label: BLIZZARD.label, url: BLIZZARD.url },
      { label: WOWHEAD.label, url: WOWHEAD.url, tag: 'Community' },
    ]);
    expect(entry.caveat).toBe('One outlet.');
    expect(entry.kindLabel).toBe('Blizzard');
  });

  it('turns a source on this site into the action that opens its page (ruling Q5)', () => {
    const entry = sortedEntries(INPUTS).find((e) => e.id === 'b-nov')!;
    expect(entry.sources).toEqual([]);
    expect(entry.actions).toEqual([{ label: 'Open the guild page ›', href: '/guild' }]);
    expect(entry.show).toBe('site');
  });

  it('fails the build for a site page with no action label', () => {
    expect(() => siteAction('https://foreversixty.gg/nowhere')).toThrow(/\/nowhere/);
  });

  it('reads a note as the body of an entry with no body, never twice', () => {
    const entry = sortedEntries(INPUTS).find((e) => e.id === 'e-sept')!;
    expect(entry.noteAsBody).toBe('Only a note.');
    expect(entry.caveat).toBeUndefined();
  });

  it('groups by month with the count each filter leaves', () => {
    const months = changelogMonths(sortedEntries(INPUTS));
    expect(months.map((m) => [m.title, m.anchor, m.counts])).toEqual([
      ['November 2026', '2026-11', { all: 3, game: 1, site: 2 }],
      ['September 2026', '2026-09', { all: 2, game: 2, site: 0 }],
    ]);
  });

  it('splits Game (Blizzard, community, datamined) from Site, and reads ?show=', () => {
    expect(['blizzard', 'community', 'datamined', 'site'].map((k) => showOf(k as Source['kind']))).toEqual([
      'game',
      'game',
      'game',
      'site',
    ]);
    expect(showCounts(sortedEntries(INPUTS))).toEqual({ all: 5, game: 3, site: 2 });
    expect(showFrom('?show=game')).toBe('game');
    expect(showFrom('?show=nope')).toBe('all');
    expect(showFrom('')).toBe('all');
  });
});
