// web/src/lib/report/heading.ts
// A report's title: the uploader's own, else the zone, else the log's date, so an untitled
// upload of a log with no zone reads "17 Sept 2026 · 42 fights" and not blank, and leads with
// the logging character when the report is credited ("Bow Jackzon · 17 Sept 2026 · 42 fights").
import { formatDate } from './format';
import type { ReportMeta } from './types';

export function reportHeading(
  meta: Pick<ReportMeta, 'title' | 'zone' | 'created_at'> & { fights: readonly { start: string }[] },
  character?: string,
): string {
  if (meta.title !== '') return meta.title;
  if (meta.zone !== '') return meta.zone;
  const date = formatDate(meta.fights[0]?.start ?? meta.created_at) || formatDate(meta.created_at);
  const count = `${meta.fights.length} ${meta.fights.length === 1 ? 'fight' : 'fights'}`;
  const when = date === '' ? count : `${date} · ${count}`;
  return character ? `${character} · ${when}` : when;
}

/**
 * The words the report page's subtitle and the /logs sample row both say about a report: where,
 * when (the first fight's date, the page's format), how many fights, and how many bosses died.
 * One function, so the two surfaces cannot drift apart.
 */
export function reportFacts(
  meta: Pick<ReportMeta, 'zone' | 'created_at'> & {
    fights: readonly { start: string; kind: string; kill: boolean }[];
  },
): { zone: string; date: string; fights: string; bosses: string } {
  const kills = meta.fights.filter((fight) => fight.kind === 'encounter' && fight.kill).length;
  return {
    zone: meta.zone,
    date: formatDate(meta.fights[0]?.start ?? meta.created_at),
    fights: `${meta.fights.length} ${meta.fights.length === 1 ? 'fight' : 'fights'}`,
    bosses: kills === 0 ? '' : `${kills} ${kills === 1 ? 'boss' : 'bosses'}`,
  };
}
