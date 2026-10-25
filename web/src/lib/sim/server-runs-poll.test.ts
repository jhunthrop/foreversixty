// @vitest-environment jsdom
import { beforeEach, describe, expect, it } from 'vitest';
import {
  LIVE_TTL_MS,
  hasInFlight,
  isLive,
  markLive,
  newlyFinished,
  readRowCount,
  reservedRowCount,
  withSeen,
  writeRowCount,
} from './server-runs-poll';
import type { SimListRow } from './types';

function row(id: string, state: SimListRow['state'], seen?: boolean): SimListRow {
  return {
    sim_id: id,
    spec: 'mage-frost',
    dps: 1204,
    engine_version: 'e',
    created_at: '2026-10-24T10:00:00Z',
    title: '',
    state,
    seen,
  };
}

beforeEach(() => localStorage.clear());

describe('in flight and newly finished', () => {
  it('sees queued and running rows', () => {
    expect(hasInFlight([row('a', 'done'), row('b', 'queued')])).toBe(true);
    expect(hasInFlight([row('a', 'running')])).toBe(true);
    expect(hasInFlight([row('a', 'done'), row('b', 'error')])).toBe(false);
    expect(hasInFlight(null)).toBe(false);
  });

  it('names only rows that flipped from in flight to done and unopened', () => {
    const before = [row('a', 'running'), row('b', 'running'), row('c', 'done', false)];
    const after = [row('a', 'done', false), row('b', 'error'), row('c', 'done', false)];
    expect(newlyFinished(before, after).map((r) => r.sim_id)).toEqual(['a']);
    expect(newlyFinished(null, after)).toEqual([]);
  });

  it('marks ids seen without changing the input', () => {
    const rows = [row('a', 'done', false)];
    expect(withSeen(rows, ['a'])[0]?.seen).toBe(true);
    expect(rows[0]?.seen).toBe(false);
  });
});

describe('the reserve hints', () => {
  it('clamps the reserved rows to one through five', () => {
    expect(reservedRowCount(0)).toBe(1);
    expect(reservedRowCount(3)).toBe(3);
    expect(reservedRowCount(9)).toBe(5);
    expect(reservedRowCount(Number.NaN)).toBe(1);
  });

  it('round-trips the row count and defaults to one', () => {
    expect(readRowCount()).toBe(1);
    writeRowCount(3);
    expect(readRowCount()).toBe(3);
  });

  it('expires the live flag two hours after it was set', () => {
    expect(isLive()).toBe(false);
    markLive(1_000_000);
    expect(isLive(1_000_000 + LIVE_TTL_MS - 1)).toBe(true);
    expect(isLive(1_000_000 + LIVE_TTL_MS)).toBe(false);
  });
});
