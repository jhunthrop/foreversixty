// web/src/test-support/raid-fixtures.ts
// One planned raid view for the Raids tab's component tests: Onyxia's Lair on Thu 10 Dec 2026,
// 8:00 pm Chicago, 23 of 40 signed up.
import type { GuildRaidView, GuildRaidViewerCharacter } from '../lib/guild/api';

export const CHICAGO = 'America/Chicago';
export const NEW_YORK = 'America/New_York';
/** Tue 8 Dec 2026, 12:00 Chicago: the raid is in 2 days 8 h. */
export const NOW = new Date('2026-12-08T18:00:00Z');

export function viewerCharacter(extra: Partial<GuildRaidViewerCharacter> = {}): GuildRaidViewerCharacter {
  return {
    character_key: 'us-pvp-kraggor',
    name: 'Kraggor',
    class: 'warrior',
    spec: 'Fury',
    verified: true,
    state: 'in',
    role: 'dps',
    ...extra,
  };
}

export function plannedRaid(extra: Partial<GuildRaidView> = {}): GuildRaidView {
  return {
    id: 7,
    zone: 'onyxias-lair',
    zone_name: "Onyxia's Lair",
    title: "Onyxia's Lair",
    starts_at: '2026-12-11T02:00:00Z',
    duration_minutes: 180,
    status: 'planned',
    size: 40,
    guild_timezone: CHICAGO,
    notes_html: '',
    counts: { in: 21, late: 2, bench: 1, out: 3, tanks: 3, healers: 7, dps: 13 },
    signups: [],
    not_signed: [],
    viewer: { characters: [viewerCharacter()] },
    discord: { connected: false, post_default: false },
    ...extra,
  };
}

export function text(markup: string): string {
  return markup
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
}

/** The markup inside the element carrying `testid`, up to its sibling marker `until`. */
export function between(markup: string, testid: string, until: string): string {
  const marker = markup.indexOf(`data-testid="${testid}"`);
  if (marker < 0) throw new Error(`no ${testid}`);
  const from = markup.indexOf('>', marker) + 1;
  const to = markup.indexOf(`data-testid="${until}"`, from);
  return markup.slice(from, to < 0 ? undefined : markup.lastIndexOf('<', to));
}
