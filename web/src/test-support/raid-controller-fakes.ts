// web/src/test-support/raid-controller-fakes.ts
// Fake deps, a fake /v1/me and fake timers for the raid page controller's tests.
import { afterEach, beforeEach, vi } from 'vitest';
import type { Me } from '../lib/account/api';
import type { GuildHome, GuildZoneView } from '../lib/guild/api';
import { RaidPageController } from '../lib/guild/raid-page-controller.svelte';
import type { RaidPageDeps } from '../lib/guild/raid-page-types';
import type { GuildPage } from '../lib/rankings/api';
import { NOW, plannedRaid } from './raid-fixtures';

export const ROUTE = { guild: { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' }, raidId: '7' } as const;

export const GUILD_PAGE = {
  guild: { id: 2, name: 'Olympus XXVII', region: 'us', ruleset: 'pvp', faction: 'horde', crest_url: null },
  progression: [],
  roster_best: [],
  reports: [],
} as unknown as GuildPage;

export function me(guild: Partial<Me['guilds'][number]> | null): Me {
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

export function deferred<T>(): {
  promise: Promise<T>;
  resolve: (v: T) => void;
  reject: (e: unknown) => void;
} {
  let resolve!: (v: T) => void;
  let reject!: (e: unknown) => void;
  const promise = new Promise<T>((res, rej) => {
    resolve = res;
    reject = rej;
  });
  return { promise, resolve, reject };
}

const RAID_ZONE = { kind: 'raid', opens: null, min_level: null, max_level: null, key_item: null } as const;

export const ZONES: GuildZoneView[] = [
  { ...RAID_ZONE, key: 'raid:barrow-deeps', name: 'Barrow Deeps', size: null },
  { ...RAID_ZONE, key: 'raid:onyxias-lair', name: "Onyxia's Lair", size: 40 },
];

/** The document's visibility the fake deps report; each test starts visible. */
export const visibility = { visible: true };

export function deps(extra: Partial<RaidPageDeps> = {}): RaidPageDeps {
  return {
    loadGuild: () => Promise.resolve(GUILD_PAGE),
    loadMe: () => Promise.resolve(me({})),
    loadHome: () =>
      Promise.resolve({
        viewer: { role: 'member', character_key: 'us-pvp-kraggor', verified: true, officer_tools: true },
      } as unknown as GuildHome),
    fetchRaid: () => Promise.resolve(plannedRaid()),
    refetchRaid: () => Promise.resolve(plannedRaid()),
    saveSignup: () =>
      Promise.resolve({ raid: plannedRaid({ counts: { ...plannedRaid().counts, in: 22 } }), moved: null }),
    setStatus: () => Promise.resolve(plannedRaid({ status: 'locked' })),
    updateRaid: () => Promise.resolve(plannedRaid({ title: 'Ony farm' })),
    loadMatchCandidates: () => Promise.resolve({ zone_known: true, rows: [] }),
    putMatch: () => Promise.resolve(plannedRaid({ status: 'done' })),
    deleteMatch: () => Promise.resolve(plannedRaid()),
    putOverride: () => Promise.resolve(plannedRaid()),
    deleteOverride: () => Promise.resolve(plannedRaid()),
    postReport: () => Promise.resolve({ posted_at: '2026-12-11T05:58:00Z' }),
    putCompTarget: () => Promise.resolve(plannedRaid()),
    deleteCompTarget: () => Promise.resolve(plannedRaid()),
    pingRaid: () =>
      Promise.resolve({
        posted: true,
        mentioned: 0,
        no_discord: 0,
        no_discord_names: [],
        next_allowed_at: null,
      }),
    postRaidNow: () => Promise.resolve(plannedRaid()),
    putAssignment: () => Promise.resolve(plannedRaid()),
    deleteAssignment: () => Promise.resolve(plannedRaid()),
    postAssignments: () =>
      Promise.resolve({
        posted: true,
        posted_at: '2026-12-10T20:00:00Z',
        next_allowed_at: '2026-12-10T20:01:00Z',
      }),
    loadZones: () => Promise.resolve(ZONES),
    isVisible: () => visibility.visible,
    now: () => new Date(Date.now()),
    storage: null,
    ...extra,
  };
}

export async function started(d: RaidPageDeps): Promise<{ c: RaidPageController; stop: () => void }> {
  const c = new RaidPageController(d, ROUTE);
  const stop = c.start();
  await vi.advanceTimersByTimeAsync(0);
  return { c, stop };
}

/** Fake timers pinned to the fixtures' clock, for every test in the calling file. */
export function useControllerClock(): void {
  beforeEach(() => {
    vi.useFakeTimers();
    vi.setSystemTime(NOW);
    visibility.visible = true;
  });

  afterEach(() => {
    vi.useRealTimers();
  });
}
