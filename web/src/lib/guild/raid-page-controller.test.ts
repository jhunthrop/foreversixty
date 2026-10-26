// web/src/lib/guild/raid-page-controller.test.ts
// The raid page's controller with fake deps and fake timers (raids spec 4 "Page states", "Poll",
// "Races", R1): who may see the raid, the 10 s poll while visible and open, the pause after three
// failed polls, the optimistic row kept until its own answer, the race codes and the plan gate.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import type { Me } from '../account/api';
import type { GuildPage } from '../rankings/api';
import { NOW, plannedRaid, viewerCharacter } from '../../test-support/raid-fixtures';
import { GuildApiError, type GuildHome, type GuildRaidView } from './api';
import { raidCopy } from './raid-copy';
import { RaidPageController, type RaidPageDeps } from './raid-page-controller.svelte';
import { RAID_POLL_MS } from './raid-page';

const ROUTE = { guild: { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' }, raidId: '7' } as const;

const GUILD_PAGE = {
  guild: { id: 2, name: 'Olympus XXVII', region: 'us', ruleset: 'pvp', faction: 'horde', crest_url: null },
  progression: [],
  roster_best: [],
  reports: [],
} as unknown as GuildPage;

function me(guild: Partial<Me['guilds'][number]> | null): Me {
  const guilds =
    guild === null
      ? []
      : [
          {
            id: 2,
            region: 'us',
            ruleset: 'pvp',
            name: 'Olympus XXVII',
            rank: 'member',
            verified: true,
            ...guild,
          },
        ];
  return { user: { id: 1, battletag: 'a#1' }, characters: [], guilds } as unknown as Me;
}

function deferred<T>(): { promise: Promise<T>; resolve: (v: T) => void; reject: (e: unknown) => void } {
  let resolve!: (v: T) => void;
  let reject!: (e: unknown) => void;
  const promise = new Promise<T>((res, rej) => {
    resolve = res;
    reject = rej;
  });
  return { promise, resolve, reject };
}

const ZONES = [
  { key: 'barrow-deeps', name: 'Barrow Deeps', size: null },
  { key: 'onyxias-lair', name: "Onyxia's Lair", size: 40 },
];

let visible = true;

function deps(extra: Partial<RaidPageDeps> = {}): RaidPageDeps {
  return {
    loadGuild: () => Promise.resolve(GUILD_PAGE),
    loadMe: () => Promise.resolve(me({})),
    loadHome: () =>
      Promise.resolve({
        viewer: { role: 'member', character_key: 'us-pvp-kraggor', verified: true, officer_tools: true },
      } as unknown as GuildHome),
    fetchRaid: () => Promise.resolve(plannedRaid()),
    refetchRaid: () => Promise.resolve(plannedRaid()),
    putSignup: () =>
      Promise.resolve({ raid: plannedRaid({ counts: { ...plannedRaid().counts, in: 22 } }), moved: null }),
    setStatus: () => Promise.resolve(plannedRaid({ status: 'locked' })),
    updateRaid: () => Promise.resolve(plannedRaid({ title: 'Ony farm' })),
    loadZones: () => Promise.resolve(ZONES),
    isVisible: () => visible,
    now: () => new Date(Date.now()),
    storage: null,
    ...extra,
  };
}

async function started(d: RaidPageDeps): Promise<{ c: RaidPageController; stop: () => void }> {
  const c = new RaidPageController(d, ROUTE);
  const stop = c.start();
  await vi.advanceTimersByTimeAsync(0);
  return { c, stop };
}

beforeEach(() => {
  vi.useFakeTimers();
  vi.setSystemTime(NOW);
  visible = true;
});

afterEach(() => {
  vi.useRealTimers();
});

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
    visible = false;
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS * 3);
    expect(refetchRaid).not.toHaveBeenCalled();
    visible = true;
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
});

