// web/src/components/guild/RaidSignupStrip.test.ts
// The viewer's sign-up strip rendered with svelte/server (initial state, no effects): one row per
// character, the current one first; the waiting row; the one spec picker (raid comp spec 3.2) and
// the note prefilled; the note counter from 120 characters; the locked and ended lines above
// disabled rows; the cancelled strip. Saving, the Saved. line, the no-spec hold and the race lines
// are clicked through in the raid e2e.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidSignup, GuildRaidView, GuildRaidViewerCharacter } from '../../lib/guild/api';
import { RAID_ENDED_LINE_ID } from '../../lib/guild/raid-page';
import { raidCopy } from '../../lib/guild/raid-copy';
import { NO_LOG_YET } from '../../test-support/raid-fixtures';
import RaidSignupStrip from './RaidSignupStrip.svelte';

function character(name: string, extra: Partial<GuildRaidViewerCharacter> = {}): GuildRaidViewerCharacter {
  return {
    character_key: `us-forever-${name.toLowerCase()}`,
    name,
    class: 'warrior',
    spec: 'Fury',
    verified: true,
    state: null,
    role: 'dps',
    ...extra,
  };
}

function raid(characters: GuildRaidViewerCharacter[], signups: GuildRaidSignup[] = []): GuildRaidView {
  return {
    id: 7,
    zone: 'onyxias-lair',
    zone_name: 'Onyxia’s Lair',
    title: 'Onyxia’s Lair',
    starts_at: '2026-12-10T02:00:00Z',
    duration_minutes: 180,
    status: 'planned',
    size: 40,
    guild_timezone: 'America/Chicago',
    notes_html: '',
    counts: { in: 0, late: 0, bench: 0, out: 0, tanks: 0, healers: 0, dps: 0 },
    signups,
    not_signed: [],
    viewer: { characters },
    discord: {
      available: true,
      connected: false,
      post_default: false,
      remind_default: true,
      channel_name: null,
    },
    discord_post: null,
    ...NO_LOG_YET,
  };
}

const KRAGGOR = character('Kraggor', {
  state: 'in',
  role: 'tank',
  spec: 'Protection',
  raid_spec: 'warrior-protection',
  default_spec: 'warrior-protection',
});
const MENDA = character('Menda', {
  class: 'priest',
  spec: 'Holy',
  raid_spec: null,
  default_spec: 'priest-holy',
});
const WAITING = character('Grimtotem', { class: 'shaman', spec: 'Enhancement', verified: false });

interface StripProps {
  view?: GuildRaidView;
  locked?: boolean;
  ended?: boolean;
  cancelled?: boolean;
  currentKey?: string | null;
}

function html({ view = raid([KRAGGOR, MENDA, WAITING]), ...flags }: StripProps = {}): string {
  return render(RaidSignupStrip, {
    props: {
      raid: view,
      locked: false,
      ended: false,
      cancelled: false,
      onSignup: () => Promise.resolve({ ok: true as const, moved: null }),
      ...flags,
    },
  }).body;
}

function text(markup: string): string {
  return markup
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
}

/** The native spec select of a row. */
function select(markup: string, key: string): string {
  return (
    markup.match(new RegExp(`<select[^>]*data-testid="signup-spec-${key}"[^>]*>[\\s\\S]*?</select>`))?.[0] ??
    ''
  );
}

/** The Who line under a row's name. */
function who(markup: string, key: string): string {
  return markup.match(new RegExp(`data-testid="signup-who-${key}">([^<]*)<`))?.[1] ?? '';
}

/** The row testids in document order. */
function rowOrder(markup: string): string[] {
  return [...markup.matchAll(/data-testid="signup-row-([^"]+)"/g)].map((m) => m[1]);
}

