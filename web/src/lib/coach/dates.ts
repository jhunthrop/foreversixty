// web/src/lib/coach/dates.ts
// "Reviewed 28 Oct": the review's timestamp as day and short month in the viewer's zone.

/** `d MMM`, e.g. `28 Oct`; the raw text when the timestamp does not parse. */
export function reviewedDate(iso: string): string {
  const date = new Date(iso);
  if (Number.isNaN(date.getTime())) return iso;
  return date.toLocaleDateString('en-GB', { day: 'numeric', month: 'short' });
}
