// web/src/lib/sim/extra-attacks.test.ts
// Spec 2026-11-04-sim-results 3.6: where the Extra attacks came from.
import { describe, expect, it } from 'vitest';
import full from '../../fixtures/sim/result-full.json';
import type { BuffNames } from './buff-names';
import { extraAttackLines } from './extra-attacks';
import type { SimResult } from './types';

const result = full as unknown as SimResult;
const names: BuffNames = { entries: { windfury_totem: { name: 'Windfury Totem', icon: '' } } };
const itemName = (id: number): string | undefined => (id === 11815 ? 'Hand of Justice' : undefined);
const actionName = (key: string): string => (key === 'spell:15600' ? 'Hand of Justice' : 'An unnamed spell');

describe('extraAttackLines', () => {
  it('without the engine field, names the request sources with no counts', () => {
    const { extra_attack_sources: _dropped, ...older } = result;
    expect(extraAttackLines(older as SimResult, names, itemName, actionName)).toEqual([
      { name: 'Windfury Totem', count: null },
    ]);
  });

  it('reads the fixture: 21.4 swings from the totem', () => {
    expect(extraAttackLines(result, names, itemName, actionName)).toEqual([
      { name: 'Windfury Totem', count: 21 },
    ]);
  });

  it('names a proc by its action key', () => {
    const proc: SimResult = { ...result, extra_attack_sources: [{ source: 'spell:15600', count: 5.6 }] };
    expect(extraAttackLines(proc, names, itemName, actionName)).toEqual([
      { name: 'Hand of Justice', count: 6 },
    ]);
  });

  it('with the engine field, carries whole counts', () => {
    const counted: SimResult = {
      ...result,
      extra_attack_sources: [
        { source: 'item:11815', count: 6.2 },
        { source: 'windfury_totem', count: 21 },
      ],
    };
    expect(extraAttackLines(counted, names, itemName, actionName)).toEqual([
      { name: 'Hand of Justice', count: 6 },
      { name: 'Windfury Totem', count: 21 },
    ]);
  });

  it('names a worn extra-attack item from the request', () => {
    const character = {
      ...result.request.character,
      buffs: [],
      gear: [{ slot: 'trinket1', item_id: 11815 }],
    };
    const { extra_attack_sources: _dropped, ...older } = result;
    const worn = { ...older, request: { ...result.request, character } } as SimResult;
    expect(extraAttackLines(worn, names, itemName, actionName)).toEqual([
      { name: 'Hand of Justice', count: null },
    ]);
  });
});
