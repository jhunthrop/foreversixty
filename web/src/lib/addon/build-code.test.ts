import { describe, expect, it } from 'vitest';
import { addonCodeFor, buildCodeFor } from './build-code';
import { decodeFSB1 } from './fsb1';
import { indexTalents } from '../planner/rules';
import type { Item, TalentFile } from '../planner/types';
import type { GearSlot } from '../sim/types';

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

// A second layout for the sim path: a tier-2 talent listed FIRST in its tree's file order
// (the order the engine's digits follow), so a test can tell file order apart from the
// tier-then-column order the code has to carry, and a second tree with points in it, so
// the trees' own position order shows too.
const simTalents: TalentFile = {
  ...talents,
  trees: [
    {
      ...talents.trees[0],
      talents: [
        {
          id: 12,
          name: 'C',
          icon: 'i',
          max_rank: 1,
          tier: 1,
          column: 0,
          prereq_talent_id: null,
          prereq_rank: null,
          ranks: [],
        },
        ...talents.trees[0].talents,
      ],
    },
    {
      ...talents.trees[1],
      talents: [
        {
          id: 20,
          name: 'D',
          icon: 'i',
          max_rank: 3,
          tier: 0,
          column: 1,
          prereq_talent_id: null,
          prereq_rank: null,
          ranks: [],
        },
      ],
    },
    talents.trees[2],
  ],
} as unknown as TalentFile;

describe('buildCodeFor: a sim result', () => {
  const simInput = (talentString: string, gear: GearSlot[], items: Map<number, Item>) =>
    ({
      kind: 'sim',
      dataBuild: '1.60.1.69893',
      classSlug: 'paladin',
      talents: talentString,
      gear,
      talentIndex: indexTalents(simTalents),
      items,
    }) as const;

  it('turns the engine’s rank string into triples, tier then column, trees in position order', () => {
    // Holy's digits follow file order: C (tier 2) = 1, A = 2, B = 5. Protection: D = 2.
    const code = buildCodeFor(simInput('125-2-', [], new Map()));
    const result = decodeFSB1(code);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    const a = { tab: 1, tier: 1, column: 1 };
    const b = { tab: 1, tier: 1, column: 2 };
    const c = { tab: 1, tier: 2, column: 1 };
    const d = { tab: 2, tier: 1, column: 2 };
    expect(result.build.order).toEqual([a, a, b, b, b, b, b, c, d, d]);
    expect(result.build.classSlug).toBe('paladin');
    expect(result.build.dataBuild).toBe('1.60.1.69893');
  });

  it('carries the winning set’s gear with each item’s own stats, not the sim’s numbers', () => {
    const code = buildCodeFor(
      simInput(
        '02',
        [
          { slot: 'finger1', item_id: 3235 },
          { slot: 'head', item_id: 12640, enchant: 2583 },
        ],
        new Map([
          [12640, item(12640, { stamina: 17, spell_power: 23 }, 400)],
          [3235, item(3235, { stamina: 4, spirit: -3 })],
        ]),
      ),
    );
    const result = decodeFSB1(code);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.build.order).toEqual([
      { tab: 1, tier: 1, column: 1 },
      { tab: 1, tier: 1, column: 1 },
    ]);
    // Slot order is the planner's SLOTS order, the same one `addonCodeFor` writes, so one
    // build is one string whichever surface copied it; the enchant does not travel.
    expect(result.build.gear).toEqual([
      { slot: 'head', itemId: 12640, stats: { armor: 400, spell_power: 23, stamina: 17 } },
      { slot: 'finger1', itemId: 3235, stats: { spirit: -3, stamina: 4 } },
    ]);
  });

  it('leaves out a slot whose item record the page does not have', () => {
    const code = buildCodeFor(
      simInput(
        '',
        [
          { slot: 'head', item_id: 9999 },
          { slot: 'chest', item_id: 1 },
        ],
        new Map([[1, item(1, { stamina: 5 })]]),
      ),
    );
    expect(code).toBe('FSB1:1.60.1.69893:paladin::chest=1:stamina=5');
  });

  it('is fine with an empty talent tree: the gear still encodes', () => {
    const code = buildCodeFor(
      simInput('--', [{ slot: 'head', item_id: 1 }], new Map([[1, item(1, { stamina: 5 })]])),
    );
    const result = decodeFSB1(code);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.build.order).toEqual([]);
    expect(result.build.gear).toHaveLength(1);
  });

  it('is empty when neither the talents nor any slot encodes, so the button disables', () => {
    expect(buildCodeFor(simInput('--', [{ slot: 'head', item_id: 9999 }], new Map()))).toBe('');
  });

  it('takes the planner path for a planner build, identical to addonCodeFor', () => {
    const input = {
      dataBuild: '1',
      classSlug: 'paladin',
      order: [10, 11],
      gear: { head: 1 },
      talents: indexTalents(talents),
      items: new Map([[1, item(1, { stamina: 5 })]]),
    };
    expect(buildCodeFor(input)).toBe(addonCodeFor(input));
  });
});
