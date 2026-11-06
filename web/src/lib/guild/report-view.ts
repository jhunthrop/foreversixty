// web/src/lib/guild/report-view.ts
// The raid page's report URL, for the copied Discord text.

/** The raid page's absolute URL with the report anchor, for the copied text. */
export function reportUrlOf(location: Pick<Location, 'origin' | 'pathname'>): string {
  return `${location.origin}${location.pathname}#report`;
}
