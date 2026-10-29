// web/src/lib/guild/raid-form.test.ts
// The raid form's pure logic (design/specs/2026-10-26-guild-raids.md 3.5, 6.1): defaults from the
// guild's last planned raid, the 6.1 field errors, the repeat helper under the cap, and the
// payload in UTC from the guild's wall clock.
import { describe, expect, it } from 'vitest';
import type { GuildRaidDiscord, GuildRaidDiscordPost, GuildRaidView } from './api';
import { NO_LOG_YET } from '../../test-support/raid-fixtures';
import { raidCopy } from './raid-copy';
import {
  RAID_CAP,
  canRepeat,
  channelLabel,
  createDefaults,
  discordFormPart,
  discordHelper,
  discordSettingsHref,
  durationLabel,
  durationOptions,
  editDefaults,
  forYou,
  lastPlannedRaid,
  notesCounter,
  openRaidCount,
  raidPayload,
  repeatHelper,
  serverFieldFor,
  validateRaidForm,
  type RaidFormValues,
} from './raid-form';

const CHICAGO = 'America/Chicago';
const NEW_YORK = 'America/New_York';
// Wed 2 Dec 2026, 12:00 Chicago.
const NOW = new Date('2026-12-02T18:00:00Z');

function raid(extra: Partial<GuildRaidView> = {}): GuildRaidView {
  return {
    id: 7,
    zone: 'onyxias-lair',
    zone_name: "Onyxia's Lair",
    title: "Onyxia's Lair",
    // Thu 26 Nov 2026, 20:30 Chicago.
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
    ...extra,
  };
}

function values(extra: Partial<RaidFormValues> = {}): RaidFormValues {
  return {
    zone: 'onyxias-lair',
    title: '',
    date: '2026-12-03',
    time: '20:00',
    duration: 180,
    timezone: CHICAGO,
    notes: '',
    repeat: false,
    postToDiscord: false,
    remindDayBefore: true,
    ...extra,
  };
}

describe('durations', () => {
  it('runs 1 h to 4 h in 30 min steps', () => {
    expect(durationOptions().map((o) => o.label)).toEqual([
      '1 h',
      '1 h 30 min',
      '2 h',
      '2 h 30 min',
      '3 h',
      '3 h 30 min',
      '4 h',
    ]);
    expect(durationLabel(90)).toBe('1 h 30 min');
  });
});

describe('lastPlannedRaid and openRaidCount', () => {
  it('takes the latest-starting raid that is not cancelled', () => {
    const a = raid({ id: 1, starts_at: '2026-12-04T02:00:00Z' });
    const b = raid({ id: 2, starts_at: '2026-12-11T02:00:00Z' });
    const c = raid({ id: 3, starts_at: '2026-12-18T02:00:00Z', status: 'cancelled' });
    expect(lastPlannedRaid([b, c, a])?.id).toBe(2);
    expect(lastPlannedRaid([])).toBeNull();
  });

  it('counts planned and locked raids only, as the cap does', () => {
    const list = [
      raid(),
      raid({ status: 'locked' }),
      raid({ status: 'cancelled' }),
      raid({ status: 'done' }),
    ];
    expect(openRaidCount(list)).toBe(2);
  });
});

const CONNECTED = {
  available: true,
  connected: true,
  post_default: true,
  remind_default: true,
  channel_name: null,
};

