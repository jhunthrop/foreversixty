// web/src/lib/home/feed.ts
// "Forever right now" (design/specs/2026-11-07-home-b.md §1, §2): the GET /v1/feed client, the
// filters, the tags, the relative times and the floor of real dated site events that fills a
// quiet filter. Types follow api/openapi.yaml (/v1/feed); written by hand like every web client.
import { requestEnvelope } from '../account/api';
import { query } from '../data/query';
import { API_BASE_URL } from '../planner/config';
import {
  sharedFeedFailed,
  sharedFeedLogsEmpty,
  sharedFeedTagBestInSlot,
  sharedFeedTagGuides,
  sharedFeedTagLogs,
  sharedFeedTagPatch,
  sharedFeedTagRankings,
  sharedFeedTagYourGuild,
  sharedFeedTagYourSpec,
  sharedFeedTitle,
  sharedFeedUploadLog,
} from '../copy/shared/feed.generated';
import { compactAgo, weekdayDayMonth } from '../dates';

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
  /** Absent: plain text. `name` carries `class`; `item` carries `quality` and `item_id`. `gain` and
   *  `loss` are home's own (a sim's change since the one before, lib/home/feed-mix.ts). */
  style?: 'strong' | 'name' | 'item' | 'gain' | 'loss';
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

/** What a feed row leads with (ui/FeedRow): the crest, else the item, else the spell icon. */
export type FeedLead = Pick<FeedRow, 'crest' | 'icon' | 'quality'>;

export interface FeedPage {
  rows: FeedRow[];
  next_cursor: string | null;
}

export interface FeedQuery {
  filter: FeedFilter;
  spec?: string;
  band?: number;
  signedIn: boolean;
  /** Rows asked for; FEED_ROWS by default. */
  limit?: number;
}

/** Rows on home (ruling P9). */
export const FEED_ROWS = 7;
/** A filter with fewer rows than this fills from the floor (ruling F1). */
export const FLOOR_MIN = 5;
const FEED_TTL_MS = 60_000;
const MAX_LEVEL = 60;
const FIRST_BAND = 20;
const BAND_STEP = 10;

export const feedCopy = {
  title: sharedFeedTitle,
  tags: {
    your_spec: sharedFeedTagYourSpec,
    logs: sharedFeedTagLogs,
    rankings: sharedFeedTagRankings,
    best_in_slot: sharedFeedTagBestInSlot,
    guides: sharedFeedTagGuides,
    your_guild: sharedFeedTagYourGuild,
    patch: sharedFeedTagPatch,
  } satisfies Record<FeedTag, string>,
  failed: sharedFeedFailed,
  logsEmpty: sharedFeedLogsEmpty,
  uploadLog: sharedFeedUploadLog,
} as const;

/** The level band a character reads against: 20 to 60 in tens. */
export function bandOf(level: number | undefined): number {
  if (level === undefined) return FIRST_BAND;
  if (level >= MAX_LEVEL) return MAX_LEVEL;
  return Math.max(FIRST_BAND, Math.floor(level / BAND_STEP) * BAND_STEP);
}

export function feedPath(q: Omit<FeedQuery, 'signedIn'>): string {
  const params = new URLSearchParams({ filter: q.filter, limit: String(q.limit ?? FEED_ROWS) });
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

/** `Wed 9 Dec`, the calendar day of a date-only event. */
export function dayText(iso: string, now: Date = new Date()): string {
  return weekdayDayMonth(new Date(iso), 'UTC', now);
}

export function whenText(row: Pick<FeedRow, 'at' | 'date_only'>, now: Date): string {
  return row.date_only ? dayText(row.at, now) : compactAgo(new Date(row.at), now);
}

/** A guild sim's tail the guild's own home already says: the guild, spec and band (ai-slop tell 22). */
const GUILD_SIM_TAIL = /, #(\d+) of (\d+) DPS in [^,(]+ \([^)]*\)$/;

/** The rows as home shows them: newest first, a guild sim reads `#2 of 14`, and with the guild's
 *  raid on the plaque its post is left out (the plaque already shows it). */
export function homeRows(rows: readonly FeedRow[], raidOnPlaque: boolean): FeedRow[] {
  return [...rows]
    .sort((a, b) => Date.parse(b.at) - Date.parse(a.at))
    .filter((row) => !(raidOnPlaque && row.kind === 'guild_raid'))
    .map((row) =>
      row.kind !== 'guild_sim'
        ? row
        : {
            ...row,
            segments: row.segments.map((seg) =>
              seg.style === undefined
                ? { ...seg, text: seg.text.replace(GUILD_SIM_TAIL, ', #$1 of $2') }
                : seg,
            ),
          },
    );
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
