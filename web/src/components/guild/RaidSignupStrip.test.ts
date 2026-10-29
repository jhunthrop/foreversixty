// web/src/components/guild/RaidSignupStrip.test.ts
// The viewer's sign-up strip rendered with svelte/server (initial state, no effects): one row per
// character, the current one first; the waiting row; the role select and note prefilled; the note
// counter from 120 characters; the locked and ended lines above disabled rows; the cancelled
// strip. Saving, the Saved. line and the race lines are clicked through in the raid e2e.
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

const KRAGGOR = character('Kraggor', { state: 'in', role: 'tank', spec: 'Protection' });
const MENDA = character('Menda', { class: 'priest', spec: 'Holy' });
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

  it('gives each verified row its segments, role select and note', () => {
    const markup = html();
    expect(markup).toContain('aria-label="Kraggor\'s sign-up"');
    expect(markup).toContain('aria-label="Menda\'s sign-up"');
    expect(markup).toContain(`placeholder="${raidCopy.placeholders.note}"`);
    expect(markup).toContain('maxlength="140"');
    expect(markup).toContain('aria-label="Role"');
  });

  it('prefills the role: the signed-up role, else the spec role', () => {
    const markup = html();
    const select = (key: string): string =>
      markup.match(
        new RegExp(`<select[^>]*data-testid="signup-role-${key}"[^>]*>[\\s\\S]*?</select>`),
      )?.[0] ?? '';
    expect(select(KRAGGOR.character_key)).toMatch(/<option value="tank"[^>]*selected/);
    expect(select(MENDA.character_key)).toMatch(/<option value="healer"[^>]*selected/);
  });

  it('mutes a character waiting for approval, with the sentence and no controls', () => {
    const markup = html();
    const row = markup.split(`data-testid="signup-row-${WAITING.character_key}"`)[1] ?? '';
    expect(text(row)).toContain(raidCopy.waitingCharacter);
    expect(row).not.toContain("Grimtotem's sign-up");
    expect(markup).not.toContain(`signup-role-${WAITING.character_key}`);
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
    expect(text(markup)).toContain(raidCopy.raidStates.cancelled);
    expect(rowOrder(markup)).toEqual([]);
  });

  describe('spec ask', () => {
    const specSelect = (markup: string, key: string): string =>
      markup.match(
        new RegExp(`<select[^>]*data-testid="signup-spec-${key}"[^>]*>[\\s\\S]*?</select>`),
      )?.[0] ?? '';
    const PALADIN = character('Aurelith', { class: 'paladin', spec: 'Holy', state: 'in', role: 'healer' });

    it('asks In, Late and Bench rows with no saved spec, naming the class specs in full', () => {
      const markup = html({ view: raid([PALADIN]) });
      const select = specSelect(markup, PALADIN.character_key);
      expect(text(markup)).toContain(raidCopy.spec.title);
      expect(select).toContain(`>${raidCopy.spec.placeholder}<`);
      expect(['Holy', 'Protection', 'Retribution'].every((name) => select.includes(`>${name}<`))).toBe(true);
      expect(markup).toContain('signup-spec-us-forever-aurelith-label');
    });

    it('shows a saved spec chosen, under the Change spec label, with its icon', () => {
      const saved = { ...PALADIN, raid_spec: 'paladin-retribution' };
      const markup = html({ view: raid([saved]) });
      expect(specSelect(markup, saved.character_key)).toMatch(
        /<option value="paladin-retribution"[^>]*selected/,
      );
      expect(text(markup)).toContain(raidCopy.spec.change);
      expect(text(markup)).not.toContain(raidCopy.spec.title);
      expect(markup).toMatch(/<img[^>]*data-testid="signup-spec-us-forever-aurelith-icon"/);
    });

    it('does not ask an unsigned row or an Out row, and reserves the block on the unsigned one', () => {
      const unsigned = html({ view: raid([character('Aurelith', { class: 'paladin' })]) });
      expect(unsigned).not.toContain('<select data-testid="signup-spec-');
      expect(unsigned).not.toMatch(/<select[^>]*signup-spec-/);
      expect(unsigned).toContain('signup-spec-reserve-us-forever-aurelith');
      const out = html({ view: raid([{ ...PALADIN, state: 'out' }]) });
      expect(out).not.toMatch(/<select[^>]*signup-spec-/);
      expect(out).not.toContain('signup-spec-reserve-');
    });

    it('collapses a saved spec into the toggle, closed', () => {
      const markup = html({ view: raid([{ ...PALADIN, raid_spec: 'paladin-retribution' }]) });
      expect(markup).toMatch(
        /<button[^>]*aria-expanded="false"[^>]*data-testid="signup-spec-us-forever-aurelith-toggle"/,
      );
      expect(text(markup)).toContain(`Retribution · ${raidCopy.spec.change}`);
      expect(markup).toMatch(/<div id="signup-spec-us-forever-aurelith-panel" hidden/);
    });

    it('disables the select with the row when the raid is locked', () => {
      expect(specSelect(html({ view: raid([PALADIN]), locked: true }), PALADIN.character_key)).toMatch(
        /^<select[^>]*disabled/,
      );
    });
  });
});
