// web/src/lib/changelog/entries.ts
// The changelog's model (pages spec 2026-11-08 section 4.1 and rulings 0A): the collection sorted
// newest first (ties in collection order), each entry with its kind's tag, its sources lane (a
// source tagged only when its kind differs from the entry's; this site's own pages as action
// links, ruling Q5), grouped by month with the count each Show filter leaves. Entries before
// launch fold to title rows that open in place (ruling CL1). Pure: the page passes the
// collection and each entry's rendered body.
import type { Source, SourceKind } from '../sources';
import { pillLabelFor } from '../sources';
import { formatDate, monthHeading } from '../dates';
import { changelogCopy } from './copy';

export const SHOWS = ['all', 'game', 'site'] as const;
export type ChangelogShow = (typeof SHOWS)[number];
/** The filter's query parameter (`?show=game`); absent is All. */
export const SHOW_PARAM = 'show';
/** Launch day: entries dated before it fold to title rows (ruling CL1). */
export const LAUNCH_DAY = '2026-11-04';
/** The site's own origin: a source here is one of its pages, linked as an action. */
const SITE_ORIGIN = 'https://foreversixty.gg';

export interface ChangelogInput {
  id: string;
  data: { title: string; date: Date; kind: SourceKind; note?: string; sources: Source[] };
  /** The rendered markdown body ('' when the file has none). */
  html: string;
}

export interface LaneSource {
  label: string;
  url: string;
  /** The source's kind, shown only when it differs from the entry's. */
  tag?: string;
}

export interface LaneAction {
  label: string;
  href: string;
}

export interface ChangelogEntry {
  id: string;
  iso: string;
  date: string;
  title: string;
  kindLabel: string;
  show: Exclude<ChangelogShow, 'all'>;
  /** Before launch: a title row that opens in place. */
  folded: boolean;
  /** The body: the rendered markdown, or (no body) the note as one paragraph's text. */
  html: string;
  noteAsBody?: string;
  sources: LaneSource[];
  actions: LaneAction[];
  /** The note, beside the sources, when the entry has a body of its own. */
  caveat?: string;
}

export interface ChangelogMonth {
  title: string;
  anchor: string;
  entries: ChangelogEntry[];
  counts: Record<ChangelogShow, number>;
}

/** Which side of the filter a kind is on: Game is Blizzard, community and datamined news. */
export function showOf(kind: SourceKind): Exclude<ChangelogShow, 'all'> {
  return kind === 'site' ? 'site' : 'game';
}

/** The Show value a query names, or All. */
export function showFrom(search: string): ChangelogShow {
  const value = new URLSearchParams(search).get(SHOW_PARAM);
  return (SHOWS as readonly string[]).includes(value ?? '') ? (value as ChangelogShow) : 'all';
}

/** A site source as the action that opens its page (`Open the guild page ›`). */
export function siteAction(url: string): LaneAction {
  const path = url.startsWith(SITE_ORIGIN) ? url.slice(SITE_ORIGIN.length) || '/' : url;
  const label = changelogCopy.siteActions[path];
  if (label === undefined)
    throw new Error(`changelog: no action label for the site page ${path} (lib/changelog/copy.ts)`);
  return { label: `${label} ›`, href: path };
}

function laneOf(kind: SourceKind, sources: readonly Source[]): Pick<ChangelogEntry, 'sources' | 'actions'> {
  const seen = new Set<string>();
  const unique = sources.filter((source) => !seen.has(source.url) && seen.add(source.url));
  return {
    actions: unique.filter((source) => source.kind === 'site').map((source) => siteAction(source.url)),
    sources: unique
      .filter((source) => source.kind !== 'site')
      .map((source) => ({
        label: source.label,
        url: source.url,
        ...(source.kind === kind ? {} : { tag: pillLabelFor(source.kind) }),
      })),
  };
}

function entryOf(input: ChangelogInput): ChangelogEntry {
  const { data } = input;
  const iso = data.date.toISOString().slice(0, 10);
  const body = input.html.trim();
  return {
    id: input.id,
    iso,
    date: formatDate(data.date),
    title: data.title,
    kindLabel: pillLabelFor(data.kind),
    show: showOf(data.kind),
    folded: iso < LAUNCH_DAY,
    html: body,
    ...(body === '' && data.note !== undefined ? { noteAsBody: data.note } : {}),
    ...(body !== '' && data.note !== undefined ? { caveat: data.note } : {}),
    ...laneOf(data.kind, data.sources),
  };
}

/** Newest first; entries of one day keep their collection order (a stable sort). */
export function sortedEntries(inputs: readonly ChangelogInput[]): ChangelogEntry[] {
  return [...inputs].sort((a, b) => b.data.date.getTime() - a.data.date.getTime()).map(entryOf);
}

/** How many entries each Show value leaves. */
export function showCounts(entries: readonly Pick<ChangelogEntry, 'show'>[]): Record<ChangelogShow, number> {
  const game = entries.filter((entry) => entry.show === 'game').length;
  return { all: entries.length, game, site: entries.length - game };
}

/** The entries by month, newest first, each with its counts under every filter. */
export function changelogMonths(entries: readonly ChangelogEntry[]): ChangelogMonth[] {
  const months: ChangelogMonth[] = [];
  for (const entry of entries) {
    const anchor = entry.iso.slice(0, 7);
    const last = months.at(-1);
    if (last?.anchor === anchor) last.entries.push(entry);
    else
      months.push({
        title: monthHeading(new Date(entry.iso)),
        anchor,
        entries: [entry],
        counts: showCounts([]),
      });
  }
  return months.map((month) => ({ ...month, counts: showCounts(month.entries) }));
}
