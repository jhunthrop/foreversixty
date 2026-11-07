// web/src/lib/home/feed.ts
// "Forever right now" (design/specs/2026-11-07-home-b.md §1, §2): the GET /v1/feed client, the
// filters, the tags, the relative times and the floor of real dated site events that fills a
// quiet filter. Types follow api/openapi.yaml (/v1/feed); written by hand like every web client.
import { requestEnvelope } from '../account/api';
import { query } from '../data/query';
import { API_BASE_URL } from '../planner/config';

export type FeedFilter = 'all' | 'spec' | 'guild' | 'logs';
/** The API's tags (one per event kind, api/internal/activity KindTag) plus `guides`, which only
 *  the floor's guide-update rows carry. */
export type FeedTag = 'your_spec' | 'logs' | 'rankings' | 'best_in_slot' | 'guides' | 'your_guild' | 'patch';
export type FeedKind =
  | 'log_upload'
  | 'boss_kill'
  | 'guild_raid'
  | 'ranking_top'
  | 'ranking_climb'
  | 'bis_change'
  | 'guild_sim'
  | 'site';

export interface FeedSegment {
  text: string;
  /** Absent: plain text. `name` carries `class`; `item` carries `quality` and `item_id`. */
  style?: 'strong' | 'name' | 'item';
  class?: string;
  quality?: number;
  item_id?: number;
}

export interface FeedRow {
  id: string;
  kind: FeedKind;
  at: string;
  tag: FeedTag;
  /** An icon file stem; drawn when there is no crest. */
  icon: string | null;
  /** A class slug: the row is about a character and leads with its crest. */
  crest: string | null;
  quality: number | null;
  href: string;
  /** A dated event shown with its date (the raid opening), not a relative time. */
  date_only: boolean;
  segments: FeedSegment[];
  /** Floor rows only: the spec the event is about, for the Your spec filter. */
  spec?: string;
}

export interface FeedPage {
  rows: FeedRow[];
  next_cursor: string | null;
}

export interface FeedQuery {
  filter: FeedFilter;
  spec?: string;
  band?: number;
  signedIn: boolean;
}

/** Rows on home (ruling P9). */
export const FEED_ROWS = 7;
/** A filter with fewer rows than this fills from the floor (ruling F1). */
export const FLOOR_MIN = 5;
const FEED_TTL_MS = 60_000;
const MAX_LEVEL = 60;
const FIRST_BAND = 20;
const BAND_STEP = 10;
const MINUTE_MS = 60_000;
const HOUR_MINUTES = 60;
const DAY_HOURS = 24;

export const feedCopy = {
  title: 'Forever right now',
  filters: { all: 'All', spec: 'Your spec', guild: 'Your guild', logs: 'Logs' } satisfies Record<
    FeedFilter,
    string
  >,
  filterLabel: 'Show',
  tags: {
    your_spec: 'Your spec',
    logs: 'Logs',
    rankings: 'Rankings',
    best_in_slot: 'Best in slot',
    guides: 'Guides',
    your_guild: 'Your guild',
    patch: 'Patch',
  } satisfies Record<FeedTag, string>,
  failed: 'We could not load what is happening on Forever.',
  guildEmpty: 'Your guild has nothing posted yet.',
  logsEmpty: 'No public logs in the last 7 days.',
  uploadLog: 'Upload a log',
} as const;

/** The feed filters a viewer gets: signed out none (All only, no bar); Your guild only in a guild. */
export function feedFilters(signedIn: boolean, inGuild: boolean): FeedFilter[] {
  if (!signedIn) return [];
  return inGuild ? ['all', 'spec', 'guild', 'logs'] : ['all', 'spec', 'logs'];
}

/** The level band a character reads against: 20 to 60 in tens. */
export function bandOf(level: number | undefined): number {
  if (level === undefined) return FIRST_BAND;
  if (level >= MAX_LEVEL) return MAX_LEVEL;
  return Math.max(FIRST_BAND, Math.floor(level / BAND_STEP) * BAND_STEP);
}

export function feedPath(q: Omit<FeedQuery, 'signedIn'>, limit = FEED_ROWS): string {
  const params = new URLSearchParams({ filter: q.filter, limit: String(limit) });
  if (q.spec !== undefined) params.set('spec', q.spec);
  if (q.band !== undefined) params.set('band', String(q.band));
  return `/v1/feed?${params.toString()}`;
}

/** One feed page, cached a minute: public signed out, private signed in (the guild rows). */
export function fetchFeed(q: FeedQuery, apiBase: string = API_BASE_URL): Promise<FeedPage> {
  const path = feedPath(q);
  return query<FeedPage>(
    `${apiBase}${path}`,
    async () => {
      const { data } = await requestEnvelope<FeedPage>(path, apiBase, {
        credentials: q.signedIn ? 'include' : 'omit',
      });
      return data ?? { rows: [], next_cursor: null };
    },
    { scope: q.signedIn ? 'private' : 'public', ttlMs: FEED_TTL_MS },
  );
}

/** `just now`, `4 min ago`, `2 h ago`, `3 d ago`. */
export function agoText(iso: string, now: Date): string {
  const minutes = Math.max(0, Math.floor((now.getTime() - new Date(iso).getTime()) / MINUTE_MS));
  if (minutes < 1) return 'just now';
  if (minutes < HOUR_MINUTES) return `${minutes} min ago`;
  const hours = Math.floor(minutes / HOUR_MINUTES);
  if (hours < DAY_HOURS) return `${hours} h ago`;
  return `${Math.floor(hours / DAY_HOURS)} d ago`;
}

/** `Wed 9 Dec`, the calendar day of a date-only event. */
export function dayText(iso: string): string {
  const date = new Date(iso);
  const day = date.toLocaleDateString('en-GB', { weekday: 'short', timeZone: 'UTC' });
  const rest = date.toLocaleDateString('en-GB', { day: 'numeric', month: 'short', timeZone: 'UTC' });
  return `${day} ${rest}`;
}

export function whenText(row: Pick<FeedRow, 'at' | 'date_only'>, now: Date): string {
  return row.date_only ? dayText(row.at) : agoText(row.at, now);
}

/** The rows a filter shows: the API's, then (below FLOOR_MIN) real site events from the floor.
 *  Your guild never floors (it says it has nothing yet); Logs never floors (site events are not
 *  logs). Your spec floors with the viewer's spec only. */
export function withFloor(
  rows: readonly FeedRow[],
  floor: readonly FeedRow[],
  filter: FeedFilter,
  spec: string | undefined,
): FeedRow[] {
  const shown = rows.slice(0, FEED_ROWS);
  if (shown.length >= FLOOR_MIN || filter === 'guild' || filter === 'logs') return shown;
  const pool = filter === 'spec' ? floor.filter((row) => row.spec !== undefined && row.spec === spec) : floor;
  const fill = pool
    .slice(0, FLOOR_MIN - shown.length)
    .map((row) => (spec !== undefined && row.spec === spec ? { ...row, tag: 'your_spec' as const } : row));
  return [...shown, ...fill];
}
