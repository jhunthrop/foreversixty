// web/src/lib/boss-count.ts
// How many bosses a night killed, said once for every page that says it (lifecycle round 8, P8-3):
// the guild pages and the account's report rows both count DISTINCT bosses, never kill pulls, so
// two Onyxia kills in one night read "1 boss", and a whelp pack is not a boss at all.
import { counted } from './plural';

/** The part of a fight these counts read. */
export interface BossFight {
  kill: boolean;
  name: string | null;
  encounter_id: number | null;
}

/**
 * The distinct bosses a night killed, by name, in the order they fell. A fight is a boss when it
 * ended in a kill and carries an encounter id (and, where the tier's boss list is known, that id
 * is on it); a trash pack has no encounter id.
 */
export function distinctBossNames(
  fights: readonly BossFight[],
  bossIds?: ReadonlySet<number> | null,
): string[] {
  const names: string[] = [];
  for (const fight of fights) {
    if (!fight.kill || fight.name === null || fight.encounter_id === null) continue;
    if (bossIds != null && !bossIds.has(fight.encounter_id)) continue;
    if (!names.includes(fight.name)) names.push(fight.name);
  }
  return names;
}

/** "1 boss", "2 bosses", "0 bosses". */
export function bossCountLabel(count: number): string {
  return counted(count, 'boss', 'bosses');
}
