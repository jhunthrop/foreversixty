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
