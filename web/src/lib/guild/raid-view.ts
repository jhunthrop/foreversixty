// web/src/lib/guild/raid-view.ts
// What a raid night counts, in the words a raid leader uses. The API's `kill_count` is every
// fight that ended in a kill, a whelp pack as much as Onyxia herself, so the page counts bosses
// and says the trash separately: "1 boss (Onyxia) · 2 kills · 2 trash packs · 0 wipes". The boss
// count is distinct bosses, "kills" the boss pulls that ended in one. Which
// encounters are bosses is the progression endpoint's tier list; with it unavailable the page
// falls back to the API's own figure under its old name.
import type { GuildProgressionPage, GuildRaidFight, GuildRaidRow } from './api';
import { counted } from '../plural';
import { bossCountLabel, distinctBossNames } from '../boss-count';

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

/** A named fight that is not one of the tier's bosses: a whelp pack, never a boss pull. */
function isTrash(fight: GuildRaidFight, bossIds: ReadonlySet<number>): boolean {
  return fight.name !== null && (fight.encounter_id === null || !bossIds.has(fight.encounter_id));
}

export function raidTally(row: GuildRaidRow, bossIds: ReadonlySet<number>): RaidTally {
  const bosses = distinctBossNames(row.fights, bossIds);
  let bossKills = 0;
  let trashPacks = 0;
  let trashWipes = 0;
  for (const fight of row.fights) {
    if (isTrash(fight, bossIds)) {
      if (fight.kill) trashPacks += 1;
      else trashWipes += 1;
    } else if (fight.kill && fight.name !== null) bossKills += 1;
  }
  // The API's wipe count is every fight that did not end in a kill; a trash pack that went
  // badly is not a wipe on the night's line (guild round 2, WP-12).
  return { bosses, bossKills, trashPacks, wipes: Math.max(0, row.wipe_count - trashWipes) };
}

/** One fight of an open night as the Raids tab draws it (guild round 2, WP-12). */
export interface NightFightView {
  index: number;
  /** `Onyxia · pull 2` (pulls counted per boss), the pack's own name, or `Pull 3` unnamed. */
  label: string;
  /** A trash pack: quiet, no pull number and no Kill or Wipe badge. */
  trash: boolean;
  /** No named encounter: the API could not name the fight. */
  unnamed: boolean;
  kill: boolean;
  duration_ms: number;
}

/** The quiet word a trash row carries in place of a Kill or Wipe badge. */
export const TRASH_LABEL = 'trash';

/** The night's fights: boss pulls numbered per boss, trash packs named and marked as trash. While
 *  the tier's boss list is unknown every named fight reads as a pull, numbered per name. */
export function nightFights(row: GuildRaidRow, bossIds: ReadonlySet<number> | null): NightFightView[] {
  const pulls = new Map<string, number>();
  return row.fights.map((fight) => {
    const base = { index: fight.index, kill: fight.kill, duration_ms: fight.duration_ms };
    if (fight.name === null)
      return { ...base, label: `Pull ${fight.index + 1}`, trash: false, unnamed: true };
    if (bossIds !== null && isTrash(fight, bossIds)) {
      return { ...base, label: fight.name, trash: true, unnamed: false };
    }
    const pull = (pulls.get(fight.name) ?? 0) + 1;
    pulls.set(fight.name, pull);
    return { ...base, label: `${fight.name} · pull ${pull}`, trash: false, unnamed: false };
  });
}

/** The facts of a night as separate pieces, so a narrow screen can wrap them whole:
 *  ["1 boss (Onyxia)", "1 kill", "2 trash packs", "1 wipe"]; trash is left out when there was none. */
export function raidTallyParts(row: GuildRaidRow, bossIds: ReadonlySet<number> | null): string[] {
  if (bossIds === null) {
    return [counted(row.kill_count, 'kill'), counted(row.wipe_count, 'wipe')];
  }
  const tally = raidTally(row, bossIds);
  const bosses = bossCountLabel(tally.bosses.length);
  return [
    tally.bosses.length > 0 ? `${bosses} (${tally.bosses.join(', ')})` : bosses,
    tally.bossKills > 0 ? counted(tally.bossKills, 'kill') : '',
    tally.trashPacks > 0 ? counted(tally.trashPacks, 'trash pack') : '',
    counted(tally.wipes, 'wipe'),
  ].filter((part) => part !== '');
}

/** "1 boss (Onyxia) · 1 kill · 2 trash packs · 1 wipe". */
export function raidTallyLine(row: GuildRaidRow, bossIds: ReadonlySet<number> | null): string {
  return raidTallyParts(row, bossIds).join(' · ');
}

/** The same figures from a report row that carries only the two counts (the public list). */
export function reportTallyLine(killCount: number, wipeCount: number): string {
  return `${counted(killCount, 'kill')} · ${counted(wipeCount, 'wipe')}`;
}
