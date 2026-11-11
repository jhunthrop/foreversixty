// web/src/components/guild/guild-events.test.ts
// Dungeon runs and guild events on the raid surfaces (guild events spec 3.4, 3.5, 5): the page's
// words follow the zone's kind, the row's zone line carries an event's size, the needs line reads
// raids only, and a refusal from the API lands under the date.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { GuildApiError, type NextRaidView } from '../../lib/guild/api';
import { plannedRowView } from '../../lib/guild/planned-row';
import { kindOf, nextRaidDeadline, raidDeadlineStart } from '../../lib/guild/raid-copy';
import { submitRefusal } from '../../lib/guild/raids-tab-controller';
import { failureLine } from '../../lib/guild/signup-strip';
import { plannedRaid } from '../../test-support/raid-fixtures';
import RaidCancelConfirm from './RaidCancelConfirm.svelte';
import RaidSignupStrip from './RaidSignupStrip.svelte';
import ZoneKeyNote from './ZoneKeyNote.svelte';

const NOW = new Date('2026-11-05T18:00:00Z');
const CHICAGO = 'America/Chicago';
const SEAL = { id: 12344, name: 'Seal of Ascension', quality: 3 };
const text = (markup: string): string =>
  markup
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();

describe('the planned row zone line', () => {
  const run = (over: Parameters<typeof plannedRaid>[0]) =>
    plannedRowView(plannedRaid(over), NOW, CHICAGO, 'en-US');

  it('reads the zone with the size an officer picked for an event', () => {
    const row = run({
      zone: 'event:guild-10',
      zone_kind: 'event',
      zone_name: 'Guild event',
      title: 'PvP premade',
      size: 10,
    });
    expect(row.zoneLine).toBe('Guild event · 10 players');
  });

  it('leaves out a no-limit event’s size, and counts it bare', () => {
    const row = run({
      zone: 'event:guild',
      zone_kind: 'event',
      zone_name: 'Guild event',
      title: 'Quest group',
      size: null,
    });
    expect(row.zoneLine).toBe('Guild event');
    expect(row.fill).toBeNull();
    expect(row.count).toMatch(/signed up$/);
  });

  it('counts a dungeon run against its five', () => {
    const counts = { ...plannedRaid().counts, in: 2, late: 0 };
    const row = run({
      zone: 'dungeon:the-deadmines',
      zone_kind: 'dungeon',
      zone_name: 'The Deadmines',
      title: 'The Deadmines',
      size: 5,
      counts,
    });
    expect(row.count).toBe('2/5');
    expect(row.fill).toBe(40);
    expect(row.zoneLine).toBeNull();
  });
});

describe('words by kind', () => {
  it('reads a view from before kinds as a raid', () => {
    expect(kindOf({})).toBe('raid');
    expect(kindOf({ zone_kind: 'dungeon' })).toBe('dungeon');
  });

  it('asks to cancel the run or the event', () => {
    const confirm = (kind: 'dungeon' | 'event'): string =>
      text(
        render(RaidCancelConfirm, {
          props: {
            kind,
            postCancel: false,
            busy: false,
            onPostChange: () => {},
            onConfirm: () => {},
            onKeep: () => {},
          },
        }).body,
      );
    expect(confirm('dungeon')).toContain('Cancel this run? Members see it as cancelled.');
    expect(confirm('dungeon')).toContain('Cancel the run');
    expect(confirm('event')).toContain('Cancel this event? Members see it as cancelled.');
    expect(confirm('event')).toContain('Cancel the event');
  });

  it('says a cancelled run was cancelled', () => {
    const raid = plannedRaid({ status: 'cancelled', zone_kind: 'dungeon' });
    const strip = render(RaidSignupStrip, {
      props: {
        raid,
        locked: false,
        ended: false,
        cancelled: true,
        onSignup: () => Promise.reject(new Error()),
      },
    }).body;
    expect(text(strip)).toContain('This run was cancelled.');
  });

  it('says an ended event has ended when a save is refused', () => {
    expect(failureLine('raid_ended', 'event')).toEqual({ message: 'This event has ended.', race: true });
    expect(failureLine('raid_ended')).toEqual({ message: 'This raid has ended.', race: true });
  });
});

describe('the key note', () => {
  it('names the key in its quality colour', () => {
    const markup = render(ZoneKeyNote, { props: { item: SEAL } }).body;
    expect(text(markup)).toMatch(/^Needs a key: Seal of Ascension ?\.$/);
    expect(markup).toContain('item-hover-12344');
  });

  it('names the door a key opens when the run has a keyless way in, never Needs a key', () => {
    const key = { id: 12382, name: 'Key to the City', quality: 2, door: 'the service entrance' };
    const markup = render(ZoneKeyNote, { props: { item: key } }).body;
    expect(text(markup)).toMatch(/^Key to the City ?opens the service entrance\.$/);
    expect(text(markup)).not.toContain('Needs a key');
    expect(markup).toContain('item-hover-12382');
  });
});

describe('the needs line deadline', () => {
  const next = (over: Partial<NextRaidView>): NextRaidView => ({
    id: 1,
    zone: 'raid:onyxias-lair',
    zone_name: "Onyxia's Lair",
    title: 'Ony',
    starts_at: '2026-12-10T02:00:00Z',
    size: 40,
    guild_timezone: CHICAGO,
    counts: { in: 0, late: 0, tanks: 0, healers: 0, dps: 0 },
    own: null,
    ...over,
  });

  it('names the raid’s weekday', () => {
    expect(nextRaidDeadline(raidDeadlineStart(next({ zone_kind: 'raid' })), CHICAGO)).toBe(
      'before Wednesday',
    );
  });

  it('falls back when a dungeon run or an event leads the band', () => {
    const run = next({ zone_kind: 'dungeon', zone: 'dungeon:the-deadmines' });
    expect(raidDeadlineStart(run)).toBeUndefined();
    expect(nextRaidDeadline(raidDeadlineStart(run), CHICAGO)).toBe('before the next raid');
    expect(raidDeadlineStart(null)).toBeUndefined();
  });
});

describe('the API refusing a raid before it opens', () => {
  it('shows its own line under the date', () => {
    const line = "Onyxia's Lair opens Wed 9 Dec. Pick that day or later.";
    const refusal = new GuildApiError('invalid', 422, 'invalid', { starts_at: line });
    expect(submitRefusal(refusal)).toEqual({ ok: false, code: 'invalid', message: line, field: 'date' });
  });

  it('keeps the plain time line for any other start refusal', () => {
    const refusal = new GuildApiError('invalid', 422, 'invalid', { starts_at: 'Pick a time after now.' });
    expect(submitRefusal(refusal)).toMatchObject({ field: 'time', message: 'Pick a time after now.' });
  });
});
