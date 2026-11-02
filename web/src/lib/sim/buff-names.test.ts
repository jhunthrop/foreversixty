import { describe, expect, it } from 'vitest';
import { EMPTY_BUFF_NAMES, buffIcon, buffLabel, imbueLabel, type BuffNames } from './buff-names';

const names = {
  entries: {
    battle_shout: { name: 'Battle Shout', icon: 'ability_warrior_battleshout' },
    'main_hand_imbue:shadow_oil': { name: 'Shadow Oil', icon: 'inv_potion_21' },
  },
};

describe('buffLabel', () => {
  it('is the build’s own name when it has one', () => {
    expect(buffLabel('battle_shout', names)).toBe('Battle Shout');
    expect(buffLabel('main_hand_imbue:shadow_oil', names)).toBe('Main hand imbue: Shadow Oil');
  });

  it('humanises the id when the build has no table, rather than showing “Unknown”', () => {
    expect(buffLabel('flask_of_supreme_power', EMPTY_BUFF_NAMES)).toBe('Flask of supreme power');
    expect(buffLabel('flask_of_supreme_power', null)).toBe('Flask of supreme power');
  });

  it('humanises a qualified id by its own half, keeping the qualifier', () => {
    expect(buffLabel('off_hand_imbue:frost_oil', EMPTY_BUFF_NAMES)).toBe('Off hand imbue: Frost oil');
  });

  // Round 10: simbuffs.json now names each weapon imbue by its oil or stone, the same
  // name for both hands, so the row says which hand and field it fills.
  it("names a weapon imbue by the build's oil and keeps which hand it is", () => {
    const names = {
      entries: { 'off_hand_stone:wizard_oil': { name: 'Wizard Oil', icon: 'inv_potion_104' } },
    };
    expect(buffLabel('off_hand_stone:wizard_oil', names)).toBe('Off hand stone: Wizard Oil');
  });

  it('leaves a client item id alone: there is nothing to humanise', () => {
    expect(buffLabel('item:13452', EMPTY_BUFF_NAMES)).toBe('item:13452');
  });
});

describe('buffIcon', () => {
  it('is the build’s icon path when the table names one', () => {
    expect(buffIcon('1.15.9', 'battle_shout', names)).toBe(
      '/data/1.15.9/icons/ability_warrior_battleshout.webp',
    );
  });

  it('is null when there is no icon, so the row renders without one', () => {
    expect(buffIcon('1.15.9', 'thorns', names)).toBeNull();
    expect(buffIcon('1.15.9', 'thorns', null)).toBeNull();
  });
});

describe('imbueLabel', () => {
  it("is the build's item name with no hand, else the humanised item half", () => {
    expect(imbueLabel('main_hand_imbue:shadow_oil', names)).toBe('Shadow Oil');
    expect(imbueLabel('off_hand_stone:frost_oil', EMPTY_BUFF_NAMES)).toBe('Frost oil');
  });
});

// Round 11: the raid preset carries `mana_spring_totem:improved` (the Restoration shaman's
// totem). The table names the plain id only, so the improved form read "Mana spring totem:
// Improved" with no icon in the preset summary.
describe('an improved buff', () => {
  const improvedNames: BuffNames = {
    entries: { mana_spring_totem: { name: 'Mana Spring Totem', icon: 'spell_nature_manaregentotem' } },
  };

  it('is named by its own buff and the grade', () => {
    expect(buffLabel('mana_spring_totem:improved', improvedNames)).toBe('Mana Spring Totem (Improved)');
    expect(buffLabel('mana_spring_totem:improved', null)).toBe('Mana spring totem (Improved)');
  });

  it("draws its own buff's icon", () => {
    expect(buffIcon('1.60.1.70707', 'mana_spring_totem:improved', improvedNames)).toBe(
      buffIcon('1.60.1.70707', 'mana_spring_totem', improvedNames),
    );
    expect(buffIcon('1.60.1.70707', 'mana_spring_totem', improvedNames)).not.toBeNull();
  });
});
