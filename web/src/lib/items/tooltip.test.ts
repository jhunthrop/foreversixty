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

  it('names a world_drop pool with and without a known level range', () => {
    const loot: LootFile = {
      sources: [
        {
          id: 'world_drop:18-25',
          kind: 'world_drop',
          name: 'World drop',
          items: [6],
          level_min: 18,
          level_max: 25,
        },
        { id: 'world_drop:unknown', kind: 'world_drop', name: 'World drop', items: [7] },
      ],
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 6 }), sources).sourceLines).toEqual([
      'World drop (BoE) · levels 18-25',
    ]);
    expect(itemTooltipModel(item({ id: 7 }), sources).sourceLines).toEqual(['World drop (BoE)']);
  });

  it('caps a long source list to the 3 highest-chance bosses plus an "and N more" line', () => {
    // The reported world-drop-pool defect's own shape: the same generic-pool item listed
    // as a named boss drop across many different dungeons -- one boss per source, since
    // sourceLinesFor takes the first matching boss PER SOURCE (its own established rule).
    const loot: LootFile = {
      sources: [
        {
          id: 'dungeon:a',
          kind: 'dungeon',
          name: 'Dungeon A',
          bosses: [{ id: 'a1', name: 'Low Chance Boss', items: [8], item_chances: { '8': 5 } }],
        },
        {
          id: 'dungeon:b',
          kind: 'dungeon',
          name: 'Dungeon B',
          bosses: [{ id: 'b1', name: 'High Chance Boss', items: [8], item_chances: { '8': 90 } }],
        },
        {
          id: 'dungeon:c',
          kind: 'dungeon',
          name: 'Dungeon C',
          bosses: [{ id: 'c1', name: 'Mid Chance Boss', items: [8], item_chances: { '8': 40 } }],
        },
        {
          id: 'dungeon:d',
          kind: 'dungeon',
          name: 'Dungeon D',
          bosses: [{ id: 'd1', name: 'No Chance Boss', items: [8] }],
        },
        {
          id: 'dungeon:e',
          kind: 'dungeon',
          name: 'Dungeon E',
          bosses: [{ id: 'e1', name: 'Another Low Boss', items: [8], item_chances: { '8': 10 } }],
        },
      ],
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 8 }), sources).sourceLines).toEqual([
      'Dungeon B — High Chance Boss',
      'Dungeon C — Mid Chance Boss',
      'Dungeon E — Another Low Boss',
      'and 2 more',
    ]);
  });

  it('does not reorder or cap a list of 3 or fewer sources', () => {
    const loot: LootFile = {
      sources: [
        { id: 'vendor:a', kind: 'vendor', name: 'Vendor A', items: [9] },
        { id: 'crafted:tailoring', kind: 'crafted', name: 'Made Item', profession: 'tailoring', items: [9] },
      ],
    };
    const sources: ItemTooltipSources = { loot, sets: [] };
    expect(itemTooltipModel(item({ id: 9 }), sources).sourceLines).toEqual([
      'Vendor A',
      'Made Item (tailoring)',
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
