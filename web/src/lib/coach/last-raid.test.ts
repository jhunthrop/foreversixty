import { beforeEach, describe, expect, it, vi } from 'vitest';
import {
  LAST_RAID_LABEL,
  LAST_RAID_RETRY_CLASS,
  cachedLastRaid,
  createLastRaidLoader,
  formatRaidDate,
  lastRaidLine,
  resetLastRaidCache,
  settleLastRaid,
  type LastRaidState,
} from './last-raid';
import type { CoachRecent, LastRaid } from './types';

// The loader remembers the page's last answer across mounts; every test starts without one.
beforeEach(() => resetLastRaidCache());

const RAID: LastRaid = {
  zone: 'Molten Core',
  fought_at: '2026-10-23T20:00:00Z',
  report_id: 'abc123',
  finding_count: 11,
  first: { stem: 'Bloodthirst casts lost', on: 5, of: 8 },
  url: '/report/abc123?fight=3&tab=coach',
};
const recent = (over: Partial<LastRaid> | null, pending = false): CoachRecent => ({
  last_raid: over === null ? null : { ...RAID, ...over },
  pending,
  reviews: [],
});
const line = (r: CoachRecent, premium: boolean) => lastRaidLine(r, premium, 'UTC');
const text = (r: CoachRecent, premium: boolean): string =>
  line(r, premium)
    .segments.map((s) => s.text)
    .join(' · ');

describe('lastRaidLine', () => {
  it('Premium with a fix on some bosses', () => {
    expect(text(recent({}), true)).toBe(
      'Last raid · Molten Core, Fri 23 Oct · Fix first: Bloodthirst casts lost, on 5 of 8 bosses · Open review',
    );
    expect(line(recent({}), true).href).toBe(RAID.url);
  });

  it('Premium one boss and all bosses', () => {
    expect(text(recent({ first: { stem: 'X', on: 1, of: 1 } }), true)).toBe(
      'Last raid · Molten Core, Fri 23 Oct · Fix first: X · Open review',
    );
    expect(text(recent({ first: { stem: 'X', on: 8, of: 8 } }), true)).toContain('X, on all 8 bosses');
  });

  it('free reads count, first and the upgrade link', () => {
    expect(text(recent({}), false)).toBe(
      'Last raid · Molten Core, Fri 23 Oct · 11 things to fix · First: Bloodthirst casts lost · Read the full review with Premium',
    );
    expect(line(recent({}), false).href).toBe('/premium#coach');
  });

  it('free with one finding, and with no first', () => {
    expect(text(recent({ finding_count: 1 }), false)).toContain('1 thing to fix');
    expect(text(recent({ first: null }), false)).toBe(
      'Last raid · Molten Core, Fri 23 Oct · 11 things to fix · Read the full review with Premium',
    );
  });

  it('clean nights', () => {
    const clean = recent({ finding_count: 0, first: null });
    expect(text(clean, true)).toBe(
      'Last raid · Molten Core, Fri 23 Oct · Nothing to fix that the log can show · Open review',
    );
    expect(text(clean, false)).toBe(
      'Last raid · Molten Core, Fri 23 Oct · Nothing to fix that the log can show',
    );
    expect(line(clean, false).href).toBeNull();
  });

  it('pending shows the line without a first and no link', () => {
    const pending = line(recent({ first: null }, true), true);
    expect(pending.segments.map((s) => s.text).join(' · ')).toBe(
      'Last raid · Molten Core, Fri 23 Oct · Reading your pulls…',
    );
    expect(pending.href).toBeNull();
  });

  it('is shown, muted, with no raid in 14 days', () => {
    const none = line(recent(null), true);
    expect(none.segments).toEqual([{ role: 'muted', text: 'Last raid · no log in the last 14 days' }]);
    expect(none.href).toBeNull();
  });

  it('marks the lead, the place and the link', () => {
    expect(line(recent({}), true).segments.map((s) => s.role)).toEqual(['lead', 'place', 'plain', 'link']);
  });
});

describe('settleLastRaid', () => {
  it('settles a good answer, carrying pending', async () => {
    const state = await settleLastRaid(async () => recent({ first: null }, true), true, 'UTC');
    expect(state.kind === 'ready' && state.pending).toBe(true);
  });

  it('settles a failure as error', async () => {
    const state = await settleLastRaid(() => Promise.reject(new Error('down')), true);
    expect(state).toEqual({ kind: 'error' });
  });

  it('is not pending when there is no raid', async () => {
    const state = await settleLastRaid(async () => recent(null, true), true);
    expect(state.kind === 'ready' && state.pending).toBe(false);
  });
});

describe('formatRaidDate', () => {
  it('is ddd d MMM in the given zone', () => {
    expect(formatRaidDate('2026-10-23T20:00:00Z', 'UTC')).toBe('Fri 23 Oct');
    expect(formatRaidDate('2026-10-23T23:30:00Z', 'Pacific/Auckland')).toBe('Sat 24 Oct');
  });
});

describe('lastRaidLine details', () => {
  it('exposes the stem so the view can clamp it, and names the line', () => {
    expect(line(recent({}), true).stem).toBe('Bloodthirst casts lost');
    expect(line(recent({ finding_count: 0, first: null }), true).stem).toBeNull();
    expect(LAST_RAID_LABEL).toBe('Last raid');
  });

  it('keys the free link on the entitlement, not the sentence', () => {
    expect(line(recent({}), false).segments.at(-1)?.role).toBe('link');
    expect(line(recent({ finding_count: 0, first: null }), false).segments.at(-1)?.role).not.toBe('link');
  });

  it('shows a pending line with a first as the review line', () => {
    expect(text(recent({}, true), true)).toContain('Fix first:');
  });
});

