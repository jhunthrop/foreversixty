// web/src/lib/report/kill-counts.ts
// How many mobs died in a pull and how many of them the logger's group killed. Engine 0.6.0
// counts both (logs/engine/fight/counts.go); a fight indexed before then carries only the
// retired npc_kills, which counted every group kill twice, so it has no figure here and the
// page shows no kill count for it rather than a wrong one.

import type { FightEntry } from './types';

export interface KillCounts {
  mobDeaths: number;
  groupKills: number;
}

type Counted = Partial<Pick<FightEntry, 'mob_deaths' | 'group_kills'>>;

const isCount = (value: number | undefined): value is number =>
  typeof value === 'number' && Number.isFinite(value) && value >= 0;

/** A fight's kill counts, or null when its line predates them. */
export function killCounts(fight: Counted): KillCounts | null {
  if (!isCount(fight.mob_deaths) || !isCount(fight.group_kills)) return null;
  return { mobDeaths: fight.mob_deaths, groupKills: fight.group_kills };
}

/** A run of pulls' counts summed, or null when any pull has none: never a partial total. */
export function sumKillCounts(fights: readonly Counted[]): KillCounts | null {
  if (fights.length === 0) return null;
  const total: KillCounts = { mobDeaths: 0, groupKills: 0 };
  for (const fight of fights) {
    const counts = killCounts(fight);
    if (counts === null) return null;
    total.mobDeaths += counts.mobDeaths;
    total.groupKills += counts.groupKills;
  }
  return total;
}

/** "177 mobs died · your group killed 26"; "No mobs died" when none did. */
export function killCountsLabel({ mobDeaths, groupKills }: KillCounts): string {
  if (mobDeaths === 0) return 'No mobs died';
  return `${mobDeaths} ${mobDeaths === 1 ? 'mob' : 'mobs'} died · your group killed ${groupKills}`;
}
