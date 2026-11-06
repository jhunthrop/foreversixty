// web/src/lib/guild/report-attendance-view.ts
// The report rail's Attendance rows (web layout spec 2026-11-05 §7.2 "Raid report", §7.4 "Manual
// attendance fix"): each raider's effective night, and when an officer fixed it, who did and what
// the log said, muted. The rows keep the attendance groups' order (not in the log, late, joined,
// present, bench and out), so the rows an officer is likely to fix come first.
import type { AttendanceOverrideStatus, RaidAttendanceRow } from './attendance-contract';
import { attendanceCopy } from './attendance-copy';
import { groupAttendance, outcomeText } from './night';
import { fixedBy, logSaid, reportNightCopy as copy } from './report-night-copy';
import { SEP, raidCopy } from './raid-copy';
import type { SelectOption } from '../ui/listbox-model';

export const FIX_STATUSES: readonly AttendanceOverrideStatus[] = ['present', 'absent', 'bench', 'excused'];

export interface FixRowModel {
  row: RaidAttendanceRow;
  /** The night every count reads: the fix, else the log's. */
  effective: string;
  /** `The log said Not in the log`, only under a fix. */
  auto: string | null;
  /** The control's value: the fix's status, `''` with none. */
  value: AttendanceOverrideStatus | '';
}

/** What the log and the sign-ups said: `Not in the log`, `Late, first seen 22 min after the first
 *  pull`, `Bench` or `Out` for a raider who stayed away; `· Excused` when the log row was excused. */
export function logOutcomeText(row: RaidAttendanceRow): string {
  const asLogged = { ...row, outcome: row.auto_outcome, excused: row.auto_excused };
  const text =
    outcomeText(asLogged) || (row.signed === 'bench' ? raidCopy.segments.bench : copy.attendance.out);
  return asLogged.excused ? `${text}${SEP}${attendanceCopy.attendance.excused}` : text;
}

export function fixRowModel(row: RaidAttendanceRow): FixRowModel {
  const fix = row.override;
  if (fix === null) return { row, effective: logOutcomeText(row), auto: null, value: '' };
  return {
    row,
    effective: fixedBy(copy.attendance.status[fix.status], fix.by_name),
    auto: logSaid(logOutcomeText(row)),
    value: fix.status,
  };
}

export function fixRows(rows: readonly RaidAttendanceRow[]): FixRowModel[] {
  return groupAttendance(rows).flatMap((group) => group.rows.map(fixRowModel));
}

/** The fix control's options. */
export function fixOptions(): SelectOption[] {
  return FIX_STATUSES.map((status) => ({ value: status, label: copy.attendance.status[status] }));
}
