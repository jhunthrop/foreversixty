// "Latest on Forever" (lib/home/feed.ts, home B spec 2026-11-07 P8, P9, F1).
import { describe, expect, it } from 'vitest';
import { agoText, bandOf, dayText, feedPath, homeRows, whenText, withFloor, type FeedRow } from './feed';

function row(id: string, over: Partial<FeedRow> = {}): FeedRow {
  return {
    id,
    kind: 'log_upload',
    at: '2026-11-06T17:00:00Z',
    tag: 'logs',
    icon: 'x',
    crest: null,
    quality: null,
    href: '/',
    date_only: false,
    segments: [{ text: id }],
    ...over,
  };
}

const NOW = new Date('2026-11-06T18:00:00Z');

describe('feed', () => {
  it('words a time ago, and a dated event by its day', () => {
    expect(agoText('2026-11-06T17:59:40Z', NOW)).toBe('just now');
    expect(agoText('2026-11-06T17:56:00Z', NOW)).toBe('4 min ago');
    expect(agoText('2026-11-06T16:00:00Z', NOW)).toBe('2 h ago');
    expect(agoText('2026-11-03T18:00:00Z', NOW)).toBe('3 d ago');
    expect(dayText('2026-12-09T00:00:00Z')).toBe('Wed 9 Dec');
    expect(whenText({ at: '2026-12-09T00:00:00Z', date_only: true }, NOW)).toBe('Wed 9 Dec');
  });

  it('reads the band in tens from 20 to 60', () => {
    expect([bandOf(12), bandOf(20), bandOf(34), bandOf(59), bandOf(60), bandOf(undefined)]).toEqual([
      20, 20, 30, 50, 60, 20,
    ]);
  });

  it('asks for seven rows of a filter, spec and band', () => {
    expect(feedPath({ filter: 'spec', spec: 'warrior-fury', band: 20 })).toBe(
      '/v1/feed?filter=spec&limit=7&spec=warrior-fury&band=20',
    );
  });

  it('fills a quiet filter to five from the floor, and only with the spec on Your spec', () => {
    const floor = [row('raid'), row('fury', { spec: 'warrior-fury', tag: 'guides' }), row('a'), row('b')];
    expect(withFloor([row('1'), row('2')], floor, 'all', undefined).map((r) => r.id)).toEqual([
      '1',
      '2',
      'raid',
      'fury',
      'a',
    ]);
    const spec = withFloor([], floor, 'spec', 'warrior-fury');
    expect(spec.map((r) => [r.id, r.tag])).toEqual([['fury', 'your_spec']]);
  });

  it('never floors Your guild or Logs, and never past seven', () => {
    const floor = [row('a')];
    expect(withFloor([], floor, 'guild', undefined)).toEqual([]);
    expect(withFloor([], floor, 'logs', undefined)).toEqual([]);
    const many = Array.from({ length: 9 }, (_, i) => row(String(i)));
    expect(withFloor(many, floor, 'all', undefined)).toHaveLength(7);
  });

  it('reads a guild sim as the guild home reads it, and leaves the raid post to the plaque', () => {
    const sim = row('sim', {
      kind: 'guild_sim',
      segments: [
        { text: 'Thalrin', style: 'name', class: 'warrior' },
        { text: ' simmed 987.1 DPS, #2 of 14 DPS in Olympus XXVII (Fury Warrior, 60 raid)' },
      ],
    });
    const post = row('post', { kind: 'guild_raid' });
    const shown = homeRows([sim, post, row('log')], true);
    expect(shown.map((r) => r.id)).toEqual(['sim', 'log']);
    expect(shown[0].segments[1].text).toBe(' simmed 987.1 DPS, #2 of 14');
    expect(homeRows([post], false).map((r) => r.id)).toEqual(['post']);
  });

  it('reads newest first, whatever order the rows came in', () => {
    const rows = [
      row('an hour', { at: '2026-11-06T17:00:00Z' }),
      row('just now', { at: '2026-11-06T17:59:00Z' }),
      row('ten', { at: '2026-11-06T17:50:00Z' }),
    ];
    expect(homeRows(rows, false).map((r) => r.id)).toEqual(['just now', 'ten', 'an hour']);
  });
});
