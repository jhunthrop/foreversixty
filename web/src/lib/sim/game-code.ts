// web/src/lib/sim/game-code.ts
// A saved sim result as the FSB1 code "Copy for the game" copies (SavedSim.svelte). Lifted
// out of the component so the bulk-versus-run choice is tested directly rather than through
// a page whose talent and item files only arrive after mount.
import { buildCodeFor } from '../addon/build-code';
import type { TalentIndex } from '../planner/rules';
import type { Item } from '../planner/types';
import { requestKind, type BulkResult } from './bulk-types';
import { winningGear, winningTalents } from './combos';
import type { SimResult } from './types';

/**
 * A bulk run (gear, talents, drops) copies its winning set: the leader's gear and talents,
 * the set its own table ranks first. A single run or a weights run copies the stored
 * request's own character. `''` (the button disabled) until the talent layout is in hand.
 */
export function gameCodeForSimResult(
  result: SimResult,
  dataBuild: string,
  talentIndex: TalentIndex | null,
  items: ReadonlyMap<number, Item>,
): string {
  if (talentIndex === null) return '';
  const kind = requestKind(result.request);
  const bulk = kind !== 'run' && kind !== 'weights' ? (result as BulkResult) : null;
  return buildCodeFor({
    kind: 'sim',
    dataBuild,
    classSlug: result.request.character.class,
    talents: bulk === null ? result.request.character.talents : winningTalents(bulk),
    gear: bulk === null ? result.request.character.gear : winningGear(bulk),
    talentIndex,
    items,
  });
}
