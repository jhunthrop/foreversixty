// @vitest-environment jsdom
import { afterEach, describe, expect, it } from 'vitest';
import {
  HISTORY_MAX_RESERVE_PX,
  HISTORY_RESERVE_KEY,
  reserveToStore,
  scrollBehavior,
  wrapReason,
  writeHistoryReserve,
} from './history-layout';

describe('scrollBehavior', () => {
  it('glides unless the visitor asked for less motion', () => {
    expect(scrollBehavior()).toBe('smooth');
    window.matchMedia = ((query: string) => ({
      matches: query.includes('reduce'),
    })) as typeof window.matchMedia;
    expect(scrollBehavior()).toBe('auto');
  });
});

afterEach(() => localStorage.clear());

describe('wrapReason', () => {
  it('keeps a short reason on one line', () => {
    expect(wrapReason('Report flagged', 200)).toEqual(['Report flagged']);
  });

  it('splits a reason that does not fit into two lines at the middle space', () => {
    expect(wrapReason('No rated fights', 72)).toEqual(['No rated', 'fights']);
  });

  it('never makes more than two lines', () => {
    expect(wrapReason('Score not kept at all today', 30)).toHaveLength(2);
  });

  it('leaves one word alone', () => {
    expect(wrapReason('Unknown', 10)).toEqual(['Unknown']);
  });
});

describe('reserveToStore', () => {
  it('keeps the settled, collapsed height', () => {
    expect(reserveToStore('ready', false, 1197)).toBe(1197);
  });

  it('keeps nothing while an owner has opened a list, in an error or while loading', () => {
    expect(reserveToStore('ready', true, 2400)).toBeNull();
    expect(reserveToStore('error', false, 120)).toBeNull();
    expect(reserveToStore('loading', false, 1197)).toBeNull();
  });
});

describe('writeHistoryReserve', () => {
  it('stores the rounded height under fs.history.reserve', () => {
    writeHistoryReserve(1059.6);
    expect(HISTORY_RESERVE_KEY).toBe('fs.history.reserve');
    expect(localStorage.getItem(HISTORY_RESERVE_KEY)).toBe('1060');
  });

  it('drops a height that is not a sane box', () => {
    writeHistoryReserve(0);
    writeHistoryReserve(Number.NaN);
    writeHistoryReserve(HISTORY_MAX_RESERVE_PX + 1);
    expect(localStorage.getItem(HISTORY_RESERVE_KEY)).toBeNull();
    writeHistoryReserve(2125);
    expect(localStorage.getItem(HISTORY_RESERVE_KEY)).toBe('2125');
  });
});
