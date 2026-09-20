// web/src/lib/sim/combos.test.ts
import { describe, expect, it } from 'vitest';
import bulkResultJson from '../../fixtures/sim/bulk-result.json';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import setsJson from '../../fixtures/planner/sets.json';
import { addonStringFor } from './addon-export';
import {
  comboRows,
  deltaLabel,
  headlineFor,
  isEmptiedOffHand,
  keepsSetBonus,
  percentOf,
  slotSummary,
  sourceNameOfCombo,
  substitutionLabel,
  winningGear,
} from './combos';
import { bulkCopy } from './copy';
import type { BulkResult } from './bulk-types';
import type { Item, ItemSet } from '../planner/types';

const result = bulkResultJson as unknown as BulkResult;
const items = new Map((itemsJson.items as unknown as Item[]).map((item) => [item.id, item]));
const sets = setsJson as unknown as ItemSet[];

describe('comboRows', () => {
  it('ranks best first and gives every member of a within-error group the same rank', () => {
    const rows = comboRows(result);
    // The fixture's 7 combos carry groups 0,0,1,1,2,2,2 (task-2 fix round 1 added a 7th
    // combo to the trailing group-2 tie). Rank is the 1-based index of a group's first
    // member, so the group-2 trio all share rank 5.
    expect(rows.map((row) => row.rank)).toEqual([1, 1, 3, 3, 5, 5, 5]);
    expect(rows[0].withinError).toBe(true);
    expect(rows[2].withinError).toBe(false);
  });

  it('carries the percent against the equipped set', () => {
    const rows = comboRows(result);
    expect(rows[0].percent).toBeCloseTo((41.2 / 1461.2) * 100, 6);
  });
});

describe('percentOf and deltaLabel', () => {
  it('reads a gain with its error and a sign', () => {
    expect(deltaLabel({ mean: 41.2, stddev: 0, error: 5.41, min: 0, max: 0 })).toBe('+41 ± 11');
    expect(deltaLabel({ mean: -1.8, stddev: 0, error: 5.38, min: 0, max: 0 })).toBe('−2 ± 11');
  });

  it('is zero percent against a zero baseline rather than infinite', () => {
    expect(percentOf(41.2, 0)).toBe(0);
  });
});

/**
 * Engine-lane rule 5's two-hander: the winner replaces main-hand and off-hand with one
 * weapon, which the engine reports as the new main hand PLUS
 * `{kind:"item", slot:"off_hand", item_id:0}`. The shared fixture has no off-hand, so the
 * case is built here from it rather than changed in a file eight other tests read.
 */
const twoHanded: BulkResult = {
  ...result,
  request: {
    ...result.request,
    character: {
      ...result.request.character,
      gear: [...result.request.character.gear, { slot: 'off_hand', item_id: 11684 }],
    },
  },
  combos: [
    {
      ...result.combos[0],
      substitutions: [
        { kind: 'item', slot: 'main_hand', item_id: 17182, name: 'Sulfuras' },
        { kind: 'item', slot: 'off_hand', item_id: 0, name: '<item removed>' },
      ],
    },
    ...result.combos.slice(1),
  ],
};

describe('winningGear', () => {
  it('writes the leader’s substitutions over the base character’s gear', () => {
    const gear = winningGear(result);
    expect(gear.find((slot) => slot.slot === 'head')?.item_id).toBe(16963);
    expect(gear.find((slot) => slot.slot === 'shoulder')?.item_id).toBe(16966);
    // untouched by the winner
    expect(gear.find((slot) => slot.slot === 'main_hand')?.item_id).toBe(12784);
  });

  it('leaves the base character’s own gear untouched', () => {
    const before = JSON.stringify(result.request.character.gear);
    winningGear(result);
    expect(JSON.stringify(result.request.character.gear)).toBe(before);
  });

  it('drops the slot a two-hander emptied rather than writing item 0 over it', () => {
    const gear = winningGear(twoHanded);
    expect(gear.find((slot) => slot.slot === 'main_hand')?.item_id).toBe(17182);
    // Not `{item_id: 0}`: nothing downstream defines 0 as "empty", so the slot is gone.
    expect(gear.some((slot) => slot.slot === 'off_hand')).toBe(false);
  });

  it('exports an addon string with no off-hand entry at all', () => {
    const addon = addonStringFor({
      dataBuild: '1.60.1',
      classSlug: 'warrior',
      raceSlug: 'orc',
      talents: '0-5530515-0',
      gear: winningGear(twoHanded),
    });
    expect(addon).toContain('main_hand=17182');
    expect(addon).not.toContain('off_hand');
    expect(addon).not.toContain('=0');
  });
});

