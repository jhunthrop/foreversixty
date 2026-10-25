// web/src/lib/sim/server-runs.test.ts
import { describe, expect, it } from 'vitest';
import {
  MAX_VISIBLE_RUNS,
  ageWords,
  elapsedSeconds,
  estimateWords,
  isOverdue,
  rowState,
  visibleRows,
} from './server-runs';
import type { SimListRow } from './types';

const NOW = new Date('2026-10-24T12:00:00Z');

function row(overrides: Partial<SimListRow> = {}): SimListRow {
  return {
    sim_id: 'abcdefghijkl',
    spec: 'mage-frost',
    dps: 1204,
    engine_version: 'v1',
    created_at: '2026-10-24T11:58:00Z',
    title: '',
    state: 'running',
    ...overrides,
  };
}

describe('estimateWords', () => {
  it.each([
    [30, 'about a minute'],
    [60, 'about a minute'],
    [90, 'about 2 minutes'],
    [170, 'about 3 minutes'],
    [5340, 'about 89 minutes'],
    [5400, 'about 2 hours'],
    [10800, 'about 3 hours'],
  ])('%i seconds reads %s', (seconds, words) => {
    expect(estimateWords(seconds)).toBe(words);
  });
});

describe('isOverdue', () => {
  it('is overdue only past one and a half times the estimate', () => {
    expect(isOverdue(180, 270)).toBe(false);
    expect(isOverdue(180, 271)).toBe(true);
  });
});

describe('elapsedSeconds', () => {
  it('counts whole seconds since creation and never goes negative', () => {
    expect(elapsedSeconds(row(), NOW)).toBe(120);
    expect(elapsedSeconds(row({ created_at: '2026-10-24T12:00:30Z' }), NOW)).toBe(0);
  });
});

describe('rowState', () => {
  it('maps queued and error straight through', () => {
    expect(rowState(row({ state: 'queued' }), NOW)).toBe('queued');
    expect(rowState(row({ state: 'error' }), NOW)).toBe('error');
  });

  it('running with an estimate that is not overdue is running', () => {
    expect(rowState(row({ estimate_seconds: 180 }), NOW)).toBe('running');
  });

  it('running past 1.5 times the estimate is overdue (271 s against 180)', () => {
    const created = '2026-10-24T11:55:29Z';
    expect(rowState(row({ estimate_seconds: 180, created_at: created }), NOW)).toBe('running-overdue');
  });

  it('running with no estimate says so', () => {
    expect(rowState(row(), NOW)).toBe('running-no-estimate');
  });

  it('done is done-unseen only when the API said seen is false', () => {
    expect(rowState(row({ state: 'done', seen: false }), NOW)).toBe('done-unseen');
    expect(rowState(row({ state: 'done', seen: true }), NOW)).toBe('done');
    expect(rowState(row({ state: 'done' }), NOW)).toBe('done');
  });

  it('a row with no state is a finished saved sim', () => {
    expect(rowState(row({ state: undefined }), NOW)).toBe('done');
  });
});

describe('ageWords', () => {
  it('is relative inside 14 days', () => {
    expect(ageWords(new Date('2026-10-24T11:46:00Z'), NOW)).toBe('14 minutes ago');
    expect(ageWords(new Date('2026-10-24T10:00:00Z'), NOW)).toBe('2 hours ago');
    expect(ageWords(new Date('2026-10-21T12:00:00Z'), NOW)).toBe('3 days ago');
    expect(ageWords(new Date('2026-10-24T11:59:50Z'), NOW)).toBe('just now');
  });

  it('is a short date from 14 days on', () => {
    expect(ageWords(new Date('2026-10-10T12:00:00Z'), NOW)).toBe('Oct 10');
    expect(ageWords(new Date('2026-09-25T12:00:00Z'), NOW)).toBe('Sep 25');
  });
});

describe('visibleRows', () => {
  const rows = Array.from({ length: 7 }, (_, i) =>
    row({ sim_id: `sim${i}`, created_at: `2026-10-2${i}T10:00:00Z` }),
  );

  it('keeps the newest five and reports more', () => {
    const view = visibleRows(rows);
    expect(MAX_VISIBLE_RUNS).toBe(5);
    expect(view.rows.map((r) => r.sim_id)).toEqual(['sim6', 'sim5', 'sim4', 'sim3', 'sim2']);
    expect(view.hasMore).toBe(true);
  });

  it('reports no more at five or fewer and does not mutate its input', () => {
    const five = rows.slice(0, 5);
    const before = five.map((r) => r.sim_id);
    expect(visibleRows(five).hasMore).toBe(false);
    expect(five.map((r) => r.sim_id)).toEqual(before);
  });
});
