// web/src/lib/guild/loot-view.ts
// Guild control-centre spec §4.F: the loot council helper. The API (contract's
// GET .../loot) already ranks candidates and names the awardee; this module only formats
// what it sent -- never re-ranks or re-decides an award client-side.
import type { GuildLootEncounter, GuildLootItem } from './api';

/** Spec §6's boss-picker copy: a killed boss with nothing left unkilled reads "next
 *  unkilled: none, farm" (today's real case, Onyxia); an unkilled boss just names itself. */
export function bossPickerLabel(selected: GuildLootEncounter | undefined): string {
  if (selected === undefined) return 'No bosses killed yet this tier';
  return selected.killed
    ? `${selected.name} · next unkilled: none, farm`
    : `${selected.name} · not yet killed`;
}

/** The first encounter in the list the API has not marked killed -- the boss picker's own
 *  default (spec §4.F); `undefined` when every listed encounter is already killed, same
 *  meaning as `bossPickerLabel`'s "none, farm" branch. */
export function nextUnkilledEncounter(
  encounters: readonly GuildLootEncounter[],
): GuildLootEncounter | undefined {
  return encounters.find((encounter) => !encounter.killed);
}

export type LootCandidateAction =
  | { kind: 'awarded' }
  | { kind: 'awarded-to-other'; label: string }
  | { kind: 'equivalent' }
  | { kind: 'award' };

/**
 * One candidate row's own action cell (spec §12.1's round-2 fix): the awardee -- the
 * ranking's own top candidate, read off `item.awarded_to` -- carries the green "Awarded"
 * tag with no button; every other candidate on an awarded item shows a muted "Awarded to
 * {name}" note and no button; an item with no `already_equivalent` candidate and not yet
 * awarded keeps the `Award` button.
 */
export function candidateAction(item: GuildLootItem, candidateKey: string): LootCandidateAction {
  if (item.awarded_to !== null) {
    return item.awarded_to.character_key === candidateKey
      ? { kind: 'awarded' }
      : { kind: 'awarded-to-other', label: `Awarded to ${item.awarded_to.name}` };
  }
  const candidate = item.candidates.find((row) => row.character_key === candidateKey);
  if (candidate?.already_equivalent === true) return { kind: 'equivalent' };
  return { kind: 'award' };
}
