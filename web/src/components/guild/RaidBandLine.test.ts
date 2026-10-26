// web/src/components/guild/RaidBandLine.test.ts
// The guild band's raid line (raids spec 3.4, 6.1 "Band line", "{when}", "Band line, none
// planned"), rendered with svelte/server; the patch after a press and the race lines are the
// component's module functions, tested on their own.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { GuildApiError, type GuildRaidView, type NextRaidView } from '../../lib/guild/api';
import { raidCopy } from '../../lib/guild/raid-copy';
import RaidBandLine, { nextRaidAfter, nextRaidAfterRefusal, pressFailedLine } from './RaidBandLine.svelte';

const TZ = 'America/Chicago';
// Monday 2026-12-07, noon in Chicago; the raid is Thursday 10 Dec, 8 pm Chicago.
const NOW = new Date('2026-12-07T18:00:00Z');
const STARTS = '2026-12-11T02:00:00Z';
const KEY = 'us/normal/kraggor';
const ALT = 'us/normal/sunderfel';

/** A `<button>` carrying the `disabled` attribute (not the `disabled:` class variant). */
const DISABLED_BUTTON = /<button[^>]*\sdisabled(?:=""|\s|>)/g;

type ViewerChar = GuildRaidView['viewer']['characters'][number];

function raid(overrides: Partial<NextRaidView> = {}): NextRaidView {
  return {
    id: 7,
    zone: 'onyxias-lair',
    zone_name: "Onyxia's Lair",
    title: "Onyxia's Lair",
    starts_at: STARTS,
    size: 40,
    guild_timezone: TZ,
    counts: { in: 20, late: 3, tanks: 3, healers: 7, dps: 13 },
    own: null,
    ...overrides,
  };
}

interface Opts {
  nextRaid?: NextRaidView | null;
  officer?: boolean;
  currentKey?: string | null;
  ownCharacters?: { character_key: string; name: string }[];
  viewerTimezone?: string;
}

function html(opts: Opts = {}): string {
  return render(RaidBandLine, {
    props: {
      guildId: 2,
      nextRaid: opts.nextRaid === undefined ? raid() : opts.nextRaid,
      officer: opts.officer ?? false,
      guildPath: '/guild/us/normal/the-last-watch',
      currentKey: opts.currentKey === undefined ? KEY : opts.currentKey,
      ownCharacters: opts.ownCharacters ?? [{ character_key: KEY, name: 'Kraggor' }],
      defaultRole: 'tank',
      onChange: () => {},
      now: NOW,
      viewerTimezone: opts.viewerTimezone ?? TZ,
      locale: 'en-US',
    },
  }).body;
}

/** The text inside the element carrying `testid`, tags and comments stripped. */
function text(markup: string, testid: string): string {
  const start = markup.indexOf(`data-testid="${testid}"`);
  if (start < 0) throw new Error(`no ${testid}`);
  const open = markup.lastIndexOf('<', start);
  const tag = /^\w+/.exec(markup.slice(open + 1))?.[0] ?? '';
  const close = markup.indexOf(`</${tag}>`, start);
  return markup
    .slice(markup.indexOf('>', start) + 1, close)
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, '')
    .replace(/\s+/g, ' ')
    .trim();
}

