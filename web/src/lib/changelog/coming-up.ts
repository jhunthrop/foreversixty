// web/src/lib/changelog/coming-up.ts
// The changelog's Coming up block (pages spec 2026-11-08 ruling Q7): /roadmap folded into the
// changelog, and its dates still ahead of launch moved here verbatim from roadmap.md's table
// (Blizzard's What's Next panel recap). A row drops itself at build time (the site rebuilds
// nightly) once its day has passed or a changelog entry names it in `covers`. Give a row an `on`
// day only when Blizzard dates one; a season stays until an entry covers it.
import type { Source } from '../sources';
import { escapeHtml } from '../html';

export interface ComingUpRow {
  /** The row's key, the value a changelog entry's `covers` names. */
  id: string;
  /** The UTC day Blizzard dates it (`2026-12-09`); the row shows through that day. */
  on?: string;
  /** As Blizzard dates it: a day ("9 Dec", the house date) or a season ("Spring 2027"). */
  when: string;
  what: string;
}

/** The table's two heads, as roadmap.md had them. */
export const COMING_UP_HEADS = ['Date', 'Content'] as const;

export const COMING_UP = [
  {
    id: 'raids-dec-9',
    on: '2026-12-09',
    when: '9 Dec',
    what: "Barrow Deeps, Hyjal Summit, and Onyxia's Lair open",
  },
  { id: 'hardcore', when: 'Winter 2026/27', what: 'The Hardcore ruleset, exact date not yet announced' },
  {
    id: 'spring-2027',
    when: 'Spring 2027',
    what: 'Two raids, two dungeons, new quests and a zone, a legendary questline, and a PvP season',
  },
  {
    id: 'summer-2027',
    when: 'Summer 2027',
    what: 'A revamped iconic raid plus one new raid, two dungeons, a PvP season, and professions and Legacy updates',
  },
] as const satisfies readonly ComingUpRow[];

/** The keys a changelog entry's `covers` may name. */
export type ComingUpId = (typeof COMING_UP)[number]['id'];
export const COMING_UP_IDS = COMING_UP.map((row) => row.id) as [ComingUpId, ...ComingUpId[]];

/** The UTC day of an instant, as `YYYY-MM-DD`. */
function utcDay(now: Date): string {
  return now.toISOString().slice(0, 10);
}

/** The rows still ahead on `now`: their day not yet passed and no changelog entry covering them. */
export function comingUpRows(
  now: Date,
  covered: Iterable<string> = [],
  rows: readonly ComingUpRow[] = COMING_UP,
): ComingUpRow[] {
  const today = utcDay(now);
  const done = new Set(covered);
  return rows.filter((row) => !done.has(row.id) && (row.on === undefined || row.on >= today));
}

export const COMING_UP_SOURCE: Source = {
  label: "Blizzard, What's Next panel recap",
  url: 'https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap',
  kind: 'blizzard',
};

/** The rows as the ruled table `Prose` draws (roadmap.md's own table, its rows ahead only). */
export function comingUpHtml(rows: readonly ComingUpRow[]): string {
  const head = COMING_UP_HEADS.map((cell) => `<th>${escapeHtml(cell)}</th>`).join('');
  const body = rows
    .map((row) => `<tr><td>${escapeHtml(row.when)}</td><td>${escapeHtml(row.what)}</td></tr>`)
    .join('');
  return `<table><thead><tr>${head}</tr></thead><tbody>${body}</tbody></table>`;
}
