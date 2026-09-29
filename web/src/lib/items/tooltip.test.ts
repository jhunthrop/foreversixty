// web/src/lib/items/tooltip.test.ts
import { describe, expect, it } from 'vitest';
import type { Item, ItemSet } from '../planner/types';
import type { LootFile } from '../sim/loot';
import { itemTooltipModel, type ItemTooltipSources } from './tooltip';

function item(overrides: Partial<Item> = {}): Item {
  return {
    id: 1,
    name: 'Test Item',
    icon: 'inv_test',
    slot: 'chest',
    quality: 4,
    required_level: 60,
    item_level: 76,
    armor: 400,
    stats: { stamina: 20, strength: 18 },
    set_id: null,
    unique: false,
    ...overrides,
  };
}

const NO_SOURCES: ItemTooltipSources = { loot: { sources: [] }, sets: [] };

describe('itemTooltipModel', () => {
  it('carries the plain fields straight off the row', () => {
    const model = itemTooltipModel(item(), NO_SOURCES);
    expect(model.id).toBe(1);
    expect(model.name).toBe('Test Item');
    expect(model.quality).toBe(4);
    expect(model.icon).toBe('inv_test');
    expect(model.itemLevel).toBe(76);
    expect(model.requiredLevel).toBe(60);
    expect(model.unique).toBe(false);
  });

  it('labels a known slot and falls back to the alias table for finger/trinket', () => {
    expect(itemTooltipModel(item({ slot: 'chest' }), NO_SOURCES).slotLabel).toBe('Chest');
    expect(itemTooltipModel(item({ slot: 'finger' }), NO_SOURCES).slotLabel).toBe('Finger');
    expect(itemTooltipModel(item({ slot: 'trinket' }), NO_SOURCES).slotLabel).toBe('Trinket');
  });

  it('renders stats in STAT_KEYS order, skipping zero and absent stats, armor excluded', () => {
    const model = itemTooltipModel(
      item({ stats: { crit: 14, strength: 18, stamina: 20, hit: 0 } }),
      NO_SOURCES,
    );
    // STAT_KEYS order: strength, agility, stamina, ..., crit, hit, ...
    expect(model.stats).toEqual(['+18 Strength', '+20 Stamina', '+14 Crit']);
  });

  it('shows armor only when the row carries some', () => {
    expect(itemTooltipModel(item({ armor: 400 }), NO_SOURCES).armor).toBe(400);
    expect(itemTooltipModel(item({ armor: 0 }), NO_SOURCES).armor).toBeNull();
  });

  it('builds a weapon line from damage, speed and dps, and omits it for non-weapons', () => {
    const sword = item({
      slot: 'main_hand',
      armor: 0,
      damage_min: 63,
      damage_max: 95,
      speed: 3.3,
      dps: 23.94,
      two_hand: false,
    });
    const model = itemTooltipModel(sword, NO_SOURCES);
    expect(model.weapon).toEqual({
      damageRange: '63 - 95 Damage',
      speed: 'Speed 3.30',
      dps: '(23.9 damage per second)',
    });
    expect(itemTooltipModel(item(), NO_SOURCES).weapon).toBeNull();
  });

  it('labels handedness and shields from fields the row actually carries, nothing invented', () => {
    const twoHander = item({
      slot: 'main_hand',
      armor: 0,
      damage_min: 1,
      damage_max: 2,
      speed: 1,
      two_hand: true,
    });
    const oneHander = item({
      slot: 'main_hand',
      armor: 0,
      damage_min: 1,
      damage_max: 2,
      speed: 1,
      two_hand: false,
    });
    const ranged = item({ slot: 'ranged', armor: 0, damage_min: 1, damage_max: 2, speed: 1 });
    const shield = item({ slot: 'off_hand', armor: 300 });
    const plainChest = item({ slot: 'chest', armor: 400 });
    expect(itemTooltipModel(twoHander, NO_SOURCES).typeLabel).toBe('Two-Handed Weapon');
    expect(itemTooltipModel(oneHander, NO_SOURCES).typeLabel).toBe('One-Handed Weapon');
    expect(itemTooltipModel(ranged, NO_SOURCES).typeLabel).toBe('Ranged Weapon');
    expect(itemTooltipModel(shield, NO_SOURCES).typeLabel).toBe('Shield');
    // No class/subclass id in the published row -- undefined, never a guessed material.
    expect(itemTooltipModel(plainChest, NO_SOURCES).typeLabel).toBeUndefined();
  });

  it('trims and collapses whitespace in effect text, and is null when the row has none', () => {
    expect(
      itemTooltipModel(item({ effect_text: '  Blasts a target for 85 Arcane damage. \r\n' }), NO_SOURCES)
        .effectText,
    ).toBe('Blasts a target for 85 Arcane damage.');
    expect(itemTooltipModel(item({ effect_text: '' }), NO_SOURCES).effectText).toBeNull();
    expect(itemTooltipModel(item(), NO_SOURCES).effectText).toBeNull();
  });

  it('names the set from set_id, and is null for an unset or unknown set', () => {
    const sets: ItemSet[] = [{ id: 550, name: 'Battlegear of Wrath', item_ids: [16963], bonuses: [] }];
    const sources: ItemTooltipSources = { loot: { sources: [] }, sets };
    expect(itemTooltipModel(item({ set_id: 550 }), sources).setName).toBe('Battlegear of Wrath');
    expect(itemTooltipModel(item({ set_id: null }), sources).setName).toBeNull();
    expect(itemTooltipModel(item({ set_id: 999 }), sources).setName).toBeNull();
  });

  it('names a raid boss, a dungeon boss, a vendor/world/zone source, a crafted profession, a reputation standing and a pvp rank', () => {
    const loot: LootFile = {
      sources: [
        {
          id: 'raid:molten-core',
          kind: 'raid',
          name: 'Molten Core',
          bosses: [{ id: 'raid:molten-core:12118', name: 'Ragnaros', items: [1] }],
        },
        { id: 'vendor:some-vendor', kind: 'vendor', name: 'Some Vendor', items: [2] },
        {
          id: 'crafted:blacksmithing',
          kind: 'crafted',
          name: 'Made item',
          profession: 'blacksmithing',
          items: [3],
        },
        {
          id: 'rep:argent-dawn:exalted',
          kind: 'rep',
          name: 'Argent Dawn',
          standing: 'exalted',
          items: [4],
        },
        { id: 'pvp:rank-10', kind: 'pvp', name: 'Warlord', rank: 10, items: [5] },
      ],
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 1 }), sources).sourceLines).toEqual(['Molten Core — Ragnaros']);
    expect(itemTooltipModel(item({ id: 2 }), sources).sourceLines).toEqual(['Some Vendor']);
    expect(itemTooltipModel(item({ id: 3 }), sources).sourceLines).toEqual(['Made item (blacksmithing)']);
    expect(itemTooltipModel(item({ id: 4 }), sources).sourceLines).toEqual(['Argent Dawn — Exalted']);
    expect(itemTooltipModel(item({ id: 5 }), sources).sourceLines).toEqual(['Warlord, rank 10']);
  });

  it('adds one line per quest reward, named with its faction', () => {
    const loot: LootFile = {
      sources: [],
      quests: {
        '744': [{ quest_id: 53, name: 'Sweet Amber', faction: 'alliance', min_level: 40, level: 44 }],
      },
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 744 }), sources).sourceLines).toEqual(['Sweet Amber (Alliance)']);
  });

  it('has no source lines when the item is not in the loot file', () => {
    expect(itemTooltipModel(item({ id: 42 }), NO_SOURCES).sourceLines).toEqual([]);
  });

  it('collapses a world or world_drop source to one "World drop" line, the two kinds treated identically', () => {
    const loot: LootFile = {
      sources: [
        { id: 'world:1', kind: 'world', name: 'Whatever this build calls it', items: [1] },
        // Not yet a real LootKind on LootSource's own type (the parallel lane's own
        // migration) -- cast the same way a build straight off that lane's pipeline would
        // arrive, so this proves the defensive read rather than assuming it away.
        {
          id: 'world:2',
          kind: 'world_drop' as LootFile['sources'][number]['kind'],
          name: 'Elsewhere',
          items: [2],
        },
      ],
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 1 }), sources).sourceLines).toEqual(['World drop']);
    expect(itemTooltipModel(item({ id: 2 }), sources).sourceLines).toEqual(['World drop']);
  });

  it('appends a world source’s level range when the row carries one, defensively (no typed field for it yet)', () => {
    const loot: LootFile = {
      sources: [
        {
          id: 'world_drop:20-30',
          kind: 'world_drop',
          name: 'World drop',
          items: [1],
          level_min: 20,
          level_max: 30,
        },
      ],
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 1 }), sources).sourceLines).toEqual([
      'World drop (BoE) · levels 20-30',
    ]);
  });

  it('puts a boss line’s known drop chance in the line as "(N%)"', () => {
    const loot: LootFile = {
      sources: [
        {
          id: 'raid:molten-core',
          kind: 'raid',
          name: 'Molten Core',
          bosses: [
            { id: 'raid:molten-core:12118', name: 'Ragnaros', items: [1], item_chances: { '1': 19.6 } },
          ],
        },
      ],
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 1 }), sources).sourceLines).toEqual(['Molten Core — Ragnaros (20%)']);
  });

  it('ranks bosses and quests ahead of vendor/crafted/rep/pvp lines, known chance descending within that group', () => {
    const loot: LootFile = {
      sources: [
        { id: 'vendor:some-vendor', kind: 'vendor', name: 'Some Vendor', items: [1] },
        {
          id: 'raid:molten-core',
          kind: 'raid',
          name: 'Molten Core',
          bosses: [{ id: 'raid:molten-core:a', name: 'Low Chance', items: [1], item_chances: { '1': 5 } }],
        },
      ],
      quests: { '1': [{ quest_id: 9, name: 'A Reward', faction: 'alliance', min_level: 1, level: 5 }] },
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 1 }), sources).sourceLines).toEqual([
      'Molten Core — Low Chance (5%)',
      'A Reward (Alliance)',
      'Some Vendor',
    ]);
  });

  it('orders two known-chance boss lines by chance descending', () => {
    const loot: LootFile = {
      sources: [
        {
          id: 'raid:molten-core',
          kind: 'raid',
          name: 'Molten Core',
          bosses: [{ id: 'raid:molten-core:a', name: 'Low Chance', items: [1], item_chances: { '1': 5 } }],
        },
        {
          id: 'raid:blackwing-lair',
          kind: 'raid',
          name: 'Blackwing Lair',
          bosses: [
            { id: 'raid:blackwing-lair:b', name: 'High Chance', items: [1], item_chances: { '1': 40 } },
          ],
        },
      ],
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 1 }), sources).sourceLines).toEqual([
      'Blackwing Lair — High Chance (40%)',
      'Molten Core — Low Chance (5%)',
    ]);
  });

  it('caps the source block at 3 lines and folds the rest into "and N more"', () => {
    const loot: LootFile = {
      sources: [
        { id: 'vendor:1', kind: 'vendor', name: 'Vendor One', items: [1] },
        { id: 'vendor:2', kind: 'vendor', name: 'Vendor Two', items: [1] },
        { id: 'vendor:3', kind: 'vendor', name: 'Vendor Three', items: [1] },
        { id: 'vendor:4', kind: 'vendor', name: 'Vendor Four', items: [1] },
        { id: 'vendor:5', kind: 'vendor', name: 'Vendor Five', items: [1] },
      ],
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 1 }), sources).sourceLines).toEqual([
      'Vendor One',
      'Vendor Two',
      'Vendor Three',
      'and 2 more',
    ]);
  });

  it('drops a duplicate line when two quest entries render identically', () => {
    // A neutral ("both") quest reward can appear as two separate entries in loot.json (one
    // recorded per side that can pick it up) that both humanise to the exact same line --
    // ItemTooltip.svelte keys its `{#each sourceLines as line (line)}` by the line's own
    // text, and a duplicate key there throws instead of rendering, so this must collapse to
    // one line rather than reach the component twice.
    const loot: LootFile = {
      sources: [],
      quests: {
        '19972': [
          { quest_id: 1, name: "Rare Fish - Keefer's Angelfish", faction: 'both', min_level: 1, level: 5 },
          { quest_id: 1, name: "Rare Fish - Keefer's Angelfish", faction: 'both', min_level: 1, level: 5 },
        ],
      },
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 19972 }), sources).sourceLines).toEqual([
      "Rare Fish - Keefer's Angelfish (Both)",
    ]);
  });
});

describe('statLines sign', () => {
  it('writes a negative stat as "-15 Parry", never "+-15"', () => {
    const model = itemTooltipModel(item({ id: 9, stats: { crit: 14, parry: -15 } }), {
      loot: { sources: [], quests: {}, factions: {} },
      sets: [],
    });
    expect(model.stats).toContain('+14 Crit');
    expect(model.stats).toContain('-15 Parry');
  });
});
