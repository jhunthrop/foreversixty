// web/src/lib/guild/raid-zones.ts
// The form's `Guild time zone` select (design/specs/2026-10-26-guild-raids.md 3.5): IANA zones
// grouped by region, labelled `Central (Chicago) · CST`. The abbreviation comes from Intl at the
// given instant, so it reads CDT in summer. A zone outside the list (a guild's stored zone, a
// browser's own) is added under `Other` with its IANA name, so the select always holds the value.

export interface TimeZoneGroup {
  region: string;
  zones: readonly { tz: string; name: string }[];
}

export const TIME_ZONE_GROUPS: readonly TimeZoneGroup[] = [
  {
    region: 'Americas',
    zones: [
      { tz: 'America/New_York', name: 'Eastern (New York)' },
      { tz: 'America/Chicago', name: 'Central (Chicago)' },
      { tz: 'America/Denver', name: 'Mountain (Denver)' },
      { tz: 'America/Phoenix', name: 'Mountain (Phoenix)' },
      { tz: 'America/Los_Angeles', name: 'Pacific (Los Angeles)' },
      { tz: 'America/Anchorage', name: 'Alaska (Anchorage)' },
      { tz: 'Pacific/Honolulu', name: 'Hawaii (Honolulu)' },
      { tz: 'America/Halifax', name: 'Atlantic (Halifax)' },
      { tz: 'America/Sao_Paulo', name: 'Brasília (São Paulo)' },
    ],
  },
  {
    region: 'Europe',
    zones: [
      { tz: 'Europe/London', name: 'UK (London)' },
      { tz: 'Europe/Paris', name: 'Central Europe (Paris)' },
      { tz: 'Europe/Berlin', name: 'Central Europe (Berlin)' },
      { tz: 'Europe/Helsinki', name: 'Eastern Europe (Helsinki)' },
      { tz: 'Europe/Moscow', name: 'Moscow' },
    ],
  },
  {
    region: 'Oceania',
    zones: [
      { tz: 'Australia/Perth', name: 'Western Australia (Perth)' },
      { tz: 'Australia/Adelaide', name: 'Central Australia (Adelaide)' },
      { tz: 'Australia/Brisbane', name: 'Eastern Australia (Brisbane)' },
      { tz: 'Australia/Sydney', name: 'Eastern Australia (Sydney)' },
      { tz: 'Pacific/Auckland', name: 'New Zealand (Auckland)' },
    ],
  },
];

const OTHER_REGION = 'Other';

export interface TimeZoneOptionGroup {
  region: string;
  options: { value: string; label: string }[];
}

function abbreviation(tz: string, at: Date): string {
  return (
    new Intl.DateTimeFormat('en-US', { timeZone: tz, timeZoneName: 'short' })
      .formatToParts(at)
      .find((part) => part.type === 'timeZoneName')?.value ?? tz
  );
}

/** The select's groups at `at`, with every zone of `include` that the list lacks under `Other`. */
export function timeZoneOptions(at: Date, include: readonly string[] = []): TimeZoneOptionGroup[] {
  const groups = TIME_ZONE_GROUPS.map(({ region, zones }) => ({
    region,
    options: zones.map(({ tz, name }) => ({ value: tz, label: `${name} · ${abbreviation(tz, at)}` })),
  }));
  const known = new Set(TIME_ZONE_GROUPS.flatMap((group) => group.zones.map((zone) => zone.tz)));
  const extra = [...new Set(include)].filter((tz) => tz !== '' && !known.has(tz));
  if (extra.length === 0) return groups;
  const other = extra.map((tz) => ({
    value: tz,
    label: `${tz.replaceAll('_', ' ')} · ${abbreviation(tz, at)}`,
  }));
  return [...groups, { region: OTHER_REGION, options: other }];
}
