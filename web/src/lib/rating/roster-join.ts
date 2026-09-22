// web/src/lib/rating/roster-join.ts
// Ruling 3 (docs/superpowers/plans/2026-09-21-rating-web.md): the rating API's per-fight
// rows carry player_key/player_name, not this fight's own unit GUID, and the report island
// has no region/ruleset available to build a characterKey join with. Join by display name
// instead -- a raid roster's names are unique within one fight. Shared by RatingPanel.svelte
// and RatingTab.svelte (Ruling 12) rather than duplicated.
import { splitUnitName } from '../characters';
import type { RatingCardPlayer } from './types';

export interface RosterUnit {
  guid: string;
  name: string;
  class?: string;
}

export function guidForPlayer(player: RatingCardPlayer, roster: RosterUnit[]): string {
  const name = splitUnitName(player.player_name).name;
  return roster.find((unit) => splitUnitName(unit.name).name === name)?.guid ?? '';
}

export function classForPlayer(
  guid: string,
  fallbackClass: string | undefined,
  roster: RosterUnit[],
): string | undefined {
  if (guid === '') return fallbackClass;
  return roster.find((unit) => unit.guid === guid)?.class ?? fallbackClass;
}
