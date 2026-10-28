// web/src/lib/coach/history-list.ts
// The character page's coach history (experience spec 3.5 and 10.3): what one row says and where
// it goes, and how the pages of the history route add up. Pure, so the list and its tests share
// one decision; CoachHistoryList.svelte owns the fetching.
import { coachCopy } from './copy';
import { reviewedDate } from './dates';
import type { CoachHistoryItem, CoachHistoryPage } from './types';

export interface HistoryRow {
  /** Unique within the list: one stored review per (report, fight). */
  key: string;
  /** `{date} · {Boss} · {Spec}`, the muted lead, in order. */
  lead: readonly string[];
  /** `fix_first.title`, or `Nothing to fix` when the review has none. */
  title: string;
  /** The muted `Wipe` tag. */
  wipe: boolean;
  /** The muted `Log only` tag: the review is the rule-only rendering. */
  logOnly: boolean;
  /** The report with the Coach tab open on that fight. */
  href: string;
}

/** One stored review per (report, fight): the row's key and the de-duplication of pages. */
const itemKey = (item: Pick<CoachHistoryItem, 'report_id' | 'fight_index'>): string =>
  `${item.report_id}|${item.fight_index}`;

/**
 * The report page with the Coach tab open on the fight. The history carries no unit GUID (a pack
 * holds none, by design), so the link leaves `source` to the tab, which picks the viewer's own
 * character in the pull, as the Discord and account links do.
 */
export function historyHref(item: Pick<CoachHistoryItem, 'report_id' | 'fight_index'>): string {
  const query = new URLSearchParams({ fight: String(item.fight_index), tab: 'coach' });
  return `/reports/${encodeURIComponent(item.report_id)}?${query.toString()}`;
}

export function historyRow(item: CoachHistoryItem): HistoryRow {
  const lead = [reviewedDate(item.generated_at), item.boss, item.spec].filter((part) => part !== '');
  return {
    key: itemKey(item),
    lead,
    title: item.fix_first?.title ?? coachCopy.historyNone,
    wipe: !item.kill,
    logOnly: item.source === 'rules',
    href: historyHref(item),
  };
}

/** The list's state: the rows read so far, the next page's cursor, and what is in flight. */
export interface HistoryListState {
  readonly status: 'loading' | 'ready' | 'error';
  readonly items: readonly CoachHistoryItem[];
  readonly cursor: string | null;
  /** `Show older` is fetching (the button is disabled and `aria-busy`). */
  readonly olderLoading: boolean;
  /** The last `Show older` failed: the `error` text under the list. */
  readonly olderFailed: boolean;
}

export const HISTORY_LOADING: HistoryListState = {
  status: 'loading',
  items: [],
  cursor: null,
  olderLoading: false,
  olderFailed: false,
};

/** A page landed: the first replaces the list, an older one is appended (a repeat is dropped). */
export function withPage(state: HistoryListState, page: CoachHistoryPage): HistoryListState {
  const seen = new Set(state.items.map(itemKey));
  const fresh = page.reviews.filter((item) => !seen.has(itemKey(item)));
  return {
    status: 'ready',
    items: [...state.items, ...fresh],
    cursor: page.next_cursor,
    olderLoading: false,
    olderFailed: false,
  };
}

export function withOlderLoading(state: HistoryListState): HistoryListState {
  return { ...state, olderLoading: true, olderFailed: false };
}

/** A failed read: the first page fails the list, an older page only says so under it. */
export function withFailure(state: HistoryListState): HistoryListState {
  if (state.status !== 'ready') return { ...HISTORY_LOADING, status: 'error' };
  return { ...state, olderLoading: false, olderFailed: true };
}

/** The section is absent with no history (spec 4.3) and while the first page is read. */
export function historyShown(state: HistoryListState): boolean {
  return state.status === 'error' || (state.status === 'ready' && state.items.length > 0);
}
