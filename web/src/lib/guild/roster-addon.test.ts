// web/src/lib/guild/roster-addon.test.ts
// Task L8.3: the Roster tab's "In game" column, read from the guild's latest addon roster
// snapshot (Task L6.14's `roster[].addon` and `addon_snapshot` on the guild home).
import { describe, expect, it } from 'vitest';
import type { GuildAddonSnapshot, GuildRosterAddon, GuildRosterRow } from './api';
import { guildRosterAddonCopy } from './copy';
import {
  addonGapLabel,
  addonIdentity,
  addonLevelSpec,
  addonSeen,
  addonSeenAt,
  addonSeenClock,
  addonSnapshotLine,
  addonSpecName,
  snapshotSenderName,
} from './roster-addon';

const ADDON: GuildRosterAddon = {
  level: 60,
  spec: 'hunter-marksmanship',
  gear_gap: 12,
  build: 'Marksmanship raid',
  seen_at: '2026-12-22 21:40',
  snapshot_at: '2026-12-22T21:43:00Z',
};

function row(overrides: Partial<GuildRosterRow> = {}): GuildRosterRow {
  return {
    character_key: 'us/pvp/zulmara',
    region: 'us',
    ruleset: 'pvp',
    name: 'Zulmara',
    class: 'hunter',
    spec: 'Marksmanship',
    rank: 'member',
    verified: true,
    logged_recently: true,
    consent: 'gear_bags',
    may_remove: false,
    addon: ADDON,
    ...overrides,
  };
}

const SNAPSHOT: GuildAddonSnapshot = {
  received_at: '2026-12-22T21:43:00Z',
  by_character: 'us/pvp/zulmara',
  members: 4,
};

describe('addonSpecName', () => {
  it('names a spec slug the way the client names its talent tree', () => {
    expect(addonSpecName('hunter-marksmanship')).toBe('Marksmanship');
    expect(addonSpecName('hunter-beast-mastery')).toBe('Beast Mastery');
  });
  it('prints an unknown slug unchanged rather than guessing', () => {
    expect(addonSpecName('hunter-new-tree')).toBe('hunter-new-tree');
  });
});

describe('addonIdentity: the level and spec from the snapshot, where they differ from the site’s', () => {
  it('is null for a row the snapshot did not carry', () => {
    expect(addonIdentity(row({ addon: null }))).toBeNull();
    expect(addonIdentity(row({ addon: undefined }))).toBeNull();
  });
  it('drops the spec when the snapshot agrees with the site’s spec', () => {
    expect(addonIdentity(row())).toEqual({ level: 60, spec: null });
  });
  it('names the snapshot’s spec when it differs from the site’s', () => {
    expect(addonIdentity(row({ addon: { ...ADDON, spec: 'hunter-survival' } }))).toEqual({
      level: 60,
      spec: 'Survival',
    });
  });
  it('names the snapshot’s spec when the site has none', () => {
    expect(addonIdentity(row({ spec: undefined }))).toEqual({ level: 60, spec: 'Marksmanship' });
  });
  it('reads a consent-withheld line (null level and spec) as null, never a guess', () => {
    expect(
      addonIdentity(row({ addon: { ...ADDON, level: null, spec: null, gear_gap: null, build: null } })),
    ).toEqual({ level: null, spec: null });
  });
  it('always carries the level: the site’s roster row has none of its own to agree with', () => {
    expect(addonIdentity(row({ addon: { ...ADDON, level: 52 } }))?.level).toBe(52);
  });
});

describe('addonGapLabel', () => {
  it('reads the addon’s integer gap under its own heading, as DPS', () => {
    expect(addonGapLabel(12)).toBe('Gear gap 12 DPS');
    expect(addonGapLabel(0)).toBe('Gear gap 0 DPS');
    // The contract allows a negative gap (a member already past the plan).
    expect(addonGapLabel(-3)).toBe('Gear gap -3 DPS');
    expect(addonGapLabel(12)).toBe(guildRosterAddonCopy.gap(12));
    expect(addonGapLabel(null)).toBeNull();
  });
});

describe('addonSeenClock', () => {
  it('takes HH:MM from the member’s local "YYYY-MM-DD HH:MM"', () => {
    expect(addonSeenClock('2026-12-22 21:40')).toBe('21:40');
    expect(addonSeenClock('2026-12-21 09:05')).toBe('09:05');
  });
  it('is null for anything else, never a guessed time', () => {
    expect(addonSeenClock('')).toBeNull();
    expect(addonSeenClock('2026-12-22T21:40:00Z')).toBeNull();
    expect(addonSeenClock('yesterday')).toBeNull();
  });
});

describe('addonSeen: how long ago, relative to now, in the viewer\u2019s own clock', () => {
  const now = new Date(2026, 11, 22, 21, 46);
  it('reads "seen N hours ago", hovering the member\u2019s clock', () => {
    expect(addonSeen('2026-12-22 19:46', now)).toEqual({
      label: 'seen 2 hours ago',
      hover: "Seen at 19:46 on the member's clock",
    });
  });
  it('reads minutes under an hour and days beyond a day', () => {
    expect(addonSeen('2026-12-22 21:40', now)?.label).toBe('seen 6 minutes ago');
    expect(addonSeen('2026-12-22 21:46', now)?.label).toBe('seen just now');
    expect(addonSeen('2026-12-20 08:05', now)?.label).toBe('seen 2 days ago');
  });
  it('a stamp ahead of the viewer\u2019s clock reads "just now", never a negative age', () => {
    expect(addonSeen('2026-12-23 00:10', now)?.label).toBe('seen just now');
  });
  it('a withholding member\u2019s hover says so', () => {
    expect(addonSeen('2026-12-22 21:40', now, true)?.hover).toBe('Details withheld by the member');
  });
  it('is null for a seen_at it cannot read', () => {
    expect(addonSeen('yesterday', now)).toBeNull();
    expect(addonSeenAt('yesterday')).toBeNull();
  });
  it('takes the stamp as the viewer\u2019s local time', () => {
    expect(addonSeenAt('2026-12-22 21:40')?.getTime()).toBe(new Date(2026, 11, 22, 21, 40).getTime());
  });
});

describe('snapshotSenderName', () => {
  it('names the sender from the roster row with that character key', () => {
    expect(snapshotSenderName('us/pvp/zulmara', [row()])).toBe('Zulmara');
  });
  it('falls back to the key’s last segment, capitalised, when the sender is not on the roster', () => {
    expect(snapshotSenderName('us/pvp/brightwick', [row()])).toBe('Brightwick');
  });
});

describe('addonSnapshotLine', () => {
  const now = new Date('2026-12-22T21:46:00Z');
  it('says who sent it and when; the counts have their own line', () => {
    expect(addonSnapshotLine(SNAPSHOT, [row()], now)).toBe(
      'Roster from the addon \u00b7 sent by Zulmara 3 minutes ago',
    );
  });
  it('says no roster has arrived yet when there is none', () => {
    expect(addonSnapshotLine(null, [row()], now)).toBe(
      'No roster from the addon yet. Log in with the addon on your guild master and it sends one.',
    );
    expect(addonSnapshotLine(undefined, [row()], now)).toBe(guildRosterAddonCopy.noSnapshot);
  });
});

describe('addonLevelSpec', () => {
  it('joins the level and the spec with one space, or reads the level alone', () => {
    expect(addonLevelSpec(60, 'Survival')).toBe('Level 60 Survival');
    expect(addonLevelSpec(60, null)).toBe('Level 60');
  });
});
