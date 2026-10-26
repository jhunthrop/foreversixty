// web/src/lib/guild/raid-page-controller.test.ts
// The raid page's controller with fake deps and fake timers (raids spec 4 "Page states", "Poll",
// "Races", R1): who may see the raid, the 10 s poll while visible and open, the pause after three
// failed polls, the optimistic row kept until its own answer and the race codes. Officer actions
// are in raid-page-controller-officer.test.ts.
import { describe, expect, it, vi } from 'vitest';
import {
  ROUTE,
  deferred,
  deps,
  me,
  started,
  useControllerClock,
  visibility,
} from '../../test-support/raid-controller-fakes';
import { plannedRaid, viewerCharacter } from '../../test-support/raid-fixtures';
import { GuildApiError, type GuildRaidView } from './api';
import { raidCopy } from './raid-copy';
import { RAID_POLL_MS } from './raid-page';
import { RaidPageController } from './raid-page-controller.svelte';

useControllerClock();

describe('access (R1)', () => {
  it('is signed-out, not-member or unverified without reading the raid', async () => {
    const fetchRaid = vi.fn(() => Promise.resolve(plannedRaid()));
    for (const [answer, access] of [
      [null, 'signed-out'],
      [me(null), 'not-member'],
      [me({ verified: false }), 'unverified'],
    ] as const) {
      const { c, stop } = await started(deps({ loadMe: () => Promise.resolve(answer), fetchRaid }));
      expect(c.access).toBe(access);
      expect(c.load).toBe('ready');
      expect(c.guild?.name).toBe('Olympus XXVII');
      stop();
    }
    expect(fetchRaid).not.toHaveBeenCalled();
  });

  it('loads the raid for a verified member, with the home viewer', async () => {
    const { c, stop } = await started(deps());
    expect(c.access).toBe('member');
    expect(c.raid?.id).toBe(7);
    expect(c.currentKey).toBe('us-pvp-kraggor');
    expect(c.role).toBe('member');
    stop();
  });

  it('reads an officer rank as officer before the home answers', async () => {
    const { c, stop } = await started(
      deps({ loadMe: () => Promise.resolve(me({ rank: 'leader' })), loadHome: () => new Promise(() => {}) }),
    );
    expect(c.role).toBe('officer');
    expect(c.officerTools).toBe(true);
    stop();
  });

  it('fails the page, never signed-out, when /v1/me does not answer', async () => {
    const { c, stop } = await started(deps({ loadMe: () => Promise.reject(new Error('offline')) }));
    expect(c.load).toBe('failed');
    expect(c.access).toBe('loading');
    stop();
  });

  it('fails the page when the raid does not load, and retries', async () => {
    let fail = true;
    const fetchRaid = vi.fn(() =>
      fail ? Promise.reject(new GuildApiError('gone', 404, 'not_found')) : Promise.resolve(plannedRaid()),
    );
    const { c, stop } = await started(deps({ fetchRaid }));
    expect(c.load).toBe('failed');
    fail = false;
    c.retry();
    await vi.advanceTimersByTimeAsync(0);
    expect(c.load).toBe('ready');
    stop();
  });
});

