// web/src/lib/guild/planner-actions.svelte.test.ts
import { describe, expect, it, vi } from 'vitest';
import { plannedRaid } from '../../test-support/raid-fixtures';
import { GuildApiError, type GuildRaidView } from './api';
import type { NightActionHost } from './night-actions.svelte';
import { PlannerActions } from './planner-actions.svelte';
import type { RaidPageDeps } from './raid-page-types';

const ROUTE = { guild: { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' }, raidId: '7' } as const;
const NOW = new Date('2026-12-08T18:00:00Z');
const TARGET = { tanks: 2, healers: 8, melee: 10, ranged: 20 };

function setup(
  deps: Partial<Pick<RaidPageDeps, 'putCompTarget' | 'deleteCompTarget' | 'pingRaid' | 'postRaidNow'>> = {},
  start: GuildRaidView = plannedRaid(),
) {
  let raid: GuildRaidView | null = start;
  const host: NightActionHost = {
    raid: () => raid,
    write: (run) => run(2),
    applyRaid: (next) => (raid = next),
    refusalLine: (_thrown, _ended, failed) => failed,
  };
  const actions = new PlannerActions(
    host,
    {
      putCompTarget: deps.putCompTarget ?? (() => Promise.resolve(plannedRaid({ title: 'saved' }))),
      deleteCompTarget: deps.deleteCompTarget ?? (() => Promise.resolve(plannedRaid({ title: 'reset' }))),
      pingRaid:
        deps.pingRaid ??
        (() =>
          Promise.resolve({
            posted: true,
            mentioned: 3,
            no_discord: 1,
            no_discord_names: ['Voidcall'],
            next_allowed_at: '2026-12-09T00:00:00Z',
          })),
      postRaidNow: deps.postRaidNow ?? (() => Promise.resolve(plannedRaid({ title: 'posted' }))),
      now: () => NOW,
    },
    ROUTE,
  );
  return { actions, raid: () => raid };
}

describe('PlannerActions', () => {
  it('saves the target and applies the answered raid', async () => {
    const put = vi.fn(() => Promise.resolve(plannedRaid({ title: 'saved' })));
    const { actions, raid } = setup({ putCompTarget: put });
    expect(await actions.saveComp(TARGET)).toBe(true);
    expect(put).toHaveBeenCalledWith(2, '7', TARGET);
    expect(raid()?.title).toBe('saved');
    expect(actions.compBusy).toBeNull();
  });

  it('keeps a 422 per field, and any other refusal as one line', async () => {
    const fields = { total: "That adds up to 41. Onyxia's Lair takes 40." };
    const refused = setup({
      putCompTarget: () => Promise.reject(new GuildApiError('bad', 422, 'invalid', fields)),
    });
    expect(await refused.actions.saveComp(TARGET)).toBe(false);
    expect(refused.actions.compFields).toEqual(fields);
    expect(refused.actions.compLine).toBeNull();

    const failed = setup({ deleteCompTarget: () => Promise.reject(new Error('down')) });
    expect(await failed.actions.resetComp()).toBe(false);
    expect(failed.actions.compLine).toBe('The target did not save. Try again.');
  });

  it('shows the browser’s own field lines', () => {
    const { actions } = setup();
    actions.refuseComp({ tanks: 'Use a whole number, 0 or more.' });
    expect(actions.compFields).toEqual({ tanks: 'Use a whole number, 0 or more.' });
  });

  it('pings, says who was named, and moves the throttle', async () => {
    const { actions, raid } = setup();
    expect(await actions.ping()).toBe(true);
    expect(actions.pingAnswer).toBe('Pinged 3 raiders on Discord. No Discord linked: Voidcall.');
    expect(raid()?.ping).toEqual({
      last_pinged_at: NOW.toISOString(),
      next_allowed_at: '2026-12-09T00:00:00Z',
    });
  });

  it('says a failed ping failed', async () => {
    const { actions } = setup({ pingRaid: () => Promise.reject(new Error('down')) });
    expect(await actions.ping()).toBe(false);
    expect(actions.pingAnswer).toBe('The ping did not go out. Try again.');
    expect(actions.pinging).toBe(false);
  });

  it('posts an unposted raid now and says Posted', async () => {
    const { actions, raid } = setup();
    expect(await actions.postNow()).toBe(true);
    expect(raid()?.title).toBe('posted');
    expect(actions.postAnswer).toBe('Posted.');
    expect(actions.posting).toBe(false);
  });

  it('says Reposted for a raid that had posted', async () => {
    const posted = plannedRaid({
      discord_post: {
        posted_at: '2026-12-01T00:00:00Z',
        post_at: '2026-12-01T00:00:00Z',
        reminded_at: null,
        remind: true,
      },
    });
    const { actions } = setup({}, posted);
    expect(await actions.postNow()).toBe(true);
    expect(actions.postAnswer).toBe('Reposted as a new message.');
  });

  it('names Settings when the guild has no channel, else says it failed', async () => {
    const noChannel = new GuildApiError('connect a Discord channel first', 409, 'no_channel');
    const { actions } = setup({ postRaidNow: () => Promise.reject(noChannel) });
    expect(await actions.postNow()).toBe(false);
    expect(actions.postAnswer).toBe('Pick a Discord channel in Settings first.');
    const down = setup({ postRaidNow: () => Promise.reject(new Error('down')) });
    expect(await down.actions.postNow()).toBe(false);
    expect(down.actions.postAnswer).toBe('The post did not go out. Try again.');
  });
});
