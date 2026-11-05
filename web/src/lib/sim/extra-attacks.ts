// web/src/lib/sim/extra-attacks.ts
// Where a run's Extra attacks came from (spec 2026-11-04-sim-results 3.6). With the engine's own
// `extra_attack_sources` the sub-line carries counts; a result saved before that field names the
// sources from its request, with no counts. The worn items are the ones whose client effect is an
// extra attack.
import { buffLabel, type BuffNames } from './buff-names';
import type { SimResult } from './types';

export const EXTRA_ATTACK_BUFF = 'windfury_totem';
export const EXTRA_ATTACK_ITEMS: readonly number[] = [11815, 11684, 17705, 871];

export interface ExtraAttackSourceLine {
  name: string;
  /** Whole extra attacks per fight; null when the result does not say. */
  count: number | null;
}

const ITEM_SOURCE = /^item:(\d+)$/;
/** A source that is a proc's action key rather than a simbuffs id. */
const SPELL_SOURCE = /^spell:\d+/;

export function extraAttackLines(
  result: SimResult,
  buffNames: BuffNames | null,
  itemName: (itemId: number) => string | undefined,
  actionName: (key: string) => string,
): ExtraAttackSourceLine[] {
  const named = (source: string): string | undefined => {
    if (SPELL_SOURCE.test(source)) return actionName(source);
    const item = ITEM_SOURCE.exec(source);
    return item === null ? buffLabel(source, buffNames) : itemName(Number(item[1]));
  };
  if (result.extra_attack_sources !== undefined) {
    return result.extra_attack_sources.flatMap(({ source, count }) => {
      const name = named(source);
      return name === undefined ? [] : [{ name, count: Math.round(count) }];
    });
  }
  const { buffs, gear } = result.request.character;
  const fromBuff = buffs.some((buff) => buff.split(':')[0] === EXTRA_ATTACK_BUFF)
    ? [{ name: buffLabel(EXTRA_ATTACK_BUFF, buffNames), count: null }]
    : [];
  const fromGear = gear.flatMap((slot) => {
    if (!EXTRA_ATTACK_ITEMS.includes(slot.item_id)) return [];
    const name = itemName(slot.item_id);
    return name === undefined ? [] : [{ name, count: null }];
  });
  return [...fromGear, ...fromBuff];
}