describe('RaidBandLine sentence', () => {
  it.each([
    ['in', 'tank', 'Thursday 8 pm, you are In as Tank, 23 of 40 signed up.', 'In'],
    ['late', 'healer', 'Thursday 8 pm, you are Late as Healer, 23 of 40 signed up.', 'Late'],
    ['bench', 'dps', 'Thursday 8 pm, you are on the Bench, 23 of 40 signed up.', 'Bench'],
    ['out', 'dps', 'Thursday 8 pm, you are Out, 23 of 40 signed up.', 'Out'],
  ] as const)('%s reads the spec sentence with the state word strong', (state, role, sentence, word) => {
    const body = html({ nextRaid: raid({ own: { character_key: KEY, state, role } }) });
    expect(text(body, 'guild-raid-line-sentence')).toBe(sentence);
    expect(body).toMatch(new RegExp(`<strong class="text-strong font-semibold">${word}</strong>`));
  });

  it('reads the unsigned sentence', () => {
    expect(text(html(), 'guild-raid-line-sentence')).toBe(
      'Thursday 8 pm, you have not signed up yet. 23 of 40 are in.',
    );
  });

  it('reads `{n} are in` while unsigned and the size is unknown', () => {
    expect(text(html({ nextRaid: raid({ size: null }) }), 'guild-raid-line-sentence')).toBe(
      'Thursday 8 pm, you have not signed up yet. 23 are in.',
    );
  });

  it('names the character when the account has two or more here', () => {
    const body = html({
      nextRaid: raid({ own: { character_key: ALT, state: 'in', role: 'healer' } }),
      ownCharacters: [
        { character_key: KEY, name: 'Kraggor' },
        { character_key: ALT, name: 'Sunderfel' },
      ],
    });
    expect(text(body, 'guild-raid-line-sentence')).toBe(
      'Thursday 8 pm, you are In on Sunderfel as Healer, 23 of 40 signed up.',
    );
  });

  it('adds the guild clock when the offsets differ', () => {
    expect(text(html({ viewerTimezone: 'America/New_York' }), 'guild-raid-line-sentence')).toBe(
      'Thursday 9 pm (8 pm guild time), you have not signed up yet. 23 of 40 are in.',
    );
  });
});

describe('RaidBandLine controls', () => {
  it('links the zone to the raid page as `{Zone} →`', () => {
    const body = html();
    expect(text(body, 'guild-raid-line-link')).toBe("Onyxia's Lair →");
    expect(body).toContain('href="/guild/us/normal/the-last-watch/raids/7"');
  });

  it('draws the band-size segments for the current character, checked from its own sign-up', () => {
    const body = html({ nextRaid: raid({ own: { character_key: KEY, state: 'late', role: 'dps' } }) });
    expect(body).toContain('data-testid="guild-raid-line-segments"');
    expect(body).toContain('md:h-9');
    expect(body).toMatch(/aria-checked="true"[^>]*data-testid="guild-raid-line-segments-late"/);
  });

  const TWO = [
    { character_key: KEY, name: 'Kraggor' },
    { character_key: ALT, name: 'Sunderfel' },
  ];

  it('acts on the character the sentence names when that one is In or Late', () => {
    for (const state of ['in', 'late'] as const) {
      const body = html({
        nextRaid: raid({ own: { character_key: ALT, state, role: 'healer' } }),
        ownCharacters: TWO,
      });
      expect(text(body, 'guild-raid-line-sentence')).toContain('on Sunderfel as Healer');
      expect(body).toMatch(
        new RegExp(`aria-checked="true"[^>]*data-testid="guild-raid-line-segments-${state}"`),
      );
      expect(body).toContain('aria-label="Sunderfel\'s sign-up"');
    }
  });

  it('acts on the current character when the other one is only on the Bench or Out', () => {
    const body = html({
      nextRaid: raid({ own: { character_key: ALT, state: 'bench', role: 'dps' } }),
      ownCharacters: TWO,
    });
    expect(body).not.toContain('aria-checked="true"');
    expect(body).toContain('aria-label="Kraggor\'s sign-up"');
  });

  it('draws no segments without a current character', () => {
    const body = html({ currentKey: null });
    expect(body).not.toContain('guild-raid-line-segments');
    expect(text(body, 'guild-raid-line-sentence')).toContain('you have not signed up');
  });

  it('draws no status line at rest, so the block stays within the 44 px cap', () => {
    expect(html()).not.toContain('guild-raid-line-status');
  });

  it('disables the segments on a locked raid and names the reason', () => {
    const body = html({ nextRaid: raid({ status: 'locked' }) });
    expect(text(body, 'guild-raid-line-locked')).toBe(raidCopy.raidStates.locked);
    expect(body).toMatch(/role="radiogroup"[^>]*aria-describedby="guild-raid-line-locked-7"/);
    expect(body).toMatch(/id="guild-raid-line-locked-7"/);
    expect([...body.matchAll(DISABLED_BUTTON)]).toHaveLength(4);
  });

  it('keeps the segments live on a planned raid', () => {
    const body = html({ nextRaid: raid({ status: 'planned' }) });
    expect(body).not.toContain('guild-raid-line-locked');
    expect(body).not.toMatch(DISABLED_BUTTON);
  });
});

