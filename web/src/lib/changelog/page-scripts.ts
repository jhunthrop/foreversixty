// web/src/lib/changelog/page-scripts.ts
// The changelog's small scripts (pages spec 2026-11-08 §4.1, §5.2, §9). Inline, so nothing moves
// after paint: the head reads `?show=` into `html[data-cl-show]` (CSS hides the other side's
// entries, empty months and their contents links); right after the bar the checked segment and
// the phone line follow it; after the list a returning visitor's entries newer than the last
// visit get `New` (localStorage, every access in try/catch: a blocked store is a first visit).
// `bootChangelog` (a module, after load) keeps all of it in step with the bar and opens a folded
// entry a link points at.
import { CONTROL_BAR_EVENT, type ControlBarChange } from '../ui/control-bar';
import { changelogCopy } from './copy';
import { SHOW_PARAM, SHOWS, showFrom, type ChangelogShow } from './entries';

/** The newest entry date a visitor has seen (`2026-11-05`). */
export const SEEN_KEY = 'fs.changelogSeen';
/** The `html` attribute the filter CSS keys off. */
export const SHOW_ATTR = 'data-cl-show';

/** Head, before paint: `?show=` into `html[data-cl-show]` (All when absent or unknown). */
export const SHOW_HEAD_SCRIPT = `(function(){var v=new URLSearchParams(location.search).get(${JSON.stringify(SHOW_PARAM)});document.documentElement.setAttribute(${JSON.stringify(SHOW_ATTR)},${JSON.stringify(SHOWS)}.indexOf(v)>=0?v:'all');})();`;

/** The phone line for each Show value (`Game changes · 25`). */
export function summaries(counts: Record<ChangelogShow, number>): Record<ChangelogShow, string> {
  return Object.fromEntries(SHOWS.map((show) => [show, changelogCopy.summary(show, counts[show])])) as Record<
    ChangelogShow,
    string
  >;
}

/** Marks the segment `show` of the group `param` checked and writes its phone line, in the bar
 *  `bar`. Self-contained (it is also serialised into the inline script): no outer names. */
export function syncBar(
  bar: ParentNode,
  param: string,
  show: string,
  lines: Readonly<Record<string, string>>,
): void {
  bar.querySelectorAll<HTMLElement>(`[data-param="${param}"] [role="radio"]`).forEach((radio) => {
    const on = radio.dataset.value === show;
    radio.setAttribute('aria-checked', String(on));
    radio.tabIndex = on ? 0 : -1;
  });
  const line = bar.querySelector('[data-control-bar-toggle] span');
  if (line !== null && lines[show] !== undefined) line.textContent = lines[show];
}

/** After the bar, before the list paints: the bar follows `html[data-cl-show]`. */
export function barSyncScript(counts: Record<ChangelogShow, number>): string {
  return `(function(){var s=document.documentElement.getAttribute(${JSON.stringify(SHOW_ATTR)})||'all';if(s==='all')return;var b=document.querySelector('[data-testid="changelog-bar"]');if(!b)return;(${syncBar.toString()})(b,${JSON.stringify(SHOW_PARAM)},s,${JSON.stringify(summaries(counts))});})();`;
}

/** After the list: `New` on entries newer than the last visit; the newest is stored on leaving. */
export function seenScript(newest: string): string {
  return `(function(){var k=${JSON.stringify(SEEN_KEY)};try{var seen=localStorage.getItem(k);if(seen){document.querySelectorAll('[data-show]').forEach(function(e){var d=e.querySelector('time');if(d&&d.getAttribute('datetime')>seen){var t=e.querySelector('[data-cl-new]');if(t)t.hidden=false;}});}}catch(e){}addEventListener('pagehide',function(){try{localStorage.setItem(k,${JSON.stringify(newest)});}catch(e){}});})();`;
}

/** Opens the folded entry the address points at, so a permalink to a pre-launch entry reads it. */
export function openTarget(hash: string): void {
  const id = decodeURIComponent(hash.replace(/^#/, ''));
  const target = id === '' ? null : document.getElementById(id);
  if (target instanceof HTMLDetailsElement && !target.open) {
    target.open = true;
    target.scrollIntoView();
  }
}

/** After load: the bar's Show choice drives the filter; a hash opens its folded entry. */
export function bootChangelog(counts: Record<ChangelogShow, number>): void {
  const lines = summaries(counts);
  const root = document.documentElement;
  const select = document.getElementById('changelog-toc') as HTMLSelectElement | null;
  const apply = (show: ChangelogShow): void => {
    root.setAttribute(SHOW_ATTR, show);
    const bar = document.querySelector<HTMLElement>('[data-testid="changelog-bar"]');
    if (bar !== null) syncBar(bar, SHOW_PARAM, show, lines);
    select?.querySelectorAll('option').forEach((option) => {
      const month = option.value === '' ? null : document.getElementById(option.value);
      option.hidden = month !== null && month.offsetParent === null;
    });
  };
  document.addEventListener(CONTROL_BAR_EVENT, (event) => {
    const { param, value } = (event as CustomEvent<ControlBarChange>).detail;
    if (param === SHOW_PARAM) apply(showFrom(`?${SHOW_PARAM}=${value}`));
  });
  apply(showFrom(location.search));
  openTarget(location.hash);
  addEventListener('hashchange', () => openTarget(location.hash));
}
