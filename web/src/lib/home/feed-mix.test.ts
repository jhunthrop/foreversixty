// The signed-in feed (lib/home/feed-mix.ts, home body pass 2026-11-09).
import { describe, expect, it } from 'vitest';
import type { FeedRow, FeedSegment } from './feed';
import { FEED_SIMS_SHOWN, foldSims, signedInRows, simDeltaText, simFigureText, simLine } from './feed-mix';

const GUILD = 'Olympus XXVII';

function row(id: string, minutesAgo: number, over: Partial<FeedRow> = {}): FeedRow {
  return {
    id,
    kind: 'log_upload',
    at: new Date(Date.UTC(2026, 10, 6, 18) - minutesAgo * 60_000).toISOString(),
    tag: 'logs',
    icon: null,
    crest: null,
    quality: null,
    href: '/',
    date_only: false,
    segments: [{ text: id }],
    ...over,
  };
}

/** A guild sim in the API's own words (api/internal/activity guildSimEvent). */
function sim(
  id: string,
  minutesAgo: number,
  name: string,
  figure: string,
  place = '',
  guild = GUILD,
): FeedRow {
  return row(id, minutesAgo, {
    kind: 'guild_sim',
    tag: 'your_guild',
    crest: 'rogue',
    segments: [
      { text: name, style: 'name', class: 'rogue' },
      { text: ` simmed ${figure} DPS${place} in ` },
      { text: guild, style: 'strong' },
      { text: ' (Combat Rogue, 30 to 39)' },
    ],
  });
}

const words = (segments: readonly FeedSegment[]) => segments.map((s) => s.text).join('');

describe('feed mix', () => {
  it('reads a sim in the API words, and leaves any other row alone', () => {
    expect(simLine(sim('a', 1, 'Vessa', '71.2', ', #1 of 6 DPS'))).toMatchObject({
      value: 71.2,
      unit: 'DPS',
      place: ', #1 of 6',
      guild: GUILD,
      tail: ' (Combat Rogue, 30 to 39)',
    });
    expect(simLine(row('log', 1))).toBeNull();
  });

  it('writes figures and changes as the API writes sims', () => {
    expect([simFigureText(58.44), simFigureText(987.4)]).toEqual(['58.4', '987']);
    expect(simDeltaText(71.2, 68)).toBe('+3.2');
    expect(simDeltaText(950, 987)).toBe('−37');
    expect(simDeltaText(71.2, 71.21)).toBeNull();
  });

  it('folds a character to its latest sim with the change, and drops the own guild name', () => {
    const folded = foldSims([sim('old', 30, 'Vessa', '68.0'), sim('new', 3, 'Vessa', '71.2')], GUILD);
    expect(folded.map((r) => r.id)).toEqual(['new']);
    expect(words(folded[0].segments)).toBe('Vessa simmed 71.2 DPS +3.2 (Combat Rogue, 30 to 39)');
    expect(folded[0].segments.find((s) => s.text === '+3.2')?.style).toBe('gain');
  });

  it('drops a sim that measured nothing, and folds to the last one that did', () => {
    const folded = foldSims([sim('real', 30, 'Vessa', '68.0'), sim('nothing', 3, 'Vessa', '0.0')], GUILD);
    expect(folded.map((r) => r.id)).toEqual(['real']);
    expect(words(folded[0].segments)).toBe('Vessa simmed 68.0 DPS (Combat Rogue, 30 to 39)');
    expect(foldSims([sim('only', 3, 'Shoulder', '0.0')], GUILD)).toEqual([]);
  });

  it('keeps another guild by name, and the ladder place', () => {
    const [shown] = foldSims([sim('a', 3, 'Vessa', '71.2', ', #1 of 6 DPS', 'Other')], GUILD);
    expect(words(shown.segments)).toBe('Vessa simmed 71.2 DPS, #1 of 6 in Other (Combat Rogue, 30 to 39)');
  });

  it('shows at most three sims, so the other kinds keep their place', () => {
    const sims = ['A', 'B', 'C', 'D', 'E'].map((name, i) => sim(name, i + 1, name, '60.0'));
    const rows = signedInRows({
      rows: [...sims, row('log', 10)],
      floor: [],
      spec: undefined,
      ownGuild: GUILD,
      raidOnPlaque: false,
    });
    expect(rows.filter((r) => r.kind === 'guild_sim')).toHaveLength(FEED_SIMS_SHOWN);
    expect(rows.map((r) => r.id)).toContain('log');
  });

  it('mixes in the viewer spec site events, never the raid opening or a BiS change the API said', () => {
    const floor = [
      row('floor:raids', -40000, { kind: 'site', date_only: true }),
      row('floor:bis:hunter-beast-mastery', 50, {
        kind: 'site',
        tag: 'best_in_slot',
        spec: 'hunter-beast-mastery',
        href: '/bis/hunter/beast-mastery',
      }),
      row('floor:guide:hunter-beast-mastery', 20, {
        kind: 'site',
        tag: 'guides',
        spec: 'hunter-beast-mastery',
      }),
      row('floor:guide:warrior-fury', 5, { kind: 'site', tag: 'guides', spec: 'warrior-fury' }),
    ];
    const api = [
      sim('s', 1, 'Vessa', '71.2'),
      row('bis', 30, { kind: 'bis_change', href: '/bis/hunter/beast-mastery' }),
      ...['a', 'b', 'c'].map((id, i) => row(id, 40 + i)),
    ];
    const rows = signedInRows({
      rows: api,
      floor,
      spec: 'hunter-beast-mastery',
      ownGuild: GUILD,
      raidOnPlaque: false,
    });
    expect(rows.map((r) => r.id)).toEqual(['s', 'floor:guide:hunter-beast-mastery', 'bis', 'a', 'b', 'c']);
    expect(rows[1].tag).toBe('your_spec');
  });

  it('fills a quiet feed to five from the floor', () => {
    const floor = [
      row('f1', 5, { kind: 'site' }),
      row('f2', 6, { kind: 'site' }),
      row('f3', 7, { kind: 'site' }),
    ];
    const rows = signedInRows({
      rows: [row('a', 1)],
      floor,
      spec: undefined,
      ownGuild: undefined,
      raidOnPlaque: true,
    });
    expect(rows.map((r) => r.id)).toEqual(['a', 'f1', 'f2', 'f3']);
  });

  it('leaves the guild raid post to the plaque', () => {
    const post = row('post', 2, { kind: 'guild_raid' });
    const rows = signedInRows({
      rows: [post],
      floor: [],
      spec: undefined,
      ownGuild: GUILD,
      raidOnPlaque: true,
    });
    expect(rows).toEqual([]);
  });
});
