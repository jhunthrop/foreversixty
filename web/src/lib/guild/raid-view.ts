// web/src/lib/guild/raid-view.ts
// What a raid night counts, in the words a raid leader uses. The API's `kill_count` is every
// fight that ended in a kill, a whelp pack as much as Onyxia herself, so the page counts bosses
// and says the trash separately: "1 boss (Onyxia) · 2 kills · 2 trash packs · 0 wipes". The boss
// count is distinct bosses, "kills" the boss pulls that ended in one. Which
// encounters are bosses is the progression endpoint's tier list; with it unavailable the page
// falls back to the API's own figure under its old name.
import type { GuildProgressionPage, GuildRaidRow } from './api';
import { counted } from '../plural';

/** The ids of the tier's bosses, from the progression page; `null` while it is unknown. */
export function bossIdsOf(page: GuildProgressionPage | null): ReadonlySet<number> | null {
  return page === null ? null : new Set(page.encounters.map((encounter) => encounter.encounter_id));
}

export interface RaidTally {
  /** Distinct boss names killed, in the order they fell. */
  bosses: string[];
  bossKills: number;
  trashPacks: number;
  wipes: number;
}

export function raidTally(row: GuildRaidRow, bossIds: ReadonlySet<number>): RaidTally {
  const bosses: string[] = [];
  let bossKills = 0;
  let trashPacks = 0;
  for (const fight of row.fights) {
    if (!fight.kill || fight.name === null) continue;
    if (fight.encounter_id !== null && bossIds.has(fight.encounter_id)) {
      bossKills += 1;
      if (!bosses.includes(fight.name)) bosses.push(fight.name);
    } else {
      trashPacks += 1;
    }
  }
  return { bosses, bossKills, trashPacks, wipes: row.wipe_count };
}

/** "1 boss (Onyxia) · 2 kills · 2 trash packs · 0 wipes"; trash is left out when there was none. */
export function raidTallyLine(row: GuildRaidRow, bossIds: ReadonlySet<number> | null): string {
  if (bossIds === null) {
    return `${counted(row.kill_count, 'kill')} · ${counted(row.wipe_count, 'wipe')}`;
  }
  const tally = raidTally(row, bossIds);
  const bosses = counted(tally.bosses.length, 'boss', 'bosses');
  return [
    tally.bosses.length > 0 ? `${bosses} (${tally.bosses.join(', ')})` : bosses,
    tally.bossKills > 0 ? counted(tally.bossKills, 'kill') : '',
    tally.trashPacks > 0 ? counted(tally.trashPacks, 'trash pack') : '',
    counted(tally.wipes, 'wipe'),
  ]
    .filter((part) => part !== '')
    .join(' · ');
}

/** The same figures from a report row that carries only the two counts (the public list). */
export function reportTallyLine(killCount: number, wipeCount: number): string {
  return `${counted(killCount, 'kill')} · ${counted(wipeCount, 'wipe')}`;
}
