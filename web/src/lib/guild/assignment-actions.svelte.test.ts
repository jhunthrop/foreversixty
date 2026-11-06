import { describe, expect, it, vi } from 'vitest';
import { MAGE, TANK, onyxia, picked, slot } from '../../test-support/assignment-fixtures';
import { plannedRaid } from '../../test-support/raid-fixtures';
import { GuildApiError, type GuildRaidView } from './api';
import { AssignmentActions, slotId } from './assignment-actions.svelte';
import { assignmentsCopy } from './assignments-view';
import type { NightActionHost } from './night-actions.svelte';
import type { RaidPageDeps } from './raid-page-types';

const ROUTE = { guild: { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' }, raidId: '7' } as const;
type Deps = Pick<RaidPageDeps, 'putAssignment' | 'deleteAssignment' | 'postAssignments'>;

function setup(deps: Partial<Deps> = {}) {
  let raid: GuildRaidView | null = plannedRaid({ assignments: onyxia() });
  const host: NightActionHost = {
    raid: () => raid,
    write: (run) => run(2),
    applyRaid: (next) => (raid = next),
    refusalLine: (_thrown, _ended, failed) => failed,
  };
  const actions = new AssignmentActions(
    host,
    {
      putAssignment: deps.putAssignment ?? (() => Promise.resolve(plannedRaid({ title: 'saved' }))),
      deleteAssignment: deps.deleteAssignment ?? (() => Promise.resolve(plannedRaid({ title: 'cleared' }))),
      postAssignments:
        deps.postAssignments ??
        (() =>
          Promise.resolve({
            posted: true,
            posted_at: '2026-12-10T20:00:00Z',
            next_allowed_at: '2026-12-10T20:01:00Z',
          })),
    },
    ROUTE,
  );
  return { actions, raid: () => raid };
}

describe('AssignmentActions', () => {
  const s = slot({ count: 3, names: [picked(TANK)] });

  it('sends the whole slot on an add and a remove, and applies the answered raid', async () => {
    const put = vi.fn(() => Promise.resolve(plannedRaid({ title: 'saved' })));
    const { actions, raid } = setup({ putAssignment: put });
    expect(await actions.add(1084, s, MAGE.character_key)).toBe(true);
    expect(put).toHaveBeenCalledWith(2, '7', 1084, 'main_tank', [TANK.character_key, MAGE.character_key]);
    expect(await actions.remove(1084, s, TANK.character_key)).toBe(true);
    expect(put).toHaveBeenLastCalledWith(2, '7', 1084, 'main_tank', []);
    expect(raid()?.title).toBe('saved');
    expect(actions.busySlot).toBeNull();
  });

  it('clears through the delete', async () => {
    const del = vi.fn(() => Promise.resolve(plannedRaid({ title: 'cleared' })));
    const { actions, raid } = setup({ deleteAssignment: del });
    expect(await actions.clear(1084, s)).toBe(true);
    expect(del).toHaveBeenCalledWith(2, '7', 1084, 'main_tank');
    expect(raid()?.title).toBe('cleared');
  });

  it('keeps the API’s field line under the slot, else the generic one', async () => {
    const field = 'Only characters signed In, Late or Bench can take a slot.';
    const refused = setup({
      putAssignment: () =>
        Promise.reject(new GuildApiError('bad', 422, 'invalid', { character_keys: field })),
    });
    expect(await refused.actions.add(1084, s, MAGE.character_key)).toBe(false);
    expect(refused.actions.slotLine).toEqual({ slot: slotId(1084, 'main_tank'), message: field });
    const failed = setup({ putAssignment: () => Promise.reject(new Error('down')) });
    await failed.actions.add(1084, s, MAGE.character_key);
    expect(failed.actions.slotLine?.message).toBe(assignmentsCopy.saveFailed);
  });

  it('posts once at a time and stamps posted_at', async () => {
    const { actions, raid } = setup();
    expect(await actions.post()).toBe(true);
    expect(actions.postAnswer).toBe(assignmentsCopy.posted);
    expect(raid()?.assignments?.posted_at).toBe('2026-12-10T20:00:00Z');
  });

  it('reads the post’s own refusals as the API says them', async () => {
    const said = 'You posted these a moment ago. Try again in a minute.';
    const { actions } = setup({
      postAssignments: () => Promise.reject(new GuildApiError(said, 429, 'assignments_throttled')),
    });
    expect(await actions.post()).toBe(false);
    expect(actions.postAnswer).toBe(said);
    const other = setup({ postAssignments: () => Promise.reject(new Error('down')) });
    await other.actions.post();
    expect(other.actions.postAnswer).toBe(assignmentsCopy.postFailed);
  });
});