describe('createLastRaidLoader', () => {
  function setup(answers: Array<CoachRecent | Error>) {
    const states: LastRaidState[] = [];
    const timers: Array<{ run: () => void; ms: number; cancelled: boolean }> = [];
    const queue = [...answers];
    const fetchRecent = vi.fn(async () => {
      const next = queue.shift()!;
      if (next instanceof Error) throw next;
      return next;
    });
    const loader = createLastRaidLoader({
      fetchRecent,
      premium: true,
      refreshDelayMs: 5000,
      onState: (state) => states.push(state),
      schedule: (run, ms) => {
        const timer = { run, ms, cancelled: false };
        timers.push(timer);
        return timer;
      },
      cancel: (handle) => {
        (handle as { cancelled: boolean }).cancelled = true;
      },
    });
    return { loader, states, timers, fetchRecent };
  }
  const flush = () => new Promise((resolve) => setTimeout(resolve, 0));
  const kinds = (states: LastRaidState[]) => states.map((s) => s.kind);

  it('a remount starts from the last answer and never shows the skeleton again', async () => {
    const first = setup([recent({})]);
    first.loader.start();
    await flush();
    expect(kinds(first.states)).toEqual(['loading', 'ready']);
    expect(cachedLastRaid()).toEqual(first.states[1]);

    // The account page re-creates the line when /v1/me refreshes: the new loader refreshes in
    // place; the component itself starts from cachedLastRaid().
    const second = setup([recent({ first: null })]);
    second.loader.start();
    await flush();
    expect(kinds(second.states)).toEqual(['ready']);
    expect(second.fetchRecent).toHaveBeenCalledTimes(1);
    expect(cachedLastRaid()).toEqual(second.states[0]);
  });

  it('arms one refresh after a pending first answer and swaps in the settled line', async () => {
    const { loader, states, timers, fetchRecent } = setup([recent({ first: null }, true), recent({})]);
    loader.start();
    await flush();
    expect(timers.map((t) => t.ms)).toEqual([5000]);
    timers[0]!.run();
    await flush();
    expect(fetchRecent).toHaveBeenCalledTimes(2);
    expect(timers).toHaveLength(1);
    expect(kinds(states)).toEqual(['loading', 'ready', 'ready']);
  });

  it('arms nothing when the first answer is settled', async () => {
    const { loader, timers } = setup([recent({})]);
    loader.start();
    await flush();
    expect(timers).toHaveLength(0);
  });

  it('clears the timer on stop and drops a late answer', async () => {
    const { loader, states, timers } = setup([recent({ first: null }, true)]);
    loader.start();
    await flush();
    loader.stop();
    expect(timers[0]!.cancelled).toBe(true);
    expect(kinds(states)).toEqual(['loading', 'ready']);
  });

  it('does not arm a refresh after an error', async () => {
    const { loader, states, timers } = setup([new Error('down')]);
    loader.start();
    await flush();
    expect(kinds(states)).toEqual(['loading', 'error']);
    expect(timers).toHaveLength(0);
  });

  it('retry shows retrying, then the answer, and arms no refresh', async () => {
    const { loader, states, timers } = setup([new Error('down'), recent({ first: null }, true)]);
    loader.start();
    await flush();
    loader.retry();
    await flush();
    expect(kinds(states)).toEqual(['loading', 'error', 'retrying', 'ready']);
    expect(timers).toHaveLength(0);
  });

  it('a failed refresh keeps the line on screen', async () => {
    const { loader, states, timers } = setup([recent({ first: null }, true), new Error('down')]);
    loader.start();
    await flush();
    timers[0]!.run();
    await flush();
    expect(kinds(states)).toEqual(['loading', 'ready']);
  });
});

describe('the Try again control', () => {
  it('is a 44 px target that does not grow the 20 px line', () => {
    const classes = LAST_RAID_RETRY_CLASS.split(' ');
    expect(classes).toContain('min-h-11');
    expect(classes).toContain('-my-3');
  });
});

describe('a viewer that flips from free to Premium', () => {
  const build = (read: () => boolean, states: LastRaidState[]) =>
    createLastRaidLoader({
      fetchRecent: async () => recent({}),
      get premium() {
        return read();
      },
      refreshDelayMs: 5000,
      onState: (state) => states.push(state),
    });
  const settle = () => new Promise((resolve) => setTimeout(resolve, 0));

  it('re-derives the copy from the kept answer: the upsell becomes the review link', async () => {
    let premium = false;
    const states: LastRaidState[] = [];
    build(() => premium, states).start();
    await settle();
    const settled = states.at(-1);
    if (settled?.kind !== 'ready') throw new Error('expected a ready state');
    expect(lastRaidLine(settled.recent, premium, 'UTC').href).toBe('/premium#coach');
    premium = true;
    const flipped = lastRaidLine(settled.recent, premium, 'UTC');
    expect(flipped.href).toBe(RAID.url);
    expect(flipped.segments.map((s) => s.text).join(' · ')).toContain('Fix first: Bloodthirst casts lost');
  });
  it('the loader reads premium at settle time, not at construction', async () => {
    let premium = false;
    const states: LastRaidState[] = [];
    const loader = build(() => premium, states);
    premium = true;
    loader.start();
    await settle();
    const settled = states.at(-1);
    expect(settled?.kind === 'ready' && settled.view.href).toBe(RAID.url);
  });
});