describe('RaidBandLine none planned', () => {
  it('reads the member line alone', () => {
    const body = html({ nextRaid: null });
    expect(text(body, 'guild-raid-line-none')).toBe(raidCopy.bandNone.member);
    expect(body).not.toContain('#raids');
  });

  it('gives an officer the Plan a raid link to #raids', () => {
    const body = html({ nextRaid: null, officer: true });
    expect(text(body, 'guild-raid-line-none')).toBe(
      `${raidCopy.bandNone.officer} ${raidCopy.bandNone.planLink}`,
    );
    expect(body).toContain('href="#raids"');
  });
});

function view(characters: ViewerChar[]): GuildRaidView {
  return {
    id: 7,
    zone: 'onyxias-lair',
    zone_name: "Onyxia's Lair",
    title: "Onyxia's Lair",
    starts_at: STARTS,
    duration_minutes: 180,
    status: 'planned',
    size: 40,
    guild_timezone: TZ,
    notes_html: '',
    counts: { in: 21, late: 3, bench: 1, out: 0, tanks: 4, healers: 7, dps: 13 },
    signups: characters.flatMap((c) =>
      c.state === null
        ? []
        : [{ ...c, state: c.state, note: c.character_key === KEY ? 'Flasks ready' : '', own: true }],
    ),
    not_signed: [],
    viewer: { characters },
    discord: { available: true, connected: false, post_default: false },
  };
}

function viewerChar(key: string, state: 'in' | 'bench' | null): ViewerChar {
  return {
    character_key: key,
    name: key,
    class: 'warrior',
    spec: 'protection',
    verified: true,
    state,
    role: 'tank',
  };
}

describe('nextRaidAfter', () => {
  it('takes the counts and the current character sign-up from the answer', () => {
    const next = nextRaidAfter(raid(), view([viewerChar(KEY, 'in'), viewerChar(ALT, 'bench')]), KEY);
    expect(next.counts).toEqual({ in: 21, late: 3, tanks: 4, healers: 7, dps: 13 });
    expect(next.own).toEqual({ character_key: KEY, state: 'in', role: 'tank', note: 'Flasks ready' });
  });

  it('takes the status from the answer', () => {
    expect(nextRaidAfter(raid(), { ...view([]), status: 'locked' }, KEY).status).toBe('locked');
  });

  it('falls back to another own sign-up, then to none', () => {
    expect(nextRaidAfter(raid(), view([viewerChar(KEY, null), viewerChar(ALT, 'bench')]), KEY).own).toEqual({
      character_key: ALT,
      state: 'bench',
      role: 'tank',
      note: '',
    });
    expect(nextRaidAfter(raid(), view([viewerChar(KEY, null)]), KEY).own).toBeNull();
  });
});

describe('pressFailedLine', () => {
  it('reads the race lines on a 409', () => {
    expect(pressFailedLine(new GuildApiError('x', 409, 'raid_locked'))).toBe(raidCopy.races.locked);
    expect(pressFailedLine(new GuildApiError('x', 409, 'raid_cancelled'))).toBe(raidCopy.races.cancelled);
  });

  it('reads the generic line otherwise', () => {
    expect(pressFailedLine(new Error('down'))).toBe(raidCopy.stripStatus.failed);
  });
});

describe('nextRaidAfterRefusal', () => {
  it('locks the band on raid_locked', () => {
    expect(nextRaidAfterRefusal(raid(), new GuildApiError('x', 409, 'raid_locked'))).toEqual({
      ...raid(),
      status: 'locked',
    });
  });

  it('drops the raid from the band on raid_cancelled', () => {
    expect(nextRaidAfterRefusal(raid(), new GuildApiError('x', 409, 'raid_cancelled'))).toBeNull();
  });

  it('leaves the band as it was on any other failure', () => {
    expect(nextRaidAfterRefusal(raid(), new Error('down'))).toBeUndefined();
  });
});
