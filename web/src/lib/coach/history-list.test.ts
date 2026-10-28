import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { coachCopy } from './copy';
import {
  HISTORY_LOADING,
  historyHref,
  historyRow,
  historyShown,
  withFailure,
  withOlderLoading,
  withPage,
} from './history-list';
import { normalizeHistory } from './normalize';
import type { CoachHistoryItem, CoachHistoryPage } from './types';

const page = (name: string): CoachHistoryPage =>
  normalizeHistory(
    JSON.parse(readFileSync(new URL(`../../fixtures/coach/${name}.json`, import.meta.url), 'utf8')),
  );
const FIRST = page('history-first');
const OLDER = page('history-older');

const item = (over: Partial<CoachHistoryItem> = {}): CoachHistoryItem => ({
  report_id: 'abc',
  fight_index: 3,
  player_key: 'us/hardcore/baelgrim',
  spec: 'Fury',
  boss: 'Onyxia',
  kill: true,
  source: 'model',
  generated_at: '2026-10-25T12:00:00Z',
  fix_first: { title: 'Bloodthirst: 9 casts lost' },
  ...over,
});

describe('historyRow', () => {
  it('reads {date} · {Boss} · {Spec} · {title} with no tags on a model-read kill', () => {
    expect(historyRow(item())).toEqual({
      key: 'abc|3',
      lead: ['25 Oct', 'Onyxia', 'Fury'],
      title: 'Bloodthirst: 9 casts lost',
      wipe: false,
      logOnly: false,
      href: '/reports/abc?fight=3&tab=coach',
    });
  });

  it('a review with no first title reads Nothing to fix', () => {
    expect(historyRow(item({ fix_first: null })).title).toBe(coachCopy.historyNone);
  });

  it('tags a wipe and a rule-only review', () => {
    expect(historyRow(item({ kill: false, source: 'rules' }))).toMatchObject({ wipe: true, logOnly: true });
  });

  it('drops an empty boss or spec rather than printing an empty segment', () => {
    expect(historyRow(item({ boss: '', spec: '' })).lead).toEqual(['25 Oct']);
  });
});

describe('historyHref', () => {
  it('encodes the report id and opens the Coach tab on the fight', () => {
    expect(historyHref({ report_id: 'a/b c', fight_index: 12 })).toBe(
      '/reports/a%2Fb%20c?fight=12&tab=coach',
    );
  });
});

describe('the history list state', () => {
  it('is hidden while the first page is read and with no history', () => {
    expect(historyShown(HISTORY_LOADING)).toBe(false);
    expect(historyShown(withPage(HISTORY_LOADING, { reviews: [], next_cursor: null }))).toBe(false);
  });

  it('the first page holds twenty rows and the cursor for Show older', () => {
    const state = withPage(HISTORY_LOADING, FIRST);
    expect(state.items).toHaveLength(20);
    expect(state.cursor).toBe(FIRST.next_cursor);
    expect(historyShown(state)).toBe(true);
  });

  it('Show older appends the next page and ends at the last', () => {
    const loading = withOlderLoading(withPage(HISTORY_LOADING, FIRST));
    expect(loading.olderLoading).toBe(true);
    const state = withPage(loading, OLDER);
    expect(state.items).toHaveLength(FIRST.reviews.length + OLDER.reviews.length);
    expect(state.cursor).toBeNull();
    expect(state.olderLoading).toBe(false);
  });

  it('a repeated row across pages is listed once', () => {
    const state = withPage(withPage(HISTORY_LOADING, FIRST), FIRST);
    expect(state.items).toHaveLength(20);
  });

  it('a failed first page is the error state; a failed older page keeps the rows', () => {
    const failedFirst = withFailure(HISTORY_LOADING);
    expect(failedFirst.status).toBe('error');
    expect(historyShown(failedFirst)).toBe(true);
    const ready = withPage(HISTORY_LOADING, FIRST);
    const failedOlder = withFailure(withOlderLoading(ready));
    expect(failedOlder).toMatchObject({ status: 'ready', olderLoading: false, olderFailed: true });
    expect(failedOlder.items).toHaveLength(20);
    expect(withOlderLoading(failedOlder).olderFailed).toBe(false);
  });
});