describe('createDefaults', () => {
  it('takes the last raid’s zone, weekday, guild time and duration', () => {
    const d = createDefaults({
      last: raid(),
      now: NOW,
      guildTimezone: CHICAGO,
      viewerTimezone: NEW_YORK,
      discord: CONNECTED,
    });
    expect(d).toMatchObject({
      zone: 'onyxias-lair',
      date: '2026-12-03',
      time: '20:30',
      duration: 150,
      timezone: CHICAGO,
      repeat: false,
      postToDiscord: true,
      title: '',
      notes: '',
    });
  });

  it('starts a first raid at 20:00 with no date and no raid chosen, in the browser zone', () => {
    const d = createDefaults({
      last: null,
      now: NOW,
      guildTimezone: null,
      viewerTimezone: NEW_YORK,
      discord: null,
    });
    expect(d).toMatchObject({ zone: '', date: '', time: '20:00', duration: 180, timezone: NEW_YORK });
    expect(d.postToDiscord).toBe(false);
  });

  it('seeds the box from the guild channel with no raid planned', () => {
    const d = createDefaults({
      last: null,
      now: NOW,
      guildTimezone: null,
      viewerTimezone: NEW_YORK,
      discord: CONNECTED,
    });
    expect(d.postToDiscord).toBe(true);
  });

  it('prefers the guild zone over the browser zone', () => {
    const d = createDefaults({
      last: null,
      now: NOW,
      guildTimezone: CHICAGO,
      viewerTimezone: NEW_YORK,
      discord: null,
    });
    expect(d.timezone).toBe(CHICAGO);
  });
});

describe('editDefaults', () => {
  it('reads the raid back on the guild wall clock', () => {
    const d = editDefaults(raid({ title: 'Ony farm' }), CHICAGO);
    expect(d).toMatchObject({ zone: 'onyxias-lair', title: 'Ony farm', date: '2026-11-26', time: '20:30' });
  });

  it('leaves the title empty when it is the zone name', () => {
    expect(editDefaults(raid(), CHICAGO).title).toBe('');
  });

  it('starts Notes from the raw markdown, empty when an older view has none', () => {
    expect(editDefaults(raid({ notes: '**Flasks**' }), CHICAGO).notes).toBe('**Flasks**');
    expect(editDefaults(raid(), CHICAGO).notes).toBe('');
  });
});

describe('validateRaidForm', () => {
  const ctx = { now: NOW, plannedCount: 0, mode: 'create' as const };

  it('passes a complete future raid', () => {
    expect(validateRaidForm(values(), ctx)).toEqual({});
  });

  it('says the 6.1 lines under each field', () => {
    expect(validateRaidForm(values({ zone: '', date: '' }), ctx)).toEqual({
      zone: raidCopy.form.errors.raid,
      date: raidCopy.form.errors.date,
    });
    expect(validateRaidForm(values({ date: '2026-12-01' }), ctx)).toEqual({
      time: raidCopy.form.errors.time,
    });
    expect(validateRaidForm(values({ time: '' }), ctx)).toEqual({ time: raidCopy.form.errors.time });
  });

  it('refuses past the cap of 12, counting the repeat as four', () => {
    expect(validateRaidForm(values(), { ...ctx, plannedCount: RAID_CAP })).toEqual({
      form: raidCopy.form.errors.cap,
    });
    expect(validateRaidForm(values({ repeat: true }), { ...ctx, plannedCount: 9 })).toEqual({
      form: raidCopy.form.errors.cap,
    });
    expect(validateRaidForm(values({ repeat: true }), { ...ctx, plannedCount: 8 })).toEqual({});
  });

  it('lets an edit keep a start already in the past', () => {
    const keep = values({ date: '2026-11-26', time: '20:30' });
    const editCtx = {
      ...ctx,
      mode: 'edit' as const,
      keepStartsAt: '2026-11-27T02:30:00Z',
      plannedCount: RAID_CAP,
    };
    expect(validateRaidForm(keep, editCtx)).toEqual({});
  });
});

describe('the repeat helper', () => {
  it('counts the slots left, and speaks only when fewer than four are', () => {
    expect(repeatHelper(8)).toBeNull();
    expect(repeatHelper(10)).toBe('Only 2 more fit: a guild plans at most 12 raids ahead.');
    expect(repeatHelper(11)).toBe('Only 1 more fit: a guild plans at most 12 raids ahead.');
    expect(repeatHelper(RAID_CAP)).toBeNull();
  });

  it('allows a repeat only when all four raids fit', () => {
    expect(canRepeat(8)).toBe(true);
    expect(canRepeat(9)).toBe(false);
  });
});

