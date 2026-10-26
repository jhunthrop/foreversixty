// web/src/components/guild/RaidForm.test.ts
// The raid form rendered with svelte/server (raids spec 3.5): defaults from the last raid and for
// a first raid, the For you line, the repeat helper at the cap, and the Discord box or its line.
// Typing, the error lines and the submit are clicked through in the raids e2e.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidView, GuildZoneView } from '../../lib/guild/api';
import { raidCopy } from '../../lib/guild/raid-copy';
import RaidForm from './RaidForm.svelte';

const CHICAGO = 'America/Chicago';
const NEW_YORK = 'America/New_York';
const NOW = new Date('2026-12-02T18:00:00Z');
const ZONES: GuildZoneView[] = [
  { key: 'barrow-deeps', name: 'Barrow Deeps', size: null },
  { key: 'hyjal-summit', name: 'Hyjal Summit', size: null },
  { key: 'onyxias-lair', name: "Onyxia's Lair", size: 40 },
];

function lastRaid(): GuildRaidView {
  return {
    id: 7,
    zone: 'onyxias-lair',
    zone_name: "Onyxia's Lair",
    title: "Onyxia's Lair",
    starts_at: '2026-11-27T02:30:00Z',
    duration_minutes: 150,
    status: 'planned',
    size: 40,
    guild_timezone: CHICAGO,
    notes_html: '',
    counts: { in: 0, late: 0, bench: 0, out: 0, tanks: 0, healers: 0, dps: 0 },
    signups: [],
    not_signed: [],
    viewer: { characters: [] },
    discord: { connected: true, post_default: true },
  };
}

interface FormProps {
  last?: GuildRaidView | null;
  plannedCount?: number;
  discordConnected?: boolean;
  viewerTimezone?: string;
  mode?: 'create' | 'edit';
}

function html({
  last = null,
  plannedCount = 0,
  discordConnected = false,
  viewerTimezone = NEW_YORK,
  mode = 'create',
}: FormProps = {}): string {
  return render(RaidForm, {
    props: {
      mode,
      zones: ZONES,
      guildTimezone: CHICAGO,
      viewerTimezone,
      lastRaid: last,
      raid: mode === 'edit' ? lastRaid() : undefined,
      plannedCount,
      discordConnected,
      settingsHref: '/guild/us/pvp/olympus-xxvii/settings',
      onSubmit: () => Promise.resolve({ ok: true as const }),
      onCancel: () => {},
      now: NOW,
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

function tag(markup: string, testid: string): string {
  const match = new RegExp(`<[a-z]+[^>]*data-testid="${testid}"[^>]*>`).exec(markup);
  if (match === null) throw new Error(`no ${testid}`);
  return match[0];
}

describe('RaidForm', () => {
  it('labels every field in the spec’s order', () => {
    const words = text(html());
    const order = [
      raidCopy.form.raid,
      raidCopy.form.title,
      raidCopy.form.date,
      raidCopy.form.time,
      raidCopy.form.duration,
      raidCopy.form.guildZone,
      raidCopy.form.notes,
      raidCopy.form.repeat,
      raidCopy.form.submitCreate,
      raidCopy.form.cancel,
    ];
    const at = order.map((label) => words.indexOf(label));
    expect(at.every((i) => i >= 0)).toBe(true);
    expect([...at].sort((a, b) => a - b)).toEqual(at);
    expect(words).toContain(raidCopy.form.markdownHelper);
  });

  it('defaults to the last raid’s zone, next weekday, guild time and duration', () => {
    const markup = html({ last: lastRaid() });
    expect(markup).toMatch(/<option value="onyxias-lair"[^>]*selected/);
    expect(tag(markup, 'raid-form-date')).toContain('value="2026-12-03"');
    expect(tag(markup, 'raid-form-time')).toContain('value="20:30"');
    expect(markup).toMatch(/<option value="150"[^>]*selected/);
    expect(markup).toMatch(/<option value="America\/Chicago"[^>]*selected/);
    expect(tag(markup, 'raid-form-title')).toContain(`placeholder="Onyxia's Lair"`);
  });

  it('starts a first raid with Choose a raid, 20:00, no date and 3 h', () => {
    const markup = html();
    expect(markup).toContain(raidCopy.placeholders.chooseRaid);
    expect(tag(markup, 'raid-form-time')).toContain('value="20:00"');
    expect(tag(markup, 'raid-form-date')).not.toMatch(/value="\d/);
    expect(markup).toMatch(/<option value="180"[^>]*selected/);
  });

  it('draws the For you line in the viewer’s zone, or same time', () => {
    expect(text(html({ last: lastRaid() }))).toContain('For you: Thu 3 Dec, 9:30 pm.');
    expect(text(html({ last: lastRaid(), viewerTimezone: CHICAGO }))).toContain('For you: same time.');
  });

  it('says how many more fit and disables the repeat at the cap', () => {
    const markup = html({ plannedCount: 10 });
    expect(text(markup)).toContain('Only 2 more fit: a guild plans at most 12 raids ahead.');
    expect(tag(markup, 'raid-form-repeat')).toMatch(/ disabled/);
    expect(html({ plannedCount: 8 })).not.toContain('raid-form-repeat-helper');
  });

  it('shows the Settings line and no box without a Discord channel', () => {
    const markup = html();
    expect(markup).not.toContain('data-testid="raid-form-discord"');
    expect(text(markup)).toContain(raidCopy.form.connectDiscord);
    expect(markup).toContain('href="/guild/us/pvp/olympus-xxvii/settings"');
  });

  it('remembers the last Post to Discord answer when a channel is connected', () => {
    const markup = html({ last: lastRaid(), discordConnected: true });
    expect(tag(markup, 'raid-form-discord')).toContain('checked');
    expect(text(markup)).toContain(raidCopy.form.postToDiscord);
    expect(markup).not.toContain('raid-form-discord-line');
  });

  it('edits with Save changes, the moved-time label and no repeat', () => {
    const markup = html({ mode: 'edit', discordConnected: true });
    const words = text(markup);
    expect(words).toContain(raidCopy.form.submitEdit);
    expect(words).toContain(raidCopy.form.postToDiscordIfMoved);
    expect(words).not.toContain(raidCopy.form.repeat);
    expect(tag(markup, 'raid-form-date')).toContain('value="2026-11-26"');
  });
});