describe('RaidSignupStrip', () => {
  it('draws the sub-head and one row per character, the current one first', () => {
    const markup = html({ currentKey: MENDA.character_key });
    expect(text(markup)).toContain(raidCopy.subheads.yourSignup);
    expect(rowOrder(markup)).toEqual([MENDA.character_key, KRAGGOR.character_key, WAITING.character_key]);
  });

  it('gives each verified row its segments, one spec picker and the note', () => {
    const markup = html();
    expect(markup).toContain('aria-label="Kraggor\'s sign-up"');
    expect(markup).toContain('aria-label="Menda\'s sign-up"');
    expect(markup).toContain(`placeholder="${raidCopy.placeholders.note}"`);
    expect(markup).toContain('maxlength="140"');
    expect(markup).toContain(`aria-label="${raidCopy.spec.label('Kraggor')}"`);
    expect(markup).not.toContain('aria-label="Role"');
    expect(markup).not.toContain('signup-role-');
  });

  it('starts on the saved spec, else the default spec, else the placeholder', () => {
    const NONE = character('Nobody', { raid_spec: null, default_spec: null });
    const markup = html({ view: raid([KRAGGOR, MENDA, NONE]) });
    expect(select(markup, KRAGGOR.character_key)).toMatch(/<option value="warrior-protection"[^>]*selected/);
    expect(select(markup, MENDA.character_key)).toMatch(/<option value="priest-holy"[^>]*selected/);
    expect(select(markup, NONE.character_key)).not.toMatch(/<option value="[^"]+"[^>]*selected/);
    expect(text(markup)).toContain(raidCopy.spec.placeholder);
  });

  it('reads the spec and class under the name, the class alone with no spec', () => {
    const NONE = character('Nobody', { raid_spec: null, default_spec: null });
    const markup = html({ view: raid([KRAGGOR, MENDA, NONE]) });
    expect(text(who(markup, KRAGGOR.character_key))).toBe('Protection Warrior');
    expect(text(who(markup, MENDA.character_key))).toBe('Holy Priest');
    expect(text(who(markup, NONE.character_key))).toBe('Warrior');
  });

  it('offers one pair per spec, Arms and Fury adding a tank pair, and starts an off-tank on it', () => {
    const offTank = character('Brokk', { state: 'in', role: 'tank', raid_spec: 'warrior-fury' });
    const markup = html({ view: raid([offTank]) });
    const options = [
      ...select(markup, offTank.character_key).matchAll(/<option value="([^"]*)"[^>]*>([^<]*)</g),
    ]
      .map((m) => [m[1], m[2]])
      .filter(([value]) => value !== '');
    expect(options).toEqual([
      ['warrior-arms', 'Arms · DPS'],
      ['warrior-arms:tank', 'Arms · Tank'],
      ['warrior-fury', 'Fury · DPS'],
      ['warrior-fury:tank', 'Fury · Tank'],
      ['warrior-protection', 'Protection · Tank'],
    ]);
    expect(select(markup, offTank.character_key)).toMatch(/<option value="warrior-fury:tank"[^>]*selected/);
    expect(text(who(markup, offTank.character_key))).toBe('Fury Warrior · Tank');
  });

  it('keeps no empty spec block under any row (the 150 px reserve is gone)', () => {
    const markup = html({ view: raid([character('Aurelith', { class: 'paladin' }), MENDA]) });
    expect(markup).not.toContain('signup-spec-reserve-');
    expect(text(markup)).not.toContain('Which spec tonight?');
  });

  it('mutes a character waiting for approval, with the sentence and no controls', () => {
    const markup = html();
    const row = markup.split(`data-testid="signup-row-${WAITING.character_key}"`)[1] ?? '';
    expect(text(row)).toContain(raidCopy.waitingCharacter);
    expect(row).not.toContain("Grimtotem's sign-up");
    expect(markup).not.toContain(`signup-spec-${WAITING.character_key}`);
    expect(markup).not.toContain(`signup-note-${WAITING.character_key}`);
  });

  it('shows the note counter only from 120 characters', () => {
    const note = (length: number): GuildRaidSignup => ({
      character_key: KRAGGOR.character_key,
      name: 'Kraggor',
      class: 'warrior',
      spec: 'Protection',
      role: 'tank',
      state: 'in',
      note: 'x'.repeat(length),
      own: true,
    });
    expect(html({ view: raid([KRAGGOR], [note(119)]) })).not.toContain('signup-counter-');
    const markup = html({ view: raid([KRAGGOR], [note(125)]) });
    expect(markup).toContain(`value="${'x'.repeat(125)}"`);
    expect(text(markup)).toContain('15 left');
  });

  it('reserves a status line under every verified row', () => {
    const markup = html();
    expect(markup).toContain(`data-testid="signup-status-${KRAGGOR.character_key}"`);
    expect(markup).toContain(`data-testid="signup-status-${MENDA.character_key}"`);
  });

  it('says the raid is locked above the rows, and disables them', () => {
    const markup = html({ locked: true });
    const [above, rows] = markup.split('data-testid="signup-row-');
    expect(text(above ?? '')).toContain(raidCopy.raidStates.locked);
    expect(rows).toContain('aria-describedby="raid-strip-state-7"');
    expect(markup).toMatch(/<select[^>]*disabled/);
    expect(markup).toMatch(/<input[^>]*disabled/);
  });

  it('disables the rows of an ended raid and points them at the band line that says why', () => {
    const markup = html({ ended: true });
    expect(markup).not.toContain('data-testid="raid-strip-state"');
    expect(markup).toMatch(/<select[^>]*disabled/);
    expect(markup).toContain(`aria-describedby="${RAID_ENDED_LINE_ID}"`);
  });

  it('is replaced by the cancelled line when the raid is cancelled', () => {
    const markup = html({ cancelled: true });
    expect(text(markup)).toContain('This raid was cancelled.');
    expect(rowOrder(markup)).toEqual([]);
  });

  it('disables the spec picker with the row when the raid is locked', () => {
    expect(select(html({ locked: true }), KRAGGOR.character_key)).toMatch(/^<select[^>]*disabled/);
  });
});
