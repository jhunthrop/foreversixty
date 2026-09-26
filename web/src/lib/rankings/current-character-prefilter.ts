// web/src/lib/rankings/current-character-prefilter.ts
// Spec 2026-09-25 section 4.2: "/rankings: pre-filtered to the character's class and
// ruleset, with the character's own row pinned at the top when it is ranked." The rankings
// API has no per-character/pin parameter and no per-character rank lookup (`fetchRankings`/
// `fetchGuildRankings` checked, neither takes one) -- adding one is an api/** change, out of
// this lane's ownership -- so the pin below only ever reorders a row already present in the
// page the normal filtered query returned. No invented rank, no second fetch.
import { parseCharacterKey } from '../characters';
import type { CurrentCharacter } from '../current-character';
import type { RankingRow } from './api';
import { type RankingsState } from './url';

/**
 * Pre-fills `class`/`ruleset` from the current pointer, but only on a URL that named
 * neither itself -- an explicit link always wins, the same rule `?code=` takes over a
 * restored planner pointer. Region is not filterable on the rankings board today
 * (`RankingsState` has no `region` fill from a pointer; the existing `region` filter stays
 * URL-only), so only class and ruleset are ever pre-filled. A non-armory pointer (a paste,
 * with no server-side ruleset) fills class alone.
 */
export function applyCurrentCharacterPrefilter(
  state: RankingsState,
  search: string,
  current: CurrentCharacter | null,
): RankingsState {
  if (current === null) return state;
  const params = new URLSearchParams(search);
  const next = { ...state };
  if (params.get('class') === null) next.class = current.classSlug;
  if (params.get('ruleset') === null && current.source === 'armory') {
    const key = parseCharacterKey(current.ref);
    if (key !== null) next.ruleset = key.ruleset;
  }
  return next;
}

/** True when the row's `player.key` names the current pointer's own armory character. */
function isCurrentCharacterRow(row: RankingRow, current: CurrentCharacter): boolean {
  return current.source === 'armory' && row.player.key === current.ref;
}

export function pinCurrentCharacterRow(rows: RankingRow[], current: CurrentCharacter | null): RankingRow[] {
  if (current === null) return rows;
  const index = rows.findIndex((row) => isCurrentCharacterRow(row, current));
  if (index <= 0) return rows;
  const pinned = rows[index];
  return [pinned, ...rows.slice(0, index), ...rows.slice(index + 1)];
}
