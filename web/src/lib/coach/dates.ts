// web/src/lib/coach/dates.ts
// "Reviewed 28 Oct": the review's timestamp as day and short month in the viewer's zone; and a
// UTC calendar day (a cap lifting) in the same form. Both read through lib/dates.ts.
import { dayMonth, utcDayMonth } from '../dates';

/** `d MMM`, e.g. `28 Oct`; the raw text when the timestamp does not parse. */
export function reviewedDate(iso: string): string {
  return dayMonth(new Date(iso)) || iso;
}

/** A `YYYY-MM-DD` UTC day as `d MMM`, e.g. `1 Nov`, in every zone; the raw text when it does not parse. */
export function utcDayDate(day: string): string {
  if (!/^\d{4}-\d{2}-\d{2}$/.test(day)) return day;
  return utcDayMonth(day) || day;
}
