// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import {
  HISTORY_MAX_RESERVE_PX,
  HISTORY_RESERVE_KEY,
  reserveToStore,
  afterLayout,
  opensOnSelection,
  scrollRowIntoView,
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

describe('scrollRowIntoView', () => {
  /** A mounted list of `rows` items, the ones at `marked` carrying data-marked, with spies. */
  function listHost(
    rows: number,
    marked: number[],
  ): { host: HTMLElement; spies: ReturnType<typeof vi.fn>[] } {
    const host = document.createElement('div');
    const spies = Array.from({ length: rows }, (_, index) => {
      const row = document.createElement('li');
      row.dataset.marked = String(marked.includes(index));
      const spy = vi.fn();
      row.scrollIntoView = spy;
      host.append(row);
      return spy;
    });
    document.body.append(host);
    return { host, spies };
  }

  afterEach(() => document.body.replaceChildren());

  it('jumps the first marked row to the middle at once, and only that row', () => {
    vi.useFakeTimers();
    const { host, spies } = listHost(12, [9, 10]);
    scrollRowIntoView(host);
    expect(spies[9]).toHaveBeenCalledExactlyOnceWith({ block: 'center', behavior: 'auto' });
    expect(spies.filter((spy) => spy.mock.calls.length > 0)).toHaveLength(1);
    expect(vi.getTimerCount()).toBe(0);
    vi.useRealTimers();
  });

  it('finds the row at scroll time, not when the list was handed over', () => {
    const { host } = listHost(3, []);
    const late = document.createElement('li');
    late.dataset.marked = 'true';
    late.scrollIntoView = vi.fn();
    host.append(late);
    scrollRowIntoView(host);
    expect(late.scrollIntoView).toHaveBeenCalledOnce();
  });

  it('does nothing for a list with no mark, an unmounted host or no host', () => {
    const { host, spies } = listHost(4, []);
    expect(() => scrollRowIntoView(host)).not.toThrow();
    const detached = listHost(4, [2]);
    detached.host.remove();
    scrollRowIntoView(detached.host);
    expect([...spies, ...detached.spies].every((spy) => spy.mock.calls.length === 0)).toBe(true);
    expect(() => scrollRowIntoView(undefined)).not.toThrow();
    expect(() => scrollRowIntoView(null)).not.toThrow();
  });
});

describe('afterLayout', () => {
  it('runs only after the DOM update and one laid-out frame', async () => {
    const order: string[] = [];
    let frame: (() => void) | undefined;
    const done = afterLayout(() => order.push('scroll'), {
      tick: async () => {
        order.push('tick');
      },
      frame: (run) => {
        order.push('frame');
        frame = run;
      },
    });
    expect(order).toEqual(['tick']);
    await done;
    expect(order).toEqual(['tick', 'frame']);
    frame?.();
    expect(order).toEqual(['tick', 'frame', 'scroll']);
  });
});

describe('opensOnSelection', () => {
  it('opens for a new selection whose rows are hidden', () => {
    expect(opensOnSelection('', 'rpt1:7', true)).toBe(true);
    expect(opensOnSelection('rpt1:7', 'rpt9:2', true)).toBe(true);
  });

  it('stays put for an equal selection handed over again (a rebuilt chart, a hover)', () => {
    expect(opensOnSelection('rpt1:7', 'rpt1:7', true)).toBe(false);
  });

  it('stays put for a selection whose rows are already shown', () => {
    expect(opensOnSelection('', 'rpt1:7', false)).toBe(false);
  });
});

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