describe('poll', () => {
  it('asks again every 10 s while visible', async () => {
    const refetchRaid = vi.fn(() =>
      Promise.resolve(plannedRaid({ counts: { ...plannedRaid().counts, in: 30 } })),
    );
    const { c, stop } = await started(deps({ refetchRaid }));
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS - 1);
    expect(refetchRaid).not.toHaveBeenCalled();
    await vi.advanceTimersByTimeAsync(1);
    expect(refetchRaid).toHaveBeenCalledTimes(1);
    expect(c.raid?.counts.in).toBe(30);
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS * 2);
    expect(refetchRaid).toHaveBeenCalledTimes(3);
    stop();
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS * 2);
    expect(refetchRaid).toHaveBeenCalledTimes(3);
  });

  it('skips while the tab is hidden', async () => {
    const refetchRaid = vi.fn(() => Promise.resolve(plannedRaid()));
    const { stop } = await started(deps({ refetchRaid }));
    visibility.visible = false;
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS * 3);
    expect(refetchRaid).not.toHaveBeenCalled();
    visibility.visible = true;
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS);
    expect(refetchRaid).toHaveBeenCalledTimes(1);
    stop();
  });

  it('never polls a cancelled or done raid', async () => {
    const refetchRaid = vi.fn(() => Promise.resolve(plannedRaid()));
    const { stop } = await started(
      deps({ fetchRaid: () => Promise.resolve(plannedRaid({ status: 'cancelled' })), refetchRaid }),
    );
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS * 3);
    expect(refetchRaid).not.toHaveBeenCalled();
    stop();
  });

  it('is silent on a failure and pauses after three in a row', async () => {
    const refetchRaid = vi.fn(() => Promise.reject(new Error('offline')));
    const { c, stop } = await started(deps({ refetchRaid }));
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS * 2);
    expect(c.pollPaused).toBe(false);
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS);
    expect(c.pollPaused).toBe(true);
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS * 3);
    expect(refetchRaid).toHaveBeenCalledTimes(3);
    stop();
  });

  it('counts only failures in a row', async () => {
    let calls = 0;
    const refetchRaid = vi.fn(() => {
      calls += 1;
      return calls === 3 ? Promise.resolve(plannedRaid()) : Promise.reject(new Error('offline'));
    });
    const { c, stop } = await started(deps({ refetchRaid }));
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS * 4);
    expect(c.pollPaused).toBe(false);
    stop();
  });

  it('keeps the optimistic row until its own answer: a poll during a save is dropped', async () => {
    const put = deferred<{ raid: GuildRaidView; moved: string | null }>();
    const refetchRaid = vi.fn(() =>
      Promise.resolve(plannedRaid({ viewer: { characters: [viewerCharacter({ state: null })] } })),
    );
    const { c, stop } = await started(deps({ refetchRaid, putSignup: () => put.promise }));
    const outcome = c.signup('us-pvp-kraggor', 'bench', 'dps', '');
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS);
    expect(refetchRaid).toHaveBeenCalledTimes(1);
    expect(c.raid?.viewer.characters[0]?.state).toBe('in');
    put.resolve({
      raid: plannedRaid({ viewer: { characters: [viewerCharacter({ state: 'bench' })] } }),
      moved: null,
    });
    await expect(outcome).resolves.toEqual({ ok: true, moved: null });
    expect(c.raid?.viewer.characters[0]?.state).toBe('bench');
    stop();
  });

  it('drops a poll that started before a save and answered after it', async () => {
    const poll = deferred<GuildRaidView>();
    const { c, stop } = await started(deps({ refetchRaid: () => poll.promise }));
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS);
    await c.signup('us-pvp-kraggor', 'in', 'dps', '');
    poll.resolve(plannedRaid({ counts: { ...plannedRaid().counts, in: 1 } }));
    await vi.advanceTimersByTimeAsync(0);
    expect(c.raid?.counts.in).toBe(22);
    stop();
  });
});

describe('sign-up races', () => {
  it('resolves the moved name after applying the answer', async () => {
    const { c, stop } = await started(
      deps({ putSignup: () => Promise.resolve({ raid: plannedRaid({ title: 'New' }), moved: 'Menda' }) }),
    );
    await expect(c.signup('us-pvp-kraggor', 'in', 'dps', '')).resolves.toEqual({ ok: true, moved: 'Menda' });
    expect(c.raid?.title).toBe('New');
    stop();
  });

  it('turns the page locked on raid_locked and cancelled on raid_cancelled', async () => {
    for (const [code, status] of [
      ['raid_locked', 'locked'],
      ['raid_cancelled', 'cancelled'],
    ] as const) {
      const refusal = new GuildApiError('refused', 409, code);
      const { c, stop } = await started(deps({ putSignup: () => Promise.reject(refusal) }));
      const outcome = await c.signup('us-pvp-kraggor', 'in', 'dps', '');
      expect(outcome).toMatchObject({ ok: false, code });
      expect(c.raid?.status).toBe(status);
      stop();
    }
  });

  it('lifts the cancelled race line to the page, where the cancelled strip has no rows', async () => {
    const cancelled = await started(
      deps({ putSignup: () => Promise.reject(new GuildApiError('refused', 409, 'raid_cancelled')) }),
    );
    expect(cancelled.c.raceLine).toBeNull();
    await cancelled.c.signup('us-pvp-kraggor', 'in', 'dps', '');
    expect(cancelled.c.raceLine).toBe(raidCopy.races.cancelled);
    cancelled.stop();
    const locked = await started(
      deps({ putSignup: () => Promise.reject(new GuildApiError('refused', 409, 'raid_locked')) }),
    );
    await locked.c.signup('us-pvp-kraggor', 'in', 'dps', '');
    expect(locked.c.raceLine).toBeNull();
    locked.stop();
  });
});

describe('lifecycle', () => {
  it('keeps one poll when started twice', async () => {
    const refetchRaid = vi.fn(() => Promise.resolve(plannedRaid()));
    const { c, stop } = await started(deps({ refetchRaid }));
    const stopAgain = c.start();
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS * 2);
    expect(refetchRaid).toHaveBeenCalledTimes(2);
    stopAgain();
    stop();
  });

  it('drops a load that answers after the stop', async () => {
    const raid = deferred<GuildRaidView>();
    const c = new RaidPageController(deps({ fetchRaid: () => raid.promise }), ROUTE);
    const stop = c.start();
    await vi.advanceTimersByTimeAsync(0);
    stop();
    raid.resolve(plannedRaid({ title: 'Late answer' }));
    await vi.advanceTimersByTimeAsync(0);
    expect(c.raid).toBeNull();
    expect(c.load).toBe('loading');
  });
});
