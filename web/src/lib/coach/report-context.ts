// web/src/lib/coach/report-context.ts
// What a row strip needs to know about the report around it, without ActorRow growing props:
// ReportView provides one getter (`provideCoachRow`) and each CoachRowLine reads it. The getter
// keeps the values live as the visitor changes fight. Trash and pulls under MinFightMS get no
// strip (spec 8: the report knows both before any fetch).
import { getContext, setContext } from 'svelte';
import type { FightEntry } from '../report/types';
import { isSampleReport } from '../report/sample';

const KEY = Symbol('coach-row');
/** Shorter pulls are not reviewed (the API's MinFightMS). */
export const MIN_REVIEW_FIGHT_MS = 30_000;

export interface CoachRowReport {
  reportId: string;
  /** The selected fight; null over the whole night. */
  fight: FightEntry | null;
  firstFight: number;
}

export type CoachRowReader = () => CoachRowReport;

export function provideCoachRow(read: CoachRowReader): void {
  setContext(KEY, read);
}

export function useCoachRow(): CoachRowReader | undefined {
  return getContext<CoachRowReader | undefined>(KEY);
}

/** A boss pull long enough to review, on a report the API knows. */
export function reviewableFight(report: CoachRowReport): FightEntry | null {
  const { fight } = report;
  if (fight === null || isSampleReport(report.reportId)) return null;
  return fight.kind === 'encounter' && fight.duration_ms >= MIN_REVIEW_FIGHT_MS ? fight : null;
}
