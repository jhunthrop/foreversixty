// web/src/lib/reports/hero.ts
// Logs landing spec (2026-10-04) §4.A's hero selection rule, pulled out of LogsHero.svelte
// so it is testable without mounting a component (the same reason tree-sizes.ts exists as
// its own module for the report page's own At-pull links): signed in with a report of your
// own -> that report, newest first; every other case -> the canonical sample
// (data/sample-report.json). `MyReport` (account/api.ts) and `ReportMeta` (report/types.ts)
// name the same five facts under different shapes (a report row from the list endpoint vs.
// the single-report endpoint the sample is fetched through), so both map onto one
// `HeroFacts` the component renders identically either way.
import type { MyReport } from '../account/api';
import type { ReportMeta } from '../report/types';

/** The ISO day a report/created_at renders as, everywhere a report shows one: the hero's
 *  own facts line and `ReportRow.svelte`'s identical fragment both call this, so the two
 *  can never drift on format. */
export function reportDay(iso: string): string {
  return iso.slice(0, 10);
}

export interface HeroFacts {
  id: string;
  title: string;
  zone: string;
  createdAt: string;
  fightCount: number;
  killCount: number;
}

/** §4.A, exact: `listMyReports(1).rows[0]` when it has a row, never any other row and
 *  never a blank hero -- every other case falls through to the sample at the call site. */
export function selectOwnHero(rows: readonly MyReport[]): MyReport | null {
  return rows.length > 0 ? rows[0] : null;
}

/** `MyReports.svelte`'s own fallback, reused so the hero and the list below never disagree:
 *  the zone when the report carries no title of its own. */
export function heroTitle(title: string, zone: string): string {
  return title === '' ? zone : title;
}

export function heroFactsFromMyReport(report: MyReport): HeroFacts {
  return {
    id: report.id,
    title: report.title,
    zone: report.zone,
    createdAt: report.created_at,
    fightCount: report.fight_count,
    killCount: report.kill_count,
  };
}

/** `ReportMeta` (the single-report endpoint the sample hero and the per-player hook both
 *  read) carries the whole `fights[]` array rather than precomputed counts -- the same two
 *  counts `ReportRow.svelte`'s own fragment shows, derived here once rather than at every
 *  call site. */
export function heroFactsFromReportMeta(meta: ReportMeta): HeroFacts {
  const kills = meta.fights.filter((fight) => fight.kind === 'encounter' && fight.kill);
  return {
    id: meta.id,
    title: meta.title,
    zone: meta.zone,
    createdAt: meta.created_at,
    fightCount: meta.fights.length,
    killCount: kills.length,
  };
}
