// web/src/lib/coach/dates.ts
// "Reviewed 28 Oct": the review's timestamp as day and short month in the viewer's zone; and a
// UTC calendar day (a cap lifting) in the same form.

const DAY_MONTH: Intl.DateTimeFormatOptions = { day: 'numeric', month: 'short' };

/** `d MMM`, e.g. `28 Oct`; the raw text when the timestamp does not parse. */
export function reviewedDate(iso: string): string {
  const date = new Date(iso);
  if (Number.isNaN(date.getTime())) return iso;
  return date.toLocaleDateString('en-GB', DAY_MONTH);
}

/** A `YYYY-MM-DD` UTC day as `d MMM`, e.g. `1 Nov`, in every zone; the raw text when it does not parse. */
export function utcDayDate(day: string): string {
  const date = new Date(`${day}T00:00:00Z`);
  if (!/^\d{4}-\d{2}-\d{2}$/.test(day) || Number.isNaN(date.getTime())) return day;
  return date.toLocaleDateString('en-GB', { ...DAY_MONTH, timeZone: 'UTC' });
}
