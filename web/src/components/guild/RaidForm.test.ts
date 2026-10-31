// web/src/components/guild/RaidForm.test.ts
// The raid form rendered with svelte/server (raids spec 3.5): defaults from the last raid and for
// a first raid, the For you line, the repeat helper at the cap, and the Discord box or its line.
// Typing, the error lines and the submit are clicked through in the raids e2e.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidDiscord, GuildRaidView, GuildZoneView } from '../../lib/guild/api';
import { raidCopy } from '../../lib/guild/raid-copy';
import { NO_LOG_YET } from '../../test-support/raid-fixtures';
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
    discord: {
      available: true,
      connected: true,
      post_default: true,
      remind_default: true,
      channel_name: null,
    },
    discord_post: null,
    ...NO_LOG_YET,
  };
}

const NOT_CONNECTED: GuildRaidDiscord = {
  available: true,
  connected: false,
  post_default: false,
  remind_default: true,
  channel_name: null,
};
const CONNECTED: GuildRaidDiscord = {
  available: true,
  connected: true,
  post_default: true,
  remind_default: true,
  channel_name: null,
};

interface FormProps {
  last?: GuildRaidView | null;
  plannedCount?: number;
  discord?: GuildRaidDiscord | null;
  viewerTimezone?: string;
  mode?: 'create' | 'edit';
  raid?: GuildRaidView;
}

function html({
  last = null,
  plannedCount = 0,
  discord = NOT_CONNECTED,
  viewerTimezone = NEW_YORK,
  mode = 'create',
  raid = lastRaid(),
}: FormProps = {}): string {
  return render(RaidForm, {
    props: {
      mode,
      zones: ZONES,
      guildTimezone: CHICAGO,
      viewerTimezone,
      lastRaid: last,
      raid: mode === 'edit' ? raid : undefined,
      plannedCount,
      discord,
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

  it('draws the repeat row as a switch', () => {
    const markup = html();
    expect(tag(markup, 'raid-form-repeat')).toContain('role="switch"');
    expect(tag(markup, 'raid-form-repeat')).toContain('aria-checked="false"');
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

  it('draws the repeat at the cap in the disabled look, described by the helper', () => {
    const capped = html({ plannedCount: 10 });
    expect(tag(capped, 'raid-form-repeat-label')).toMatch(/opacity-50/);
    expect(tag(capped, 'raid-form-repeat')).toContain('aria-describedby="raid-form-create-repeat-helper"');
    expect(tag(html({ plannedCount: 8 }), 'raid-form-repeat-label')).not.toMatch(/opacity-50/);
  });

  it('draws every select as the shared Select, with its chevron', () => {
    const markup = html({ last: lastRaid() });
    for (const field of ['zone', 'duration', 'timezone']) {
      expect(tag(markup, `raid-form-${field}`)).toMatch(/class="sel-native"/);
    }
    expect(markup.match(/data-testid="select-chevron"/g)?.length).toBe(3);
  });

  it('shows the Settings line and no box without a Discord channel', () => {
    const markup = html();
    expect(markup).not.toContain('data-testid="raid-form-discord"');
    expect(text(markup)).toContain(raidCopy.form.connectDiscord);
    expect(markup).toContain('href="/guild/us/pvp/olympus-xxvii/settings#discord"');
    // The inline link takes the 44 px ::after phone hit area of the gate link.
    expect(markup).toMatch(
      /class="[^"]*after:h-11[^"]*md:after:hidden[^"]*"[^>]*data-testid="raid-form-settings-link"/,
    );
  });

  it('remembers the last Post to Discord answer when a channel is connected', () => {
    const markup = html({ last: lastRaid(), discord: CONNECTED });
    expect(tag(markup, 'raid-form-discord')).toContain('checked');
    expect(text(markup)).toContain(raidCopy.form.postToDiscord);
    expect(markup).not.toContain('raid-form-discord-line');
  });

  it('offers the box from the page-level channel state with no raid planned', () => {
    const markup = html({ last: null, discord: CONNECTED });
    expect(tag(markup, 'raid-form-discord')).toContain('checked');
    expect(markup).not.toContain('raid-form-discord-line');
  });

  it('offers the Settings line, not the box, for an available but unconnected channel with no raid planned', () => {
    const markup = html({ last: null, discord: NOT_CONNECTED });
    expect(markup).not.toContain('data-testid="raid-form-discord"');
    expect(markup).toContain('raid-form-discord-line');
  });

  it('draws neither the box nor the line when connect is off or the channel is unknown', () => {
    for (const discord of [{ ...NOT_CONNECTED, available: false }, null]) {
      const markup = html({ discord });
      expect(markup).not.toContain('data-testid="raid-form-discord"');
      expect(markup).not.toContain('raid-form-discord-line');
    }
  });

  it('edits with Save changes, the switch rows and no repeat', () => {
    const markup = html({ mode: 'edit', discord: CONNECTED });
    const words = text(markup);
    expect(words).toContain(raidCopy.form.submitEdit);
    expect(words).toContain(raidCopy.form.postToDiscord);
    expect(words).toContain(raidCopy.form.remindDayBefore);
    expect(words).not.toContain(raidCopy.form.repeat);
    expect(tag(markup, 'raid-form-date')).toContain('value="2026-11-26"');
  });

  it('draws the create switch on, the reminder under it, with the channel in the helper', () => {
    const discord = { ...CONNECTED, channel_name: 'raid-signups' };
    const markup = html({ last: lastRaid(), discord });
    expect(tag(markup, 'raid-form-discord')).toContain('role="switch"');
    expect(tag(markup, 'raid-form-remind')).toContain('role="switch"');
    expect(text(markup)).toContain(raidCopy.form.postNow('#raid-signups'));
    expect(text(markup)).toContain(raidCopy.form.remind);
  });

  it('edit of a bot post shows the posted line and the live reminder', () => {
    const discord = { ...CONNECTED, channel_name: 'raid-signups' };
    const stamp = '2026-11-20T18:00:00Z';
    const post = { posted_at: stamp, post_at: stamp, reminded_at: null, remind: true };
    const markup = html({ mode: 'edit', discord, raid: { ...lastRaid(), discord, discord_post: post } });
    expect(markup).toContain('raid-form-discord-posted');
    expect(text(markup)).toContain(raidCopy.form.posted('#raid-signups'));
    expect(markup).not.toContain('data-testid="raid-form-discord"');
    expect(tag(markup, 'raid-form-remind')).not.toMatch(/ disabled(=|\s|\/|>)/);
  });

  it('edit of a webhook post keeps Post to Discord if the time changes', () => {
    const stamp = '2026-11-20T18:00:00Z';
    const post = { posted_at: stamp, post_at: stamp, reminded_at: null, remind: true };
    const markup = html({ mode: 'edit', discord: CONNECTED, raid: { ...lastRaid(), discord_post: post } });
    expect(text(markup)).toContain(raidCopy.form.postToDiscordIfMoved);
  });
});
