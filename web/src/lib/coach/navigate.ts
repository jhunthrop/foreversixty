// web/src/lib/coach/navigate.ts
// Opens a report link (`?fight=3&tab=coach&source=GUID`) without a page load: push the url and
// raise `popstate`, which ReportView already listens to (it re-reads the query string). A click
// with a modifier, or a non-primary button, keeps the browser's own behaviour.
import { reportSearch, type ReportState } from '../report/url';

export function reportHref(state: ReportState, firstFight: number): string {
  return reportSearch(state, firstFight) || '?';
}

export function openReportLink(href: string): void {
  window.history.pushState(null, '', `${window.location.pathname}${href}`);
  window.dispatchEvent(new PopStateEvent('popstate'));
}

export function followReportLink(event: MouseEvent, href: string): void {
  if (event.defaultPrevented || event.button !== 0) return;
  if (event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return;
  event.preventDefault();
  openReportLink(href);
}