describe('forYou', () => {
  it('renders the guild wall time in the viewer zone', () => {
    expect(forYou(values({ date: '2026-12-10' }), NEW_YORK, NOW)).toBe('For you: Thu 10 Dec, 9:00 pm.');
    expect(forYou(values({ date: '2026-12-10' }), CHICAGO, NOW)).toBe('For you: same time.');
  });

  it('reads an undated form as today on the guild calendar', () => {
    expect(forYou(values({ date: '' }), NEW_YORK, NOW)).toBe('For you: Wed 2 Dec, 9:00 pm.');
  });
});

describe('raidPayload', () => {
  it('converts the guild wall clock to UTC and repeats three weeks on create', () => {
    expect(
      raidPayload(values({ title: '  Ony  ', repeat: true, postToDiscord: true }), 'create', true),
    ).toEqual({
      zone: 'onyxias-lair',
      title: 'Ony',
      starts_at: '2026-12-04T02:00:00.000Z',
      duration_minutes: 180,
      notes: '',
      timezone: CHICAGO,
      post_to_discord: true,
      remind_day_before: true,
      repeat_weeks: 3,
    });
  });

  it('carries no repeat on edit', () => {
    expect(raidPayload(values({ repeat: true }), 'edit', true)).not.toHaveProperty('repeat_weeks');
  });

  it('never sends a remembered Discord flag the form does not draw', () => {
    expect(raidPayload(values({ postToDiscord: true }), 'create', false).post_to_discord).toBe(false);
  });
});

describe('discordFormPart', () => {
  it('draws the box when connected, the line when not, nothing when connect is off or unknown', () => {
    expect(
      discordFormPart({
        available: true,
        connected: true,
        post_default: false,
        remind_default: true,
        channel_name: null,
      }),
    ).toBe('box');
    expect(
      discordFormPart({
        available: true,
        connected: false,
        post_default: false,
        remind_default: true,
        channel_name: null,
      }),
    ).toBe('line');
    expect(
      discordFormPart({
        available: false,
        connected: false,
        post_default: false,
        remind_default: true,
        channel_name: null,
      }),
    ).toBe('none');
    expect(discordFormPart(null)).toBe('none');
  });

  it("links the line to the settings page's Discord section", () => {
    expect(discordSettingsHref('/guild/us/pvp/x/settings')).toBe('/guild/us/pvp/x/settings#discord');
  });
});

describe('notesCounter and serverFieldFor', () => {
  it('counts down from 1,880 characters', () => {
    expect(notesCounter(1879)).toBeNull();
    expect(notesCounter(1880)).toBe('120 left');
  });

  it('maps the API’s field names onto the form’s lines', () => {
    expect(serverFieldFor('zone')).toBe('zone');
    expect(serverFieldFor('starts_at')).toBe('time');
    expect(serverFieldFor('notes')).toBeNull();
  });
});

const BOT: GuildRaidDiscord = {
  available: true,
  connected: true,
  post_default: true,
  remind_default: true,
  channel_name: 'raid-signups',
};
const WEBHOOK: GuildRaidDiscord = { ...BOT, channel_name: null };
const posted = (extra: Partial<GuildRaidDiscordPost> = {}): GuildRaidDiscordPost => ({
  posted_at: '2026-11-30T18:00:00Z',
  post_at: '2026-11-30T18:00:00Z',
  reminded_at: null,
  remind: true,
  ...extra,
});

interface HelperArgs {
  values?: Partial<RaidFormValues>;
  discord?: GuildRaidDiscord;
  post?: GuildRaidDiscordPost | null;
  mode?: 'create' | 'edit';
}

function helper({ values: v = {}, discord = BOT, post = null, mode = 'create' }: HelperArgs = {}) {
  const filled = values({ postToDiscord: true, ...v });
  const startsAt = new Date(`${filled.date}T${filled.time}:00-06:00`);
  return discordHelper({ values: filled, discord, post, mode, startsAt, now: NOW, timezone: CHICAGO });
}