describe('isEmptiedOffHand', () => {
  it('is the rule-5 sentinel and nothing else', () => {
    expect(isEmptiedOffHand({ kind: 'item', slot: 'off_hand', item_id: 0 })).toBe(true);
    expect(isEmptiedOffHand({ kind: 'item', slot: 'off_hand', item_id: 11684 })).toBe(false);
    expect(isEmptiedOffHand({ kind: 'item', slot: 'main_hand', item_id: 0 })).toBe(false);
    expect(isEmptiedOffHand({ kind: 'talents', name: 'Deep Fury' })).toBe(false);
  });
});

describe('slotSummary', () => {
  it('names the winner’s item per slot and the gain that slot contributed alone', () => {
    const rows = slotSummary(result);
    const head = rows.find((row) => row.slot === 'head')!;
    expect(head.item_id).toBe(16963);
    expect(head.name).toBe('Helm of Wrath');
    // the single-substitution combo for that slot
    expect(head.gain).toBeCloseTo(38.6, 6);
    const shoulder = rows.find((row) => row.slot === 'shoulder')!;
    expect(shoulder.gain).toBeCloseTo(18.9, 6);
  });

  it('lists a slot the winner changed even when nothing measured it alone', () => {
    const trimmed: BulkResult = { ...result, combos: [result.combos[0]] };
    const rows = slotSummary(trimmed);
    expect(rows.map((row) => row.slot).sort()).toEqual(['head', 'shoulder']);
    expect(rows.every((row) => row.gain === null || typeof row.gain === 'number')).toBe(true);
  });

  // Engine-lane rule 5, the second leak (Task 16's re-review, routed to this task): a
  // two-hander replacing a main-plus-off-hand pair emits a second substitution
  // `{kind:"item", slot:"off_hand", item_id:0, name:"<item removed>"}`. SubstitutionChips
  // already renders this as an emptied slot; slotSummary fed the same substitution's raw
  // `name` straight through substitutionLabel, which prints "<item removed>" verbatim since
  // the field is non-empty. This is the "By slot" panel's own path onto the same sentinel.
  it('renders the emptied off-hand from a two-hander swap as an emptied slot, never the literal engine text', () => {
    const twoHander: BulkResult = {
      ...result,
      combos: [
        {
          substitutions: [
            { kind: 'item', slot: 'main_hand', item_id: 19351, name: 'Sulfuras, Hand of Ragnaros' },
            { kind: 'item', slot: 'off_hand', item_id: 0, name: '<item removed>' },
          ],
          dps: result.equipped,
          delta: result.equipped,
          group: 0,
        },
      ],
    };
    const rows = slotSummary(twoHander);
    const offHand = rows.find((row) => row.slot === 'off_hand')!;
    expect(offHand.name).toBe(bulkCopy.offHandEmptied);
    expect(offHand.name).not.toBe('<item removed>');
    const mainHand = rows.find((row) => row.slot === 'main_hand')!;
    expect(mainHand.name).toBe('Sulfuras, Hand of Ragnaros');
  });
});

describe('keepsSetBonus', () => {
  it('is false when the combination does not reach the piece count', () => {
    expect(keepsSetBonus(result.combos[0], result, items, sets, 4)).toBe(false);
  });

  it('is true at a piece count the combination does reach', () => {
    expect(keepsSetBonus(result.combos[0], result, items, sets, 2)).toBe(true);
  });
});

describe('headlineFor and substitutionLabel', () => {
  it('names the leader’s biggest change from the result’s own name field', () => {
    expect(headlineFor(result)).toBe('+41 DPS from Helm of Wrath');
  });

  it('labels an item, a loadout, a set and a consumable list (contract 10.8)', () => {
    expect(substitutionLabel({ kind: 'item', slot: 'head', item_id: 16963, name: 'Helm of Wrath' })).toBe(
      'Helm of Wrath',
    );
    expect(substitutionLabel({ kind: 'talents', name: 'Deep Fury' })).toBe('Deep Fury');
    expect(substitutionLabel({ kind: 'set', name: 'AQ set' })).toBe('AQ set');
    expect(substitutionLabel({ kind: 'consumes', name: 'flask_of_supreme_power' })).toBe(
      'With flask of supreme power',
    );
  });

  it('falls back to the item id only when the engine sent no name', () => {
    expect(substitutionLabel({ kind: 'item', slot: 'head', item_id: 16963 })).toBe('Item 16963');
  });
});

describe('sourceNameOfCombo', () => {
  it('reads the boss off the substitution the engine copied it onto (contract 10.1 A6)', () => {
    expect(
      sourceNameOfCombo({
        substitutions: [{ kind: 'item', item_id: 1, source_name: 'Ragnaros' }],
        dps: result.equipped,
        delta: result.equipped,
        group: 0,
      }),
    ).toBe('Ragnaros');
  });
});
