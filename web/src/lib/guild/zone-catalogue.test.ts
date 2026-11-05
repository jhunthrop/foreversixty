// web/src/lib/guild/zone-catalogue.test.ts
// The What and Group size selects (guild events spec 3.1, 3.3). The real catalogue is read from
// data/curated/raids/events.json, so the groups and level order are the client's own rows.
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import type { GuildZoneView } from './api';
import {
  GUILD_EVENT_KEY,
  eventKeyOf,
  eventSizeOf,
  groupSizeOptions,
  isEventKey,
  isRaidKey,
  opensDateOf,
  opensDayOf,
  whatOptions,
  whatValueOf,
  zoneIsOpenOn,
  zoneOptionLabel,
} from './zone-catalogue';

const BEFORE = new Date('2026-11-05T12:00:00Z');
const AFTER = new Date('2026-12-09T12:00:00Z');
const OPENS = '2026-12-09T00:00:00Z';

const events = JSON.parse(
  readFileSync(new URL('../../../../data/curated/raids/events.json', import.meta.url), 'utf8'),
) as { zones: (Omit<GuildZoneView, 'opens'> & { opens?: string | null })[] };

const ony: GuildZoneView = {
  key: 'raid:onyxias-lair',
  name: "Onyxia's Lair",
  kind: 'raid',
  size: 40,
  opens: OPENS,
  min_level: null,
  max_level: null,
  key_item: null,
};
const catalogue: GuildZoneView[] = [...events.zones.map((zone) => ({ ...zone, opens: null })), ony];
const byKey = (key: string): GuildZoneView => {
  const found = catalogue.find((zone) => zone.key === key);
  if (found === undefined) throw new Error(`no zone ${key}`);
  return found;
};

describe('the event keys', () => {
  it('folds the six keys into one What value and a size', () => {
    expect(whatValueOf('event:guild-10')).toBe(GUILD_EVENT_KEY);
    expect(whatValueOf('dungeon:the-deadmines')).toBe('dungeon:the-deadmines');
    expect(eventSizeOf('event:guild-10')).toBe('10');
    expect(eventSizeOf('event:guild')).toBe('');
    expect(eventKeyOf('40')).toBe('event:guild-40');
    expect(eventKeyOf('')).toBe('event:guild');
    expect(isEventKey('event:guild-5')).toBe(true);
    expect(isEventKey('dungeon:the-deadmines')).toBe(false);
    expect(isRaidKey('raid:onyxias-lair')).toBe(true);
    expect(isRaidKey('event:guild')).toBe(false);
  });

  it('offers No limit and the catalogue sizes', () => {
    expect(groupSizeOptions(catalogue).map((option) => option.label)).toEqual([
      'No limit',
      '5',
      '10',
      '15',
      '20',
      '40',
    ]);
  });
});

describe('the option labels', () => {
  it('reads name, size and levels, each part only when known', () => {
    expect(zoneOptionLabel(byKey('dungeon:the-deadmines'), BEFORE)).toBe('The Deadmines · 5 · 17 to 26');
    expect(zoneOptionLabel(byKey('dungeon:upper-blackrock-spire'), BEFORE)).toBe(
      'Upper Blackrock Spire · 10 · 59 to 60',
    );
    expect(zoneOptionLabel({ ...ony, size: null, name: 'Barrow Deeps' }, BEFORE)).toBe(
      'Barrow Deeps · opens 9 Dec',
    );
  });

  it('drops the opens suffix from the open day on', () => {
    expect(zoneOptionLabel(ony, BEFORE)).toBe("Onyxia's Lair · 40 · opens 9 Dec");
    expect(zoneOptionLabel(ony, AFTER)).toBe("Onyxia's Lair · 40");
  });
});

describe('the What groups', () => {
  it('runs Guild event, Dungeons in level order, Raids', () => {
    const options = whatOptions(catalogue, BEFORE);
    expect([...new Set(options.map((option) => option.group))]).toEqual(['Guild event', 'Dungeons', 'Raids']);
    expect(options[0]).toEqual({ value: GUILD_EVENT_KEY, label: 'Guild event', group: 'Guild event' });
    expect(options.filter((option) => option.value.startsWith('event:'))).toHaveLength(1);
    const levels = catalogue.filter((zone) => zone.kind === 'dungeon').map((zone) => zone.min_level ?? 0);
    expect(levels).toEqual([...levels].sort((a, b) => a - b));
  });

  it('lists each wing as its own run', () => {
    const names = catalogue.map((zone) => zone.name);
    expect(names).toEqual(
      expect.arrayContaining([
        'Scarlet Monastery: Graveyard',
        'Scarlet Monastery: Library',
        'Scarlet Monastery: Armory',
        'Scarlet Monastery: Cathedral',
        'Stratholme: Live',
        'Stratholme: Undead',
        'Lower Blackrock Spire',
        'Upper Blackrock Spire',
      ]),
    );
  });
});

describe('a raid’s open day', () => {
  it('reads as a calendar day, whatever the viewer’s zone', () => {
    expect(opensDayOf(ony)).toBe('Wed 9 Dec');
    expect(opensDateOf(ony)).toBe('2026-12-09');
    expect(opensDayOf(byKey('dungeon:the-deadmines'))).toBeNull();
  });

  it('opens on the day itself, not before', () => {
    expect(zoneIsOpenOn(ony, '2026-12-08')).toBe(false);
    expect(zoneIsOpenOn(ony, '2026-12-09')).toBe(true);
    expect(zoneIsOpenOn(byKey('dungeon:the-deadmines'), '2026-11-06')).toBe(true);
  });
});
