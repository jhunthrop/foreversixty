import { describe, expect, it } from 'vitest';
import { addonCodeFor } from './build-code';
import { decodeFSB1 } from './fsb1';
import { indexTalents } from '../planner/rules';
import type { Item, TalentFile } from '../planner/types';

const talents: TalentFile = {
  build: '1.60.1.69893',
  class_id: 2,
  class_slug: 'paladin',
  trees: [
    {
      id: 1,
      name: 'Holy',
      position: 0,
      talents: [
        {
          id: 10,
          name: 'A',
          icon: 'i',
          max_rank: 2,
          tier: 0,
          column: 0,
          prereq_talent_id: null,
          prereq_rank: null,
          ranks: [],
        },
        {
          id: 11,
          name: 'B',
          icon: 'i',
          max_rank: 5,
          tier: 0,
          column: 1,
          prereq_talent_id: null,
          prereq_rank: null,
          ranks: [],
        },
      ],
    },
    { id: 2, name: 'Protection', position: 1, talents: [] },
    { id: 3, name: 'Retribution', position: 2, talents: [] },
  ],
} as unknown as TalentFile;

const item = (id: number, stats: Record<string, number>, armor = 0): Item =>
  ({
    id,
    name: `item ${id}`,
    icon: 'i',
    slot: 'head',
    quality: 4,
    required_level: 0,
    item_level: 60,
    armor,
    stats,
    set_id: null,
    unique: false,
  }) as unknown as Item;

describe('addonCodeFor', () => {
  it('turns the planner’s talent id order into 1-based tab/tier/column triples', () => {
    const code = addonCodeFor({
      dataBuild: '1.60.1.69893',
      classSlug: 'paladin',
      order: [10, 11, 10],
      gear: {},
      talents: indexTalents(talents),
      items: new Map(),
    });
    const result = decodeFSB1(code);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.build.order).toEqual([
      { tab: 1, tier: 1, column: 1 },
      { tab: 1, tier: 1, column: 2 },
      { tab: 1, tier: 1, column: 1 },
    ]);
  });

  it('carries each equipped item’s stats, armour included', () => {
    const code = addonCodeFor({
      dataBuild: '1.60.1.69893',
      classSlug: 'paladin',
      order: [],
      gear: { head: 12640 },
      talents: indexTalents(talents),
      items: new Map([[12640, item(12640, { stamina: 17, spell_power: 23 }, 400)]]),
    });
    const result = decodeFSB1(code);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.build.gear).toEqual([
      { slot: 'head', itemId: 12640, stats: { armor: 400, spell_power: 23, stamina: 17 } },
    ]);
  });

  it('maps the planner’s own stat names onto the contract vocabulary', () => {
    // The item files say `healing` and `fire_res`; Data.lua's weights say
    // `healing_power` and `fire_resistance`. The addon must get the second.
    const code = addonCodeFor({
      dataBuild: '1',
      classSlug: 'paladin',
      order: [],
      gear: { head: 1 },
      talents: indexTalents(talents),
      items: new Map([[1, item(1, { healing: 40, fire_res: 10 })]]),
    });
    expect(code).toContain('head=1:fire_resistance=10;healing_power=40');
  });

  it('leaves out a slot whose item is not in the index rather than writing a bare id', () => {
    const code = addonCodeFor({
      dataBuild: '1',
      classSlug: 'paladin',
      order: [],
      gear: { head: 9999 },
      talents: indexTalents(talents),
      items: new Map(),
    });
    expect(code).toBe('FSB1:1:paladin::');
  });

  it('round-trips an item with a negative stat, sign intact', () => {
    // Ring of Scorn (spirit -3) and 42 other negative stat values are real equippable
    // items. The encoder copies item.stats through verbatim, so the decoder has to read
    // what it wrote: a player wearing one must not get a code the addon refuses, and the
    // penalty must not be quietly dropped on the way either.
    const code = addonCodeFor({
      dataBuild: '1.60.1.69893',
      classSlug: 'warrior',
      order: [],
      gear: { finger1: 3235 },
      talents: indexTalents(talents),
      items: new Map([[3235, item(3235, { stamina: 4, spirit: -3 })]]),
    });
    // `spirit` leads `stamina` because a gear entry's stats go on the wire in name order
    // ("spi" < "sta"), not in contract 10.8's vocabulary order and not in the order the
    // item's own stats table happens to list them.
    expect(code).toBe('FSB1:1.60.1.69893:warrior::finger1=3235:spirit=-3;stamina=4');
    const result = decodeFSB1(code);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.build.gear).toEqual([{ slot: 'finger1', itemId: 3235, stats: { stamina: 4, spirit: -3 } }]);
  });

  it('omits a zero armour value rather than writing armor=0', () => {
    const code = addonCodeFor({
      dataBuild: '1',
      classSlug: 'paladin',
      order: [],
      gear: { head: 1 },
      talents: indexTalents(talents),
      items: new Map([[1, item(1, { stamina: 5 }, 0)]]),
    });
    expect(code).toBe('FSB1:1:paladin::head=1:stamina=5');
  });
});