describe('officer actions', () => {
  const officerMe = (): Promise<Me> => Promise.resolve(me({ rank: 'officer' }));
  const officerHome = (): Promise<GuildHome> =>
    Promise.resolve({
      viewer: { role: 'officer', character_key: null, verified: true, officer_tools: true },
    } as unknown as GuildHome);

  it('applies the lock answer', async () => {
    const setStatus = vi.fn(() => Promise.resolve(plannedRaid({ status: 'locked' })));
    const { c, stop } = await started(deps({ loadMe: officerMe, loadHome: officerHome, setStatus }));
    await expect(c.runStatus('lock')).resolves.toBe(true);
    expect(c.raid?.status).toBe('locked');
    expect(setStatus).toHaveBeenCalledWith(2, '7', 'lock', undefined);
    stop();
  });

  it('sends the Discord choice with a cancel', async () => {
    const setStatus = vi.fn(() => Promise.resolve(plannedRaid({ status: 'cancelled' })));
    const { c, stop } = await started(deps({ loadMe: officerMe, loadHome: officerHome, setStatus }));
    await c.runStatus('cancel', true);
    expect(setStatus).toHaveBeenCalledWith(2, '7', 'cancel', { post_to_discord: true });
    stop();
  });

  it('ends the officer tools on guild_plan_required, with the line under the control used', async () => {
    const refusal = new GuildApiError('plan', 403, 'guild_plan_required');
    const { c, stop } = await started(
      deps({ loadMe: officerMe, loadHome: officerHome, setStatus: () => Promise.reject(refusal) }),
    );
    await expect(c.runStatus('lock')).resolves.toBe(false);
    expect(c.officerTools).toBe(false);
    expect(c.line).toEqual({ control: 'lock', message: raidCopy.races.officerToolsEnded });
    stop();
  });

  it('re-reads the raid after a conflict and says it did not save', async () => {
    const refetchRaid = vi.fn(() => Promise.resolve(plannedRaid({ status: 'cancelled' })));
    const refusal = new GuildApiError('conflict', 409, 'conflict');
    const { c, stop } = await started(
      deps({
        loadMe: officerMe,
        loadHome: officerHome,
        refetchRaid,
        setStatus: () => Promise.reject(refusal),
      }),
    );
    await c.runStatus('lock');
    await vi.advanceTimersByTimeAsync(0);
    expect(c.line).toEqual({ control: 'lock', message: raidCopy.officerRow.failed });
    expect(c.raid?.status).toBe('cancelled');
    stop();
  });

  it('sets a roster row with the spec role and keeps its note', async () => {
    const putSignup = vi.fn(() => Promise.resolve({ raid: plannedRaid(), moved: null }));
    const { c, stop } = await started(deps({ loadMe: officerMe, loadHome: officerHome, putSignup }));
    const row = {
      character_key: 'us/pvp/mira',
      name: 'Mira',
      class: 'mage',
      spec: 'Frost',
      role: 'dps' as const,
      state: null,
      note: 'late',
      own: false,
    };
    await expect(c.setRosterState(row, 'bench')).resolves.toBe(true);
    expect(putSignup).toHaveBeenCalledWith(2, '7', 'us/pvp/mira', {
      state: 'bench',
      role: 'dps',
      note: 'late',
    });
    stop();
  });

  it('puts a roster refusal under the roster, a plan refusal under that row', async () => {
    const row = {
      character_key: 'us/pvp/mira',
      name: 'Mira',
      class: 'mage',
      spec: 'Frost',
      role: 'dps' as const,
      state: null,
      note: '',
      own: false,
    };
    const failing = await started(
      deps({ loadMe: officerMe, loadHome: officerHome, putSignup: () => Promise.reject(new Error('x')) }),
    );
    await expect(failing.c.setRosterState(row, 'in')).resolves.toBe(false);
    expect(failing.c.line).toEqual({ control: 'roster', message: raidCopy.officerRow.failed });
    failing.stop();
    const plan = new GuildApiError('plan', 403, 'guild_plan_required');
    const gated = await started(
      deps({ loadMe: officerMe, loadHome: officerHome, putSignup: () => Promise.reject(plan) }),
    );
    await gated.c.setRosterState(row, 'in');
    expect(gated.c.line).toEqual({
      control: 'roster:us/pvp/mira',
      message: raidCopy.races.officerToolsEnded,
    });
    expect(gated.c.officerTools).toBe(false);
    gated.stop();
  });

  it('saves an edit and maps a refusal onto the form', async () => {
    const { c, stop } = await started(deps({ loadMe: officerMe }));
    const payload = {
      zone: 'onyxias-lair',
      title: 'Ony farm',
      starts_at: '2026-12-11T02:00:00Z',
      duration_minutes: 180,
      notes: '',
      timezone: 'America/Chicago',
      post_to_discord: false,
    };
    c.openEdit();
    expect(c.editing).toBe(true);
    expect(c.zones.map((z) => z.key)).toEqual(['onyxias-lair']);
    await vi.advanceTimersByTimeAsync(0);
    expect(c.zones).toEqual(ZONES);
    await expect(c.saveEdit(payload)).resolves.toEqual({ ok: true });
    expect(c.raid?.title).toBe('Ony farm');
    expect(c.editing).toBe(false);
    stop();
    const plan = new GuildApiError('plan', 403, 'guild_plan_required');
    const gated = await started(
      deps({ loadMe: officerMe, loadHome: officerHome, updateRaid: () => Promise.reject(plan) }),
    );
    await expect(gated.c.saveEdit(payload)).resolves.toMatchObject({
      ok: false,
      code: 'guild_plan_required',
    });
    expect(gated.c.officerTools).toBe(false);
    gated.stop();
  });
});
