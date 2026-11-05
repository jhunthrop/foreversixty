// web/src/lib/guild/zone-catalogue.ts
// The form's `What` and `Group size` selects from the API's curated zones (guild events spec 3.1,
// 3.3): the groups in their order, each option's suffixes, the six guild event keys folded into one
// option plus a size, and a raid's open day. The API holds the facts; this only words them.
import type { GuildZoneView } from './api';
import { raidCopy, SEP } from './raid-copy';

/** The `What` option every guild event key folds into; its size is `Group size`'s. */
export const GUILD_EVENT_KEY = 'event:guild';
const RAID_KEY_PREFIX = 'raid:';
const DUNGEON_KEY_PREFIX = 'dungeon:';
/** `Group size`'s value when the event has no cap. */
const NO_LIMIT = '';

export const isRaidKey = (key: string): boolean => key.startsWith(RAID_KEY_PREFIX);
export const isEventKey = (key: string): boolean =>
  key === GUILD_EVENT_KEY || key.startsWith(`${GUILD_EVENT_KEY}-`);
export const isDungeonKey = (key: string): boolean => key.startsWith(DUNGEON_KEY_PREFIX);

/** The `What` value a stored zone key shows as. */
export const whatValueOf = (key: string): string => (isEventKey(key) ? GUILD_EVENT_KEY : key);

/** `Group size`'s value for an event key: `10`, or the no-limit value. */
export const eventSizeOf = (key: string): string =>
  key.startsWith(`${GUILD_EVENT_KEY}-`) ? key.slice(GUILD_EVENT_KEY.length + 1) : NO_LIMIT;

/** The key the pair (`Guild event`, size) maps to. */
export const eventKeyOf = (size: string): string =>
  size === NO_LIMIT ? GUILD_EVENT_KEY : `${GUILD_EVENT_KEY}-${size}`;

/** `No limit`, `5`, `10`, `15`, `20`, `40`, from the keys the API lists. */
export function groupSizeOptions(zones: readonly GuildZoneView[]): { value: string; label: string }[] {
  const sizes = zones.filter((zone) => zone.kind === 'event' && zone.size !== null);
  return [
    { value: NO_LIMIT, label: raidCopy.form.noLimit },
    ...sizes.map((zone) => ({ value: String(zone.size), label: String(zone.size) })),
  ];
}

/** `{ day: 'Wed 9 Dec', short: '9 Dec' }`, read in UTC: the open date is a calendar day, whatever
 *  zone the viewer is in. */
function opensParts(opens: string): { day: string; short: string } {
  const at = new Date(opens);
  const part = (option: Intl.DateTimeFormatOptions): string =>
    new Intl.DateTimeFormat('en-GB', { timeZone: 'UTC', ...option }).format(at);
  const short = `${part({ day: 'numeric' })} ${part({ month: 'short' })}`;
  return { day: `${part({ weekday: 'short' })} ${short}`, short };
}

/** `Wed 9 Dec`; `null` for a zone with no open day. */
export function opensDayOf(zone: GuildZoneView): string | null {
  return zone.opens === null ? null : opensParts(zone.opens).day;
}

/** The open day as a date input's value, `2026-12-09`; `null` for a zone with no open day. */
export function opensDateOf(zone: GuildZoneView): string | null {
  return zone.opens === null ? null : zone.opens.slice(0, 10);
}

/** Whether a `YYYY-MM-DD` on the guild's calendar is on or after the zone's open day. */
export function zoneIsOpenOn(zone: GuildZoneView, date: string): boolean {
  const opens = opensDateOf(zone);
  return opens === null || date >= opens;
}

/** `The Deadmines · 5 · 17 to 26`, `Onyxia's Lair · 40 · opens 9 Dec`; each part only when known,
 *  and the opens part only until the day arrives. */
export function zoneOptionLabel(zone: GuildZoneView, now: Date): string {
  const parts = [zone.name];
  if (zone.size !== null) parts.push(String(zone.size));
  if (zone.min_level !== null && zone.max_level !== null) {
    parts.push(`${zone.min_level} to ${zone.max_level}`);
  }
  if (zone.opens !== null && Date.parse(zone.opens) > now.getTime()) {
    parts.push(`opens ${opensParts(zone.opens).short}`);
  }
  return parts.join(SEP);
}

/** The `What` select: `Guild event`, then `Dungeons` in the API's level order, then `Raids`. */
export function whatOptions(
  zones: readonly GuildZoneView[],
  now: Date,
): { value: string; label: string; group: string }[] {
  const { groups } = raidCopy.form;
  const ofKind = (kind: GuildZoneView['kind'], group: string) =>
    zones
      .filter((zone) => zone.kind === kind)
      .map((zone) => ({ value: zone.key, label: zoneOptionLabel(zone, now), group }));
  return [
    { value: GUILD_EVENT_KEY, label: groups.event, group: groups.event },
    ...ofKind('dungeon', groups.dungeons),
    ...ofKind('raid', groups.raids),
  ];
}
