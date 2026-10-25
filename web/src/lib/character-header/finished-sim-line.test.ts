import { describe, expect, it, vi } from 'vitest';
import type { FinishedSim, Me } from '../account/api';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import { finishedRunsFor, markAllSeen } from './finished-sim-line';

const RUN: FinishedSim = {
  sim_id: 'sim-1',
  spec: 'mage-frost',
  title: '',
  dps: 1204,
  finished_at: '2026-10-24T11:00:00Z',
};

function meWith(server: boolean, finished: FinishedSim[] | undefined): Me {
  return {
    user: { id: 7 },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ server_sims: server }),
    finished_sims: finished,
  } as unknown as Me;
}

describe('finishedRunsFor', () => {
  it('names the runs for a member with server sims', () => {
    expect(finishedRunsFor(meWith(true, [RUN, { ...RUN, sim_id: 'sim-2' }]))?.map((r) => r.sim_id)).toEqual([
      'sim-1',
      'sim-2',
    ]);
  });

  it.each([
    ['no server sims', meWith(false, [RUN])],
    ['an empty list', meWith(true, [])],
    ['an older API with no list', meWith(true, undefined)],
    ['no answer yet', null],
  ])('is absent for %s', (_name, me) => {
    expect(finishedRunsFor(me)).toBeNull();
  });
});

describe('markAllSeen', () => {
  it('posts one seen call per run and succeeds when all answer', async () => {
    const markSeen = vi.fn().mockResolvedValue(undefined);
    expect(await markAllSeen([RUN, { ...RUN, sim_id: 'sim-2' }], markSeen)).toBe(true);
    expect(markSeen.mock.calls.map(([id]) => id)).toEqual(['sim-1', 'sim-2']);
  });

  it('fails when any one call fails, after trying them all', async () => {
    const markSeen = vi.fn().mockRejectedValueOnce(new Error('x')).mockResolvedValue(undefined);
    expect(await markAllSeen([RUN, { ...RUN, sim_id: 'sim-2' }], markSeen)).toBe(false);
    expect(markSeen).toHaveBeenCalledTimes(2);
  });
});
