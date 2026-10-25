import { describe, expect, it, vi } from 'vitest';
import type { Me } from '../account/api';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import { markSimSeenOnOpen } from './mark-seen-on-open';

function me(server: boolean, unseen: string[] = ['sim-1']): Me {
  return {
    user: { id: 7 },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ server_sims: server }),
    finished_sims: unseen.map((sim_id) => ({
      sim_id,
      spec: 'mage-frost',
      title: '',
      dps: 1,
      finished_at: '',
    })),
  } as unknown as Me;
}

describe('markSimSeenOnOpen', () => {
  it('marks the sim seen once when it is in the member’s unseen list', async () => {
    const markSeen = vi.fn().mockResolvedValue(undefined);
    await markSimSeenOnOpen('sim-1', () => Promise.resolve(me(true)), markSeen);
    expect(markSeen).toHaveBeenCalledTimes(1);
    expect(markSeen).toHaveBeenCalledWith('sim-1');
  });

  it('never calls for a signed-out visitor', async () => {
    const markSeen = vi.fn();
    await markSimSeenOnOpen('sim-1', () => Promise.resolve(null), markSeen);
    expect(markSeen).not.toHaveBeenCalled();
  });

  it('never calls for a member without server sims (an empty list)', async () => {
    const markSeen = vi.fn();
    await markSimSeenOnOpen('sim-1', () => Promise.resolve(me(false, [])), markSeen);
    expect(markSeen).not.toHaveBeenCalled();
  });

  it('never calls for a payer opening a sim that is not in the list', async () => {
    const markSeen = vi.fn();
    await markSimSeenOnOpen('sim-9', () => Promise.resolve(me(true, ['sim-1', 'sim-2'])), markSeen);
    expect(markSeen).not.toHaveBeenCalled();
  });

  it('never calls without an id (an inline result)', async () => {
    const markSeen = vi.fn();
    await markSimSeenOnOpen('', () => Promise.resolve(me(true)), markSeen);
    expect(markSeen).not.toHaveBeenCalled();
  });

  it('ignores a failed read and a failed seen call', async () => {
    await expect(
      markSimSeenOnOpen('sim-1', () => Promise.reject(new Error('down')), vi.fn()),
    ).resolves.toBeUndefined();
    await expect(
      markSimSeenOnOpen(
        'sim-1',
        () => Promise.resolve(me(true)),
        () => Promise.reject(new Error('404')),
      ),
    ).resolves.toBeUndefined();
  });
});
