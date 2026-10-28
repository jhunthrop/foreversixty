import { describe, expect, it } from 'vitest';
import { formatRaidDate, lastRaidLine, settleLastRaid } from './last-raid';
import type { CoachRecent, LastRaid } from './types';

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
    expect(line(recent({}), false).href).toBe('/premium');
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