describe('discordHelper', () => {
  it('within a week posts now, further out names the day, off says so', () => {
    const channel = '#raid-signups';
    expect(helper({ values: { date: '2026-12-05' } }).post).toBe(raidCopy.form.postNow(channel));
    expect(helper({ values: { date: '2026-12-12' } }).post).toBe(
      raidCopy.form.postLater(channel, 'Saturday 5 December'),
    );
    expect(helper({ values: { date: '2026-12-12' }, discord: WEBHOOK }).post).toBe(
      raidCopy.form.postLater(null, 'Saturday 5 December'),
    );
    const off = helper({ values: { postToDiscord: false } });
    expect(off.post).toBe(raidCopy.form.postOff);
    expect(off.remind).toBe(raidCopy.form.remindNeedsPost);
    expect(helper().remind).toBe(raidCopy.form.remind);
  });

  it('names the guild-zone day a week before the start', () => {
    const later = helper({ values: { date: '2026-12-10', time: '20:00' } });
    expect(later.post).toBe(raidCopy.form.postLater('#raid-signups', 'Thursday 3 December'));
  });

  it('repeat says each raid posts a week before, and whether the first goes now', () => {
    expect(helper({ values: { repeat: true, date: '2026-12-05' } }).post).toBe(
      raidCopy.form.postRepeat(true),
    );
    expect(helper({ values: { repeat: true, date: '2026-12-20' } }).post).toBe(
      raidCopy.form.postRepeat(false),
    );
  });

  it('edit of a bot post shows the posted line; reminded shows Reminder sent', () => {
    const edit = helper({ mode: 'edit', post: posted() });
    expect(edit.postRow).toBe('posted');
    expect(edit.post).toBe(raidCopy.form.posted('#raid-signups'));
    expect(edit.remindRow).toBe('live');
    const done = helper({ mode: 'edit', post: posted({ reminded_at: '2026-12-01T00:00:00Z' }) });
    expect(done.remindRow).toBe('reminded');
    expect(done.remind).toBe(raidCopy.form.reminded);
  });

  it('edit of a webhook post keeps the if-the-time-changes box and a live reminder', () => {
    const edit = helper({ mode: 'edit', discord: WEBHOOK, post: posted(), values: { postToDiscord: false } });
    expect(edit.postRow).toBe('webhookMoved');
    expect(edit.post).toBeNull();
    expect(edit.remind).toBe(raidCopy.form.remind);
  });

  it('edit of a raid not yet posted draws the switch like create', () => {
    const edit = helper({ mode: 'edit', post: posted({ posted_at: null }) });
    expect(edit.postRow).toBe('switch');
  });

  it('channel name is cut at 60', () => {
    expect(channelLabel('x'.repeat(70))).toBe('#' + 'x'.repeat(60) + '…');
    expect(channelLabel('x'.repeat(60))).toBe('#' + 'x'.repeat(60));
    expect(channelLabel('raid-signups')).toBe('#raid-signups');
    expect(channelLabel(null)).toBeNull();
  });
});

describe('remind defaults', () => {
  const dflt = { last: null, now: NOW, guildTimezone: CHICAGO, viewerTimezone: CHICAGO };

  it('create reads the guild default', () => {
    expect(createDefaults({ ...dflt, discord: { ...BOT, remind_default: false } }).remindDayBefore).toBe(
      false,
    );
    expect(createDefaults({ ...dflt, discord: BOT }).remindDayBefore).toBe(true);
  });

  it('edit reads the raid’s own answer before the default', () => {
    const view = (extra: Partial<GuildRaidView>): GuildRaidView =>
      raid({ discord: { ...BOT, remind_default: true }, ...extra });
    expect(editDefaults(view({ discord_post: posted({ remind: false }) }), CHICAGO).remindDayBefore).toBe(
      false,
    );
    expect(editDefaults(view({ discord_post: null }), CHICAGO).remindDayBefore).toBe(true);
  });
});
